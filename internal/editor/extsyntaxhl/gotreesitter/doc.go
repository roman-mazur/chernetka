// Package tree_sitter is a copy of github.com/tree-sitter/go-tree-sitter with
// a newer core library. The latest release of the bindings ships the core
// v0.25.0, which may loop forever in the error recovery, like the CUE grammar
// does on "{R(z&[". The core v0.25.3 fixes it (tree-sitter/tree-sitter#4257).
//
// The core API is the same in the v0.25 patch releases, so the bindings work
// with it unchanged. A replace directive in go.mod is not an option: go install
// refuses the modules having one. Drop the copy once the bindings release
// a newer core. To update it, change the revisions below and run go generate.
// The bindings are v0.25.0: its tag is removed upstream, the commit is the one
// the Go module proxy has for it.
package tree_sitter

//go:generate go run ./fetchcore adc13ffd8b2c0b01b878fda9f7c422ce0df5fad3 v0.25.10
