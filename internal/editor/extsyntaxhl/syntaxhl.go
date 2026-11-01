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
	code.Go:            {newHighlighter: newTreeSitter(goGrammar)},
	code.Nix:           {newHighlighter: newTreeSitter(nixGrammar)},
	code.CUE:           {newHighlighter: newTreeSitter(cueGrammar)},
	code.JSON:          {newHighlighter: newTreeSitter(jsonGrammar)},
	code.JSONC:         {newHighlighter: newTreeSitter(jsonGrammar)},
	code.YAML:          {newHighlighter: newTreeSitter(yamlGrammar)},
	code.Shell:         {newHighlighter: newTreeSitter(bashGrammar)},
	code.Markdown:      {newHighlighter: newMarkdown},
	code.GitCommitMsg:  {newHighlighter: func() highlighter { return new(gitMessage) }},
	code.GitRebaseTodo: {newHighlighter: func() highlighter { return new(gitRebase) }},
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
	doc.notified = buf.Version()
	doc.parse(nil)
	return doc
}

func (in *Integration) AfterEdit(_ editor.Sender, buf *editor.Buffer, changes []change.Item) {
	if doc, ok := buf.ExtensionData(in.ID()).(*document); ok {
		doc.edited(changes)
	}
}

// document is the per-buffer extension data. It owns the language highlighter
// and caches the flattened spans of the whole document until the next edit.
//
// The document is parsed when its spans or code blocks are requested, not on every edit:
// the keys typed between two renders are parsed at once, with the changes they made.
// It keeps no text of its own: the buffer's text is parsed, and the changes locate
// the edits in it with the lengths of the lines of the revision parsed the last time.
type document struct {
	*logger.LogEmbed
	buf  *editor.Buffer
	lang *language
	hl   highlighter

	parsed   uint64 // the buffer version parsed the last time
	lens     []int  // the lengths of the lines of the parsed version
	parseErr error  // of the last parse

	// pending are the changes made since the parsed version, up to the notified one,
	// unless changesLost is set: then the next parse is from scratch.
	pending     []change.Item
	notified    uint64 // the buffer version of the last AfterEdit
	changesLost bool

	builder spanBuilder
	spans   []editor.SyntaxSpan // sorted by line number, then by start offset
	built   bool
}

// edited records the changes made to the buffer, the document is parsed again when
// the highlights are requested.
func (d *document) edited(changes []change.Item) {
	d.notified = d.buf.Version()
	if d.changesLost {
		if d.notified == d.parsed {
			// The changes were parsed with the code blocks requested before the notification.
			d.changesLost = false
		}
		return
	}
	if len(d.pending)+len(changes) > maxIncrementalChanges {
		// So many changes are parsed from scratch anyway.
		clear(d.pending)
		d.pending, d.changesLost = d.pending[:0], true
		return
	}
	d.pending = append(d.pending, changes...)
}

// ensureParsed brings the document up to date with the buffer.
func (d *document) ensureParsed() {
	switch {
	case d.buf.Version() == d.parsed:
		return
	case d.buf.Version() != d.notified:
		// The code blocks are requested before the edit notification reached this extension,
		// so not all the changes are known. The ones coming with it are parsed already.
		d.parse(nil)
		d.changesLost = true
	case d.changesLost:
		d.parse(nil)
	default:
		d.parse(d.pending)
	}
}

// parse brings the document up to date with the buffer. The changes, if known, are the ones
// made since the version parsed the last time, they let the highlighter parse incrementally.
// The pending changes are dropped.
func (d *document) parse(changes []change.Item) {
	lines := d.buf.Content.Lines()
	var edits []treesitter.InputEdit
	if changes != nil && d.parseErr == nil {
		var ok bool
		if edits, ok = textEdits(d.lens, changes, lines); !ok {
			d.Debugf("%d changes don't match the last parsed version of %s, parsing from scratch", len(changes), d.buf.Path)
		}
	}
	err := d.hl.reparse(newSource(d.buf.Text()), edits)
	if err != nil && d.parseErr == nil {
		// Logged once: the document may stay broken for many edits.
		d.Logf("cannot highlight %s: %s", d.buf.Path, err)
	}
	d.parseErr = err
	d.parsed, d.lens = d.buf.Version(), lineLengths(lines)
	clear(d.pending) // Don't hold the lines.
	d.pending, d.changesLost = d.pending[:0], false
	// Spans are rebuilt lazily: a document that is never rendered never gets flattened.
	d.spans, d.built = d.spans[:0], false
}

// build flattens the spans of the parsed version, which must be the buffer's.
func (d *document) build() {
	d.built = true
	src := newSource(d.buf.Text())
	d.builder.reset()
	d.hl.spans(src, func(rs rawSpan) { d.builder.add(src, rs) })
	d.spans = d.builder.build(d.spans)
	d.Debugf("built %d spans for %d %s lines", len(d.spans), len(src.lines), d.lang.name())
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
	d.ensureParsed()
	if !d.built {
		d.build()
	}
	return blocks.CodeBlocks()
}

func (d *document) Close() error { return d.hl.Close() }
