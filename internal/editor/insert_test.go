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
	insertInput(&buf, []byte{0x1b, '[', 'C'}, &RenderPrefs{TabSize: 2})
	if buf.resetMutated() {
		t.Error("unexpected mutation")
	}
	if buf.c.Col != 1 {
		t.Errorf("cx didn't change after arrow right")
	}
}

func TestInsertInput_AutomateBrackets(t *testing.T) {
	const firstLineContent = "line 1"
	data, err := content.LoadFullText(strings.NewReader(
		strings.Join([]string{firstLineContent, "something"}, "\n"),
	))
	if err != nil {
		t.Fatal(err)
	}
	buf := Buffer{
		Content: &data,
		c:       content.Position{Col: data.Lines()[0].Len()}, // at the end of the first line
	}
	prefs := RenderPrefs{TabSize: 2}

	const expectedBrackets = "{(['`\"\"`'])}"

	for i := range expectedBrackets[:len(expectedBrackets)/2] {
		insertInput(&buf, []byte{expectedBrackets[i]}, &prefs)
	}

	if !buf.resetMutated() {
		t.Error("mutations expected but didn't seem to happen")
	}

	res, expected := buf.Content.Lines()[0].String(), firstLineContent+expectedBrackets
	if res != expected {
		t.Errorf("got %q, want %q", res, expected)
	}
	expectedCurPos := content.Position{Col: len(firstLineContent) + len(expectedBrackets)/2}
	if buf.c != expectedCurPos {
		t.Errorf("buf.c=%v, want %v", buf.c, expectedCurPos)
	}
}

func TestInsertInput_HandleAutoClosingBracket(t *testing.T) {
	var buf Buffer
	buf.Content = content.Empty()
	prefs := RenderPrefs{TabSize: 2}

	var expected string
	for i, pair := range []string{
		"{}", "()", "[]", "''", "``", "\"\"",
	} {
		expected += pair
		insertInput(&buf, []byte{pair[0]}, &prefs)
		insertInput(&buf, []byte{pair[1]}, &prefs)
		if !buf.resetMutated() {
			t.Error("mutations expected but didn't seem to happen")
		}
		if buf.Text() != expected {
			t.Errorf("i=%d, buf.Text()=%q, want %q", i, buf.Text(), pair)
		}
	}

	t.Log("regression check: insert at 0 pos")
	MoveHome.DoOnBuffer(&buf, prefs)
	insertInput(&buf, []byte{'}'}, &prefs)
	MoveHome.DoOnBuffer(&buf, prefs)
	insertInput(&buf, []byte{'}'}, &prefs)
	if buf.Text() != "}}"+expected {
		t.Errorf("buf.Text()=%q, want %q", buf.Text(), "}}"+expected)
	}
}

func TestInsertInput_AcceptUTF8(t *testing.T) {
	buf := Buffer{
		Content: content.Empty(),
	}
	prefs := RenderPrefs{TabSize: 2}

	const sample = "кохання вічне"
	for _, sym := range sample {
		insertInput(&buf, []byte(string(sym)), &prefs)
	}

	if !buf.resetMutated() {
		t.Error("mutations expected but didn't seem to happen")
	}

	if res := buf.Text(); res != sample {
		t.Errorf("buf.Text()=%q, want %q", res, sample)
	}
	expectedPos := content.Position{Col: len(sample)}
	if buf.c != expectedPos {
		t.Errorf("buf.c=%v, want %v", buf.c, expectedPos)
	}
}
