package markdown

import (
	"strings"
	"testing"
	"time"

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

func TestScan_LongDelimiterRun(t *testing.T) {
	// A run of stars longer than three, or of tildes other than two, is not emphasis,
	// and is scanned once. When each of its suffixes is scanned again, a run as long
	// as this one takes many seconds.
	const n = 300_000
	for _, c := range []string{"*", "~"} {
		line := "a " + strings.Repeat(c, n) + " b" + c
		start := time.Now()
		Scan([]string{line}, func(s Span) {
			if s.Token == code.TtEmphasis || s.Token == code.TtStrong || s.Token == code.TtComment {
				t.Errorf("%s: emphasis at %d-%d", c, s.Start, s.End)
			}
		})
		if d := time.Since(start); d > time.Second {
			t.Errorf("scanning a run of %d %s takes %v", n, c, d)
		}
	}
}

func TestScan_UnclosedLinear(t *testing.T) {
	// The openers with no closer on the line are scanned in linear time, and so are
	// the nested constructs. Each of these lines takes many seconds when the rest of
	// the line is scanned again from each opener, or for each nesting level.
	const n = 100_000
	r := strings.Repeat
	for name, line := range map[string]string{
		"emphasis":      r("*a ", n),
		"strong":        r("**a ", n),
		"underscore":    r(" _a", n),
		"strikethrough": r("~~a ", n),
		"code span":     "a " + r("`", 3*n),
		"link":          r("[", 3*n),
		"image":         r("![", n),
		"nested links":  r("[", n) + r("]", n),
		"link target":   r("[a](", n),
		"link ref":      r("[a][", n),
		"link label":    r("[*a ", n) + r("]", n),
		"angle":         r("<", 3*n),
		"block quote":   r("> ", n) + "a",
	} {
		painted := 0
		start := time.Now()
		Scan([]string{line}, func(s Span) { painted += s.End - s.Start })
		if d := time.Since(start); d > time.Second {
			t.Errorf("%s: scanning %d bytes takes %v", name, len(line), d)
		}
		// The spans are painted for the editor, a nested span over its parent.
		if painted > 4*len(line) {
			t.Errorf("%s: the spans of %d bytes cover %d", name, len(line), painted)
		}
	}
}
