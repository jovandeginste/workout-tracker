package ast

type Node interface {
	Type() Token
}

type OperandExpr struct{}

type ValueExpr struct {
	Value int64
}

type BinaryExpr struct {
	X  Node
	Op Token // and, or, ==, !=
	Y  Node
}

type QuestionMarkExpr struct {
	Cond Node
	T    Node
	F    Node
}

func (OperandExpr) Type() Token      { return Operand }
func (ValueExpr) Type() Token        { return Value }
func (QuestionMarkExpr) Type() Token { return Question }
func (e BinaryExpr) Type() Token     { return e.Op }
