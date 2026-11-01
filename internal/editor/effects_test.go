package editor

import (
	"errors"
	"path/filepath"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/watch/dirwatch"
)

// IsolateEffects lets the files be saved only in dir, until the test ends. With the clipboard
// kept in memory by the harness, it allows typing any input, like in fuzz tests, without
// changing the files outside dir.
func IsolateEffects(t testing.TB, dir string) {
	prevSave := saveContent
	t.Cleanup(func() { saveContent = prevSave })

	saveContent = func(doc content.Document, dst string) (string, error) {
		abs, err := filepath.Abs(dst)
		if err != nil {
			return "", err
		}
		if rel, err := filepath.Rel(dir, abs); err != nil || !filepath.IsLocal(rel) {
			return "", errors.New("saving outside the test directory")
		}
		return content.Save(doc, dst)
	}
}

// DisableFileWatching keeps the editor from watching the files and directories it opens, until
// the test ends: the changes made outside the editor are not reloaded. The watcher waits for
// the file system events in a goroutine that synctest never sees as blocked, so a test bubble
// opening a file with the watching would never settle.
func DisableFileWatching(t testing.TB) {
	prev := newWatcher
	t.Cleanup(func() { newWatcher = prev })
	newWatcher = func(func(error)) (*dirwatch.Watcher, error) {
		return nil, errors.New("file watching is disabled in the test")
	}
}
