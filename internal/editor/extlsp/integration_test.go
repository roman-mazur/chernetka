package extlsp

import (
	"context"
	"errors"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/lsp"
)

func TestExtractSuggestions(t *testing.T) {
	label := func(labels ...string) (res []protocol.CompletionItem) {
		for _, l := range labels {
			res = append(res, protocol.CompletionItem{Label: l})
		}
		return
	}
	// edit replaces [from, to) of the line with text.
	edit := func(lbl string, from, to uint32, text string) protocol.CompletionItem {
		return protocol.CompletionItem{Label: lbl, TextEdit: &protocol.TextEdit{
			Range:   protocol.Range{Start: protocol.Position{Character: from}, End: protocol.Position{Character: to}},
			NewText: text,
		}}
	}
	cases := []struct {
		name  string
		items []protocol.CompletionItem
		line  string
		cx    int
		want  []string
	}{
		{name: "strips typed prefix", items: label("Println"), line: "Pri", cx: 3, want: s("ntln")},
		{name: "no prefix shows whole label", items: label("Println"), line: "x.", cx: 2, want: s("Println")},
		{name: "fully typed yields nothing", items: label("Println"), line: "Println", cx: 7, want: nil},
		{name: "complete word is not extended", items: label("func", "functions"), line: "func", cx: 4, want: nil},
		{name: "mismatched prefix bails out", items: label("Println"), line: "Foo", cx: 3, want: nil},
		{name: "no items", items: nil, line: "Pri", cx: 3, want: nil},
		{name: "prefix only counts identifier chars", items: label("foo"), line: "a + f", cx: 5, want: s("oo")},
		{name: "cx past line end is clamped", items: label("Println"), line: "Pri", cx: 99, want: s("ntln")},
		{name: "common part goes first", items: label("Printf", "Println", "Print", "PrintX"), line: "Pri", cx: 3, want: s("nt", "ntf", "ntln")},
		{name: "common prefix of top ranked items", items: label("Printf", "Println", "PrintX", "Pr"), line: "P", cx: 1, want: s("rint", "rintf", "rintln", "rintX")},
		{name: "nothing in common", items: label("Printf", "Panic"), line: "P", cx: 1, want: nil},
		{name: "text edit", items: []protocol.CompletionItem{edit("cmdChannel", 2, 5, "cmdChannel")}, line: "e.cmd", cx: 5, want: s("Channel")},
		{
			name:  "text edit replacing more than the identifier",
			items: []protocol.CompletionItem{edit("e.ch", 3, 5, "<-e.ch")},
			line:  "x <-", cx: 4, want: nil,
		},
		{
			name:  "text edit reaching past the cursor",
			items: []protocol.CompletionItem{edit("foobar", 0, 5, "foobar")},
			line:  "foo()", cx: 3, want: nil,
		},
		{
			name:  "text edit with utf16 offsets",
			items: []protocol.CompletionItem{edit("Println", 3, 6, "Println")},
			line:  "日本 Pri", cx: len("日本 Pri"), want: s("ntln"),
		},
		{name: "duplicates are merged", items: label("Println", "Println"), line: "Pri", cx: 3, want: s("ntln")},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got []string
			for _, sug := range extractSuggestions(tc.items, tc.line, tc.cx) {
				got = append(got, sug.text)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("extractSuggestions(%q, cx=%d) = %q, want %q", tc.line, tc.cx, got, tc.want)
			}
		})
	}
}

func TestShouldComplete(t *testing.T) {
	cases := []struct {
		before string
		want   bool
	}{
		{"", false},
		{"\t", false},
		{"x := ", false},
		{"foo(", false},
		{"f", true},
		{"fm", true},
		{"fmt.", true},
		{"fmt.P", true},
		{"_", true},
		{"fmt.Pr", true},
		{"x := 1.", false},
		{"x := 12", false},
		{"a.b.", true},
		{"s[0].", false},
	}
	for _, tc := range cases {
		if got := shouldComplete(tc.before); got != tc.want {
			t.Errorf("shouldComplete(%q) = %t, want %t", tc.before, got, tc.want)
		}
	}
}

