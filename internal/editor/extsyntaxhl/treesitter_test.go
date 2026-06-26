package extsyntaxhl

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
)

func TestIntegration_AfterEdit(t *testing.T) {
	t.Run("no syntax highlight", func(t *testing.T) {
		shl := new(Integration)
		var (
			edit editor.Editor
			buf  editor.Buffer
		)
		shl.AfterEdit(&edit, &buf) // should not crash
	})
}

func TestIntegration(t *testing.T) {
	example := testExamples{
		{line: "// Command comment.", tokens: []string{"0:19:TtComment"}},
		{line: "package main", tokens: []string{"0:7:TtKeyword", "8:12:TtIdentifier"}},
		{line: "import \"fmt\"", tokens: []string{"0:6:TtKeyword", "7:12:TtStringLiteral"}},
		{line: "import \"time\"", tokens: []string{"0:6:TtKeyword", "7:13:TtStringLiteral"}},
		{line: "func main() {", tokens: []string{"0:4:TtKeyword", "5:9:TtIdentifier"}},
		{line: "\tfmt.Println(\"Hello World\")", tokens: []string{"1:4:TtIdentifier", "5:12:TtIdentifier", "13:26:TtStringLiteral"}},
		{line: "\tif time.Now().Day() == time.Sunday {", tokens: []string{"1:3:TtKeyword", "4:8:TtIdentifier", "9:12:TtIdentifier", "15:18:TtIdentifier", "24:28:TtIdentifier", "29:35:TtIdentifier"}},
		{line: "\t\tfmt.Println(\"It's Sunday!\")", tokens: []string{"2:5:TtIdentifier", "6:13:TtIdentifier", "14:28:TtStringLiteral"}},
		{line: "\t} else {", tokens: []string{"3:7:TtKeyword"}},
		{line: "\t\tfmt.Println(\"It's not Sunday!\")", tokens: []string{"2:5:TtIdentifier", "6:13:TtIdentifier", "14:32:TtStringLiteral"}},
		{line: "\t}", tokens: []string{"0:2:TtNothing"}},
		{line: "}", tokens: []string{"0:1:TtNothing"}},
	}

	lines := example.lines()

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

	syntaxHighlighter := new(Integration)
	syntaxHighlighter.LogDebug = true
	syntaxHighlighter.LogF = t.Logf

	edit.Extend(syntaxHighlighter)

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

	example.match(t, data)

	insertEmptyLine := func(buf *editor.Buffer, hl editor.Extension) int {
		mut := buf.Mutate()
		idx := min(17, buf.Content.Len()/2)
		mut.Insert(idx, content.TextLine(""))
		hl.AfterEdit(&edit, buf)
		return idx
	}

	t.Run("edit/insert-empty-line", func(t *testing.T) {
		syntaxHighlighter.LogF = t.Logf
		insertEmptyLine(&buf, syntaxHighlighter)

		exampleV2 := example.copy()
		exampleV2 = slices.Insert(exampleV2, len(example)/2, testCase{tokens: []string{"0:0:TtNothing"}})
		exampleV2.match(t, buf.ExtensionData(syntaxHighlighter.ID()).(*syntaxTree))
	})

	t.Run("real-sources", func(t *testing.T) {
		entries, err := os.ReadDir(".")
		if err != nil {
			t.Fatal(err)
		}
		for _, entry := range entries {
			if !entry.IsDir() && filepath.Ext(entry.Name()) == ".go" {
				t.Run(entry.Name(), func(t *testing.T) {
					var edit editor.Editor
					syntaxHighlighter := new(Integration)
					syntaxHighlighter.LogF = t.Logf
					edit.Extend(syntaxHighlighter)

					f, err := os.Open(entry.Name())
					must(t, err)
					t.Cleanup(func() { _ = f.Close() })

					must(t, edit.OpenReader(entry.Name(), f))

					checkHl := func(buf *editor.Buffer, ctxMsg string) {
						t.Helper()
						hl := buf.ExtensionData(syntaxHighlighter.ID()).(editor.SyntaxHighlighter)
						for i, line := range buf.Content.Lines() {
							textLine := line.String()
							spans := hl.SyntaxSpans(i, textLine)
							textLen := len(strings.TrimSpace(textLine))
							if textLen > 2 && len(spans) < 2 {
								if len(spans) == 1 && spans[0].TokenType != editor.TtNothing {
									continue
								}
								t.Log("line:", textLine)
								t.Log("spans:", spans)
								t.Errorf("%s > line %d: got %d spans, want at least %d",
									ctxMsg, i, len(spans), 2)
							}
						}
					}

					buf := edit.Top()
					// TODO: something must be wrong with Buffer Text() implementation.

					t.Log("initial check")
					checkHl(buf, "initial check")

					oldLen := buf.Content.Len()
					t.Logf("insert empty line at %d and check", insertEmptyLine(buf, syntaxHighlighter))
					if buf.Content.Len() == oldLen {
						t.Error("content Len() didn't change after insert")
					}
					if len(buf.Content.Lines()) != buf.Content.Len() {
						t.Error("inconsistent content length")
					}
					checkHl(buf, "after insert")
				})
			}
		}
	})
}

func must(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

type testCase struct {
	line   string
	tokens []string
}

func (tc testCase) copy() testCase {
	res := testCase{
		line:   tc.line,
		tokens: make([]string, len(tc.tokens)),
	}
	copy(res.tokens, tc.tokens)
	return res
}

type testExamples []testCase

func (te testExamples) lines() []string {
	lines := make([]string, len(te))
	for i, ex := range te {
		lines[i] = ex.line
	}
	return lines
}

func (te testExamples) expectedTokens(ln int) []string {
	res := make([]string, len(te[ln].tokens))
	for i, suffix := range te[ln].tokens {
		res[i] = fmt.Sprintf("%d:%s", ln, suffix)
	}
	return res
}

func (te testExamples) match(t *testing.T, data *syntaxTree) {
	t.Helper()
	for i, ex := range te {
		spans := data.SyntaxSpans(i, ex.line)
		if ss, es := fmt.Sprint(spans), fmt.Sprint(te.expectedTokens(i)); ss != es {
			t.Errorf("line %d: got %s, want %s", i, ss, es)
		}
	}
}

func (te testExamples) copy() testExamples {
	res := make(testExamples, len(te))
	for i := range res {
		res[i] = te[i].copy()
	}
	return res
}
