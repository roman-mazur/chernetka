// Package typescript provides the tree-sitter grammar for TypeScript.
//
// The sources under typescript/src are copied from the typescript grammar of
// https://github.com/tree-sitter/tree-sitter-typescript by fetchgrammar,
// which moves the parse tables to tables.zlib to keep them compressed in the
// binary. Upstream records the commit. To update them, change the reference
// below, run go generate and check that the highlight query still compiles.
package typescript

//go:generate go run ../fetchgrammar https://github.com/tree-sitter/tree-sitter-typescript v0.23.2 typescript
