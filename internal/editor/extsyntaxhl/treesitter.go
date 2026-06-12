package extsyntaxhl

import (
	"path/filepath"

	treesitter "github.com/tree-sitter/go-tree-sitter"
	gositter "github.com/tree-sitter/tree-sitter-go/bindings/go"
	"rmazur.io/chernetka/internal/editor"
)

// Integration implements an editor.Extension by integrating with the tree-sitter for syntax highlight.
type Integration struct {
}

func (in *Integration) ID() string { return "syntaxhl" }

func (in *Integration) MakeBufferData(buf *editor.Buffer) editor.BufferExtData {
	if buf.Path == "" {
		return nil
	}
	if filepath.Ext(buf.Path) != ".go" {
		return nil
	}

	p := treesitter.NewParser()
	defer p.Close()
	if err := p.SetLanguage(treesitter.NewLanguage(gositter.Language())); err != nil {
		return nil
	}
	tree := p.Parse([]byte(buf.Text()), nil)
	return &syntaxTree{
		tree: tree,
	}
}

func (in *Integration) AfterEdit(*editor.Editor, *editor.Buffer) {
	// TODO: update the syntax tree.
}

func (in *Integration) HandleInsertInput(*editor.Buffer, *editor.RenderPrefs, []byte) (handled, changed bool) {
	return
}

type syntaxTree struct {
	tree     *treesitter.Tree
	lastNode *treesitter.Node
}

func (st *syntaxTree) Close() error {
	st.tree.Close()
	return nil
}

func (st *syntaxTree) SyntaxSpans(ln int, line string) []editor.SyntaxSpan {
	//log.Println(ln, line, " <<<<<<<<<<<<<<<<")
	node := st.locateFirstNode(ln)
	if node == nil {
		// Didn't find anything.
		return noHighlightSpan(ln, line)
	}

	toTraverse := []*treesitter.Node{node}
	p := node.Parent()
	if p != nil {
		for i := range p.ChildCount() {
			sibling := p.Child(i)
			if sibling.StartPosition() != node.StartPosition() && sibling.StartPosition().Row == uint(ln) {
				toTraverse = append(toTraverse, sibling)
			}
		}
	}

	var out []editor.SyntaxSpan
	for _, n := range toTraverse {
		collectSpans(n, &out, ln, line)
	}
	if len(out) == 0 {
		return noHighlightSpan(ln, line)
	}
	return out
}

func noHighlightSpan(ln int, line string) []editor.SyntaxSpan {
	return []editor.SyntaxSpan{{LineNumber: ln, End: len(line)}}
}

func appendSpan(out *[]editor.SyntaxSpan, span editor.SyntaxSpan, line string) {
	if span.Start >= len(line) {
		return
	}
	*out = append(*out, span)
}

func collectSpans(node *treesitter.Node, out *[]editor.SyntaxSpan, ln int, line string) {
	if node.StartPosition().Row > uint(ln) {
		return
	}
	//log.Print(node.StartPosition())
	//log.Print(node.EndPosition())
	//log.Print(node.GrammarName())

	switch node.GrammarName() {
	// See https://go.dev/ref/spec#Keywords
	case "break", "default", "func", "interface", "select",
		"case", "defer", "go", "map", "struct",
		"chan", "else", "goto", "package", "switch",
		"const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var":
		appendSpan(out, spanFromNode(node, editor.TtKeyword), line)
		return

	case "identifier":
		appendSpan(out, spanFromNode(node, editor.TtIdentifier), line)
		return

	case "interpreted_string_literal":
		appendSpan(out, spanFromNode(node, editor.TtStringLiteral), line)
		return
	}

	for i := range node.ChildCount() {
		collectSpans(node.Child(i), out, ln, line)
	}
}

func spanFromNode(node *treesitter.Node, typ editor.TokenType) editor.SyntaxSpan {
	return editor.SyntaxSpan{
		LineNumber: int(node.StartPosition().Row),
		Start:      int(node.StartPosition().Column),
		End:        int(node.EndPosition().Column),
		TokenType:  typ,
	}
}

func (st *syntaxTree) locateFirstNode(ln int) *treesitter.Node {
	if st.lastNode == nil {
		st.lastNode = st.tree.RootNode()
	}
	if st.lastNode.StartPosition().Row > uint(ln) {
		st.lastNode = st.tree.RootNode()
	}

	stack := []*treesitter.Node{st.lastNode}
	var node *treesitter.Node
	for len(stack) > 0 {
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if node.StartPosition().Row == uint(ln) && node.EndPosition().Row == uint(ln) {
			break
		}

		if node.StartPosition().Row > uint(ln) || node.EndPosition().Row < uint(ln) {
			if p := node.Parent(); p != nil {
				stack = append(stack, p)
			}
			continue
		}

		n := node.ChildCount()
		for i := n - 1; i >= 0 && i < n; i-- {
			ch := node.Child(i)
			if ch.StartPosition().Row <= uint(ln) && ch.EndPosition().Row >= uint(ln) {
				stack = append(stack, ch)
			}
		}
	}

	st.lastNode = node
	if node == nil || node.StartPosition().Row != uint(ln) {
		return nil
	}

	//log.Println("next node:")
	//log.Println(node.StartPosition())
	//log.Println(node.GrammarName())
	return node
}
