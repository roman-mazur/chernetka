package editor

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/content"
)

func TestEditor_WatchFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "a.txt")
	writeFile(t, path, "first")

	h := NewTestHarness()
	h.Run(t)
	h.Post(t, &OpenFile{Path: path})

	t.Run("write", func(t *testing.T) {
		writeFile(t, path, "second\nline")
		waitForText(t, h, "second\nline")
	})

	t.Run("replace", func(t *testing.T) {
		tmp := path + ".tmp"
		writeFile(t, tmp, "third")
		if err := os.Rename(tmp, path); err != nil {
			t.Fatal(err)
		}
		waitForText(t, h, "third")
	})

	t.Run("unsaved changes", func(t *testing.T) {
		h.Post(t, CommandFunc(func(e *Editor) {
			e.Top().Mutate().Insert(0, content.TextLine("unsaved"))
		}))
		writeFile(t, path, "fourth")
		waitForText(t, h, "fourth")
		h.Post(t, CommandFunc(func(e *Editor) {
			if e.Top().dirty {
				t.Error("the reloaded buffer is dirty")
			}
		}))
	})
}

func TestEditor_WatchFile_OwnSave(t *testing.T) {
	path := filepath.Join(t.TempDir(), "a.txt")
	writeFile(t, path, "first")

	h := NewTestHarness()
	h.Run(t)
	h.Post(t, &OpenFile{Path: path})

	h.Post(t, CommandFunc(func(e *Editor) {
		buf := e.Top()
		buf.Mutate().Update(0, content.TextLine("saved"))
		(&Save{DstPath: path}).DoOnBuffer(buf, e.rPrefs)
		// Typed right after the save, before the watcher reports it.
		buf.Mutate().Update(0, content.TextLine("typed"))
	}))

	time.Sleep(10 * fileChangeDelay)
	h.Post(t, CommandFunc(func(e *Editor) {
		if got := e.Top().Text(); got != "typed" {
			t.Errorf("the own save replaced the content: got %q", got)
		}
		if !e.Top().dirty {
			t.Error("the buffer with unsaved changes is not dirty")
		}
	}))
}

func TestEditor_ReloadShorter(t *testing.T) {
	path := filepath.Join(t.TempDir(), "a.txt")
	writeFile(t, path, strings.Repeat("line\n", 100))

	var e Editor
	(&OpenFile{Path: path}).DoOnEditor(&e)
	defer func() { _ = e.watcher.Close() }()
	buf := e.Top()
	buf.h = 10
	buf.updateCursor(content.Position{Line: 99, Col: 4})
	buf.offset = 95
	buf.noKeyboard = true // Scrolled by the mouse.

	e.reloadBuffer(buf, []byte("a\nb"))
	MoveEnd.DoOnBuffer(buf, e.rPrefs)
	if buf.c() != (content.Position{Line: 1, Col: 1}) {
		t.Errorf("cursor is %v after the reload", buf.c())
	}
	buf.clampCursor(e.rPrefs.TabSize)
	var out strings.Builder
	buf.Render(&out, &e.rPrefs)
	if buf.offset > 1 {
		t.Errorf("offset is %d after the reload", buf.offset)
	}
}

func writeFile(t *testing.T, path, text string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(text), 0600); err != nil {
		t.Fatal(err)
	}
}

func waitForText(t *testing.T, h *TestHarness, want string) {
	t.Helper()
	var got string
	for range 100 {
		h.Post(t, CommandFunc(func(e *Editor) { got = e.Top().Text() }))
		if got == want {
			return
		}
		time.Sleep(fileChangeDelay)
	}
	t.Fatalf("buffer content is not reloaded: got %q, want %q", got, want)
}

func TestEditor_WatchDir(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "a.txt"), "a")

	h := NewTestHarness()
	h.Run(t)
	h.Post(t, CommandFunc(func(e *Editor) { e.OpenDir(dir, nil) }))

	writeFile(t, filepath.Join(dir, "b.txt"), "b")
	var text string
	for range 100 {
		h.Post(t, CommandFunc(func(e *Editor) { text = textOf(e.Top().Content) }))
		if strings.Contains(text, "b.txt") {
			return
		}
		time.Sleep(dirChangeDelay)
	}
	t.Fatalf("the directory listing is not reloaded: %q", text)
}

func TestEditor_WatchStopsOnClose(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "a.txt")
	writeFile(t, path, "a")

	var e Editor
	(&OpenFile{Path: path}).DoOnEditor(&e)
	e.OpenDir(dir, nil)
	if e.bufs[0].unwatch == nil || e.bufs[1].unwatch == nil {
		t.Fatal("the buffers are not watched")
	}
	stopped := 0
	for _, b := range e.bufs {
		unwatch := b.unwatch
		b.unwatch = func() { stopped++; unwatch() }
	}
	for !e.pop() {
	}
	if stopped != 2 {
		t.Errorf("stopped %d watches, want 2", stopped)
	}
	_ = e.watcher.Close()
}
