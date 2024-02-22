package ast

import (
	"fmt"
	"strconv"
	"strings"
)

// The parser can translate a single rule into an abstract
// tree of nodes. It implements the following grammar::
//
//	   condition     = and_condition ('or' and_condition)*
//	                   ('@integer' samples)?
//	                   ('@decimal' samples)?
//	   and_condition = relation ('and' relation)*
//	   relation      = expr ('=' | '!=') range_list
//	   expr          = operand ('%' value)?
//	   operand       = 'n' | 'i' | 'f' | 't' | 'v' | 'w'
//	   range_list    = (range | value) (',' range_list)*
//	   value         = digit+
//	   digit         = 0|1|2|3|4|5|6|7|8|9
//	   range         = value'..'value
//	   samples       = sampleRange (',' sampleRange)* (',' ('…'|'...'))?
//	   sampleRange   = decimalValue '~' decimalValue
//	   decimalValue  = value ('.' value)?
//
//	- Whitespace can occur between or around any of the above tokens.
//	- Rules should be mutually exclusive; for a given numeric value, only one
//	  rule should apply (i.e. the condition should only be true for one of
//	  the plural rule elements).
//	- The in and within relations can take comma-separated lists, such as:
//	  'n in 3,5,7..15'.
//	- Samples are ignored.
//	  The translator parses the expression on instanciation into an attribute
//	  called `ast`.
//
// See: http://unicode.org/reports/tr35/tr35-numbers.html#51-plural-rules-syntax
type parser struct {
	s           *scanner
	lastToken   Token  // last read token
	lastLiteral string // last read literal
	n           int    // buffer size (max=1)
}

func MustParse(rawRule string) *Rule {
	r, err := Parse(rawRule)
	if err != nil {
		panic(err)
	}
	return r
}

func Parse(rawRule string) (*Rule, error) {
	p := &parser{s: newScanner(rawRule)}

	a, err := p.Parse()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (p *parser) Parse() (*Rule, error) {
	if errScan := p.scanNext(); errScan != nil {
		return nil, errScan
	}

	rule := &Rule{}
	if p.lastToken != sample {
		n, err := p.condition()
		if err != nil {
			return nil, err
		}
		rule.Root = n
	}

	samples, errS := p.samples()
	if errS != nil {
		return nil, errS
	}
	rule.Samples = samples

	return rule, nil
}

func (p *parser) condition() (Node, error) {
	ln, err := p.andCondition()
	if err != nil {
		return nil, err
	}

	if p.lastToken == Or {
		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}

		rn, errRn := p.condition() // right
		if errRn != nil {
			return nil, errRn
		}

		orNode := &BinaryExpr{
			X:  ln,
			Op: Or,
			Y:  rn,
		}
		return orNode, nil
	}

	return ln, nil
}

func (p *parser) andCondition() (Node, error) {
	ln, err := p.relation() // left
	if err != nil {
		return nil, err
	}

	if p.lastToken == And {
		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}

		rn, errRn := p.andCondition()
		if errRn != nil {
			return nil, errRn
		}

		andNode := &BinaryExpr{
			X:  ln,
			Op: And,
			Y:  rn,
		}
		return andNode, nil
	}

	return ln, nil
}

func (p *parser) relation() (Node, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}

	//  require ('=' | '!=')
	if p.lastToken != Equal && p.lastToken != NotEqual {
		return nil, fmt.Errorf("expected value (\\d+), got %s", p.lastLiteral)
	}

	compareNode := &InRelationExpr{Op: p.lastToken}

	if errScan := p.scanNext(); errScan != nil {
		return nil, errScan
	}

	rangeLis, errList := p.rangeList()
	if errList != nil {
		return nil, errList
	}

	compareNode.X = expr
	compareNode.Y = rangeLis
	return compareNode, nil
}

func (p *parser) expression() (Node, error) {
	if p.lastToken != Operand {
		return nil, fmt.Errorf("expected operand, got %q", p.lastLiteral)
	}

	operandNode := &OperandExpr{Operand: p.lastLiteral}

	if errScan := p.scanNext(); errScan != nil {
		return nil, errScan
	}

	if p.lastToken == Remainder {
		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}

		valueNode, errValue := p.value()
		if errValue != nil {
			return nil, errValue
		}

		reminderNode := &ModuloExpr{
			Op:    operandNode,
			Value: valueNode.Value,
		}
		return reminderNode, nil
	}

	return operandNode, nil
}

