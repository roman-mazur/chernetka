package editor

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/vt/escape"
)

func TestFuzzyScore(t *testing.T) {
	for _, tc := range []struct {
		query, candidate string
		ok               bool
	}{
		{"", "any.go", true},
		{"bf", "buffer.go", true},
		{"BUF", "internal/editor/buffer.go", true},
		{"edbuf", "internal/editor/buffer.go", true},
		{"fub", "buffer.go", false},
		{"buffers", "buffer.go", false},
	} {
		if _, ok := fuzzyScore(tc.query, tc.candidate); ok != tc.ok {
			t.Errorf("fuzzyScore(%q, %q) ok = %t, want %t", tc.query, tc.candidate, ok, tc.ok)
		}
	}

	// Pairs of candidates for a query: the first one must be ranked higher.
	for _, tc := range []struct {
		query, better, worse string
	}{
		{"buf", "internal/editor/buffer.go", "internal/buf_test/x.go"}, // file name
		{"edit", "internal/editor/x.go", "internal/extd2/idx_test.go"}, // consecutive
		{"ed", "internal/editor/x.go", "internal/code/x.go"},           // segment start
	} {
		better, _ := fuzzyScore(tc.query, filepath.FromSlash(tc.better))
		worse, _ := fuzzyScore(tc.query, filepath.FromSlash(tc.worse))
		if better <= worse {
			t.Errorf("query %q: %q scored %d, %q scored %d", tc.query, tc.better, better, tc.worse, worse)
		}
	}
}

func TestQuickOpen_Refresh(t *testing.T) {
	root := t.TempDir()
	for _, p := range []string{"cmd.go", "buffer_test.go", "buffer.go", filepath.Join("sub", "b.txt")} {
		p = filepath.Join(root, p)
		if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, nil, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	e := newPickerTestEditor(root)
	for _, name := range []string{"cmd.go", "notes.txt", "buffer.go"} {
		e.OpenBuffer(&Buffer{Path: filepath.Join(root, name), Content: content.Empty()})
	}
	e.OpenBuffer(&Buffer{Content: content.Empty()}) // Scratch buffers are not listed.
	e.OpenReader(filepath.Join(root, "buffer.go"), nil)

	displayed := func() (res []string) {
		for _, m := range e.status.quick.matches {
			res = append(res, m.display)
		}
		return
	}

	e.startQuickOpen(e.Top())
	runQueuedCommand(t, e)
	// Open buffers from the most recent one, the current buffer is the last one;
	// then the other files.
	want := []string{"notes.txt", "cmd.go", "buffer.go", filepath.Join("sub", "b.txt"), "buffer_test.go"}
	if got := displayed(); !slices.Equal(got, want) {
		t.Errorf("matches = %q, want %q", got, want)
	}
	if e.status.quick.total != 5 {
		t.Errorf("total = %d", e.status.quick.total)
	}

	e.Top().cmdline += "bt"
	e.syncQuickOpen(e.Top())
	// Matching the file name is preferred.
	if got, want := displayed(), []string{filepath.Join("sub", "b.txt"), "buffer_test.go"}; !slices.Equal(got, want) {
		t.Errorf("matches = %q, want %q", got, want)
	}
	if got, want := e.status.quick.matches[0].path, filepath.Join(root, "sub", "b.txt"); got != want {
		t.Errorf("path to open = %q, want %q", got, want)
	}
	if e.status.quick.matches[0].open {
		t.Errorf("file is marked as open")
	}
}

func TestListProjectFiles(t *testing.T) {
	root := t.TempDir()
	for _, p := range []string{"a.go", "sub/b.go", ".git-like/c", "node_modules/d.js", "ignored.log"} {
		p = filepath.Join(root, filepath.FromSlash(p))
		if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, nil, 0o600); err != nil {
			t.Fatal(err)
		}
	}

	t.Run("walk", func(t *testing.T) {
		got := listProjectFiles(root, 100)
		want := []string{"a.go", "ignored.log", filepath.Join("sub", "b.go")}
		if !slices.Equal(got, want) {
			t.Errorf("files = %q, want %q", got, want)
		}
		if got := listProjectFiles(root, 1); len(got) != 1 {
			t.Errorf("files over the limit: %q", got)
		}
	})

	t.Run("git", func(t *testing.T) {
		if _, err := exec.LookPath("git"); err != nil {
			t.Skip("no git")
		}
		if err := exec.Command("git", "-C", root, "init", "-q").Run(); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(root, ".gitignore"), []byte("*.log\nnode_modules/\n"), 0o600); err != nil {
			t.Fatal(err)
		}
		got := listProjectFiles(filepath.Join(root, "sub"), 100)
		if want := []string{"b.go"}; !slices.Equal(got, want) {
			t.Errorf("files in sub = %q, want %q", got, want)
		}
		got = listProjectFiles(root, 100)
		slices.Sort(got)
		want := []string{".git-like/c", ".gitignore", "a.go", "sub/b.go"}
		for i := range want {
			want[i] = filepath.FromSlash(want[i])
		}
		if !slices.Equal(got, want) {
			t.Errorf("files = %q, want %q", got, want)
		}
	})
}

