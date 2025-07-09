package expr

import (
	"github.com/gsanchezgavier/craftinginterpreters/src/token"
)

func Example_plus() {
	p := Printer{}
	p.Print(Binary{
		Left:     Literal{Value: 1},
		Operator: token.New(token.PLUS, "+", 1),
		Right:    Literal{Value: 2},
	})
	// Output: (+ 1 2)
}
func Example_unary() {
	p := Printer{}
	p.Print(Unary{
		Operator: token.New(token.MINUS, "-", 1),
		Right:    Literal{Value: 1},
	})
	// Output: (- 1)
}
func Example_grouping() {
	p := Printer{}
	p.Print(Grouping{
		expression: Binary{
			Left:     Literal{Value: 1},
			Operator: token.New(token.PLUS, "+", 1),
			Right:    Literal{Value: 2},
		},
	})
	// Output: (group (+ 1 2))
}
