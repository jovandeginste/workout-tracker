package poplural

import (
	"math"

	"github.com/vorlif/spreak/catalog/poplural/ast"
	"github.com/vorlif/spreak/internal/util"
)

type Form struct {
	NPlurals int
	FormFunc func(n int64) int
}

func (f *Form) Evaluate(a interface{}) int {
	num, err := util.ToNumber(a)
	if err != nil {
		return 0
	}

	num = math.RoundToEven(math.Abs(num))
	return f.FormFunc(int64(num))
}

func MustParse(rule string) *Form {
	form, err := Parse(rule)
	if err != nil {
		panic(err)
	}
	return form
}

// Parse parses a plural forms header and returns a function to evaluate this header.
// If for a header there is already a predefined function, this function will be returned.
func Parse(rule string) (*Form, error) {
	parsed, err := ast.Parse(rule)
	if err != nil {
		return nil, err
	}

	// Use of built-in functions, if available
	compiledRaw := ast.CompileToString(parsed)
	if form, ok := rawToBuiltIn[compiledRaw]; ok {
		return form, nil
	}

	f := &Form{
		NPlurals: parsed.NPlurals,
		FormFunc: generateFormFunc(parsed),
	}
	return f, nil
}

func generateFormFunc(forms *ast.Forms) func(n int64) int {
	if forms.Root == nil {
		return func(n int64) int { return 0 }
	}

	return func(n int64) int {
		return int(evaluateNode(forms.Root, n))
	}
}

func evaluateNode(node ast.Node, num int64) int64 {
	var conditionTrue bool

	switch v := node.(type) {
	case *ast.ValueExpr:
		return v.Value
	case *ast.OperandExpr:
		return num
	case *ast.QuestionMarkExpr:
		if evaluateNode(v.Cond, num) == 1 {
			return evaluateNode(v.T, num)
		}
		return evaluateNode(v.F, num)
	case *ast.BinaryExpr:
		switch v.Type() {
		case ast.LogicalAnd:
			conditionTrue = evaluateNode(v.X, num) == 1 && evaluateNode(v.Y, num) == 1
		case ast.LogicalOr:
			conditionTrue = evaluateNode(v.X, num) == 1 || evaluateNode(v.Y, num) == 1
		case ast.Equal:
			conditionTrue = evaluateNode(v.X, num) == evaluateNode(v.Y, num)
		case ast.NotEqual:
			conditionTrue = evaluateNode(v.X, num) != evaluateNode(v.Y, num)
		case ast.Greater:
			conditionTrue = evaluateNode(v.X, num) > evaluateNode(v.Y, num)
		case ast.GreaterOrEqual:
			conditionTrue = evaluateNode(v.X, num) >= evaluateNode(v.Y, num)
		case ast.Less:
			conditionTrue = evaluateNode(v.X, num) < evaluateNode(v.Y, num)
		case ast.LessOrEqual:
			conditionTrue = evaluateNode(v.X, num) <= evaluateNode(v.Y, num)
		case ast.Reminder:
			rightVal := evaluateNode(v.Y, num)
			if rightVal == 0 {
				return 0
			}
			return evaluateNode(v.X, num) % rightVal
		}
	default:
		return 0
	}

	if conditionTrue {
		return 1
	}

	return 0
}
