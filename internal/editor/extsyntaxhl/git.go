package extsyntaxhl

import (
	"iter"

	"rmazur.io/chernetka/internal/content/code"
)

// gitHighlighter highlights a file git opens in the editor, a commit message or
// a rebase todo list, with a scanner of the git package.
type gitHighlighter struct {
	scan func(lines []string) iter.Seq[code.Token]
}

func newGitHighlighter(scan func([]string) iter.Seq[code.Token]) func() highlighter {
	return func() highlighter { return gitHighlighter{scan: scan} }
}

// reparse does nothing: like with Markdown, scanning a document is cheap enough
// that all the work happens in spans, which also runs once per revision.
func (gitHighlighter) reparse(*source, []textEdit) error { return nil }

func (gitHighlighter) Close() error { return nil }

func (h gitHighlighter) spans(src *source) iter.Seq[code.Token] {
	return h.scan(src.lines)
}
