package tree_sitter

/*
#cgo CFLAGS: -Iinclude -Isrc -std=c11 -D_POSIX_C_SOURCE=200112L -D_DEFAULT_SOURCE
#include <tree_sitter/api.h>
*/
import "C"

// UseLibcAllocator makes the core library allocate with the C library directly.
//
// The bindings route every allocation of the core through Go, even with the
// default allocator set by SetAllocator(nil, nil, nil, nil): each one is a cgo
// callback that ends up in the same C malloc. Parsing allocates a lot, and the
// callbacks take about a third of a parse. Since the memory comes from the same
// malloc either way, the switch is safe at any time, but it's not synchronized
// with the core: call it before parsing, and not along with SetAllocator.
//
// It's not part of the upstream bindings: see ownFiles in fetchcore.
func UseLibcAllocator() {
	C.ts_set_allocator(nil, nil, nil, nil)
}
