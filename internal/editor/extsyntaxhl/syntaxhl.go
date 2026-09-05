// Package extsyntaxhl implements an editor extension that colorizes buffers.
//
// A buffer is matched to a language by its file name, and the language provides
// a highlighter: either a tree-sitter grammar driven by a highlight query, or a
// hand written scanner for the formats that have no Go grammar bindings.
// Whatever the source, a highlighter reports coarse rawSpans that the
// spanBuilder flattens into the per-line spans the renderer consumes.
package extsyntaxhl

import (
	"io"
	"path/filepath"
	"slices"
	"strings"

	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/logger"
)

// highlighter reports the highlighted regions of a document in one language.
type highlighter interface {
	// reparse rebuilds whatever state the highlighter keeps for a new revision
	// of the document. It is called on every edit, so it holds the expensive
	// work, while spans is called during rendering.
	reparse(src *source)
	// spans reports every highlighted region of the document. Reported spans may
	// cover several lines and may be nested inside one another; the caller
	// flattens them.
	spans(src *source, emit func(rawSpan))

	io.Closer
}

// language ties a file type to the highlighter that understands it.
type language struct {
	name           string
	matcher        func(filePath string) bool
	newHighlighter func() highlighter
}

var languages = []language{
	{
		name:           "go",
		matcher:        matchExtensions(".go"),
		newHighlighter: newTreeSitter(goGrammar),
	},
	{
		name:           "markdown",
		matcher:        matchExtensions(".md", ".markdown"),
		newHighlighter: newMarkdown,
	},
	{
		name: "git_commit_msg",
		matcher: func(filePath string) bool {
			return strings.HasSuffix(filePath, ".git/COMMIT_EDITMSG")
		},
		newHighlighter: func() highlighter { return new(gitMessage) },
	},
}

func matchExtensions(extensions ...string) func(filePath string) bool {
	return func(filePath string) bool {
		ext := strings.ToLower(filepath.Ext(filePath))
		return slices.Contains(extensions, ext)
	}
}

func languageForPath(path string) *language {
	for i := range languages {
		if languages[i].matcher(path) {
			return &languages[i]
		}
	}
	return nil
}

// holdsPlainText reports whether the buffer holds ordinary text lines. A
// directory listing buffer is named after its directory, which may well end in
// a recognized extension, and colorizing file names as code would be wrong.
func holdsPlainText(buf *editor.Buffer) bool {
	lines := buf.Content.Lines()
	return len(lines) == 0 || lines[0].MimeType() == "text/plain"
}

// Integration implements an editor.Extension that colorizes buffers of the
// languages it recognizes.
type Integration struct {
	logImpl
}

func (in *Integration) ID() string { return "syntaxhl" }

func (in *Integration) MakeBufferData(buf *editor.Buffer) editor.BufferExtData {
	lang := languageForPath(buf.Path)
	if lang == nil || !holdsPlainText(buf) {
		return nil
	}

	in.logf(false, "highlighting %s as %s", buf.Path, lang.name)
	doc := &document{logImpl: &in.logImpl, lang: lang, hl: lang.newHighlighter()}
	doc.reparse(buf.Text())
	return doc
}

func (in *Integration) AfterEdit(_ *editor.Editor, buf *editor.Buffer) {
	doc, ok := buf.ExtensionData(in.ID()).(*document)
	if !ok {
		return
	}
	in.logf(true, "AfterEdit(_, %q)", buf.Path)
	doc.reparse(buf.Text())
}

func (in *Integration) HandleInsertInput(*editor.Buffer, *editor.RenderPrefs, []byte) (handled bool) {
	return false
}

// document is the per-buffer extension data. It owns the language highlighter
// and caches the flattened spans of the whole document until the next edit.
type document struct {
	*logImpl
	lang *language
	hl   highlighter

	src     *source
	builder spanBuilder
	spans   []editor.SyntaxSpan // sorted by line number, then by start offset
	built   bool
}

func (d *document) reparse(text string) {
	d.src = newSource(text)
	d.hl.reparse(d.src)
	// Spans are rebuilt lazily: a buffer that is edited several times between
	// two renders is only flattened once, and one that is never rendered never
	// gets flattened at all.
	d.spans, d.built = d.spans[:0], false
}

func (d *document) build() {
	d.built = true
	d.builder.reset()
	d.hl.spans(d.src, func(rs rawSpan) { d.builder.add(d.src, rs) })
	d.spans = d.builder.build(d.spans)
	d.logf(true, "built %d spans for %d %s lines", len(d.spans), len(d.src.lines), d.lang.name)
}

// SyntaxSpans implements editor.SyntaxHighlighter.
func (d *document) SyntaxSpans(ln int, line string) []editor.SyntaxSpan {
	if !d.built {
		d.build()
	}
	return clipSpans(spansOfLine(d.spans, ln), line)
}

func (d *document) Close() error { return d.hl.Close() }

type logImpl struct {
	LogDebug bool
	LogF     logger.Func
}

func (li *logImpl) logf(debug bool, fmt string, args ...any) {
	if li.LogF == nil {
		return
	}
	if debug && !li.LogDebug {
		return
	}
	li.LogF(fmt, args...)
}
