package editor_test

import (
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/editor"
)

// recordingExt is a stub Extension that counts how often each method is invoked
// and lets a test dictate what HandleInsertInput reports back to the editor.
type recordingExt struct {
	id string

	makeBufferData int
	afterEdit      int
	handleInsert   int

	handled, changed bool // what HandleInsertInput returns
	lastInsert       []byte
}

func (r *recordingExt) ID() string { return r.id }

func (r *recordingExt) MakeBufferData(*editor.Buffer) editor.BufferExtData {
	r.makeBufferData++
	return stubExtData{}
}

// stubExtData is a non-nil BufferExtData so tests can confirm the editor stored
// it under the extension's ID.
type stubExtData struct{}

func (stubExtData) TextSuggestion() string { return "" }

func (r *recordingExt) AfterEdit(*editor.Editor, *editor.Buffer) { r.afterEdit++ }

func (r *recordingExt) HandleInsertInput(_ *editor.Buffer, _ *editor.RenderPrefs, b []byte) (handled, changed bool) {
	r.handleInsert++
	r.lastInsert = append([]byte(nil), b...)
	return r.handled, r.changed
}

// TestEditor_MakeBufferDataOnOpen verifies the editor asks each extension to
// build its per-buffer data whenever a buffer is opened.
func TestEditor_MakeBufferDataOnOpen(t *testing.T) {
	ext := &recordingExt{id: "rec"}
	h := editor.NewTestHarness()
	h.Extend(ext)

	if err := h.OpenReader("a.txt", strings.NewReader("hi")); err != nil {
		t.Fatalf("OpenReader: %v", err)
	}

	if ext.makeBufferData != 1 {
		t.Errorf("MakeBufferData called %d times, want 1", ext.makeBufferData)
	}
	// The returned data must be reachable via the extension's ID, which the
	// editor uses as the storage key.
	if _, ok := h.ActiveBuffer().ExtensionData(ext.id).(stubExtData); !ok {
		t.Errorf("ExtensionData(%q) did not return the data built by the extension", ext.id)
	}
}

// TestEditor_InsertInputDispatch verifies that insert-mode input always reaches
// HandleInsertInput, and that AfterEdit fires exactly when the resulting edit
// changed the buffer — whether the change came from the extension or the
// editor's own insertion.
func TestEditor_InsertInputDispatch(t *testing.T) {
	cases := []struct {
		name       string
		handled    bool // HandleInsertInput's reported "handled"
		changed    bool // HandleInsertInput's reported "changed"
		wantHandle int
		wantAfter  int
		wantText   string
	}{
		{
			name:       "unhandled falls through to editor insertion",
			handled:    false,
			changed:    false,
			wantHandle: 1,
			wantAfter:  1,
			wantText:   "ahi", // 'a' typed at the start of "hi"
		},
		{
			name:       "handled without change skips insertion and AfterEdit",
			handled:    true,
			changed:    false,
			wantHandle: 1,
			wantAfter:  0,
			wantText:   "hi",
		},
		{
			name:       "handled with change triggers AfterEdit only",
			handled:    true,
			changed:    true,
			wantHandle: 1,
			wantAfter:  1,
			wantText:   "hi", // extension claims the input, so no editor insertion
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ext := &recordingExt{id: "rec", handled: tc.handled, changed: tc.changed}
			h := editor.NewTestHarness()
			h.Extend(ext)
			if err := h.OpenReader("a.txt", strings.NewReader("hi")); err != nil {
				t.Fatalf("OpenReader: %v", err)
			}
			h.SetMode(editor.ModeInsert)

			h.SendInput([]byte{'a'})

			if ext.handleInsert != tc.wantHandle {
				t.Errorf("HandleInsertInput called %d times, want %d", ext.handleInsert, tc.wantHandle)
			}
			if ext.afterEdit != tc.wantAfter {
				t.Errorf("AfterEdit called %d times, want %d", ext.afterEdit, tc.wantAfter)
			}
			if string(ext.lastInsert) != "a" {
				t.Errorf("HandleInsertInput saw %q, want %q", ext.lastInsert, "a")
			}
			if got := h.ActiveBuffer().Text(); got != tc.wantText {
				t.Errorf("buffer text = %q, want %q", got, tc.wantText)
			}
		})
	}
}
