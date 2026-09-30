package main

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/cheimg"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
)

// recordingViewer collects the items it's asked to show.
type recordingViewer []cheimg.Item

func (rv *recordingViewer) Show(it cheimg.Item) { *rv = append(*rv, it) }

// newTestDelegate returns a delegate that records the items shown and the replacements with che-img.
func newTestDelegate(t *testing.T) (*editDelegate, *recordingViewer, *[]string) {
	var (
		viewer   recordingViewer
		replaced []string
	)
	return &editDelegate{
		edit:   new(editor.Editor),
		logf:   t.Logf,
		viewer: &viewer,
		replaceWithViewer: func(path string) error {
			replaced = append(replaced, path)
			return errors.New("not found")
		},
	}, &viewer, &replaced
}

func writeTestFiles(t *testing.T, names ...string) string {
	t.Helper()
	dir := t.TempDir()
	for _, name := range names {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(name), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

// topPath returns the path of the active buffer, or "" if there is none.
func topPath(e *editor.Editor) string {
	if top := e.Top(); top != nil {
		return top.Path
	}
	return ""
}

func TestEditDelegate_OpenFile(t *testing.T) {
	dir := writeTestFiles(t, "notes.txt", "arch.d2", "photo.png")
	textPath, diagramPath, imagePath := filepath.Join(dir, "notes.txt"), filepath.Join(dir, "arch.d2"), filepath.Join(dir, "photo.png")

	t.Run("text", func(t *testing.T) {
		ed, viewer, replaced := newTestDelegate(t)
		ed.openFile(textPath)
		if len(*viewer) != 0 || len(*replaced) != 0 {
			t.Errorf("text file sent to che-img: shown %v, replaced %v", *viewer, *replaced)
		}
		if got := topPath(ed.edit); got != textPath {
			t.Errorf("active buffer = %q, want %q", got, textPath)
		}
	})

	t.Run("diagram is edited and shown", func(t *testing.T) {
		ed, viewer, replaced := newTestDelegate(t)
		ed.openFile(diagramPath)
		if want := []cheimg.Item{{Path: diagramPath}}; len(*viewer) != 1 || (*viewer)[0] != want[0] {
			t.Errorf("shown %v, want %v", *viewer, want)
		}
		if len(*replaced) != 0 {
			t.Errorf("editor replaced for a diagram")
		}
		if got := topPath(ed.edit); got != diagramPath {
			t.Errorf("active buffer = %q, want %q", got, diagramPath)
		}
	})

	t.Run("image with open buffers is shown", func(t *testing.T) {
		ed, viewer, replaced := newTestDelegate(t)
		ed.openFile(textPath)
		ed.openFile(imagePath)
		if want := (cheimg.Item{Path: imagePath}); len(*viewer) != 1 || (*viewer)[0] != want {
			t.Errorf("shown %v, want %v", *viewer, want)
		}
		if len(*replaced) != 0 {
			t.Errorf("editor replaced while buffers are open")
		}
		if got := topPath(ed.edit); got != textPath {
			t.Errorf("image opened in the editor: active buffer = %q", got)
		}
	})

	t.Run("image without buffers replaces the editor", func(t *testing.T) {
		ed, viewer, replaced := newTestDelegate(t)
		ed.openFile(imagePath)
		if len(*replaced) != 1 || (*replaced)[0] != imagePath {
			t.Errorf("replaced with %v, want che-img for %s", *replaced, imagePath)
		}
		if len(*viewer) != 0 {
			t.Errorf("shown %v", *viewer)
		}
		// The replacement is faked to fail: the reason is shown in the editor.
		top := ed.edit.Top()
		if top == nil {
			t.Fatal("no buffer with the error")
		}
		errContent, ok := top.Content.(*content.ErrorContent)
		if !ok || !strings.Contains(errContent.Error.Error(), "cannot start che-img: not found") {
			t.Errorf("top buffer content = %#v", top.Content)
		}
	})

	t.Run("relative paths are shown as absolute", func(t *testing.T) {
		t.Chdir(dir)
		ed, viewer, _ := newTestDelegate(t)
		ed.openFile(textPath) // Keep che-img from replacing the editor.
		ed.openFile("photo.png")
		if want := (cheimg.Item{Path: imagePath}); len(*viewer) != 1 || (*viewer)[0] != want {
			t.Errorf("shown %v, want %v", *viewer, want)
		}
	})
}

func TestShellQuote(t *testing.T) {
	for _, s := range []string{"", "plain", "my dir", "it's", "''", "$HOME", "`date`", `a\b`, "x'; echo injected; '", "a\nb"} {
		out, err := exec.Command("sh", "-c", "printf %s "+shellQuote(s)).CombinedOutput()
		if err != nil {
			t.Errorf("%q: %s: %s", s, err, out)
			continue
		}
		if string(out) != s {
			t.Errorf("%s printed %q, want %q", shellQuote(s), out, s)
		}
	}
}

func TestGitDiffCommand(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git is not installed")
	}
	repo := filepath.Join(t.TempDir(), "it's my $HOME `repo`")
	path := filepath.Join(repo, "notes.txt")
	git := func(args ...string) {
		t.Helper()
		cmd := exec.Command("git", append([]string{"-C", repo}, args...)...)
		cmd.Env = append(os.Environ(), "GIT_AUTHOR_NAME=t", "GIT_AUTHOR_EMAIL=t@t", "GIT_COMMITTER_NAME=t", "GIT_COMMITTER_EMAIL=t@t")
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %s: %s", args, err, out)
		}
	}
	write := func(text string) {
		t.Helper()
		if err := os.WriteFile(path, []byte(text), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.Mkdir(repo, 0o700); err != nil {
		t.Fatal(err)
	}
	git("init", "-q")
	write("committed\n")
	git("add", ".")
	git("commit", "-qm", "init")
	write("staged\n")
	git("add", ".")
	write("staged\nunstaged\n")

	t.Chdir(t.TempDir()) // The editor may run outside the repository.
	out, err := exec.Command("sh", "-c", gitDiffCommand(path)).CombinedOutput()
	if err != nil {
		t.Fatalf("%s: %s", err, out)
	}
	for _, want := range []string{"-committed", "+staged", "+unstaged"} {
		if !strings.Contains(string(out), want) {
			t.Errorf("diff misses %q:\n%s", want, out)
		}
	}
}
