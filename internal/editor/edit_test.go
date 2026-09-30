package editor

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/input"
	"rmazur.io/chernetka/internal/vt"
	"rmazur.io/chernetka/internal/vt/escape"
)

func TestEditor_OpenReader(t *testing.T) {
	var e Editor

	if err := e.OpenReader("first.txt", strings.NewReader("hello\nworld")); err != nil {
		t.Fatalf("OpenReader: %v", err)
	}
	assertLayoutTop(t, &e, "first.txt")
	assertBufferCount(t, &e, 1)

	if err := e.OpenReader("second.txt", strings.NewReader("foo")); err != nil {
		t.Fatalf("OpenReader: %v", err)
	}
	assertLayoutTop(t, &e, "second.txt")
	assertBufferCount(t, &e, 2)
}

func TestEditor_OpenReader_ReuseExisting(t *testing.T) {
	var e Editor

	for _, name := range []string{"a.txt", "b.txt", "c.txt"} {
		if err := e.OpenReader(name, strings.NewReader(name)); err != nil {
			t.Fatalf("OpenReader(%q): %v", name, err)
		}
	}
	// Stack top→bottom: c.txt, b.txt, a.txt
	assertLayoutTop(t, &e, "c.txt")
	assertBufferCount(t, &e, 3)

	// Re-open a non-top buffer — should move to top, not create a new entry.
	if err := e.OpenReader("a.txt", strings.NewReader("ignored")); err != nil {
		t.Fatalf("OpenReader(a.txt): %v", err)
	}
	assertLayoutTop(t, &e, "a.txt")
	assertBufferCount(t, &e, 3)

	// Re-opening the already-top buffer is a no-op.
	if err := e.OpenReader("a.txt", strings.NewReader("ignored")); err != nil {
		t.Fatalf("OpenReader(a.txt) no-op: %v", err)
	}
	assertLayoutTop(t, &e, "a.txt")
	assertBufferCount(t, &e, 3)
}

func TestEditor_OpenReader_ReuseSamePath(t *testing.T) {
	t.Chdir(t.TempDir())
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	var e Editor
	for _, name := range []string{"a.txt", "b.txt", filepath.Join(wd, "a.txt"), "./a.txt"} {
		if err := e.OpenReader(name, strings.NewReader(name)); err != nil {
			t.Fatalf("OpenReader(%q): %v", name, err)
		}
	}
	assertLayoutTop(t, &e, "a.txt")
	assertBufferCount(t, &e, 2)
}

func assertLayoutTop(t *testing.T, e *Editor, wantPath string) {
	t.Helper()
	var bufs []*Buffer
	for buf := range e.layout() {
		bufs = append(bufs, buf)
	}
	if len(bufs) != 1 {
		t.Fatalf("layout returned %d buffers, want 1", len(bufs))
	}
	if bufs[0].Path != wantPath {
		t.Errorf("layout top path = %q, want %q", bufs[0].Path, wantPath)
	}
}

// assertBufferCount walks the stack via next pointers and fails if the count
// doesn't match or exceeds a safety cap (which would indicate a cycle).
func assertBufferCount(t *testing.T, e *Editor, want int) {
	t.Helper()
	const maxSafe = 1000
	n := 0
	for range e.buffers() {
		n++
		if n > maxSafe {
			t.Fatalf("buffer count exceeds %d — likely a cycle in the stack", maxSafe)
		}
	}
	if n != want {
		t.Errorf("buffer count = %d, want %d", n, want)
	}
}

