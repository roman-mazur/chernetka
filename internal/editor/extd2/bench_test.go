package extd2

import (
	"fmt"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/cheimg"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl"
)

type noopViewer struct{}

func (noopViewer) Show(cheimg.Item) {}

// benchSection is a piece of a Markdown document with a d2 block, repeated to get a document of the needed size.
const benchSection = `## Section %d

Some *text* with a [link](https://example.com), ` + "`code`" + `, and **bold** words.
- [ ] a task
- an item with a longer description that keeps going for a while

` + "```d2" + `
a -> b: label
b -> c
` + "```" + `

` + "```go" + `
func main() {}
` + "```" + `
`

const benchVisibleLines = 50

// benchDoc is a Markdown document open in an editor with the syntax highlighter and, optionally, the d2 extension.
type benchDoc struct {
	buf     *editor.Buffer
	hl      *extsyntaxhl.Integration
	d2      *Integration // nil when the extension is not registered
	actions content.LineActions
	editLn  int // the line changed by the edits, it's the first visible one
	edits   int
}

func newBenchDoc(b *testing.B, lines int, withD2 bool) *benchDoc {
	b.Helper()
	var sb strings.Builder
	for i := 0; strings.Count(sb.String(), "\n") < lines; i++ {
		fmt.Fprintf(&sb, benchSection, i)
	}

	var edit editor.Editor
	d := &benchDoc{hl: new(extsyntaxhl.Integration)}
	// The registration order is the one of cmd/che.
	edit.Extend(d.hl)
	if withD2 {
		d.d2 = &Integration{Viewer: noopViewer{}}
		edit.Extend(d.d2)
	}
	if err := edit.OpenReader("doc.md", strings.NewReader(sb.String())); err != nil {
		b.Fatal(err)
	}
	d.buf = edit.Top()
	if d.d2 != nil {
		d.actions = d.buf.ExtensionData(d.d2.ID()).(content.LineActions)
	}
	// Edit the text line in the middle of the document.
	d.editLn = d.buf.Content.Len() / 2
	for !strings.HasPrefix(d.buf.Content.Lines()[d.editLn].String(), "Some ") {
		d.editLn++
	}
	d.edit() // The first sync.
	return d
}

// edit changes a line and runs the AfterEdit hooks like the editor loop does after an input.
func (d *benchDoc) edit() {
	d.edits++
	line := d.buf.Content.Lines()[d.editLn].String()
	d.buf.Mutate().Update(d.editLn, content.TextLine(fmt.Sprintf("%s%d", line[:len("Some")], d.edits%10)+line[len("Some")+1:]))
	d.hl.AfterEdit(nil, d.buf)
	if d.d2 != nil {
		d.d2.AfterEdit(nil, d.buf)
	}
}

// lineActions requests the actions of the visible lines, like rendering does.
func (d *benchDoc) lineActions() {
	for ln := d.editLn; ln < d.editLn+benchVisibleLines; ln++ {
		d.actions.LineAction(ln)
	}
}

// highlight requests the syntax spans of the visible lines, like rendering does.
func (d *benchDoc) highlight() {
	hl, _ := editor.FindExtData[editor.SyntaxHighlighter](d.buf)
	lines := d.buf.Content.Lines()
	for ln := d.editLn; ln < d.editLn+benchVisibleLines; ln++ {
		hl.SyntaxSpans(ln, lines[ln].String())
	}
}

// BenchmarkMarkdown measures the cost of the d2 extension in a Markdown document.
// The d2 blocks are taken from the syntax highlighter. With d2=false, the extension is not registered.
//
// AfterEdit measures an edit followed by the AfterEdit hooks. It includes highlighting
// the whole document: the d2 extension requests the code blocks, which is done for rendering anyway.
// LineActions measures requesting the actions of the visible lines, with no edits.
// Keystroke measures what the editor does for a typed character: the edit, the AfterEdit hooks,
// and rendering the visible lines with their highlight and actions.
func BenchmarkMarkdown(b *testing.B) {
	for _, lines := range []int{300, 5000} {
		for _, withD2 := range []bool{false, true} {
			name := fmt.Sprintf("lines=%d/d2=%t", lines, withD2)
			if withD2 {
				b.Run(name+"/AfterEdit", func(b *testing.B) {
					d := newBenchDoc(b, lines, withD2)
					for b.Loop() {
						d.edit()
					}
				})
				b.Run(name+"/LineActions", func(b *testing.B) {
					d := newBenchDoc(b, lines, withD2)
					for b.Loop() {
						d.lineActions()
					}
				})
			}
			b.Run(name+"/Keystroke", func(b *testing.B) {
				if lines == 300 && withD2 {
					// A noisy one, see internal/cmd/benchmarks.
					b.Log("tolerance=20%")
				}
				d := newBenchDoc(b, lines, withD2)
				for b.Loop() {
					d.edit()
					d.highlight()
					if d.actions != nil {
						d.lineActions()
					}
				}
			})
		}
	}
}
