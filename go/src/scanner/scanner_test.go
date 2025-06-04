package scanner_test

import (
	"reflect"
	"testing"

	"github.com/gsanchezgavier/craftinginterpreters/src/scanner"
)

func TestScanner_ScanToken_SingleCharacterTokens(t *testing.T) {
	source := "(){}.,-+;*"
	expectedTokens := []scanner.Token{
		scanner.NewToken(scanner.LEFT_PAREN, "(", 1),
		scanner.NewToken(scanner.RIGHT_PAREN, ")", 1),
		scanner.NewToken(scanner.LEFT_BRACE, "{", 1),
		scanner.NewToken(scanner.RIGHT_BRACE, "}", 1),
		scanner.NewToken(scanner.DOT, ".", 1),
		scanner.NewToken(scanner.COMMA, ",", 1),
		scanner.NewToken(scanner.MINUS, "-", 1),
		scanner.NewToken(scanner.PLUS, "+", 1),
		scanner.NewToken(scanner.SEMICOLON, ";", 1),
		scanner.NewToken(scanner.STAR, "*", 1),
		scanner.NewToken(scanner.EOF, "", 1),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}

func TestScanner_ScanToken_Operators(t *testing.T) {
	source := "! != = == < <= > >="
	expectedTokens := []scanner.Token{
		scanner.NewToken(scanner.BANG, "!", 1),
		scanner.NewToken(scanner.BANG_EQUAL, "!=", 1),
		scanner.NewToken(scanner.EQUAL, "=", 1),
		scanner.NewToken(scanner.EQUAL_EQUAL, "==", 1),
		scanner.NewToken(scanner.LESS, "<", 1),
		scanner.NewToken(scanner.LESS_EQUAL, "<=", 1),
		scanner.NewToken(scanner.GREATER, ">", 1),
		scanner.NewToken(scanner.GREATER_EQUAL, ">=", 1),
		scanner.NewToken(scanner.EOF, "", 1),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}

func TestScanner_ScanToken_Comments(t *testing.T) {
	source := "// this is a comment\n+ -"
	expectedTokens := []scanner.Token{
		scanner.NewToken(scanner.PLUS, "+", 2),
		scanner.NewToken(scanner.MINUS, "-", 2),
		scanner.NewToken(scanner.EOF, "", 2),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}
func TestScanner_ScanToken_MultilineComments(t *testing.T) {
	source := "/* this is a \n multiline comment */ + -"
	expectedTokens := []scanner.Token{
		scanner.NewToken(scanner.PLUS, "+", 2),
		scanner.NewToken(scanner.MINUS, "-", 2),
		scanner.NewToken(scanner.EOF, "", 2),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}
func TestScanner_ScanToken_Multiline(t *testing.T) {
	source := " /*a*/ + -"
	expectedTokens := []scanner.Token{
		scanner.NewToken(scanner.PLUS, "+", 1),
		scanner.NewToken(scanner.MINUS, "-", 1),
		scanner.NewToken(scanner.EOF, "", 1),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}

func TestScanner_ScanToken_UnexpectedCharacter(t *testing.T) {
	source := "@"
	expectedTokens := []scanner.Token{
		scanner.NewToken(scanner.EOF, "", 1),
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
	expectedTokens := []scanner.Token{
		scanner.NewLiteralToken(scanner.STRING, `"Hello, World!"`, 1, "Hello, World!"),
		scanner.NewLiteralToken(scanner.STRING, `"Another string"`, 2, "Another string"),
		scanner.NewToken(scanner.EOF, "", 2),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}

func TestScanner_ScanToken_Numbers(t *testing.T) {
	source := "123 45.67 0.89 0. "
	expectedTokens := []scanner.Token{
		scanner.NewLiteralToken(scanner.NUMBER, "123", 1, 123.0),
		scanner.NewLiteralToken(scanner.NUMBER, "45.67", 1, 45.67),
		scanner.NewLiteralToken(scanner.NUMBER, "0.89", 1, 0.89),
		scanner.NewLiteralToken(scanner.NUMBER, "0", 1, 0.0),
		scanner.NewToken(scanner.DOT, ".", 1),
		scanner.NewToken(scanner.EOF, "", 1),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}

func TestScanner_ScanToken_Identifiers(t *testing.T) {
	source := "varName anotherVar _privateVar and andalucia"
	expectedTokens := []scanner.Token{
		scanner.NewToken(scanner.IDENTIFIER, "varName", 1),
		scanner.NewToken(scanner.IDENTIFIER, "anotherVar", 1),
		scanner.NewToken(scanner.IDENTIFIER, "_privateVar", 1),
		scanner.NewToken(scanner.AND, "and", 1),
		scanner.NewToken(scanner.IDENTIFIER, "andalucia", 1),
		scanner.NewToken(scanner.EOF, "", 1),
	}

	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, but got %v", expectedTokens, tokens)
	}
}
