package extlsp

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"go.lsp.dev/uri"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/logger"
)

// The navigation and refactoring tools with a real gopls, skipped unless
// CHE_LSP_EVAL is set (see eval_test.go).

// realToolsWorkspace is a module with a function used in two files.
func realToolsWorkspace(t *testing.T) (dir, mainPath, otherPath string) {
	t.Helper()
	dir = t.TempDir()
	if abs, err := filepath.EvalSymlinks(dir); err == nil {
		dir = abs // gopls reports the real paths.
	}
	_ = os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/x\n\ngo 1.22\n"), 0o644)
	mainPath = filepath.Join(dir, "main.go")
	_ = os.WriteFile(mainPath, []byte("package main\n\nfunc foo(x int, s string) int { return x }\n\nfunc main() { foo(1, \"a\") }\n"), 0o644)
	otherPath = filepath.Join(dir, "other.go")
	_ = os.WriteFile(otherPath, []byte("package main\n\nfunc other() int {\n\treturn foo(2, \"b\")\n}\n"), 0o644)
	return dir, mainPath, otherPath
}

func openRealGopls(t *testing.T, le *Integration, path string) (*editor.TestHarness, *BufferData) {
	t.Helper()
	src, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	h := editor.NewTestHarness(t)
	h.LogEmbed = logger.Embed(logger.Prefix(t.Logf, "editor: "))
	h.Extend(le)
	if err := h.OpenReader(path, strings.NewReader(string(src))); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = le.Close() })
	data := h.Top().ExtensionData(le.ID()).(*BufferData)
	for !data.srv.ready.Load() {
		if data.srv.failed.Load() {
			t.Fatal("gopls failed to start")
		}
		time.Sleep(20 * time.Millisecond)
	}
	return h, data
}

