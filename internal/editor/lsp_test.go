package editor

import (
	"context"
	"reflect"
	"testing"
	"time"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
	"rmazur.io/x/edit/internal/content"
)

func TestExtractLspSuggestions(t *testing.T) {
	item := func(label string) []protocol.CompletionItem {
		return []protocol.CompletionItem{{Label: label}}
	}
	cases := []struct {
		name  string
		items []protocol.CompletionItem
		line  string
		cx    int
		want  []string
	}{
		{name: "strips typed prefix", items: item("Println"), line: "Pri", cx: 3, want: s("ntln")},
		{name: "no prefix shows whole label", items: item("Println"), line: "x.", cx: 2, want: s("Println")},
		{name: "fully typed yields empty", items: item("Println"), line: "Println", cx: 7, want: nil},
		{name: "mismatched prefix bails out", items: item("Println"), line: "Foo", cx: 3, want: nil},
		{name: "no items", items: nil, line: "Pri", cx: 3, want: nil},
		{name: "prefix only counts identifier chars", items: item("foo"), line: "a + f", cx: 5, want: s("oo")},
		{name: "cx past line end is clamped", items: item("Println"), line: "Pri", cx: 99, want: s("ntln")},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := extractLspSuggestions(tc.items, tc.line, tc.cx); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("extractLspSuggestions(%q, cx=%d) = %v, want %v", tc.line, tc.cx, got, tc.want)
			}
		})
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

func TestUTF16Len(t *testing.T) {
	cases := []struct {
		in   string
		want uint32
	}{
		{"", 0},
		{"abc", 3},
		{"日本", 2},          // BMP runes: one code unit each.
		{"a\U0001F600", 3}, // astral rune: surrogate pair (2 units) + 'a'.
	}
	for _, tc := range cases {
		if got := utf16Len(tc.in); got != tc.want {
			t.Errorf("utf16Len(%q) = %d, want %d", tc.in, got, tc.want)
		}
	}
}

// fakeLSP records the calls afterEdit makes and returns a canned completion.
type fakeLSP struct {
	changedVersion int32
	completionCh   chan struct{}
	items          []protocol.CompletionItem
}

func (f *fakeLSP) DidOpen(context.Context, uri.URI, string, string, int32) error { return nil }
func (f *fakeLSP) DidChange(_ context.Context, _ uri.URI, _ string, v int32) error {
	f.changedVersion = v
	return nil
}
func (f *fakeLSP) Completion(context.Context, uri.URI, uint32, uint32) ([]protocol.CompletionItem, error) {
	if f.completionCh != nil {
		f.completionCh <- struct{}{}
	}
	return f.items, nil
}
func (f *fakeLSP) Shutdown(context.Context) error { return nil }

func TestEditor_AfterEditSetsSuggestion(t *testing.T) {
	fake := &fakeLSP{
		completionCh: make(chan struct{}, 1),
		items:        []protocol.CompletionItem{{Label: "Println"}},
	}
	buf := &Buffer{
		content: &content.FullText{content.TextLine("Pri")},
		mode:    ModeInsert,
		cx:      3,
		lsp: lspBufferData{
			docUri:  uri.File("/tmp/x.go"),
			version: 1,
		},
	}
	e := &Editor{
		lsp:        lspIntegration{client: fake},
		cmdChannel: make(chan Command, 1), // buffered so the async Post never blocks.
	}

	e.afterEdit(buf)

	// Wait for the goroutine to issue the completion request.
	select {
	case <-fake.completionCh:
	case <-time.After(2 * time.Second):
		t.Fatal("completion request never issued")
	}
	if fake.changedVersion != 2 {
		t.Errorf("didChange version = %d, want 2", fake.changedVersion)
	}

	// Drain the posted command and apply it on this (stand-in main-loop) thread.
	select {
	case cmd := <-e.cmdChannel:
		cmd.DoOnEditor(e)
	case <-time.After(2 * time.Second):
		t.Fatal("no command posted back")
	}
	if buf.lsp.suggestions[0] != "ntln" {
		t.Errorf("suggestion = %q, want %q", buf.lsp.suggestions[0], "ntln")
	}
}

func TestEditor_AfterEditStaleResultDropped(t *testing.T) {
	fake := &fakeLSP{
		completionCh: make(chan struct{}, 1),
		items:        []protocol.CompletionItem{{Label: "Println"}},
	}
	buf := &Buffer{
		content: &content.FullText{content.TextLine("Pri")},
		mode:    ModeInsert,
		cx:      3,
		lsp: lspBufferData{
			docUri:  uri.File("/tmp/x.go"),
			version: 1,
		},
	}
	e := &Editor{lsp: lspIntegration{client: fake}, cmdChannel: make(chan Command, 1)}

	e.afterEdit(buf)
	<-fake.completionCh

	// Simulate a newer edit arriving before the result is applied.
	buf.lsp.reqCompletion++

	select {
	case cmd := <-e.cmdChannel:
		cmd.DoOnEditor(e)
	case <-time.After(2 * time.Second):
		t.Fatal("no command posted back")
	}
	if len(buf.lsp.suggestions) != 0 {
		t.Errorf("stale suggestion applied: %v", buf.lsp.suggestions)
	}
}
