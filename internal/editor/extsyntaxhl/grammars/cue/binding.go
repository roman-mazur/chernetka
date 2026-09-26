// Package cue provides the tree-sitter grammar for CUE.
//
// The sources under src are copied from https://github.com/eonpatapon/tree-sitter-cue,
// as the upstream repository has no Go bindings. Upstream records the commit.
// To update them, run go generate and check that the highlight query still
// compiles.
package cue

//go:generate go run ../fetchgrammar https://github.com/eonpatapon/tree-sitter-cue main

// #cgo CFLAGS: -std=c11 -fPIC -I${SRCDIR}/src
// #include "src/parser.c"
// #include "src/scanner.c"
import "C"

import "unsafe"

// Language returns the tree-sitter language, to be passed to treesitter.NewLanguage.
func Language() unsafe.Pointer {
	return unsafe.Pointer(C.tree_sitter_cue())
}
