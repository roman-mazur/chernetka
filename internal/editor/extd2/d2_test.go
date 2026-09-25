package extd2

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/cheimg"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl"
	"rmazur.io/chernetka/internal/logger"
)

// recordingViewer collects the diagrams it's asked to show.
type recordingViewer chan cheimg.Item

func (rv recordingViewer) Show(d cheimg.Item) { rv <- d }

func (rv recordingViewer) receive(t *testing.T) cheimg.Item {
	t.Helper()
	select {
	case d := <-rv:
		return d
	case <-time.After(time.Second):
		t.Fatal("no diagram shown")
		return cheimg.Item{}
	}
}

func newIntegration(t *testing.T) (*Integration, recordingViewer) {
	viewer := make(recordingViewer, 1)
	ext := &Integration{Viewer: viewer}
	ext.LogEmbed = logger.Embed(t.Logf)
	return ext, viewer
}

// openDoc opens a buffer in an editor with the extension and the syntax highlighter providing the code blocks.
func openDoc(t *testing.T, ext *Integration, path, text string) *editor.Buffer {
	t.Helper()
	var edit editor.Editor
	edit.Extend(new(extsyntaxhl.Integration))
	edit.Extend(ext)
	if err := edit.OpenReader(path, strings.NewReader(text)); err != nil {
		t.Fatal(err)
	}
	return edit.Top()
}

// actionLines returns the indexes of the lines the extension provides actions for.
func actionLines(t *testing.T, buf *editor.Buffer, ext *Integration) []int {
	t.Helper()
	actions, ok := buf.ExtensionData(ext.ID()).(content.LineActions)
	if !ok {
		t.Fatalf("no line actions for %q", buf.Path)
	}
	var res []int
	for i := range buf.Content.Len() + 1 {
		if actions.LineAction(i) != nil {
			res = append(res, i)
		}
	}
	return res
}

func engage(t *testing.T, buf *editor.Buffer, ext *Integration, lineNumber int) {
	t.Helper()
	action := buf.ExtensionData(ext.ID()).(content.LineActions).LineAction(lineNumber)
	if action == nil {
		t.Fatalf("no action on line %d", lineNumber)
	}
	action.Engage()
}

