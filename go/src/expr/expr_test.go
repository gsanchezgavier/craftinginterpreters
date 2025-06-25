package expr

// type Printer struct{}

// func (p Printer) print(expr Expr[T]) {
// 	expr.Accept(p)
// }

// func (p Printer) VisitBinary(b Binary) {
// 	p.parenthesize(b.operator.lexeme, b.left, b.right)
// }
// func (p Printer) VisitGrouping(g Grouping) {}
// func (p Printer) VisitLiteral(l Literal) {
// 	if l.value != nil {
// 		fmt.Printf("%v", l.value)
// 	}
// }
// func (p Printer) VisitUnary(u Unary) {}

// func (p Printer) parenthesize(name string, expr ...Expr) {
// 	fmt.Print("(")
// 	fmt.Print(name)
// 	for _, expr := range expr {
// 		fmt.Print(" ")
// 		expr.Accept(p)
// 	}
// 	fmt.Print(")")
// }

// func Test_print(t *testing.T) {
// 	p := Printer{}
// 	p.print(Binary{
// 		left:     Literal{value: 1},
// 		operator: NewToken(PLUS, "+", 1),
// 		right:    Literal{value: 2},
// 	})
// 	t.FailNow()
// }
