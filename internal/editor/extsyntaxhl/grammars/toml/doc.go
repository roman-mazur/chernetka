// Package toml provides the tree-sitter grammar for TOML.
//
// The sources under src are copied from https://github.com/tree-sitter-grammars/tree-sitter-toml
// by fetchgrammar, which moves the parse tables to tables.zlib to keep them
// compressed in the binary. Upstream records the commit. To update them,
// change the reference below, run go generate and check that the highlight
// query still compiles.
package toml

//go:generate go run ../fetchgrammar https://github.com/tree-sitter-grammars/tree-sitter-toml v0.7.0
