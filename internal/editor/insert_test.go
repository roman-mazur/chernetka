package editor

import (
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/input"
)

func TestInsertInput_HandleCursor(t *testing.T) {
	data, err := content.LoadFullText(strings.NewReader("line 1\nline 2"))
	if err != nil {
		t.Fatal(err)
	}
	buf := Buffer{
		Content: &data,
	}
	insertInput(&buf, input.Move(input.CursorArrowRight, 0), &RenderPrefs{TabSize: 2})
	if buf.resetMutated() {
		t.Error("unexpected mutation")
	}
	if buf.c().Col != 1 {
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
		_c:      content.Position{Col: data.Lines()[0].Len()}, // at the end of the first line
	}
	prefs := RenderPrefs{TabSize: 2}

	const expectedBrackets = "{(['`\"\"`'])}"

	for i := range expectedBrackets[:len(expectedBrackets)/2] {
		insertInput(&buf, input.Rune(rune(expectedBrackets[i])), &prefs)
	}

	if !buf.resetMutated() {
		t.Error("mutations expected but didn't seem to happen")
	}

	res, expected := buf.Content.Lines()[0].String(), firstLineContent+expectedBrackets
	if res != expected {
		t.Errorf("got %q, want %q", res, expected)
	}
	expectedCurPos := content.Position{Col: len(firstLineContent) + len(expectedBrackets)/2}
	if buf.c() != expectedCurPos {
		t.Errorf("cursor=%v, want %v", buf.c(), expectedCurPos)
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
		insertInput(&buf, input.Rune(rune(pair[0])), &prefs)
		insertInput(&buf, input.Rune(rune(pair[1])), &prefs)
		if !buf.resetMutated() {
			t.Error("mutations expected but didn't seem to happen")
		}
		if buf.Text() != expected {
			t.Errorf("i=%d, buf.Text()=%q, want %q", i, buf.Text(), pair)
		}
	}

	t.Log("regression check: insert at 0 pos")
	MoveHome.DoOnBuffer(&buf, prefs)
	insertInput(&buf, input.Rune('}'), &prefs)
	MoveHome.DoOnBuffer(&buf, prefs)
	insertInput(&buf, input.Rune('}'), &prefs)
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
		insertInput(&buf, input.Rune(sym), &prefs)
	}

	if !buf.resetMutated() {
		t.Error("mutations expected but didn't seem to happen")
	}

	if res := buf.Text(); res != sample {
		t.Errorf("buf.Text()=%q, want %q", res, sample)
	}
	expectedPos := content.Position{Col: len(sample)}
	if buf.c() != expectedPos {
		t.Errorf("cursor=%v, want %v", buf.c(), expectedPos)
	}
}

func TestInsertInput_ReplaceMultilineSelection(t *testing.T) {
	for _, tc := range []struct {
		name string
		key  input.Key
		want string
	}{
		{"backspace", input.Of(input.Backspace), "first 12\nlast"},
		{"enter", input.Of(input.Enter), "first 12\n\nlast"},
		{"rune", input.Rune('x'), "first 12x\nlast"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			data, err := content.LoadFullText(strings.NewReader("first 12345\nmiddle\nab\nlast"))
			if err != nil {
				t.Fatal(err)
			}
			// The selection ends on a line shorter than its start column.
			end := content.Position{Line: 2, Col: 2}
			buf := Buffer{
				Content: &data,
				_c:      end,
				sel:     []content.Span{{Start: content.Position{Col: 8}, End: end}},
			}
			insertInput(&buf, tc.key, &RenderPrefs{TabSize: 2})
			if got := buf.Text(); got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
