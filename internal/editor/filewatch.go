package editor

import (
	"bytes"
	"os"
	"path/filepath"
	"slices"
	"time"

	"github.com/fsnotify/fsnotify"
	"rmazur.io/chernetka/internal/content"
)

// fileChangeDelay groups the events of a single file update, so a file that is
// being written is not read in the middle.
const fileChangeDelay = 50 * time.Millisecond

// watchFile reloads the buffer content every time its file is changed outside the editor.
// The watcher is stopped when the buffer is closed.
func (e *Editor) watchFile(buf *Buffer) {
	path, err := filepath.Abs(buf.Path)
	if err != nil {
		e.Logf("cannot watch %s: %s", buf.Path, err)
		return
	}
	w, err := fsnotify.NewWatcher()
	if err != nil {
		e.Logf("cannot watch %s: %s", buf.Path, err)
		return
	}
	// The directory is watched since the file is replaced on save by many tools (and this editor),
	// so a watch on the file itself would stop at the first save.
	if err := w.Add(filepath.Dir(path)); err != nil {
		e.Logf("cannot watch %s: %s", buf.Path, err)
		_ = w.Close()
		return
	}
	buf.watcher = w
	go e.handleFileEvents(w, path, buf)
}

func (e *Editor) handleFileEvents(w *fsnotify.Watcher, path string, buf *Buffer) {
	var (
		timer     *time.Timer
		timerChan <-chan time.Time
	)
	for {
		select {
		case ev, ok := <-w.Events:
			if !ok {
				return // Buffer closed.
			}
			if filepath.Clean(ev.Name) != path || !ev.Has(fsnotify.Write) && !ev.Has(fsnotify.Create) {
				continue
			}
			if timer == nil {
				timer = time.NewTimer(fileChangeDelay)
			} else {
				timer.Reset(fileChangeDelay)
			}
			timerChan = timer.C

		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			e.Logf("watch %s: %s", path, err)

		case <-timerChan:
			timerChan = nil
			data, err := os.ReadFile(path)
			if err != nil {
				e.Logf("cannot reload %s: %s", path, err)
				continue
			}
			e.Send(CommandFunc(func(e *Editor) { e.reloadBuffer(buf, data) }))
		}
	}
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
