package extindent

import (
	"slices"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
)

const (
	down       = "\x1b[B"
	right      = "\x1b[C"
	shiftLeft  = "\x1b[1;2D"
	shiftRight = "\x1b[1;2C"
	shiftDown  = "\x1b[1;2B"
	enter      = "\r"
	tab        = "\t"
	shiftTab   = "\x1b[Z"
	esc        = "\x1b"
	undo       = "\x1a"
)

func TestSplitLine(t *testing.T) {
	for _, tc := range []struct {
		name string
		text string // with | at the cursor
		want string // the text after Enter with | at the cursor
	}{
		{"no indentation", "a|b", "a\n|b"},
		{"tabs", "\t\ta|b", "\t\ta\n\t\t|b"},
		{"spaces", "    a|b", "    a\n    |b"},
		{"blank rest is dropped", "\tab | cd", "\tab\n\t|cd"},
		{"end of line", "\tab|", "\tab\n\t|"},
		{"blank indented line", "\t\t|", "\n\t\t|"},
		{"cursor within indentation", "\t|\tab", "\n\t|ab"},
		{"after bracket", "\tif x {|", "\tif x {\n\t\t|"},
		{"between brackets", "\tf(){|}", "\tf(){\n\t\t|\n\t}"},
		{"between brackets with blanks", "\tf(){ | }", "\tf(){\n\t\t|\n\t}"},
		{"before other text", "x := []int{|1, 2}", "x := []int{\n\t|1, 2}"},
		{"empty line takes indent from above", "\tx := 1\n\n|", "\tx := 1\n\n\n\t|"},
		{"empty line after bracket above", "func f() {\n|\n}", "func f() {\n\n\t|\n}"},
		{"empty line at the top", "|\n\tx", "\n|\n\tx"},
		{"spaces unit is the smallest indentation", "a:\n  b:\n    c: [|", "a:\n  b:\n    c: [\n      |"},
		{"spaces unit of a single line", "  f(|", "  f(\n    |"},
		{"spaces unit skips tab lines", "x\n\ty\n   f(|", "x\n\ty\n   f(\n      |"},
		{"unindented line with spaces in the file", "f(|\n  x", "f(\n  |\n  x"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := editor.NewTestHarness(t)
			h.Extend(new(Integration))
			text, line, col := cutCursor(tc.text)
			if err := h.OpenReader("a.go", strings.NewReader(text)); err != nil {
				t.Fatal(err)
			}
			h.Run(t)
			h.SendInput(t, []byte("i"+strings.Repeat(down, line)+strings.Repeat(right, col)+enter))
			if got := topTextWithCursor(t, h); got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestSplitLine_Undo(t *testing.T) {
	h := editor.NewTestHarness(t)
	h.Extend(new(Integration))
	if err := h.OpenReader("a.go", strings.NewReader("\tf(){}")); err != nil {
		t.Fatal(err)
	}
	h.Run(t)
	h.SendInput(t, []byte("i"+strings.Repeat(right, 5)+enter+"\x1b"+undo))
	if got, want := topTextWithCursor(t, h), "\tf(){|}"; got != want {
		t.Errorf("after undo got %q, want %q", got, want)
	}
}

func TestSplitLine_Selection(t *testing.T) {
	h := editor.NewTestHarness(t)
	h.Extend(new(Integration))
	if err := h.OpenReader("a.go", strings.NewReader("\tabc")); err != nil {
		t.Fatal(err)
	}
	h.Run(t)
	// The editor replaces the selection with a plain line break.
	h.SendInput(t, []byte("i"+strings.Repeat(right, 3)+shiftLeft+enter))
	if got, want := topTextWithCursor(t, h), "\ta\n|c"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestIndentLines(t *testing.T) {
	for _, tc := range []struct {
		name  string
		path  string
		text  string
		input []string // each one is sent at once, the text is checked after each
		want  []string
	}{
		{
			name:  "selected lines",
			text:  "a\n\tb\nc",
			input: []string{"i" + shiftDown + shiftRight + tab, tab, shiftTab, shiftTab, shiftTab},
			want:  []string{"\ta\n\t\tb\nc", "\t\ta\n\t\t\tb\nc", "\ta\n\t\tb\nc", "a\n\tb\nc", "a\nb\nc"},
		},
		{
			name:  "part of a line",
			text:  "\tabc",
			input: []string{"i" + right + right + shiftRight + tab, shiftTab},
			want:  []string{"\t\tabc", "\tabc"},
		},
		{
			name:  "selection ending at a line start",
			text:  "a\nb",
			input: []string{"i" + shiftDown + tab},
			want:  []string{"\ta\nb"},
		},
		{
			name:  "blank lines are kept",
			text:  "a\n\n  \nb",
			input: []string{"i" + strings.Repeat(shiftDown, 3) + shiftRight + tab},
			want:  []string{"\ta\n\n  \n\tb"},
		},
		{
			name:  "smallest spaces of the file",
			path:  "a.yaml",
			text:  "a:\n  b:\n    c: 1\n  d: 2",
			input: []string{"i" + down + down + down + shiftRight + tab, shiftTab, shiftTab},
			want:  []string{"a:\n  b:\n    c: 1\n    d: 2", "a:\n  b:\n    c: 1\n  d: 2", "a:\n  b:\n    c: 1\nd: 2"},
		},
		{
			name:  "spaces of the file for an unindented line",
			path:  "a.yaml",
			text:  "a:\n  b: 1\nc: 2",
			input: []string{"i" + down + down + shiftRight + tab},
			want:  []string{"a:\n  b: 1\n  c: 2"},
		},
		{
			name:  "shift+tab takes off fewer spaces",
			path:  "a.yaml",
			text:  "a:\n  b:\n    c\n   d",
			input: []string{"i" + down + down + shiftDown + shiftRight + shiftTab},
			want:  []string{"a:\n  b:\n  c\n d"},
		},
		{
			name:  "shift+tab on the cursor line",
			text:  "\t\tx\n\ty",
			input: []string{"i" + shiftTab, shiftTab, shiftTab},
			want:  []string{"\tx\n\ty", "x\n\ty", "x\n\ty"},
		},
		{
			name:  "tab without a selection is typed",
			text:  "ab",
			input: []string{"i" + right + tab},
			want:  []string{"a\tb"},
		},
		{
			name:  "undone at once",
			text:  "a\nb",
			input: []string{"i" + shiftDown + shiftRight + tab + esc + undo},
			want:  []string{"a\nb"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := editor.NewTestHarness(t)
			h.Extend(new(Integration))
			path := tc.path
			if path == "" {
				path = "a.go"
			}
			if err := h.OpenReader(path, strings.NewReader(tc.text)); err != nil {
				t.Fatal(err)
			}
			h.Run(t)
			for i, in := range tc.input {
				h.SendInput(t, []byte(in))
				var got string
				h.Post(t, editor.CommandFunc(func(e *editor.Editor) { got = e.Top().Text() }))
				if got != tc.want[i] {
					t.Errorf("text after %q = %q, want %q", in, got, tc.want[i])
				}
			}
		})
	}
}

func TestSelectedSpans(t *testing.T) {
	span := func(l1, c1, l2, c2 int) content.Span {
		return content.Span{Start: content.Position{Line: l1, Col: c1}, End: content.Position{Line: l2, Col: c2}}
	}
	for _, tc := range []struct {
		name      string
		sel       []content.Span
		wantSpans []content.Span
		wantLines []int // visited by forLines
	}{
		{"no selection", nil, []content.Span{span(5, 0, 5, 0)}, []int{5}},
		{"one line", []content.Span{span(1, 2, 1, 3)}, []content.Span{span(1, 2, 1, 3)}, []int{1}},
		{"block", []content.Span{span(1, 2, 3, 1)}, []content.Span{span(1, 2, 3, 1)}, []int{1, 2, 3}},
		{"ending at a line start", []content.Span{span(1, 0, 3, 0)}, []content.Span{span(1, 0, 2, 0)}, []int{1, 2}},
		{
			"apart",
			[]content.Span{span(1, 0, 2, 1), span(7, 0, 8, 1)},
			[]content.Span{span(1, 0, 2, 1), span(7, 0, 8, 1)},
			[]int{1, 2, 7, 8},
		},
		{
			"sharing a line",
			[]content.Span{span(1, 0, 2, 1), span(2, 3, 3, 1)},
			[]content.Span{span(1, 0, 2, 1), span(2, 3, 3, 1)},
			[]int{1, 2, 3},
		},
		{
			"sharing the line before a line start",
			[]content.Span{span(1, 2, 3, 0), span(4, 0, 4, 1)},
			[]content.Span{span(1, 2, 2, 0), span(4, 0, 4, 1)},
			[]int{1, 2, 4},
		},
		{
			"next span on the line ending at a line start",
			[]content.Span{span(0, 0, 1, 0), span(1, 1, 1, 2)},
			[]content.Span{span(0, 0, 0, 0), span(1, 1, 1, 2)},
			[]int{0, 1},
		},
		{
			"only a line break",
			[]content.Span{span(1, 3, 2, 0)},
			[]content.Span{span(1, 3, 1, 3)},
			[]int{1},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			spans := selectedSpans(slices.Clone(tc.sel), 5)
			if !slices.Equal(spans, tc.wantSpans) {
				t.Errorf("spans %v, want %v", spans, tc.wantSpans)
			}
			lines := slices.Collect(forLines(spans))
			if !slices.Equal(lines, tc.wantLines) {
				t.Errorf("lines %v, want %v", lines, tc.wantLines)
			}
		})
	}
}

func TestIndentSpans_Apart(t *testing.T) {
	h := editor.NewTestHarness(t)
	h.Extend(new(Integration))
	if err := h.OpenReader("a.go", strings.NewReader("a\nb\n\tc\nd\ne")); err != nil {
		t.Fatal(err)
	}
	h.Run(t)
	line := func(l int) content.Span {
		return content.Span{Start: content.Position{Line: l}, End: content.Position{Line: l, Col: 1}}
	}
	indent := func(dir int) (text string) {
		h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
			buf := e.Top()
			d := buf.ExtensionData(new(Integration).ID()).(*document)
			d.indentSpans(selectedSpans([]content.Span{line(0), line(2), line(4)}, 0), dir)
			text = buf.Text()
		}))
		return text
	}
	if got, want := indent(1), "\ta\nb\n\t\tc\nd\n\te"; got != want {
		t.Errorf("indented %q, want %q", got, want)
	}
	if got, want := indent(-1), "a\nb\n\tc\nd\ne"; got != want {
		t.Errorf("unindented %q, want %q", got, want)
	}
	h.SendInput(t, []byte(undo))
	var got string
	h.Post(t, editor.CommandFunc(func(e *editor.Editor) { got = e.Top().Text() }))
	if want := "\ta\nb\n\t\tc\nd\n\te"; got != want {
		t.Errorf("after undo %q, want %q: the lines are undone at once", got, want)
	}
}

func cutCursor(s string) (text string, line, col int) {
	i := strings.Index(s, "|")
	before := s[:i]
	line = strings.Count(before, "\n")
	col = len(before) - strings.LastIndex(before, "\n") - 1
	return before + s[i+1:], line, col
}

func topTextWithCursor(t *testing.T, h *editor.TestHarness) (text string) {
	t.Helper()
	h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
		buf := e.Top()
		cx, cy := buf.Pos()
		lines := strings.Split(buf.Text(), "\n")
		lines[cy] = lines[cy][:cx] + "|" + lines[cy][cx:]
		text = strings.Join(lines, "\n")
	}))
	return text
}
