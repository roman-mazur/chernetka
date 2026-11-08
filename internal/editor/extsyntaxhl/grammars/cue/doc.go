// Package cue provides the tree-sitter grammar for CUE.
//
// The sources under src are copied from https://github.com/eonpatapon/tree-sitter-cue
// by fetchgrammar, which moves the parse tables to tables.zlib to keep them
// compressed in the binary. Upstream records the commit. To update them,
// change the reference below, run go generate and check that the highlight
// query still compiles.
package cue

//go:generate go run ../fetchgrammar https://github.com/eonpatapon/tree-sitter-cue dd7b90e0770ff18070c515937ba3c3d6d93db00e
