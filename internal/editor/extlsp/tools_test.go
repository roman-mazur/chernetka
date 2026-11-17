package extlsp

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/lsp"
)

// The navigation and refactoring requests of the fake server.

type fakeTools struct {
	references   []protocol.Location
	callers      map[string][]protocol.CallHierarchyIncomingCall // by function name
	renameTarget *lsp.RenameTarget
	renameEdit   *lsp.WorkspaceEdit
	renameErr    error
	actions      []lsp.CodeAction
	resolveEdit  *lsp.WorkspaceEdit // the edit of the resolved action
	applyEdit    func(lsp.ApplyEditRequest) error

	closed   []uri.URI
	renamed  []string // the new names requested
	executed []string // the commands executed
	watched  []protocol.FileEvent
	events   chan struct{} // signaled on the notifications
}

func (f *fakeLSP) signal() {
	if f.events != nil {
		select {
		case f.events <- struct{}{}:
		default:
		}
	}
}

func (f *fakeLSP) DidClose(_ context.Context, u uri.URI) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.closed = append(f.closed, u)
	f.signal()
	return nil
}

func (f *fakeLSP) DidChangeWatchedFiles(_ context.Context, changes []protocol.FileEvent) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.watched = append(f.watched, changes...)
	f.signal()
	return nil
}

func (f *fakeLSP) References(context.Context, uri.URI, protocol.Position) ([]protocol.Location, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.references, nil
}

func (f *fakeLSP) PrepareCallHierarchy(_ context.Context, u uri.URI, pos protocol.Position) ([]protocol.CallHierarchyItem, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.callers == nil {
		return nil, nil
	}
	return []protocol.CallHierarchyItem{{Name: "root", URI: u, SelectionRange: protocol.Range{Start: pos, End: pos}}}, nil
}

func (f *fakeLSP) IncomingCalls(_ context.Context, item protocol.CallHierarchyItem) ([]protocol.CallHierarchyIncomingCall, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.callers[item.Name], nil
}

func (f *fakeLSP) PrepareRename(context.Context, uri.URI, protocol.Position, string) (*lsp.RenameTarget, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.renameTarget, nil
}

func (f *fakeLSP) Rename(_ context.Context, _ uri.URI, _ protocol.Position, newName string) (*lsp.WorkspaceEdit, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.renamed = append(f.renamed, newName)
	return f.renameEdit, f.renameErr
}

func (f *fakeLSP) CodeActions(context.Context, uri.URI, protocol.Range, []protocol.Diagnostic) ([]lsp.CodeAction, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.actions, nil
}

func (f *fakeLSP) ResolveCodeAction(_ context.Context, action lsp.CodeAction) (lsp.CodeAction, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	action.Edit, action.Data = f.resolveEdit, nil
	return action, nil
}

func (f *fakeLSP) ExecuteCommand(_ context.Context, cmd protocol.Command) error {
	f.mu.Lock()
	apply := f.applyEdit
	f.executed = append(f.executed, cmd.Command)
	f.mu.Unlock()
	if apply == nil {
		return errors.New("unknown command")
	}
	// The server makes the edits through the client.
	return apply(lsp.ApplyEditRequest{Label: cmd.Title, Edit: lsp.WorkspaceEdit{Changes: []lsp.DocumentChange{
		{URI: uri.File(cmd.Arguments[0].(string)), Edits: []protocol.TextEdit{{NewText: "// by the command\n"}}},
	}}})
}

// waitEvent waits for a notification of the fake server.
func (f *fakeLSP) waitEvent(t *testing.T) {
	t.Helper()
	select {
	case <-f.events:
	case <-time.After(2 * time.Second):
		t.Fatal("no notification")
	}
}

// waitTool waits for the tool buffer to be shown and returns its lines.
func waitTool(t *testing.T, h *editor.TestHarness) []string {
	t.Helper()
	for range 400 {
		var lines []string
		h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
			if tb := h.ToolBuffer(); tb != nil {
				for _, l := range tb.Content.Lines() {
					lines = append(lines, l.String())
				}
			}
		}))
		if lines != nil {
			return lines
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("no tool buffer")
	return nil
}

func loc(path string, line, char uint32) protocol.Location {
	p := protocol.Position{Line: line, Character: char}
	return protocol.Location{URI: uri.File(path), Range: protocol.Range{Start: p, End: p}}
}

