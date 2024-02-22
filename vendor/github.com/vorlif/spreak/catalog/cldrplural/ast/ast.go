package ast

type Rule struct {
	Root    Node
	Samples []string
}

type Node interface {
	Type() Token
}

type OperandExpr struct {
	Operand string
}

type ValueExpr struct {
	Value int64
}

type RangeExpr struct {
	From int64
	To   int64
}

type RangeListExpr struct {
	X Node // (range | value)
	Y Node // (',' range_list)*  Optional!
}

type ModuloExpr struct {
	Op    *OperandExpr
	Value int64
}

type BinaryExpr struct {
	X  Node
	Op Token // and, or
	Y  Node
}

type InRelationExpr struct {
	X  Node           // operand (('mod' | '%') value)?
	Op Token          // ==, !=
	Y  *RangeListExpr //
}

func (OperandExpr) Type() Token      { return Operand }
func (ValueExpr) Type() Token        { return Value }
func (RangeExpr) Type() Token        { return ValueRange }
func (RangeListExpr) Type() Token    { return RangeList }
func (ModuloExpr) Type() Token       { return Remainder }
func (e BinaryExpr) Type() Token     { return e.Op }
func (e InRelationExpr) Type() Token { return e.Op }
