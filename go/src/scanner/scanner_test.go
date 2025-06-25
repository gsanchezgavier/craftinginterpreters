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
		token.NewToken(token.LEFT_PAREN, "(", 1),
		token.NewToken(token.RIGHT_PAREN, ")", 1),
		token.NewToken(token.LEFT_BRACE, "{", 1),
		token.NewToken(token.RIGHT_BRACE, "}", 1),
		token.NewToken(token.DOT, ".", 1),
		token.NewToken(token.COMMA, ",", 1),
		token.NewToken(token.MINUS, "-", 1),
		token.NewToken(token.PLUS, "+", 1),
		token.NewToken(token.SEMICOLON, ";", 1),
		token.NewToken(token.STAR, "*", 1),
		token.NewToken(token.EOF, "", 1),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}

func TestScanner_ScanToken_Operators(t *testing.T) {
	source := "! != = == < <= > >="
	expectedTokens := []token.Token{
		token.NewToken(token.BANG, "!", 1),
		token.NewToken(token.BANG_EQUAL, "!=", 1),
		token.NewToken(token.EQUAL, "=", 1),
		token.NewToken(token.EQUAL_EQUAL, "==", 1),
		token.NewToken(token.LESS, "<", 1),
		token.NewToken(token.LESS_EQUAL, "<=", 1),
		token.NewToken(token.GREATER, ">", 1),
		token.NewToken(token.GREATER_EQUAL, ">=", 1),
		token.NewToken(token.EOF, "", 1),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}

func TestScanner_ScanToken_Comments(t *testing.T) {
	source := "// this is a comment\n+ -"
	expectedTokens := []token.Token{
		token.NewToken(token.PLUS, "+", 2),
		token.NewToken(token.MINUS, "-", 2),
		token.NewToken(token.EOF, "", 2),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}
func TestScanner_ScanToken_MultilineComments(t *testing.T) {
	source := "/* this is a \n multiline comment */ + -"
	expectedTokens := []token.Token{
		token.NewToken(token.PLUS, "+", 2),
		token.NewToken(token.MINUS, "-", 2),
		token.NewToken(token.EOF, "", 2),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}
func TestScanner_ScanToken_Multiline(t *testing.T) {
	source := " /*a*/ + -"
	expectedTokens := []token.Token{
		token.NewToken(token.PLUS, "+", 1),
		token.NewToken(token.MINUS, "-", 1),
		token.NewToken(token.EOF, "", 1),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}

func TestScanner_ScanToken_UnexpectedCharacter(t *testing.T) {
	source := "@"
	expectedTokens := []token.Token{
		token.NewToken(token.EOF, "", 1),
	}

	scanner := scanner.NewScanner(source)
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
		token.NewToken(token.EOF, "", 2),
	}

	scanner := scanner.NewScanner(source)
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
		token.NewToken(token.DOT, ".", 1),
		token.NewToken(token.EOF, "", 1),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}

func TestScanner_ScanToken_Identifiers(t *testing.T) {
	source := "varName anotherVar _privateVar and andalucia"
	expectedTokens := []token.Token{
		token.NewToken(token.IDENTIFIER, "varName", 1),
		token.NewToken(token.IDENTIFIER, "anotherVar", 1),
		token.NewToken(token.IDENTIFIER, "_privateVar", 1),
		token.NewToken(token.AND, "and", 1),
		token.NewToken(token.IDENTIFIER, "andalucia", 1),
		token.NewToken(token.EOF, "", 1),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}
