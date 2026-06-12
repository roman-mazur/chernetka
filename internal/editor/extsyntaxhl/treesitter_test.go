package extsyntaxhl

import (
	"fmt"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
)

func TestIntegration(t *testing.T) {
	example := []struct {
		line   string
		tokens []string
	}{
		{line: "package main", tokens: []string{"0:0:7:TtKeyword", "0:8:12:TtIdentifier"}},
		{line: "import \"fmt\"", tokens: []string{"1:0:6:TtKeyword", "1:7:12:TtStringLiteral"}},
		{line: "func main() {", tokens: []string{"2:0:4:TtKeyword", "2:5:9:TtIdentifier"}},
		{line: "\tfmt.Println(\"Hello World\")", tokens: []string{"3:1:4:TtIdentifier", "3:5:12:TtIdentifier", "3:13:26:TtStringLiteral"}},
		{line: "}", tokens: []string{"4:0:1:TtNothing"}},
	}
	lines := make([]string, len(example))
	for i := range example {
		lines[i] = example[i].line
	}

	prog := strings.Join(lines, "\n")
	text, err := content.LoadFullText(strings.NewReader(prog))
	if err != nil {
		t.Fatal(err)
	}

	buf := editor.Buffer{
		Path:    "test.go",
		Content: &text,
	}
	var edit editor.Editor
	edit.Extend(new(Integration))

	edit.OpenBuffer(&buf)

	data, ok := buf.ExtensionData("syntaxhl").(*syntaxTree)
	if !ok {
		t.Fatal("didn't get the syntax tree for the buffer")
	}
	root := data.tree.RootNode()
	rootStart, rootEnd := root.ByteRange()
	if rootStart != 0 || rootEnd != uint(len(prog)) {
		t.Errorf("rootStart: %d, rootEnd: %d", rootStart, rootEnd)
	}

	for i, line := range lines {
		spans := data.SyntaxSpans(i, line)
		if ss := fmt.Sprint(spans); fmt.Sprint(example[i].tokens) != ss {
			t.Errorf("line %d: got %s, want %s", i, ss, fmt.Sprint(example[i].tokens))
		}
	}
}
