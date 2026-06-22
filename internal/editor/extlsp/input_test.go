package extlsp

import (
	"context"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/editor"
)

// Raw byte sequences for the inputs HandleInsertInput reacts to.
var (
	keyArrowUp    = []byte{0x1b, '[', 'A'}
	keyArrowDown  = []byte{0x1b, '[', 'B'}
	keyArrowLeft  = []byte{0x1b, '[', 'D'}
	keyArrowRight = []byte{0x1b, '[', 'C'}
	keyEscape     = []byte{0x1b}
	keyTab        = []byte{'\t'}
)

func TestIntegration_HandleInsertInput(t *testing.T) {
	cases := []struct {
		name        string
		suggestions []string // assigned before the input; nil leaves the list empty
		input       []byte
		wantHandled bool
		wantCurrent string // CurrentSuggestion afterwards; "" means none remain
		wantText    string // buffer contents afterwards
	}{
		{
			name:     "no suggestions ignores tab",
			input:    keyTab,
			wantText: "Pri",
		},
		{
			name:        "arrow down selects next suggestion",
			suggestions: s("a", "b", "c"),
			input:       keyArrowDown,
			wantHandled: true,
			wantCurrent: "b",
			wantText:    "Pri",
		},
		{
			name:        "arrow up wraps to last suggestion",
			suggestions: s("a", "b", "c"),
			input:       keyArrowUp,
			wantHandled: true,
			wantCurrent: "c",
			wantText:    "Pri",
		},
		{
			name:        "left arrow dismisses suggestions",
			suggestions: s("a", "b"),
			input:       keyArrowLeft,
			wantText:    "Pri",
		},
		{
			name:        "right arrow dismisses suggestions",
			suggestions: s("a", "b"),
			input:       keyArrowRight,
			wantText:    "Pri",
		},
		{
			name:        "escape dismisses suggestions",
			suggestions: s("a", "b"),
			input:       keyEscape,
			wantHandled: true,
			wantText:    "Pri",
		},
		{
			name:        "tab accepts current suggestion",
			suggestions: s("ntln"),
			input:       keyTab,
			wantHandled: true,
			wantText:    "Println",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fake := &fakeLSP{}
			var le Integration
			h, buf, data := newGoBuffer(t, &le, fake, "Pri")
			h.MoveCursorToLineEnd() // cursor right after "Pri"
			if tc.suggestions != nil {
				data.Assign(tc.suggestions)
			}

			handled := le.HandleInsertInput(buf, &editor.RenderPrefs{}, tc.input)

			if handled != tc.wantHandled {
				t.Errorf("HandleInsertInput = (handled %t), want (%t)", handled, tc.wantHandled)
			}
			if tc.wantCurrent == "" {
				if data.HasSuggestions() {
					t.Errorf("suggestions still present: %v", data.suggestions)
				}
			} else if got := data.CurrentSuggestion(); got != tc.wantCurrent {
				t.Errorf("CurrentSuggestion = %q, want %q", got, tc.wantCurrent)
			}
			if got := buf.Text(); got != tc.wantText {
				t.Errorf("buffer text = %q, want %q", got, tc.wantText)
			}
		})
	}
}

// TestIntegration_HandleInsertInput_NoBufferData verifies the extension stays
// out of the way for buffers it never attached data to (e.g. non-Go files).
func TestIntegration_HandleInsertInput_NoBufferData(t *testing.T) {
	var le Integration
	le.Starter = func(context.Context, string) (lspClient, error) {
		t.Fatal("LSP should not start for a non-Go buffer")
		return nil, nil
	}

	h := editor.NewTestHarness()
	h.Extend(&le)
	if err := h.OpenReader("notes.txt", strings.NewReader("hello")); err != nil {
		t.Fatalf("OpenReader: %v", err)
	}
	buf := h.Top()

	handled := le.HandleInsertInput(buf, &editor.RenderPrefs{}, keyTab)
	if handled {
		t.Errorf("HandleInsertInput = (%t), want (false)", handled)
	}
}
