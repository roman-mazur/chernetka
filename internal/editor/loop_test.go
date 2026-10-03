package editor

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
)

func TestEditor_RenderToolBuffer(t *testing.T) {
	var e Editor
	mainLines := make(content.FullText, 100)
	for i := range mainLines {
		mainLines[i] = content.TextLine(fmt.Sprintf("main %d", i+1))
	}
	e.OpenBuffer(&Buffer{Path: "a.txt", Content: &mainLines})
	toolLines := content.FullText{content.TextLine("tool 1"), content.TextLine("tool 2"), content.TextLine("tool 3")}
	e.toolBuf = &Buffer{Content: &toolLines}

	var out bytes.Buffer
	e.render(bufio.NewWriter(&out))
	rows := screenRows(out.String(), 40)
	t.Log("\n" + strings.Join(rows, "\n"))

	// The screen is 80x40 in tests: 7 rows of the tool buffer at the bottom,
	// the status bar of the main buffer right above it, and the main content on top.
	const toolH = 7
	statusRow := len(rows) - toolH - 1
	for i := range statusRow {
		if want := fmt.Sprintf("main %d", i+1); !strings.Contains(rows[i], want) {
			t.Errorf("row %d = %q, want %q", i, rows[i], want)
		}
	}
	if !strings.HasPrefix(rows[statusRow], " NORMAL  a.txt") {
		t.Errorf("status row %d = %q", statusRow, rows[statusRow])
	}
	for i, row := range rows[statusRow+1:] {
		want := ""
		if i < len(toolLines) {
			want = toolLines[i].String()
		}
		if got := strings.TrimSpace(row); !strings.HasSuffix(got, want) || (want == "" && got != "") {
			t.Errorf("tool row %d = %q, want %q", i, row, want)
		}
	}
}

var (
	cursorPositionRE = regexp.MustCompile(`\x1b\[(?:(\d+);(\d+))?H`)
	csiRE            = regexp.MustCompile(`\x1b\[[0-9;?]* ?[A-Za-z]`)
	oscRE            = regexp.MustCompile(`\x1b\][^\x07]*\x07`)
)

// screenRows emulates a terminal of height rows writing the output to it and
// returns the text of every row. It follows the cursor positioning and the line
// endings, the other escape sequences are dropped.
func screenRows(output string, height int) []string {
	screen := make([][]rune, height)
	row, col := 0, 0
	write := func(text string) {
		text = csiRE.ReplaceAllString(text, "")
		for len(text) > 0 {
			if strings.HasPrefix(text, "\r\n") {
				row, col = row+1, 0
				text = text[2:]
				if row == height {
					// A line ending on the last row scrolls the screen.
					screen = append(screen[1:], nil)
					row--
				}
				continue
			}
			r := []rune(text)[0]
			text = text[len(string(r)):]
			for len(screen[row]) <= col {
				screen[row] = append(screen[row], ' ')
			}
			screen[row][col] = r
			col++
		}
	}
	// The terminal title and the mouse shape are not shown on the screen.
	output = oscRE.ReplaceAllString(output, "")
	last := 0
	for _, m := range cursorPositionRE.FindAllStringSubmatchIndex(output, -1) {
		write(output[last:m[0]])
		last = m[1]
		row, col = 0, 0
		if m[2] >= 0 {
			row, _ = strconv.Atoi(output[m[2]:m[3]])
			col, _ = strconv.Atoi(output[m[4]:m[5]])
			row, col = row-1, col-1
		}
	}
	write(output[last:])

	res := make([]string, height)
	for i := range screen {
		res[i] = strings.TrimRight(string(screen[i]), " ")
	}
	return res
}

func BenchmarkLayoutState_Pass(b *testing.B) {
	ls := layoutState{editor: setupBenchmarkEditor(b)}
	for b.Loop() {
		for range ls.Pass() {
		}
	}
}

func setupBenchmarkEditor(b *testing.B) *Editor {
	b.Helper()
	f, err := os.Open("loop.go")
	if err != nil {
		b.Fatal(err)
	}
	b.Cleanup(func() { _ = f.Close() })

	var edit Editor
	if err := edit.OpenReader("test.txt", f); err != nil {
		b.Fatal(err)
	}
	return &edit
}
