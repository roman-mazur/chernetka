package extindent

import (
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/editor"
)

const (
	down      = "\x1b[B"
	right     = "\x1b[C"
	shiftLeft = "\x1b[1;2D"
	enter     = "\r"
	undo      = "\x1a"
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
		{"spaces unit from the line above", "a:\n  b: [|", "a:\n  b: [\n    |"},
		{"spaces unit without lines above", "  f(|", "  f(\n    |"},
		{"spaces unit skips tab lines", "x\n\ty\n   f(|", "x\n\ty\n   f(\n      |"},
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
