package content

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Save writes the content to the destination path. The content is written to a temporary file
// in the destination directory first, which then replaces the destination, so the destination
// is never left half written. An existing destination keeps its permissions, and a new one
// gets newFileMode.
func Save(content Document, dst string) (err error) {
	mode := os.FileMode(newFileMode)
	if info, err := os.Stat(dst); err == nil {
		mode = info.Mode().Perm()
	}

	f, err := os.CreateTemp(filepath.Dir(dst), "."+filepath.Base(dst)+".*.tmp")
	if err != nil {
		return fmt.Errorf("could not create temporary file: %w", err)
	}
	defer func() {
		if err != nil {
			_ = f.Close()
			_ = os.Remove(f.Name())
		}
	}()

	w := bufio.NewWriter(f)
	if err := SaveToWriter(content, w); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}
	if err := f.Chmod(mode); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(f.Name(), dst)
}

// newFileMode is the permissions of the files created by Save.
const newFileMode = 0o644

func SaveToWriter(content Document, out io.Writer) error {
	if content == nil {
		return nil
	}
	for i, line := range content.Lines() {
		if i > 0 {
			if _, err := out.Write([]byte("\n")); err != nil {
				return err
			}
		}
		if _, err := out.Write([]byte(line.String())); err != nil {
			return err
		}
	}
	return nil
}
