// Package c provides the tree-sitter grammar for C.
//
// The sources under src are copied from https://github.com/tree-sitter/tree-sitter-c
// by fetchgrammar, which moves the parse tables to tables.zlib to keep them
// compressed in the binary. Upstream records the commit. To update them,
// change the reference below, run go generate and check that the highlight
// query still compiles.
package c

//go:generate go run ../fetchgrammar https://github.com/tree-sitter/tree-sitter-c v0.24.2
