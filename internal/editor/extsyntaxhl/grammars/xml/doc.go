// Package xml provides the tree-sitter grammar for XML, which also reads SVG.
//
// The sources under xml/src are copied from the xml grammar of
// https://github.com/tree-sitter-grammars/tree-sitter-xml by fetchgrammar,
// which moves the parse tables to tables.zlib to keep them compressed in the
// binary. Upstream records the commit. To update them, change the reference
// below, run go generate and check that the highlight query still compiles.
package xml

//go:generate go run ../fetchgrammar https://github.com/tree-sitter-grammars/tree-sitter-xml v0.7.0 xml
