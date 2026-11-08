// Package yaml provides the tree-sitter grammar for YAML.
//
// The sources under src are copied from https://github.com/tree-sitter-grammars/tree-sitter-yaml
// by fetchgrammar, which moves the parse tables to tables.zlib to keep them
// compressed in the binary. Upstream records the commit. To update them,
// change the reference below, run go generate and check that the highlight
// query still compiles.
package yaml

//go:generate go run ../fetchgrammar https://github.com/tree-sitter-grammars/tree-sitter-yaml v0.7.2
