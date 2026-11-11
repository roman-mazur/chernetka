// Package html provides the tree-sitter grammar for HTML.
//
// The sources under src are copied from https://github.com/tree-sitter/tree-sitter-html
// by fetchgrammar, which moves the parse tables to tables.zlib to keep them
// compressed in the binary. Upstream records the commit. To update them,
// change the reference below, run go generate and check that the highlight
// query still compiles.
package html

//go:generate go run ../fetchgrammar https://github.com/tree-sitter/tree-sitter-html v0.23.2