func TestEditor_Run(t *testing.T) {
	src, err := os.Open("edit.go")
	if err != nil {
		t.Fatalf("cannot open test input: %s", err)
	}
	t.Cleanup(func() { _ = src.Close() })

	te := NewTestHarness()
	if err := te.OpenReader("test.txt", src); err != nil {
		t.Fatalf("OpenReader: %s", err)
	}
	te.Run(t)

	t.Run("scrolling", func(t *testing.T) {
		te.SendInput(t, []byte("\x1b[B")) // Cursor down.
		var cursorY int
		te.Post(t, CommandFunc(func(e *Editor) {
			cursorY = e.Top().c.Line
		}))
		if cursorY != 1 {
			t.Errorf("cursor y = %d, want 1 after cursor down", cursorY)
		}

		for range 5 {
			te.SendInput(t, []byte(input.Scroll(input.ScrollDirectionDown, 10, 20).Encode()))
		}
		var offset int
		te.Post(t, CommandFunc(func(e *Editor) {
			offset = e.Top().offset
		}))
		if offset != 5 {
			t.Errorf("offset = %d, some scroll down events seem to be ignored", offset)
		}
	})

	t.Run("clipboard paste", func(t *testing.T) {
		// Land the cursor at a known position for a deterministic assertion.
		te.Post(t, CommandFunc(func(e *Editor) {
			e.Top().c = content.Position{}
		}))

		te.SendInput(t, []byte("\x1b[200~pasted \x1b[201~"))

		var line0 string
		var cursor content.Position
		te.Post(t, CommandFunc(func(e *Editor) {
			line0 = e.Top().Content.Lines()[0].String()
			cursor = e.Top().c
		}))
		if want := "pasted "; !strings.HasPrefix(line0, want) {
			t.Errorf("line 0 = %q, want prefix %q", line0, want)
		}
		if want := (content.Position{Col: len("pasted ")}); cursor != want {
			t.Errorf("cursor after paste = %+v, want %+v", cursor, want)
		}
	})

	t.Run("clipboard paste split across reads", func(t *testing.T) {
		te.Post(t, CommandFunc(func(e *Editor) {
			e.Top().c = content.Position{}
		}))

		// Split both the pasted text and the terminator marker across
		// separate writes, exercising the reader continuation that
		// readAndHandleInput relies on to reassemble a paste that
		// straddles two terminal reads.
		te.WriteInput(t, []byte("\x1b[200~spl"))
		te.WriteInput(t, []byte("it\x1b[2"))
		te.SendInput(t, []byte("01~"))

		var line0 string
		te.Post(t, CommandFunc(func(e *Editor) {
			line0 = e.Top().Content.Lines()[0].String()
		}))
		if want := "split"; !strings.HasPrefix(line0, want) {
			t.Errorf("line 0 = %q, want prefix %q", line0, want)
		}
	})

	t.Run("empty clipboard paste is a no-op", func(t *testing.T) {
		var before string
		te.Post(t, CommandFunc(func(e *Editor) {
			before = e.Top().Text()
		}))

		te.SendInput(t, []byte("\x1b[200~\x1b[201~"))

		var after string
		te.Post(t, CommandFunc(func(e *Editor) {
			after = e.Top().Text()
		}))
		if before != after {
			t.Errorf("buffer changed after an empty paste:\nbefore: %q\nafter:  %q", before, after)
		}
	})

	t.Run("select line", func(t *testing.T) {
		te.SendInput(t, []byte("\x1b[1;10D"))
		var selText string
		te.Post(t, CommandFunc(func(e *Editor) {
			selText = e.Top().SelectedText()
		}))
		if selText == "" {
			t.Errorf("selected text is empty")
		}
		te.Post(t, CommandFunc(func(e *Editor) {
			e.execBufferCmd(MoveHome)
		}))
	})

	t.Run("clipboard copy", func(t *testing.T) {
		te.SendInputSequence(t, "ihe") // Insert mode, then "he"
		te.Post(t, CommandFunc(func(e *Editor) {
			// Select these letters.
			e.Top().sel = []content.Span{
				{End: content.Position{Col: 2}},
			}
		}))
		te.SendInput(t, []byte{0x3}) // Ctrl+C

		clipData := clipboard.Read()
		t.Log("clipboard data:", clipData)
		if clipData != "he" {
			t.Errorf("clipboard not copied")
		}
	})
}

// TestEditor_LayoutWindowSize covers the terminal size resolution used by layout.
// A horizontally split terminal pane is shorter than the mock fallback height, and
// the editor used to keep that fallback because it treated fd 0 (the tty arriving on
// stdin) as "no terminal". It then rendered more rows than the pane could hold,
// scrolling the top of the buffer out of view.
func TestEditor_LayoutWindowSize(t *testing.T) {
	for _, tc := range []struct {
		name         string
		termSize     func() (vt.WindowSize, error)
		wantW, wantH int
	}{
		{
			name:  "no terminal",
			wantW: 80, wantH: 40,
		},
		{
			name:     "split pane on stdin",
			termSize: func() (vt.WindowSize, error) { return vt.WindowSize{Cols: 120, Rows: 21}, nil },
			wantW:    120, wantH: 21,
		},
		{
			name:     "size query fails",
			termSize: func() (vt.WindowSize, error) { return vt.WindowSize{}, os.ErrInvalid },
			wantW:    80, wantH: 42,
		},
		{
			name:     "zero size reported",
			termSize: func() (vt.WindowSize, error) { return vt.WindowSize{}, nil },
			wantW:    80, wantH: 42,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var e Editor
			if tc.termSize != nil {
				term := vt.TestTerminal(0, 0, nil)
				term.SizeFunc = tc.termSize
				e.term = term
			}
			if err := e.OpenReader("a.txt", strings.NewReader("hello\nworld")); err != nil {
				t.Fatalf("OpenReader: %v", err)
			}

			for buf := range e.layout() {
				// One row is left for the status bar.
				if buf.w != tc.wantW || buf.h != tc.wantH-1 {
					t.Errorf("buffer size = %dx%d, want %dx%d", buf.w, buf.h, tc.wantW, tc.wantH-1)
				}
			}
		})
	}
}

