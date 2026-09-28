// Package extmd implements an editor extension that makes Markdown task list items actionable.
//
// A line like "- [ ] task" can be engaged to check the item turning it into "- [x] task", and back.
// Lines inside fenced code blocks are not actionable if another extension, like the syntax highlighter,
// provides code.Blocks for the buffer.
package extmd

import (
	"path/filepath"
	"slices"
	"strings"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/content/parser/markdown"
	"rmazur.io/chernetka/internal/editor"
)

// Integration implements an editor.Extension that provides line actions for the Markdown task list items.
type Integration struct{}

func (in *Integration) ID() string { return "md" }

func (in *Integration) MakeBufferData(buf *editor.Buffer) editor.BufferExtData {
	switch strings.ToLower(filepath.Ext(buf.Path)) {
	case ".md", ".markdown":
	default:
		return nil
	}
	lines := buf.Content.Lines()
	if len(lines) > 0 && lines[0].MimeType() != content.MimeTypeTextPlain {
		return nil // Not a file, like a directory listing.
	}
	if _, ok := buf.Content.(content.Mutable); !ok {
		return nil
	}
	return &document{buf: buf}
}

func (in *Integration) AfterEdit(editor.Sender, *editor.Buffer) {}

// document is the per-buffer extension data.
type document struct {
	buf *editor.Buffer
}

func (d *document) LineAction(lineNumber int) content.LineAction {
	lines := d.buf.Content.Lines()
	if lineNumber < 0 || lineNumber >= len(lines) {
		return nil
	}
	if markdown.TaskCheckbox(lines[lineNumber].String()) < 0 || d.inCodeBlock(lineNumber) {
		return nil
	}
	return toggle{d: d, line: lineNumber}
}

// inCodeBlock reports whether the line is between the fences of a code block.
func (d *document) inCodeBlock(lineNumber int) bool {
	provider, ok := editor.FindExtData[code.Blocks](d.buf)
	if !ok {
		return false
	}
	blocks := provider.CodeBlocks()
	// The last block starting above the line is the only one that may contain it.
	i, _ := slices.BinarySearchFunc(blocks, lineNumber, func(b code.Block, ln int) int { return b.Start.Line - ln })
	if i == 0 {
		return false
	}
	start, end := blocks[i-1].ContentLines()
	return start <= lineNumber && lineNumber < end
}

// toggle checks or unchecks the task list item on the line.
type toggle struct {
	d    *document
	line int
}

func (t toggle) Engage() {
	lines := t.d.buf.Content.Lines()
	if t.line >= len(lines) {
		return
	}
	text := lines[t.line].String()
	i := markdown.TaskCheckbox(text)
	if i < 0 {
		return
	}
	mark := "x"
	if text[i] != ' ' {
		mark = " "
	}
	t.d.buf.Mutate().Update(t.line, content.TextLine(text[:i]+mark+text[i+1:]))
}

// SingleShot tells the editor to not re-run the toggle as the last engaged action.
func (toggle) SingleShot() {}
