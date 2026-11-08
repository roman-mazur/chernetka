// Package cpp provides the tree-sitter grammar for C++.
//
// The sources under src are copied from https://github.com/tree-sitter/tree-sitter-cpp
// by fetchgrammar, which moves the parse tables to tables.zlib to keep them
// compressed in the binary. Upstream records the commit. To update them,
// change the reference below, run go generate and check that the highlight
// query still compiles.
package cpp

//go:generate go run ../fetchgrammar https://github.com/tree-sitter/tree-sitter-cpp v0.23.4