func TestEditor_Run_ConfiguresTerminal(t *testing.T) {
	var out bytes.Buffer
	term := vt.TestTerminal(80, 40, struct {
		io.Reader
		io.Writer
	}{strings.NewReader(""), &out}) // Input EOF makes the editor quit.

	var e Editor
	runFinished := make(chan struct{})
	go func() {
		defer close(runFinished)
		e.Run(term)
	}()
	select {
	case <-runFinished:
	case <-time.After(5 * time.Second):
		t.Fatal("Run did not finish on input EOF")
	}
	got := out.String()

	for _, tc := range []struct {
		name string
		f    escape.ConfigFunc
	}{
		{"mouse", escape.EnableMouse},
		{"alternative buffer", escape.EnableAlternativeBuffer},
		{"line wrapping", escape.DisableLineWrapping},
		{"bracketed paste", escape.EnableBracketedPasteMode},
		{"focus reporting", escape.EnableFocusReporting},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var seq bytes.Buffer
			restore := tc.f(&seq)
			enable := seq.String()
			seq.Reset()
			restore()
			disable := seq.String()

			on := strings.Index(got, enable)
			if on < 0 {
				t.Fatalf("terminal is not configured: missing %q in the output", enable)
			}
			if off := strings.LastIndex(got, disable); off < on {
				t.Errorf("terminal is not restored on exit: missing %q after %q", disable, enable)
			}
		})
	}
}

func TestEditor_ShowDiff(t *testing.T) {
	const ctrlD = "\x04"
	dir := t.TempDir()
	path := filepath.Join(dir, "a.txt")
	if err := os.WriteFile(path, []byte("a"), 0o600); err != nil {
		t.Fatal(err)
	}
	newEditor := func() (*Editor, *[]string) {
		var shown []string
		e := &Editor{ShowDiff: func(path string) { shown = append(shown, path) }}
		return e, &shown
	}

	t.Run("file in any mode", func(t *testing.T) {
		e, shown := newEditor()
		(&OpenFile{Path: path}).DoOnEditor(e)
		e.handleInput([]byte(ctrlD))
		e.Top().mode = ModeInsert
		e.handleInput([]byte(ctrlD))
		if len(*shown) != 2 || (*shown)[0] != path || (*shown)[1] != path {
			t.Errorf("shown %q, want %q twice", *shown, path)
		}
		if got := e.Top().Content.Lines()[0].String(); got != "a" {
			t.Errorf("Ctrl+D changed the content to %q", got)
		}
	})

	t.Run("changes are saved first", func(t *testing.T) {
		var saved []string
		e := &Editor{ShowDiff: func(path string) {
			data, _ := os.ReadFile(path)
			saved = append(saved, string(data))
		}}
		(&OpenFile{Path: path}).DoOnEditor(e)
		t.Cleanup(func() { _ = os.WriteFile(path, []byte("a"), 0o600) })
		for _, k := range []string{"i", "b", ctrlD} {
			e.handleInput([]byte(k))
		}
		if len(saved) != 1 || saved[0] != "ba" {
			t.Errorf("diff shown for %q, want the saved %q", saved, "ba")
		}
		if e.Top().dirty {
			t.Error("buffer is still dirty")
		}
	})

	t.Run("relative path is shown as absolute", func(t *testing.T) {
		t.Chdir(dir)
		e, shown := newEditor()
		(&OpenFile{Path: "a.txt"}).DoOnEditor(e)
		e.handleInput([]byte(ctrlD))
		if len(*shown) != 1 || (*shown)[0] != path {
			t.Errorf("shown %q, want %q", *shown, path)
		}
	})

	t.Run("buffers without a file", func(t *testing.T) {
		e, shown := newEditor()
		e.New()
		e.handleInput([]byte(ctrlD))
		e.OpenDir(dir, nil)
		e.handleInput([]byte(ctrlD))
		if len(*shown) != 0 {
			t.Errorf("shown %q", *shown)
		}
	})

	t.Run("no hook", func(t *testing.T) {
		var e Editor
		(&OpenFile{Path: path}).DoOnEditor(&e)
		e.handleInput([]byte(ctrlD)) // Must not panic.
	})
}

