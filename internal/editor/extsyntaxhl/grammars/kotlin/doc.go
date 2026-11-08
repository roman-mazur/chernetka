// Package kotlin provides the tree-sitter grammar for Kotlin.
//
// The sources under src are copied from https://github.com/fwcd/tree-sitter-kotlin
// by fetchgrammar, which moves the parse tables to tables.zlib to keep them
// compressed in the binary. Upstream records the commit. To update them,
// change the reference below, run go generate and check that the highlight
// query still compiles.
package kotlin

//go:generate go run ../fetchgrammar https://github.com/fwcd/tree-sitter-kotlin 0.3.8
