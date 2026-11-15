package extsyntaxhl

import (
	"fmt"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/change"
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/editor"
)

func codeBlocksOf(t *testing.T, buf *editor.Buffer, ext *Integration) code.Blocks {
	t.Helper()
	blocks, ok := buf.ExtensionData(ext.ID()).(code.Blocks)
	if !ok {
		t.Fatalf("no code blocks for %q", buf.Path)
	}
	return blocks
}

// formatBlocks describes the blocks as "lang startLine:startCol-endLine:endCol", with "open" for the unterminated ones.
func formatBlocks(blocks []code.Block) []string {
	var res []string
	for _, b := range blocks {
		desc := fmt.Sprintf("%q %d:%d-%d:%d", b.Lang, b.Start.Line, b.Start.Col, b.End.Line, b.End.Col)
		if !b.Closed {
			desc += " open"
		}
		res = append(res, desc)
	}
	return res
}

func TestMarkdown_FindCodeBlocks(t *testing.T) {
	for _, tc := range []struct {
		name string
		md   string
		want []string
	}{
		{name: "no blocks", md: "# Title\n\ntext"},
		{name: "one block", md: "text\n```d2\na -> b\n```\ntext", want: []string{`"d2" 1:0-3:3`}},
		{name: "empty block", md: "```d2\n```", want: []string{`"d2" 0:0-1:3`}},
		{name: "no language", md: "```\nplain\n```", want: []string{`"" 0:0-2:3`}},
		{name: "info string attributes", md: "``` d2 {title=x}\na\n```", want: []string{`"d2" 0:0-2:3`}},
		{name: "tildes", md: "~~~go\na\n~~~", want: []string{`"go" 0:0-2:3`}},
		{name: "indented", md: "   ```d2\na\n   ```", want: []string{`"d2" 0:3-2:6`}},
		{name: "indented code, not a fence", md: "    ```d2\n    a\n    ```"},
		{
			name: "several blocks",
			md:   "```d2\na\n```\n```go\n```\n```d2\nb\n```",
			want: []string{`"d2" 0:0-2:3`, `"go" 3:0-4:3`, `"d2" 5:0-7:3`},
		},
		{name: "unterminated", md: "```d2\na -> b\nb -> c", want: []string{`"d2" 0:0-2:6 open`}},
		{name: "unterminated at the fence", md: "text\n```d2", want: []string{`"d2" 1:0-1:5 open`}},
		{
			name: "fence inside another block",
			md:   "````md\n```d2\na\n```\n````\n```d2\nb\n```",
			want: []string{`"md" 0:0-4:4`, `"d2" 5:0-7:3`},
		},
		{name: "closing fence is longer", md: "```d2\na\n`````", want: []string{`"d2" 0:0-2:5`}},
		{name: "closing fence is shorter", md: "````d2\na\n```\n````", want: []string{`"d2" 0:0-3:4`}},
		{name: "closing fence of another char", md: "```d2\na\n~~~\n```", want: []string{`"d2" 0:0-3:3`}},
		{name: "closing fence with info", md: "```d2\na\n```d2\n```", want: []string{`"d2" 0:0-3:3`}},
		{name: "inline code", md: "```d2`\na\n```", want: []string{`"" 2:0-2:3 open`}},
		{name: "tilde fence may have backticks", md: "~~~d2`\na\n~~~", want: []string{"\"d2`\" 0:0-2:3"}},
		{name: "not a fence", md: "``d2\na\n``"},
		{name: "in a block quote", md: "> ```d2\n> a\n> ```\ntext", want: []string{`"d2" 0:2-2:5`}},
		{name: "in nested block quotes", md: "> > ```d2\n>> a\n> > ```", want: []string{`"d2" 0:4-2:7`}},
		// The fence ends with the block quote it is in.
		{name: "block quote ends", md: "> ```d2\n> a\ntext\n```", want: []string{`"d2" 0:2-1:3 open`, `"" 3:0-3:3 open`}},
		{name: "block quote ends at the fence", md: "> ```d2\ntext", want: []string{`"d2" 0:2-0:7 open`}},
		{name: "closing fence in a block quote", md: "```d2\n> ```\n```", want: []string{`"d2" 0:0-2:3`}},
		{name: "in a list item", md: "- item\n\n    ```d2\n    a\n    ```", want: []string{`"d2" 2:4-4:7`}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			buf, ext := openDoc(t, "notes.md", tc.md)
			got := formatBlocks(codeBlocksOf(t, buf, ext).CodeBlocks())
			if fmt.Sprint(got) != fmt.Sprint(tc.want) {
				t.Errorf("CodeBlocks() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestCodeBlocks_OtherLanguages(t *testing.T) {
	buf, ext := openDoc(t, "main.go", "package main\n\n// ```d2\n// ```")
	if got := codeBlocksOf(t, buf, ext).CodeBlocks(); got != nil {
		t.Errorf("CodeBlocks() = %v for Go", got)
	}
}

// TestCodeBlocks_ParsedOnce checks that another extension reading the blocks in its AfterEdit
// gets the blocks of the edited buffer, as the syntax highlighter is notified first, while the
// document is still parsed once per revision.
func TestCodeBlocks_ParsedOnce(t *testing.T) {
	reader := new(blocksReader)
	h, buf, ext := openHarness(t, "notes.md", "```d2\na\n```", reader)
	doc := buf.ExtensionData(ext.ID()).(*document)
	doc.SyntaxSpans(0, "```d2") // Like the first render does.

	buf.Mutate().Insert(0, content.TextLine("# Title"))
	h.NotifyExtensions(buf)
	if got, want := reader.blocks, [][]string{{`"d2" 1:0-3:3`}}; fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("the blocks read in AfterEdit = %v, want %v", got, want)
	}
	parses := doc.parses
	if spans := doc.SyntaxSpans(0, "# Title"); len(spans) == 0 {
		t.Error("the edited line is not highlighted")
	}
	if doc.parses != parses {
		t.Error("the same revision is parsed again on render")
	}
}

// blocksReader is an extension reading the code blocks in AfterEdit, like extmd and extd2 do.
type blocksReader struct{ blocks [][]string }

func (r *blocksReader) ID() string { return "blocks reader" }
func (r *blocksReader) MakeBufferData(editor.Sender, *editor.Buffer) editor.BufferExtData {
	return nil
}
func (r *blocksReader) AfterEdit(_ editor.Sender, buf *editor.Buffer, _ []change.Item) {
	if provider, ok := editor.FindExtData[code.Blocks](buf); ok {
		r.blocks = append(r.blocks, formatBlocks(provider.CodeBlocks()))
	}
}
