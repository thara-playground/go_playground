package json

import (
	"testing"
)

func TestNextToken(t *testing.T) {

	var tests = []struct {
		input    string
		expected []TokenType
	}{
		{"    []   ", []TokenType{LBRACKET, RBRACKET, EOF}},
		{"    [123]   ", []TokenType{LBRACKET, NUMBER, RBRACKET, EOF}},
		{`["abc"    ]`, []TokenType{LBRACKET, STRING, RBRACKET, EOF}},
		{`[123,"abc"]`, []TokenType{LBRACKET, NUMBER, COMMA, STRING, RBRACKET, EOF}},

		{"{}", []TokenType{LBRACE, RBRACE, EOF}},
		{"{     }   ", []TokenType{LBRACE, RBRACE, EOF}},
		{`{"k": 123     }   `, []TokenType{LBRACE, STRING, COLON, NUMBER, RBRACE, EOF}},
		{`{"k1": 3, "k2": "abc"}`, []TokenType{LBRACE, STRING, COLON, NUMBER, COMMA, STRING, COLON, STRING, RBRACE, EOF}},
	}

	for i, tt := range tests {
		l := NewLexer(tt.input)
		for j, expectedType := range tt.expected {
			actual := l.nextToken()
			if actual.Type != expectedType {
				t.Errorf("[%d] %d: actual type '%s' does not match expected type '%s'", i, j, actual.Type, expectedType)
			}
		}
	}
}
