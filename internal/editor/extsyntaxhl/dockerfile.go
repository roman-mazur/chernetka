package extsyntaxhl

import (
	"rmazur.io/chernetka/internal/content/parser/dockerfile"
	treesitter "rmazur.io/chernetka/internal/editor/extsyntaxhl/gotreesitter"
)

// dockerfileHighlighter highlights Dockerfiles with the hand written scanner of the
// dockerfile package.
type dockerfileHighlighter struct{}

func newDockerfile() highlighter { return dockerfileHighlighter{} }

// reparse does nothing: like with Markdown, scanning a document is cheap enough
// that all the work happens in spans, which also runs once per revision.
func (dockerfileHighlighter) reparse(*source, []treesitter.InputEdit) error { return nil }

func (dockerfileHighlighter) Close() error { return nil }

func (dockerfileHighlighter) spans(src *source, emit func(rawSpan)) {
	dockerfile.Scan(src.lines, func(s dockerfile.Span) {
		emit(lineSpan(s.Line, s.Start, s.End, s.Token))
	})
}
