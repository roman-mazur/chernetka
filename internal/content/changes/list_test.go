package changes

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"rmazur.io/chernetka/internal/content"
)

func TestHistory_Track(t *testing.T) {
	var h History

	m := h.Track(content.Empty().(content.Mutable))
	m.Insert(0, content.TextLine("line 1"))
	if len(h) != 1 {
		t.Fatal("history didn't start on fist edit")
	}

	m.Insert(1, content.TextLine("line 2"))
	m.Update(1, content.TextLine("update line 2"))
	m.Delete(0)
	m.Update(0, content.TextLine("update line 2 v2"))
	m.Update(0, content.TextLine("update line 2 v3")) // modifies the last change

	if editsLen := len(h[0].Items); editsLen != 5 {
		t.Errorf("expected history of 5 items, got %d", editsLen)
	}

	var out bytes.Buffer
	if err := content.SaveToWriter(m, &out); err != nil {
		t.Fatal(err)
	}

	var replayOut bytes.Buffer
	if err := content.SaveToWriter(h[0].Replay(), &replayOut); err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(out.String(), replayOut.String()); diff != "" {
		t.Errorf("replay mismatch (-want +got):\n%s", diff)
	}
}
