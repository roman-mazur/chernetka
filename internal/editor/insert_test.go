package editor

import (
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
)

func TestInsertInput_HandleCursor(t *testing.T) {
	data, err := content.LoadFullText(strings.NewReader("line 1\nline 2"))
	if err != nil {
		t.Fatal(err)
	}
	buf := Buffer{
		Content: &data,
	}
	changed := insertInput(&buf, []byte{0x1b, '[', 'C'}, &RenderPrefs{TabSize: 2})
	if changed {
		t.Errorf("no changes expected after handling cursor keys")
	}
	if buf.cx != 1 {
		t.Errorf("cx didn't change after arrow right")
	}
}
