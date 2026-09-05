package extsyntaxhl

import (
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/editor"
)

// TestRenderDrivesRealEditor pushes highlighted buffers through the editor's
// own render loop. The renderer slices every line by the offsets its spans
// carry and panics on one that reaches past the end of the line, so this covers
// the contract between the highlighter and the renderer rather than the span
// offsets themselves.
func TestRenderDrivesRealEditor(t *testing.T) {
	for _, tc := range []struct{ path, text, insert string }{{
		path:   "render.go",
		insert: `var added = "x"`,
		text: strings.Join([]string{
			"package main",
			"",
			`import "fmt"`,
			"",
			`// A comment with a "quote" in it.`,
			"var raw = `a multi",
			"line raw string`",
			"",
			"func main() {",
			"\tfmt.Println(\"tab:\\there\", 1.5, nil)",
			"}",
			"",
			"func broken( {", // a parse error must not break rendering
		}, "\n"),
	}, {
		path:   "render.md",
		insert: "# a **new** line",
		text: strings.Join([]string{
			"# Heading with `code` and **bold**",
			"",
			"Setext",
			"------",
			"",
			"- [x] a task with [a link](https://example.com)",
			"  - nested *item* with ~~strike~~",
			"",
			"> quoted **text**",
			"",
			"| A | B |",
			"|---|---|",
			"| 1 | 2 |",
			"",
			"```go",
			"func main() {}",
			"```",
			"",
			"unclosed `code span and *emphasis",
			"a line with a tab\tin the middle",
			"unicode: däta — ünïcode *emphasis* ok",
		}, "\n"),
	}} {
		t.Run(tc.path, func(t *testing.T) {
			h := editor.NewTestHarness()
			h.Extend(new(Integration))
			if err := h.OpenReader(tc.path, strings.NewReader(tc.text)); err != nil {
				t.Fatal(err)
			}
			h.Run(t)

			buf := h.Top()
			lines := buf.Content.Len()

			// send delivers keys the way a terminal does, one at a time: the
			// editor ignores insert mode input that is not a single byte, so a
			// batched string would be dropped.
			send := func(keys string) {
				t.Helper()
				for i := range len(keys) {
					h.SendInput(t, []byte{keys[i]})
				}
			}

			// Walk the cursor down the whole buffer. The line under the cursor
			// renders with a background, which is the path where the renderer
			// interleaves the syntax spans with the background spans.
			send(strings.Repeat("j", lines+1))

			// Select a few lines, layering selection backgrounds over the
			// syntax spans.
			send("gg")
			for range 5 {
				h.SendInput(t, []byte("\x1b[1;2B")) // shift+down
			}

			// Type into the buffer so the highlighter reparses between renders.
			send("G")
			send("o" + tc.insert)
			h.SendInput(t, []byte("\x1b"))

			last := buf.Content.Len() - 1
			if last != lines {
				t.Fatalf("content length %d, want %d", buf.Content.Len(), lines+1)
			}
			if got := buf.Content.Lines()[last].String(); got != tc.insert {
				t.Fatalf("last line is %q, want %q", got, tc.insert)
			}

			// The inserted line has to come back highlighted, which shows the
			// edit invalidated the cached spans instead of serving stale ones.
			hl, ok := buf.ExtensionData("syntaxhl").(editor.SyntaxHighlighter)
			if !ok {
				t.Fatal("no syntax highlighter for the buffer")
			}
			if spans := hl.SyntaxSpans(last, tc.insert); len(spans) == 0 {
				t.Error("the inserted line came back with no highlight")
			}
		})
	}
}