func TestRealGoplsFindsUsagesAndCallers(t *testing.T) {
	skipUnlessEval(t)
	_, mainPath, _ := realToolsWorkspace(t)
	var le Integration
	h, data := openRealGopls(t, &le, mainPath)
	h.Run(t)

	started := time.Now()
	h.Post(t, editor.CommandFunc(func(e *editor.Editor) { data.FindUsages(e, content.Position{Line: 2, Col: 6}) }))
	got := waitTool(t, h)
	fmt.Printf("usages in %s:\n%s\n", time.Since(started).Round(time.Millisecond), strings.Join(got, "\n"))
	want := []string{
		"  3 usages of foo",
		"~ main.go",
		"    3:6  func foo(x int, s string) int { return x }",
		"    5:15  func main() { foo(1, \"a\") }",
		"~ other.go",
		"    4:9  return foo(2, \"b\")",
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("usages:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}

	h.Post(t, editor.CommandFunc(func(e *editor.Editor) { data.FindCallers(e, content.Position{Line: 2, Col: 6}) }))
	for range 400 {
		if got = waitTool(t, h); strings.HasSuffix(got[0], "callers of foo") {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Printf("callers:\n%s\n", strings.Join(got, "\n"))
	want = []string{
		"  2 callers of foo",
		"~ main  main.go:5",
		"    5:15  func main() { foo(1, \"a\") }",
		"  > callers of main",
		"~ other  other.go:3",
		"    4:9  return foo(2, \"b\")",
		"  > callers of other",
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("callers:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}

func TestRealGoplsRenames(t *testing.T) {
	skipUnlessEval(t)
	_, mainPath, otherPath := realToolsWorkspace(t)
	var le Integration
	h, data := openRealGopls(t, &le, mainPath)
	buf := h.Top()
	h.Run(t)

	rename := func(pos content.Position, wantPlaceholder, typed string) {
		t.Helper()
		h.Post(t, editor.CommandFunc(func(e *editor.Editor) { data.Rename(e, pos) }))
		deadline := time.Now().Add(30 * time.Second)
		for !onLoop(t, h, func() bool { return h.Prompting(buf) }) {
			if time.Now().After(deadline) {
				t.Fatal("no rename prompt")
			}
			time.Sleep(10 * time.Millisecond)
		}
		if got := onLoop(t, h, h.PromptText); got != wantPlaceholder {
			t.Errorf("rename placeholder %q, want %q", got, wantPlaceholder)
		}
		h.SendInput(t, []byte(strings.Repeat("\x7f", len(wantPlaceholder))+typed+"\r"))
	}
	waitText := func(want string) {
		t.Helper()
		deadline := time.Now().Add(30 * time.Second)
		for onLoop(t, h, buf.Text) != want {
			if time.Now().After(deadline) {
				t.Fatalf("buffer text %q, want %q", onLoop(t, h, buf.Text), want)
			}
			time.Sleep(10 * time.Millisecond)
		}
	}

	// A symbol: the other file is changed on the disk.
	rename(content.Position{Line: 2, Col: 6}, "foo", "bar")
	waitText("package main\n\nfunc bar(x int, s string) int { return x }\n\nfunc main() { bar(1, \"a\") }\n")
	if data, _ := os.ReadFile(otherPath); string(data) != "package main\n\nfunc other() int {\n\treturn bar(2, \"b\")\n}\n" {
		t.Errorf("other.go: %q", data)
	}

	// The signature of a function, on its func keyword: the parameters are reordered.
	rename(content.Position{Line: 2, Col: 1}, "func(x int, s string) int", "func(s string, x int) int")
	waitText("package main\n\nfunc bar(s string, x int) int { return x }\n\nfunc main() { bar(\"a\", 1) }\n")
	if data, _ := os.ReadFile(otherPath); !strings.Contains(string(data), "bar(\"b\", 2)") {
		t.Errorf("other.go: %q", data)
	}
}

func TestRealGoplsCodeActions(t *testing.T) {
	skipUnlessEval(t)
	_, mainPath, _ := realToolsWorkspace(t)
	var le Integration
	h, data := openRealGopls(t, &le, mainPath)
	buf := h.Top()
	h.Run(t)

	// The call of foo in main is inlined.
	h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
		data.ShowCodeActions(e, content.Span{Start: content.Position{Line: 4, Col: 15}, End: content.Position{Line: 4, Col: 15}})
	}))
	deadline := time.Now().Add(30 * time.Second)
	for !onLoop(t, h, func() bool { return h.Prompting(buf) }) {
		if time.Now().After(deadline) {
			t.Fatal("no code actions prompt")
		}
		time.Sleep(10 * time.Millisecond)
	}
	items := onLoop(t, h, h.PromptItems)
	fmt.Printf("code actions: %q\n", items)
	h.SendInput(t, []byte("inline\r"))
	want := "package main\n\nfunc foo(x int, s string) int { return x }\n\nfunc main() { _ = 1 }\n"
	deadline = time.Now().Add(30 * time.Second)
	for onLoop(t, h, buf.Text) != want {
		if time.Now().After(deadline) {
			t.Fatalf("buffer text %q, want %q", onLoop(t, h, buf.Text), want)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// TestRealGoplsSeesFileChanges checks that a file changed on the disk, like by
// another editor, is seen by gopls: the problem of using what it defines goes away.
func TestRealGoplsSeesFileChanges(t *testing.T) {
	skipUnlessEval(t)
	dir, mainPath, otherPath := realToolsWorkspace(t)
	src := "package main\n\nfunc foo(x int, s string) int { return x }\n\nfunc main() { foo(1, \"a\"); baz() }\n"
	_ = os.WriteFile(mainPath, []byte(src), 0o644)
	var le Integration
	h, data := openRealGopls(t, &le, mainPath)
	h.Run(t)

	waitDiags := func(want int) {
		t.Helper()
		deadline := time.Now().Add(30 * time.Second)
		for {
			diags := onLoop(t, h, data.Diagnostics)
			if len(diags) == want {
				return
			}
			if time.Now().After(deadline) {
				t.Fatalf("diagnostics %+v, want %d", diags, want)
			}
			time.Sleep(20 * time.Millisecond)
		}
	}
	waitDiags(1) // baz is undefined.

	// Watching the files takes a moment to start after gopls asks for it.
	deadline := time.Now().Add(10 * time.Second)
	for {
		data.srv.watcher.mu.Lock()
		n := len(data.srv.watcher.stops)
		data.srv.watcher.mu.Unlock()
		if n > 0 || time.Now().After(deadline) {
			break
		}
		time.Sleep(20 * time.Millisecond)
	}
	time.Sleep(200 * time.Millisecond)
	_ = os.WriteFile(filepath.Join(dir, "baz.go"), []byte("package main\n\nfunc baz() {}\n"), 0o644)
	started := time.Now()
	waitDiags(0)
	fmt.Printf("the new file is seen in %s\n", time.Since(started).Round(time.Millisecond))

	// A closed buffer is read from the disk again: other.go is opened, closed, and changed.
	otherSrc, _ := os.ReadFile(otherPath)
	if err := h.OpenReader(otherPath, strings.NewReader(string(otherSrc))); err != nil {
		t.Fatal(err)
	}
	other := h.Top()
	h.Post(t, editor.CommandFunc(func(e *editor.Editor) { e.CloseBuffer(other) }))
	_ = os.WriteFile(otherPath, []byte("package main\n\nfunc other() int {\n\treturn foo(2) + undefinedThing\n}\n"), 0o644)
	// The problem is in other.go, which is not open: gopls publishes it for the package anyway.
	deadline = time.Now().Add(30 * time.Second)
	for len(data.srv.diagnostics(uri.File(otherPath))) == 0 {
		if time.Now().After(deadline) {
			t.Fatal("the change of the closed file is not seen")
		}
		time.Sleep(20 * time.Millisecond)
	}
}
