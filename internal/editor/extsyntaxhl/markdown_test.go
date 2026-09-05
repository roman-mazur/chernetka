package extsyntaxhl

import "testing"

func TestMarkdownBlocks(t *testing.T) {
	doc := hlDoc{
		{"# Title", []string{"0:1:Punctuation", "1:7:Heading"}},
		{"###### Deep ###", []string{"0:6:Punctuation", "6:12:Heading", "12:15:Punctuation"}},
		{"####### Too deep", nil},
		{"#hashtag", nil},
		{"", nil},
		{"Setext", []string{"0:6:Heading"}},
		{"======", []string{"0:6:Punctuation"}},
		{"Dashed", []string{"0:6:Heading"}},
		{"---", []string{"0:3:Punctuation"}},
		{"", nil},
		// The same '---' with no paragraph above it is a thematic break.
		{"---", []string{"0:3:ListMarker"}},
		{"***", []string{"0:3:ListMarker"}},
		{"", nil},
		{"- bullet", []string{"0:1:ListMarker"}},
		{"+ bullet", []string{"0:1:ListMarker"}},
		{"* bullet", []string{"0:1:ListMarker"}},
		{"12. ordered", []string{"0:3:ListMarker"}},
		{"3) ordered", []string{"0:2:ListMarker"}},
		{"- [x] done", []string{"0:1:ListMarker", "2:5:Constant"}},
		{"- [ ] todo", []string{"0:1:ListMarker", "2:5:Constant"}},
		{"not.a list", nil},
		{"", nil},
		{"> quoted", []string{"0:1:Punctuation", "1:8:Quote"}},
		{"> ## in quote", []string{"0:1:Punctuation", "1:2:Quote", "2:4:Punctuation", "4:13:Heading"}},
		// The two markers are adjacent punctuation, so they come back as one span.
		{">> nested", []string{"0:2:Punctuation", "2:9:Quote"}},
	}

	buf, ext := openDoc(t, "notes.md", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestMarkdownInline(t *testing.T) {
	doc := hlDoc{
		{"*italic* and **bold** and ***both***", []string{
			"0:1:Punctuation", "1:7:Emphasis", "7:8:Punctuation",
			"13:15:Punctuation", "15:19:Strong", "19:21:Punctuation",
			"26:29:Punctuation", "29:33:Strong", "33:36:Punctuation",
		}},
		{"_under_ but not snake_case_word", []string{
			"0:1:Punctuation", "1:6:Emphasis", "6:7:Punctuation",
		}},
		{"~~struck~~", []string{"0:2:Punctuation", "2:8:Comment", "8:10:Punctuation"}},
		{"a `code span` and ``a ` inside``", []string{
			"2:3:Punctuation", "3:12:RawText", "12:13:Punctuation",
			"18:20:Punctuation", "20:30:RawText", "30:32:Punctuation",
		}},
		// An unterminated delimiter is left alone rather than swallowing the line.
		{"a * lone star and `unclosed", nil},
		{"**bold at line start**", []string{
			"0:2:Punctuation", "2:20:Strong", "20:22:Punctuation",
		}},
		// Code binds tighter than emphasis, and nests inside it.
		{"**bold `code`**", []string{
			"0:2:Punctuation", "2:7:Strong", "7:8:Punctuation",
			"8:12:RawText", "12:15:Punctuation",
		}},
		{"escaped \\*stars\\* stay", []string{"8:10:Escape", "15:17:Escape"}},
	}

	buf, ext := openDoc(t, "notes.md", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestMarkdownLinks(t *testing.T) {
	doc := hlDoc{
		{"[label](https://a.example)", []string{
			"0:1:Punctuation", "1:6:Link", "6:8:Punctuation", "8:25:URL", "25:26:Punctuation",
		}},
		{"[l](/p \"Title\")", []string{
			"0:1:Punctuation", "1:2:Link", "2:4:Punctuation", "4:6:URL",
			"6:14:StringLiteral", "14:15:Punctuation",
		}},
		{"![alt](/img.png)", []string{
			"0:2:Punctuation", "2:5:Link", "5:7:Punctuation", "7:15:URL", "15:16:Punctuation",
		}},
		{"[label][ref]", []string{
			"0:1:Punctuation", "1:6:Link", "6:8:Punctuation", "8:11:URL", "11:12:Punctuation",
		}},
		{"<https://angle.example>", []string{
			"0:1:Punctuation", "1:22:URL", "22:23:Punctuation",
		}},
		// An HTML tag is punctuation throughout, so the whole tag is one span.
		{"a <br> tag", []string{"2:6:Punctuation"}},
		{"bare https://x.example/a, ok", []string{"5:24:URL"}},
		{"nourl inhttps://x.example", nil},
		{"[ref]: https://a.example \"T\"", []string{
			"0:1:Punctuation", "1:4:Link", "4:6:Punctuation", "7:24:URL", "25:28:StringLiteral",
		}},
	}

	buf, ext := openDoc(t, "notes.md", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestMarkdownCodeBlocks(t *testing.T) {
	doc := hlDoc{
		{"```go", []string{"0:3:Punctuation", "3:5:Keyword"}},
		// Everything inside a fence is raw text, markup and all.
		{"# not a heading", []string{"0:15:RawText"}},
		{"", nil},
		{"~~~ not a close", []string{"0:15:RawText"}},
		{"```", []string{"0:3:Punctuation"}},
		{"# a heading again", []string{"0:1:Punctuation", "1:17:Heading"}},
		{"", nil},
		{"    indented code", []string{"4:17:RawText"}},
		{"", nil},
		{"a paragraph", nil},
		// Four spaces under a paragraph continue it instead of opening a block.
		{"    still the paragraph", nil},
		{"", nil},
		{"- a list item", []string{"0:1:ListMarker"}},
		{"    a continuation line, not code", nil},
	}

	buf, ext := openDoc(t, "notes.md", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestMarkdownTables(t *testing.T) {
	doc := hlDoc{
		{"a | pipe in a sentence", nil},
		{"", nil},
		{"| A | B |", []string{
			"0:1:Punctuation", "1:4:Strong", "4:5:Punctuation", "5:8:Strong", "8:9:Punctuation",
		}},
		{"|---|--:|", []string{"0:9:Punctuation"}},
		{"| 1 | `x` |", []string{
			"0:1:Punctuation", "4:5:Punctuation", "6:7:Punctuation",
			"7:8:RawText", "8:9:Punctuation", "10:11:Punctuation",
		}},
		{"", nil},
		{"| after the table", nil},
	}

	buf, ext := openDoc(t, "notes.md", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

// TestMarkdownUnterminatedFence checks that a fence left open colorizes the
// rest of the document instead of tripping over the end of it.
func TestMarkdownUnterminatedFence(t *testing.T) {
	doc := hlDoc{
		{"text", nil},
		{"```", []string{"0:3:Punctuation"}},
		{"code", []string{"0:4:RawText"}},
		{"more", []string{"0:4:RawText"}},
	}

	buf, ext := openDoc(t, "notes.md", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestMarkdownAfterEdit(t *testing.T) {
	doc := hlDoc{
		{"# Title", []string{"0:1:Punctuation", "1:7:Heading"}},
		{"- item", []string{"0:1:ListMarker"}},
	}

	buf, ext := openDoc(t, "notes.md", doc.text())
	hl := highlighterOf(t, buf, ext)
	doc.check(t, hl)

	insertLine(t, buf, ext, 1)
	hlDoc{doc[0], {"", nil}, doc[1]}.check(t, hl)
}

// TestMarkdownProjectDocs runs the scanner over the repository's own Markdown
// files, checking only that every span stays inside its line.
func TestMarkdownProjectDocs(t *testing.T) {
	for _, path := range []string{"../../../README.md", "../../../TODO.md"} {
		t.Run(path, func(t *testing.T) {
			buf, ext := openDoc(t, path, readFile(t, path))
			hl := highlighterOf(t, buf, ext)
			for i, line := range buf.Content.Lines() {
				text := line.String()
				for _, s := range hl.SyntaxSpans(i, text) {
					if s.LineNumber != i || s.Start < 0 || s.End > len(text) || s.Start >= s.End {
						t.Errorf("line %d %q: span %s is out of range", i, text, s)
					}
				}
			}
		})
	}
}
