//go:build darwin || linux

package main

import (
	"bytes"
	"io/fs"
	"math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"rmazur.io/chernetka/internal/cheimg"
	"rmazur.io/chernetka/internal/vt/emulation"
)

const (
	monkeyDuration = 20 * time.Second
	// monkeyHangTimeout is how long the editor may not respond before it's considered hung.
	// Saving a file waits for the language server to format it.
	monkeyHangTimeout = 10 * time.Second
)

// TestMonkey types random input into che, checking that it neither crashes nor hangs.
// It builds che and che-img, and runs them in pseudo-terminals, isolated from the user's
// home directory and terminal.
// Set CHEMONKEY=1 to run it, and CHEMONKEY_SEED to repeat the same input
// (the timing of the editor still differs from run to run).
// Pass -count=1 to go test to avoid reusing a cached result:
//
//	CHEMONKEY=1 go test -count=1 -run TestMonkey ./cmd/che
func TestMonkey(t *testing.T) {
	if os.Getenv("CHEMONKEY") != "1" {
		t.Skip("set CHEMONKEY=1 to run the monkey test")
	}
	seed := uint64(time.Now().UnixNano())
	if s := os.Getenv("CHEMONKEY_SEED"); s != "" {
		var err error
		if seed, err = strconv.ParseUint(s, 10, 64); err != nil {
			t.Fatalf("bad CHEMONKEY_SEED: %s", err)
		}
	}
	t.Logf("CHEMONKEY_SEED=%d", seed)
	rnd := rand.New(rand.NewPCG(seed, 0))

	projectRoot, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	// che looks for che-img next to its executable.
	bin := t.TempDir()
	build := exec.Command("go", "build", "-o", bin+string(filepath.Separator),
		"rmazur.io/chernetka/cmd/che", "rmazur.io/chernetka/cmd/che-img")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build: %s\n%s", err, out)
	}
	env := monkeyEnv(t)

	t.Run("new file", func(t *testing.T) {
		exts := []string{".go", ".md", ".d2", ".txt", ".json", ".yaml", ".sh", ".cue"}
		path := filepath.Join(t.TempDir(), "new"+exts[rnd.IntN(len(exts))])
		if err := os.WriteFile(path, nil, 0o600); err != nil {
			t.Fatal(err)
		}
		runMonkey(t, rnd, bin, env, path)
	})

	t.Run("project file copy", func(t *testing.T) {
		src := randomProjectFile(t, rnd, projectRoot)
		t.Logf("editing a copy of %s", src)
		data, err := os.ReadFile(src)
		if err != nil {
			t.Fatal(err)
		}
		path := filepath.Join(t.TempDir(), filepath.Base(src))
		if err := os.WriteFile(path, data, 0o600); err != nil {
			t.Fatal(err)
		}
		runMonkey(t, rnd, bin, env, path)
	})
}

// monkeyEnv returns the environment of the tested processes without HOME, which is
// set per run: the logs and the sockets of che are kept in a temporary directory.
// The Go environment of the user is kept for the language server.
// Panes are never opened in the user's terminal: osascript is replaced with a failing stub.
func monkeyEnv(t *testing.T) []string {
	t.Helper()
	goVarNames := []string{"GOPATH", "GOMODCACHE", "GOCACHE", "GOENV"}
	goEnv, err := exec.Command("go", append([]string{"env"}, goVarNames...)...).Output()
	if err != nil {
		t.Fatalf("go env: %s", err)
	}
	goVars := strings.Split(strings.TrimSpace(string(goEnv)), "\n")
	stubs := t.TempDir()
	if err := os.WriteFile(filepath.Join(stubs, "osascript"), []byte("#!/bin/sh\nexit 1\n"), 0o700); err != nil {
		t.Fatal(err)
	}

	var env []string
	for _, kv := range os.Environ() {
		switch k, _, _ := strings.Cut(kv, "="); k {
		case "HOME", "PATH", "TERM", "TERM_PROGRAM", "CHE_TAB", "CHEDEBUG", "GOPATH", "GOMODCACHE", "GOCACHE", "GOENV":
		default:
			env = append(env, kv)
		}
	}
	for i, k := range goVarNames {
		env = append(env, k+"="+goVars[i])
	}
	return append(env,
		"PATH="+stubs+string(filepath.ListSeparator)+os.Getenv("PATH"),
		"TERM=xterm-256color",
		"CHE_TAB=monkey",
	)
}

