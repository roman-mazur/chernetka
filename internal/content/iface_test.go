package content

import (
	"fmt"
	"testing"
)

func TestIsText(t *testing.T) {
	for _, tc := range []struct {
		name string
		doc  Document
		text bool
	}{
		{name: "text", doc: &FullText{TextLine("test")}, text: true},
		{name: "folder", doc: LoadFolder(".", nil), text: false},
		{name: "empty", doc: Empty(), text: true},
		{name: "error", doc: &ErrorContent{Error: fmt.Errorf("test")}, text: false},
	} {
		if res := IsText(tc.doc); res != tc.text {
			t.Errorf("IsText() = %t, want %t", res, tc.text)
		}
	}
}
