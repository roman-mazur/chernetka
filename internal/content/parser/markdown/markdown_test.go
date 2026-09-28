package markdown

import (
	"testing"

	"rmazur.io/chernetka/internal/content/code"
)

func TestTaskCheckbox(t *testing.T) {
	for line, want := range map[string]int{
		"- [ ] task":      3,
		"- [x] task":      3,
		"* [X] task":      3,
		"+ [ ]":           3,
		"  - [ ] nested":  5,
		"\t-  [ ] spaces": 5,
		"-\t[ ] tab":      3,
		"1. [ ] numbered": 4,
		"12) [x] paren":   5,
		"> - [ ] quoted":  5,
		"> > 1. [x] deep": 8,

		"":                -1,
		"   ":             -1,
		">":               -1,
		"- task":          -1,
		"-[ ] task":       -1,
		"- [] task":       -1,
		"- [y] task":      -1,
		"- [ ]task":       -1,
		"- [ ](link)":     -1,
		"[ ] task":        -1,
		"1 [ ] task":      -1,
		". [ ] task":      -1,
		"1234567890. [ ]": -1,
		"text - [ ] no":   -1,
	} {
		if got := TaskCheckbox(line); got != want {
			t.Errorf("TaskCheckbox(%q) = %d, want %d", line, got, want)
		}
	}
}

func TestScan_TaskCheckboxConsistent(t *testing.T) {
	// The scanner marks the checkbox of a task list item where TaskCheckbox finds it.
	lines := []string{"- [ ] task", "> 2. [x] quoted", "- [ ](link)", "-\t[ ] tab"}
	var boxes []int
	Scan(lines, func(s Span) {
		if s.Token == code.TtConstant {
			boxes = append(boxes, s.Line)
			if want := TaskCheckbox(lines[s.Line]); s.Start+1 != want {
				t.Errorf("checkbox of %q at %d, TaskCheckbox = %d", lines[s.Line], s.Start+1, want)
			}
		}
	})
	if len(boxes) != 3 || boxes[0] != 0 || boxes[1] != 1 || boxes[2] != 3 {
		t.Errorf("checkboxes on lines %v, want [0 1 3]", boxes)
	}
}
