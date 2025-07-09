package expr

import (
	"fmt"
	"strings"

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
	Accept(v Visitor) any
}

type Visitor interface {
	VisitBinary(b Binary) any
	VisitGrouping(g Grouping) any
	VisitLiteral(l Literal) any
	VisitUnary(u Unary) any
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func NewBinary(left Expr, operator token.Token, right Expr) Binary {
	return Binary{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (b Binary) Accept(v Visitor) any {
	return v.VisitBinary(b)
}

type Literal struct {
	Value any
}

func NewLiteral(value any) Literal {
	return Literal{
		Value: value,
	}
}

func (l Literal) Accept(v Visitor) any {
	return v.VisitLiteral(l)
}

type Grouping struct {
	expression Expr
}

func NewGrouping(expression Expr) Grouping {
	return Grouping{
		expression: expression,
	}
}
func (g Grouping) Accept(v Visitor) any {
	return v.VisitGrouping(g)
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

func NewUnary(operator token.Token, right Expr) Unary {
	return Unary{
		Operator: operator,
		Right:    right,
	}
}
func (u Unary) Accept(v Visitor) any {
	return v.VisitUnary(u)
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

func (p Printer) Print(expr Expr) any {
	return expr.Accept(p)
}

func (p Printer) VisitBinary(b Binary) any {
	return p.parenthesize(b.Operator.Lexeme, b.Left, b.Right)
}
func (p Printer) VisitGrouping(g Grouping) any {
	return p.parenthesize("group", g.expression)
}
func (p Printer) VisitLiteral(l Literal) any {
	if l.Value != nil {
		return fmt.Sprintf("%v", l.Value)
	} else {
		return "nil"
	}
}
func (p Printer) VisitUnary(u Unary) any {
	return p.parenthesize(u.Operator.Lexeme, u.Right)
}

func (p Printer) parenthesize(name string, expr ...Expr) any {
	var buf strings.Builder
	buf.WriteString("(")
	buf.WriteString(name)
	for _, expr := range expr {
		buf.WriteString(" ")
		buf.WriteString(expr.Accept(p).(string))
	}
	buf.WriteString(")")
	return buf.String()
}
