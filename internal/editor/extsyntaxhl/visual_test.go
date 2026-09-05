package extsyntaxhl

import (
	"fmt"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/editor"
)

// TestVisual renders highlighted buffers into the test log so the colors can be
// inspected by eye with `go test -v ./internal/editor/extsyntaxhl/`. It asserts
// nothing beyond the buffer being highlighted at all: it is here to be looked
// at, not to catch regressions. The exact offsets are covered by the tests in
// markdown_test.go and treesitter_test.go.
func TestVisual(t *testing.T) {
	t.Log("token colors:\n" + editor.TokenLegend())

	for _, tc := range []struct{ path, text string }{
		{"visual.md", visualMarkdown},
		{"visual.go", visualGo},
	} {
		t.Run(tc.path, func(t *testing.T) {
			h := editor.NewTestHarness()
			h.Extend(new(Integration))
			if err := h.OpenReader(tc.path, strings.NewReader(tc.text)); err != nil {
				t.Fatal(err)
			}
			buf := h.Top()

			hl, ok := buf.ExtensionData("syntaxhl").(editor.SyntaxHighlighter)
			if !ok {
				t.Fatalf("no syntax highlighter for %q", tc.path)
			}

			t.Log("rendered:\n" + h.RenderBuffer(100))

			// The same content again, annotated with the token type behind every
			// colorized run, for when the colors alone are ambiguous.
			var spans strings.Builder
			for i, line := range buf.Content.Lines() {
				text := line.String()
				names := make([]string, 0, 8)
				for _, s := range hl.SyntaxSpans(i, text) {
					names = append(names, fmt.Sprintf("%s=%q", s.TokenType, text[s.Start:s.End]))
				}
				fmt.Fprintf(&spans, "%2d | %-52s | %s\n", i+1, text, strings.Join(names, " "))
			}
			t.Log("spans:\n" + spans.String())
		})
	}
}

const visualMarkdown = "# Chernetka notes\n" + `
Plain text with **bold**, *italic*, ` + "`inline code`" + `, ~~struck~~ and a
backslash \*escape\* that stays literal.

Setext Heading
--------------

## Lists

- [x] markdown syntax highlight
- [ ] commit description highlight
  - nested item with *emphasis*
1. ordered item
2) also ordered

> A block quote with **emphasis** inside.
> ### and a heading nested in the quote

| Language | Backend      | Done |
|----------|--------------|:----:|
| Go       | tree-sitter  | yes  |
| Markdown | ` + "`mdScanner`" + `  | yes  |

` + "```go" + `
func main() {
	fmt.Println("inside a fence, not highlighted as Go")
}
` + "```" + `

    an indented code block

Links: [inline](https://example.com "Title"), [ref][r], <https://autolink.example>,
a bare https://go.dev/ref/spec URL, and an ![image](./demo.png).

[r]: https://example.com/reference "Reference Title"

---

An <b>HTML</b> tag, and snake_case_words that must not turn italic.
`

const visualGo = `// Package visual is here to be looked at.
package visual

import (
	"fmt"
	"strings"
)

// Greeter greets.
type Greeter struct {
	Name  string
	Count int
}

const defaultName = "world"

func (g *Greeter) Greet(w *strings.Builder) error {
	if g.Name == "" {
		g.Name = defaultName
	}
	for i := range g.Count {
		fmt.Fprintf(w, "hello %s\t#%d\n", g.Name, i+1)
	}
	return nil
}

var banner = ` + "`" + `a raw string
that spans
three lines` + "`" + `

/*
A block comment
across lines.
*/

func main() {
	g := &Greeter{Name: "chernetka", Count: 3}
	literals := []any{0x1f, 1.5e3, 'r', true, false, nil}
	_, _ = g, literals
}
`
