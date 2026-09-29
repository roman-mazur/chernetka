package editor

import (
	"path/filepath"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/input"
)

// definitionExt gives every buffer a DefinitionFinder recording the positions it's asked about.
type definitionExt struct {
	asked []content.Position
}

func (de *definitionExt) ID() string                           { return "definition" }
func (de *definitionExt) MakeBufferData(*Buffer) BufferExtData { return definitionFinder{de} }
func (de *definitionExt) AfterEdit(Sender, *Buffer)            {}

type definitionFinder struct{ *definitionExt }

func (df definitionFinder) FindDefinition(_ Sender, pos content.Position) {
	df.asked = append(df.asked, pos)
}

func TestEditor_CtrlClickFindsDefinition(t *testing.T) {
	var ext definitionExt
	h := NewTestHarness()
	h.Extend(&ext)
	if err := h.OpenReader("a.go", strings.NewReader("package a\n\nfunc foo() {}\n\nvar x = foo()\n")); err != nil {
		t.Fatal(err)
	}
	h.Run(t)

	// Screen coordinates are 1-based, the click is on the last line over "foo".
	const row = 5
	var col int
	h.Post(t, CommandFunc(func(e *Editor) { col = e.Top().lineNumberPrefixWidth() + len("var x = f") }))
	want := content.Position{Line: 4, Col: len("var x = f") - 1}

	click := func(button input.MouseButton, mod input.Modifier) {
		press, release := input.Press(button, col, row).With(mod), input.Release(button, col, row).With(mod)
		h.SendInput(t, []byte(press.Encode()+release.Encode()))
	}
	askedCount := func() (n int) {
		h.Post(t, CommandFunc(func(*Editor) { n = len(ext.asked) }))
		return n
	}

	click(input.MouseButtonLeft, 0) // A plain click only moves the cursor.
	if n := askedCount(); n != 0 {
		t.Fatalf("definition asked on a plain click: %v", ext.asked)
	}

	click(input.MouseButtonLeft, input.ModCtrl)  // Ctrl+left.
	click(input.MouseButtonRight, input.ModCtrl) // Ctrl+right, as some terminals report Ctrl+click on macOS.
	if n := askedCount(); n != 2 {
		t.Fatalf("definition asked %d times, want 2", n)
	}
	for _, pos := range ext.asked {
		if pos != want {
			t.Errorf("definition asked at %v, want %v", pos, want)
		}
	}
	h.Post(t, CommandFunc(func(e *Editor) {
		if e.Top().c != want {
			t.Errorf("cursor at %v, want %v", e.Top().c, want)
		}
	}))
}

func TestEditor_GoTo(t *testing.T) {
	dir := t.TempDir()
	a, b := filepath.Join(dir, "a.txt"), filepath.Join(dir, "b.txt")
	writeFile(t, a, "a\n")
	writeFile(t, b, strings.Repeat("line\n", 100))

	h := NewTestHarness()
	h.Run(t)
	h.Post(t, &OpenFile{Path: a})

	pos := content.Position{Line: 60, Col: 2}
	h.Post(t, CommandFunc(func(e *Editor) {
		// Check the scroll before the loop renders the buffer.
		(&GoTo{Path: b, Pos: pos}).DoOnEditor(e)
		buf := e.Top()
		if buf.Path != b {
			t.Fatalf("top buffer is %s, want %s", buf.Path, b)
		}
		if buf.c != pos {
			t.Errorf("cursor at %v, want %v", buf.c, pos)
		}
		buf.h = 30
		buf.clampCursor(e.rPrefs.TabSize)
		if buf.offset != 50 {
			t.Errorf("offset = %d, want the line in the upper third of the screen", buf.offset)
		}
	}))

	// An open buffer is activated.
	h.Post(t, &GoTo{Path: a, Pos: content.Position{Col: 1}})
	h.Post(t, CommandFunc(func(e *Editor) {
		if got := len(e.bufs); got != 2 {
			t.Errorf("%d buffers open, want 2", got)
		}
		if buf := e.Top(); buf.Path != a || buf.c != (content.Position{Col: 1}) {
			t.Errorf("top buffer %s at %v, want %s at 0:1", buf.Path, buf.c, a)
		}
	}))

	// The harness quits with one buffer left.
	h.Post(t, CommandFunc(func(e *Editor) { e.pop() }))
}
