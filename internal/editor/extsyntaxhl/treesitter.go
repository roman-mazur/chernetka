package extsyntaxhl

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"

	treesitter "github.com/tree-sitter/go-tree-sitter"
	gositter "github.com/tree-sitter/tree-sitter-go/bindings/go"
	"rmazur.io/chernetka/internal/editor"
)

//go:embed queries/go.scm
var goQuery string

// goGrammar highlights Go using the upstream tree-sitter grammar.
var goGrammar = &tsGrammar{
	load:   func() *treesitter.Language { return treesitter.NewLanguage(gositter.Language()) },
	query:  goQuery,
	tokens: defaultCaptureTokens,
}

// defaultCaptureTokens maps the capture names used by tree-sitter highlight
// queries to editor token types. A dotted name falls back to its prefix, so
// "function.call" without an entry of its own is treated as "function".
var defaultCaptureTokens = map[string]editor.TokenType{
	"keyword":  editor.TtKeyword,
	"variable": editor.TtIdentifier,
	"type":     editor.TtTypeRef,
	"module":   editor.TtImportRef,
	"function": editor.TtFuncDeclaration,
	// A call site is a reference, not a declaration.
	"function.call":   editor.TtCall,
	"function.method": editor.TtCall,
	"property":        editor.TtField,
	"string":          editor.TtStringLiteral,
	"escape":          editor.TtEscape,
	"number":          editor.TtNumberLiteral,
	"constant":        editor.TtConstant,
	"comment":         editor.TtComment,
}

// tsGrammar is a tree-sitter grammar paired with its highlight query. Both are
// immutable and shared by every buffer of the language, so they are built once.
type tsGrammar struct {
	load   func() *treesitter.Language
	query  string
	tokens map[string]editor.TokenType

	once     sync.Once
	lang     *treesitter.Language
	compiled *treesitter.Query
	// captureTokens resolves a capture index to its token type, which avoids a
	// map lookup for every capture of every parse.
	captureTokens []editor.TokenType
	err           error
}

// prepare compiles the grammar and its query on first use. The error is sticky:
// a query that fails to compile will not compile later either.
func (g *tsGrammar) prepare() error {
	g.once.Do(func() {
		g.lang = g.load()
		if g.lang == nil {
			g.err = fmt.Errorf("grammar is not available")
			return
		}
		// NewQuery returns a concrete *QueryError rather than an error, so it
		// has to be nil checked before being treated as one.
		compiled, queryErr := treesitter.NewQuery(g.lang, g.query)
		if queryErr != nil {
			g.err = fmt.Errorf("highlight query: %s", queryErr.Error())
			return
		}
		g.compiled = compiled

		names := compiled.CaptureNames()
		g.captureTokens = make([]editor.TokenType, len(names))
		for i, name := range names {
			g.captureTokens[i] = tokenForCapture(g.tokens, name)
		}
	})
	return g.err
}

func (g *tsGrammar) tokenType(captureIndex uint32) editor.TokenType {
	if int(captureIndex) >= len(g.captureTokens) {
		return editor.TtNothing
	}
	return g.captureTokens[captureIndex]
}

// tokenForCapture resolves a capture name against the mapping, falling back to
// ever less specific names. Names with no mapping at all yield TtNothing and
// their captures are dropped, so a query may capture more than we colorize.
func tokenForCapture(tokens map[string]editor.TokenType, name string) editor.TokenType {
	for {
		if t, ok := tokens[name]; ok {
			return t
		}
		dot := strings.LastIndexByte(name, '.')
		if dot < 0 {
			return editor.TtNothing
		}
		name = name[:dot]
	}
}

// tsHighlighter highlights a document by running the grammar's highlight query
// over a full tree-sitter parse.
type tsHighlighter struct {
	grammar *tsGrammar
	tree    *treesitter.Tree
}

func newTreeSitter(g *tsGrammar) func() highlighter {
	return func() highlighter { return &tsHighlighter{grammar: g} }
}

func (h *tsHighlighter) reparse(src *source) {
	h.closeTree()
	if err := h.grammar.prepare(); err != nil {
		return
	}

	parser := treesitter.NewParser()
	defer parser.Close()
	if err := parser.SetLanguage(h.grammar.lang); err != nil {
		return
	}
	// TODO: feed the edited ranges to Parse to reparse incrementally.
	h.tree = parser.Parse([]byte(src.text), nil)
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
		if token == editor.TtNothing {
			continue
		}
		// Next reuses the match memory, so the positions are read out now.
		emit(spanFromNode(&capture.Node, token))
	}
}

func (h *tsHighlighter) Close() error {
	h.closeTree()
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
func spanFromNode(node *treesitter.Node, token editor.TokenType) rawSpan {
	start, end := node.StartPosition(), node.EndPosition()
	return rawSpan{
		StartLine: int(start.Row),
		StartCol:  int(start.Column),
		EndLine:   int(end.Row),
		EndCol:    int(end.Column),
		TokenType: token,
	}
}
