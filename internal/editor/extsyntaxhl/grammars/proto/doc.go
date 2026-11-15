// Package proto provides the tree-sitter grammar for Protocol Buffers.
//
// The sources under src are copied from https://github.com/coder3101/tree-sitter-proto
// by fetchgrammar, which moves the parse tables to tables.zlib to keep them
// compressed in the binary. Upstream records the commit. To update them,
// change the reference below, run go generate and check that the highlight
// query still compiles.
package proto

//go:generate go run ../fetchgrammar https://github.com/coder3101/tree-sitter-proto 0.6.0
