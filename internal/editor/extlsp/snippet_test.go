package extlsp

import "testing"

func TestExpandSnippet(t *testing.T) {
	cases := []struct {
		in     string
		text   string
		cursor int
	}{
		{"Println", "Println", 7},
		{"Println(${1:})", "Println()", 8},
		{"DeleteSpan(${1:})", "DeleteSpan()", 11},
		{`Position{$0\}`, "Position{}", 9},
		{"Line: ${1:}", "Line: ", 6},
		{"f(${1:a}, ${2:b})$0", "f(a, b)", 2},
		{"f(${2:b}, ${1:a})", "f(b, a)", 5},
		{"x${1}y", "xy", 1},
		{`a\$b\\c`, `a$b\c`, 5},
		{"${1:outer ${2:inner}}", "outer inner", 0},
		{"${1|one,two|} end", "one end", 0},
		{"price $", "price $", 7},
	}
	for _, tc := range cases {
		text, cursor := expandSnippet(tc.in)
		if text != tc.text || cursor != tc.cursor {
			t.Errorf("expandSnippet(%q) = (%q, %d), want (%q, %d)", tc.in, text, cursor, tc.text, tc.cursor)
		}
	}
}
