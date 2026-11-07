//go:build unix

package editor

import (
	"syscall"
	"testing"
)

// limitOpenFiles lowers the number of files the test process may open to n until the test ends.
func limitOpenFiles(t *testing.T, n uint64) {
	var prev syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &prev); err != nil {
		t.Fatal(err)
	}
	limit := prev
	limit.Cur = min(n, prev.Cur)
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &prev); err != nil {
			t.Error(err)
		}
	})
}
