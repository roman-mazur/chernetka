package editor

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/change"
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

func (ae actionsEditorExt) ID() string                               { return "actions" }
func (ae actionsEditorExt) AfterEdit(Sender, *Buffer, []change.Item) {}
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
		// The second paste starts in the read of the first one. It's not typed in the review mode.
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

	// Getting the focus back does not scroll the text: the text is scrolled by the mouse
	// while the terminal is in the background, often leaving the cursor out of the screen.
	// With the focus brought by a key combination, the cursor stays where it was, and the
	// next key reveals it. With the focus brought by a click, the cursor moves to the clicked
	// position. The property is checked with random texts, scrolls, and clicks.
	t.Run("focus in keeps the scroll", func(t *testing.T) {
		const (
			focusIn  = "\x1b[I"
			focusOut = "\x1b[O"
		)
		for seed := range uint64(50) {
			rnd := rand.New(rand.NewPCG(seed, 0))
			lines := 1 + rnd.IntN(300)
			cursorLine := 1 + rnd.IntN(lines)
			scrolls := rnd.IntN(2 * lines)
			scrollDown := rnd.IntN(4) > 0
			byClick := rnd.IntN(2) == 0
			row, col := 1+rnd.IntN(45), 1+rnd.IntN(80) // possibly over the status bar or past the screen
			name := fmt.Sprintf("seed=%d/lines=%d/cursor=%d/scrolls=%d/down=%t/click=%t", seed, lines, cursorLine, scrolls, scrollDown, byClick)

			t.Run(name, func(t *testing.T) {
				h := NewTestHarness(t)
				if err := h.OpenReader("", strings.NewReader(strings.Repeat("some text\n", lines))); err != nil {
					t.Fatal(err)
				}
				h.Run(t)

				// Scroll the text in the background.
				h.SendInput(t, fmt.Appendf(nil, ":%d\r", cursorLine))
				h.SendInput(t, []byte(focusOut))
				wheel := "\x1b[<64;5;5M"
				if scrollDown {
					wheel = "\x1b[<65;5;5M"
				}
				h.SendInput(t, []byte(strings.Repeat(wheel, scrolls)))

				type view struct {
					offset, xoff int
					c            content.Position
				}
				var before view
				var height int
				var clicked content.Position // where the click lands with the scroll kept
				h.Post(t, CommandFunc(func(e *Editor) {
					b := e.Top()
					before, height = view{b.offset, b.xoff, b.c()}, b.h
					clicked = b.screenToContentPosition(row-1, col-1, e.rPrefs.TabSize)
				}))

				want := before
				if byClick {
					h.SendInput(t, fmt.Appendf(nil, "%s\x1b[<0;%d;%dM\x1b[<0;%d;%dm", focusIn, col, row, col, row))
					want.c = clicked
				} else {
					h.SendInput(t, []byte(focusIn))
				}

				var after view
				h.Post(t, CommandFunc(func(e *Editor) {
					b := e.Top()
					after = view{b.offset, b.xoff, b.c()}
				}))
				if after != want {
					t.Fatalf("after the focus in: %+v, want %+v", after, want)
				}

				// The next key reveals the cursor.
				h.SendInput(t, []byte("l"))
				h.Post(t, CommandFunc(func(e *Editor) {
					b := e.Top()
					if ln := b.c().Line; ln < b.offset || ln >= b.offset+height {
						t.Errorf("cursor line %d is out of the screen after a key: offset %d, height %d", ln, b.offset, height)
					}
				}))
			})
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
	shiftPress := func(col, line int) string {
		return input.Press(left, col+3, line+1).With(input.ModShift).Encode()
	}
	shiftRelease := func(col, line int) string {
		return input.Release(left, col+3, line+1).With(input.ModShift).Encode()
	}
	shiftClick := func(col, line int) string {
		return shiftPress(col, line) + shiftRelease(col, line)
	}
	shiftDrag := func(col, line int) string {
		return input.Drag(left, col+3, line+1).With(input.ModShift).Encode()
	}

	for _, tc := range []struct {
		name   string
		text   string
		inputs []string
		want   string
	}{
		{
			name:   "shift+click selects forward from the cursor",
			text:   "hello world\nsecond line",
			inputs: []string{"i", click(2, 0), shiftClick(3, 1), "x"},
			want:   "hexond line",
		},
		{
			name:   "shift+click selects backward from the cursor",
			text:   "hello world\nsecond line",
			inputs: []string{"i", click(3, 1), shiftClick(2, 0), "x"},
			want:   "hexond line",
		},
		{
			name:   "shift+click again moves the end of the selection",
			text:   "hello world",
			inputs: []string{"i", click(2, 0), shiftClick(4, 0), shiftClick(6, 0), "x"},
			want:   "hexworld",
		},
		{
			name:   "shift+click extends the keyboard selection",
			text:   "hello world",
			inputs: []string{"i" + shiftRight + shiftRight, shiftClick(5, 0), "x"},
			want:   "x world",
		},
		{
			name:   "shift+click extends the selected word",
			text:   "hello world",
			inputs: []string{"i", click(1, 0) + click(1, 0), shiftClick(8, 0), "x"},
			want:   "xrld",
		},
		{
			name:   "shift+drag continues from the shift+click",
			text:   "hello world",
			inputs: []string{"i", click(2, 0), shiftPress(5, 0) + shiftDrag(6, 0) + shiftDrag(7, 0) + shiftRelease(7, 0), "x"},
			want:   "hexorld",
		},
		{
			name:   "click after shift+click clears the selection",
			text:   "hello world",
			inputs: []string{"i", click(2, 0), shiftClick(4, 0), click(6, 0), "x"},
			want:   "hello xworld",
		},
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
			name:   "delete a cluster under the cursor",
			text:   "a" + testFlag + "b",
			inputs: []string{"lx"},
			want:   "ab",
		},
		{
			name:   "append after a cluster",
			text:   testMermaid + "b",
			inputs: []string{"aX"},
			want:   testMermaid + "Xb",
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

func TestPasteText_CursorStaysAfterJoinedText(t *testing.T) {
	ft := content.FullText{content.TextLine("\u0301z")}
	buf := &Buffer{Content: &ft}
	PasteText("\u0430").DoOnBuffer(buf, RenderPrefs{TabSize: 4})
	if got, want := buf.c(), pos(0, len(testAccent)); got != want {
		t.Errorf("cursor at %s, want %s", got, want)
	}
}

func TestPasteText_LineBreaks(t *testing.T) {
	for _, tc := range []struct{ name, text string }{
		{name: "LF", text: "one\ntwo\n"},
		{name: "CRLF", text: "one\r\ntwo\r\n"},
		{name: "CR", text: "one\rtwo\r"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ft := content.FullText{content.TextLine("x")}
			buf := &Buffer{Content: &ft}
			PasteText(tc.text).DoOnBuffer(buf, RenderPrefs{TabSize: 4})
			if got := buf.Text(); got != "one\ntwo\nx" {
				t.Errorf("text %q", got)
			}
			if buf.c() != pos(2, 0) {
				t.Errorf("cursor at %s", buf.c())
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

func TestDragSelectionScrolls(t *testing.T) {
	const left = input.MouseButtonLeft
	newEditor := func(t *testing.T) (*Editor, *Buffer) {
		t.Helper()
		e := &Editor{cmdChannel: make(chan Command, 16)}
		var lines []string
		for i := range 30 {
			lines = append(lines, fmt.Sprintf("line %d", i))
		}
		if err := e.OpenReader("", strings.NewReader(strings.Join(lines, "\n"))); err != nil {
			t.Fatal(err)
		}
		b := e.Top()
		b.w, b.h = 40, 10
		t.Cleanup(e.stopDragScroll)
		return e, b
	}
	// waitTick runs the next queued scroll tick.
	waitTick := func(t *testing.T, e *Editor) {
		t.Helper()
		select {
		case cmd := <-e.cmdChannel:
			cmd.DoOnEditor(e)
		case <-time.After(time.Second):
			t.Fatal("no scroll tick")
		}
	}
	selection := func(b *Buffer) content.Span { return b.sel[len(b.sel)-1] }

	t.Run("down past the bottom edge", func(t *testing.T) {
		e, b := newEditor(t)
		sendChunks(t, e, input.Press(left, 5, 3).Encode()+input.Drag(left, 6, 3).Encode())
		// Over the status bar, the row below the buffer.
		sendChunks(t, e, input.Drag(left, 6, 11).Encode())
		if b.offset != 1 {
			t.Errorf("offset %d, want 1", b.offset)
		}
		if got := selection(b); got.Start.Line != 2 || got.End.Line != 10 {
			t.Errorf("selection %v, want lines 2-10", got)
		}

		// The text keeps scrolling while the mouse stays there.
		waitTick(t, e)
		waitTick(t, e)
		if b.offset != 3 {
			t.Errorf("offset %d after ticks, want 3", b.offset)
		}
		if got := selection(b).End.Line; got != 12 {
			t.Errorf("selection end line %d, want 12", got)
		}

		// Further down, it's faster, and it stops at the end of the text.
		sendChunks(t, e, input.Drag(left, 6, 40).Encode())
		if b.offset != 20 {
			t.Errorf("offset %d, want 20", b.offset)
		}
		waitTick(t, e)
		if e.dragScrolling {
			t.Error("still scrolling at the end of the text")
		}
		sendChunks(t, e, input.Release(left, 6, 40).Encode())
		if got := selection(b); got.Start.Line != 2 || got.End.Line != 29 {
			t.Errorf("selection %v, want lines 2-29", got)
		}
	})

	t.Run("up at the top row", func(t *testing.T) {
		e, b := newEditor(t)
		b.offset = 10
		sendChunks(t, e, input.Press(left, 5, 5).Encode()+input.Drag(left, 6, 5).Encode())
		sendChunks(t, e, input.Drag(left, 6, 1).Encode())
		if b.offset != 9 {
			t.Errorf("offset %d, want 9", b.offset)
		}
		waitTick(t, e)
		if b.offset != 8 {
			t.Errorf("offset %d after a tick, want 8", b.offset)
		}
		if got := selection(b); got.Start.Line != 14 || got.End.Line != 8 {
			t.Errorf("selection %v, want lines 14-8", got)
		}
	})

	t.Run("stops when back inside or released", func(t *testing.T) {
		e, b := newEditor(t)
		sendChunks(t, e, input.Press(left, 5, 3).Encode()+input.Drag(left, 6, 3).Encode()+input.Drag(left, 6, 12).Encode())
		if !e.dragScrolling {
			t.Fatal("not scrolling")
		}
		sendChunks(t, e, input.Drag(left, 6, 5).Encode())
		if e.dragScrolling {
			t.Error("still scrolling inside the buffer")
		}
		offset := b.offset
		sendChunks(t, e, input.Drag(left, 6, 12).Encode()+input.Release(left, 6, 12).Encode())
		if e.dragScrolling {
			t.Error("still scrolling after the release")
		}
		time.Sleep(2 * mouseDragScrollInterval)
		for len(e.cmdChannel) > 0 {
			(<-e.cmdChannel).DoOnEditor(e)
		}
		if b.offset != offset+2 {
			t.Errorf("offset %d, want %d", b.offset, offset+2)
		}
	})
}
