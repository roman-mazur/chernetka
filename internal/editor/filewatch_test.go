package editor

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/content"
)

func TestEditor_WatchFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "a.txt")
	writeFile(t, path, "first")

	h := NewTestHarness(t)
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

	h := NewTestHarness(t)
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

func TestEditor_ReloadBuffer(t *testing.T) {
	open := func(t *testing.T, text string) (*Editor, *Buffer, string) {
		t.Helper()
		path := filepath.Join(t.TempDir(), "a.txt")
		writeFile(t, path, text)
		e := &Editor{cmdChannel: make(chan Command, 16)}
		if err := e.OpenReader(path, strings.NewReader(text)); err != nil {
			t.Fatal(err)
		}
		return e, e.Top(), path
	}
	check := func(t *testing.T, buf *Buffer, text string, dirty bool) {
		t.Helper()
		if got := buf.Text(); got != text {
			t.Errorf("text %q, want %q", got, text)
		}
		if buf.dirty != dirty {
			t.Errorf("dirty %t, want %t", buf.dirty, dirty)
		}
	}

	t.Run("file not changed", func(t *testing.T) {
		e, buf, _ := open(t, "first")
		buf.Mutate().Update(0, content.TextLine("typed"))
		e.reloadBuffer(buf, []byte("first"))
		check(t, buf, "typed", true)
	})

	t.Run("file changed", func(t *testing.T) {
		e, buf, _ := open(t, "first")
		buf.Mutate().Update(0, content.TextLine("typed"))
		e.reloadBuffer(buf, []byte("reloaded\nfile"))
		check(t, buf, "reloaded\nfile", false)
	})

	t.Run("edit after reload", func(t *testing.T) {
		e, buf, _ := open(t, "first")
		buf.Mutate().Update(0, content.TextLine("typed"))
		e.reloadBuffer(buf, []byte("reloaded\nfile"))
		buf.Mutate().Update(0, content.TextLine("edited"))
		check(t, buf, "edited\nfile", true)
	})

	t.Run("undo after reload", func(t *testing.T) {
		e, buf, _ := open(t, "first")
		buf.Mutate().Update(0, content.TextLine("typed"))
		e.reloadBuffer(buf, []byte("reloaded"))
		sendChunks(t, e, "\x1a") // The changes before the reload are not undone.
		check(t, buf, "reloaded", false)

		buf.Mutate().Update(0, content.TextLine("edited"))
		sendChunks(t, e, "\x1a")
		check(t, buf, "reloaded", true)
	})

	t.Run("own save", func(t *testing.T) {
		// The saved file differs from the opened one: the lines end with \n now.
		e, buf, path := open(t, "one\r\ntwo")
		buf.Mutate().Update(0, content.TextLine("saved"))
		(&Save{DstPath: path}).DoOnBuffer(buf, e.rPrefs)
		saved, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		// Typed before the watcher reports the save.
		buf.Mutate().Update(0, content.TextLine("typed"))
		e.reloadBuffer(buf, saved)
		check(t, buf, "typed\ntwo", true)
	})

	t.Run("save elsewhere", func(t *testing.T) {
		e, buf, path := open(t, "first")
		buf.Mutate().Update(0, content.TextLine("copied"))
		(&Save{DstPath: path + ".copy"}).DoOnBuffer(buf, e.rPrefs)
		// The file itself is changed outside the editor to the text of the copy.
		e.reloadBuffer(buf, []byte("copied"))
		check(t, buf, "copied", false)
	})
}