func TestIntegration_FindUsages(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	other := filepath.Join(dir, "other.go")
	if err := os.WriteFile(other, []byte("package main\n\nfunc other() {\n\tfoo(2) // 日本\n}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	main := filepath.Join(dir, "main.go")
	fake := &fakeLSP{}
	fake.references = []protocol.Location{loc(other, 3, 1), loc(main, 2, 5), loc(main, 4, 8), loc(main, 4, 8)}
	var le Integration
	h, _, data := newBuffer(t, &le, fake, main, "package main\n\nfunc foo(x int) {}\n\nvar _ = foo(1)\n")
	h.Run(t)

	h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
		data.FindUsages(e, content.Position{Line: 2, Col: 6})
	}))
	got := waitTool(t, h)
	want := []string{
		"  4 usages of foo",
		"~ main.go",
		"    3:6  func foo(x int) {}",
		"    5:9  var _ = foo(1)",
		"~ other.go",
		"    4:2  foo(2) // 日本",
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("tool buffer:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}

	// Enter on a usage jumps to it, in the other file too.
	h.SendInput(t, []byte("jjjjj\r"))
	h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
		if top := e.Top(); top.Path != other || top.Mode() != editor.ModeReview {
			t.Errorf("top buffer %s", top.Path)
		} else if cx, cy := top.Pos(); cx != 1 || cy != 3 {
			t.Errorf("cursor at %d:%d", cy, cx)
		}
	}))

	// No usages.
	fake.mu.Lock()
	fake.references = nil
	fake.mu.Unlock()
	h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
		e.FindBuffer(main) // The data belongs to main.go, which is not the top buffer now.
		data.FindUsages(e, content.Position{Line: 4, Col: 8})
	}))
	for range 400 {
		if got = waitTool(t, h); len(got) == 1 {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	if want := "  no usages of foo"; len(got) != 1 || got[0] != want {
		t.Errorf("tool buffer %q, want %q", got, want)
	}
}

func TestIntegration_FindCallers(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	main := filepath.Join(dir, "main.go")
	text := "package main\n\nfunc foo() {}\n\nfunc bar() { foo() }\n\nfunc baz() { bar(); bar() }\n"
	if err := os.WriteFile(main, []byte(text), 0o644); err != nil {
		t.Fatal(err)
	}
	fn := func(name string, line uint32) protocol.CallHierarchyItem {
		p := protocol.Position{Line: line, Character: 5}
		return protocol.CallHierarchyItem{Name: name, URI: uri.File(main), SelectionRange: protocol.Range{Start: p, End: p}}
	}
	rng := func(line, char uint32) protocol.Range {
		p := protocol.Position{Line: line, Character: char}
		return protocol.Range{Start: p, End: p}
	}
	fake := &fakeLSP{}
	fake.callers = map[string][]protocol.CallHierarchyIncomingCall{
		"root": {{From: fn("bar", 4), FromRanges: []protocol.Range{rng(4, 13)}}},
		"bar":  {{From: fn("baz", 6), FromRanges: []protocol.Range{rng(6, 13), rng(6, 20)}}},
	}
	var le Integration
	h, _, data := newBuffer(t, &le, fake, main, text)
	h.Run(t)

	h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
		data.FindCallers(e, content.Position{Line: 2, Col: 6})
	}))
	got := waitTool(t, h)
	want := []string{
		"  1 callers of root",
		"~ bar  main.go:5",
		"    5:14  func bar() { foo() }",
		"  > callers of bar",
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Fatalf("tool buffer:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}

	// Expanding the callers of a caller loads them.
	h.SendInput(t, []byte("jjj\r"))
	for range 400 {
		if got = waitTool(t, h); len(got) > 4 {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	want = append(want[:3], "  ~ callers of bar",
		"    ~ baz  main.go:7",
		"        7:14  func baz() { bar(); bar() }",
		"        7:21  func baz() { bar(); bar() }",
		"      > callers of baz",
	)
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("tool buffer:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}

func TestIntegration_Rename(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	main, other := filepath.Join(dir, "main.go"), filepath.Join(dir, "other.go")
	if err := os.WriteFile(other, []byte("package main\n\nvar _ = foo(2)\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	edit := func(path string, line, from, to uint32, text string) lsp.DocumentChange {
		return lsp.DocumentChange{URI: uri.File(path), Edits: []protocol.TextEdit{{
			Range:   protocol.Range{Start: protocol.Position{Line: line, Character: from}, End: protocol.Position{Line: line, Character: to}},
			NewText: text,
		}}}
	}
	fake := &fakeLSP{}
	fake.renameTarget = &lsp.RenameTarget{Placeholder: "foo"}
	fake.renameEdit = &lsp.WorkspaceEdit{Changes: []lsp.DocumentChange{
		edit(main, 2, 5, 8, "bar"), edit(main, 4, 8, 11, "bar"), edit(other, 2, 8, 11, "bar"),
	}}
	var le Integration
	h, buf, data := newBuffer(t, &le, fake, main, "package main\n\nfunc foo(x int) {}\n\nvar _ = foo(1)\n")
	h.Run(t)

	h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
		data.Rename(e, content.Position{Line: 2, Col: 6})
	}))
	// The prompt opens with the current name.
	for range 400 {
		if onLoop(t, h, func() bool { return h.Prompting(buf) }) {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	h.SendInput(t, []byte("\x7f\x7f\x7fbar\r"))
	want := "package main\n\nfunc bar(x int) {}\n\nvar _ = bar(1)\n"
	for range 400 {
		if onLoop(t, h, buf.Text) == want {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	if got := onLoop(t, h, buf.Text); got != want {
		t.Errorf("buffer text %q, want %q", got, want)
	}
	if data, _ := os.ReadFile(other); string(data) != "package main\n\nvar _ = bar(2)\n" {
		t.Errorf("other.go: %q", data)
	}
	// The server got the edit of the open buffer, and the change of the other file.
	for range 400 {
		fake.mu.Lock()
		synced := fake.text == want && len(fake.watched) > 0
		fake.mu.Unlock()
		if synced {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	fake.mu.Lock()
	renamed, watched := fake.renamed, fake.watched
	fake.mu.Unlock()
	if strings.Join(renamed, ",") != "bar" {
		t.Errorf("renamed to %q", renamed)
	}
	if len(watched) != 1 || watched[0].URI != uri.File(other) {
		t.Errorf("watched files changed: %+v", watched)
	}
	// Undo reverts the renaming in the buffer at once.
	h.SendInput(t, []byte{0x1a})
	if got := onLoop(t, h, buf.Text); got != "package main\n\nfunc foo(x int) {}\n\nvar _ = foo(1)\n" {
		t.Errorf("after undo: %q", got)
	}

	// A failed renaming is reported.
	fake.mu.Lock()
	fake.renameErr = errors.New("conflicts with x")
	fake.mu.Unlock()
	h.Post(t, editor.CommandFunc(func(e *editor.Editor) { data.Rename(e, content.Position{Line: 2, Col: 6}) }))
	for range 400 {
		if onLoop(t, h, func() bool { return h.Prompting(buf) }) {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	h.SendInput(t, []byte("x\r"))
	if got := waitTool(t, h); len(got) != 1 || got[0] != "  rename: conflicts with x" {
		t.Errorf("tool buffer %q", got)
	}
}

func TestIntegration_CodeActions(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	main := filepath.Join(dir, "main.go")
	newFile := filepath.Join(dir, "sub", "new.go")
	fake := &fakeLSP{}
	fake.actions = []lsp.CodeAction{
		{Title: "Inline call", Kind: protocol.RefactorInline, Data: []byte(`{"x": 1}`)},
		{Title: "Extract to new file", Kind: protocol.RefactorExtract, Edit: &lsp.WorkspaceEdit{Changes: []lsp.DocumentChange{
			{Kind: lsp.ChangeCreate, URI: uri.File(newFile)},
			{URI: uri.File(newFile), Edits: []protocol.TextEdit{{NewText: "package main\n"}}},
		}}},
		{Title: "Browse doc", Kind: protocol.Source, Command: &protocol.Command{Title: "doc", Command: "gopls.doc", Arguments: []any{main}}},
	}
	fake.resolveEdit = &lsp.WorkspaceEdit{Changes: []lsp.DocumentChange{{URI: uri.File(main), Edits: []protocol.TextEdit{{
		Range:   protocol.Range{Start: protocol.Position{Line: 2}, End: protocol.Position{Line: 2, Character: 4}},
		NewText: "// inlined\nfunc",
	}}}}}
	var le Integration
	h, buf, data := newBuffer(t, &le, fake, main, "package main\n\nfunc foo() {}\n")
	fake.applyEdit = data.srv.applyEditRequest
	h.Run(t)

	pick := func(query string) {
		t.Helper()
		h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
			data.ShowCodeActions(e, content.Span{Start: content.Position{Line: 2}, End: content.Position{Line: 2}})
		}))
		for range 400 {
			if onLoop(t, h, func() bool { return h.Prompting(buf) }) {
				break
			}
			time.Sleep(5 * time.Millisecond)
		}
		h.SendInput(t, []byte(query+"\r"))
	}
	waitText := func(want string) {
		t.Helper()
		for range 400 {
			if onLoop(t, h, buf.Text) == want {
				return
			}
			time.Sleep(5 * time.Millisecond)
		}
		t.Errorf("buffer text %q, want %q", onLoop(t, h, buf.Text), want)
	}

	// An action resolved to its edit.
	pick("inline")
	waitText("package main\n\n// inlined\nfunc foo() {}\n")

	// An action creating a file.
	pick("new file")
	for range 400 {
		if data, err := os.ReadFile(newFile); err == nil && string(data) == "package main\n" {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	if data, err := os.ReadFile(newFile); err != nil || string(data) != "package main\n" {
		t.Errorf("new file: %q, %v", data, err)
	}

	// An action run by the server, which edits through the client.
	pick("doc")
	waitText("// by the command\npackage main\n\n// inlined\nfunc foo() {}\n")
	fake.mu.Lock()
	executed := fake.executed
	fake.mu.Unlock()
	if strings.Join(executed, ",") != "gopls.doc" {
		t.Errorf("executed %q", executed)
	}
}

func TestIntegration_ClosesDocument(t *testing.T) {
	fake := &fakeLSP{}
	fake.events = make(chan struct{}, 1)
	var le Integration
	h, buf, _ := newGoBuffer(t, &le, fake, "package main\n")
	fake.waitVersion(t, 1)
	h.OpenReader("other.txt", strings.NewReader("x"))
	if !h.CloseBuffer(buf) {
		t.Fatal("not closed")
	}
	fake.waitEvent(t)
	fake.mu.Lock()
	defer fake.mu.Unlock()
	if len(fake.closed) != 1 || fake.closed[0] != uri.File(filepath.Join(mustAbs(t, "."), "completion_buf.go")) {
		t.Errorf("closed %v", fake.closed)
	}
}

func mustAbs(t *testing.T, p string) string {
	t.Helper()
	abs, err := filepath.Abs(p)
	if err != nil {
		t.Fatal(err)
	}
	return abs
}

func TestIntegration_WatchesFiles(t *testing.T) {
	dir := mustAbs(t, t.TempDir())
	t.Chdir(dir)
	fake := &fakeLSP{}
	fake.events = make(chan struct{}, 1)
	var le Integration
	_, _, data := newBuffer(t, &le, fake, filepath.Join(dir, "main.go"), "package main\n")
	fake.waitVersion(t, 1)

	// The server asks to watch the Go files of the workspace, like gopls does.
	data.srv.watcher.watch([]string{filepath.ToSlash(dir) + "/**/*.{go,mod}"})
	for range 400 {
		data.srv.watcher.mu.Lock()
		n := len(data.srv.watcher.stops)
		data.srv.watcher.mu.Unlock()
		if n == 1 {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	time.Sleep(50 * time.Millisecond) // The watch takes a moment to see the files.
	if err := os.WriteFile(filepath.Join(dir, "other.go"), []byte("package main\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "notes.txt"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	fake.waitEvent(t)
	fake.mu.Lock()
	defer fake.mu.Unlock()
	if len(fake.watched) != 1 || fake.watched[0].URI != uri.File(filepath.Join(dir, "other.go")) || fake.watched[0].Type != protocol.FileChangeTypeChanged {
		t.Errorf("watched changes %+v", fake.watched)
	}
}

func TestWatchedDirs(t *testing.T) {
	root := filepath.Join(string(filepath.Separator), "ws")
	dirs := watchedDirs([]string{
		"/ws/mod/**/*.{go,mod,sum,work}",
		"/ws/mod/vendor/**/*.go",
		"**/*.{go,mod}",
		"/ws/other/go.work",
		"/ws/other/**/*.tmpl",
	}, root)
	var got []string
	for _, d := range dirs {
		got = append(got, fmt.Sprintf("%s go=%t txt=%t work=%t", filepath.ToSlash(d.dir),
			d.accept("/x/a.go"), d.accept("/x/a.txt"), d.accept("/x/go.work")))
	}
	want := []string{
		"/ws go=true txt=false work=false",
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("watched dirs:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}

	dirs = watchedDirs([]string{"/ws/other/go.work", "/ws/mod/**/*.go"}, root)
	got = got[:0]
	for _, d := range dirs {
		got = append(got, fmt.Sprintf("%s go=%t work=%t", filepath.ToSlash(d.dir), d.accept("/x/a.go"), d.accept("/x/go.work")))
	}
	want = []string{"/ws/mod go=true work=false", "/ws/other go=true work=true"}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("watched dirs:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}

func TestIdentAround(t *testing.T) {
	for _, tc := range []struct {
		line string
		col  int
		want string
	}{
		{"foo(bar)", 0, "foo"},
		{"foo(bar)", 2, "foo"},
		{"foo(bar)", 3, "foo"},
		{"foo(bar)", 4, "bar"},
		{"foo(bar)", 7, "bar"},
		{"x = y", 2, "column 3"},
		{"日本", 3, "日本"},
	} {
		if got := identAround(tc.line, tc.col); got != tc.want {
			t.Errorf("identAround(%q, %d) = %q, want %q", tc.line, tc.col, got, tc.want)
		}
	}
}
