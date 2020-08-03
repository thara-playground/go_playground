package json

import (
	"fmt"
	"strings"
	"unicode"
)

type lexer struct {
	input    []rune
	position int
	char     rune
}

func NewLexer(input string) *lexer {
	l := &lexer{[]rune(input), 0, 0}
	l.readChar()
	return l
}

func (l *lexer) nextToken() Token {
	l.skipWhitespaces()

	if l.position < 0 {
		return NewToken(EOF, "")
	}

	var t Token
	char := string(l.char)
	switch char {
	case "{":
		t = NewToken(LBRACE, char)
	case "}":
		t = NewToken(RBRACE, char)
	case "[":
		t = NewToken(LBRACKET, char)
	case "]":
		t = NewToken(RBRACKET, char)
	case ",":
		t = NewToken(COMMA, char)
	case ":":
		t = NewToken(COLON, char)
	case `"`:
		l.readChar()
		t = NewToken(STRING, l.readString())
	case "-":
		t = NewToken(MINUS, char)
	default:
		if unicode.IsDigit(l.char) {
			return NewToken(NUMBER, l.readNumber())
		} else {
			t = NewToken(ILLEGAL, "")
		}
	}

	l.readChar()
	return t
}

func (l *lexer) readChar() bool {
	if len(l.input) <= l.position {
		l.char = 0
		l.position = -1
		return false
	} else {
		l.char = l.input[l.position]
		l.position++
		return true
	}
}

func (l *lexer) skipWhitespaces() {
	for IsWhiteSpace(l.char) {
		if !l.readChar() {
			break
		}
	}
}

func (l *lexer) readNumber() string {
	//TODO support fraction, exponent, and scientific notation
	var b strings.Builder
	for unicode.IsDigit(l.char) {
		fmt.Fprintf(&b, "%s", string(l.char))
		if !l.readChar() {
			break
		}
	}
	return b.String()
}

func (l *lexer) readString() string {
	var b strings.Builder
	for l.char != '"' {
		fmt.Fprintf(&b, "%s", string(l.char))
		if !l.readChar() {
			break
		}
	}
	return b.String()
}
