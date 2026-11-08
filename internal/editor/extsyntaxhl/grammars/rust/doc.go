// Package rust provides the tree-sitter grammar for Rust.
//
// The sources under src are copied from https://github.com/tree-sitter/tree-sitter-rust
// by fetchgrammar, which moves the parse tables to tables.zlib to keep them
// compressed in the binary. Upstream records the commit. To update them,
// change the reference below, run go generate and check that the highlight
// query still compiles.
package rust

//go:generate go run ../fetchgrammar https://github.com/tree-sitter/tree-sitter-rust v0.24.2
