package expr

import (
	"fmt"

	"github.com/gsanchezgavier/craftinginterpreters/src/token"
)

// expression     → literal
//                | unary
//                | binary
//                | grouping ;

// literal        → NUMBER | STRING | "true" | "false" | "nil" ;
// grouping       → "(" expression ")" ;
// unary          → ( "-" | "!" ) expression ;
// binary         → expression operator expression ;
// operator       → "==" | "!=" | "<" | "<=" | ">" | ">="
//                | "+"  | "-"  | "*" | "/" ;

type Expr interface {
	Accept(v Visitor)
}

type Visitor interface {
	VisitBinary(b Binary)
	VisitGrouping(g Grouping)
	VisitLiteral(l Literal)
	VisitUnary(u Unary)
}

type Binary struct {
	left     Expr
	operator token.Token
	right    Expr
}

func NewBinary(left Expr, operator token.Token, right Expr) Binary {
	return Binary{
		left:     left,
		operator: operator,
		right:    right,
	}
}

func (b Binary) Accept(v Visitor) {
	v.VisitBinary(b)
}

type Literal struct {
	value any
}

func NewLiteral(value any) Literal {
	return Literal{
		value: value,
	}
}

func (l Literal) Accept(v Visitor) {
	v.VisitLiteral(l)
}

type Grouping struct {
	expression Expr
}

func NewGrouping(expression Expr) Grouping {
	return Grouping{
		expression: expression,
	}
}
func (g Grouping) Accept(v Visitor) {
	v.VisitGrouping(g)
}

type Unary struct {
	operator token.Token
	right    Expr
}

func NewUnary(operator token.Token, right Expr) Unary {
	return Unary{
		operator: operator,
		right:    right,
	}
}
func (u Unary) Accept(v Visitor) {
	v.VisitUnary(u)
}

// Generic version
// type Expr[T any] interface {
// 	Accept(v Visitor[T]) T
// }

// type Visitor[T any] interface {
// 	VisitBinary(b Binary[T])
// 	VisitGrouping(g Grouping[T])
// 	VisitLiteral(l Literal[T])
// 	VisitUnary(u Unary[T])
// }

// type Binary[T any] struct {
// 	left     Expr[T]
// 	operator token.Token
// 	right    Expr[T]
// }

// func (b Binary[T]) Accept(v Visitor[T]) {
// 	v.VisitBinary(b)
// }

// type Literal[T any] struct {
// 	value any
// }

// func (l Literal[T]) Accept(v Visitor[T]) {
// 	v.VisitLiteral(l)
// }

// type Grouping[T any] struct {
// 	expression Expr[T]
// }

// func (g Grouping[T]) Accept(v Visitor[T]) {
// 	v.VisitGrouping(g)
// }

// type Unary[T any] struct {
// 	operator token.Token
// 	right    Expr[T]
// }

// func (u Unary[T]) Accept(v Visitor[T]) {
// 	v.VisitUnary(u)
// }

type Printer struct{}

func (p Printer) Print(expr Expr) {
	expr.Accept(p)
}

func (p Printer) VisitBinary(b Binary) {
	p.parenthesize(b.operator.Lexeme, b.left, b.right)
}
func (p Printer) VisitGrouping(g Grouping) {
	p.parenthesize("group", g.expression)
}
func (p Printer) VisitLiteral(l Literal) {
	if l.value != nil {
		fmt.Printf("%v", l.value)
	} else {
		fmt.Print("nil")
	}
}
func (p Printer) VisitUnary(u Unary) {
	p.parenthesize(u.operator.Lexeme, u.right)
}

func (p Printer) parenthesize(name string, expr ...Expr) {
	fmt.Print("(")
	fmt.Print(name)
	for _, expr := range expr {
		fmt.Print(" ")
		expr.Accept(p)
	}
	fmt.Print(")")
}
