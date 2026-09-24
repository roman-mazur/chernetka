package logger

import (
	"fmt"
	"testing"
)

func TestWriter(t *testing.T) {
	w := Writer(Prefix(t.Logf, "test: "))
	const text = "hello world"
	n, err := fmt.Fprint(w, text)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(text) {
		t.Errorf("got %d, want %d", n, len(text))
	}
}
