package editor

import (
	"fmt"
	"slices"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/input"
)

// handleInput handles the keys decoded from the terminal input like the editor loop does.
func (e *Editor) handleInput(b []byte) (quit bool) {
	keys, _ := input.Keys(b)
	return slices.ContainsFunc(keys, e.handleKey)
}

// sendChunks passes the chunks to sendInput like they are read from the terminal,
// and runs the commands sent to the editor loop.
func sendChunks(t *testing.T, e *Editor, chunks ...string) {
	t.Helper()
	var pending []byte
	for _, chunk := range chunks {
		var err error
		pending, err = e.sendInput(append(pending, chunk...), strings.NewReader(""))
		if err != nil {
			t.Fatal(err)
		}
		for len(e.cmdChannel) > 0 {
			(<-e.cmdChannel).DoOnEditor(e)
		}
	}
	if len(pending) > 0 {
		t.Errorf("pending input %q", pending)
	}
}

func TestEditor_SendInput(t *testing.T) {
	newEditor := func(t *testing.T) *Editor {
		e := &Editor{cmdChannel: make(chan Command, 16)}
		if err := e.OpenReader("", strings.NewReader("one\ntwo\nthree\nfour")); err != nil {
			t.Fatal(err)
		}
		return e
	}

	t.Run("keys read at once", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "\x1b[B\x1b[B", "jl")
		if got := e.Top().c; got != pos(3, 1) {
			t.Errorf("cursor at %s", got)
		}
	})

	t.Run("split sequence", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "\x1b[1;", "2B")
		if b := e.Top(); b.c != pos(1, 0) || len(b.sel) == 0 {
			t.Errorf("cursor at %s, selection %v", b.c, b.sel)
		}
	})

	t.Run("split mouse event", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "j\x1b[<0;3", ";1Mi\x1b[<0;3;1m")
		if e.quitRequested {
			t.Fatal("quit on a split mouse event")
		}
		// The click at the first row moves the cursor back from the second line.
		if b := e.Top(); b.mode != ModeInsert || b.c.Line != 0 {
			t.Errorf("mode %s, cursor at %s after the click", b.mode, b.c)
		}
	})

	t.Run("paste", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "i\x1b[200~x\x1b[201~")
		if got := e.Top().Content.Lines()[0].String(); got != "xone" {
			t.Errorf("line %q", got)
		}
	})

	t.Run("keys after quit", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "qj")
		if !e.quitRequested || e.Top().c != pos(0, 0) {
			t.Errorf("quit %t, cursor at %s", e.quitRequested, e.Top().c)
		}
	})
}

func TestEditor_SelectionEdits(t *testing.T) {
	const (
		shiftDown  = "\x1b[1;2B"
		shiftRight = "\x1b[1;2C"
		shiftLeft  = "\x1b[1;2D"
		backspace  = "\x7f"
	)
	// click returns the mouse press and release at the 0-based text column and line.
	// The line numbers take 2 columns.
	click := func(col, line int) string {
		return fmt.Sprintf("\x1b[<0;%d;%dM\x1b[<0;%[1]d;%[2]dm", col+3, line+1)
	}

	for _, tc := range []struct {
		name   string
		text   string
		inputs []string
		want   string
	}{
		{
			name:   "select down to a shorter line",
			text:   "hello world\nab\nlast",
			inputs: []string{"$i" + shiftDown, backspace},
			want:   "hello world\nlast",
		},
		{
			name:   "select down past the last line",
			text:   "one\ntwo",
			inputs: []string{"i" + shiftDown + shiftDown + shiftDown, "x"},
			want:   "xtwo",
		},
		{
			name:   "backspace with an empty selection",
			text:   "abc",
			inputs: []string{"$i" + shiftLeft + shiftRight, backspace},
			want:   "ab",
		},
		{
			name:   "select a line",
			text:   "one\ntwo\nthree",
			inputs: []string{"j", click(1, 1) + click(1, 1) + click(1, 1), "x"},
			want:   "one\nthree",
		},
		{
			name:   "select the last line",
			text:   "one\ntwo",
			inputs: []string{"j", click(1, 1) + click(1, 1) + click(1, 1), "x"},
			want:   "one\n",
		},
		{
			name:   "click after selecting a word",
			text:   "one two",
			inputs: []string{click(1, 0) + click(1, 0), click(5, 0), "iX"},
			want:   "one tXwo",
		},
		{
			name:   "extend a selected word",
			text:   "one two",
			inputs: []string{click(1, 0) + click(1, 0), shiftRight + shiftRight, "x"},
			want:   "wo",
		},
		{
			name:   "drag over the line numbers",
			text:   "one\ntwo\nthree",
			inputs: []string{"\x1b[<0;6;2M\x1b[<32;4;2M\x1b[<32;1;1M\x1b[<0;1;1m", "x"},
			want:   "\nthree",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			e := &Editor{cmdChannel: make(chan Command, 16)}
			if err := e.OpenReader("", strings.NewReader(tc.text)); err != nil {
				t.Fatal(err)
			}
			b := e.Top()
			b.w, b.h = 40, 10
			now := time.Now()
			e.now = func() time.Time { return now }
			for _, in := range tc.inputs {
				now = now.Add(time.Second) // Separate inputs are not double clicks.
				sendChunks(t, e, in)
				b.clampCursor(e.rPrefs.TabSize) // Like the render between the inputs.
			}
			if got := b.Text(); got != tc.want {
				t.Errorf("text %q, want %q", got, tc.want)
			}
		})
	}
}

func TestPasteText_ReplacesSelection(t *testing.T) {
	ft := content.FullText{content.TextLine("one"), content.TextLine("two")}
	buf := &Buffer{Content: &ft, sel: []content.Span{{Start: pos(0, 1), End: pos(1, 1)}}}
	PasteText("X").DoOnBuffer(buf, RenderPrefs{TabSize: 4})
	if got := buf.Text(); got != "oXwo" {
		t.Errorf("text %q", got)
	}
	if buf.c != pos(0, 2) || len(buf.sel) != 0 {
		t.Errorf("cursor at %s, selection %v", buf.c, buf.sel)
	}
}
