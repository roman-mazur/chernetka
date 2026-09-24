package code

import "rmazur.io/chernetka/internal/content"

// Block is a code block embedded into a markup document, like a fenced code block in Markdown.
//
// The Span covers the whole block including its fences: it starts at the opening fence and ends
// after the closing one. A block that is not closed runs to the end of the document.
type Block struct {
	content.Span
	Lang   string // language of the block, e.g. the first word of a Markdown fence info string
	Closed bool   // whether the block has a closing fence
}

// ContentLines returns the half-open range of the lines between the fences.
func (b Block) ContentLines() (start, end int) {
	start, end = b.Start.Line+1, b.End.Line
	if !b.Closed {
		end++ // The last line belongs to the content unless it's the opening fence itself.
	}
	return start, max(start, end)
}

// Blocks is implemented by the parsed documents that may embed code blocks.
type Blocks interface {
	// CodeBlocks returns the blocks of the current document content sorted by the start position.
	// The result must not be modified.
	CodeBlocks() []Block
}
