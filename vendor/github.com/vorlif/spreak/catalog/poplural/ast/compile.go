package ast

import (
	"fmt"
	"strconv"
	"strings"
)

func CompileToString(forms *Forms) string {
	return fmt.Sprintf("nplurals=%d; plural=%s;", forms.NPlurals, compileNode(forms.Root))
}

func compileNode(node Node) string {
	switch v := node.(type) {
	case *ValueExpr:
		return strconv.FormatInt(v.Value, 10)
	case *OperandExpr:
		return "n"
	case *QuestionMarkExpr:
		cond := compileNode(v.Cond)
		then := compileNode(v.T)
		other := compileNode(v.F)

		if len(cond) > 1 && !(strings.HasPrefix(cond, "(") && strings.HasSuffix(cond, ")")) {
			cond = "(" + cond + ")"
		}
		if len(then) > 1 {
			then = "(" + then + ")"
		}
		if len(other) > 1 {
			other = "(" + other + ")"
		}

		return fmt.Sprintf("%s ? %s : %s", cond, then, other)
	case *BinaryExpr:
		switch v.Type() {
		case LogicalAnd:
			return compileNode(v.X) + " && " + compileNode(v.Y)
		case LogicalOr:
			return fmt.Sprintf("(%s || %s)", compileNode(v.X), compileNode(v.Y))
		case Equal:
			return compileNode(v.X) + " == " + compileNode(v.Y)
		case NotEqual:
			return compileNode(v.X) + " != " + compileNode(v.Y)
		case Greater:
			return compileNode(v.X) + " > " + compileNode(v.Y)
		case GreaterOrEqual:
			return compileNode(v.X) + " >= " + compileNode(v.Y)
		case Less:
			return compileNode(v.X) + " < " + compileNode(v.Y)
		case LessOrEqual:
			return compileNode(v.X) + " <= " + compileNode(v.Y)
		case Reminder:
			return compileNode(v.X) + " % " + compileNode(v.Y)
		}
	}
	return ""
}