// randomProjectFile returns a random text file of the project.
func randomProjectFile(t *testing.T, rnd *rand.Rand, root string) string {
	t.Helper()
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if path != root && strings.HasPrefix(d.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if !d.Type().IsRegular() || cheimg.KindOf(path) == cheimg.KindImage {
			return nil
		}
		if info, err := d.Info(); err != nil || info.Size() == 0 || info.Size() > 100<<10 {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil || !utf8.Valid(data) || bytes.IndexByte(data, 0) >= 0 {
			return nil // Not a text file.
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Fatal("no project files found")
	}
	return files[rnd.IntN(len(files))]
}

// runMonkey starts che-img, then che with the file at path, and types random input to che.
// The files the input may create stay in the directory of the file.
func runMonkey(t *testing.T, rnd *rand.Rand, bin string, env []string, path string) {
	dir := filepath.Dir(path)
	// The sockets are in the home directory: t.TempDir makes their paths too long.
	home, err := os.MkdirTemp("", "che")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(home) })
	env = append(env, "HOME="+home)
	command := func(name string, args ...string) *exec.Cmd {
		cmd := exec.Command(filepath.Join(bin, name), args...)
		cmd.Dir, cmd.Env = dir, env
		return cmd
	}

	img := emulation.Start(t, command("che-img"), 80, 40)
	img.WaitReady(t, monkeyHangTimeout) // It listens for the diagrams before it sets up its terminal.

	m := emulation.Monkey{
		Cmd:         command("che", filepath.Base(path)),
		Duration:    monkeyDuration,
		HangTimeout: monkeyHangTimeout,
		Rand:        rnd,
		Snippets:    cheSnippets,
		Ban: emulation.Ban{
			Runes:     "qy",  // Quitting ends the test early. Copying uses the system clipboard.
			Ctrl:      "cxv", // The clipboard.
			CtrlMouse: true,  // Going to a definition opens files outside the test directory.
		},
		// Close the prompts, leave the insert mode, and quit the buffer.
		Quit:  "\x1b\x1b\x1b:q\r",
		Watch: []*emulation.Process{img},
		OnFailure: func(t testing.TB) {
			t.Logf("editor log tail:\n%s", logTail(filepath.Join(home, ".chernetka", "logs", "editor.log"), 100))
		},
	}
	m.Run(t)
	img.Quit(t, "q", monkeyHangTimeout)
}

// cheSnippets are typed along with the random input: the code in the supported languages,
// the normal mode commands, search and replace, and the command line.
var cheSnippets = []string{
	"package main\r", "import \"fmt\"\r", "func main() {\r", "}\r", "fmt.", "os.", "strings.",
	"for i := range 10 {", "if err != nil {\r\treturn err\r}", "// comment ", "/* ", "*/",
	"# Heading\r", "- [ ] item\r", "```d2\r", "```\r", "a -> b\r", "\"str", "'c'", "`raw",
	"{\"k\": [1, 2]}", "name: value\r", "echo $HOME\r", "\t\t", "\r\r\r",
	"\x1bh", "\x1bj", "\x1bk", "\x1bl", "\x1b0", "\x1b$", "\x1bg", "\x1bG", "\x1b ", "\x1bb",
	"\x1bn", "\x1bN", "\x1bi", "\x1ba", "\x1bA", "\x1bo", "\x1bx", "\x1b+", "\x1b-", "\x1b:",
	"\x1b/", "\x1b/a\r", "\x1b/[a-z]+/X\r", "\x1b/(/\r", "\x1b:w\r", "\x0fnew\r",
}

// logTail returns the last n lines of the log file.
func logTail(path string, n int) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return err.Error()
	}
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	return strings.Join(lines[max(0, len(lines)-n):], "\n")
}
