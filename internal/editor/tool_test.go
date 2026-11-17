package editor

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/change"
	"rmazur.io/chernetka/internal/editor/navigation"
)

// newToolEditor returns an editor showing a 100-line file with a listing of the
// jumps to its lines in the tool buffer.
func newToolEditor(t *testing.T) (*Editor, *content.Tree) {
	t.Helper()
	e := Editor{cmdChannel: make(chan Command, 16)}
	lines := make(content.FullText, 100)
	for i := range lines {
		lines[i] = content.TextLine(fmt.Sprintf("line %d", i+1))
	}
	e.OpenBuffer(&Buffer{Path: "a.txt", Content: &lines})
	tree := content.NewTree(&content.TreeItem{Text: "2 usages"})
	for _, ln := range []int{10, 50} {
		goTo := &GoTo{Location: navigation.Location{Path: "a.txt", Position: content.Position{Line: ln - 1, Col: 2}}}
		tree.Add(&content.TreeItem{
			Text:   fmt.Sprintf("%d: line %d", ln, ln),
			Action: content.LineActionFunc(func() { go e.Send(goTo) }), // Like the extensions do.
		})
	}
	(&ShowTool{Content: tree}).DoOnEditor(&e)
	return &e, tree
}

// runSent runs the commands sent to the editor from the actions, waiting for them shortly.
func runSent(e *Editor) {
	for {
		select {
		case cmd := <-e.cmdChannel:
			cmd.DoOnEditor(e)
		case <-time.After(50 * time.Millisecond):
			return
		}
	}
}

func TestEditor_ToolBuffer(t *testing.T) {
	e, tree := newToolEditor(t)
	main := e.Top()
	if e.toolBuf == nil || e.toolBuf.Content != tree || e.focused() != e.toolBuf {
		t.Fatalf("tool buffer %+v focused %t", e.toolBuf, e.toolFocused)
	}

	// The movement keys move in the tool buffer, Enter jumps in the main buffer
	// keeping the focus, Esc returns it, Ctrl+T switches it.
	e.handleInput([]byte("j"))
	e.handleInput([]byte("\r"))
	runSent(e)
	if got := main.c(); got != pos(9, 2) {
		t.Errorf("main cursor %s after Enter on the 1st usage", got)
	}
	if e.focused() != e.toolBuf || e.toolBuf.c() != pos(1, 0) {
		t.Errorf("focused %v at %s after the jump", e.focused() == e.toolBuf, e.toolBuf.c())
	}
	e.handleInput([]byte("G"))
	e.handleInput([]byte("\r"))
	runSent(e)
	if got := main.c(); got != pos(49, 2) {
		t.Errorf("main cursor %s after Enter on the 2nd usage", got)
	}
	e.handleInput([]byte("\x1b"))
	if e.focused() != main {
		t.Error("Esc did not return the focus")
	}
	e.handleInput([]byte("j"))
	if got := main.c(); got != pos(50, 2) {
		t.Errorf("main cursor %s after j with the main buffer focused", got)
	}
	e.handleInput([]byte{0x14}) // Ctrl+T
	if e.focused() != e.toolBuf {
		t.Error("Ctrl+T did not focus the tool buffer")
	}
	e.handleInput([]byte("k"))
	if e.toolBuf.c() != pos(1, 0) || main.c() != pos(50, 2) {
		t.Errorf("k moved: tool at %s, main at %s", e.toolBuf.c(), main.c())
	}

	// The tool buffer is not edited or quit with the keys of the review mode.
	e.handleInput([]byte("ix"))
	if e.toolBuf.mode != ModeReview || tree.Lines()[1].String() != "  10: line 10" {
		t.Errorf("tool buffer in %s showing %q", e.toolBuf.mode, tree.Lines()[1])
	}
	if quit := e.handleInput([]byte(":q\r")); quit || e.quitRequested {
		t.Error(": opened the command line for the tool buffer")
	}

	// q closes it.
	e.handleInput([]byte("q"))
	if e.toolBuf != nil || e.focused() != main || e.quitRequested {
		t.Errorf("tool buffer %v after q, quit %t", e.toolBuf, e.quitRequested)
	}
	e.handleInput([]byte{0x14})
	if e.toolFocused {
		t.Error("Ctrl+T focused a closed tool buffer")
	}
}

func TestEditor_ToolBufferMouse(t *testing.T) {
	e, _ := newToolEditor(t)
	main := e.Top()
	var out bytes.Buffer
	e.render(bufio.NewWriter(&out)) // Lays the buffers out.
	toolRow := e.toolBuf.y + 1      // The first row of the tool buffer, 1-based.

	click := func(x, y int) {
		sendChunks(t, e, fmt.Sprintf("\x1b[<0;%d;%dM\x1b[<0;%d;%dm", x, y, x, y))
	}
	click(6, 5) // After the line numbers.
	if e.focused() != main || main.c() != pos(4, 2) {
		t.Errorf("after a click on the main buffer: focused tool %t, cursor %s", e.toolFocused, main.c())
	}
	click(3, toolRow+2)
	if e.focused() != e.toolBuf || e.toolBuf.c() != pos(2, 2) {
		t.Errorf("after a click on the tool buffer: focused tool %t, cursor %s", e.toolFocused, e.toolBuf.c())
	}
	e.handleInput([]byte("\r"))
	runSent(e)
	if main.c() != pos(49, 2) || e.focused() != e.toolBuf {
		t.Errorf("after Enter: main cursor %s, focused tool %t", main.c(), e.toolFocused)
	}

	// The cursor is rendered in the focused buffer.
	out.Reset()
	e.render(bufio.NewWriter(&out))
	positions := cursorPositionRE.FindAllStringSubmatch(out.String(), -1)
	last := positions[len(positions)-1]
	if want := fmt.Sprintf("%d;%d", toolRow+2, 3); last[1]+";"+last[2] != want {
		t.Errorf("cursor rendered at %s, want %s", last[1:], want)
	}
}

