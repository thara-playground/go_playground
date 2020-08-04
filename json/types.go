package json

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type Value interface {
	fmt.Stringer

	Equals(Value) bool
}

// Object is an unordered set of name/value pairs
type Object struct {
	values map[string]Value
}

var _ Value = &Object{values: nil}

func (o *Object) Equals(v Value) bool {
	if (o == nil) != (v == nil) {
		return false
	}
	o2, ok := v.(*Object)
	if !ok {
		return false
	}
	if len(o.values) != len(o2.values) {
		return false
	}

	keys := sortedKeys(o.values)
	keys2 := sortedKeys(o2.values)

	for i, k := range keys {
		if k != keys2[i] {
			return false
		}
		v1 := o.values[k]
		v2 := o2.values[k]
		if !v1.Equals(v2) {
			return false
		}
	}

	return true
}

func (o *Object) String() string {
	keys := sortedKeys(o.values)

	var b strings.Builder
	fmt.Fprintf(&b, "{")
	for _, k := range keys {
		fmt.Fprintf(&b, "%s:%s,", k, o.values[k])
	}
	fmt.Fprintf(&b, "}")
	return b.String()
}

func sortedKeys(v map[string]Value) []string {
	keys := make([]string, 0)
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Array is an ordered collection of values
type Array struct {
	values []Value
}

var _ Value = &Array{values: nil}

func (a *Array) Equals(v Value) bool {
	if (a == nil) != (v == nil) {
		return false
	}
	a2, ok := v.(*Array)
	if !ok {
		return false
	}
	if len(a.values) != len(a2.values) {
		return false
	}

	for i, v1 := range a.values {
		v2 := a2.values[i]
		if !v1.Equals(v2) {
			return false
		}
	}
	return true
}

func (a *Array) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "[")
	for _, v := range a.values {
		fmt.Fprintf(&b, "%s,", v)
	}
	fmt.Fprintf(&b, "]")
	return b.String()
}

// String is a sequence of zero or more Unicode characters
type String struct {
	value string
}

var _ Value = &String{value: ""}

func (s *String) Equals(v Value) bool {
	if (s == nil) != (v == nil) {
		return false
	}
	s2, ok := v.(*String)
	if !ok {
		return false
	}
	return s.value == s2.value
}

func (s *String) String() string {
	return `"` + s.value + `"`
}

// Number is very much like a C or Java number, except that the octal and hexadecimal formats are not used.
type Number struct {
	value *big.Float
}

var _ Value = &Number{value: new(big.Float)}

func (n *Number) Equals(v Value) bool {
	if (n == nil) != (v == nil) {
		return false
	}
	n2, ok := v.(*Number)
	if !ok {
		return false
	}
	return n.value.Cmp(n2.value) == 0
}

func (n *Number) String() string {
	return n.value.String()
}

type Bool struct {
	value bool
}

var _ Value = &Bool{value: false}

func (b *Bool) Equals(v Value) bool {
	if (b == nil) != (v == nil) {
		return false
	}
	b2, ok := v.(*Bool)
	if !ok {
		return false
	}
	return b.value == b2.value
}

func (b *Bool) String() string {
	return strconv.FormatBool(b.value)
}

type Null struct{}

var NULL_OBJECT Value = &Null{}

func (n *Null) Equals(v Value) bool {
	if (n == nil) != (v == nil) {
		return false
	}
	_, ok := v.(*Null)
	return ok
}

func (n *Null) String() string {
	return "null"
}
