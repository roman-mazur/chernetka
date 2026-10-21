package extlsp

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.lsp.dev/protocol"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
)

// The edits of the server are undone separately from the typing just before them.

func TestIntegration_UndoAcceptedSuggestion(t *testing.T) {
	text := "package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\tstrings.S\n}"
	var le Integration
	_, buf, data := newGoBuffer(t, &le, &fakeLSP{}, text)
	at := content.Position{Line: 7, Col: len("\tstrings.S")}
	buf.ReplaceText(content.Span{Start: at, End: at}, "p") // Typed shortly before accepting.
	typed := buf.Text()

	data.assign([]suggestion{{
		text:   "lit",
		cursor: 3,
		label:  "Split",
		edits: []protocol.TextEdit{{
			Range:   protocol.Range{Start: protocol.Position{Line: 3, Character: 6}, End: protocol.Position{Line: 3, Character: 6}},
			NewText: "\n\t\"strings\"",
		}},
	}}, anchor{line: "\tstrings.Sp", cx: 11, cy: 7})
	if !le.HandleInsertKey(buf, keyTab) {
		t.Fatal("tab not handled")
	}
	accepted := buf.Text()

	checkUndoSteps(t, buf, accepted, typed, text)
}

func TestIntegration_UndoFormatting(t *testing.T) {
	t.Chdir(t.TempDir())
	fake := &fakeLSP{format: func(text string) string {
		return strings.ReplaceAll(text, "x:=1", "x := 1")
	}}
	var le Integration
	text := "package main\n\nfunc main() {\n\tx:=\n}\n"
	h, buf, _ := newGoBuffer(t, &le, fake, text)
	for range 3 {
		editor.RelMove{Dy: 1}.DoOnBuffer(buf, editor.RenderPrefs{})
	}
	h.MoveCursorToLineEnd()
	h.SetMode(editor.ModeInsert)
	fake.waitVersion(t, 1) // The server is ready.

	h.Run(t)
	h.SendInput(t, []byte("1"))
	h.SendInput(t, keySave) // Shortly after typing.

	want := "package main\n\nfunc main() {\n\tx := 1\n}\n"
	if got := waitSaved(t, "completion_buf.go"); got != want {
		t.Fatalf("saved\n%s\nwant\n%s", got, want)
	}
	h.Post(t, editor.CommandFunc(func(*editor.Editor) {
		checkUndoSteps(t, buf, want, "package main\n\nfunc main() {\n\tx:=1\n}\n", text)
	}))
}

// checkUndoSteps checks that the buffer has the texts one by one when undoing its changes,
// starting from the current one, and that redoing them returns to it.
func checkUndoSteps(t *testing.T, buf *editor.Buffer, texts ...string) {
	t.Helper()
	ur, ok := buf.Mutate().(content.UndoRedo)
	if !ok {
		t.Fatal("the buffer does not support undo")
	}
	check := func(what, want string) {
		t.Helper()
		if diff := cmp.Diff(want, buf.Text()); diff != "" {
			t.Errorf("text %s (-want +got):\n%s", what, diff)
		}
	}
	check("before undoing", texts[0])
	for i, want := range texts[1:] {
		editor.Undo.DoOnBuffer(buf, editor.RenderPrefs{})
		check("after undo #"+strconv.Itoa(i+1), want)
	}
	if ur.HasUndo() {
		t.Error("more changes to undo")
	}
	for range texts[1:] {
		editor.Redo.DoOnBuffer(buf, editor.RenderPrefs{})
	}
	check("after redoing all", texts[0])
}
