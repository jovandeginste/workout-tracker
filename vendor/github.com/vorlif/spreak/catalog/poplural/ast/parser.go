package ast

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Forms struct {
	NPlurals int
	Root     Node
}

type parser struct {
	s           *scanner
	lastToken   Token  // last read Token
	lastLiteral string // last read literal
	n           int    // buffer size (max=1)
}

func MustParse(rule string) *Forms {
	f, err := Parse(rule)
	if err != nil {
		panic(err)
	}

	return f
}

func Parse(rule string) (*Forms, error) {
	p := &parser{
		s: newScanner(strings.NewReader(rule)),
	}
	return p.Parse()
}

func (p *parser) Parse() (*Forms, error) {
	forms := &Forms{}

	if tok, lit := p.scanIgnoreWhitespace(); tok != nPlurals {
		return nil, fmt.Errorf("po parser: found %q, expected 'nplurals'", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != assign {
		return nil, fmt.Errorf("po parser: found %q, expected '=' after nplurals", lit)
	}

	//nolint:revive
	if tok, lit := p.scanIgnoreWhitespace(); tok != Value {
		return nil, fmt.Errorf("po parser: found %q, expected '=' after nplurals", lit)
	} else {
		forms.NPlurals, _ = strconv.Atoi(lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != semicolon {
		return nil, fmt.Errorf("po parser: found %q, expected ';' after 'nplurals=%d'", lit, forms.NPlurals)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != plural {
		return nil, fmt.Errorf("po parser: found %q, expected 'plural' after 'nplurals=%d; '", lit, forms.NPlurals)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != assign {
		return nil, fmt.Errorf("po parser: found %q, expected '=' after 'nplurals=%d; plural'", lit, forms.NPlurals)
	}

	if errScan := p.scanNext(); errScan != nil {
		return nil, errScan
	}

	n, err := p.expression()
	if err != nil {
		return nil, err
	}

	forms.Root = n

	if p.lastToken != semicolon {
		return nil, fmt.Errorf("po parser: found %q, expected ';' at end", p.lastLiteral)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != eof {
		return nil, fmt.Errorf("po parser: found %q, expected end", lit)
	}

	return forms, nil

}

func (p *parser) expression() (Node, error) {
	n, err := p.logicalOrExpression()
	if err != nil {
		return nil, err
	}

	if p.lastToken == Question {
		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}

		condTrue, errCt := p.expression()
		if errCt != nil {
			return nil, errCt
		}

		if p.lastToken != colon {
			return nil, fmt.Errorf("po parser: found %q, expected \":\"", p.lastLiteral)
		}

		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}

		condFalse, errCf := p.expression()
		if errCf != nil {
			return nil, errCf
		}

		qmNode := &QuestionMarkExpr{
			Cond: n,
			T:    condTrue,
			F:    condFalse,
		}
		return qmNode, nil
	}

	return n, nil
}

func (p *parser) logicalOrExpression() (Node, error) {
	ln, err := p.logicalAndExpression()
	if err != nil {
		return nil, err
	}

	if p.lastToken == LogicalOr {
		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}

		rn, errRn := p.logicalOrExpression() // right
		if errRn != nil {
			return nil, errRn
		}

		orNode := &BinaryExpr{
			X:  ln,
			Op: LogicalOr,
			Y:  rn,
		}

		if rn.Type() == LogicalOr {
			rnNode := rn.(*BinaryExpr)
			orNode.Y = rnNode.X
			rnNode.X = orNode
			return rnNode, nil
		}

		return orNode, nil
	}

	return ln, nil
}

func (p *parser) logicalAndExpression() (Node, error) {
	ln, err := p.equalityExpression() // left
	if err != nil {
		return nil, err
	}

	if p.lastToken == LogicalAnd {
		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}

		rn, errRn := p.logicalAndExpression()
		if errRn != nil {
			return nil, errRn
		}

		andNode := &BinaryExpr{
			X:  ln,
			Op: LogicalAnd,
			Y:  rn,
		}

		if rn.Type() == LogicalAnd {
			rnNode := rn.(*BinaryExpr)
			andNode.Y = rnNode.X
			rnNode.X = andNode
			return rnNode, nil
		}

		return andNode, nil
	}

	return ln, nil
}

func (p *parser) equalityExpression() (Node, error) {
	n, err := p.relationalExpression()
	if err != nil {
		return nil, err
	}

	if p.lastToken == Equal || p.lastToken == NotEqual {
		compareNode := &BinaryExpr{Op: p.lastToken}

		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}

		re, errRe := p.relationalExpression()
		if errRe != nil {
			return nil, errRe
		}

		compareNode.X = n
		compareNode.Y = re
		return compareNode, nil
	}

	return n, nil
}

func (p *parser) relationalExpression() (Node, error) {
	n, err := p.multiplicativeExpression()
	if err != nil {
		return nil, err
	}

	if p.lastToken == Greater || p.lastToken == Less || p.lastToken == GreaterOrEqual || p.lastToken == LessOrEqual {
		compareNode := &BinaryExpr{Op: p.lastToken}

		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}

		me, errMe := p.multiplicativeExpression()
		if errMe != nil {
			return nil, errMe
		}

		compareNode.X = n
		compareNode.Y = me
		return compareNode, nil
	}

	return n, nil
}

func (p *parser) multiplicativeExpression() (Node, error) {
	n, err := p.pmExpression()
	if err != nil {
		return nil, err
	}

	if p.lastToken == Reminder {
		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}

		pm, errPm := p.pmExpression()
		if errPm != nil {
			return nil, errPm
		}

		modNode := &BinaryExpr{
			X:  n,
			Op: Reminder,
			Y:  pm,
		}
		return modNode, nil
	}

	return n, nil
}

func (p *parser) pmExpression() (Node, error) {
	switch p.lastToken {
	case Operand:
		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}
		return &OperandExpr{}, nil
	case Value:
		value, _ := strconv.ParseInt(p.lastLiteral, 10, 64)
		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}
		return &ValueExpr{Value: value}, nil
	case leftBracket:
		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}

		exprNode, errEp := p.expression()
		if errEp != nil {
			return nil, errEp
		}

		if p.lastToken != rightBracket {
			return nil, fmt.Errorf("found %q, expected )", p.lastLiteral)
		}

		if errScan := p.scanNext(); errScan != nil {
			return nil, errScan
		}

		return exprNode, nil
	default:
		return nil, fmt.Errorf("found %q, expected something other", p.lastLiteral)
	}
}

func (p *parser) scanNext() error {
	if tok, _ := p.scanIgnoreWhitespace(); tok == eof || tok == failure {
		return errors.New("eof reached without result")
	}
	return nil
}

func (p *parser) scan() (tok Token, lit string) {
	if p.n != 0 {
		p.n = 0
		return p.lastToken, p.lastLiteral
	}

	tok, lit = p.s.scan()

	p.lastToken, p.lastLiteral = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == whitespace {
		tok, lit = p.scan()
	}
	return
}