func TestIdentTrailing(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Pri", "Pri"},
		{"x.Pri", "Pri"},
		{"a + foo", "foo"},
		{"foo(", ""},
		{"under_score", "under_score"},
		{"a1b2", "a1b2"},
		{"", ""},
	}
	for _, tc := range cases {
		if got := identTrailing(tc.in); got != tc.want {
			t.Errorf("identTrailing(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func s(s ...string) []string { return s }

// fakeLSP is a stand-in lspClient that keeps the document as a server would
// and returns canned completion items.
type fakeLSP struct {
	items []protocol.CompletionItem

	mu          sync.Mutex
	languageID  string           // the language of the opened document
	version     int32            // the latest document version the server knows
	text        string           // the latest text the server knows
	completions []string         // line prefixes before the cursor for every completion request
	changes     []lsp.TextChange // all the changes received
	release     chan struct{}    // when set, completions block until it's closed

	organize  func(text string) string // what organize imports makes of the text
	organized int                      // organize imports requests

	format     func(text string) string // what formatting makes of the text
	formatOpts protocol.FormattingOptions

	definitions []protocol.Location // what a definition request returns
	definedAt   []string            // the text before the position of every definition request
}

func (f *fakeLSP) DidOpen(_ context.Context, _ uri.URI, languageID, text string, v int32) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.languageID = languageID
	f.version = v
	f.text = text
	return nil
}

func (f *fakeLSP) DidChange(_ context.Context, _ uri.URI, v int32, changes ...lsp.TextChange) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.version = v
	for _, c := range changes {
		f.text = applyChange(f.text, c)
		f.changes = append(f.changes, c)
	}
	return nil
}

func (f *fakeLSP) Completion(_ context.Context, _ uri.URI, line, char uint32) ([]protocol.CompletionItem, error) {
	f.mu.Lock()
	var prefix string
	if lines := strings.Split(f.text, "\n"); int(line) < len(lines) {
		l := lines[line]
		prefix = l[:byteOffset(l, char)]
	}
	f.completions = append(f.completions, prefix)
	release := f.release
	f.mu.Unlock()

	if release != nil {
		<-release
	}
	return f.items, nil
}

func (f *fakeLSP) OrganizeImports(context.Context, uri.URI) ([]protocol.TextEdit, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.organized++
	if f.organize == nil {
		return nil, nil
	}
	change := diff(f.text, f.organize(f.text))
	return []protocol.TextEdit{{Range: *change.Range, NewText: change.Text}}, nil
}

func (f *fakeLSP) Formatting(_ context.Context, _ uri.URI, opts protocol.FormattingOptions) ([]protocol.TextEdit, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.formatOpts = opts
	if f.format == nil {
		return nil, errors.New("no formatting")
	}
	change := diff(f.text, f.format(f.text))
	return []protocol.TextEdit{{Range: *change.Range, NewText: change.Text}}, nil
}

func (f *fakeLSP) Definition(_ context.Context, _ uri.URI, line, char uint32) ([]protocol.Location, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	lines := strings.Split(f.text, "\n")
	f.definedAt = append(f.definedAt, strings.Join(lines[:line], "\n")+"\n"+lines[line][:byteOffset(lines[line], char)])
	return f.definitions, nil
}

func (f *fakeLSP) Shutdown(context.Context) error { return nil }

// newGoBuffer wires le into a fresh editor backed by fake, opens a Go buffer
// holding text, and returns the buffer together with its LSP extension data.
// The starter of le is kept if set.
func newGoBuffer(t *testing.T, le *Integration, fake *fakeLSP, text string) (*editor.TestHarness, *editor.Buffer, *BufferData) {
	t.Helper()
	return newBuffer(t, le, fake, "completion_buf.go", text)
}

// newBuffer is like newGoBuffer for a buffer of any language, recognized by its path.
func newBuffer(t *testing.T, le *Integration, fake *fakeLSP, path, text string) (*editor.TestHarness, *editor.Buffer, *BufferData) {
	t.Helper()
	if le.Starter == nil {
		le.Starter = func(context.Context, string, string) (lspClient, error) { return fake, nil }
	}

	h := editor.NewTestHarness()
	// Set before the extension starts logging from its goroutine.
	h.LogEmbed = logger.Embed(logger.Prefix(t.Logf, "editor: "))
	h.Extend(le)
	t.Cleanup(func() { _ = le.Close() }) // Stop its goroutine (and logging) with the test.
	if err := h.OpenReader(path, strings.NewReader(text)); err != nil {
		t.Fatalf("OpenReader: %v", err)
	}
	buf := h.Top()
	data, _ := buf.ExtensionData(le.ID()).(*BufferData)
	if data == nil {
		t.Fatal("buffer has no LSP extension data")
	}
	return h, buf, data
}

// runPosted drains and applies the single command the integration posts back to
// the editor from its completion goroutine.
func runPosted(t *testing.T, h *editor.TestHarness) {
	t.Helper()
	select {
	case cmd := <-h.Commands():
		cmd.DoOnEditor(h.Editor)
	case <-time.After(2 * time.Second):
		t.Fatal("no command posted back")
	}
}

func TestIntegration_AfterEditSetsSuggestion(t *testing.T) {
	fake := &fakeLSP{items: []protocol.CompletionItem{{Label: "Println"}}}
	var le Integration
	h, buf, data := newGoBuffer(t, &le, fake, "Pri")
	h.SetMode(editor.ModeInsert)
	h.MoveCursorToLineEnd() // cursor sits right after "Pri"
	fake.waitVersion(t, 1)  // opened in the background

	h.Run(t)
	h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
		le.AfterEdit(e, buf)
	}))

	waitSuggestion(t, h, data, "ntln")
	fake.waitVersion(t, 2)
}

