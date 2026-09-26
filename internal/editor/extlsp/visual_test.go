package extlsp

import (
	"context"
	"strings"
	"testing"

	"go.lsp.dev/protocol"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl"
	"rmazur.io/chernetka/internal/vt/escape"
)

// TestVisual renders inline suggestions into the test log so they can be
// inspected by eye with `go test -v -run TestVisual ./internal/editor/extlsp/`.
// Each case shows the line being typed with the suggestion as ghost text after
// the cursor and its info on the right. It asserts nothing beyond the ghost text
// being rendered: it is here to be looked at.
func TestVisual(t *testing.T) {
	const src = "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tname := \"world\"\n\t%s\n}\n"
	imp := []protocol.TextEdit{{NewText: "\n\t\"strings\""}}

	for _, tc := range []struct {
		name        string
		typed       string // the line being typed, the cursor is at its end...
		after       string // ...followed by this text
		suggestions []suggestion
		next        int // how many times to press ↓
	}{
		{
			name:  "candidates disagree: their common part, alternatives on the right",
			typed: "fmt.Pri",
			suggestions: []suggestion{
				{text: "nt", cursor: 2, info: "Print · Printf · Println"},
				{text: "nt()", cursor: 3, label: "Print", info: "func(a ...any) (n int, err error)"},
				{text: "ntf()", cursor: 4, label: "Printf", info: "func(format string, a ...any) (n int, err error)"},
				{text: "ntln()", cursor: 5, label: "Println", info: "func(a ...any) (n int, err error)"},
			},
		},
		{
			name:  "cycled with ↓ to the third alternative",
			typed: "fmt.Pri",
			suggestions: []suggestion{
				{text: "nt", cursor: 2, info: "Print · Printf · Println"},
				{text: "nt()", cursor: 3, label: "Print", info: "func(a ...any) (n int, err error)"},
				{text: "ntf()", cursor: 4, label: "Printf", info: "func(format string, a ...any) (n int, err error)"},
				{text: "ntln()", cursor: 5, label: "Println", info: "func(a ...any) (n int, err error)"},
			},
			next: 2,
		},
		{
			name:        "single candidate: a call with its signature",
			typed:       "fmt.Println(na",
			after:       ")",
			suggestions: []suggestion{{text: "me", cursor: 2, label: "name", info: "string"}},
		},
		{
			name:  "adds an import when accepted",
			typed: "strings.ToU",
			suggestions: []suggestion{
				{text: "pper()", cursor: 5, label: "ToUpper", info: "func(s string) string", edits: imp},
			},
		},
		{
			name:  "in the middle of a line",
			typed: "fmt.Println(fmt.Spr",
			after: `, "!")`,
			suggestions: []suggestion{
				{text: "int", cursor: 3, info: "Sprint · Sprintf · Sprintln"},
				{text: "intf()", cursor: 5, label: "Sprintf", info: "func(format string, a ...any) string"},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var le Integration
			le.Starter = func(context.Context, string) (lspClient, error) { return &fakeLSP{}, nil }
			t.Cleanup(func() { _ = le.Close() })

			h := editor.NewTestHarness()
			h.Extend(&le)
			h.Extend(new(extsyntaxhl.Integration))
			line := tc.typed + tc.after
			text := strings.Replace(src, "%s", line, 1)
			if err := h.OpenReader("visual.go", strings.NewReader(text)); err != nil {
				t.Fatal(err)
			}
			buf := h.Top()
			for range 6 {
				editor.RelMove{Dy: 1}.DoOnBuffer(buf, editor.RenderPrefs{})
			}
			h.SetMode(editor.ModeInsert)
			h.MoveCursorToLineEnd()
			for range len(tc.after) {
				editor.RelMove{Dx: -1}.DoOnBuffer(buf, editor.RenderPrefs{})
			}

			data := buf.ExtensionData(le.ID()).(*BufferData)
			cx, cy := buf.Pos()
			data.assign(tc.suggestions, anchor{line: "\t" + line, cx: cx, cy: cy})
			for range tc.next {
				data.SuggestNext()
			}

			out := h.RenderBuffer(100)
			if !strings.Contains(escape.Clean(out), tc.typed+data.TextSuggestion().Text+tc.after) {
				t.Errorf("no ghost text in:\n%s", escape.Clean(out))
			}
			t.Log("rendered:\n" + out)
		})
	}
}
