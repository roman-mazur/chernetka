package extlsp_test

import (
	"testing"

	"rmazur.io/chernetka/internal/editor/extlsp"
)

func TestBufferData(t *testing.T) {
	var d extlsp.BufferData
	d.Assign([]string{"s1", "s2", "s3"})
	if s := d.TextSuggestion().Text; s != "s1" {
		t.Fatal("bad first suggestion", s)
	}

	d.SuggestPrev()
	t.Log("suggestPrev")
	if s := d.TextSuggestion().Text; s != "s3" {
		t.Fatal("unexpected current suggestion", s)
	}

	d.SuggestPrev()
	t.Log("suggestPrev")
	if s := d.TextSuggestion().Text; s != "s2" {
		t.Fatal("unexpected current suggestion", s)
	}

	d.SuggestNext()
	t.Log("suggestNext")
	if s := d.TextSuggestion().Text; s != "s3" {
		t.Fatal("unexpected current suggestion", s)
	}

	d.SuggestNext()
	t.Log("suggestNext")
	if s := d.TextSuggestion().Text; s != "s1" {
		t.Fatal("unexpected current suggestion", s)
	}

	// Checking suggestion info.
	d.Assign(nil)
	if got := d.TextSuggestion().Info; got != "" {
		t.Errorf("info without suggestions = %q", got)
	}
	d.Assign([]string{"s1"})
	if got := d.TextSuggestion().Info; got != "" {
		t.Errorf("info for a single plain suggestion = %q", got)
	}
	d.Assign([]string{"s1", "s2", "s3"})
	d.SuggestNext()
	if got, want := d.TextSuggestion().Info, "↑↓ 2/3"; got != want {
		t.Errorf("info = %q, want %q", got, want)
	}
}
