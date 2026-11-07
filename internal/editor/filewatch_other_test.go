//go:build !unix

package editor

import "testing"

// limitOpenFiles does nothing: the number of open files cannot be limited on this platform.
func limitOpenFiles(*testing.T, uint64) {}
