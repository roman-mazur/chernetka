package editor

import (
	"os"
	"testing"
)

func BenchmarkLayoutState_Pass(b *testing.B) {
	ls := layoutState{editor: setupBenchmarkEditor(b)}
	for b.Loop() {
		for range ls.Pass() {
		}
	}
}

func setupBenchmarkEditor(b *testing.B) *Editor {
	b.Helper()
	f, err := os.Open("loop.go")
	if err != nil {
		b.Fatal(err)
	}
	b.Cleanup(func() { _ = f.Close() })

	var edit Editor
	if err := edit.OpenReader("test.txt", f); err != nil {
		b.Fatal(err)
	}
	return &edit
}
