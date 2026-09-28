package extsyntaxhl

import (
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/content/parser/markdown"
)

// mdHighlighter highlights Markdown with the hand written scanner of the markdown package.
// Besides the highlighted regions, the scanner finds the fenced code blocks,
// which mdHighlighter exposes implementing code.Blocks.
type mdHighlighter struct {
	blocks []code.Block // found by the last spans call
}

func newMarkdown() highlighter { return new(mdHighlighter) }

// reparse does nothing: scanning a document is cheap enough that all the work
// happens in spans, which also runs once per revision.
func (*mdHighlighter) reparse(*source) {}

func (*mdHighlighter) Close() error { return nil }

func (m *mdHighlighter) spans(src *source, emit func(rawSpan)) {
	// A new slice for every revision: the previous one may still be in use.
	m.blocks = markdown.Scan(src.lines, func(s markdown.Span) {
		emit(lineSpan(s.Line, s.Start, s.End, s.Token))
	})
}

// CodeBlocks implements code.Blocks. It reports the blocks found by the last spans call.
func (m *mdHighlighter) CodeBlocks() []code.Block { return m.blocks }
