package tree_sitter

/*
#cgo CFLAGS: -Iinclude -Isrc -std=c11 -D_POSIX_C_SOURCE=200112L -D_DEFAULT_SOURCE
#include <tree_sitter/api.h>
*/
import "C"

import "unsafe"

// ParseBuffer parses the UTF8 text like Parse does, reading it from the slice directly.
//
// Parse gives the core a callback reading the text, which copies the rest of the text from the
// requested offset to a new C string on every call, and keeps them all till the parse ends.
// The core reads the text again wherever it goes back, which it does a lot in the error recovery
// and while reusing the old tree, so the copies take most of the time of a parse.
// The core does not keep the text after parsing it, and neither does the tree.
//
// It's not part of the upstream bindings: see ownFiles in fetchcore.
func (p *Parser) ParseBuffer(text []byte, oldTree *Tree) *Tree {
	var cOldTree *C.TSTree
	if oldTree != nil {
		cOldTree = oldTree._inner
	}
	var cText *C.char
	if len(text) > 0 {
		cText = (*C.char)(unsafe.Pointer(&text[0]))
	}
	cNewTree := C.ts_parser_parse_string(p._inner, cOldTree, cText, C.uint32_t(len(text)))
	if cNewTree == nil {
		return nil
	}
	return newTree(cNewTree)
}
