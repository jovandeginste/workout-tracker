package cldrplural

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/vorlif/spreak/internal/util"
)

type Category int

const (
	Zero Category = iota
	One
	Two
	Few
	Many
	Other
)

var CategoryNames = map[Category]string{
	Zero:  "Zero",
	One:   "One",
	Two:   "Two",
	Few:   "Few",
	Many:  "Many",
	Other: "Other",
}

func (cat Category) String() string {
	if name, ok := CategoryNames[cat]; ok {
		return name
	}

	return "unknown"
}

type Operand int

const (
	OperandN Operand = iota // the absolute value of N.*
	OperandI                // the integer digits of N.*
	OperandV                // the number of visible fraction digits in N, with trailing zeros.*
	OperandW                // the number of visible fraction digits in N, without trailing zeros.*
	OperandF                // the visible fraction digits in N, with trailing zeros, expressed as an integer.*
	OperandT                // the visible fraction digits in N, without trailing zeros, expressed as an integer.*
	OperandC                // compact decimal exponent value: exponent of the power of 10 used in compact decimal formatting.
)

func (op Operand) String() string {
	if name, ok := operandNames[op]; ok {
		return name
	}

	return "unknown operand"
}

var operandNames = map[Operand]string{
	OperandN: "n",
	OperandI: "i",
	OperandV: "v",
	OperandW: "w",
	OperandF: "f",
	OperandT: "t",
	OperandC: "c",
}

var OperandMap = map[string]Operand{
	"n": OperandN,
	"i": OperandI,
	"v": OperandV,
	"w": OperandW,
	"f": OperandF,
	"t": OperandT,
	"c": OperandC,
	"e": OperandC,
}

type FormFunc func(ops *Operands) Category

type RuleSet struct {
	Categories []Category
	FormFunc   FormFunc
}

func (rs *RuleSet) Evaluate(a interface{}) Category {
	ops, err := NewOperands(a)
	if err != nil {
		ops = newOperandsInt(0)
	}
	return rs.FormFunc(ops)
}

// The Operands are numeric values corresponding to features of the source number.
type Operands struct {
	N float64
	I int64
	V int64
	W int64
	F int64
	T int64
	C int64
}

func MustNewOperands(a interface{}) *Operands {
	ops, err := NewOperands(a)
	if err != nil {
		panic(err)
	}
	return ops
}

// NewOperands converts the representation of a float value into the appropriate Operands.
func NewOperands(a interface{}) (*Operands, error) {
	a = util.Indirect(a)
	if a == nil {
		return nil, errors.New("operands value is nil")
	}

	switch v := a.(type) {
	case string:
		return newOperandsString(v)
	case int64:
		return newOperandsInt(v), nil
	case int:
		return newOperandsInt(int64(v)), nil
	case float32:
		return newOperandsString(fmt.Sprintf("%v", v))
	case float64:
		return newOperandsString(fmt.Sprintf("%v", v))
	default:
		num, err := util.ToNumber(v)
		if err != nil {
			return nil, err
		}
		return newOperandsString(fmt.Sprintf("%v", num))
	}
}

func newOperandsInt(i int64) *Operands {
	if i < 0 {
		i = -i
	}
	return &Operands{float64(i), i, 0, 0, 0, 0, 0}
}

func newOperandsString(raw string) (*Operands, error) {
	op := &Operands{}

	if cIdx := strings.Index(raw, "c"); cIdx >= 0 {
		c, err := strconv.Atoi(raw[cIdx+1:])
		if err != nil {
			return nil, err
		}
		op.C = int64(c)
		raw = shiftDecimalPoint(raw[:cIdx], c)
	}

	src, errP := strconv.ParseFloat(raw, 64)
	if errP != nil {
		return nil, errP
	}

	op.N = math.Abs(src)
	op.I = int64(src)

	if pointIdx := strings.Index(raw, "."); pointIdx >= 0 {
		fractionDigits := raw[pointIdx+1:]
		if fractionDigits != "" {
			op.V = int64(len(fractionDigits))
			i, err := strconv.ParseInt(fractionDigits, 10, 64)
			if err != nil {
				return nil, err
			}
			op.F = i
		}

		withoutZeros := strings.TrimRight(fractionDigits, "0")
		if withoutZeros != "" {
			op.W = int64(len(withoutZeros))
			i, err := strconv.ParseInt(withoutZeros, 10, 64)
			if err != nil {
				return nil, err
			}

			op.T = i
		}
	}

	return op, nil
}

func shiftDecimalPoint(raw string, c int) string {
	var s strings.Builder

	shift := false
	for _, r := range raw {
		if r == '.' {
			shift = true
			continue
		}
		if c == 0 && shift {
			s.WriteRune('.')
			shift = false
		}
		if shift {
			c--
		}
		s.WriteRune(r)
	}

	for i := 0; i < c; i++ {
		s.WriteRune('0')
	}
	return s.String()
}