func TestEditor_CloseBuffer(t *testing.T) {
	var e Editor
	for _, p := range []string{"a.txt", "b.txt", "c.txt"} {
		e.OpenBuffer(&Buffer{Path: p, Content: content.Empty()})
	}
	a, b := e.FindBuffer("a.txt"), e.FindBuffer("b.txt")
	if a == nil || b == nil || e.FindBuffer("d.txt") != nil || e.FindBuffer("") != nil {
		t.Fatalf("FindBuffer: a %v, b %v", a, b)
	}
	if !e.CloseBuffer(b) || e.FindBuffer("b.txt") != nil || len(e.bufs) != 2 || e.Top().Path != "c.txt" {
		t.Errorf("closing a buffer below the top: %v", e.bufs)
	}
	if !e.CloseBuffer(e.Top()) || e.Top() != a {
		t.Errorf("closing the top buffer: %v", e.bufs)
	}
	if e.CloseBuffer(a) || e.Top() != a {
		t.Errorf("closing the last buffer: %v", e.bufs)
	}
}

func TestEditor_EditBuffer(t *testing.T) {
	dir := t.TempDir()
	var e Editor
	saved := make(map[string]string)
	prev := saveContent
	saveContent = func(c content.Document, dst string) (string, error) {
		saved[dst] = content.Sprint(c, 0, c.Len())
		return "fp", nil
	}
	t.Cleanup(func() { saveContent = prev })

	ext := &editRecordingExt{}
	e.Extend(ext)
	for _, name := range []string{"a.txt", "b.txt"} {
		ft := content.FullText{content.TextLine("hello")}
		e.OpenBuffer(&Buffer{Path: dir + "/" + name, Content: &ft})
	}
	a, b := e.FindBuffer(dir+"/a.txt"), e.Top()
	for _, buf := range []*Buffer{a, b} {
		e.EditBuffer(buf, func(buf *Buffer) {
			buf.ReplaceText(content.Span{Start: pos(0, 0), End: pos(0, 5)}, "bye")
		})
	}
	// The background buffer is saved, the active one is left dirty.
	if got := saved[dir+"/a.txt"]; got != "bye\n" || a.dirty {
		t.Errorf("a.txt saved as %q, dirty %t", got, a.dirty)
	}
	if _, ok := saved[dir+"/b.txt"]; ok || !b.dirty {
		t.Errorf("b.txt saved %t, dirty %t", ok, b.dirty)
	}
	if want := []string{a.Path, b.Path}; strings.Join(ext.edited, ",") != strings.Join(want, ",") {
		t.Errorf("extensions notified about %q, want %q", ext.edited, want)
	}
	e.SaveBuffer(b)
	if got := saved[dir+"/b.txt"]; got != "bye\n" || b.dirty {
		t.Errorf("b.txt saved as %q, dirty %t", got, b.dirty)
	}
}

// editRecordingExt records the buffers it's notified about.
type editRecordingExt struct{ edited []string }

func (x *editRecordingExt) ID() string                                   { return "rec" }
func (x *editRecordingExt) MakeBufferData(Sender, *Buffer) BufferExtData { return nil }
func (x *editRecordingExt) AfterEdit(_ Sender, buf *Buffer, _ []change.Item) {
	x.edited = append(x.edited, buf.Path)
}

func TestEditor_Prompts(t *testing.T) {
	e, input := newCmdEditor(&Buffer{Content: content.Empty()})
	buf := e.Top()

	var typed []string
	e.PromptInput(buf, "rename: ", "old", func(text string) { typed = append(typed, text) })
	if !e.Prompting(buf) || e.status.cmd.prompt.prefix() != "rename: " || e.status.cmd.text != "old" {
		t.Fatalf("prompt %+v", e.status.cmd)
	}
	input(backspace, backspace, backspace, "new", "\r")
	if e.status.cmd != nil || strings.Join(typed, ",") != "new" {
		t.Errorf("typed %q, prompt %+v", typed, e.status.cmd)
	}
	e.PromptInput(buf, "rename: ", "old", func(text string) { typed = append(typed, text) })
	input("\x1b")
	if e.status.cmd != nil || strings.Join(typed, ",") != "new" {
		t.Errorf("typed %q after Esc, prompt %+v", typed, e.status.cmd)
	}

	var picked []int
	items := []string{"Extract function", "Inline call", "Extract variable"}
	e.PromptChoice(buf, "refactor: ", items, func(i int) { picked = append(picked, i) })
	input("\t", "\r")
	e.PromptChoice(buf, "refactor: ", items, func(i int) { picked = append(picked, i) })
	input("var", "\r") // Narrowed down to "Extract variable".
	e.PromptChoice(buf, "refactor: ", items, func(i int) { picked = append(picked, i) })
	input("zzz", "\r") // No matches: nothing is picked.
	e.PromptChoice(buf, "refactor: ", items, func(i int) { picked = append(picked, i) })
	input("\x1b")
	if want := "1,2"; fmt.Sprint(picked) != "["+strings.ReplaceAll(want, ",", " ")+"]" {
		t.Errorf("picked %v, want %s", picked, want)
	}
	if e.Prompting(buf) {
		t.Error("prompting after Esc")
	}
}