// waitVersion waits until the server knows the given document version.
func (f *fakeLSP) waitVersion(t *testing.T, want int32) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for {
		f.mu.Lock()
		got := f.version
		f.mu.Unlock()
		if got == want {
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("server document version = %d, want %d", got, want)
		}
		time.Sleep(5 * time.Millisecond)
	}
}

func TestIntegration_AfterEditDropsStaleSuggestion(t *testing.T) {
	fake := &fakeLSP{items: []protocol.CompletionItem{{Label: "Println"}}}
	var le Integration
	h, buf, data := newGoBuffer(t, &le, fake, "Pri")
	h.SetMode(editor.ModeInsert)
	h.MoveCursorToLineEnd()

	le.AfterEdit(h.Editor, buf)
	// A newer edit bumps the request counter before the in-flight completion
	// result is applied, marking that result stale.
	data.reqCompletion++

	runPosted(t, h)

	if data.HasSuggestions() {
		t.Errorf("stale suggestion applied: %v", data.suggestions)
	}
}

// TestIntegration_CompletionSeesLatestText types through the editor and checks
// that every completion request is made against the text that includes the
// last keystroke. The server must not compute completions for a stale text.
func TestIntegration_CompletionSeesLatestText(t *testing.T) {
	fake := &fakeLSP{items: []protocol.CompletionItem{{Label: "Println"}}}
	var le Integration
	h, _, data := newGoBuffer(t, &le, fake, "")
	h.Run(t)
	h.Post(t, editor.CommandFunc(func(*editor.Editor) { h.SetMode(editor.ModeInsert) }))

	h.SendInputSequence(t, "fmt.Pr")
	waitSuggestion(t, h, data, "intln")

	fake.mu.Lock()
	for _, prefix := range fake.completions {
		// Every request happened at the end of the text the server had.
		if !strings.HasPrefix("fmt.Pr", prefix) || !shouldComplete(prefix) {
			t.Errorf("completion requested for %q", prefix)
		}
	}
	fake.mu.Unlock()
	_ = le.Close()
}

// TestIntegration_TypeThrough verifies that typing the next characters of the
// suggestion keeps the rest of it visible without waiting for the server, and
// that other edits clear it right away.
func TestIntegration_TypeThrough(t *testing.T) {
	fake := &fakeLSP{items: []protocol.CompletionItem{{Label: "Println"}}}
	var le Integration
	h, _, data := newGoBuffer(t, &le, fake, "")
	h.Run(t)
	h.Post(t, editor.CommandFunc(func(*editor.Editor) { h.SetMode(editor.ModeInsert) }))

	h.SendInputSequence(t, "Pri")
	waitSuggestion(t, h, data, "ntln")

	// Hold the server answers back: the suggestion must follow the typing.
	fake.mu.Lock()
	fake.release = make(chan struct{})
	fake.mu.Unlock()

	h.SendInput(t, []byte("n"))
	if got := onLoop(t, h, data.TextSuggestion); got.Text != "tln" {
		t.Errorf("after typing n: suggestion = %q, want %q", got, "tln")
	}
	h.SendInput(t, []byte("x"))
	if got := onLoop(t, h, data.TextSuggestion); got.Text != "" {
		t.Errorf("after a mismatching char: suggestion = %q, want none", got)
	}
	close(fake.release)
	_ = le.Close()
}

