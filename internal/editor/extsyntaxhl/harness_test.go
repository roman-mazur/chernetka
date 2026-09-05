package extsyntaxhl

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
)

// hlLine is one line of a test document together with the spans expected for
// it, each written as "start:end:TokenType". A nil tokens list means the line
// is expected to carry no highlight at all.
type hlLine struct {
	line   string
	tokens []string
}

type hlDoc []hlLine

func (d hlDoc) text() string {
	lines := make([]string, len(d))
	for i, l := range d {
		lines[i] = l.line
	}
	return strings.Join(lines, "\n")
}

// check compares the spans reported for every line against the expectations.
func (d hlDoc) check(t *testing.T, hl editor.SyntaxHighlighter) {
	t.Helper()
	for i, expected := range d {
		var got []string
		for _, s := range hl.SyntaxSpans(i, expected.line) {
			if s.LineNumber != i {
				t.Errorf("line %d: span %s reported for the wrong line", i, s)
			}
			if s.Start < 0 || s.End > len(expected.line) || s.Start >= s.End {
				t.Errorf("line %d %q: span %s is out of range", i, expected.line, s)
			}
			got = append(got, fmt.Sprintf("%d:%d:%s", s.Start, s.End, s.TokenType))
		}
		if fmt.Sprint(got) != fmt.Sprint(expected.tokens) {
			t.Errorf("line %d %q:\n got %v\nwant %v", i, expected.line, got, expected.tokens)
		}
	}
}

// openDoc opens a buffer through a real editor, so the test covers the
// extension wiring as well as the highlighter itself.
func openDoc(t *testing.T, path, text string) (*editor.Buffer, *Integration) {
	t.Helper()

	ext := new(Integration)
	ext.LogDebug = true
	ext.LogF = t.Logf

	var edit editor.Editor
	edit.Extend(ext)
	if err := edit.OpenReader(path, strings.NewReader(text)); err != nil {
		t.Fatal(err)
	}
	buf := edit.Top()
	t.Cleanup(func() { _ = buf.Close() })
	return buf, ext
}

// highlighterOf returns the buffer's highlighter, failing the test if the
// extension decided not to highlight the buffer at all.
func highlighterOf(t *testing.T, buf *editor.Buffer, ext *Integration) editor.SyntaxHighlighter {
	t.Helper()
	hl, ok := buf.ExtensionData(ext.ID()).(editor.SyntaxHighlighter)
	if !ok {
		t.Fatalf("no syntax highlighter for %q", buf.Path)
	}
	return hl
}

// insertLine inserts an empty line in the middle of the buffer and lets the
// extension react, which is what invalidates the cached spans.
func insertLine(t *testing.T, buf *editor.Buffer, ext *Integration, at int) {
	t.Helper()
	buf.Mutate().Insert(at, content.TextLine(""))
	ext.AfterEdit(nil, buf)
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}