func TestEditor_AutoSave(t *testing.T) {
	dir := t.TempDir()
	writeFile := func(t *testing.T, name string) string {
		t.Helper()
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte("a"), 0o600); err != nil {
			t.Fatal(err)
		}
		return path
	}
	readFile := func(t *testing.T, path string) string {
		t.Helper()
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		return string(data)
	}
	edit := func(e *Editor) {
		for _, k := range []string{"i", "b", "\x1b"} {
			e.handleInput([]byte(k))
		}
	}

	t.Run("focus out", func(t *testing.T) {
		path := writeFile(t, "focus.txt")
		var e Editor
		(&OpenFile{Path: path}).DoOnEditor(&e)
		edit(&e)
		e.handleInput([]byte("\x1b[I"))
		if got := readFile(t, path); got != "a" {
			t.Errorf("saved on focus in: %q", got)
		}
		e.handleInput([]byte("\x1b[O"))
		if got := readFile(t, path); got != "ba" {
			t.Errorf("file content = %q, want %q", got, "ba")
		}
		if e.Top().dirty {
			t.Error("buffer is still dirty")
		}
	})

	t.Run("open another buffer", func(t *testing.T) {
		path := writeFile(t, "push.txt")
		var e Editor
		(&OpenFile{Path: path}).DoOnEditor(&e)
		edit(&e)
		e.New()
		if got := readFile(t, path); got != "ba" {
			t.Errorf("file content = %q, want %q", got, "ba")
		}
	})

	t.Run("select another buffer", func(t *testing.T) {
		path1, path2 := writeFile(t, "sel1.txt"), writeFile(t, "sel2.txt")
		var e Editor
		(&OpenFile{Path: path1}).DoOnEditor(&e)
		(&OpenFile{Path: path2}).DoOnEditor(&e)
		edit(&e)
		(&OpenFile{Path: path1}).DoOnEditor(&e)
		if e.Top().Path != path1 {
			t.Fatalf("top buffer is %q, want %q", e.Top().Path, path1)
		}
		if got := readFile(t, path2); got != "ba" {
			t.Errorf("file content = %q, want %q", got, "ba")
		}
	})

	t.Run("scratch buffer", func(t *testing.T) {
		var e Editor
		e.New()
		edit(&e)
		e.handleInput([]byte("\x1b[O")) // Must not panic.
		e.New()
		if !e.bufs[0].dirty {
			t.Error("scratch buffer is not dirty anymore")
		}
	})
}

// actionsExt is an Extension providing actions for the listed lines of every buffer.
type actionsExt struct {
	noopExt
	actions testActionsExt
}

func (ae *actionsExt) MakeBufferData(*Buffer) BufferExtData { return ae.actions }

type noopExt struct{}

func (noopExt) ID() string                { return "actions" }
func (noopExt) AfterEdit(Sender, *Buffer) {}

func TestEditor_RerunAction(t *testing.T) {
	const ctrlR = 0x12
	first, second := new(countingAction), new(countingAction)
	h := NewTestHarness()
	h.Extend(&actionsExt{actions: testActionsExt{0: first, 2: second}})
	if err := h.OpenReader("a.txt", strings.NewReader("action\ntext\naction")); err != nil {
		t.Fatal(err)
	}
	h.Run(t)

	check := func(wantFirst, wantSecond int) {
		t.Helper()
		// Commands are drained by SendInput: the counters are not modified concurrently.
		if first.engaged != wantFirst || second.engaged != wantSecond {
			t.Errorf("engaged %d and %d times, want %d and %d",
				first.engaged, second.engaged, wantFirst, wantSecond)
		}
	}

	h.SendInput(t, []byte{ctrlR})
	check(0, 0) // Nothing to re-run yet.

	h.SendInput(t, []byte{'\r'})
	check(1, 0)

	h.SendInputSequence(t, "jj")
	h.SendInput(t, []byte{ctrlR})
	check(2, 0) // The cursor position does not matter.

	h.SendInput(t, []byte{'\r'})
	h.SendInput(t, []byte{ctrlR})
	check(2, 2)

	h.SendInput(t, []byte{'i'})
	h.SendInput(t, []byte{ctrlR})
	check(2, 3) // Works in insert mode too.

	h.Post(t, CommandFunc(func(e *Editor) {
		if text := e.Top().Text(); text != "action\ntext\naction" {
			t.Errorf("Ctrl+R changed the text: %q", text)
		}
	}))
}