func TestEditor_QuickOpenInput(t *testing.T) {
	newEditor := func(t *testing.T) (*Editor, string) {
		root := t.TempDir()
		for _, name := range []string{"a.txt", "b.txt", "c.txt"} {
			if err := os.WriteFile(filepath.Join(root, name), []byte(name), 0o600); err != nil {
				t.Fatal(err)
			}
		}
		e := newPickerTestEditor(root)
		for _, name := range []string{"a.txt", "b.txt"} {
			(&OpenFile{Path: filepath.Join(root, name)}).DoOnEditor(e)
		}
		return e, root
	}
	input := func(e *Editor, keys ...string) {
		for _, k := range keys {
			e.handleInput([]byte(k))
			if q := e.status.quick; q != nil && q.loading {
				runQueuedCommand(t, e)
			}
		}
	}
	const ctrlO = "\x0f"

	t.Run("switch to the previous buffer", func(t *testing.T) {
		e, root := newEditor(t)
		input(e, ctrlO)
		if top := e.Top(); top.mode != ModeCommand || top.cmdline != "e " || e.status.quick == nil {
			t.Fatalf("mode %s, cmdline %q", top.mode, top.cmdline)
		}
		input(e, "\r")
		assertTop(t, e, filepath.Join(root, "a.txt"), ModeNormal)
		input(e, ctrlO, "\r")
		assertTop(t, e, filepath.Join(root, "b.txt"), ModeNormal)
		assertBufferCount(t, e, 2)
	})

	t.Run("select a file", func(t *testing.T) {
		e, root := newEditor(t)
		input(e, ctrlO, "\t", "\t") // a.txt, b.txt, c.txt
		input(e, "\x1b[Z", ctrlO)   // back and forth
		input(e, "\r")
		assertTop(t, e, filepath.Join(root, "c.txt"), ModeNormal)
		assertBufferCount(t, e, 3)
	})

	t.Run("typed command", func(t *testing.T) {
		e, root := newEditor(t)
		input(e, ":", "e", " ", "c", "\r")
		assertTop(t, e, filepath.Join(root, "c.txt"), ModeNormal)
	})

	t.Run("backspace closes the picker", func(t *testing.T) {
		e, root := newEditor(t)
		input(e, ctrlO, "\x7f")
		if e.status.quick != nil {
			t.Error("picker is still shown")
		}
		assertTop(t, e, filepath.Join(root, "b.txt"), ModeCommand)
		input(e, "\x7f", "\x7f")
		assertTop(t, e, filepath.Join(root, "b.txt"), ModeNormal)
	})

	t.Run("cancel", func(t *testing.T) {
		e, root := newEditor(t)
		top := e.Top()
		top.offset = 3
		input(e, "i", ctrlO)
		if e.status.quick == nil {
			t.Fatal("no picker in insert mode")
		}
		e.status.quick.tall, top.offset = true, 5 // Rendered with two status rows.
		input(e, "\x1b")
		assertTop(t, e, filepath.Join(root, "b.txt"), ModeNormal)
		if e.status.quick != nil || top.offset != 3 {
			t.Errorf("picker %v, offset %d", e.status.quick, top.offset)
		}
	})

	t.Run("new file", func(t *testing.T) {
		e, root := newEditor(t)
		input(e, ctrlO)
		input(e, strings.Split("new.txt", "")...)
		if len(e.status.quick.matches) != 0 {
			t.Fatalf("matches: %v", e.status.quick.matches)
		}
		input(e, "\r")
		path := filepath.Join(root, "new.txt")
		assertTop(t, e, path, ModeNormal)
		if !e.Top().canEdit() {
			t.Error("cannot edit the new file")
		}
	})

	t.Run("open with the hook", func(t *testing.T) {
		e, root := newEditor(t)
		var opened []string
		e.OpenPath = func(path string) { opened = append(opened, path) }
		input(e, ctrlO, "\r") // An open buffer is activated.
		input(e, ctrlO, "c", "\r")
		if want := []string{filepath.Join(root, "c.txt")}; !slices.Equal(opened, want) {
			t.Errorf("opened %q, want %q", opened, want)
		}
	})
}

