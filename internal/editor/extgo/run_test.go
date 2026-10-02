package extgo

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/logger"
)

// recordingRunner collects the commands it's asked to run.
type recordingRunner struct{ cmds []Command }

func (rr *recordingRunner) Run(cmd Command) { rr.cmds = append(rr.cmds, cmd) }

func newIntegration(t *testing.T) (*Integration, *recordingRunner) {
	runner := new(recordingRunner)
	ext := &Integration{Runner: runner}
	ext.LogEmbed = logger.Embed(t.Logf)
	return ext, runner
}

func openDoc(t *testing.T, ext *Integration, path, text string) *editor.Buffer {
	t.Helper()
	var edit editor.Editor
	edit.Extend(ext)
	if err := edit.OpenReader(path, strings.NewReader(text)); err != nil {
		t.Fatal(err)
	}
	return edit.Top()
}

// actionCommands returns the command lines of the actions provided for the buffer lines, by line index.
func actionCommands(t *testing.T, buf *editor.Buffer, ext *Integration, runner *recordingRunner) map[int]string {
	t.Helper()
	actions, ok := buf.ExtensionData(ext.ID()).(content.LineActions)
	if !ok {
		t.Fatalf("no line actions for %q", buf.Path)
	}
	res := make(map[int]string)
	for i := range buf.Content.Len() + 1 {
		if action := actions.LineAction(i); action != nil {
			runner.cmds = nil
			action.Engage()
			if len(runner.cmds) != 1 {
				t.Fatalf("line %d action ran %d commands", i, len(runner.cmds))
			}
			res[i] = runner.cmds[0].Line
		}
	}
	return res
}

func TestIntegration_LineActions(t *testing.T) {
	for _, tc := range []struct {
		name, path, text string
		want             map[int]string
	}{
		{
			name: "tests",
			path: "a_test.go",
			text: `package a

import "testing"

func TestA(t *testing.T) {}
func Test(t *testing.T) {}
func Test_b(tt *testing.T) {}
func Testing(t *testing.T) {}
func TestHelper(tb testing.TB) {}
func BenchmarkA(b *testing.B) {}
func helper(t *testing.T) {}
  func TestIndented(t *testing.T) {}`,
			want: map[int]string{
				4: "go test -v -run '^TestA$' .",
				5: "go test -v -run '^Test$' .",
				6: "go test -v -run '^Test_b$' .",
				9: "go test -v -run '^&' -benchmem -bench '^BenchmarkA$' .",
			},
		},
		{
			name: "main",
			path: "cmd/main.go",
			text: "// Command.\npackage main\n\nfunc main() {\n}\n\nfunc other() {}",
			want: map[int]string{3: "go run ."},
		},
		{
			name: "main in a test file",
			path: "main_test.go",
			text: "package main\n\nfunc main() {}",
			want: map[int]string{},
		},
		{
			name: "main of a library",
			path: "lib.go",
			text: "package lib\n\nfunc main() {}",
			want: map[int]string{},
		},
		{
			name: "tests in a regular file",
			path: "lib.go",
			text: "package lib\n\nfunc TestA(t *testing.T) {}",
			want: map[int]string{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ext, runner := newIntegration(t)
			buf := openDoc(t, ext, tc.path, tc.text)
			got := actionCommands(t, buf, ext, runner)
			if len(got) != len(tc.want) {
				t.Errorf("got actions %v, want %v", got, tc.want)
			}
			for ln, want := range tc.want {
				if got[ln] != want {
					t.Errorf("line %d: got %q, want %q", ln, got[ln], want)
				}
			}
		})
	}
}

func TestIntegration_CommandDir(t *testing.T) {
	ext, runner := newIntegration(t)
	buf := openDoc(t, ext, filepath.Join("cmd", "x", "main.go"), "package main\nfunc main() {}")
	buf.ExtensionData(ext.ID()).(content.LineActions).LineAction(1).Engage()
	want, err := filepath.Abs(filepath.Join("cmd", "x"))
	if err != nil {
		t.Fatal(err)
	}
	if len(runner.cmds) != 1 || runner.cmds[0].Dir != want {
		t.Errorf("got commands %v, want one in %s", runner.cmds, want)
	}
}

