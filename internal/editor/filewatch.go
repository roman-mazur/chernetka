package editor

import (
	"bytes"
	"os"
	"path/filepath"
	"slices"
	"time"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/watch/dirwatch"
)

const (
	// fileChangeDelay groups the events of a single file update, so a file that is
	// being written is not read in the middle.
	fileChangeDelay = 50 * time.Millisecond
	// dirChangeDelay groups the changes in a directory tree, like the ones of a git checkout.
	dirChangeDelay = 100 * time.Millisecond
)

// fsWatcher returns the watcher shared by the buffers, starting it on the first use.
// It returns nil if the watcher cannot be started.
func (e *Editor) fsWatcher() *dirwatch.Watcher {
	if e.watcher == nil {
		w, err := dirwatch.New(func(err error) { e.Logf("watch: %s", err) })
		if err != nil {
			e.Logf("cannot watch files: %s", err)
			return nil
		}
		e.watcher = w
	}
	return e.watcher
}

// watchFile reloads the buffer content every time its file is changed outside the editor.
// The watch is stopped when the buffer is closed.
func (e *Editor) watchFile(buf *Buffer) {
	w := e.fsWatcher()
	if w == nil {
		return
	}
	path, err := filepath.Abs(buf.Path)
	if err != nil {
		e.Logf("cannot watch %s: %s", buf.Path, err)
		return
	}
	stop, err := w.WatchFile(path, fileChangeDelay, func() {
		data, err := os.ReadFile(path)
		if err != nil {
			e.Logf("cannot reload %s: %s", path, err)
			return
		}
		e.Send(CommandFunc(func(e *Editor) { e.reloadBuffer(buf, data) }))
	})
	if err != nil {
		e.Logf("cannot watch %s: %s", buf.Path, err)
		return
	}
	buf.unwatch = stop
}

// watchDir reloads the directory listing of the buffer every time the directory tree changes.
// The watch is stopped when the buffer is closed.
func (e *Editor) watchDir(buf *Buffer, dir string, open content.OpenFile) {
	w := e.fsWatcher()
	if w == nil {
		return
	}
	stop, err := w.WatchTree(dir, dirChangeDelay, func([]string) {
		folder := content.LoadFolder(dir, open)
		e.Send(CommandFunc(func(e *Editor) {
			if !slices.Contains(e.bufs, buf) {
				return // Closed in the meantime.
			}
			folder.SyncState(buf.Content.(*content.FsContent))
			buf.Content = folder
			buf.clampPos()
			e.renderRequested = true
		}))
	})
	if err != nil {
		e.Logf("cannot watch %s: %s", dir, err)
		return
	}
	buf.unwatch = stop
}

// reloadBuffer replaces the buffer content with the file data read from the disk.
// Unsaved changes in the buffer are discarded.
// TODO: Merge the outside changes with the unsaved ones instead.
func (e *Editor) reloadBuffer(buf *Buffer, data []byte) {
	if !slices.Contains(e.bufs, buf) {
		return // Closed in the meantime.
	}
	ft, err := content.LoadFullText(bytes.NewReader(data))
	if err != nil {
		return
	}
	text := textOf(&ft)
	if text == buf.fileText {
		return // Not changed since the last load or save, for example, saved by this editor.
	}
	buf.fileText = text
	if text == buf.Text() {
		buf.dirty = false
		return
	}

	e.Logf("reload %s", buf.Path)
	buf.Content = &ft
	buf.cancelSelection()
	buf.setMutated()
	buf.dirty = false
	e.renderRequested = true
	e.handleAfterEdit(buf)
}

// textOf returns the document text as it's saved to a file.
func textOf(doc content.Document) string {
	var out bytes.Buffer
	_ = content.SaveToWriter(doc, &out)
	return out.String()
}