// newPickerTestEditor returns an editor with a command queue to receive the listed
// project files without running the event loop.
func newPickerTestEditor(root string) *Editor {
	return &Editor{Root: root, cmdChannel: make(chan Command, 1)}
}

// runQueuedCommand waits for a command posted to the editor and runs it like the event loop.
func runQueuedCommand(t *testing.T, e *Editor) {
	t.Helper()
	select {
	case cmd := <-e.cmdChannel:
		cmd.DoOnEditor(e)
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for a command")
	}
}

func assertTop(t *testing.T, e *Editor, path string, mode Mode) {
	t.Helper()
	top := e.Top()
	if top.Path != path || top.mode != mode {
		t.Errorf("top buffer %q in %s, want %q in %s", top.Path, top.mode, path, mode)
	}
}

func TestQuickOpen_Render(t *testing.T) {
	newPicker := func(n int) *quickOpen {
		q := &quickOpen{buf: &Buffer{mode: ModeCommand, cmdline: "e q"}}
		for i := range n {
			q.matches = append(q.matches, quickMatch{display: strings.Repeat(string(rune('a'+i)), 5)})
		}
		q.total = n + 10
		return q
	}
	render := func(q *quickOpen, w int) (rows []string, raw string) {
		var out bytes.Buffer
		q.render(&out, w)
		raw = out.String()
		return strings.Split(escape.Clean(raw), "\r\n"), raw
	}

	t.Run("one row", func(t *testing.T) {
		q := newPicker(3)
		q.sel = 1
		if h := q.height(40); h != 1 {
			t.Fatalf("height = %d", h)
		}
		rows, raw := render(q, 40)
		want := ":e q" + strings.Repeat(" ", 10) + " aaaaa  bbbbb  ccccc  +10 "
		if len(rows) != 1 || rows[0] != want {
			t.Errorf("rendered %q, want %q", rows, want)
		}
		if !strings.Contains(raw, "\x1b[0m bbbbb \x1b[7m") {
			t.Errorf("selected match is not highlighted: %q", raw)
		}
	})

	t.Run("two rows", func(t *testing.T) {
		q := newPicker(8)
		q.sel = 7
		if h := q.height(30); h != 2 {
			t.Fatalf("height = %d", h)
		}
		rows, _ := render(q, 30)
		// The selected match is the last one: the window is shifted to show it.
		want := []string{
			" fffff  ggggg  hhhhh  +10     ",
			":e q" + strings.Repeat(" ", 21) + "8/18 ",
		}
		if !slices.Equal(rows, want) {
			t.Errorf("rendered %q, want %q", rows, want)
		}

		// The height does not change back while the picker is shown.
		q.matches = q.matches[:1]
		if h := q.height(30); h != 2 {
			t.Errorf("height = %d", h)
		}
	})

	t.Run("no matches", func(t *testing.T) {
		q := newPicker(0)
		q.query, q.total = "q", 0
		rows, _ := render(q, 30)
		if want := ":e q" + strings.Repeat(" ", 17) + "new file "; len(rows) != 1 || rows[0] != want {
			t.Errorf("rendered %q, want %q", rows, want)
		}
	})

	t.Run("long path", func(t *testing.T) {
		q := newPicker(0)
		q.matches = []quickMatch{{display: strings.Repeat("x/", 30) + "file.go"}}
		q.total = 1
		rows, _ := render(q, 30)
		if want := ":e q     …x/x/x/x/x/x/file.go "; len(rows) != 1 || rows[0] != want {
			t.Errorf("rendered %q, want %q", rows, want)
		}
	})
}

func TestEditor_RenderQuickOpen(t *testing.T) {
	e := newPickerTestEditor(t.TempDir())
	lines := make(content.FullText, 100)
	for i := range lines {
		lines[i] = content.TextLine("line")
	}
	e.OpenBuffer(&Buffer{Path: "a.txt", Content: &lines})
	top := e.Top()
	top.c.Line = 38 // The last visible line with one status row.

	e.handleInput([]byte{0x0f})
	runQueuedCommand(t, e)
	e.status.quick.tall = true
	var out bytes.Buffer
	e.render(bufio.NewWriter(&out))
	if top.h != 38 || top.offset != 1 {
		t.Errorf("content height %d, offset %d", top.h, top.offset)
	}
	rows := strings.Split(escape.Clean(out.String()), "\r\n")
	if len(rows) != 40 || !strings.HasPrefix(rows[39], ":e ") {
		t.Errorf("rendered %d rows, the last one %q", len(rows), rows[len(rows)-1])
	}

	e.handleInput([]byte{0x1b})
	for range e.layout() {
	}
	if top.h != 39 || top.offset != 0 {
		t.Errorf("content height %d, offset %d after the picker is closed", top.h, top.offset)
	}
}
