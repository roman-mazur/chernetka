package content

import (
	"bufio"
	"strings"
	"testing"
)

func TestLoadFolder(t *testing.T) {
	fc := LoadFolder("testdata/dirs/sample1")
	expectedStructure := []struct {
		name string
		dir  bool
	}{
		{"child-empty", true},
		{"child1", true},
		{"child2", true},
		{"fulltext-example-3.txt", false},
	}

	if len(expectedStructure) != fc.Len() {
		t.Fatalf("length mismatch: expected %d, got %d", len(expectedStructure), fc.Len())
	}

	lines := fc.Lines()
	for i, entry := range expectedStructure {
		if !strings.Contains(lines[i].String(), entry.name) {
			t.Errorf("line %d: expected to contain %q, got %q", i, entry.name, lines[i])
		}
		if mimeType := lines[i].MimeType(); mimeType != "text/filename" {
			t.Errorf("lines[%d].MimeType() = %q, want %q", i, mimeType, "text/filename")
		}
		if entry.dir {
			action, ok := lines[i].(LineAction)
			if !ok {
				t.Errorf("line %d: expected LineAction, got %T", i, lines[i])
				continue
			}

			action.Engage() // Expand.
			if fc.Len() <= len(lines) && !strings.HasSuffix(entry.name, "-empty") {
				t.Errorf("expanding dir %q didn't work", lines[i])
				(*testWriter)(t).printContent(fc)
			}
			if expandedName := lines[i].String(); !strings.Contains(expandedName, "- "+entry.name) {
				t.Errorf("bad display of expanded dir: %q", expandedName)
			}
			if nextLine := fc.Lines()[i+1]; nextLine.String()[0] != '\t' {
				t.Errorf("unexpected state after expanding line %s, next line: %q", lines[i], nextLine)
				(*testWriter)(t).printContent(fc)
			}

			action.Engage() // Collapse.
			if fc.Len() != len(lines) {
				t.Errorf("collapsing dir %q didn't work", lines[i])
				(*testWriter)(t).printContent(fc)
			}
		}
	}
}

type testWriter testing.T

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.Log(string(p))
	return len(p), nil
}

func (tw *testWriter) printContent(c Interface) {
	out := bufio.NewWriter(tw)
	defer out.Flush()
	Print(out, c, 0, c.Len())
}
