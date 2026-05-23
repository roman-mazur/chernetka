package editor

import (
	"strings"
	"testing"
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
	if bufs[0].path != wantPath {
		t.Errorf("layout top path = %q, want %q", bufs[0].path, wantPath)
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
