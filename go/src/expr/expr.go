package expr

import "github.com/gsanchezgavier/craftinginterpreters/src/token"

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

type Expr[T any] interface {
	Accept(v Visitor[T]) T
}

type Visitor[T any] interface {
	VisitBinary(b Binary[T])
	VisitGrouping(g Grouping[T])
	VisitLiteral(l Literal[T])
	VisitUnary(u Unary[T])
}

type Binary[T any] struct {
	left     Expr[T]
	operator token.Token
	right    Expr[T]
}

func (b Binary[T]) Accept(v Visitor[T]) {
	v.VisitBinary(b)
}

type Literal[T any] struct {
	value any
}

func (l Literal[T]) Accept(v Visitor[T]) {
	v.VisitLiteral(l)
}

type Grouping[T any] struct {
	expression Expr[T]
}

func (g Grouping[T]) Accept(v Visitor[T]) {
	v.VisitGrouping(g)
}

type Unary[T any] struct {
	operator token.Token
	right    Expr[T]
}

func (u Unary[T]) Accept(v Visitor[T]) {
	v.VisitUnary(u)
}
