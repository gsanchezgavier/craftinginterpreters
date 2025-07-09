package parser

import (
	"fmt"

	"github.com/gsanchezgavier/craftinginterpreters/src/expr"
	t "github.com/gsanchezgavier/craftinginterpreters/src/token"
)

type Parser struct {
	tokens  []t.Token
	current int
}

func New(tokens []t.Token) Parser {
	return Parser{
		tokens: tokens,
	}
}

func (p *Parser) Parse() expr.Expr {
	return p.expression()
}

func (p *Parser) expression() expr.Expr {
	return p.equality()
}

func (p *Parser) equality() expr.Expr {
	e := p.comparison()

	for p.match(t.BANG_EQUAL, t.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		e = expr.NewBinary(e, operator, right)
	}

	return e
}
func (p *Parser) comparison() expr.Expr {
	e := p.term()

	for p.match(t.GREATER, t.GREATER_EQUAL, t.LESS, t.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		e = expr.NewBinary(e, operator, right)
	}

	return e
}
func (p *Parser) term() expr.Expr {
	e := p.factor()

	for p.match(t.MINUS, t.PLUS) {
		operator := p.previous()
		right := p.factor()
		e = expr.NewBinary(e, operator, right)
	}

	return e
}
func (p *Parser) factor() expr.Expr {
	e := p.unary()

	for p.match(t.SLASH, t.STAR) {
		operator := p.previous()
		right := p.unary()
		e = expr.NewBinary(e, operator, right)
	}

	return e
}
func (p *Parser) unary() expr.Expr {
	if p.match(t.BANG, t.MINUS) {
		operator := p.previous()
		right := p.unary()
		return expr.NewUnary(operator, right)
	}
	return p.primary()
}
func (p *Parser) primary() expr.Expr {
	if p.match(t.FALSE) {
		return expr.NewLiteral(false)
	}
	if p.match(t.TRUE) {
		return expr.NewLiteral(true)
	}
	if p.match(t.NIL) {
		return expr.NewLiteral(nil)
	}

	if p.match(t.NUMBER, t.STRING) {
		return expr.NewLiteral(p.previous().Literal)
	}

	if p.match(t.LEFT_PAREN) {
		e := p.expression()
		_, _ = p.consume(t.RIGHT_PAREN, "Expect ')' after expression.")
		return expr.NewGrouping(e)
	}
	printError(p.peek(), "Expect expression.")
	// TODO
	panic("Unexpected token")
}

func (p *Parser) consume(tokenType t.TokenType, message string) (t.Token, error) {
	if !p.check(tokenType) {
		return t.Token{}, fmt.Errorf("checking")
	}
	return p.advance(), nil

}

func (p *Parser) match(types ...t.TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}
func (p *Parser) check(t t.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().TokenType == t
}
func (p *Parser) advance() t.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}
func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == t.EOF

}
func (p *Parser) peek() t.Token {
	return p.tokens[p.current]
}
func (p *Parser) previous() t.Token {
	if p.current == 0 {
		return p.tokens[0]
	}
	return p.tokens[p.current-1]
}

// private void synchronize() {
//     advance();

//     while (!isAtEnd()) {
//       if (previous().type == SEMICOLON) return;

//       switch (peek().type) {
//         case CLASS:
//         case FUN:
//         case VAR:
//         case FOR:
//         case IF:
//         case WHILE:
//         case PRINT:
//         case RETURN:
//           return;
//       }

//       advance();
//     }
//   }

func printError(token t.Token, message string) {
	if token.TokenType == t.EOF {
		fmt.Printf("%d at end: %s\n", token.Line, message)
	} else {
		fmt.Printf("%d at '%s': %s\n", token.Line, token.Lexeme, message)
	}
}
