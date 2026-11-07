package content

import (
	"bufio"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestLoadFolder(t *testing.T) {
	var openFile testOpenFile
	fc := LoadFolder("testdata/dirs/sample1", &openFile)
	expectedStructure := []struct {
		name string
		dir  bool
	}{
		{"child-empty", true},
		{"child1", true},
		{"child2", true},
		{"fulltext-example-3.txt", false},
	}

	if len(expectedStructure) != fc.Len() {
		t.Fatalf("length mismatch: expected %d, got %d", len(expectedStructure), fc.Len())
	}

	lines := fc.Lines()
	for i, entry := range expectedStructure {
		if !strings.Contains(lines[i].String(), entry.name) {
			t.Errorf("line %d: expected to contain %q, got %q", i, entry.name, lines[i])
		}
		if mimeType := lines[i].MimeType(); mimeType != "text/filename" {
			t.Errorf("lines[%d].MimeType() = %q, want %q", i, mimeType, "text/filename")
		}

		action, ok := lines[i].(LineAction)
		if !ok {
			t.Errorf("line %d: expected LineAction, got %T", i, lines[i])
			continue
		}

		if entry.dir {
			action.Engage() // Expand.
			if fc.Len() <= len(lines) && !strings.HasSuffix(entry.name, "-empty") {
				t.Errorf("expanding dir %q didn't work", lines[i])
				(*testWriter)(t).printContent(fc)
			}
			if expandedName := lines[i].String(); !strings.Contains(expandedName, "~ "+entry.name) {
				t.Errorf("bad display of expanded dir: %q", expandedName)
			}
			if nextLine := fc.Lines()[i+1]; nextLine.String()[0] != '\t' {
				t.Errorf("unexpected state after expanding line %s, next line: %q", lines[i], nextLine)
				(*testWriter)(t).printContent(fc)
			}

			action.Engage() // Collapse.
			if fc.Len() != len(lines) {
				t.Errorf("collapsing dir %q didn't work", lines[i])
				(*testWriter)(t).printContent(fc)
			}

			if openFile.lastPath != "" {
				t.Errorf("unexpected open file %q after engaging with %q", openFile.lastPath, lines[i])
			}
		} else {
			action.Engage() // Open file.
			if openFile.lastPath != entry.name {
				t.Errorf("opening file %s didn't work; lastPath = %q", entry.name, openFile.lastPath)
			}
		}

		openFile.lastPath = ""
	}
}

func TestFsContent_SyncState(t *testing.T) {
	var openFile testOpenFile
	const samplePath = "testdata/dirs/sample1"
	origin := LoadFolder(samplePath, &openFile)

	t.Log("expand child 2")
	origin.Lines()[2].(LineAction).Engage()
	t.Log("expand child 1")
	origin.Lines()[1].(LineAction).Engage()
	t.Log("expand child 11")
	origin.Lines()[2].(LineAction).Engage()

	expectedStructure := []struct {
		name string
		dir  bool
	}{
		{"child-empty", true},
		{"child1", true},
		{"child11", true},
		{"fulltext-1.txt", false},
		{"fulltext-empty-0.txt", false},
		{"fulltext-example-3.txt", false},
		{"child2", true},
		{"fulltext-1.txt", false},
		{"fulltext-example-3.txt", false},
	}
	lines := origin.Lines()
	t.Log("current state", lines)

	verifyStructure := func(origin *FsContent) {
		t.Helper()
		if len(expectedStructure) != origin.Len() {
			t.Errorf("length mismatch: expected %d, got %d", len(expectedStructure), origin.Len())
		}
		for i, entry := range expectedStructure {
			if !strings.Contains(lines[i].String(), entry.name) {
				t.Errorf("line %d: expected to contain %q, got %q", i, entry.name, lines[i])
			}
		}
	}
	verifyStructure(origin)

	fc := LoadFolder(samplePath, &openFile)
	fc.SyncState(origin)
	t.Log("synced state", fc.Lines())
	verifyStructure(fc)

	t.Log("open", expectedStructure[3].name)
	fc.Lines()[3].(LineAction).Engage()
	if filepath.Base(openFile.lastPath) != expectedStructure[3].name {
		t.Errorf("expected to open %q, got %q", expectedStructure[3].name, openFile.lastPath)
	}
}

func TestFsContent_DirAt(t *testing.T) {
	fc := LoadFolder("testdata/dirs/sample1", nil)
	if fc.RootPath() != "testdata/dirs/sample1" {
		t.Errorf("root path %q", fc.RootPath())
	}
	fc.Lines()[1].(LineAction).Engage() // Expand child1: child11, then its files.

	for line, want := range []string{"child-empty", "child1", "child1/child11", "child1", "child1", "child2", "."} {
		if got := fc.DirAt(line); got != want {
			t.Errorf("line %d (%s): got %q, want %q", line, fc.Lines()[line], got, want)
		}
	}
	for _, line := range []int{-1, fc.Len()} {
		if got := fc.DirAt(line); got != "." {
			t.Errorf("line %d: got %q", line, got)
		}
	}
}

func TestFsContent_DirAt_RootFile(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "a.txt"), nil, 0o600); err != nil {
		t.Fatal(err)
	}
	if got := LoadFolder(dir, nil).DirAt(0); got != "." {
		t.Errorf("got %q", got)
	}
}

// TestFsContent_SyncState_Error covers the directory being removed while its content is displayed.
func TestFsContent_SyncState_Error(t *testing.T) {
	var openFile testOpenFile
	origin := LoadFolder("testdata/dirs/sample1", &openFile)
	origin.Lines()[1].(LineAction).Engage()
	missing := LoadFolder("testdata/dirs/missing", &openFile)
	// The message is platform specific.
	_, readErr := fs.ReadDir(os.DirFS("testdata/dirs/missing"), ".")
	if readErr == nil {
		t.Fatal("missing directory exists")
	}
	notExist := errors.Unwrap(readErr).Error()

	confirmMissinState := func() {
		if missing.Len() != 1 || !strings.Contains(missing.Lines()[0].String(), notExist) {
			t.Errorf("missing directory content: %v", missing.Lines())
		}
	}
	confirmMissinState()

	missing.SyncState(origin)
	confirmMissinState()

	fc := LoadFolder("testdata/dirs/sample1", &openFile)
	fc.SyncState(missing)
	if fc.Len() != 4 {
		t.Errorf("synced content length = %d, want 4 collapsed entries", fc.Len())
	}
}

type testOpenFile struct {
	lastPath string
}

func (tof *testOpenFile) OpenFile(path string) {
	tof.lastPath = path
}

type testWriter testing.T

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.Log(string(p))
	return len(p), nil
}

func (tw *testWriter) printContent(c Document) {
	out := bufio.NewWriter(tw)
	defer out.Flush()
	Print(out, c, 0, c.Len())
}

func TestFsContent_ExpandedDirs(t *testing.T) {
	fc := LoadFolder("testdata/dirs/sample1", nil)
	check := func(t *testing.T, want ...string) {
		t.Helper()
		if got := fc.ExpandedDirs(); !slices.Equal(got, want) {
			t.Errorf("shown dirs %q, want %q", got, want)
		}
	}
	engage := func(name string) {
		for _, l := range fc.Lines() {
			if strings.TrimSpace(l.String()[1:]) == name {
				l.(LineAction).Engage()
				return
			}
		}
		t.Fatalf("no line %q", name)
	}

	check(t, ".")
	engage("child1")
	engage("child2")
	check(t, ".", "child1", "child2")
	engage("child1")
	check(t, ".", "child2")
}
