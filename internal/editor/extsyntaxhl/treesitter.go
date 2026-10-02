package extsyntaxhl

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"

	yamlsitter "github.com/tree-sitter-grammars/tree-sitter-yaml/bindings/go"
	treesitter "rmazur.io/chernetka/internal/editor/extsyntaxhl/gotreesitter"
	bashsitter "github.com/tree-sitter/tree-sitter-bash/bindings/go"
	gositter "github.com/tree-sitter/tree-sitter-go/bindings/go"
	jsonsitter "github.com/tree-sitter/tree-sitter-json/bindings/go"
	"rmazur.io/chernetka/internal/content/code"
	cuesitter "rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/cue"
	nixsitter "rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/nix"
)

func init() {
	go goGrammar.prepare()
}

//go:embed queries/go.scm
var goQuery string

// goGrammar highlights Go using the upstream tree-sitter grammar.
var goGrammar = &tsGrammar{
	load:   func() *treesitter.Language { return treesitter.NewLanguage(gositter.Language()) },
	query:  goQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/json.scm
var jsonQuery string

// jsonGrammar highlights JSON using the upstream tree-sitter grammar.
var jsonGrammar = &tsGrammar{
	load:   func() *treesitter.Language { return treesitter.NewLanguage(jsonsitter.Language()) },
	query:  jsonQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/yaml.scm
var yamlQuery string

// yamlGrammar highlights YAML using the upstream tree-sitter grammar.
var yamlGrammar = &tsGrammar{
	load:   func() *treesitter.Language { return treesitter.NewLanguage(yamlsitter.Language()) },
	query:  yamlQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/bash.scm
var bashQuery string

// bashGrammar highlights shell scripts using the upstream tree-sitter-bash grammar.
var bashGrammar = &tsGrammar{
	load:   func() *treesitter.Language { return treesitter.NewLanguage(bashsitter.Language()) },
	query:  bashQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/nix.scm
var nixQuery string

// nixGrammar highlights Nix using the grammar copied from tree-sitter-nix.
var nixGrammar = &tsGrammar{
	load:   func() *treesitter.Language { return treesitter.NewLanguage(nixsitter.Language()) },
	query:  nixQuery,
	tokens: defaultCaptureTokens,
}

//go:embed queries/cue.scm
var cueQuery string

// cueGrammar highlights CUE using the grammar copied from tree-sitter-cue.
var cueGrammar = &tsGrammar{
	load:   func() *treesitter.Language { return treesitter.NewLanguage(cuesitter.Language()) },
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
	// Code embedded in a string, like an interpolated expression, is plain
	// text again rather than a part of the string.
	"embedded": code.TtIdentifier,
}

// tsGrammar is a tree-sitter grammar paired with its highlight query. Both are
// immutable and shared by every buffer of the language, so they are built once.
type tsGrammar struct {
	load   func() *treesitter.Language
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
// over a full tree-sitter parse.
type tsHighlighter struct {
	grammar *tsGrammar
	tree    *treesitter.Tree
}

func newTreeSitter(g *tsGrammar) func() highlighter {
	return func() highlighter { return &tsHighlighter{grammar: g} }
}

func (h *tsHighlighter) reparse(src *source) error {
	h.closeTree()
	if err := h.grammar.prepare(); err != nil {
		return err
	}

	parser := treesitter.NewParser()
	defer parser.Close()
	if err := parser.SetLanguage(h.grammar.lang); err != nil {
		return err
	}
	// TODO: feed the edited ranges to Parse to reparse incrementally.
	h.tree = parser.Parse([]byte(src.text), nil)
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