// TestIntegration_NewLineClearsSuggestion guards against a suggestion made for
// one line being shown (and accepted with Tab) on the next one.
func TestIntegration_NewLineClearsSuggestion(t *testing.T) {
	fake := &fakeLSP{items: []protocol.CompletionItem{{Label: "Println"}}}
	var le Integration
	h, buf, data := newGoBuffer(t, &le, fake, "")
	h.Run(t)
	h.Post(t, editor.CommandFunc(func(*editor.Editor) { h.SetMode(editor.ModeInsert) }))

	h.SendInputSequence(t, "Pri")
	waitSuggestion(t, h, data, "ntln")

	h.SendInput(t, []byte("\r"))
	if got := onLoop(t, h, data.TextSuggestion); got.Text != "" {
		t.Errorf("suggestion %q survived a new line", got)
	}
	h.SendInput(t, []byte("\t"))
	if got, want := onLoop(t, h, buf.Text), "Pri\n\t"; got != want {
		t.Errorf("text = %q, want %q", got, want)
	}
	_ = le.Close()
}

func TestIntegration_AcceptAppliesImport(t *testing.T) {
	text := "package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\tstrings.Sp\n}"
	fake := &fakeLSP{}
	var le Integration
	h, buf, data := newGoBuffer(t, &le, fake, text)
	for range 7 {
		editor.RelMove{Dy: 1}.DoOnBuffer(buf, editor.RenderPrefs{})
	}
	h.MoveCursorToLineEnd()

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
	want := "package main\n\nimport (\n\t\"fmt\"\n\t\"strings\"\n)\n\nfunc main() {\n\tstrings.Split\n}"
	if got := buf.Text(); got != want {
		t.Errorf("text =\n%s\nwant\n%s", got, want)
	}
	if cx, cy := buf.Pos(); cx != 14 || cy != 8 {
		t.Errorf("cursor = %d:%d, want 14:8", cy, cx)
	}
}

// waitSuggestion waits until the editor shows the wanted suggestion.
func waitSuggestion(t *testing.T, h *editor.TestHarness, data *BufferData, want string) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for {
		got := onLoop(t, h, data.TextSuggestion)
		if got.Text == want {
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("suggestion = %q, want %q", got, want)
		}
		time.Sleep(5 * time.Millisecond)
	}
}

// onLoop evaluates f on the editor loop, which owns the buffer state.
func onLoop[T any](t *testing.T, h *editor.TestHarness, f func() T) (res T) {
	t.Helper()
	h.Post(t, editor.CommandFunc(func(*editor.Editor) { res = f() }))
	return res
}

func TestIntegration_NoCompletionOutsideInsertMode(t *testing.T) {
	fake := &fakeLSP{items: []protocol.CompletionItem{{Label: "Println"}}}
	var le Integration
	h, buf, data := newGoBuffer(t, &le, fake, "Pri")
	h.MoveCursorToLineEnd()

	le.AfterEdit(h.Editor, buf) // normal mode edit, like x or a paste
	fake.waitVersion(t, 2)
	_ = le.Close() // waits for the sync loop

	fake.mu.Lock()
	defer fake.mu.Unlock()
	if len(fake.completions) != 0 {
		t.Errorf("completions requested in normal mode: %q", fake.completions)
	}
	if data.HasSuggestions() {
		t.Errorf("suggestions assigned in normal mode")
	}
}

func TestExtractSuggestions_Snippets(t *testing.T) {
	snippet := func(label string, from, to uint32, text string) protocol.CompletionItem {
		return protocol.CompletionItem{
			Label:            label,
			InsertTextFormat: protocol.InsertTextFormatSnippet,
			TextEdit: &protocol.TextEdit{
				Range:   protocol.Range{Start: protocol.Position{Character: from}, End: protocol.Position{Character: to}},
				NewText: text,
			},
		}
	}

	got := extractSuggestions([]protocol.CompletionItem{snippet("Println", 4, 7, "Println(${1:})")}, "fmt.Pri", 7)
	if len(got) != 1 || got[0].text != "ntln()" || got[0].cursor != 5 {
		t.Errorf("call snippet: got %+v, want text \"ntln()\" with the cursor at 5", got)
	}

	got = extractSuggestions([]protocol.CompletionItem{snippet("Println", 4, 11, "Println(${1:})")}, "fmt.Println", 11)
	if len(got) != 0 {
		t.Errorf("complete name: got %+v, want no suggestions", got)
	}
}

