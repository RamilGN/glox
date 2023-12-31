// Code generated by "go run tools/ast/main.go"; DO NOT EDIT.

package main

type Visitor interface {
	VisitBinary(*Binary) (any, error)
	VisitGrouping(*Grouping) (any, error)
	VisitLiteral(*Literal) (any, error)
	VisitUnary(*Unary) (any, error)
}

type Expr interface {
	Accept(Visitor) (any, error)
}

type Binary struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func NewBinary(
	Left Expr,
	Operator Token,
	Right Expr,
) *Binary {

	return &Binary{
		Left:     Left,
		Operator: Operator,
		Right:    Right,
	}
}

func (e *Binary) Accept(visitor Visitor) (any, error) {
	return visitor.VisitBinary(e)
}

type Grouping struct {
	Expression Expr
}

func NewGrouping(
	Expression Expr,
) *Grouping {

	return &Grouping{
		Expression: Expression,
	}
}

func (e *Grouping) Accept(visitor Visitor) (any, error) {
	return visitor.VisitGrouping(e)
}

type Literal struct {
	Object any
}

func NewLiteral(
	Object any,
) *Literal {

	return &Literal{
		Object: Object,
	}
}

func (e *Literal) Accept(visitor Visitor) (any, error) {
	return visitor.VisitLiteral(e)
}

type Unary struct {
	Operator Token
	Right    Expr
}

func NewUnary(
	Operator Token,
	Right Expr,
) *Unary {

	return &Unary{
		Operator: Operator,
		Right:    Right,
	}
}

func (e *Unary) Accept(visitor Visitor) (any, error) {
	return visitor.VisitUnary(e)
}
