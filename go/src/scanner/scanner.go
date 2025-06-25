package scanner

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/gsanchezgavier/craftinginterpreters/src/token"
)

const (
	quotesSize = 1
	//TODO nil rune?
	end = -1
)

type Scanner struct {
	source  string
	start   int
	current int
	line    int
	tokens  []token.Token
}

func New(source string) Scanner {
	return Scanner{
		source: source,
		line:   1,
	}
}

func (s *Scanner) ScanTokens() []token.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, token.New(token.EOF, "", s.line))

	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
func (s *Scanner) advance() rune {
	// TODO handle error
	r, size := utf8.DecodeRuneInString(s.source[s.current:])
	// advance the right number of bytes depending on the rune size
	s.current += size
	return r
}
func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return end
	}
	r, _ := utf8.DecodeRuneInString(s.source[s.current:])
	return r
}
func (s *Scanner) peekNext() rune {
	if s.isAtEnd() {
		return end
	}
	_, size := utf8.DecodeRuneInString(s.source[s.current:])
	// if peek() is at the end
	if s.current+size >= len(s.source) {
		return end
	}
	r, _ := utf8.DecodeRuneInString(s.source[s.current+size:])

	return r
}
func (s *Scanner) advanceIfMatch(value rune) bool {
	if s.isAtEnd() {
		return false
	}
	r, size := utf8.DecodeRuneInString(s.source[s.current:])
	if r != value {
		return false
	}
	s.current += size
	return true
}
func (s *Scanner) addToken(t token.TokenType) {
	lexeme := string(s.source[s.start:s.current])
	s.tokens = append(s.tokens, token.New(t, lexeme, s.line))
}
func (s *Scanner) addLiteralToken(t token.TokenType, literal any) {
	lexeme := string(s.source[s.start:s.current])
	s.tokens = append(s.tokens, token.NewLiteralToken(t, lexeme, s.line, literal))
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	if tokenType, ok := token.Keywords[text]; ok {
		s.addToken(tokenType)
	} else {
		s.addToken(token.IDENTIFIER)
	}

}
func (s *Scanner) number() {
	// different logic that the book.
	for {
		peek := s.peek()
		if isDigit(peek) {
			// consume the number int or fractional
			s.advance()
			continue
		}
		// is fractional, consume '.'
		if peek == '.' && isDigit(s.peekNext()) {
			s.advance()
		} else {
			// neither a digit nor a fractional point.
			break
		}
	}
	val, _ := strconv.ParseFloat(s.source[s.start:s.current], 64)
	s.addLiteralToken(token.NUMBER, val)
}
func (s *Scanner) string() {
	for {
		r := s.peek()
		if r == end {
			fmt.Printf("Unterminated string: %d", s.line)
			return
		}
		s.advance()
		if r == '\n' {
			s.line++
		}
		if r == '"' {
			break
		}
	}
	// The book logic
	// for s.peek() != '"' && !s.isAtEnd() {
	// 	if s.peek() == '\n' {
	// 		s.line++
	// 	}
	// 	s.advance()
	// }
	// if s.isAtEnd() {
	// 	fmt.Printf("Unterminated string: %d", s.line)
	// 	return
	// }
	// s.advance()

	value := s.source[s.start+quotesSize : s.current-quotesSize]
	s.addLiteralToken(token.STRING, value)
}
func (s *Scanner) scanToken() {
	r := s.advance()
	switch r {
	// single char token
	case '(':
		s.addToken(token.LEFT_PAREN)
	case ')':
		s.addToken(token.RIGHT_PAREN)
	case '{':
		s.addToken(token.LEFT_BRACE)
	case '}':
		s.addToken(token.RIGHT_BRACE)
	case ',':
		s.addToken(token.COMMA)
	case '.':
		s.addToken(token.DOT)
	case '-':
		s.addToken(token.MINUS)
	case '+':
		s.addToken(token.PLUS)
	case ';':
		s.addToken(token.SEMICOLON)
	case '*':
		s.addToken(token.STAR)
	// double chart tokens
	case '!':
		if s.advanceIfMatch('=') {
			s.addToken(token.BANG_EQUAL)
		} else {
			s.addToken(token.BANG)
		}
	case '=':
		if s.advanceIfMatch('=') {
			s.addToken(token.EQUAL_EQUAL)
		} else {
			s.addToken(token.EQUAL)
		}
	case '<':
		if s.advanceIfMatch('=') {
			s.addToken(token.LESS_EQUAL)
		} else {
			s.addToken(token.LESS)
		}
	case '>':
		if s.advanceIfMatch('=') {
			s.addToken(token.GREATER_EQUAL)
		} else {
			s.addToken(token.GREATER)
		}
	// comments
	case '/':
		switch {
		case s.advanceIfMatch('/'):
			// A comment goes until the end of the line.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		case s.advanceIfMatch('*'):
		commentLoop:
			for {
				r := s.peek()
				switch r {
				case end:
					fmt.Printf("Unterminated comment: %d", s.line)
					return
				case '\n':
					s.line++
				case '*':
					if s.peekNext() == '/' {
						s.advance()
						s.advance()
						break commentLoop
					}
				}
				s.advance()
			}
		default:
			s.addToken(token.SLASH)
		}
	// white spaces
	case ' ', '\r', '\t':
		break
	// newline
	case '\n':
		s.line++
	// literals
	case '"':
		s.string()
	default:
		switch {
		case isDigit(r):
			s.number()
		case isAlpha(r):
			s.identifier()
		default:
			// TODO log
			fmt.Printf("Unexpected character: %d", s.line)
		}
	}
}

func isAlphaNumeric(r rune) bool {
	return isDigit(r) || isAlpha(r)
}
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
func isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}