// TestIntegration_IncrementalSync types, deletes, and accepts a suggestion with
// an import through the editor and checks that the server, receiving only the
// changes, ends up with exactly the buffer text.
func TestIntegration_IncrementalSync(t *testing.T) {
	text := "package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\t\n}\n"
	fake := &fakeLSP{items: []protocol.CompletionItem{{
		Label: "Split",
		AdditionalTextEdits: []protocol.TextEdit{{
			Range:   protocol.Range{Start: protocol.Position{Line: 3, Character: 6}, End: protocol.Position{Line: 3, Character: 6}},
			NewText: "\n\t\"strings\"",
		}},
	}}}
	var le Integration
	h, buf, data := newGoBuffer(t, &le, fake, text)
	h.Run(t)
	h.Post(t, editor.CommandFunc(func(*editor.Editor) {
		for range 7 {
			editor.RelMove{Dy: 1}.DoOnBuffer(buf, editor.RenderPrefs{})
		}
		h.MoveCursorToLineEnd()
		h.SetMode(editor.ModeInsert)
	}))

	h.SendInputSequence(t, "strings.Sx\x7fp") // with a backspace
	waitSuggestion(t, h, data, "lit")
	h.SendInput(t, []byte("\t"))    // accept, adding the import
	h.SendInputSequence(t, "(\"\r") // auto-paired brackets and a new line
	want := onLoop(t, h, buf.Text)
	deadline := time.Now().Add(2 * time.Second)
	for {
		fake.mu.Lock()
		got := fake.text
		fake.mu.Unlock()
		if got == want {
			break
		}
		if time.Now().After(deadline) {
			t.Fatalf("server text:\n%s\nbuffer text:\n%s", got, want)
		}
		time.Sleep(5 * time.Millisecond)
	}
	_ = le.Close()

	fake.mu.Lock()
	defer fake.mu.Unlock()
	for _, c := range fake.changes {
		if c.Range == nil {
			t.Errorf("full text sent: %q", c.Text)
		}
	}
	if !strings.Contains(want, "\"strings\"") || !strings.Contains(want, "strings.Split(\"") {
		t.Errorf("unexpected buffer text:\n%s", want)
	}
}

// TestIntegration_AddsImportWhenTyped types a call of a package that is not
// imported and checks the import is added, while an unused import that
// "organize imports" would remove stays.
func TestIntegration_AddsImportWhenTyped(t *testing.T) {
	text := "package main\n\nimport (\n\t\"fmt\"\n\t\"io\"\n)\n\nfunc main() {\n\tfmt.Println()\n}\n"
	fake := &fakeLSP{organize: func(text string) string {
		// Like gopls: add strconv, remove io.
		return strings.Replace(text, "\t\"io\"\n", "\t\"strconv\"\n", 1)
	}}
	var le Integration
	h, buf, _ := newGoBuffer(t, &le, fake, text)
	h.Run(t)
	h.Post(t, editor.CommandFunc(func(*editor.Editor) {
		for range 8 {
			editor.RelMove{Dy: 1}.DoOnBuffer(buf, editor.RenderPrefs{})
		}
		h.MoveCursorToLineEnd()
		editor.RelMove{Dx: -1}.DoOnBuffer(buf, editor.RenderPrefs{}) // inside ()
		h.SetMode(editor.ModeInsert)
	}))

	h.SendInputSequence(t, "strconv.Itoa(1")
	want := "package main\n\nimport (\n\t\"fmt\"\n\t\"io\"\n\t\"strconv\"\n)\n\nfunc main() {\n\tfmt.Println(strconv.Itoa(1))\n}\n"
	deadline := time.Now().Add(2 * time.Second)
	for onLoop(t, h, buf.Text) != want {
		if time.Now().After(deadline) {
			t.Fatalf("text:\n%s\nwant:\n%s", onLoop(t, h, buf.Text), want)
		}
		time.Sleep(5 * time.Millisecond)
	}
	// The cursor follows the text it was in.
	if cx, cy := onLoop2(t, h, buf.Pos); cy != 9 || cx != len("\tfmt.Println(strconv.Itoa(1") {
		t.Errorf("cursor = %d:%d", cy, cx)
	}
	// And the server gets the import too.
	fake.waitText(t, want)

	fake.mu.Lock()
	defer fake.mu.Unlock()
	if fake.organized != 1 {
		t.Errorf("organize imports requested %d times, want once: after \"strconv.Itoa(\"", fake.organized)
	}
}

