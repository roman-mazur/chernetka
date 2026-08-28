package editor

import (
	"os"
	"path/filepath"
	"testing"

	"rmazur.io/chernetka/internal/content"
)

func TestCommandInput_WriteRunsSave(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")
	ft := content.FullText{content.TextLine("hello")}
	buf := &Buffer{Path: path, Content: &ft, dirty: true, mode: ModeCommand}

	prefs := RenderPrefs{}
	for _, b := range []byte{'w'} {
		commandInput(buf, []byte{b}, &prefs)
	}
	commandInput(buf, []byte{'\r'}, &prefs)

	if buf.mode != ModeNormal {
		t.Errorf("mode = %v, want ModeNormal", buf.mode)
	}
	if buf.cmdline != "" {
		t.Errorf("cmdline = %q, want empty", buf.cmdline)
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
	buf := &Buffer{mode: ModeCommand, cmdline: "wq"}
	prefs := RenderPrefs{}

	commandInput(buf, []byte{0x1b}, &prefs)

	if buf.mode != ModeNormal {
		t.Errorf("mode = %v, want ModeNormal", buf.mode)
	}
	if buf.cmdline != "" {
		t.Errorf("cmdline = %q, want empty", buf.cmdline)
	}
}

func TestCommandInput_BackspaceExitsWhenEmpty(t *testing.T) {
	buf := &Buffer{mode: ModeCommand, cmdline: "w"}
	prefs := RenderPrefs{}

	commandInput(buf, []byte{0x7f}, &prefs) // "w" -> ""
	if buf.cmdline != "" {
		t.Errorf("cmdline after 1st backspace = %q, want empty", buf.cmdline)
	}
	if buf.mode != ModeCommand {
		t.Errorf("mode after 1st backspace = %v, want ModeCommand", buf.mode)
	}

	commandInput(buf, []byte{0x7f}, &prefs) // empty -> back to normal
	if buf.mode != ModeNormal {
		t.Errorf("mode after 2nd backspace = %v, want ModeNormal", buf.mode)
	}
}

func TestCommandInput_QuitReturnsQuit(t *testing.T) {
	buf := &Buffer{mode: ModeCommand, cmdline: "q"}
	prefs := RenderPrefs{}

	if quit := commandInput(buf, []byte{'\r'}, &prefs); !quit {
		t.Errorf("commandInput(:q) returned quit=false, want true")
	}
}

func TestCommandInput_Clipboard(t *testing.T) {
	ft := content.FullText{content.TextLine("hello")}
	buf := &Buffer{Content: &ft, dirty: true, mode: ModeCommand}
	prefs := RenderPrefs{}

	StartTextSelection.DoOnBuffer(buf, prefs)
	buf.c.Col = ft.Lines()[0].Len()
	StopTextSelection.DoOnBuffer(buf, prefs)
	if selText := buf.SelectedText(); selText != "hello" {
		t.Errorf("SelectedText() = %q, want %q", selText, "hello")
	}

	buf.cmdline = "pbcopy"
	q := commandInput(buf, []byte("\r"), &prefs)
	if q {
		t.Error("quit flag returned true on pbcopy")
	}
	if res := clipboard.Read(); res != "hello" {
		t.Errorf("clipboard.Read() = %q, expected %q", res, "hello")
	}
}

func TestCommandInput_ClipboardCutAndPaste(t *testing.T) {
	ft := content.FullText{content.TextLine("hello")}
	buf := &Buffer{Content: &ft, mode: ModeCommand}
	prefs := RenderPrefs{}

	StartTextSelection.DoOnBuffer(buf, prefs)
	buf.c.Col = ft.Lines()[0].Len()
	StopTextSelection.DoOnBuffer(buf, prefs)

	buf.cmdline = "pbcut"
	if q := commandInput(buf, []byte("\r"), &prefs); q {
		t.Error("quit flag returned true on pbcut")
	}
	if res := clipboard.Read(); res != "hello" {
		t.Errorf("clipboard.Read() after cut = %q, want %q", res, "hello")
	}
	if got := ft.Lines()[0].String(); got != "" {
		t.Errorf("line after cut = %q, want empty", got)
	}

	buf.cmdline = "pbpaste"
	if q := commandInput(buf, []byte("\r"), &prefs); q {
		t.Error("quit flag returned true on pbpaste")
	}
	if got := ft.Lines()[0].String(); got != "hello" {
		t.Errorf("line after paste = %q, want %q", got, "hello")
	}
}
