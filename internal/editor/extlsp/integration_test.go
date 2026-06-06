package extlsp

import (
	"context"
	"reflect"
	"strings"
	"testing"
	"time"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
	"rmazur.io/chernetka/internal/editor"
)

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

func s(s ...string) []string { return s }

// fakeLSP is a stand-in lspClient that records the document versions it is told
// about and returns canned completion items.
type fakeLSP struct {
	openVersion    int32
	changedVersion int32
	items          []protocol.CompletionItem
}

func (f *fakeLSP) DidOpen(_ context.Context, _ uri.URI, _, _ string, v int32) error {
	f.openVersion = v
	return nil
}

func (f *fakeLSP) DidChange(_ context.Context, _ uri.URI, _ string, v int32) error {
	f.changedVersion = v
	return nil
}

func (f *fakeLSP) Completion(context.Context, uri.URI, uint32, uint32) ([]protocol.CompletionItem, error) {
	return f.items, nil
}

func (f *fakeLSP) Shutdown(context.Context) error { return nil }

// newGoBuffer wires le into a fresh editor backed by fake, opens a Go buffer
// holding text, and returns the buffer together with its LSP extension data.
func newGoBuffer(t *testing.T, le *Integration, fake *fakeLSP, text string) (*editor.TestHarness, *editor.Buffer, *BufferData) {
	t.Helper()
	le.Starter = func(context.Context, string) (lspClient, error) { return fake, nil }

	h := editor.NewTestHarness()
	h.Extend(le)
	if err := h.OpenReader("completion_buf.go", strings.NewReader(text)); err != nil {
		t.Fatalf("OpenReader: %v", err)
	}
	buf := h.ActiveBuffer()
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
	h.MoveCursorToLineEnd() // cursor sits right after "Pri"

	if fake.openVersion != 1 {
		t.Errorf("DidOpen version = %d, want 1", fake.openVersion)
	}

	le.AfterEdit(h.Editor, buf)

	if fake.changedVersion != 2 {
		t.Errorf("DidChange version = %d, want 2", fake.changedVersion)
	}

	runPosted(t, h)

	if !data.HasSuggestions() {
		t.Fatal("no suggestion assigned")
	}
	if got := data.CurrentSuggestion(); got != "ntln" {
		t.Errorf("suggestion = %q, want %q", got, "ntln")
	}
}

func TestIntegration_AfterEditDropsStaleSuggestion(t *testing.T) {
	fake := &fakeLSP{items: []protocol.CompletionItem{{Label: "Println"}}}
	var le Integration
	h, buf, data := newGoBuffer(t, &le, fake, "Pri")
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