// textOf returns the document text as it's saved to a file.
func textOf(doc content.Document) string {
	var out strings.Builder
	_ = content.SaveToWriter(doc, &out)
	return out.String()
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

	h := NewTestHarness(t)
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

// TestEditor_WatchLargeTree opens the directory browser on a tree with more entries than
// the process may open files, like a checkout of Chromium (500k files) or AOSP is for
// the limit of macOS (61440 per process). The open files are limited for the test to
// make a small tree large enough.
func TestEditor_WatchLargeTree(t *testing.T) {
	const openFilesLimit = 1024
	dir := t.TempDir()
	dirs, files := generateTree(t, dir, 3, 8, 4)
	t.Logf("generated %d directories and %d files in %s", dirs, files, dir)
	limitOpenFiles(t, openFilesLimit)

	h := NewTestHarness(t)
	h.Run(t)
	// Post fails if the directory is not opened in a second: the UI loop is blocked meanwhile.
	h.Post(t, CommandFunc(func(e *Editor) { e.OpenDir(dir, nil) }))

	// Expand dir0, dir0/dir0, and dir0/dir0/dir0 (the first line of each listing) like Enter does.
	listing := func(e *Editor) string { return textOf(e.Top().Content) }
	for depth := range 3 {
		h.Post(t, CommandFunc(func(e *Editor) {
			e.Top().updateCursor(content.Position{Line: depth})
			e.maybeEngage(e.Top())
		}))
	}
	h.Post(t, CommandFunc(func(e *Editor) { t.Logf("directory buffer:\n%s", listing(e)) }))

	t.Run("change in an expanded directory", func(t *testing.T) {
		writeFile(t, filepath.Join(dir, "dir0", "dir0", "dir0", "new.txt"), "new")
		var text string
		for range 50 {
			h.Post(t, CommandFunc(func(e *Editor) { text = listing(e) }))
			if strings.Contains(text, "new.txt") {
				return
			}
			time.Sleep(dirChangeDelay)
		}
		t.Errorf("the listing is not reloaded:\n%s", text)
	})

	watched := func(e *Editor) []string {
		return slices.Sorted(maps.Keys(e.bufs[0].dirWatch.stops))
	}
	t.Run("watched directories", func(t *testing.T) {
		h.Post(t, CommandFunc(func(e *Editor) {
			if got, want := watched(e), []string{".", "dir0", "dir0/dir0", "dir0/dir0/dir0"}; !slices.Equal(got, want) {
				t.Errorf("watching %q, want %q", got, want)
			}
		}))
	})

	t.Run("collapse", func(t *testing.T) {
		h.Post(t, CommandFunc(func(e *Editor) {
			e.Top().updateCursor(content.Position{Line: 1}) // dir0/dir0
			e.maybeEngage(e.Top())
			if got, want := watched(e), []string{".", "dir0"}; !slices.Equal(got, want) {
				t.Errorf("watching %q, want %q", got, want)
			}
		}))
	})

	t.Run("open a file", func(t *testing.T) {
		path := filepath.Join(dir, "dir7", "dir7", "dir7", "file3.txt")
		h.Post(t, &OpenFile{Path: path})
		h.Post(t, CommandFunc(func(e *Editor) {
			if got := e.Top().Text(); got != path {
				t.Errorf("opened %q, got:\n%s", path, got)
			}
		}))
	})
}

// generateTree creates directories nested depth levels deep, with fanout subdirectories
// and the given number of files in each directory. A file contains its path.
func generateTree(t *testing.T, root string, depth, fanout, filesPerDir int) (dirs, files int) {
	t.Helper()
	for i := range filesPerDir {
		path := filepath.Join(root, fmt.Sprintf("file%d.txt", i))
		writeFile(t, path, path)
		files++
	}
	if depth == 0 {
		return dirs, files
	}
	for i := range fanout {
		sub := filepath.Join(root, fmt.Sprintf("dir%d", i))
		if err := os.Mkdir(sub, 0700); err != nil {
			t.Fatal(err)
		}
		d, f := generateTree(t, sub, depth-1, fanout, filesPerDir)
		dirs, files = dirs+d+1, files+f
	}
	return dirs, files
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

func TestWatchInBackground(t *testing.T) {
	type result struct {
		stop func()
		err  error
	}
	start := func() (chan result, func() (func(), error)) {
		results := make(chan result)
		return results, func() (func(), error) { r := <-results; return r.stop, r.err }
	}

	t.Run("stopped after it started", func(t *testing.T) {
		results, startFn := start()
		stopped := make(chan struct{})
		stop := watchInBackground(startFn, func(err error) { t.Error(err) })
		results <- result{stop: func() { close(stopped) }}
		for i := 0; ; i++ { // Started on the other goroutine.
			stop()
			select {
			case <-stopped:
				return
			case <-time.After(10 * time.Millisecond):
				if i == 100 {
					t.Fatal("the watch is not stopped")
				}
			}
		}
	})

	t.Run("stopped while starting", func(t *testing.T) {
		results, startFn := start()
		stopped := make(chan struct{})
		stop := watchInBackground(startFn, func(err error) { t.Error(err) })
		stop()
		results <- result{stop: func() { close(stopped) }}
		select {
		case <-stopped:
		case <-time.After(time.Second):
			t.Fatal("the watch is not stopped once started")
		}
	})

	t.Run("failed", func(t *testing.T) {
		results, startFn := start()
		failed := make(chan error, 1)
		stop := watchInBackground(startFn, func(err error) { failed <- err })
		results <- result{err: os.ErrNotExist}
		if err := <-failed; err != os.ErrNotExist {
			t.Errorf("failed with %v", err)
		}
		stop() // No-op.
	})
}