// TestEditor_RerunAction_CloseBuffer checks that closing a buffer forgets the action
// engaged in it, but keeps the one engaged in another buffer.
func TestEditor_RerunAction_CloseBuffer(t *testing.T) {
	const ctrlR = 0x12
	a, b := new(countingAction), new(countingAction)
	h := NewTestHarness()
	h.Extend(&pathActionsExt{actions: map[string]*countingAction{"a.txt": a, "b.txt": b}})
	for _, path := range []string{"a.txt", "b.txt"} {
		if err := h.OpenReader(path, strings.NewReader("action")); err != nil {
			t.Fatal(err)
		}
	}
	h.Run(t)

	check := func(wantA, wantB int) {
		t.Helper()
		// Commands are drained by SendInput: the counters are not modified concurrently.
		if a.engaged != wantA || b.engaged != wantB {
			t.Errorf("engaged %d and %d times, want %d and %d", a.engaged, b.engaged, wantA, wantB)
		}
	}

	h.SendInput(t, []byte{'\r'})
	check(0, 1)

	h.Post(t, CommandFunc(func(e *Editor) {
		if err := e.OpenReader("c.txt", strings.NewReader("no action")); err != nil {
			t.Error(err)
		}
	}))
	h.SendInput(t, []byte{ctrlR})
	check(0, 2) // Re-run from another buffer.

	h.SendInput(t, []byte{'q'})
	h.SendInput(t, []byte{ctrlR})
	check(0, 3) // Closing another buffer keeps the action.

	h.SendInput(t, []byte{'q'})
	h.SendInput(t, []byte{ctrlR})
	check(0, 3) // Closing its buffer forgets the action.
}

// pathActionsExt provides the actions for the first lines of the buffers by their paths.
type pathActionsExt struct {
	noopExt
	actions map[string]*countingAction
}

func (pe *pathActionsExt) MakeBufferData(buf *Buffer) BufferExtData {
	if action, ok := pe.actions[buf.Path]; ok {
		return testActionsExt{0: action}
	}
	return nil
}

// singleShotAction is a countingAction that is not re-run.
type singleShotAction struct{ countingAction }

func (*singleShotAction) SingleShot() {}

// singleShotExt provides a countingAction on the first line and a singleShotAction on the second.
type singleShotExt struct {
	noopExt
	rerun  *countingAction
	single *singleShotAction
}

func (se *singleShotExt) MakeBufferData(*Buffer) BufferExtData { return se }

func (se *singleShotExt) LineAction(lineNumber int) content.LineAction {
	switch lineNumber {
	case 0:
		return se.rerun
	case 1:
		return se.single
	}
	return nil
}

func TestEditor_RerunAction_SkipsSingleShot(t *testing.T) {
	const ctrlR = 0x12
	rerun, single := new(countingAction), new(singleShotAction)
	h := NewTestHarness()
	h.Extend(&singleShotExt{rerun: rerun, single: single})
	if err := h.OpenReader("a.txt", strings.NewReader("action\nsingle")); err != nil {
		t.Fatal(err)
	}
	h.Run(t)

	h.SendInput(t, []byte{'\r'}) // Engage the first action, the cursor stays.
	h.SendInputSequence(t, "j")
	h.SendInput(t, []byte{'\r'})
	h.SendInput(t, []byte{ctrlR})
	if rerun.engaged != 2 || single.engaged != 1 {
		t.Errorf("engaged %d and %d times, want 2 and 1", rerun.engaged, single.engaged)
	}
}

func TestEditor_EngageSaves(t *testing.T) {
	path := filepath.Join(t.TempDir(), "a.txt")
	if err := os.WriteFile(path, []byte("a\nb"), 0o600); err != nil {
		t.Fatal(err)
	}
	action := new(countingAction)
	e := new(Editor)
	e.Extend(&actionsExt{actions: testActionsExt{1: action}})
	(&OpenFile{Path: path}).DoOnEditor(e)

	saved := func() string {
		data, _ := os.ReadFile(path)
		return string(data)
	}
	for _, k := range []string{"x", "j", "\r"} {
		e.handleInput([]byte(k))
	}
	if action.engaged != 1 || saved() != "\nb" {
		t.Errorf("action engaged %d times with the file %q, want once with the changes saved", action.engaged, saved())
	}
	e.handleInput([]byte("x"))
	e.handleInput([]byte{0x12}) // Ctrl+R
	if action.engaged != 2 || saved() != "\n" {
		t.Errorf("re-run action engaged %d times with the file %q, want twice with the changes saved", action.engaged, saved())
	}
}