func absPath(t *testing.T, p string) string {
	t.Helper()
	res, err := filepath.Abs(p)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func TestDiagramFile(t *testing.T) {
	ext, viewer := newIntegration(t)
	const src = "a -> b\nb -> c"
	buf := openDoc(t, ext, "docs/arch.D2", src)

	if got := actionLines(t, buf, ext); len(got) != 1 || got[0] != 0 {
		t.Errorf("action lines = %v, want [0]", got)
	}

	engage(t, buf, ext, 0)
	want := cheimg.Item{Path: absPath(t, "docs/arch.D2"), Source: src}
	if got := viewer.receive(t); got != want {
		t.Errorf("shown %+v, want %+v", got, want)
	}

	// Unsaved changes are shown.
	buf.Mutate().Insert(2, content.TextLine("c -> d"))
	engage(t, buf, ext, 0)
	if got := viewer.receive(t); got.Source != src+"\nc -> d" {
		t.Errorf("shown source %q", got.Source)
	}
}

func TestMarkdown(t *testing.T) {
	ext, viewer := newIntegration(t)
	buf := openDoc(t, ext, "/notes/README.md", strings.Join([]string{
		"# Notes", // 0
		"```d2",   // 1
		"a -> b",  // 2
		"b -> c",  // 3
		"```",     // 4
		"```go",   // 5
		"```",     // 6
		"~~~ d2",  // 7
		"x -> y",  // 8
	}, "\n"))

	if got, want := actionLines(t, buf, ext), []int{1, 7}; !slices.Equal(got, want) {
		t.Errorf("action lines = %v, want %v", got, want)
	}

	engage(t, buf, ext, 1)
	want := cheimg.Item{Path: "/notes/README.md", Source: "a -> b\nb -> c"}
	if got := viewer.receive(t); got != want {
		t.Errorf("shown %+v, want %+v", got, want)
	}

	engage(t, buf, ext, 7)
	if got := viewer.receive(t); got.Source != "x -> y" {
		t.Errorf("shown source of the unterminated block %q", got.Source)
	}

	// Edits move the blocks.
	buf.Mutate().Insert(0, content.TextLine("```d2"))
	buf.Mutate().Insert(1, content.TextLine("```"))
	ext.AfterEdit(nil, buf)
	if got, want := actionLines(t, buf, ext), []int{0, 3, 9}; !slices.Equal(got, want) {
		t.Errorf("action lines after edit = %v, want %v", got, want)
	}

	// The block is gone. Its closing fence opens a block without a language now.
	buf.Mutate().Update(3, content.TextLine("text"))
	ext.AfterEdit(nil, buf)
	if got, want := actionLines(t, buf, ext), []int{0, 9}; !slices.Equal(got, want) {
		t.Errorf("action lines after edit = %v, want %v", got, want)
	}
}

func TestNoActions(t *testing.T) {
	t.Run("other files", func(t *testing.T) {
		ext, _ := newIntegration(t)
		buf := openDoc(t, ext, "main.go", "```d2\n```")
		if data := buf.ExtensionData(ext.ID()); data != nil {
			t.Errorf("got extension data %#v", data)
		}
	})

	t.Run("no viewer", func(t *testing.T) {
		ext := new(Integration)
		buf := openDoc(t, ext, "a.d2", "a -> b")
		if data := buf.ExtensionData(ext.ID()); data != nil {
			t.Errorf("got extension data %#v", data)
		}
	})

	t.Run("markdown without syntax highlighter", func(t *testing.T) {
		ext, _ := newIntegration(t)
		var edit editor.Editor
		edit.Extend(ext)
		if err := edit.OpenReader("a.md", strings.NewReader("```d2\na -> b\n```")); err != nil {
			t.Fatal(err)
		}
		if got := actionLines(t, edit.Top(), ext); got != nil {
			t.Errorf("action lines = %v, want none", got)
		}
	})

	t.Run("directory listing", func(t *testing.T) {
		dir := filepath.Join(t.TempDir(), "diagrams.d2")
		if err := os.MkdirAll(dir, 0o700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, "a.d2"), []byte("a"), 0o600); err != nil {
			t.Fatal(err)
		}
		ext, _ := newIntegration(t)
		var edit editor.Editor
		edit.Extend(ext)
		edit.OpenDir(dir, nil)
		if data := edit.Top().ExtensionData(ext.ID()); data != nil {
			t.Errorf("got extension data %#v for a directory listing", data)
		}
	})
}

// TestEngageWithEnter checks the wiring with the editor: Enter in normal mode engages the action.
func TestEngageWithEnter(t *testing.T) {
	ext, viewer := newIntegration(t)
	h := editor.NewTestHarness()
	h.Extend(new(extsyntaxhl.Integration))
	h.Extend(ext)
	if err := h.OpenReader("a.md", strings.NewReader("text\n```d2\na -> b\n```")); err != nil {
		t.Fatal(err)
	}
	h.Run(t)

	h.SendInput(t, []byte{'\r'}) // Moves the cursor down from a plain line.
	h.SendInput(t, []byte{'\r'})
	if got := viewer.receive(t); got.Source != "a -> b" {
		t.Errorf("shown source %q", got.Source)
	}
}

// TestMarkdown_ExtensionsOrder checks that the blocks are up to date even if the syntax highlighter
// is registered after this extension and has not been notified about an edit yet.
func TestMarkdown_ExtensionsOrder(t *testing.T) {
	ext, _ := newIntegration(t)
	var edit editor.Editor
	edit.Extend(ext)
	edit.Extend(new(extsyntaxhl.Integration))
	if err := edit.OpenReader("a.md", strings.NewReader("```d2\na -> b\n```")); err != nil {
		t.Fatal(err)
	}
	buf := edit.Top()
	if got, want := actionLines(t, buf, ext), []int{0}; !slices.Equal(got, want) {
		t.Errorf("action lines = %v, want %v", got, want)
	}

	buf.Mutate().Insert(0, content.TextLine("# Title"))
	ext.AfterEdit(nil, buf)
	if got, want := actionLines(t, buf, ext), []int{1}; !slices.Equal(got, want) {
		t.Errorf("action lines after edit = %v, want %v", got, want)
	}
}
