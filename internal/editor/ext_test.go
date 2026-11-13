package editor_test

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/content/change"
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/input"
)

// recordingExt is a stub Extension that counts how often each method is invoked
// and lets a test dictate what HandleInsertKey reports back to the editor.
type recordingExt struct {
	id string

	makeBufferData chan struct{}
	afterEdit      chan struct{}
	handleInsert   chan struct{}
	lastInsert     chan input.Key

	handled bool // what HandleInsertKey returns
}

func initRecordingExt(id string, handled bool) *recordingExt {
	return &recordingExt{
		id:      id,
		handled: handled,

		makeBufferData: make(chan struct{}, 1),
		afterEdit:      make(chan struct{}, 1),
		handleInsert:   make(chan struct{}, 1),
		lastInsert:     make(chan input.Key, 1),
	}
}

func (r *recordingExt) ID() string { return r.id }

func (r *recordingExt) MakeBufferData(editor.Sender, *editor.Buffer) editor.BufferExtData {
	r.makeBufferData <- struct{}{}
	return stubExtData{}
}

func (r *recordingExt) AfterEdit(editor.Sender, *editor.Buffer, []change.Item) {
	r.afterEdit <- struct{}{}
}

func (r *recordingExt) HandleInsertKey(_ *editor.Buffer, k input.Key) (handled bool) {
	r.handleInsert <- struct{}{}
	r.lastInsert <- k
	return r.handled
}

// stubExtData is a non-nil BufferExtData so tests can confirm the editor stored
// it under the extension's ID.
type stubExtData struct{}

func (stubExtData) TextSuggestion() code.Suggestion { return code.Suggestion{} }

