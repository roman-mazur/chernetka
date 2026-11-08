package extsyntaxhl

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"
	"unsafe"

	"rmazur.io/chernetka/internal/content/code"
	treesitter "rmazur.io/chernetka/internal/editor/extsyntaxhl/gotreesitter"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/bash"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/c"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/cpp"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/csharp"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/cue"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/golang"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/java"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/json"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/kotlin"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/nix"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/yaml"
)

func init() {
	// Before the grammar is prepared: the allocator is not switched safely while the core runs.
	treesitter.UseLibcAllocator()
	go goGrammar.prepare()
}

//go:embed queries/go.scm
var goQuery string

// goGrammar highlights Go using the tree-sitter grammar copied from upstream.
var goGrammar = &tsGrammar{
	load:   golang.Language,
	query:  goQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/json.scm
var jsonQuery string

// jsonGrammar highlights JSON using the tree-sitter grammar copied from upstream.
var jsonGrammar = &tsGrammar{
	load:   json.Language,
	query:  jsonQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/yaml.scm
var yamlQuery string

// yamlGrammar highlights YAML using the tree-sitter grammar copied from upstream.
var yamlGrammar = &tsGrammar{
	load:   yaml.Language,
	query:  yamlQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/bash.scm
var bashQuery string

// bashGrammar highlights shell scripts using the tree-sitter-bash grammar copied from upstream.
var bashGrammar = &tsGrammar{
	load:   bash.Language,
	query:  bashQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/c.scm
var cQuery string

// cGrammar highlights C using the tree-sitter grammar copied from upstream.
var cGrammar = &tsGrammar{
	load:   c.Language,
	query:  cQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/cpp.scm
var cppQuery string

// cppGrammar highlights C++ using the tree-sitter grammar copied from upstream. Its grammar
// extends the C one, and so does its query: the C patterns follow the C++ ones.
var cppGrammar = &tsGrammar{
	load:   cpp.Language,
	query:  cppQuery + "\n" + cQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/csharp.scm
var csharpQuery string

// csharpGrammar highlights C# using the tree-sitter grammar copied from upstream.
var csharpGrammar = &tsGrammar{
	load:   csharp.Language,
	query:  csharpQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/java.scm
var javaQuery string

// javaGrammar highlights Java using the tree-sitter grammar copied from upstream.
var javaGrammar = &tsGrammar{
	load:   java.Language,
	query:  javaQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/kotlin.scm
var kotlinQuery string

// kotlinGrammar highlights Kotlin using the grammar copied from fwcd/tree-sitter-kotlin.
var kotlinGrammar = &tsGrammar{
	load:   kotlin.Language,
	query:  kotlinQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/nix.scm
var nixQuery string

// nixGrammar highlights Nix using the grammar copied from tree-sitter-nix.
var nixGrammar = &tsGrammar{
	load:   nix.Language,
	query:  nixQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/cue.scm
var cueQuery string

// cueGrammar highlights CUE using the grammar copied from tree-sitter-cue.
var cueGrammar = &tsGrammar{
	load:   cue.Language,
	query:  cueQuery,
	tokens: defaultCaptureTokens,
}

// defaultCaptureTokens maps the capture names used by tree-sitter highlight
// queries to editor token types. A dotted name falls back to its prefix, so
// "function.call" without an entry of its own is treated as "function".
var defaultCaptureTokens = map[string]code.TokenType{
	"keyword":  code.TtKeyword,
	"variable": code.TtIdentifier,
	"type":     code.TtTypeRef,
	"module":   code.TtImportRef,
	"function": code.TtFuncDeclaration,
	// A call site is a reference, not a declaration.
	"function.call":   code.TtCall,
	"function.method": code.TtCall,
	"property":        code.TtField,
	"string":          code.TtStringLiteral,
	"escape":          code.TtEscape,
	"number":          code.TtNumberLiteral,
	"constant":        code.TtConstant,
	"comment":         code.TtComment,
	"annotation":      code.TtAnnotation,
	"directive":       code.TtDirective,
	// Code embedded in a string, like an interpolated expression, is plain
	// text again rather than a part of the string.
	"embedded": code.TtIdentifier,
}

// tsGrammar is a tree-sitter grammar paired with its highlight query. Both are
// immutable and shared by every buffer of the language, so they are built once.
type tsGrammar struct {
	load   func() (unsafe.Pointer, error)
	query  string
	tokens map[string]code.TokenType

	once     sync.Once
	lang     *treesitter.Language
	compiled *treesitter.Query
	// captureTokens resolves a capture index to its token type, which avoids a
	// map lookup for every capture of every parse.
	captureTokens []code.TokenType
	err           error
}

// prepare compiles the grammar and its query on first use. The error is sticky:
// a query that fails to compile will not compile later either.
func (g *tsGrammar) prepare() error {
	g.once.Do(func() {
		ptr, err := g.load()
		if err != nil {
			g.err = err
			return
		}
		g.lang = treesitter.NewLanguage(ptr)
		// NewQuery returns a concrete *QueryError rather than an error, so it
		// has to be nil checked before being treated as one.
		compiled, queryErr := treesitter.NewQuery(g.lang, g.query)
		if queryErr != nil {
			g.err = fmt.Errorf("highlight query: %s", queryErr.Error())
			return
		}
		g.compiled = compiled

		names := compiled.CaptureNames()
		g.captureTokens = make([]code.TokenType, len(names))
		for i, name := range names {
			g.captureTokens[i] = tokenForCapture(g.tokens, name)
		}
	})
	return g.err
}

func (g *tsGrammar) tokenType(captureIndex uint32) code.TokenType {
	if int(captureIndex) >= len(g.captureTokens) {
		return code.TtNothing
	}
	return g.captureTokens[captureIndex]
}

// tokenForCapture resolves a capture name against the mapping, falling back to
// ever less specific names. Names with no mapping at all yield TtNothing and
// their captures are dropped, so a query may capture more than we colorize.
func tokenForCapture(tokens map[string]code.TokenType, name string) code.TokenType {
	for {
		if t, ok := tokens[name]; ok {
			return t
		}
		dot := strings.LastIndexByte(name, '.')
		if dot < 0 {
			return code.TtNothing
		}
		name = name[:dot]
	}
}

// tsHighlighter highlights a document by running the grammar's highlight query
// over a tree-sitter parse. The tree of the previous revision is reused for
// parsing the next one when the edits between them are known.
type tsHighlighter struct {
	grammar *tsGrammar
	parser  *treesitter.Parser
	tree    *treesitter.Tree
}

func newTreeSitter(g *tsGrammar) func() highlighter {
	return func() highlighter { return &tsHighlighter{grammar: g} }
}

func (h *tsHighlighter) reparse(src *source, edits []treesitter.InputEdit) error {
	if err := h.grammar.prepare(); err != nil {
		h.closeTree()
		return err
	}
	if h.parser == nil {
		parser := treesitter.NewParser()
		if err := parser.SetLanguage(h.grammar.lang); err != nil {
			parser.Close()
			h.closeTree()
			return err
		}
		h.parser = parser
	}

	// The old tree is reused only if it's edited to match the new text:
	// the parser takes it as the parse of the new text otherwise.
	old := h.tree
	if old != nil && len(edits) > 0 {
		for i := range edits {
			old.Edit(&edits[i])
		}
	} else {
		h.closeTree()
		old = nil
	}
	tree := h.parser.ParseBuffer([]byte(src.text), old)
	h.closeTree()
	h.tree = tree
	return nil
}

func (h *tsHighlighter) spans(src *source, emit func(rawSpan)) {
	if h.tree == nil {
		return
	}
	root := h.tree.RootNode()
	if root == nil {
		return
	}

	cursor := treesitter.NewQueryCursor()
	defer cursor.Close()

	captures := cursor.Captures(h.grammar.compiled, root, []byte(src.text))
	for {
		match, i := captures.Next()
		if match == nil {
			return
		}
		if int(i) >= len(match.Captures) {
			continue
		}
		capture := match.Captures[i]
		token := h.grammar.tokenType(capture.Index)
		if token == code.TtNothing {
			continue
		}
		// Next reuses the match memory, so the positions are read out now.
		emit(spanFromNode(&capture.Node, token))
	}
}

func (h *tsHighlighter) Close() error {
	h.closeTree()
	if h.parser != nil {
		h.parser.Close()
		h.parser = nil
	}
	return nil
}

func (h *tsHighlighter) closeTree() {
	if h.tree != nil {
		h.tree.Close()
		h.tree = nil
	}
}

// spanFromNode converts a node's position to a rawSpan. A tree-sitter column is
// a byte offset within its row, which is what the editor slices lines by. Nodes
// covering several rows stay multi-line here and are cut up by the spanBuilder.
func spanFromNode(node *treesitter.Node, token code.TokenType) rawSpan {
	start, end := node.StartPosition(), node.EndPosition()
	return rawSpan{
		StartLine: int(start.Row),
		StartCol:  int(start.Column),
		EndLine:   int(end.Row),
		EndCol:    int(end.Column),
		TokenType: token,
	}
}
