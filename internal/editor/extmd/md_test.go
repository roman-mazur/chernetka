package extmd

import (
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
)

func TestCheckboxAt(t *testing.T) {
	for line, want := range map[string]int{
		"- [ ] task":      3,
		"- [x] task":      3,
		"* [X] task":      3,
		"+ [ ]":           3,
		"  - [ ] nested":  5,
		"\t-  [ ] spaces": 5,
		"1. [ ] numbered": 4,
		"12) [x] paren":   5,

		"":              -1,
		"- task":        -1,
		"-[ ] task":     -1,
		"- [] task":     -1,
		"- [y] task":    -1,
		"- [ ]task":     -1,
		"[ ] task":      -1,
		"1 [ ] task":    -1,
		". [ ] task":    -1,
		"- [ ](link)":   -1,
		"text - [ ] no": -1,
	} {
		if got := checkboxAt(line); got != want {
			t.Errorf("checkboxAt(%q) = %d, want %d", line, got, want)
		}
	}
}

func openDoc(t *testing.T, path, text string) (*editor.Buffer, content.LineActions) {
	t.Helper()
	var edit editor.Editor
	ext := new(Integration)
	edit.Extend(ext)
	if err := edit.OpenReader(path, strings.NewReader(text)); err != nil {
		t.Fatal(err)
	}
	buf := edit.Top()
	actions, _ := buf.ExtensionData(ext.ID()).(content.LineActions)
	return buf, actions
}

func TestToggle(t *testing.T) {
	buf, actions := openDoc(t, "todo.md", "# TODO\n- [ ] one\n  - [x] two\n")
	if actions == nil {
		t.Fatal("no line actions")
	}
	if actions.LineAction(0) != nil || actions.LineAction(3) != nil {
		t.Error("action on a line without a checkbox")
	}

	actions.LineAction(1).Engage()
	actions.LineAction(2).Engage()
	if got, want := buf.Text(), "# TODO\n- [x] one\n  - [ ] two\n"; got != want {
		t.Errorf("text after toggling = %q, want %q", got, want)
	}
	if _, ok := actions.LineAction(1).(content.LineActionSingleShot); !ok {
		t.Error("toggle is re-run as the last action")
	}
}

func TestNotMarkdown(t *testing.T) {
	if _, actions := openDoc(t, "todo.txt", "- [ ] one\n"); actions != nil {
		t.Error("line actions for a text file")
	}
}
