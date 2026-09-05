package extsyntaxhl

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unicode"

	"rmazur.io/chernetka/internal/editor"
)

func TestGoHighlight(t *testing.T) {
	doc := hlDoc{
		{"// Command comment.", []string{"0:19:Comment"}},
		{"package main", []string{"0:7:Keyword", "8:12:Identifier"}},
		{"", nil},
		{"import \"fmt\"", []string{"0:6:Keyword", "7:12:ImportRef"}},
		{"", nil},
		{"type T struct{ N int }", []string{
			"0:4:Keyword", "5:6:TypeRef", "7:13:Keyword", "15:16:Field", "17:20:TypeRef",
		}},
		{"", nil},
		{"func (t *T) Do(s string) error {", []string{
			"0:4:Keyword", "6:7:Identifier", "9:10:TypeRef",
			"12:14:FuncDeclaration", "15:16:Identifier", "17:23:TypeRef", "25:30:TypeRef",
		}},
		{"\tconst x = 1.5", []string{"1:6:Keyword", "7:8:Identifier", "11:14:NumberLiteral"}},
		// The \t escape sequence sits inside the string literal and overrides it.
		{"\tfmt.Println(\"a\\tb\", t.N, x, nil)", []string{
			"1:4:Identifier", "5:12:Call", "13:15:StringLiteral", "15:17:Escape",
			"17:19:StringLiteral", "21:22:Identifier", "23:24:Field", "26:27:Identifier",
			"29:32:Constant",
		}},
		{"\treturn nil", []string{"1:7:Keyword", "8:11:Constant"}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "main.go", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

// TestGoMultiLineLiteral covers a node that starts on one line and ends on
// another: it has to be cut into one span per line, each clipped to its own
// line, instead of being reported with the start line's number and the end
// line's column.
func TestGoMultiLineLiteral(t *testing.T) {
	doc := hlDoc{
		{"package main", []string{"0:7:Keyword", "8:12:Identifier"}},
		{"", nil},
		{"var raw = `first", []string{"0:3:Keyword", "4:7:Identifier", "10:16:StringLiteral"}},
		{"second line is longer", []string{"0:21:StringLiteral"}},
		{"third`", []string{"0:6:StringLiteral"}},
		{"", nil},
		{"/* a block", []string{"0:10:Comment"}},
		{"   comment */", []string{"0:13:Comment"}},
	}

	buf, ext := openDoc(t, "main.go", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestGoHighlightAfterEdit(t *testing.T) {
	doc := hlDoc{
		{"package main", []string{"0:7:Keyword", "8:12:Identifier"}},
		{"", nil},
		{"func main() {}", []string{"0:4:Keyword", "5:9:FuncDeclaration"}},
	}

	buf, ext := openDoc(t, "main.go", doc.text())
	hl := highlighterOf(t, buf, ext)
	doc.check(t, hl)

	// Inserting a line shifts everything below it down by one.
	insertLine(t, buf, ext, 1)
	shifted := hlDoc{doc[0], {"", nil}, doc[1], doc[2]}
	shifted.check(t, hl)
}

// TestLanguageForPath checks that only the recognized file types get a
// highlighter, and that an unknown one is left alone rather than mishandled.
func TestLanguageForPath(t *testing.T) {
	for path, want := range map[string]string{
		"main.go":         "go",
		"a/b/main.GO":     "go",
		"notes.md":        "markdown",
		"notes.markdown":  "markdown",
		"":                "",
		"Makefile":        "",
		"query.sql":       "",
		"go":              "",
		"archive.go.bak":  "",
		"dir.go/file.txt": "",
	} {
		got := ""
		if lang := languageForPath(path); lang != nil {
			got = lang.name
		}
		if got != want {
			t.Errorf("languageForPath(%q) = %q, want %q", path, got, want)
		}
	}
}

func TestNoHighlightForUnknownType(t *testing.T) {
	buf, ext := openDoc(t, "notes.txt", "some text\n")
	if data := buf.ExtensionData(ext.ID()); data != nil {
		t.Errorf("got extension data %#v for an unknown file type, want none", data)
	}
	// AfterEdit has to tolerate a buffer it never made data for.
	ext.AfterEdit(nil, buf)
	// So does a buffer that was never passed through MakeBufferData at all.
	ext.AfterEdit(nil, new(editor.Buffer))
}

// TestGoRealSources runs the highlighter over this package's own sources. It is
// a sanity check rather than an exact comparison: every line with real content
// should come back with some highlight, and no span may fall outside its line.
func TestGoRealSources(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".go" {
			continue
		}
		t.Run(entry.Name(), func(t *testing.T) {
			src, err := os.ReadFile(entry.Name())
			if err != nil {
				t.Fatal(err)
			}
			buf, ext := openDoc(t, entry.Name(), string(src))
			hl := highlighterOf(t, buf, ext)

			check := func(stage string) {
				t.Helper()
				for i, line := range buf.Content.Lines() {
					text := line.String()
					spans := hl.SyntaxSpans(i, text)
					for _, s := range spans {
						if s.LineNumber != i || s.Start < 0 || s.End > len(text) || s.Start >= s.End {
							t.Fatalf("%s > line %d %q: span %s is out of range", stage, i, text, s)
						}
					}
					// Any line with a word on it is a keyword, an identifier, a
					// comment or a string, so it has to come back colorized.
					// Lines of pure punctuation such as "}}," legitimately do not.
					if len(spans) == 0 && strings.ContainsFunc(text, unicode.IsLetter) {
						t.Errorf("%s > line %d %q: no highlight", stage, i, text)
					}
				}
			}

			check("initial")

			before := buf.Content.Len()
			insertLine(t, buf, ext, before/2)
			if buf.Content.Len() != before+1 {
				t.Fatalf("content length %d after insert, want %d", buf.Content.Len(), before+1)
			}
			check("after insert")
		})
	}
}

// TestNoHighlightForDirectoryListing covers a directory whose own name ends in
// a recognized extension: the buffer lists file names, not code.
func TestNoHighlightForDirectoryListing(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "notes.md")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "a.txt"), []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}

	ext := new(Integration)
	ext.LogF = t.Logf
	var edit editor.Editor
	edit.Extend(ext)
	edit.OpenDir(dir, noopOpener{})

	buf := edit.Top()
	if buf.Path != "notes.md" {
		t.Fatalf("buffer path is %q, want the directory name", buf.Path)
	}
	if data := buf.ExtensionData(ext.ID()); data != nil {
		t.Errorf("got extension data %#v for a directory listing, want none", data)
	}
}

// noopOpener satisfies content.OpenFile for a test that never opens anything.
type noopOpener struct{}

func (noopOpener) OpenFile(string) {}
