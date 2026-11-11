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

	"rmazur.io/chernetka/internal/content/change"
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/editor"
	treesitter "rmazur.io/chernetka/internal/editor/extsyntaxhl/gotreesitter"
	"rmazur.io/chernetka/internal/logger"
)

// highlighter reports the highlighted regions of a document in one language.
type highlighter interface {
	// reparse rebuilds whatever state the highlighter keeps for a new revision
	// of the document. It is called once for the edits made between two renders,
	// so it holds the expensive work, while spans is called during rendering.
	// On an error, the document has no highlighting until the next successful reparse.
	//
	// The edits turn the previous revision into the new one, if they are known,
	// letting the highlighter reuse what it found in the unchanged parts. They
	// are nil for the first revision, and when the whole text is to be parsed.
	reparse(src *source, edits []treesitter.InputEdit) error
	// spans reports every highlighted region of the document. Reported spans may
	// cover several lines and may be nested inside one another; the caller
	// flattens them.
	spans(src *source, emit func(rawSpan))

	io.Closer
}

// language ties a syntax to the highlighter that understands it.
type language struct {
	syntax         *code.Syntax
	newHighlighter func() highlighter
}

func (l *language) name() string { return l.syntax.Name }

func languageForPath(path string) *language {
	return languages[code.SyntaxForPath(path)]
}

var languages = map[*code.Syntax]*language{
	code.Go:              {newHighlighter: newTreeSitter(goGrammar)},
	code.Nix:             {newHighlighter: newTreeSitter(nixGrammar)},
	code.CUE:             {newHighlighter: newTreeSitter(cueGrammar)},
	code.JSON:            {newHighlighter: newTreeSitter(jsonGrammar)},
	code.JSONC:           {newHighlighter: newTreeSitter(jsonGrammar)},
	code.YAML:            {newHighlighter: newTreeSitter(yamlGrammar)},
	code.Shell:           {newHighlighter: newTreeSitter(bashGrammar)},
	code.C:               {newHighlighter: newTreeSitter(cGrammar)},
	code.CPP:             {newHighlighter: newTreeSitter(cppGrammar)},
	code.CSharp:          {newHighlighter: newTreeSitter(csharpGrammar)},
	code.Java:            {newHighlighter: newTreeSitter(javaGrammar)},
	code.Kotlin:          {newHighlighter: newTreeSitter(kotlinGrammar)},
	code.TypeScript:      {newHighlighter: newTreeSitter(typescriptGrammar)},
	code.TSX:             {newHighlighter: newTreeSitter(tsxGrammar)},
	code.JavaScript:      {newHighlighter: newTreeSitter(tsxGrammar)},
	code.Rust:            {newHighlighter: newTreeSitter(rustGrammar)},
	code.Python:          {newHighlighter: newTreeSitter(pythonGrammar)},
	code.CommonLisp:      {newHighlighter: newTreeSitter(commonLispGrammar)},
	code.TreeSitterQuery: {newHighlighter: newTreeSitter(treeSitterQueryGrammar)},
	code.Markdown:        {newHighlighter: newMarkdown},
	code.GitCommitMsg:    {newHighlighter: func() highlighter { return new(gitMessage) }},
	code.GitRebaseTodo:   {newHighlighter: func() highlighter { return new(gitRebase) }},
}

func init() {
	for s, l := range languages {
		l.syntax = s
	}
}

// Integration implements an editor.Extension that colorizes buffers of the
// languages it recognizes.
type Integration struct {
	logger.LogEmbed
}

func (in *Integration) ID() string { return "syntaxhl" }

func (in *Integration) MakeBufferData(_ editor.Sender, buf *editor.Buffer) editor.BufferExtData {
	lang := languageForPath(buf.Path)
	// A directory listing buffer is named after its directory, which may well end in
	// a recognized extension, and colorizing file names as code would be wrong.
	if lang == nil || !buf.HoldsTextFile() {
		return nil
	}

	in.Logf("highlighting %s as %s", buf.Path, lang.name())
	doc := &document{LogEmbed: &in.LogEmbed, buf: buf, lang: lang, hl: lang.newHighlighter()}
	doc.parse(buf.Text())
	return doc
}

func (in *Integration) AfterEdit(_ editor.Sender, buf *editor.Buffer, _ []change.Item) {
	if doc, ok := buf.ExtensionData(in.ID()).(*document); ok {
		doc.stale = true
	}
}

// document is the per-buffer extension data. It owns the language highlighter
// and caches the flattened spans of the whole document until the next edit.
//
// The document is parsed when its spans or code blocks are requested, not on every edit:
// the keys typed between two renders are parsed at once.
type document struct {
	*logger.LogEmbed
	buf  *editor.Buffer
	lang *language
	hl   highlighter

	src      *source // the revision parsed the last time
	parseErr error   // of the last parse
	stale    bool    // the buffer is edited after src was parsed

	builder spanBuilder
	spans   []editor.SyntaxSpan // sorted by line number, then by start offset
	built   bool
}

// ensureParsed brings the document up to date with the buffer.
func (d *document) ensureParsed() {
	if d.stale {
		d.parse(d.buf.Text())
	}
}

// parse brings the document up to date with the text, reusing the last parse
// for the text both revisions share.
func (d *document) parse(text string) {
	d.stale = false
	if d.src != nil && d.src.text == text {
		return // Like a line deleted and inserted back.
	}
	var edits []treesitter.InputEdit
	if d.src != nil && d.parseErr == nil {
		edits = []treesitter.InputEdit{textEdit(d.src.text, text)}
	}
	d.src = newSource(text)
	err := d.hl.reparse(d.src, edits)
	if err != nil && d.parseErr == nil {
		// Logged once: the document may stay broken for many edits.
		d.Logf("cannot highlight %s: %s", d.buf.Path, err)
	}
	d.parseErr = err
	// Spans are rebuilt lazily: a document that is never rendered never gets flattened.
	d.spans, d.built = d.spans[:0], false
}

func (d *document) build() {
	d.built = true
	d.builder.reset()
	d.hl.spans(d.src, func(rs rawSpan) { d.builder.add(d.src, rs) })
	d.spans = d.builder.build(d.spans)
	d.Debugf("built %d spans for %d %s lines", len(d.spans), len(d.src.lines), d.lang.name())
}

// SyntaxSpans implements editor.SyntaxHighlighter.
func (d *document) SyntaxSpans(ln int, line string) []editor.SyntaxSpan {
	d.ensureParsed()
	if !d.built {
		d.build()
	}
	return clipSpans(spansOfLine(d.spans, ln), line)
}

// CodeBlocks implements code.Blocks for the languages embedding code blocks, it's empty for others.
// The document is brought up to date with the buffer first. So, another extension can use the
// blocks in its AfterEdit regardless of whether this extension has been notified about the edit yet.
func (d *document) CodeBlocks() []code.Block {
	blocks, ok := d.hl.(code.Blocks)
	if !ok {
		return nil
	}
	// The edit notification may not have reached this extension yet.
	d.parse(d.buf.Text())
	if !d.built {
		d.build()
	}
	return blocks.CodeBlocks()
}

func (d *document) Close() error { return d.hl.Close() }
