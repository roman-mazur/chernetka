// Package tsx provides the tree-sitter grammar for TypeScript with JSX, which
// also reads JavaScript, with or without JSX.
//
// The sources under tsx/src are copied from the tsx grammar of
// https://github.com/tree-sitter/tree-sitter-typescript by fetchgrammar,
// which moves the parse tables to tables.zlib to keep them compressed in the
// binary. Upstream records the commit. To update them, change the reference
// below, run go generate and check that the highlight query still compiles.
package tsx

//go:generate go run ../fetchgrammar https://github.com/tree-sitter/tree-sitter-typescript v0.23.2 tsx
