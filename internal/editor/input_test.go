package editor

import (
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

type actionsEditorExt []int

func (ae actionsEditorExt) ID() string                { return "actions" }
func (ae actionsEditorExt) AfterEdit(Sender, *Buffer) {}
func (ae actionsEditorExt) MakeBufferData(Sender, *Buffer) BufferExtData {
	res := make(testActionsExt)
	for _, ln := range ae {
		res[ln] = new(countingAction)
	}
	return res
}

func TestEditor_SendInput(t *testing.T) {
	newEditor := func(t *testing.T, actionLines ...int) *Editor {
		e := &Editor{cmdChannel: make(chan Command, 16)}
		e.Extend(actionsEditorExt(actionLines))
		if err := e.OpenReader("", strings.NewReader("one\ntwo\nthree\nfour")); err != nil {
			t.Fatal(err)
		}
		return e
	}

	t.Run("keys read at once", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "\x1b[B\x1b[B", "jl")
		if got := e.Top().c(); got != pos(3, 1) {
			t.Errorf("cursor at %s", got)
		}
	})

	t.Run("split sequence", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "\x1b[1;", "2B")
		if b := e.Top(); b.c() != pos(1, 0) || len(b.sel) == 0 {
			t.Errorf("cursor at %s, selection %v", b.c(), b.sel)
		}
	})

	t.Run("split mouse event", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "j\x1b[<0;3", ";1Mi\x1b[<0;3;1m")
		if e.quitRequested {
			t.Fatal("quit on a split mouse event")
		}
		// The click at the first row moves the cursor back from the second line.
		if b := e.Top(); b.mode != ModeInsert || b.c().Line != 0 {
			t.Errorf("mode %s, cursor at %s after the click", b.mode, b.c())
		}
	})

	t.Run("paste", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "i\x1b[200~x\x1b[201~")
		if got := e.Top().Content.Lines()[0].String(); got != "xone" {
			t.Errorf("line %q", got)
		}
	})

	t.Run("keys after paste", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "i\x1b[200~x\x1b[201~yz")
		if got := e.Top().Content.Lines()[0].String(); got != "xyzone" {
			t.Errorf("line %q", got)
		}
	})

	t.Run("paste after paste", func(t *testing.T) {
		e := newEditor(t)
		// The second paste starts in the read of the first one. It's not typed in the normal mode.
		sendChunks(t, e, "\x1b[200~x\x1b[201~\x1b[20", "0~qj\x1b[201~")
		if e.quitRequested {
			t.Fatal("quit on a paste")
		}
		if got := e.Top().Content.Lines()[0].String(); got != "xqjone" {
			t.Errorf("line %q", got)
		}
	})

	t.Run("keys after quit", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "qj")
		if !e.quitRequested || e.Top().c() != pos(0, 0) {
			t.Errorf("quit %t, cursor at %s", e.quitRequested, e.Top().c())
		}
	})

	t.Run("engage action", func(t *testing.T) {
		e := newEditor(t, 1)
		sendChunks(t, e, "\r")

		buf := e.Top()
		if buf.c().Line != 1 {
			t.Fatalf("Enter on a plain line moved the cursor to %d, want 1", buf.c().Line)
		}

		sendChunks(t, e, "\r")
		if got := buf.lineAction(1).(*countingAction).engaged; got != 1 {
			t.Errorf("action engaged %d times, want 1", got)
		}
		if buf.c().Line != 1 {
			t.Errorf("Enter on an actionable line moved the cursor to %d", buf.c().Line)
		}
	})
}

func TestEditor_DoubleClick(t *testing.T) {
	for _, tc := range []struct {
		name         string
		engage       bool // engageOnDoubleClick of the buffer
		line         int
		wantEngaged  int
		wantSelected bool
	}{
		{name: "engages action", engage: true, line: 1, wantEngaged: 1},
		{name: "selects word without action", engage: true, line: 2, wantSelected: true},
		{name: "selects word in text buffer", engage: false, line: 1, wantSelected: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			e := &Editor{cmdChannel: make(chan Command, 16)}
			e.Extend(actionsEditorExt{1})
			if err := e.OpenReader("", strings.NewReader("one\ntwo\nthree\nfour")); err != nil {
				t.Fatal(err)
			}
			buf := e.Top()
			buf.engageOnDoubleClick = tc.engage

			// Screen coordinates are 1-based: click over the second character of the line.
			row, col := tc.line+1, buf.lineNumberPrefixWidth()+2
			click := input.Click(input.MouseButtonLeft, col, row)
			sendChunks(t, e, click, click)

			action := buf.lineAction(1).(*countingAction)
			if action.engaged != tc.wantEngaged {
				t.Errorf("action engaged %d times, want %d", action.engaged, tc.wantEngaged)
			}
			if tc.wantEngaged > 0 && e.lastAction.action != action {
				t.Errorf("last action %v, want the engaged one", e.lastAction)
			}
			if got := len(buf.sel) > 0; got != tc.wantSelected {
				t.Errorf("selection %v, want selected %t", buf.sel, tc.wantSelected)
			}
			if buf.c().Line != tc.line {
				t.Errorf("cursor at %s, want line %d", buf.c(), tc.line)
			}
		})
	}
}

func TestEditor_SelectionEdits(t *testing.T) {
	const (
		shiftDown  = "\x1b[1;2B"
		shiftRight = "\x1b[1;2C"
		shiftLeft  = "\x1b[1;2D"
		backspace  = "\x7f"
		left       = input.MouseButtonLeft
	)
	// click returns the mouse press and release at the 0-based text column and line.
	// The line numbers take 2 columns.
	click := func(col, line int) string {
		return input.Click(left, col+3, line+1)
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
			name: "drag over the line numbers",
			text: "one\ntwo\nthree",
			inputs: []string{
				input.Press(left, 6, 2).Encode() + input.Drag(left, 4, 2).Encode() +
					input.Drag(left, 1, 1).Encode() + input.Release(left, 1, 1).Encode(),
				"x",
			},
			want: "\nthree",
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
	if buf.c() != pos(0, 2) || len(buf.sel) != 0 {
		t.Errorf("cursor at %s, selection %v", buf.c(), buf.sel)
	}
}
