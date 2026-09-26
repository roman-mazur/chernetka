package editor

import (
	"os"
	"path/filepath"
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
