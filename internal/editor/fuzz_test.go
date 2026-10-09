package editor_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/editor"
)

// fuzzFileExts are the extensions of the fuzzed file, picking the extensions of the editor engaged.
var fuzzFileExts = []string{".txt", ".go", ".md", ".d2", ".json", ".yaml", ".sh", ".cue"}

// FuzzEditor opens a file with the fuzzed text, and types the fuzzed input into the editor
// as the terminal input: keys, mouse events, and pastes. The extensions of cmd/che are
// enabled (see cheExtensions), and the fuzzed kind picks the file extension in fuzzFileExts.
// It checks that the editor neither panics nor hangs.
//
// The editor runs without its loop, see TestHarness.HandleInput: an input is handled
// the same way every time, and quickly. Files are saved only in the test directory,
// and the clipboard is kept in memory.
//
//	go test -run '^$' -fuzz FuzzEditor ./internal/editor
func FuzzEditor(f *testing.F) {
	for _, seed := range []struct {
		ext         string
		text, input string
	}{
		{".go", "", "ipackage main\r\rfunc main() {\r\tprintln(1)\r}\x1b:w\r"},
		{".txt", "line one\nline two\nline three\n", "\x1b[1;2B\x1b[1;2C\x03\x1b[B\x16\x18\x1b[1;5D"},
		{".txt", "foo bar foo\nbaz\n", "/foo\rnN/[a-z]+/X\r/(/\r\x1b/a\r"},
		{".txt", strings.Repeat("some text\n", 60), "\x1b[<0;5;3M\x1b[<0;5;3m\x1b[<0;5;3M\x1b[<32;12;6M\x1b[<0;12;6m\x1b[<65;1;1M\x1b[<64;1;1M\x1b[6~G\x1b[H"},
		{".md", "a\n", "A\x1b[200~pasted\n\ttext\x1b[201~\x1bxx0$o# Heading\r- [ ] item\r```d2\ra -> b\r```\r\x1b"},
		{".txt", "x\n", ":2\r:w copy.txt\r\x0fnew\ritext\x1b:w\r:q\r"},
		{".go", "package p\n\n// f does\nfunc f() {\n\t\tx := 1\n}\n", "jjli\t\x7f\x7f\x1b[3~\x1b[1;3C\x1b[1;10A\x1b[Z\x1b\x1b[1;7D\x1b[1;7C"},
		{".d2", "a -> b\n", "Go\rc: {\r  shape: circle\r}\x1b"},
		{".cue", "", "iname: \"value\"\rlist: [1, 2]\x1b"},
	} {
		f.Add(byte(fuzzExtIndex(seed.ext)), seed.text, []byte(seed.input))
	}
	f.Fuzz(func(t *testing.T, kind byte, text string, input []byte) {
		runFuzzedInput(t, t.TempDir(), fuzzFileExts[int(kind)%len(fuzzFileExts)], text, input)
	})
}

func fuzzExtIndex(ext string) int {
	for i, e := range fuzzFileExts {
		if e == ext {
			return i
		}
	}
	panic("unknown extension " + ext)
}

func TestRunFuzzedInput(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "sandbox")
	if err := os.Mkdir(dir, 0o700); err != nil {
		t.Fatal(err)
	}
	runFuzzedInput(t, dir, ".txt", "text\n", []byte("ihello \x1b:w\r:w ../outside.txt\r:w inside.txt\r"))

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

// runFuzzedInput opens a file with the extension and the text in dir, and handles the input.
func runFuzzedInput(t *testing.T, dir, ext, text string, input []byte) {
	t.Chdir(dir)
	editor.IsolateEffects(t, dir)

	// A hung editor never returns: crash with the goroutines to report the input.
	hang := time.AfterFunc(fuzzHangTimeout, func() {
		debug.SetTraceback("all")
		panic(fmt.Sprintf("the editor does not handle the input for %s", fuzzHangTimeout))
	})
	defer hang.Stop()

	h := editor.NewTestHarness()
	h.Root = dir
	for _, ext := range cheExtensions() {
		h.Extend(ext)
	}

	path := filepath.Join(dir, "file"+ext)
	if err := os.WriteFile(path, []byte(text), 0o600); err != nil {
		t.Fatal(err)
	}
	// Unlike OpenFile, the file is not watched: starting a watcher takes most of the time.
	if err := h.OpenReader(path, strings.NewReader(text)); err != nil {
		t.Fatal(err)
	}
	h.HandleInput(t, input)
}
