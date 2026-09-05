package extsyntaxhl

import (
	"fmt"
	"testing"

	"rmazur.io/chernetka/internal/editor"
)

func buildSpans(src *source, raw ...rawSpan) []editor.SyntaxSpan {
	var sb spanBuilder
	for _, rs := range raw {
		sb.add(src, rs)
	}
	return sb.build(nil)
}

func formatSpans(spans []editor.SyntaxSpan) string {
	out := make([]string, len(spans))
	for i, s := range spans {
		out[i] = s.String()
	}
	return fmt.Sprint(out)
}

func TestSpanBuilder(t *testing.T) {
	src := newSource("hello world\nsecond line\nthird")

	for _, tc := range []struct {
		name string
		raw  []rawSpan
		want string
	}{{
		name: "single line",
		raw:  []rawSpan{lineSpan(0, 0, 5, editor.TtKeyword)},
		want: "[0:0:5:Keyword]",
	}, {
		// A node spanning several lines becomes one span per line, each clipped
		// to its own line rather than to the end column of the last one.
		name: "multi line is split per line",
		raw:  []rawSpan{{StartLine: 0, StartCol: 6, EndLine: 2, EndCol: 3, TokenType: editor.TtStringLiteral}},
		want: "[0:6:11:StringLiteral 1:0:11:StringLiteral 2:0:3:StringLiteral]",
	}, {
		// A nested span wins over the one enclosing it, which is what puts an
		// escape sequence on top of its string literal.
		name: "nested span splits its parent",
		raw: []rawSpan{
			lineSpan(0, 0, 11, editor.TtStringLiteral),
			lineSpan(0, 5, 7, editor.TtEscape),
		},
		want: "[0:0:5:StringLiteral 0:5:7:Escape 0:7:11:StringLiteral]",
	}, {
		name: "order of nested spans does not matter",
		raw: []rawSpan{
			lineSpan(0, 5, 7, editor.TtEscape),
			lineSpan(0, 0, 11, editor.TtStringLiteral),
		},
		want: "[0:0:5:StringLiteral 0:5:7:Escape 0:7:11:StringLiteral]",
	}, {
		// Two spans over the identical range: the first one reported wins, which
		// is how a highlight query orders its specific patterns first.
		name: "identical ranges resolve to the first",
		raw: []rawSpan{
			lineSpan(0, 0, 5, editor.TtFuncDeclaration),
			lineSpan(0, 0, 5, editor.TtIdentifier),
		},
		want: "[0:0:5:FuncDeclaration]",
	}, {
		name: "partial overlap is resolved, never duplicated",
		raw: []rawSpan{
			lineSpan(0, 0, 6, editor.TtKeyword),
			lineSpan(0, 4, 11, editor.TtIdentifier),
		},
		want: "[0:0:4:Keyword 0:4:11:Identifier]",
	}, {
		name: "adjacent spans of one type merge",
		raw: []rawSpan{
			lineSpan(0, 0, 5, editor.TtComment),
			lineSpan(0, 5, 11, editor.TtComment),
		},
		want: "[0:0:11:Comment]",
	}, {
		name: "spans are clipped to the line",
		raw:  []rawSpan{lineSpan(2, 3, 99, editor.TtKeyword)},
		want: "[2:3:5:Keyword]",
	}, {
		name: "spans outside the document are dropped",
		raw: []rawSpan{
			lineSpan(9, 0, 3, editor.TtKeyword),
			lineSpan(-1, 0, 3, editor.TtKeyword),
			lineSpan(0, 20, 30, editor.TtKeyword),
			lineSpan(0, 4, 4, editor.TtKeyword),
			lineSpan(0, 6, 2, editor.TtKeyword),
		},
		want: "[]",
	}, {
		name: "unmapped token types are dropped",
		raw:  []rawSpan{lineSpan(0, 0, 5, editor.TtNothing)},
		want: "[]",
	}, {
		name: "spans are sorted by line and offset",
		raw: []rawSpan{
			lineSpan(2, 0, 5, editor.TtKeyword),
			lineSpan(0, 6, 11, editor.TtComment),
			lineSpan(0, 0, 5, editor.TtKeyword),
		},
		want: "[0:0:5:Keyword 0:6:11:Comment 2:0:5:Keyword]",
	}} {
		t.Run(tc.name, func(t *testing.T) {
			if got := formatSpans(buildSpans(src, tc.raw...)); got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestSpansOfLine(t *testing.T) {
	src := newSource("aaa\nbbb\nccc\nddd")
	spans := buildSpans(src,
		lineSpan(0, 0, 3, editor.TtKeyword),
		lineSpan(2, 0, 1, editor.TtKeyword),
		lineSpan(2, 2, 3, editor.TtComment),
	)

	for _, tc := range []struct {
		line int
		want string
	}{
		{0, "[0:0:3:Keyword]"},
		{1, "[]"},
		{2, "[2:0:1:Keyword 2:2:3:Comment]"},
		{3, "[]"},
		{99, "[]"},
		{-1, "[]"},
	} {
		if got := formatSpans(spansOfLine(spans, tc.line)); got != tc.want {
			t.Errorf("spansOfLine(%d) = %s, want %s", tc.line, got, tc.want)
		}
	}
}

// TestClipSpans covers the case where the editor renders a line that is shorter
// than it was at the last parse. Without clipping the renderer would be handed
// an offset past the end of the line.
func TestClipSpans(t *testing.T) {
	spans := []editor.SyntaxSpan{
		{LineNumber: 0, Start: 0, End: 4, TokenType: editor.TtKeyword},
		{LineNumber: 0, Start: 5, End: 9, TokenType: editor.TtIdentifier},
	}

	for _, tc := range []struct {
		line string
		want string
	}{
		{"0123456789", "[0:0:4:Keyword 0:5:9:Identifier]"},
		{"012345678", "[0:0:4:Keyword 0:5:9:Identifier]"},
		{"01234567", "[0:0:4:Keyword 0:5:8:Identifier]"},
		{"012345", "[0:0:4:Keyword 0:5:6:Identifier]"},
		{"01234", "[0:0:4:Keyword]"},
		{"0123", "[0:0:4:Keyword]"},
		{"01", "[0:0:2:Keyword]"},
		{"", "[]"},
	} {
		if got := formatSpans(clipSpans(spans, tc.line)); got != tc.want {
			t.Errorf("clipSpans(%q) = %s, want %s", tc.line, got, tc.want)
		}
	}

	// Clipping must not corrupt the cached spans it was handed.
	if got := formatSpans(spans); got != "[0:0:4:Keyword 0:5:9:Identifier]" {
		t.Errorf("clipSpans modified its input: %s", got)
	}
}

func TestNewSourceMatchesBufferLines(t *testing.T) {
	for _, tc := range []struct {
		text string
		want string
	}{
		{"", `[""]`},
		{"one", `["one"]`},
		{"one\ntwo", `["one" "two"]`},
		{"one\n", `["one" ""]`},
		{"\n\n", `["" "" ""]`},
	} {
		if got := fmt.Sprintf("%q", newSource(tc.text).lines); got != tc.want {
			t.Errorf("newSource(%q).lines = %s, want %s", tc.text, got, tc.want)
		}
	}
}
