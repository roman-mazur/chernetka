package extsyntaxhl

import (
	"iter"

	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/content/parser/dockerfile"
)

// dockerfileHighlighter highlights Dockerfiles with the hand written scanner of the
// dockerfile package.
type dockerfileHighlighter struct{}

func newDockerfile() highlighter { return dockerfileHighlighter{} }

// reparse does nothing: like with Markdown, scanning a document is cheap enough
// that all the work happens in spans, which also runs once per revision.
func (dockerfileHighlighter) reparse(*source, []textEdit) error { return nil }

func (dockerfileHighlighter) Close() error { return nil }

func (dockerfileHighlighter) spans(src *source) iter.Seq[code.Token] {
	return dockerfile.Scan(src.lines)
}
