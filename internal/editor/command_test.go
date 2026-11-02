package editor

import (
	"bytes"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/vt/escape"
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

	input(":w", backspace) // "w" -> ""
	if e.status.cmd == nil || e.status.cmd.text != "" {
		t.Fatalf("command line after 1st backspace = %+v, want open and empty", e.status.cmd)
	}

	input(backspace) // empty -> closed
	if e.status.cmd != nil {
		t.Errorf("command line after 2nd backspace = %+v, want closed", e.status.cmd)
	}
}

func TestCommandInput_UTF8(t *testing.T) {
	e, input := newCmdEditor(&Buffer{Content: content.Empty()})
	input(":w кіт", backspace)
	if got := e.status.cmd.text; got != "w кі" {
		t.Errorf("text = %q", got)
	}
}

func TestCommandInput_KeepsEmojiJoiners(t *testing.T) {
	e, input := newCmdEditor(&Buffer{Content: content.Empty()})
	const mermaid = "\U0001F9DC\u200D\u2640\uFE0F"
	input(":w " + mermaid)
	if got, want := e.status.cmd.text, "w "+mermaid; got != want {
		t.Errorf("text = %q, want %q", got, want)
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
	buf.updateCursor(content.Position{Line: buf.c().Line, Col: ft.Lines()[0].Len()})
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
	buf.updateCursor(content.Position{Line: buf.c().Line, Col: ft.Lines()[0].Len()})
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

func TestCommandInput_GoToLine(t *testing.T) {
	var ft content.FullText
	for range 100 {
		ft = append(ft, content.TextLine("line"))
	}
	ft[59] = content.TextLine("\t  indented")
	buf := &Buffer{Content: &ft, h: 30}
	e, input := newCmdEditor(buf)

	input(":60", "\r")
	if want := (content.Position{Line: 59, Col: 3}); buf.c() != want {
		t.Errorf("cursor at %v, want %v", buf.c(), want)
	}
	buf.clampCursor(e.rPrefs.TabSize)
	if buf.offset != 49 {
		t.Errorf("offset = %d, want the line in the upper third of the screen", buf.offset)
	}

	for cmd, line := range map[string]int{"1": 0, "0": 0, "1000": 99} {
		input(":"+cmd, "\r")
		if buf.c().Line != line {
			t.Errorf(":%s moved to line %d, want %d", cmd, buf.c().Line, line)
		}
	}
}

func TestCommandInput_GoToLineFromInsertMode(t *testing.T) {
	ft := content.FullText{content.TextLine("a"), content.TextLine("b"), content.TextLine("c")}
	buf := &Buffer{Content: &ft}
	e, input := newCmdEditor(buf)

	input("i", "\x0c") // Ctrl+L
	if e.status.cmd == nil {
		t.Fatal("Ctrl+L did not open the command line")
	}
	input("3", "\r")
	if buf.c().Line != 2 || buf.mode != ModeInsert {
		t.Errorf("cursor line %d, mode %s", buf.c().Line, buf.mode)
	}
}

const (
	ctrlN     = "\x0e"
	backspace = "\x08"
)

// newFileTree creates the files: root/a.txt, root/sub/b.txt, and root/sub/deep/c.txt.
func newFileTree(t *testing.T) (root string) {
	t.Helper()
	root = t.TempDir()
	for _, p := range []string{"a.txt", "sub/b.txt", "sub/deep/c.txt"} {
		p = filepath.Join(root, filepath.FromSlash(p))
		if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte("text"), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

func typeKeys(e *Editor, text string) {
	for _, r := range text {
		e.handleInput([]byte(string(r)))
	}
}

func newFileEditor(root string, paths ...string) *Editor {
	e := &Editor{Root: root}
	for _, p := range paths {
		(&OpenFile{Path: p}).DoOnEditor(e)
	}
	return e
}

func TestCommandInput_NewFile(t *testing.T) {
	t.Run("relative to the current file", func(t *testing.T) {
		root := newFileTree(t)
		e := newFileEditor(root, filepath.Join(root, "sub", "b.txt"))
		e.handleInput([]byte(ctrlN))

		if !prompting[*newFilePrompt](e) {
			t.Fatalf("didn't enter prompting mode: %+v", e.status.cmd)
		}
		if e.status.cmd.text != "n sub/" {
			t.Fatalf("e.status.cmd.text = %q, want %q", e.status.cmd.text, "n sub/")
		}

		typeKeys(e, "new.txt\r")
		assertTop(t, e, filepath.Join(root, "sub", "new.txt"), ModeNormal)
		assertBufferCount(t, e, 2)
		if !e.Top().canEdit() {
			t.Error("cannot edit the new file")
		}
		if _, err := os.Stat(e.Top().Path); !os.IsNotExist(err) {
			t.Errorf("the file exists before saving: %v", err)
		}
	})

	t.Run("relative path of the current file", func(t *testing.T) {
		root := newFileTree(t)
		t.Chdir(root)
		e := newFileEditor(root, filepath.Join("sub", "b.txt"))
		e.handleInput([]byte(ctrlN))
		typeKeys(e, "../new.txt\r")
		assertTop(t, e, filepath.Join(root, "new.txt"), ModeNormal)
	})

	t.Run("buffer without a file uses the project root", func(t *testing.T) {
		root := newFileTree(t)
		t.Chdir(filepath.Join(root, "sub")) // Not the working directory.
		e := newFileEditor(root)
		e.New()
		e.handleInput([]byte(ctrlN))
		if c := e.status.cmd; c == nil || c.text != "n " {
			t.Fatalf("command line %+v", c)
		}
		typeKeys(e, "new.txt\r")
		assertTop(t, e, filepath.Join(root, "new.txt"), ModeNormal)
	})

	t.Run("absolute path", func(t *testing.T) {
		root := newFileTree(t)
		e := newFileEditor(root, filepath.Join(root, "sub", "b.txt"))
		path := filepath.Join(root, "new.txt")
		e.handleInput([]byte(ctrlN))
		typeKeys(e, strings.Repeat(backspace, 4)+path+"\r")
		assertTop(t, e, path, ModeNormal)
	})

	t.Run("typed command", func(t *testing.T) {
		root := newFileTree(t)
		e := newFileEditor(root, filepath.Join(root, "a.txt"))
		typeKeys(e, ":n n.txt\r")
		assertTop(t, e, filepath.Join(root, "n.txt"), ModeNormal)
	})

	t.Run("saving creates the parent directories", func(t *testing.T) {
		root := newFileTree(t)
		e := newFileEditor(root, filepath.Join(root, "a.txt"))
		e.handleInput([]byte(ctrlN))
		typeKeys(e, filepath.Join("x", "y", "new.txt")+"\r")
		typeKeys(e, "ihello\x1b")
		e.handleInput([]byte{0x13}) // Ctrl+S.
		data, err := os.ReadFile(filepath.Join(root, "x", "y", "new.txt"))
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "hello" {
			t.Errorf("saved %q", data)
		}
	})

	t.Run("existing file is opened", func(t *testing.T) {
		root := newFileTree(t)
		e := newFileEditor(root, filepath.Join(root, "a.txt"))
		e.handleInput([]byte(ctrlN))
		typeKeys(e, filepath.Join("sub", "b.txt")+"\r")
		assertTop(t, e, filepath.Join(root, "sub", "b.txt"), ModeNormal)
		if got := e.Top().Content.Lines()[0].String(); got != "text" {
			t.Errorf("content %q", got)
		}
	})

	t.Run("open buffer is activated", func(t *testing.T) {
		root := newFileTree(t)
		a, b := filepath.Join(root, "a.txt"), filepath.Join(root, "sub", "b.txt")
		e := newFileEditor(root, a, b)
		e.handleInput([]byte(ctrlN))
		typeKeys(e, "../a.txt\r")
		assertTop(t, e, a, ModeNormal)
		assertBufferCount(t, e, 2)
	})

	t.Run("nothing to create", func(t *testing.T) {
		root := newFileTree(t)
		a := filepath.Join(root, "a.txt")
		for _, text := range []string{"", "x/", "sub"} {
			t.Logf("checking %q", text)
			e := newFileEditor(root, a)
			e.handleInput([]byte(ctrlN))
			typeKeys(e, text+"\r")
			assertTop(t, e, a, ModeNormal)
			assertBufferCount(t, e, 1)
		}
	})

	t.Run("open hook", func(t *testing.T) {
		root := newFileTree(t)
		e := newFileEditor(root, filepath.Join(root, "a.txt"))
		var opened []string
		e.OpenPath = func(path string) { opened = append(opened, path) }
		e.handleInput([]byte(ctrlN))
		typeKeys(e, "new.txt\r")
		e.handleInput([]byte(ctrlN))
		typeKeys(e, filepath.Join("sub", "b.txt")+"\r")
		if want := []string{filepath.Join(root, "new.txt"), filepath.Join(root, "sub", "b.txt")}; !slices.Equal(opened, want) {
			t.Errorf("opened %q, want %q", opened, want)
		}
		assertBufferCount(t, e, 1)
	})

	t.Run("backspace returns to the ex command", func(t *testing.T) {
		root := newFileTree(t)
		e := newFileEditor(root, filepath.Join(root, "a.txt"))
		e.handleInput([]byte(ctrlN))
		e.handleInput([]byte(backspace))
		if c := e.status.cmd; c == nil || c.text != "n" || prompting[*newFilePrompt](e) {
			t.Fatalf("command line %+v", c)
		}
		e.handleInput([]byte(" "))
		if !prompting[*newFilePrompt](e) {
			t.Fatalf("command line %+v", e.status.cmd)
		}
	})

	t.Run("no path does not run other commands", func(t *testing.T) {
		root := newFileTree(t)
		path := filepath.Join(root, "a.txt")
		e := newFileEditor(root, path)
		typeKeys(e, "ix\x1b")
		typeKeys(e, ":n\r")
		if !e.Top().dirty {
			t.Error("the buffer is saved")
		}
		assertBufferCount(t, e, 1)
	})

	t.Run("insert mode is restored on cancel", func(t *testing.T) {
		root := newFileTree(t)
		a := filepath.Join(root, "a.txt")
		e := newFileEditor(root, a)
		typeKeys(e, "i")
		e.handleInput([]byte(ctrlN))
		e.handleInput([]byte(ctrlN)) // Keeps the prompt open.
		typeKeys(e, "x")
		if c := e.status.cmd; c == nil || c.text != "n x" {
			t.Fatalf("command line %+v", c)
		}
		e.handleInput([]byte{0x1b})
		assertTop(t, e, a, ModeInsert)
	})
}

func TestCommandInput_NewFile_DirBuffer(t *testing.T) {
	sep := string(filepath.Separator)
	open := func(t *testing.T) (*Editor, string) {
		root := newFileTree(t)
		e := &Editor{Root: root}
		e.OpenDir(root, nil)
		return e, root
	}
	engage := func(e *Editor, line int) {
		e.Top().updateCursor(content.Position{Line: line})
		e.handleInput([]byte{'\r'})
	}

	t.Run("cursor on a directory", func(t *testing.T) {
		e, root := open(t)
		e.handleInput([]byte(ctrlN)) // The lines are: sub, a.txt.
		if got, want := e.status.cmd.text, "n sub"+sep; got != want {
			t.Fatalf("text = %q, want %q", got, want)
		}
		typeKeys(e, "new.txt\r")
		assertTop(t, e, filepath.Join(root, "sub", "new.txt"), ModeNormal)
	})

	t.Run("cursor on a file in the root", func(t *testing.T) {
		e, root := open(t)
		e.Top().updateCursor(content.Position{Line: 1})
		e.handleInput([]byte(ctrlN))
		if got := e.status.cmd.text; got != "n " {
			t.Fatalf("text = %q", got)
		}
		typeKeys(e, "new.txt\r")
		assertTop(t, e, filepath.Join(root, "new.txt"), ModeNormal)
	})

	t.Run("nested directories", func(t *testing.T) {
		e, _ := open(t)
		engage(e, 0) // sub: deep, b.txt.
		engage(e, 1) // deep: c.txt.

		e.Top().updateCursor(content.Position{Line: 1})
		e.handleInput([]byte(ctrlN))
		if got, want := e.status.cmd.text, "n "+filepath.Join("sub", "deep")+sep; got != want {
			t.Errorf("on the directory: text = %q, want %q", got, want)
		}
		e.handleInput([]byte{0x1b})

		e.Top().updateCursor(content.Position{Line: 2})
		e.handleInput([]byte(ctrlN))
		if got, want := e.status.cmd.text, "n "+filepath.Join("sub", "deep")+sep; got != want {
			t.Errorf("on the file: text = %q, want %q", got, want)
		}
		e.handleInput([]byte{0x1b})

		e.Top().updateCursor(content.Position{Line: 3})
		e.handleInput([]byte(ctrlN))
		if got, want := e.status.cmd.text, "n sub"+sep; got != want {
			t.Errorf("on the file in the parent: text = %q, want %q", got, want)
		}
	})

	t.Run("typed command", func(t *testing.T) {
		e, _ := open(t)
		typeKeys(e, ":n ")
		if got, want := e.status.cmd.text, "n sub"+sep; got != want {
			t.Fatalf("text = %q, want %q", got, want)
		}
	})

	t.Run("without a name nothing is created", func(t *testing.T) {
		e, _ := open(t)
		e.handleInput([]byte(ctrlN))
		e.handleInput([]byte{'\r'})
		assertBufferCount(t, e, 1)
	})

	t.Run("relative tree root", func(t *testing.T) {
		root := newFileTree(t)
		t.Chdir(filepath.Dir(root))
		e := &Editor{Root: root}
		e.OpenDir(filepath.Base(root), nil)
		var opened []string
		e.OpenPath = func(path string) { opened = append(opened, path) }
		e.handleInput([]byte(ctrlN))
		typeKeys(e, "new.txt\r")
		// The path is passed to a process that may run in another directory.
		if want := []string{filepath.Join(root, "sub", "new.txt")}; !slices.Equal(opened, want) {
			t.Errorf("opened %q, want %q", opened, want)
		}
	})

	t.Run("working directory inside the project", func(t *testing.T) {
		e, root := open(t) // The lines are: sub, a.txt.
		t.Chdir(filepath.Join(root, "sub"))
		e.handleInput([]byte(ctrlN))
		if got, want := e.status.cmd.text, "n sub/"; got != want {
			t.Errorf("on the directory: text = %q, want %q", got, want)
		}
		e.handleInput([]byte{0x1b})

		e.Top().updateCursor(content.Position{Line: 1})
		e.handleInput([]byte(ctrlN))
		if got, want := e.status.cmd.text, "n "; got != want {
			t.Errorf("on the root file: text = %q, want %q", got, want)
		}
	})

	t.Run("tree below the project root", func(t *testing.T) {
		root := newFileTree(t)
		e := &Editor{Root: root}
		e.OpenDir(filepath.Join(root, "sub"), nil) // The lines are: deep, b.txt.
		var opened []string
		e.OpenPath = func(path string) { opened = append(opened, path) }
		e.handleInput([]byte(ctrlN))
		if got, want := e.status.cmd.text, "n sub/deep/"; got != want {
			t.Errorf("text = %q, want %q", got, want)
		}
		typeKeys(e, "new.txt\r")
		if want := []string{filepath.Join(root, "sub", "deep", "new.txt")}; !slices.Equal(opened, want) {
			t.Errorf("opened %q, want %q", opened, want)
		}
	})
}

func TestOpenFile_Errors(t *testing.T) {
	t.Run("missing file starts an empty buffer", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "x", "new.txt")
		e := &Editor{}
		(&OpenFile{Path: path}).DoOnEditor(e)
		assertTop(t, e, path, ModeNormal)
		if !e.Top().canEdit() || e.Top().dirty {
			t.Errorf("canEdit %t, dirty %t", e.Top().canEdit(), e.Top().dirty)
		}
	})

	t.Run("other errors are shown", func(t *testing.T) {
		file := filepath.Join(t.TempDir(), "a.txt")
		if err := os.WriteFile(file, nil, 0o600); err != nil {
			t.Fatal(err)
		}
		e := &Editor{}
		(&OpenFile{Path: filepath.Join(file, "b.txt")}).DoOnEditor(e) // The parent is a file.
		if _, ok := e.Top().Content.(*content.ErrorContent); !ok {
			t.Errorf("content %#v", e.Top().Content)
		}
	})
}

func TestNewFilePrompt_Render(t *testing.T) {
	root := newFileTree(t)
	e := newFileEditor(root, filepath.Join(root, "sub", "b.txt"))
	e.handleInput([]byte(ctrlN))

	render := func(w int) string {
		var out bytes.Buffer
		e.status.cmd.prompt.render(&out, w)
		return strings.TrimRight(escape.Clean(out.String()), " ")
	}
	typeKeys(e, "n.txt")

	want := ":n sub/n.txt"
	if got := render(20); got != want {
		t.Errorf("rendered %q, want %q", got, want)
	}

	for range 5 {
		e.handleInput([]byte(backspace))
	}
	typeKeys(e, "b.txt")
	if got := render(50); !strings.HasSuffix(got, "(exists)") {
		t.Errorf("rendered %q", got)
	}
	if got := render(8); strings.HasSuffix(got, "(exists)") {
		t.Errorf("rendered %q", got)
	}

	for range 5 {
		e.handleInput([]byte(backspace))
	}
	typeKeys(e, "deep")
	if got := render(50); !strings.HasSuffix(got, "(dir)") {
		t.Errorf("rendered %q", got)
	}
}
