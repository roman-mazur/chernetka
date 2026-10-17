package editor

import (
	"errors"
	"path/filepath"
	"testing"

	"rmazur.io/chernetka/internal/content"
)

// IsolateEffects lets the files be saved only in dir, until the test ends. With the clipboard
// kept in memory by the harness, it allows typing any input, like in fuzz tests, without
// changing the files outside dir.
func IsolateEffects(t testing.TB, dir string) {
	prevSave := saveContent
	t.Cleanup(func() { saveContent = prevSave })

	saveContent = func(doc content.Document, dst string) error {
		abs, err := filepath.Abs(dst)
		if err != nil {
			return err
		}
		if rel, err := filepath.Rel(dir, abs); err != nil || !filepath.IsLocal(rel) {
			return errors.New("saving outside the test directory")
		}
		return content.Save(doc, dst)
	}
}
