package scanner_test

import (
	"reflect"
	"testing"

	"github.com/gsanchezgavier/craftinginterpreters/src/scanner"
	"github.com/gsanchezgavier/craftinginterpreters/src/token"
)

func TestScanner_ScanToken_SingleCharacterTokens(t *testing.T) {
	source := "(){}.,-+;*"
	expectedTokens := []token.Token{
		token.New(token.LEFT_PAREN, "(", 1),
		token.New(token.RIGHT_PAREN, ")", 1),
		token.New(token.LEFT_BRACE, "{", 1),
		token.New(token.RIGHT_BRACE, "}", 1),
		token.New(token.DOT, ".", 1),
		token.New(token.COMMA, ",", 1),
		token.New(token.MINUS, "-", 1),
		token.New(token.PLUS, "+", 1),
		token.New(token.SEMICOLON, ";", 1),
		token.New(token.STAR, "*", 1),
		token.New(token.EOF, "", 1),
	}

	scanner := scanner.New(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}

func TestScanner_ScanToken_Operators(t *testing.T) {
	source := "! != = == < <= > >="
	expectedTokens := []token.Token{
		token.New(token.BANG, "!", 1),
		token.New(token.BANG_EQUAL, "!=", 1),
		token.New(token.EQUAL, "=", 1),
		token.New(token.EQUAL_EQUAL, "==", 1),
		token.New(token.LESS, "<", 1),
		token.New(token.LESS_EQUAL, "<=", 1),
		token.New(token.GREATER, ">", 1),
		token.New(token.GREATER_EQUAL, ">=", 1),
		token.New(token.EOF, "", 1),
	}

	scanner := scanner.New(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}

func TestScanner_ScanToken_Comments(t *testing.T) {
	source := "// this is a comment\n+ -"
	expectedTokens := []token.Token{
		token.New(token.PLUS, "+", 2),
		token.New(token.MINUS, "-", 2),
		token.New(token.EOF, "", 2),
	}

	scanner := scanner.New(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}
func TestScanner_ScanToken_MultilineComments(t *testing.T) {
	source := "/* this is a \n multiline comment */ + -"
	expectedTokens := []token.Token{
		token.New(token.PLUS, "+", 2),
		token.New(token.MINUS, "-", 2),
		token.New(token.EOF, "", 2),
	}

	scanner := scanner.New(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}
func TestScanner_ScanToken_Multiline(t *testing.T) {
	source := " /*a*/ + -"
	expectedTokens := []token.Token{
		token.New(token.PLUS, "+", 1),
		token.New(token.MINUS, "-", 1),
		token.New(token.EOF, "", 1),
	}

	scanner := scanner.New(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}

func TestScanner_ScanToken_UnexpectedCharacter(t *testing.T) {
	source := "@"
	expectedTokens := []token.Token{
		token.New(token.EOF, "", 1),
	}

	scanner := scanner.New(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}
func TestScanner_ScanToken_Strings(t *testing.T) {
	source := `"Hello, World!" 
	"Another string" "unterminated`
	expectedTokens := []token.Token{
		token.NewLiteralToken(token.STRING, `"Hello, World!"`, 1, "Hello, World!"),
		token.NewLiteralToken(token.STRING, `"Another string"`, 2, "Another string"),
		token.New(token.EOF, "", 2),
	}

	scanner := scanner.New(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}

func TestScanner_ScanToken_Numbers(t *testing.T) {
	source := "123 45.67 0.89 0. "
	expectedTokens := []token.Token{
		token.NewLiteralToken(token.NUMBER, "123", 1, 123.0),
		token.NewLiteralToken(token.NUMBER, "45.67", 1, 45.67),
		token.NewLiteralToken(token.NUMBER, "0.89", 1, 0.89),
		token.NewLiteralToken(token.NUMBER, "0", 1, 0.0),
		token.New(token.DOT, ".", 1),
		token.New(token.EOF, "", 1),
	}

	scanner := scanner.New(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}

func TestScanner_ScanToken_Identifiers(t *testing.T) {
	source := "varName anotherVar _privateVar and andalucia"
	expectedTokens := []token.Token{
		token.New(token.IDENTIFIER, "varName", 1),
		token.New(token.IDENTIFIER, "anotherVar", 1),
		token.New(token.IDENTIFIER, "_privateVar", 1),
		token.New(token.AND, "and", 1),
		token.New(token.IDENTIFIER, "andalucia", 1),
		token.New(token.EOF, "", 1),
	}

	scanner := scanner.New(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}
