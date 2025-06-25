package expr

import (
	"github.com/gsanchezgavier/craftinginterpreters/src/token"
)

func Example_plus() {
	p := Printer{}
	p.Print(Binary{
		left:     Literal{value: 1},
		operator: token.New(token.PLUS, "+", 1),
		right:    Literal{value: 2},
	})
	// Output: (+ 1 2)
}
func Example_unary() {
	p := Printer{}
	p.Print(Unary{
		operator: token.New(token.MINUS, "-", 1),
		right:    Literal{value: 1},
	})
	// Output: (- 1)
}
func Example_grouping() {
	p := Printer{}
	p.Print(Grouping{
		expression: Binary{
			left:     Literal{value: 1},
			operator: token.New(token.PLUS, "+", 1),
			right:    Literal{value: 2},
		},
	})
	// Output: (group (+ 1 2))
}
