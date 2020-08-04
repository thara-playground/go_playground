package json

import (
	"math/big"
)

func Parse(input string) Value {
	l := NewLexer(input)
	t := l.nextToken()
	return newValue(l, t)
}

func newValue(l *lexer, t Token) Value {
	for t.Type != EOF {
		switch t.Type {
		case LBRACKET:
			return newArray(l)
		case LBRACE:
			return newObject(l)
		default:
			v, ok := newSingleValue(t)
			if ok {
				return v
			}
			t = l.nextToken()
		}
	}
	return nil
}

func isArrayOrObject(t Token) bool {
	return t.Type == LBRACKET || t.Type == LBRACE
}

func newArray(l *lexer) Value {
	var values []Value
	t := l.nextToken()

	for t.Type != RBRACKET {
		if t.Type == COMMA {
			t = l.nextToken()
			continue
		}

		if isArrayOrObject(t) {
			v := newValue(l, t)
			values = append(values, v)
		} else {
			v, ok := newSingleValue(t)
			if ok {
				values = append(values, v)
			}
		}
		t = l.nextToken()
	}
	return &Array{values}
}

func newObject(l *lexer) Value {
	values := make(map[string]Value)

	t := l.nextToken()
	for t.Type != RBRACE {
		if t.Type != STRING {
			// COMMA
			t = l.nextToken()
			continue
		}
		key := t.Value

		t = l.nextToken()
		if t.Type != COLON {
			panic("illegal format (lexer bug)")
		}

		t = l.nextToken()
		v := newValue(l, t)
		if v == nil {
			break
		}
		values[key] = v

		t = l.nextToken()
	}
	return &Object{values}
}

func newSingleValue(t Token) (Value, bool) {
	switch t.Type {
	case STRING:
		return newString(t), true
	case NUMBER:
		return newNumber(t), true
	case TRUE, FALSE:
		return newBool(t), true
	case NULL:
		return NULL_OBJECT, true
	default:
		return nil, false
	}
}

func newString(t Token) Value {
	return &String{t.Value}
}

func newNumber(t Token) Value {
	v, ok := new(big.Float).SetString(t.Value)
	if !ok {
		panic("big.Float SetString is failed") // bug
	}
	return &Number{v}
}

func newBool(t Token) Value {
	if t.Type == TRUE {
		return &Bool{true}
	} else {
		return &Bool{false}
	}
}
