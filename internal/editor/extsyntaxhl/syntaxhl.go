// Package extsyntaxhl implements an editor extension that colorizes buffers.
//
// A buffer is matched to a language by its file name, and the language provides
// a highlighter: either a tree-sitter grammar driven by a highlight query, or a
// hand written scanner for the formats that have no Go grammar bindings.
// Whatever the source, a highlighter reports coarse tokens that the
// spanBuilder flattens into the per-line spans the renderer consumes.
package extsyntaxhl

import (
	"io"
	"iter"

	"rmazur.io/chernetka/internal/content/change"
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/content/parser/git"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/bash"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/c"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/commonlisp"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/cpp"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/csharp"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/css"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/cue"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/golang"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/hcl"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/html"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/java"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/json"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/kotlin"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/makefile"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/nix"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/proto"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/python"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/rust"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/toml"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/tsquery"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/tsx"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/typescript"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/xml"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/yaml"
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
	reparse(src *source, edits []textEdit) error
	// spans returns the highlighted regions of the document, which may span
	// lines and nest; the caller flattens them. It's iterated once per revision.
	spans(src *source) iter.Seq[code.Token]

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

// The grammars shared by several languages, or prepared at the start, see init.
var (
	goGrammar   = newGrammar(golang.Language, "golang")
	jsonGrammar = newGrammar(json.Language, "json")
	// TSX extends TypeScript, and so does its query.
	tsxGrammar = newGrammar(tsx.Language, "tsx", "typescript")
)

// languages are the highlighted languages. A language highlighted with a tree-sitter
// grammar takes a package of the grammars directory, which holds the grammar and
// its highlight query, highlights.scm.
var languages = map[*code.Syntax]*language{
	code.Go:              {newHighlighter: newTreeSitter(goGrammar)},
	code.JSON:            {newHighlighter: newTreeSitter(jsonGrammar)},
	code.JSONC:           {newHighlighter: newTreeSitter(jsonGrammar)},
	code.TSX:             {newHighlighter: newTreeSitter(tsxGrammar)},
	code.JavaScript:      {newHighlighter: newTreeSitter(tsxGrammar)},
	code.Nix:             treeSitter(nix.Language, "nix"),
	code.CUE:             treeSitter(cue.Language, "cue"),
	code.YAML:            treeSitter(yaml.Language, "yaml"),
	code.Shell:           treeSitter(bash.Language, "bash"),
	code.C:               treeSitter(c.Language, "c"),
	code.CPP:             treeSitter(cpp.Language, "cpp", "c"), // C++ extends C
	code.CSharp:          treeSitter(csharp.Language, "csharp"),
	code.Java:            treeSitter(java.Language, "java"),
	code.Kotlin:          treeSitter(kotlin.Language, "kotlin"),
	code.TypeScript:      treeSitter(typescript.Language, "typescript"),
	code.Rust:            treeSitter(rust.Language, "rust"),
	code.Python:          treeSitter(python.Language, "python"),
	code.TOML:            treeSitter(toml.Language, "toml"),
	code.HCL:             treeSitter(hcl.Language, "hcl"),
	code.HTML:            treeSitter(html.Language, "html"),
	code.CSS:             treeSitter(css.Language, "css"),
	code.XML:             treeSitter(xml.Language, "xml"),
	code.Proto:           treeSitter(proto.Language, "proto"),
	code.Makefile:        treeSitter(makefile.Language, "makefile"),
	code.CommonLisp:      treeSitter(commonlisp.Language, "commonlisp"),
	code.TreeSitterQuery: treeSitter(tsquery.Language, "tsquery"),
	code.Markdown:        {newHighlighter: newMarkdown},
	code.Dockerfile:      {newHighlighter: newDockerfile},
	code.GitCommitMsg:    {newHighlighter: newGitHighlighter(git.ScanCommitMessage)},
	code.GitRebaseTodo:   {newHighlighter: newGitHighlighter(git.ScanRebaseTodo)},
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

// HighlightsSyntax implements editor.SyntaxHighlighterExtension.
func (in *Integration) HighlightsSyntax() {}

func (in *Integration) MakeBufferData(_ editor.Sender, buf *editor.Buffer) editor.BufferExtData {
	lang := languageForPath(buf.Path)
	// A directory listing buffer is named after its directory, which may well end in
	// a recognized extension, and colorizing file names as code would be wrong.
	if lang == nil || !buf.HoldsTextFile() {
		return nil
	}

	in.Logf("highlighting %s as %s", buf.Path, lang.name())
	// The document is parsed when it's rendered the first time.
	return &document{LogEmbed: &in.LogEmbed, buf: buf, lang: lang, hl: lang.newHighlighter()}
}

func (in *Integration) AfterEdit(_ editor.Sender, buf *editor.Buffer, changes []change.Item) {
	if doc, ok := buf.ExtensionData(in.ID()).(*document); ok {
		doc.edited(changes)
	}
}

// document is the per-buffer extension data. It owns the language highlighter
// and the flattened spans of the whole document.
//
// The document is parsed when its spans or code blocks are requested, not on every edit:
// the keys typed between two renders are parsed at once, with all their changes.
// The text is not kept between the parses: the changes are located in the text
// with the lengths of the lines parsed the last time.
type document struct {
	*logger.LogEmbed
	buf  *editor.Buffer
	lang *language
	hl   highlighter

	lineLens []int // the lengths of the lines parsed the last time, nil until the first parse
	parseErr error // of the last parse
	parses   int   // counts the parses

	// stale is set when the buffer is edited after the last parse. The pending changes turn
	// the parsed revision into the edited one, unless changesLost is set.
	stale       bool
	pending     []change.Item
	changesLost bool

	builder spanBuilder
	spans   []editor.SyntaxSpan // sorted by line number, then by start offset
}

// edited records the changes made to the buffer, the document is parsed again when
// the highlights are requested.
func (d *document) edited(changes []change.Item) {
	d.stale = true
	switch {
	case d.lineLens == nil:
		// Not parsed yet, the first parse is from scratch.
	case changes == nil:
		// Not known, like when the content is reloaded.
		d.dropPending()
		d.changesLost = true
	case d.changesLost:
	case len(d.pending)+len(changes) > maxIncrementalChanges:
		// So many changes are parsed from scratch anyway.
		d.dropPending()
		d.changesLost = true
	default:
		d.pending = append(d.pending, changes...)
	}
}

func (d *document) dropPending() {
	clear(d.pending) // Don't hold the lines.
	d.pending, d.changesLost = d.pending[:0], false
}

// ensureParsed brings the document up to date with the edits it was notified about. The changes are parsed
// incrementally if they are all known, the text is parsed from scratch otherwise.
func (d *document) ensureParsed() {
	if d.lineLens != nil && !d.stale {
		return
	}
	var changes []change.Item
	if d.lineLens != nil && !d.changesLost {
		changes = d.pending
		if changes == nil {
			changes = []change.Item{} // Known, but none.
		}
	}
	d.parse(changes)
}

// parse brings the document up to date with the buffer, and builds its spans.
// The changes, unless nil, are the ones made to the revision parsed the last time,
// they let the highlighter parse incrementally.
func (d *document) parse(changes []change.Item) {
	src := newSource(d.buf.Text())
	var edits []textEdit
	incremental := false
	if changes != nil && d.parseErr == nil {
		if edits, incremental = textEdits(d.lineLens, src, changes); !incremental {
			d.Debugf("%d changes don't match the last parsed revision of %s, parsing from scratch", len(changes), d.buf.Path)
		}
	}

	d.stale = false
	d.dropPending()
	if d.lineLens == nil {
		d.lineLens = make([]int, 0, len(src.lines))
	}
	d.lineLens = d.lineLens[:0]
	for _, line := range src.lines {
		d.lineLens = append(d.lineLens, len(line))
	}
	if incremental && len(edits) == 0 {
		return // The text is the same.
	}

	d.parses++
	err := d.hl.reparse(src, edits)
	if err != nil && d.parseErr == nil {
		// Logged once: the document may stay broken for many edits.
		d.Logf("cannot highlight %s: %s", d.buf.Path, err)
	}
	d.parseErr = err

	d.builder.reset()
	for token := range d.hl.spans(src) {
		d.builder.add(src, token)
	}
	d.spans = d.builder.build(d.spans[:0])
	d.Debugf("built %d spans for %d %s lines", len(d.spans), len(src.lines), d.lang.name())
}

// SyntaxSpans implements editor.SyntaxHighlighter.
func (d *document) SyntaxSpans(ln int, line string) []editor.SyntaxSpan {
	d.ensureParsed()
	return clipSpans(spansOfLine(d.spans, ln), line)
}

// CodeBlocks implements code.Blocks for the languages embedding code blocks, it's empty for others.
// The document is brought up to date with the edits it was notified about first: as a syntax
// highlighter, it's notified before the other extensions, which can use the blocks in AfterEdit.
func (d *document) CodeBlocks() []code.Block {
	blocks, ok := d.hl.(code.Blocks)
	if !ok {
		return nil
	}
	d.ensureParsed()
	return blocks.CodeBlocks()
}

func (d *document) Close() error { return d.hl.Close() }