// TestEditor_MakeBufferDataOnOpen verifies the editor asks each extension to
// build its per-buffer data whenever a buffer is opened.
func TestEditor_MakeBufferDataOnOpen(t *testing.T) {
	ext := initRecordingExt("rec", false)
	h := editor.NewTestHarness(t)
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
// HandleInsertKey, and that AfterEdit fires exactly when the resulting edit
// changed the buffer — whether the change came from the extension or the
// editor's own insertion.
func TestEditor_InputHandle(t *testing.T) {
	cases := []struct {
		name       string
		handled    bool // HandleInsertKey's reported "handled"
		changed    bool // HandleInsertKey's reported "changed"
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
			h := editor.NewTestHarness(t)
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
				t.Errorf("HandleInsertKey called = %t, want %t", called, tc.wantHandle)
			}
			if called := checkCallbackInvoked(ext.afterEdit, time.Second); called != tc.wantAfter {
				t.Errorf("AfterEdit called = %t, want %t", called, tc.wantAfter)
			}
			lastInsert := <-ext.lastInsert
			if lastInsert != input.Rune('a') {
				t.Errorf("HandleInsertKey saw %s, want a", lastInsert)
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
	h := editor.NewTestHarness(t)
	h.Extend(ext)
	if err := h.OpenReader("a.txt", strings.NewReader("first\nsecond")); err != nil {
		t.Fatalf("OpenReader: %v", err)
	}
	h.Run(t)

	selectFirstLine := editor.CommandFunc(func(e *editor.Editor) {
		editor.MoveContentStart.DoOnBuffer(e.Top(), editor.RenderPrefs{})
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
			wantText:  "second", // The line break is cut with the line.
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

// suggestingExt provides the buffer data suggesting its text.
type suggestingExt string

func (s suggestingExt) ID() string                                                        { return string(s) }
func (s suggestingExt) AfterEdit(editor.Sender, *editor.Buffer, []change.Item)            {}
func (s suggestingExt) MakeBufferData(editor.Sender, *editor.Buffer) editor.BufferExtData { return s }
func (s suggestingExt) TextSuggestion() code.Suggestion                                   { return code.Suggestion{Text: string(s)} }

// TestFindExtData_Order verifies the data of the first registered extension is found
// when several extensions implement the same interface.
func TestFindExtData_Order(t *testing.T) {
	for range 20 { // Iterating over a map would pick a random one.
		h := editor.NewTestHarness(t)
		h.Extend(suggestingExt("first"))
		h.Extend(suggestingExt("second"))
		if err := h.OpenReader("a.txt", strings.NewReader("hi")); err != nil {
			t.Fatal(err)
		}
		ca, ok := editor.FindExtData[editor.CodeAssist](h.Top())
		if !ok || ca.TextSuggestion().Text != "first" {
			t.Fatalf("found %v, want the first extension data", ca)
		}
	}
}

// TestAfterEditChanges checks the number of changes passed to the extensions for the kinds of edits
// made through the editor loop: every change of a line, or of the lines inserted or deleted at once,
// undo and redo included. The changes of an edit splitting or joining lines are an update of a line
// and an insertion or a deletion of the others.
func TestAfterEditChanges(t *testing.T) {
	for _, tc := range []struct {
		name, text string
		inputs     []string
		changes    int
	}{
		{name: "typing", text: "one", inputs: []string{"A two\rthree\x1b"}, changes: 4 + 2 + 5},
		{name: "joining lines", text: "one\ntwo", inputs: []string{"ji\x7f\x7f\x1b"}, changes: 2 + 1},
		{name: "deleting the selected lines", text: "one\ntwo\nthree\nfour", inputs: []string{"l\x1b[1;2B\x1b[1;2Bx", "G\x1b[1;2Ax"}, changes: 2 + 2},
		{name: "deleting the text", text: "one\ntwo", inputs: []string{"\x1b[1;2B\x1b[1;2Fx"}, changes: 2},
		{name: "paste", text: "one\n", inputs: []string{"A\x1b[200~ pasted\n\ttext\nmore\x1b[201~\x1b"}, changes: 2},
		{name: "undo and redo", text: "one\ntwo", inputs: []string{"otext\x1b", "\x1a", "\x1b[122;6u", "\x1a\x1a"}, changes: 5 + 2 + 2 + 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := editor.NewTestHarness(t)
			rec := new(changesRecorder)
			h.Extend(rec)
			if err := h.OpenReader("file.txt", strings.NewReader(tc.text)); err != nil {
				t.Fatal(err)
			}
			h.Run(t)
			for _, in := range tc.inputs {
				h.SendInput(t, []byte(in))
			}
			calls := rec.take(t, h)
			changes := 0
			for _, c := range calls {
				if c == nil {
					t.Errorf("unknown changes passed for an edit")
				}
				changes += len(c)
			}
			if changes != tc.changes {
				t.Errorf("%d changes passed in %d calls, want %d: %v", changes, len(calls), tc.changes, calls)
			}
		})
	}
}

// TestAfterEditChanges_Reload checks the changes passed to the extensions when a file is reloaded,
// changed outside the editor: they are nil, the content is new, and the next edits are known again.
func TestAfterEditChanges_Reload(t *testing.T) {
	path := filepath.Join(t.TempDir(), "file.txt")
	if err := os.WriteFile(path, []byte("one\ntwo"), 0o600); err != nil {
		t.Fatal(err)
	}
	h := editor.NewTestHarness(t)
	rec := new(changesRecorder)
	h.Extend(rec)
	h.Run(t)
	h.Post(t, &editor.OpenFile{Path: path})

	h.SendInput(t, []byte("Azero\x1b"))
	if calls := rec.take(t, h); !allKnown(calls) {
		t.Fatalf("got the changes %v of typing, want them known", calls)
	}

	if err := os.WriteFile(path, []byte("reloaded\ncontent"), 0o600); err != nil {
		t.Fatal(err)
	}
	var calls [][]change.Item
	for range 200 {
		if calls = rec.take(t, h); len(calls) > 0 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	if len(calls) != 1 || calls[0] != nil {
		t.Fatalf("got the changes %v of the reload, want nil", calls)
	}

	h.SendInput(t, []byte("Anew\x1b"))
	if calls := rec.take(t, h); !allKnown(calls) {
		t.Fatalf("got the changes %v of typing after the reload, want them known", calls)
	}
}

// allKnown reports whether the extensions were notified, and got known changes every time.
func allKnown(calls [][]change.Item) bool {
	for _, changes := range calls {
		if len(changes) == 0 {
			return false
		}
	}
	return len(calls) > 0
}

// changesRecorder is an extension recording the changes passed to AfterEdit.
type changesRecorder struct {
	calls [][]change.Item
}

func (r *changesRecorder) ID() string { return "changes recorder" }
func (r *changesRecorder) MakeBufferData(editor.Sender, *editor.Buffer) editor.BufferExtData {
	return nil
}
func (r *changesRecorder) AfterEdit(_ editor.Sender, _ *editor.Buffer, changes []change.Item) {
	r.calls = append(r.calls, changes)
}

// take returns the changes recorded since the previous call, reading them on the editor loop.
func (r *changesRecorder) take(t *testing.T, h *editor.TestHarness) (calls [][]change.Item) {
	h.Post(t, editor.CommandFunc(func(*editor.Editor) { calls, r.calls = r.calls, nil }))
	return calls
}

// TestAfterEditOrder checks that the syntax highlighters are notified first, whatever order
// they are registered in.
func TestAfterEditOrder(t *testing.T) {
	var order []string
	h := editor.NewTestHarness(t)
	h.Extend(orderExt{id: "plain", order: &order})
	h.Extend(highlighterOrderExt{orderExt{id: "highlighter", order: &order}})
	h.Extend(orderExt{id: "plain2", order: &order})
	h.Extend(highlighterOrderExt{orderExt{id: "highlighter2", order: &order}})
	if err := h.OpenReader("file.txt", strings.NewReader("text")); err != nil {
		t.Fatal(err)
	}
	h.Run(t)
	h.SendInput(t, []byte("ix\x1b"))
	var got []string
	h.Post(t, editor.CommandFunc(func(*editor.Editor) { got = slices.Clone(order) }))
	if want := []string{"highlighter", "highlighter2", "plain", "plain2"}; !slices.Equal(got, want) {
		t.Errorf("notified %v, want %v", got, want)
	}
}

// orderExt records the order of the AfterEdit calls.
type orderExt struct {
	id    string
	order *[]string
}

func (oe orderExt) ID() string                                                        { return oe.id }
func (oe orderExt) MakeBufferData(editor.Sender, *editor.Buffer) editor.BufferExtData { return nil }
func (oe orderExt) AfterEdit(editor.Sender, *editor.Buffer, []change.Item) {
	*oe.order = append(*oe.order, oe.id)
}

// highlighterOrderExt is an orderExt that is a syntax highlighter.
type highlighterOrderExt struct{ orderExt }

func (highlighterOrderExt) HighlightsSyntax() {}
