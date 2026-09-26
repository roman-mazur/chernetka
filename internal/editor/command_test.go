package editor

import (
	"os"
	"path/filepath"
	"testing"

	"rmazur.io/chernetka/internal/content"
)

// newCmdEditor returns an editor showing buf, and the function to type the keys.
// A key longer than one character that is not an escape sequence is typed rune by rune.
func newCmdEditor(buf *Buffer) (*Editor, func(keys ...string) (quit bool)) {
	var e Editor
	e.OpenBuffer(buf)
	return &e, func(keys ...string) (quit bool) {
		for _, k := range keys {
			if len(k) > 1 && k[0] != '\x1b' {
				for _, r := range k {
					quit = e.handleInput([]byte(string(r)))
				}
				continue
			}
			quit = e.handleInput([]byte(k))
		}
		return quit
	}
}

func TestCommandInput_WriteRunsSave(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")
	ft := content.FullText{content.TextLine("hello")}
	buf := &Buffer{Path: path, Content: &ft, dirty: true}
	e, input := newCmdEditor(buf)

	input(":w", "\r")

	if e.status.cmd != nil || buf.mode != ModeNormal {
		t.Errorf("command line %+v, mode %s", e.status.cmd, buf.mode)
	}
	if buf.dirty {
		t.Errorf("dirty = true, want false after :w")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if got, want := string(data), "hello"; got != want {
		t.Errorf("file content = %q, want %q", got, want)
	}
}

func TestCommandInput_EscCancels(t *testing.T) {
	e, input := newCmdEditor(&Buffer{Content: content.Empty()})
	if quit := input(":wq", "\x1b"); quit {
		t.Error("Esc quits")
	}
	if e.status.cmd != nil {
		t.Errorf("command line %+v", e.status.cmd)
	}
}

func TestCommandInput_BackspaceExitsWhenEmpty(t *testing.T) {
	e, input := newCmdEditor(&Buffer{Content: content.Empty()})

	input(":w", "\x7f") // "w" -> ""
	if e.status.cmd == nil || e.status.cmd.text != "" {
		t.Fatalf("command line after 1st backspace = %+v, want open and empty", e.status.cmd)
	}

	input("\x7f") // empty -> closed
	if e.status.cmd != nil {
		t.Errorf("command line after 2nd backspace = %+v, want closed", e.status.cmd)
	}
}

func TestCommandInput_UTF8(t *testing.T) {
	e, input := newCmdEditor(&Buffer{Content: content.Empty()})
	input(":w кіт", "\x7f")
	if got := e.status.cmd.text; got != "w кі" {
		t.Errorf("text = %q", got)
	}
}

func TestCommandInput_ResumesInsertMode(t *testing.T) {
	buf := &Buffer{Content: content.Empty()}
	e, input := newCmdEditor(buf)
	input("i", "\x0f") // Ctrl+O
	if buf.mode != ModeNormal || e.status.cmd == nil {
		t.Fatalf("mode %s, command line %+v", buf.mode, e.status.cmd)
	}
	input("\x1b")
	if buf.mode != ModeInsert {
		t.Errorf("mode %s after the command line is closed", buf.mode)
	}
}

func TestCommandInput_QuitReturnsQuit(t *testing.T) {
	_, input := newCmdEditor(&Buffer{Content: content.Empty()})
	if quit := input(":q", "\r"); !quit {
		t.Errorf(":q returned quit=false, want true")
	}
}

func TestCommandInput_Clipboard(t *testing.T) {
	ft := content.FullText{content.TextLine("hello")}
	buf := &Buffer{Content: &ft, dirty: true}
	prefs := RenderPrefs{}
	_, input := newCmdEditor(buf)

	StartTextSelection.DoOnBuffer(buf, prefs)
	buf.c.Col = ft.Lines()[0].Len()
	StopTextSelection.DoOnBuffer(buf, prefs)
	if selText := buf.SelectedText(); selText != "hello" {
		t.Errorf("SelectedText() = %q, want %q", selText, "hello")
	}

	if q := input(":pbcopy", "\r"); q {
		t.Error("quit flag returned true on pbcopy")
	}
	if res := clipboard.Read(); res != "hello" {
		t.Errorf("clipboard.Read() = %q, expected %q", res, "hello")
	}
}

func TestCommandInput_ClipboardCutAndPaste(t *testing.T) {
	ft := content.FullText{content.TextLine("hello")}
	buf := &Buffer{Content: &ft}
	prefs := RenderPrefs{}
	_, input := newCmdEditor(buf)

	StartTextSelection.DoOnBuffer(buf, prefs)
	buf.c.Col = ft.Lines()[0].Len()
	StopTextSelection.DoOnBuffer(buf, prefs)

	if q := input(":pbcut", "\r"); q {
		t.Error("quit flag returned true on pbcut")
	}
	if res := clipboard.Read(); res != "hello" {
		t.Errorf("clipboard.Read() after cut = %q, want %q", res, "hello")
	}
	if got := ft.Lines()[0].String(); got != "" {
		t.Errorf("line after cut = %q, want empty", got)
	}

	if q := input(":pbpaste", "\r"); q {
		t.Error("quit flag returned true on pbpaste")
	}
	if got := ft.Lines()[0].String(); got != "hello" {
		t.Errorf("line after paste = %q, want %q", got, "hello")
	}
}
