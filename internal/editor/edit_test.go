package editor

import (
	"os"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
)

func TestEditor_OpenReader(t *testing.T) {
	var e Editor

	if err := e.OpenReader("first.txt", strings.NewReader("hello\nworld")); err != nil {
		t.Fatalf("OpenReader: %v", err)
	}
	assertLayoutTop(t, &e, "first.txt")
	assertBufferCount(t, &e, 1)

	if err := e.OpenReader("second.txt", strings.NewReader("foo")); err != nil {
		t.Fatalf("OpenReader: %v", err)
	}
	assertLayoutTop(t, &e, "second.txt")
	assertBufferCount(t, &e, 2)
}

func TestEditor_OpenReader_ReuseExisting(t *testing.T) {
	var e Editor

	for _, name := range []string{"a.txt", "b.txt", "c.txt"} {
		if err := e.OpenReader(name, strings.NewReader(name)); err != nil {
			t.Fatalf("OpenReader(%q): %v", name, err)
		}
	}
	// Stack top→bottom: c.txt, b.txt, a.txt
	assertLayoutTop(t, &e, "c.txt")
	assertBufferCount(t, &e, 3)

	// Re-open a non-top buffer — should move to top, not create a new entry.
	if err := e.OpenReader("a.txt", strings.NewReader("ignored")); err != nil {
		t.Fatalf("OpenReader(a.txt): %v", err)
	}
	assertLayoutTop(t, &e, "a.txt")
	assertBufferCount(t, &e, 3)

	// Re-opening the already-top buffer is a no-op.
	if err := e.OpenReader("a.txt", strings.NewReader("ignored")); err != nil {
		t.Fatalf("OpenReader(a.txt) no-op: %v", err)
	}
	assertLayoutTop(t, &e, "a.txt")
	assertBufferCount(t, &e, 3)
}

func assertLayoutTop(t *testing.T, e *Editor, wantPath string) {
	t.Helper()
	var bufs []*Buffer
	for buf := range e.layout() {
		bufs = append(bufs, buf)
	}
	if len(bufs) != 1 {
		t.Fatalf("layout returned %d buffers, want 1", len(bufs))
	}
	if bufs[0].Path != wantPath {
		t.Errorf("layout top path = %q, want %q", bufs[0].Path, wantPath)
	}
}

// assertBufferCount walks the stack via next pointers and fails if the count
// doesn't match or exceeds a safety cap (which would indicate a cycle).
func assertBufferCount(t *testing.T, e *Editor, want int) {
	t.Helper()
	const maxSafe = 1000
	n := 0
	for range e.buffers() {
		n++
		if n > maxSafe {
			t.Fatalf("buffer count exceeds %d — likely a cycle in the stack", maxSafe)
		}
	}
	if n != want {
		t.Errorf("buffer count = %d, want %d", n, want)
	}
}

func TestEditor_Run(t *testing.T) {
	input, err := os.Open("edit.go")
	if err != nil {
		t.Fatalf("cannot open test input: %s", err)
	}
	t.Cleanup(func() { _ = input.Close() })

	te := NewTestHarness()
	if err := te.OpenReader("test.txt", input); err != nil {
		t.Fatalf("OpenReader: %s", err)
	}
	te.Run(t)

	t.Run("scrolling", func(t *testing.T) {
		te.SendInput(t, []byte("\x1b[B")) // Cursor down.
		var cursorY int
		te.Post(t, CommandFunc(func(e *Editor) {
			cursorY = e.Top().c.Line
		}))
		if cursorY != 1 {
			t.Errorf("cursor y = %d, want 1 after cursor down", cursorY)
		}

		for range 5 {
			te.SendInput(t, []byte("\x1b[<65;10;20M")) // Scroll down.
		}
		var offset int
		te.Post(t, CommandFunc(func(e *Editor) {
			offset = e.Top().offset
		}))
		if offset != 5 {
			t.Errorf("offset = %d, some scroll down events seem to be ignored", offset)
		}
	})

	t.Run("clipboard paste", func(t *testing.T) {
		// Land the cursor at a known position for a deterministic assertion.
		te.Post(t, CommandFunc(func(e *Editor) {
			e.Top().c = content.Position{}
		}))

		te.SendInput(t, []byte("\x1b[200~pasted \x1b[201~"))

		var line0 string
		var cursor content.Position
		te.Post(t, CommandFunc(func(e *Editor) {
			line0 = e.Top().Content.Lines()[0].String()
			cursor = e.Top().c
		}))
		if want := "pasted "; !strings.HasPrefix(line0, want) {
			t.Errorf("line 0 = %q, want prefix %q", line0, want)
		}
		if want := (content.Position{Col: len("pasted ")}); cursor != want {
			t.Errorf("cursor after paste = %+v, want %+v", cursor, want)
		}
	})

	t.Run("clipboard paste split across reads", func(t *testing.T) {
		te.Post(t, CommandFunc(func(e *Editor) {
			e.Top().c = content.Position{}
		}))

		// Split both the pasted text and the terminator marker across
		// separate writes, exercising the reader continuation that
		// readAndHandleInput relies on to reassemble a paste that
		// straddles two terminal reads.
		te.SendInput(t, []byte("\x1b[200~spl"))
		te.SendInput(t, []byte("it\x1b[2"))
		te.SendInput(t, []byte("01~"))

		var line0 string
		te.Post(t, CommandFunc(func(e *Editor) {
			line0 = e.Top().Content.Lines()[0].String()
		}))
		if want := "split"; !strings.HasPrefix(line0, want) {
			t.Errorf("line 0 = %q, want prefix %q", line0, want)
		}
	})

	t.Run("empty clipboard paste is a no-op", func(t *testing.T) {
		var before string
		te.Post(t, CommandFunc(func(e *Editor) {
			before = e.Top().Text()
		}))

		te.SendInput(t, []byte("\x1b[200~\x1b[201~"))

		var after string
		te.Post(t, CommandFunc(func(e *Editor) {
			after = e.Top().Text()
		}))
		if before != after {
			t.Errorf("buffer changed after an empty paste:\nbefore: %q\nafter:  %q", before, after)
		}
	})
}
