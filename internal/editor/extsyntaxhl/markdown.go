package extsyntaxhl

import (
	"iter"

	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/content/parser/markdown"
)

// mdHighlighter highlights Markdown with the hand written scanner of the markdown package.
// Besides the highlighted regions, the scanner finds the fenced code blocks,
// which mdHighlighter exposes implementing code.Blocks.
type mdHighlighter struct {
	scanner markdown.Scanner
	blocks  []code.Block // found by the last spans iteration
}

func newMarkdown() highlighter { return new(mdHighlighter) }

// reparse does nothing: the cheap scan happens in spans.
func (*mdHighlighter) reparse(*source, []textEdit) error { return nil }

func (*mdHighlighter) Close() error { return nil }

func (m *mdHighlighter) spans(src *source) iter.Seq[code.Token] {
	return func(yield func(code.Token) bool) {
		defer func() { m.blocks = m.scanner.Blocks() }()
		for t := range m.scanner.Tokens(src.lines) {
			if !yield(t) {
				return
			}
		}
	}
}

// CodeBlocks implements code.Blocks.
func (m *mdHighlighter) CodeBlocks() []code.Block { return m.blocks }
