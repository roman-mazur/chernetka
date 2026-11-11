// Package hcl provides the tree-sitter grammar for HCL.
//
// The sources under src are copied from https://github.com/tree-sitter-grammars/tree-sitter-hcl
// by fetchgrammar, which moves the parse tables to tables.zlib to keep them
// compressed in the binary. Upstream records the commit. To update them,
// change the reference below, run go generate and check that the highlight
// query still compiles.
package hcl

//go:generate go run ../fetchgrammar https://github.com/tree-sitter-grammars/tree-sitter-hcl v1.2.0
