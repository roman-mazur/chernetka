package extlsp

import (
	"testing"

	"go.lsp.dev/protocol"
	"rmazur.io/chernetka/internal/content"
)

func TestSpanOf(t *testing.T) {
	lines := []content.Line{content.TextLine("package main"), content.TextLine("x := \"日本\" // ok")}
	pos := func(line, char uint32) protocol.Position { return protocol.Position{Line: line, Character: char} }
	cases := []struct {
		name string
		r    protocol.Range
		want content.Span
		ok   bool
	}{
		{"ascii", protocol.Range{Start: pos(0, 0), End: pos(0, 7)}, content.Span{End: content.Position{Col: 7}}, true},
		{
			// 日 and 本 are 1 UTF-16 unit and 3 bytes each.
			"utf16 columns become bytes", protocol.Range{Start: pos(1, 6), End: pos(1, 9)},
			content.Span{Start: content.Position{Line: 1, Col: 6}, End: content.Position{Line: 1, Col: 13}}, true,
		},
		{"across lines", protocol.Range{Start: pos(0, 8), End: pos(1, 1)},
			content.Span{Start: content.Position{Col: 8}, End: content.Position{Line: 1, Col: 1}}, true},
		{"past the line end is clamped", protocol.Range{Start: pos(0, 99), End: pos(0, 99)},
			content.Span{Start: content.Position{Col: 12}, End: content.Position{Col: 12}}, true},
		{"past the last line", protocol.Range{Start: pos(1, 0), End: pos(2, 0)}, content.Span{}, false},
	}
	for _, tc := range cases {
		got, ok := spanOf(lines, tc.r)
		if ok != tc.ok || (ok && got != tc.want) {
			t.Errorf("%s: spanOf(%v) = %v, %t; want %v, %t", tc.name, tc.r, got, ok, tc.want, tc.ok)
		}
	}
}

func TestLSPPosition(t *testing.T) {
	cases := []struct {
		line string
		p    content.Position
		want protocol.Position
	}{
		{"fmt.Pri", content.Position{Line: 3, Col: 7}, protocol.Position{Line: 3, Character: 7}},
		{"日本 Pri", content.Position{Line: 0, Col: len("日本 Pri")}, protocol.Position{Character: 6}},
		{"a\U0001F600b", content.Position{Col: 5}, protocol.Position{Character: 3}}, // a surrogate pair
		{"ab", content.Position{Col: 9}, protocol.Position{Character: 2}},           // clamped
	}
	for _, tc := range cases {
		if got := lspPosition(tc.line, tc.p); got != tc.want {
			t.Errorf("lspPosition(%q, %v) = %v, want %v", tc.line, tc.p, got, tc.want)
		}
	}
}

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

func TestByteOffset(t *testing.T) {
	cases := []struct {
		line string
		u16  uint32
		want int
	}{
		{"abc", 0, 0},
		{"abc", 2, 2},
		{"abc", 9, 3},
		{"日本x", 2, len("日本")},
		{"\U0001F600x", 2, 4},
	}
	for _, tc := range cases {
		if got := byteOffset(tc.line, tc.u16); got != tc.want {
			t.Errorf("byteOffset(%q, %d) = %d, want %d", tc.line, tc.u16, got, tc.want)
		}
	}
}
