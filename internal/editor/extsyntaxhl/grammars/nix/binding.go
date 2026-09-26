// Package nix provides the tree-sitter grammar for Nix.
//
// The sources under src are copied from https://github.com/nix-community/tree-sitter-nix,
// as the upstream repository has no Go bindings. Upstream records the commit.
// To update them, run go generate and check that the highlight query still
// compiles.
package nix

//go:generate go run ../fetchgrammar https://github.com/nix-community/tree-sitter-nix master

// #cgo CFLAGS: -std=c11 -fPIC -I${SRCDIR}/src
// #include "src/parser.c"
// #include "src/scanner.c"
import "C"

import "unsafe"

// Language returns the tree-sitter language, to be passed to treesitter.NewLanguage.
func Language() unsafe.Pointer {
	return unsafe.Pointer(C.tree_sitter_nix())
}
