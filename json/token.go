package json

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

func NewToken(t TokenType, v string) Token {
	return Token{t, v}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"
	COMMA    = ","
	COLON    = ":"

	DOUBLE_QUOTE = `"`
	MINUS        = "-"

	// literals
	STRING = "STRING"
	NUMBER = "number"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	NULL   = "NULL"
)

var WHITESPACES map[rune]bool = map[rune]bool{
	0x0020: true, // space
	0x000A: true, // linefeed
	0x000D: true, // carriage return
	0x0009: true, // horizontal tab
}

func IsWhiteSpace(r rune) bool {
	_, ok := WHITESPACES[r]
	return ok
}
