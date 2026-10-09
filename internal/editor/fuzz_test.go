package editor

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/content"
)

// FuzzEditor opens a file with the fuzzed text, and types the fuzzed input into the editor
// as the terminal input: keys, mouse events, and pastes. It checks that the editor
// neither panics nor hangs.
//
// The editor runs without its loop: the input is read and the commands are run in the test
// goroutine, so an input is handled the same way every time, and quickly.
// Files are saved only in the test directory, and the clipboard is kept in memory.
//
//	go test -run '^$' -fuzz FuzzEditor ./internal/editor
func FuzzEditor(f *testing.F) {
	for _, seed := range []struct{ text, input string }{
		{"", "ipackage main\r\rfunc main() {\r\tprintln(1)\r}\x1b:w\r"},
		{"line one\nline two\nline three\n", "\x1b[1;2B\x1b[1;2C\x03\x1b[B\x16\x18\x1b[1;5D"},
		{"foo bar foo\nbaz\n", "/foo\rnN/[a-z]+/X\r/(/\r\x1b/a\r"},
		{strings.Repeat("some text\n", 60), "\x1b[<0;5;3M\x1b[<0;5;3m\x1b[<0;5;3M\x1b[<32;12;6M\x1b[<0;12;6m\x1b[<65;1;1M\x1b[<64;1;1M\x1b[6~G\x1b[H"},
		{"a\n", "A\x1b[200~pasted\n\ttext\x1b[201~\x1bxx0$o# Heading\r- [ ] item\r\x1b"},
		{"x\n", ":2\r:w copy.txt\r\x0fnew\ritext\x1b:w\r:q\r"},
		{"\tindented\n\t\tmore\n", "jli\t\x7f\x7f\x1b[3~\x1b[1;3C\x1b[1;10A\x1b[Z\x1b"},
	} {
		f.Add(seed.text, []byte(seed.input))
	}
	f.Fuzz(func(t *testing.T, text string, input []byte) {
		runFuzzedInput(t, t.TempDir(), text, input)
	})
}

func TestRunFuzzedInput(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "sandbox")
	if err := os.Mkdir(dir, 0o700); err != nil {
		t.Fatal(err)
	}
	runFuzzedInput(t, dir, "text\n", []byte("ihello \x1b:w\r:w ../outside.txt\r:w inside.txt\r"))

	for _, name := range []string{"file.txt", "inside.txt"} {
		data, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil || !strings.Contains(string(data), "hello text") {
			t.Errorf("%s has %q (%v), want the typed text saved", name, data, err)
		}
	}
	if _, err := os.Stat(filepath.Join(root, "outside.txt")); err == nil {
		t.Error("a file is saved outside the test directory")
	}
}

// fuzzHangTimeout is how long an input may take before the editor is considered hung.
const fuzzHangTimeout = 10 * time.Second

// runFuzzedInput opens a file with the text in dir, and handles the input.
func runFuzzedInput(t *testing.T, dir, text string, input []byte) {
	t.Chdir(dir)
	sandboxEffects(t, dir)

	// A hung editor never returns: crash with the goroutines to report the input.
	hang := time.AfterFunc(fuzzHangTimeout, func() {
		debug.SetTraceback("all")
		panic(fmt.Sprintf("the editor does not handle the input for %s", fuzzHangTimeout))
	})
	defer hang.Stop()

	// Every input byte sends at most one command, and the commands may send more:
	// keep room for them, so sending never blocks without the loop.
	e := &Editor{Root: dir, cmdChannel: make(chan Command, len(input)+64)}
	e.rPrefs = newRenderPrefs()
	defer func() {
		close(e.loopDone()) // Unblocks the goroutines sending commands, like the file watcher.
		if e.watcher != nil {
			_ = e.watcher.Close()
		}
	}()

	path := filepath.Join(dir, "file.txt")
	if err := os.WriteFile(path, []byte(text), 0o600); err != nil {
		t.Fatal(err)
	}
	// Unlike OpenFile, the file is not watched: starting a watcher takes most of the time.
	if err := e.OpenReader(path, strings.NewReader(text)); err != nil {
		t.Fatal(err)
	}
	e.Top().fileText = e.Top().Text()

	// The reader sends the commands for the whole input, and the quit at its end.
	e.readAndHandleInput(context.Background(), bufio.NewReader(bytes.NewReader(input)))

	out := bufio.NewWriter(io.Discard)
	if e.update(out) {
		return
	}
	for {
		select {
		case cmd := <-e.cmdChannel:
			cmd.DoOnEditor(e)
			if e.update(out) {
				return
			}
		default:
			return
		}
	}
}

// sandboxEffects keeps the clipboard in memory, and lets the files be saved only in dir.
func sandboxEffects(t *testing.T, dir string) {
	prevClipboard, prevSave := clipboard, saveContent
	t.Cleanup(func() { clipboard, saveContent = prevClipboard, prevSave })

	clipboard = new(memoryClipboard)
	saveContent = func(doc content.Document, dst string) error {
		abs, err := filepath.Abs(dst)
		if err != nil {
			return err
		}
		if rel, err := filepath.Rel(dir, abs); err != nil || !filepath.IsLocal(rel) {
			return errors.New("saving outside the test directory")
		}
		return content.Save(doc, dst)
	}
}

type memoryClipboard struct{ text string }

func (c *memoryClipboard) Write(s string) { c.text = s }
func (c *memoryClipboard) Read() string   { return c.text }
