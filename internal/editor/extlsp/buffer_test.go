package extlsp_test

import (
	"testing"

	"rmazur.io/x/edit/internal/editor/extlsp"
)

func TestBufferData(t *testing.T) {
	var d extlsp.BufferData
	d.Assign([]string{"s1", "s2", "s3"})
	if s := d.CurrentSuggestion(); s != "s1" {
		t.Fatal("bad first suggestion", s)
	}

	d.SuggestPrev()
	t.Log("suggestPrev")
	if s := d.CurrentSuggestion(); s != "s3" {
		t.Fatal("unexpected current suggestion", s)
	}

	d.SuggestPrev()
	t.Log("suggestPrev")
	if s := d.CurrentSuggestion(); s != "s2" {
		t.Fatal("unexpected current suggestion", s)
	}

	d.SuggestNext()
	t.Log("suggestNext")
	if s := d.CurrentSuggestion(); s != "s3" {
		t.Fatal("unexpected current suggestion", s)
	}

	d.SuggestNext()
	t.Log("suggestNext")
	if s := d.CurrentSuggestion(); s != "s1" {
		t.Fatal("unexpected current suggestion", s)
	}
}
