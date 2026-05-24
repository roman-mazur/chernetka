package content

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSave(t *testing.T) {
	content, err := LoadFullText(strings.NewReader("hello world 1\nhello world 2\nhello world 3"))
	if err != nil {
		t.Fatal(err)
	}

	dst := filepath.Join(t.TempDir(), "out.txt")
	for range 4 {
		if err := Save(&content, dst); err != nil {
			t.Errorf("Save(content, %q) => %s", dst, err)
		}
	}

	f, err := os.Open(dst)
	if err != nil {
		t.Fatal("cannot open dst file after Save:", err)
	}
	t.Cleanup(func() {
		_ = f.Close()
	})
	outContent, err := LoadFullText(f)
	if err != nil {
		t.Fatal(err)
	}
	if outContent.Len() != content.Len() {
		t.Errorf("content length mismatch: got %d, want %d", outContent.Len(), content.Len())
	}
}
