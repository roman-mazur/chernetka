// Package csharp provides the tree-sitter grammar for C#.
//
// The sources under src are copied from https://github.com/tree-sitter/tree-sitter-c-sharp
// by fetchgrammar, which moves the parse tables to tables.zlib to keep them
// compressed in the binary. Upstream records the commit. To update them,
// change the reference below, run go generate and check that the highlight
// query still compiles.
package csharp

//go:generate go run ../fetchgrammar https://github.com/tree-sitter/tree-sitter-c-sharp v0.23.5
