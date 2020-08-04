package json

import (
	"fmt"
	"math/big"
	"testing"
)

func TestParse(t *testing.T) {

	var tests = []struct {
		input    string
		expected Value
	}{
		{"", nil},
		{`  "123"`, &String{"123"}},
		{"  123    ", toNumber("123")},

		{"[]", &Array{values: make([]Value, 0)}},
		{`["123"]`, &Array{values: []Value{&String{"123"}}}},
		{`["123", "345"]`, &Array{values: []Value{&String{"123"}, &String{"345"}}}},
		{`[123, 345]`, &Array{values: []Value{toNumber("123"), toNumber("345")}}},
		{`[{"x": 3, "y": 4}]`, &Array{values: []Value{
			&Object{values: map[string]Value{
				"x": toNumber("3"),
				"y": toNumber("4"),
			}},
		}}},

		{"{}", &Object{values: make(map[string]Value, 0)}},
		{`{"k": "abc"}`, &Object{values: map[string]Value{
			"k": &String{"abc"},
		}}},
		{`{"k": "abc"}`, &Object{values: map[string]Value{
			"k": &String{"abc"},
		}}},
		{`{"k1": "abc", "k2": 123}`, &Object{values: map[string]Value{
			"k1": &String{"abc"},
			"k2": toNumber("123"),
		}}},
		{`{"k": ["abc", "def"]}`, &Object{values: map[string]Value{
			"k": &Array{values: []Value{&String{"abc"}, &String{"def"}}},
		}}},
		{`{"k1": [{"x": 3, "y": 4}, {"x": 9, "y": 7}], "k2": {"a": 123, "b": [123,456]}}`, &Object{values: map[string]Value{
			"k1": &Array{values: []Value{
				&Object{values: map[string]Value{
					"x": toNumber("3"),
					"y": toNumber("4"),
				}},
				&Object{values: map[string]Value{
					"x": toNumber("9"),
					"y": toNumber("7"),
				}},
			}},
			"k2": &Object{values: map[string]Value{
				"a": toNumber("123"),
				"b": &Array{values: []Value{toNumber("123"), toNumber("456")}},
			}},
		}}},
	}

	for i, tt := range tests {
		actual := Parse(tt.input)
		if actual == nil {
			if tt.expected != nil {
				t.Errorf("[%d] this actual result is nil but the expected result is not", i)
			}
		} else {
			if !actual.Equals(tt.expected) {
				fmt.Printf("actual: %v\n", actual)
				fmt.Printf("expected: %v\n", tt.expected)
				t.Errorf("[%d] this actual result does not match the expected result", i)
			}
		}
	}
}

func toNumber(s string) *Number {
	v, ok := new(big.Float).SetString(s)
	if !ok {
		panic("big.Float SetString is failed")
	}
	return &Number{v}
}
