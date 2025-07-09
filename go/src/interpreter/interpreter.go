package interpreter

import (
	"fmt"

	e "github.com/gsanchezgavier/craftinginterpreters/src/expr"
	t "github.com/gsanchezgavier/craftinginterpreters/src/token"
)

type Interpreter struct{}

func (i Interpreter) Interpret(e e.Expr) {
	val := i.evaluate(e)
	fmt.Print(val)
}

func (i Interpreter) VisitLiteral(l e.Literal) any {
	return l.Value
}
func (i Interpreter) VisitGrouping(g e.Grouping) any {
	return i.evaluate(g)
}
func (i Interpreter) VisitUnary(u e.Unary) any {
	right := i.evaluate(u.Right)
	switch u.Operator.TokenType {
	case t.MINUS:
		return -right.(float64)
	case t.BANG:
		return !isTruthy(right)
	default:
		return nil
	}
}
func (i Interpreter) VisitBinary(b e.Binary) any {
	left := i.evaluate(b.Left)
	right := i.evaluate(b.Right)

	switch b.Operator.TokenType {
	case t.MINUS:
		return left.(float64) - right.(float64)
	case t.PLUS:
		if l, ok := left.(float64); ok {
			if r, ok := right.(float64); ok {
				return l + r
			}
		}
		if l, ok := left.(string); ok {
			if r, ok := right.(string); ok {
				return l + r
			}
		}
		return nil
	case t.SLASH:
		return left.(float64) / right.(float64)
	case t.STAR:
		return left.(float64) * right.(float64)
	case t.EQUAL_EQUAL:
		return left == right
	case t.BANG_EQUAL:
		return left != right
	case t.GREATER:
		return left.(float64) > right.(float64)
	case t.GREATER_EQUAL:
		return left.(float64) >= right.(float64)
	case t.LESS:
		return left.(float64) < right.(float64)
	case t.LESS_EQUAL:
		return left.(float64) <= right.(float64)

	default:
		return nil
	}
}

func (i Interpreter) evaluate(e e.Expr) any {
	return e.Accept(i)
}

func isTruthy(obj any) bool {
	if obj == nil {
		return false
	}
	if b, ok := obj.(bool); ok {
		return b
	}
	return true
}
