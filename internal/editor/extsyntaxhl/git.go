package extsyntaxhl

import (
	"rmazur.io/chernetka/internal/content/parser/git"
	treesitter "rmazur.io/chernetka/internal/editor/extsyntaxhl/gotreesitter"
)

// gitHighlighter highlights a file git opens in the editor, a commit message or
// a rebase todo list, with a scanner of the git package.
type gitHighlighter struct {
	scan func(lines []string, emit func(git.Span))
}

func newGitHighlighter(scan func([]string, func(git.Span))) func() highlighter {
	return func() highlighter { return gitHighlighter{scan: scan} }
}

// reparse does nothing: like with Markdown, scanning a document is cheap enough
// that all the work happens in spans, which also runs once per revision.
func (gitHighlighter) reparse(*source, []treesitter.InputEdit) error { return nil }

func (gitHighlighter) Close() error { return nil }

func (h gitHighlighter) spans(src *source, emit func(rawSpan)) {
	h.scan(src.lines, func(s git.Span) {
		emit(lineSpan(s.Line, s.Start, s.End, s.Token))
	})
}