// TestIntegration_NoImportRequestWhenImported checks the server isn't asked
// about imports for packages that are imported.
func TestIntegration_NoImportRequestWhenImported(t *testing.T) {
	text := "package main\n\nimport \"fmt\"\n\nfunc main() {\n\t\n}\n"
	fake := &fakeLSP{}
	var le Integration
	h, buf, _ := newGoBuffer(t, &le, fake, text)
	h.Run(t)
	h.Post(t, editor.CommandFunc(func(*editor.Editor) {
		for range 5 {
			editor.RelMove{Dy: 1}.DoOnBuffer(buf, editor.RenderPrefs{})
		}
		h.MoveCursorToLineEnd()
		h.SetMode(editor.ModeInsert)
	}))
	h.SendInputSequence(t, "fmt.Println(x.y, ")
	fake.waitText(t, onLoop(t, h, buf.Text))

	fake.mu.Lock()
	defer fake.mu.Unlock()
	if fake.organized != 1 { // for "x.y," only: x may be a package
		t.Errorf("organize imports requested %d times, want once", fake.organized)
	}
}

// waitText waits until the server has the given text.
func (f *fakeLSP) waitText(t *testing.T, want string) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for {
		f.mu.Lock()
		got := f.text
		f.mu.Unlock()
		if got == want {
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("server text:\n%s\nwant:\n%s", got, want)
		}
		time.Sleep(5 * time.Millisecond)
	}
}

func onLoop2[A, B any](t *testing.T, h *editor.TestHarness, f func() (A, B)) (a A, b B) {
	t.Helper()
	h.Post(t, editor.CommandFunc(func(*editor.Editor) { a, b = f() }))
	return a, b
}

var keySave = []byte{0x13} // Ctrl+S

func TestIntegration_SaveFormats(t *testing.T) {
	t.Chdir(t.TempDir())
	fake := &fakeLSP{format: func(text string) string {
		return strings.ReplaceAll(text, "x:=1", "x := 1")
	}}
	var le Integration
	h, buf, _ := newGoBuffer(t, &le, fake, "package main\n\nfunc main() {\n\tx:=1\n}\n")
	for range 3 {
		editor.RelMove{Dy: 1}.DoOnBuffer(buf, editor.RenderPrefs{})
	}
	h.MoveCursorToLineEnd()
	fake.waitVersion(t, 1) // The server is ready.

	h.Run(t)
	h.SendInput(t, keySave)

	want := "package main\n\nfunc main() {\n\tx := 1\n}\n"
	if got := waitSaved(t, "completion_buf.go"); got != want {
		t.Errorf("saved\n%s\nwant\n%s", got, want)
	}
	if got := onLoop(t, h, buf.Text); got != want {
		t.Errorf("text =\n%s\nwant\n%s", got, want)
	}
	if cx, cy := onLoop2(t, h, buf.Pos); cx != 7 || cy != 3 {
		t.Errorf("cursor = %d:%d, want 3:7", cy, cx)
	}
	// The server gets the formatted text.
	fake.waitVersion(t, 3)
	fake.mu.Lock()
	defer fake.mu.Unlock()
	if fake.text != want {
		t.Errorf("server text =\n%s\nwant\n%s", fake.text, want)
	}
	if fake.formatOpts.TabSize != 4 {
		t.Errorf("formatted with tab size %d, want the editor's 4", fake.formatOpts.TabSize)
	}
}

func TestIntegration_SaveUnformatted(t *testing.T) {
	const text = "package main\n\nfunc main() {\n\tx:=\n}\n"
	for _, tc := range []struct {
		name  string
		ready bool
	}{
		{name: "server failure", ready: true}, // As with syntax errors.
		{name: "server not ready"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Chdir(t.TempDir())
			fake := &fakeLSP{}
			var le Integration
			release := make(chan struct{})
			defer close(release)
			if !tc.ready {
				le.Starter = func(context.Context, string, string) (lspClient, error) {
					<-release
					return fake, nil
				}
			}
			h, _, _ := newGoBuffer(t, &le, fake, text)
			if tc.ready {
				fake.waitVersion(t, 1)
			}

			h.Run(t)
			h.SendInput(t, keySave)

			if got := waitSaved(t, "completion_buf.go"); got != text {
				t.Errorf("saved %q, want %q", got, text)
			}
		})
	}
}

// waitSaved waits until the file at path is written and returns its content.
func waitSaved(t *testing.T, path string) string {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for {
		data, err := os.ReadFile(path)
		if err == nil {
			return string(data)
		}
		if time.Now().After(deadline) {
			t.Fatalf("not saved: %s", err)
		}
		time.Sleep(5 * time.Millisecond)
	}
}
