// Package nix provides the tree-sitter grammar for Nix.
//
// The sources under src are copied from https://github.com/nix-community/tree-sitter-nix
// by fetchgrammar, which moves the parse tables to tables.zlib to keep them
// compressed in the binary. Upstream records the commit. To update them,
// change the reference below, run go generate and check that the highlight
// query still compiles.
package nix

//go:generate go run ../fetchgrammar https://github.com/nix-community/tree-sitter-nix a2cd7f4011c6e5830c0c9af5aa35441b3ddd5fba
