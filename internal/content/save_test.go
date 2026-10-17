package content

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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

	t.Run("real-sources", func(t *testing.T) {
		const basePath = "../editor"
		entries, err := os.ReadDir(basePath)
		must(t, err)
		for _, entry := range entries {
			if !entry.IsDir() && filepath.Ext(entry.Name()) == ".go" {
				t.Run(entry.Name(), func(t *testing.T) {
					f, err := os.Open(filepath.Join(basePath, entry.Name()))
					must(t, err)
					t.Cleanup(func() { _ = f.Close() })

					stat, err := f.Stat()
					must(t, err)
					content, err := LoadFullText(f)
					must(t, err)

					var out bytes.Buffer
					must(t, SaveToWriter(&content, &out))

					if stat.Size() != int64(out.Len()) {
						t.Errorf("content size mismatch: got %d, want %d", stat.Size(), out.Len())
					}
				})
			}
		}
	})

	t.Run("empty", func(t *testing.T) {
		var out bytes.Buffer
		must(t, SaveToWriter(Empty(), &out))
		if out.Len() != 0 {
			t.Errorf("content size mismatch: got %d, want 0", out.Len())
		}
	})
}

// TestSave_SameNameConcurrently checks that the files of the same name saved at the same time,
// like by the editors of different projects, keep their own content.
func TestSave_SameNameConcurrently(t *testing.T) {
	const saves = 8
	dsts := make([]string, saves)
	errs := make(chan error, saves)
	for i := range saves {
		dsts[i] = filepath.Join(t.TempDir(), "out.txt")
		go func() {
			content, err := LoadFullText(strings.NewReader(fmt.Sprintf("content %d", i)))
			if err == nil {
				err = Save(&content, dsts[i])
			}
			errs <- err
		}()
	}
	for range saves {
		if err := <-errs; err != nil {
			t.Error(err)
		}
	}
	for i, dst := range dsts {
		data, err := os.ReadFile(dst)
		if want := fmt.Sprintf("content %d", i); err != nil || string(data) != want {
			t.Errorf("%s has %q (%v), want %q", dst, data, err, want)
		}
	}
}

func TestSave_Mode(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("the permissions are not kept on Windows")
	}
	content, err := LoadFullText(strings.NewReader("#!/bin/sh"))
	must(t, err)
	dir := t.TempDir()

	for _, tc := range []struct {
		name     string
		existing os.FileMode // 0 for a new file
		want     os.FileMode
	}{
		{name: "new file", want: newFileMode},
		{name: "executable", existing: 0o755, want: 0o755},
		{name: "private", existing: 0o600, want: 0o600},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dst := filepath.Join(dir, strings.ReplaceAll(tc.name, " ", "-"))
			if tc.existing != 0 {
				must(t, os.WriteFile(dst, nil, tc.existing))
				must(t, os.Chmod(dst, tc.existing)) // Not limited by the umask.
			}
			must(t, Save(&content, dst))
			info, err := os.Stat(dst)
			must(t, err)
			if got := info.Mode().Perm(); got != tc.want {
				t.Errorf("mode %v, want %v", got, tc.want)
			}
		})
	}
	assertNoTempFiles(t, dir)
}

func TestSave_Error(t *testing.T) {
	content, err := LoadFullText(strings.NewReader("text"))
	must(t, err)
	dir := t.TempDir()
	// Renaming a file over a directory fails.
	dst := filepath.Join(dir, "dir")
	must(t, os.Mkdir(dst, 0o700))
	must(t, os.WriteFile(filepath.Join(dst, "file"), nil, 0o600))

	if err := Save(&content, dst); err == nil {
		t.Error("saved over a directory")
	}
	assertNoTempFiles(t, dir)
}

// assertNoTempFiles checks that only the saved files are in dir.
func assertNoTempFiles(t *testing.T, dir string) {
	t.Helper()
	entries, err := os.ReadDir(dir)
	must(t, err)
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".tmp") {
			t.Errorf("temporary file %s is left", e.Name())
		}
	}
}
