package extlsp

import (
	"reflect"
	"testing"

	"go.lsp.dev/protocol"
)

func TestUTF16Len(t *testing.T) {
	cases := []struct {
		in   string
		want uint32
	}{
		{"", 0},
		{"abc", 3},
		{"日本", 2},          // BMP runes: one code unit each.
		{"a\U0001F600", 3}, // astral rune: surrogate pair (2 units) + 'a'.
	}
	for _, tc := range cases {
		if got := utf16Len(tc.in); got != tc.want {
			t.Errorf("utf16Len(%q) = %d, want %d", tc.in, got, tc.want)
		}
	}
}

func TestExtractLspSuggestions(t *testing.T) {
	item := func(label string) []protocol.CompletionItem {
		return []protocol.CompletionItem{{Label: label}}
	}
	cases := []struct {
		name  string
		items []protocol.CompletionItem
		line  string
		cx    int
		want  []string
	}{
		{name: "strips typed prefix", items: item("Println"), line: "Pri", cx: 3, want: s("ntln")},
		{name: "no prefix shows whole label", items: item("Println"), line: "x.", cx: 2, want: s("Println")},
		{name: "fully typed yields empty", items: item("Println"), line: "Println", cx: 7, want: nil},
		{name: "mismatched prefix bails out", items: item("Println"), line: "Foo", cx: 3, want: nil},
		{name: "no items", items: nil, line: "Pri", cx: 3, want: nil},
		{name: "prefix only counts identifier chars", items: item("foo"), line: "a + f", cx: 5, want: s("oo")},
		{name: "cx past line end is clamped", items: item("Println"), line: "Pri", cx: 99, want: s("ntln")},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := extractLspSuggestions(tc.items, tc.line, tc.cx); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("extractLspSuggestions(%q, cx=%d) = %v, want %v", tc.line, tc.cx, got, tc.want)
			}
		})
	}
}

func TestIdentTrailing(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Pri", "Pri"},
		{"x.Pri", "Pri"},
		{"a + foo", "foo"},
		{"foo(", ""},
		{"under_score", "under_score"},
		{"a1b2", "a1b2"},
		{"", ""},
	}
	for _, tc := range cases {
		if got := identTrailing(tc.in); got != tc.want {
			t.Errorf("identTrailing(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func s(s ...string) []string { return s }

// TODO: Add tests that verify how inputs are handled (in input_test.go), and here for checking how stale suggestions
//       are dropped.
