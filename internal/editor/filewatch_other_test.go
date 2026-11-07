//go:build !unix

package editor

import "testing"

// limitOpenFiles skips the test: the number of open files cannot be limited on this platform.
func limitOpenFiles(t *testing.T, _ uint64) {
	t.Skip("the number of open files cannot be limited on this platform")
}
