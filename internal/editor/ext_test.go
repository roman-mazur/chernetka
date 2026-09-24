package editor_test

import (
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/editor"
)

// recordingExt is a stub Extension that counts how often each method is invoked
// and lets a test dictate what HandleInsertInput reports back to the editor.
type recordingExt struct {
	id string

	makeBufferData chan struct{}
	afterEdit      chan struct{}
	handleInsert   chan struct{}
	lastInsert     chan []byte

	handled bool // what HandleInsertInput returns
}

func initRecordingExt(id string, handled bool) *recordingExt {
	return &recordingExt{
		id:      id,
		handled: handled,

		makeBufferData: make(chan struct{}, 1),
		afterEdit:      make(chan struct{}, 1),
		handleInsert:   make(chan struct{}, 1),
		lastInsert:     make(chan []byte, 1),
	}
}

func (r *recordingExt) ID() string { return r.id }

func (r *recordingExt) MakeBufferData(*editor.Buffer) editor.BufferExtData {
	r.makeBufferData <- struct{}{}
	return stubExtData{}
}

func (r *recordingExt) AfterEdit(*editor.Editor, *editor.Buffer) {
	r.afterEdit <- struct{}{}
}

func (r *recordingExt) HandleInsertInput(_ *editor.Buffer, _ *editor.RenderPrefs, b []byte) (handled bool) {
	r.handleInsert <- struct{}{}
	r.lastInsert <- append([]byte(nil), b...)
	return r.handled
}

// stubExtData is a non-nil BufferExtData so tests can confirm the editor stored
// it under the extension's ID.
type stubExtData struct{}

func (stubExtData) TextSuggestion() string { return "" }

// TestEditor_MakeBufferDataOnOpen verifies the editor asks each extension to
// build its per-buffer data whenever a buffer is opened.
func TestEditor_MakeBufferDataOnOpen(t *testing.T) {
	ext := initRecordingExt("rec", false)
	h := editor.NewTestHarness()
	h.Extend(ext)

	if err := h.OpenReader("a.txt", strings.NewReader("hi")); err != nil {
		t.Fatalf("OpenReader: %v", err)
	}

	if !checkCallbackInvoked(ext.makeBufferData, time.Millisecond) {
		t.Errorf("MakeBufferData note called")
	}
	// The returned data must be reachable via the extension's ID, which the
	// editor uses as the storage key.
	if _, ok := h.Top().ExtensionData(ext.id).(stubExtData); !ok {
		t.Errorf("ExtensionData(%q) did not return the data built by the extension", ext.id)
	}
}

// TestEditor_InputHandle verifies that insert-mode input always reaches
// HandleInsertInput, and that AfterEdit fires exactly when the resulting edit
// changed the buffer — whether the change came from the extension or the
// editor's own insertion.
func TestEditor_InputHandle(t *testing.T) {
	cases := []struct {
		name       string
		handled    bool // HandleInsertInput's reported "handled"
		changed    bool // HandleInsertInput's reported "changed"
		wantHandle bool
		wantAfter  bool
		wantText   string
	}{
		{
			name:       "unhandled falls through to editor insertion",
			handled:    false,
			changed:    false,
			wantHandle: true,
			wantAfter:  true,
			wantText:   "hi\na",
		},
		{
			name:       "handled without change skips insertion and AfterEdit",
			handled:    true,
			changed:    false,
			wantHandle: true,
			wantAfter:  false,
			wantText:   "hi\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ext := initRecordingExt("rec", tc.handled)
			h := editor.NewTestHarness()
			h.Extend(ext)
			if err := h.OpenReader("a.txt", strings.NewReader("hi")); err != nil {
				t.Fatalf("OpenReader: %v", err)
			}

			h.Run(t)

			h.SendInput(t, []byte{'o'})
			if called := checkCallbackInvoked(ext.afterEdit, time.Second); !called {
				t.Errorf("AfterEdit called after line insertion = %t", called)
			}
			h.SendInput(t, []byte{'a'})

			if called := checkCallbackInvoked(ext.handleInsert, time.Second); called != tc.wantHandle {
				t.Errorf("HandleInsertInput called = %t, want %t", called, tc.wantHandle)
			}
			if called := checkCallbackInvoked(ext.afterEdit, time.Second); called != tc.wantAfter {
				t.Errorf("AfterEdit called = %t, want %t", called, tc.wantAfter)
			}
			lastInsert := <-ext.lastInsert
			if string(lastInsert) != "a" {
				t.Errorf("HandleInsertInput saw %q, want %q", string(lastInsert), "a")
			}

			editorText := make(chan string, 1)
			h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
				editorText <- e.Top().Text()
			}))
			if got := checkData(t, editorText); got != tc.wantText {
				t.Errorf("buffer text = %q, want %q", got, tc.wantText)
			}
		})
	}
}

func checkCallbackInvoked(ch chan struct{}, timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case <-ch:
		return true
	case <-timer.C:
		return false
	}
}

func checkData[T comparable](t *testing.T, ch chan T) T {
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	var zero T
	select {
	case d := <-ch:
		return d
	case <-timer.C:
		t.Error("timeout waiting for data from the callback")
		return zero
	}
}

// TestEditor_AfterEditOnClipboard verifies that the edits made with the clipboard
// notify the extensions like the typed ones do, while copying text does not.
func TestEditor_AfterEditOnClipboard(t *testing.T) {
	ext := initRecordingExt("rec", false)
	h := editor.NewTestHarness()
	h.Extend(ext)
	if err := h.OpenReader("a.txt", strings.NewReader("first\nsecond")); err != nil {
		t.Fatalf("OpenReader: %v", err)
	}
	h.Run(t)

	selectFirstLine := editor.CommandFunc(func(e *editor.Editor) {
		editor.SelectLine.DoOnBuffer(e.Top(), editor.RenderPrefs{})
	})
	checkText := func(want string) {
		t.Helper()
		text := make(chan string, 1)
		h.Post(t, editor.CommandFunc(func(e *editor.Editor) { text <- e.Top().Text() }))
		if got := checkData(t, text); got != want {
			t.Errorf("buffer text = %q, want %q", got, want)
		}
	}

	for _, step := range []struct {
		name      string
		prepare   editor.Command
		input     string
		wantAfter bool
		wantText  string
	}{
		{
			name:      "bracketed paste",
			input:     "\x1b[200~pasted \x1b[201~",
			wantAfter: true,
			wantText:  "pasted first\nsecond",
		},
		{
			name:      "copy",
			prepare:   selectFirstLine,
			input:     "\x03", // Ctrl+C
			wantAfter: false,
			wantText:  "pasted first\nsecond",
		},
		{
			name:      "cut",
			prepare:   selectFirstLine,
			input:     "\x18", // Ctrl+X
			wantAfter: true,
			wantText:  "\nsecond",
		},
		{
			name:      "paste",
			input:     "\x16", // Ctrl+V
			wantAfter: true,
			wantText:  "pasted first\nsecond",
		},
	} {
		if step.prepare != nil {
			h.Post(t, step.prepare)
		}
		h.SendInput(t, []byte(step.input))

		timeout := time.Second
		if !step.wantAfter {
			timeout = 50 * time.Millisecond
		}
		if called := checkCallbackInvoked(ext.afterEdit, timeout); called != step.wantAfter {
			t.Errorf("%s: AfterEdit called = %t, want %t", step.name, called, step.wantAfter)
		}
		checkText(step.wantText)
	}
}
