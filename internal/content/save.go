package content

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Save writes the content to the destination path.
func Save(content Interface, dst string) error {
	tmpPath := filepath.Join(tmpDir(), fmt.Sprintf("%s-%d", filepath.Base(dst), time.Now().Unix()))
	f, err := os.Create(tmpPath)
	if err != nil {
		return fmt.Errorf("could not create temporary file: %w", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for i, line := range content.Lines() {
		if i > 0 {
			if _, err := w.WriteString("\n"); err != nil {
				return err
			}
		}
		if _, err := w.WriteString(line.String()); err != nil {
			return err
		}
	}
	if err := w.Flush(); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return os.Rename(tmpPath, dst)
}

func tmpDir() string {
	dir := filepath.Join(os.TempDir(), "edit")
	_ = os.MkdirAll(dir, 0700)
	return dir
}
