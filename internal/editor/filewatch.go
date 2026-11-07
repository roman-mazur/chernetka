package editor

import (
	"bytes"
	"os"
	"path/filepath"
	"slices"
	"sync"
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

// newWatcher starts a file system watcher. Tests may replace it to not watch the files.
var newWatcher = dirwatch.New

// fsWatcher returns the watcher shared by the buffers, starting it on the first use.
// It returns nil if the watcher cannot be started.
func (e *Editor) fsWatcher() *dirwatch.Watcher {
	if e.watcher == nil {
		w, err := newWatcher(func(err error) { e.Logf("watch: %s", err) })
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
	reload := func(data []byte) {
		e.Send(CommandFunc(func(e *Editor) { e.reloadBuffer(buf, data) }))
	}
	loaded := buf.srcFingerprint
	buf.unwatch = watchInBackground(func() (stop func(), err error) {
		stop, err = w.WatchFile(path, fileChangeDelay, func() {
			data, err := os.ReadFile(path)
			if err != nil {
				e.Logf("cannot reload %s: %s", path, err)
				return
			}
			reload(data)
		})
		if err != nil {
			return nil, err
		}
		// The file might be changed before the watch started.
		if data, err := os.ReadFile(path); err == nil {
			if _, fp, err := content.LoadFullText(bytes.NewReader(data)); err == nil && fp != loaded {
				reload(data)
			}
		}
		return stop, nil
	}, func(err error) { e.Logf("cannot watch %s: %s", buf.Path, err) })
}

// watchInBackground starts a watch on another goroutine: starting it takes a while for a large
// directory, as it's listed, or its entries are opened. The returned function stops the watch,
// also if it's still starting. The error of starting it is passed to failed.
func watchInBackground(start func() (stop func(), err error), failed func(error)) (stop func()) {
	var (
		mu      sync.Mutex
		stopped bool
		stopFn  func()
	)
	go func() {
		s, err := start()
		if err != nil {
			failed(err)
			return
		}
		mu.Lock()
		defer mu.Unlock()
		if stopped {
			s()
			return
		}
		stopFn = s
	}()
	return func() {
		mu.Lock()
		defer mu.Unlock()
		stopped = true
		if stopFn != nil {
			stopFn()
			stopFn = nil
		}
	}
}

// watchDir reloads the directory listing of the buffer every time the listing of a directory
// it shows changes. The watch is stopped when the buffer is closed.
func (e *Editor) watchDir(buf *Buffer, dir string, open content.OpenFile) {
	w := e.fsWatcher()
	if w == nil {
		return
	}
	dw := &dirWatch{e: e, w: w, buf: buf, root: dir, open: open, stops: make(map[string]func())}
	buf.dirWatch, buf.unwatch = dw, dw.stop
	dw.sync()
}

// dirWatch watches the directories shown in a directory buffer: the root and the expanded ones.
// The whole tree is not watched: watching a directory may take an open file for each of its
// entries (kqueue on macOS does so), and a large tree, like a checkout of Chromium or AOSP,
// has more entries than a process may open.
type dirWatch struct {
	e    *Editor
	w    *dirwatch.Watcher
	buf  *Buffer
	root string
	open content.OpenFile

	stops map[string]func() // stops watching a shown directory, by its path relative to the root
}

// sync watches the directories the buffer shows now, and stops watching the others.
// It's called after a directory is expanded or collapsed, and after the listing is reloaded.
func (dw *dirWatch) sync() {
	fc, ok := dw.buf.Content.(*content.FsContent)
	if !ok || dw.stops == nil {
		return
	}
	shown := make(map[string]bool)
	for _, d := range fc.ExpandedDirs() {
		shown[d] = true
	}
	for d, stop := range dw.stops {
		if !shown[d] {
			stop()
			delete(dw.stops, d)
		}
	}
	for d := range shown {
		if dw.stops[d] != nil {
			continue
		}
		path, listed := filepath.Join(dw.root, filepath.FromSlash(d)), fc.DirEntries(d)
		dw.stops[d] = watchInBackground(func() (stop func(), err error) {
			if stop, err = dw.w.WatchDir(path, dirChangeDelay, dw.reload); err != nil {
				return nil, err
			}
			// The directory might be changed before the watch started.
			if entries, err := os.ReadDir(path); err == nil && !slices.Equal(entryNames(entries), listed) {
				dw.reload(nil)
			}
			return stop, nil
		}, func(err error) { dw.e.Logf("cannot watch %s: %s", d, err) })
	}
}

func entryNames(entries []os.DirEntry) []string {
	names := make([]string, len(entries))
	for i, e := range entries {
		names[i] = e.Name()
	}
	slices.Sort(names)
	return names
}

// reload reads the listing again, keeping the same directories expanded.
// It's called on the watcher goroutine.
func (dw *dirWatch) reload([]string) {
	folder := content.LoadFolder(dw.root, dw.open)
	dw.e.Send(CommandFunc(func(e *Editor) {
		if !slices.Contains(e.bufs, dw.buf) {
			return // Closed in the meantime.
		}
		folder.SyncState(dw.buf.Content.(*content.FsContent))
		dw.buf.replaceContent(folder)
		dw.sync()
		e.renderRequested = true
	}))
}

// stop stops watching all the directories.
func (dw *dirWatch) stop() {
	for _, stop := range dw.stops {
		stop()
	}
	dw.stops = nil
}

// reloadBuffer replaces the buffer content with the file data read from the disk.
// Unsaved changes in the buffer are discarded.
// TODO: Consider merging the outside changes with the unsaved ones. Or presenting a choice what version to keep.
func (e *Editor) reloadBuffer(buf *Buffer, data []byte) {
	if !slices.Contains(e.bufs, buf) {
		return // Closed in the meantime.
	}
	ft, fp, err := content.LoadFullText(bytes.NewReader(data))
	if err != nil {
		return
	}
	if fp == buf.srcFingerprint {
		return // Not changed since the last load or save, for example, saved by this editor.
	}

	e.Logf("reload %s", buf.Path)
	buf.srcFingerprint = fp
	buf.replaceContent(&ft)
	e.renderRequested = true
	e.handleAfterEdit(buf)
}
