package ast

type Visitor interface {
	Visit(node Node) (w Visitor)
}

func Walk(v Visitor, node Node) {
	if v = v.Visit(node); v == nil {
		return
	}

	switch n := node.(type) {
	case *RangeListExpr:
		Walk(v, n.X)
		if n.Y != nil {
			Walk(v, n.Y)
		}
	case *ModuloExpr:
		Walk(v, n.Op)
	case *BinaryExpr:
		Walk(v, n.X)
		Walk(v, n.Y)
	case *InRelationExpr:
		Walk(v, n.X)
		Walk(v, n.Y)
	}

	v.Visit(nil)
}

type inspector func(Node) bool

func (f inspector) Visit(node Node) Visitor {
	if f(node) {
		return f
	}
	return nil
}

func Inspect(node Node, f func(Node) bool) {
	Walk(inspector(f), node)
}