// value = digit+.
func (p *parser) value() (*ValueExpr, error) {
	if p.lastToken != Value {
		return nil, fmt.Errorf("expected value (\\d+), got %q", p.lastLiteral)
	}

	value, _ := strconv.ParseInt(p.lastLiteral, 10, 64)

	if errScan := p.scanNext(); errScan != nil {
		return nil, errScan
	}

	return &ValueExpr{Value: value}, nil
}

// range_list = (range | value) (',' range_list)*.
func (p *parser) rangeList() (*RangeListExpr, error) {
	listNode := &RangeListExpr{}

	if p.lastToken == Value {
		n, err := p.value()
		if err != nil {
			return nil, err
		}
		listNode.X = n
	} else if p.lastToken == ValueRange {
		n, err := p.rangeExpression()
		if err != nil {
			return nil, err
		}
		listNode.X = n
	} else {
		return nil, fmt.Errorf("expected value range or value, got %q", p.lastLiteral)
	}

	if p.lastToken == ValueRange || p.lastToken == Value {
		right, errR := p.rangeList()
		if errR != nil {
			return nil, errR
		}

		listNode.Y = right
	}

	return listNode, nil
}

// range = value'..'value.
func (p *parser) rangeExpression() (*RangeExpr, error) {
	if p.lastToken != ValueRange {
		return nil, fmt.Errorf("expected range expression (\\d+..\\d+), got %q", p.lastLiteral)
	}

	valueStr := strings.Split(p.lastLiteral, "..")
	if len(valueStr) != 2 {
		return nil, fmt.Errorf("the value range has an invalid formatting %s", p.lastLiteral)
	}
	left, errL := strconv.ParseInt(valueStr[0], 10, 64)
	if errL != nil {
		return nil, fmt.Errorf("the value range has an invalid start %s", p.lastLiteral)
	}

	right, errR := strconv.ParseInt(valueStr[1], 10, 64)
	if errR != nil {
		return nil, fmt.Errorf("the value range has an invalid end %s", p.lastLiteral)
	}

	if right < left {
		return nil, fmt.Errorf("the end of the value range is larger than the beginning: %s", p.lastLiteral)
	}

	rangeNode := &RangeExpr{
		From: left,
		To:   right,
	}

	if errScan := p.scanNext(); errScan != nil {
		return nil, errScan
	}

	return rangeNode, nil
}

func (p *parser) samples() ([]string, error) {
	var samples []string

	if p.lastToken != eof && p.lastToken != sample {
		return nil, fmt.Errorf("end of the rule expected got %q", p.lastLiteral)
	} else if p.lastToken == eof {
		return samples, nil
	}

	for p.lastToken != eof {
		switch p.lastToken {
		case sample: // "@decimal", "@integer"
			if errScan := p.scanNext(); errScan != nil {
				return nil, errScan
			}
			continue
		case sampleValue, Value:
			samples = append(samples, p.lastLiteral)
		case sampleRange:
			rangeValues := parseSampleRange(p.lastLiteral)
			samples = append(samples, rangeValues...)
		default:
			return nil, fmt.Errorf("expected sample, got %q", p.lastLiteral)
		}

		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}
	}

	return samples, nil
}

func (p *parser) scanNext() error {
	if tok, _ := p.scan(); tok == unknown {
		return fmt.Errorf("unknown token readed %s", p.lastLiteral)
	}
	return nil
}

func (p *parser) scan() (tok Token, lit string) {
	if p.n != 0 {
		p.n = 0
		return p.lastToken, p.lastLiteral
	}

	tok, lit = p.s.Scan()

	p.lastToken, p.lastLiteral = tok, lit

	return
}

// sampleRange = sampleValue ('~' sampleValue)?
// sampleValue = value ('.' digit+)? ([ce] digitPos digit+)?
func parseSampleRange(sampleRange string) []string {
	var sampleValues []string // sampleValue, sampleValue, sampleValue, ...
	if parts := strings.Split(sampleRange, "~"); len(parts) == 2 {
		for ex := parts[0]; ; ex = increment(ex) {
			sampleValues = append(sampleValues, ex)
			if ex == parts[1] {
				break
			}
		}
	} else {
		sampleValues = append(sampleValues, parts...)
	}
	return sampleValues
}

func increment(dec string) string {
	runes := []rune(dec)
	carry := true
	for i := len(runes) - 1; carry && i >= 0; i-- {
		switch runes[i] {
		case '.':
			continue
		case '9':
			runes[i] = '0'
		default:
			runes[i]++
			carry = false
		}
	}
	if carry {
		runes = append([]rune{'1'}, runes...)
	}
	return string(runes)
}
