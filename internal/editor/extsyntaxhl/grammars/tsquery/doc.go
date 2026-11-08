// Package tsquery provides the tree-sitter grammar for tree-sitter queries, the
// language of the highlight queries in ../../queries.
//
// The sources under src are copied from https://github.com/tree-sitter-grammars/tree-sitter-query
// by fetchgrammar, which moves the parse tables to tables.zlib to keep them
// compressed in the binary. Upstream records the commit. To update them,
// change the reference below, run go generate and check that the highlight
// query still compiles.
package tsquery

//go:generate go run ../fetchgrammar https://github.com/tree-sitter-grammars/tree-sitter-query v0.8.0