func TestIntegration_NoData(t *testing.T) {
	ext, _ := newIntegration(t)
	for _, path := range []string{"", "a.txt", "a.go.md"} {
		if data := openDoc(t, ext, path, "package main\nfunc main() {}").ExtensionData(ext.ID()); data != nil {
			t.Errorf("%q has extension data", path)
		}
	}

	var edit editor.Editor
	edit.Extend(ext)
	dir := filepath.Join(t.TempDir(), "pkg.go")
	if err := os.Mkdir(dir, 0o700); err != nil {
		t.Fatal(err)
	}
	edit.OpenDir(dir, nil)
	if data := edit.Top().ExtensionData(ext.ID()); data != nil {
		t.Error("directory listing has extension data")
	}

	edit = editor.Editor{}
	edit.Extend(&Integration{})
	if err := edit.OpenReader("main.go", strings.NewReader("package main\nfunc main() {}")); err != nil {
		t.Fatal(err)
	}
	if data := edit.Top().ExtensionData(ext.ID()); data != nil {
		t.Error("extension without a runner has data")
	}
}

func TestIntegration_SavesBeforeRun(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "a_test.go")
	if err := os.WriteFile(path, []byte("package a\nfunc TestA(t *testing.T) {}"), 0o600); err != nil {
		t.Fatal(err)
	}
	var saved []string
	ext := &Integration{Runner: runnerFunc(func(Command) {
		data, _ := os.ReadFile(path)
		saved = append(saved, string(data))
	})}
	h := editor.NewTestHarness()
	h.Extend(ext)
	(&editor.OpenFile{Path: path}).DoOnEditor(h.Editor)
	h.Run(t)

	h.SendInputSequence(t, "j")
	h.SendInput(t, []byte{'\r'})
	h.SendInputSequence(t, "A")
	h.SendInputSequence(t, " // edited")
	h.SendInput(t, []byte{0x12}) // Ctrl+R
	want := []string{"package a\nfunc TestA(t *testing.T) {}", "package a\nfunc TestA(t *testing.T) {} // edited"}
	if !slices.Equal(saved, want) {
		t.Errorf("run with the files %q, want %q", saved, want)
	}
}

type runnerFunc func(Command)

func (f runnerFunc) Run(cmd Command) { f(cmd) }

// TestIntegration_RerunFromAnotherFile checks that Ctrl+R runs the test again after
// another file is edited, and after that file is closed.
func TestIntegration_RerunFromAnotherFile(t *testing.T) {
	dir := t.TempDir()
	testPath, libPath := filepath.Join(dir, "a_test.go"), filepath.Join(dir, "a.go")
	if err := os.WriteFile(testPath, []byte("package a\nfunc TestA(t *testing.T) {}"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(libPath, []byte("package a"), 0o600); err != nil {
		t.Fatal(err)
	}
	ext, runner := newIntegration(t)
	h := editor.NewTestHarness()
	h.Extend(ext)
	(&editor.OpenFile{Path: testPath}).DoOnEditor(h.Editor)
	h.Run(t)

	checkRuns := func(want int) {
		t.Helper()
		if len(runner.cmds) != want {
			t.Fatalf("ran %d commands, want %d: %v", len(runner.cmds), want, runner.cmds)
		}
		if got := runner.cmds[want-1].Line; got != "go test -v -run '^TestA$' ." {
			t.Errorf("ran %q, want the test", got)
		}
	}

	h.SendInputSequence(t, "j")
	h.SendInput(t, []byte{'\r'})
	checkRuns(1)

	h.Post(t, &editor.OpenFile{Path: libPath})
	h.SendInputSequence(t, "A // edited")
	h.SendInput(t, []byte{0x1b}) // Esc
	h.SendInput(t, []byte{0x12}) // Ctrl+R saves the changes and runs the test.
	checkRuns(2)

	h.SendInput(t, []byte{'q'}) // Close the other file, the test file is on top again.
	h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
		if e.Top().Path != testPath {
			t.Errorf("top buffer is %q, want %q", e.Top().Path, testPath)
		}
	}))
	h.SendInput(t, []byte{0x12}) // Ctrl+R
	checkRuns(3)
}
