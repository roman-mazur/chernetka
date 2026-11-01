package extsyntaxhl

import (
	"math/rand/v2"
	"os"
	"slices"
	"strings"
	"testing"
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/change"
	"rmazur.io/chernetka/internal/editor"
	treesitter "rmazur.io/chernetka/internal/editor/extsyntaxhl/gotreesitter"
)

// changesRecorder is an extension collecting the changes passed to AfterEdit.
type changesRecorder struct{ changes []change.Item }

func (cr *changesRecorder) ID() string { return "changes recorder" }
func (cr *changesRecorder) MakeBufferData(editor.Sender, *editor.Buffer) editor.BufferExtData {
	return nil
}
func (cr *changesRecorder) AfterEdit(_ editor.Sender, _ *editor.Buffer, changes []change.Item) {
	cr.changes = append(cr.changes, changes...)
}

// snippets are typed into the documents by randomEdit: they open and close
// the strings, the comments and the blocks, which changes much of the tree.
var snippets = []string{"x", " ", "{", "}", "(", ")", "\"", "`", "/*", "*/", "// ", "func f() {", "\t", ":", "- ", "1"}

// randomEdit changes the buffer like typing, deleting, or pasting does.
func randomEdit(rnd *rand.Rand, m content.Mutable) {
	lines := m.Lines()
	if len(lines) == 0 { // All the lines were deleted.
		m.Insert(0, content.TextLine(snippets[rnd.IntN(len(snippets))]))
		return
	}
	pos := func() content.Position {
		ln := rnd.IntN(len(lines))
		line := lines[ln].String()
		col := rnd.IntN(len(line) + 1)
		for col > 0 && col < len(line) && !utf8.RuneStart(line[col]) {
			col--
		}
		return content.Position{Line: ln, Col: col}
	}
	switch rnd.IntN(6) {
	case 0, 1: // Typing.
		content.InsertText(m, pos(), snippets[rnd.IntN(len(snippets))])
	case 2: // A new line, or a paste of a few.
		text := "\n"
		for range rnd.IntN(3) {
			text += snippets[rnd.IntN(len(snippets))] + "\n"
		}
		content.InsertText(m, pos(), text)
	case 3: // Deleting the text between two positions, maybe joining the lines.
		start, end := pos(), pos()
		if end.Line < start.Line || (end.Line == start.Line && end.Col < start.Col) {
			start, end = end, start
		}
		end.Line = min(end.Line, start.Line+3)
		end.Col = min(end.Col, lines[end.Line].Len())
		if end.Line == start.Line {
			end.Col = max(end.Col, start.Col)
		}
		content.DeleteSpan(m, content.Span{Start: start, End: end})
	case 4: // Deleting whole lines.
		if len(lines) > 1 {
			start := rnd.IntN(len(lines) - 1)
			m.Delete(start, min(len(lines), start+1+rnd.IntN(3)))
		}
	case 5: // Inserting whole lines.
		at := rnd.IntN(len(lines) + 1)
		m.Insert(at, content.TextLine(snippets[rnd.IntN(len(snippets))]), content.TextLine(""))
	}
}

// replayLines applies the change to the lines of a text.
func replayLines(lines []string, ch change.Item) []string {
	strs := func(ls []content.Line) (res []string) {
		for _, l := range ls {
			res = append(res, l.String())
		}
		return res
	}
	pos := ch.Span.Start.Line
	switch ch.Operation {
	case change.OperationInsert:
		return slices.Insert(lines, pos, strs(ch.New)...)
	case change.OperationDelete:
		return slices.Delete(lines, pos, pos+len(ch.Old))
	default:
		lines[pos] = ch.New[0].String()
		return lines
	}
}

// checkEdit checks that the edit describes how before turned into after, the way a tree-sitter
// edit does: the text before StartByte and the one after the old and new ends is the same,
// and the positions are where the offsets are.
func checkEdit(t *testing.T, before, after string, e treesitter.InputEdit) {
	t.Helper()
	start, oldEnd, newEnd := int(e.StartByte), int(e.OldEndByte), int(e.NewEndByte)
	if start > oldEnd || start > newEnd || oldEnd > len(before) || newEnd > len(after) {
		t.Fatalf("edit %+v is out of the texts of %d and %d bytes", e, len(before), len(after))
	}
	if before[:start] != after[:start] || before[oldEnd:] != after[newEnd:] {
		t.Fatalf("edit %+v does not turn\n%q\ninto\n%q", e, before, after)
	}
	pointAt := func(text string, offset int) treesitter.Point {
		lineStart := strings.LastIndexByte(text[:offset], '\n') + 1
		return treesitter.NewPoint(uint(strings.Count(text[:offset], "\n")), uint(offset-lineStart))
	}
	if got := pointAt(before, start); got != e.StartPosition {
		t.Errorf("edit %+v: the start is at %+v", e, got)
	}
	if got := pointAt(before, oldEnd); got != e.OldEndPosition {
		t.Errorf("edit %+v: the old end is at %+v", e, got)
	}
	if got := pointAt(after, newEnd); got != e.NewEndPosition {
		t.Errorf("edit %+v: the new end is at %+v", e, got)
	}
}

func TestTextEdits(t *testing.T) {
	rec := new(changesRecorder)
	var edit editor.Editor
	edit.Extend(rec)
	if err := edit.OpenReader("text.txt", strings.NewReader("one\ntwo\nthree")); err != nil {
		t.Fatal(err)
	}
	buf := edit.Top()

	rnd := rand.New(rand.NewPCG(1, 2))
	for i := range 2000 {
		old, oldLens := buf.Text(), lineLengths(buf.Content.Lines())
		var lines []string // the document may have no lines, unlike the text split by "\n"
		for _, line := range buf.Content.Lines() {
			lines = append(lines, line.String())
		}
		for range 1 + rnd.IntN(3) {
			randomEdit(rnd, buf.Mutate())
		}
		rec.changes = nil
		editor.NotifyExtensions(nil, buf, rec)
		if rec.changes == nil {
			continue // Not mutated, like deleting an empty span.
		}
		cur := buf.Text()

		edits, ok := textEdits(oldLens, rec.changes, buf.Content.Lines())
		if !ok {
			t.Fatalf("edit %d: changes %+v are not converted, turning\n%q\ninto\n%q", i, rec.changes, old, cur)
		}
		if len(edits) != len(rec.changes) {
			t.Fatalf("edit %d: %d edits for %d changes", i, len(edits), len(rec.changes))
		}
		before := old
		for j, ch := range rec.changes {
			lines = replayLines(lines, ch)
			after := strings.Join(lines, "\n")
			checkEdit(t, before, after, edits[j])
			before = after
		}
		if t.Failed() {
			t.Fatalf("edit %d: changes %+v", i, rec.changes)
		}

		if i%100 == 0 {
			// Start over now and then: the document grows.
			buf.Mutate().Delete(0, buf.Content.Len())
			buf.Mutate().Insert(0, content.TextLine("a"), content.TextLine("b"))
			editor.NotifyExtensions(nil, buf, rec)
		}
	}
}

func TestTextEdits_Mismatch(t *testing.T) {
	lens, lines := []int{3, 3}, []content.Line{content.TextLine("one"), content.TextLine("two!")}
	for _, tc := range []struct {
		name    string
		changes []change.Item
	}{
		{"no changes", nil},
		{"another line", []change.Item{change.Replacement(0, content.TextLine("one"), content.TextLine("one!"))}},
		{"other old line", []change.Item{change.Replacement(1, content.TextLine("two?"), content.TextLine("two!"))}},
		{"out of the text", []change.Item{change.Deletion(2, []content.Line{content.TextLine("three")})}},
		{"not all changes", []change.Item{change.Insertion(2, []content.Line{content.TextLine("three")})}},
	} {
		if edits, ok := textEdits(lens, tc.changes, lines); ok {
			t.Errorf("%s: got edits %+v for the changes that don't match the texts", tc.name, edits)
		}
	}
}

// TestIncrementalParse edits the documents, checking that the incremental parse highlights them
// as the parse from scratch does.
func TestIncrementalParse(t *testing.T) {
	for _, tc := range []struct{ path, src string }{
		{"edited.go", "syntaxhl.go"},
		{"edited.yaml", "../../../.github/workflows/test.yml"},
		{"edited.cue", "../../ci/test.cue"},
		{"edited.json", ""},
	} {
		t.Run(tc.path, func(t *testing.T) {
			text := `{"a": [1, "two", {"b": null}]}`
			if tc.src != "" {
				data, err := os.ReadFile(tc.src)
				if err != nil {
					t.Fatal(err)
				}
				text = string(data)
			}
			ext, rec := new(Integration), new(changesRecorder)
			var edit editor.Editor
			edit.Extend(ext)
			if err := edit.OpenReader(tc.path, strings.NewReader(text)); err != nil {
				t.Fatal(err)
			}
			buf := edit.Top()
			t.Cleanup(func() { _ = buf.Close() })
			doc := buf.ExtensionData(ext.ID()).(*document)
			if _, ok := doc.hl.(*tsHighlighter); !ok {
				t.Fatalf("%s is not highlighted with tree-sitter", tc.path)
			}

			rnd := rand.New(rand.NewPCG(3, 4))
			for i := range 200 {
				oldLens := doc.lens
				// The keys handled between two renders.
				rec.changes = nil
				for range 1 + rnd.IntN(3) {
					for range 1 + rnd.IntN(3) {
						randomEdit(rnd, buf.Mutate())
					}
					editor.NotifyExtensions(nil, buf, ext, rec)
				}
				doc.ensureParsed() // Like the render does.
				if rec.changes == nil {
					continue // Not mutated, like deleting an empty span.
				}
				if doc.parsed != buf.Version() {
					t.Fatalf("edit %d: the document is not parsed", i)
				}
				if _, ok := textEdits(oldLens, rec.changes, buf.Content.Lines()); !ok {
					t.Fatalf("edit %d: the document is parsed from scratch", i)
				}

				fresh := &document{LogEmbed: &ext.LogEmbed, buf: buf, lang: doc.lang, hl: doc.lang.newHighlighter()}
				fresh.parse(nil)
				fresh.build()
				_ = fresh.Close()
				if !doc.built {
					doc.build()
				}
				if !slices.Equal(doc.spans, fresh.spans) {
					t.Fatalf("edit %d: the incremental parse highlights the text differently:\n%s", i, buf.Text())
				}
			}
		})
	}
}

// BenchmarkKeystroke measures typing a character into a Go file: the edit, and the parse.
func BenchmarkKeystroke(b *testing.B) {
	data, err := os.ReadFile("treesitter.go")
	if err != nil {
		b.Fatal(err)
	}
	ext := new(Integration)
	var edit editor.Editor
	edit.Extend(ext)
	if err := edit.OpenReader("edited.go", strings.NewReader(string(data))); err != nil {
		b.Fatal(err)
	}
	buf := edit.Top()
	b.Cleanup(func() { _ = buf.Close() })
	doc := buf.ExtensionData(ext.ID()).(*document)
	ln := buf.Content.Len() / 2

	for i := 0; b.Loop(); i++ {
		line := buf.Content.Lines()[ln].String()
		if i%2 == 0 {
			line += "x"
		} else {
			line = line[:len(line)-1]
		}
		buf.Mutate().Update(ln, content.TextLine(line))
		editor.NotifyExtensions(nil, buf, ext)
		doc.ensureParsed() // Like the render does.
	}
}

// TestCodeBlocksBeforeAfterEdit checks the highlights after the code blocks are requested
// by another extension before the edit notification reaches this one.
func TestCodeBlocksBeforeAfterEdit(t *testing.T) {
	buf, ext := openDoc(t, "notes.md", "# Title\n\n```go\nx := 1\n```\n")
	doc := buf.ExtensionData(ext.ID()).(*document)

	buf.Mutate().Insert(0, content.TextLine("```d2"), content.TextLine("a -> b"))
	if blocks := doc.CodeBlocks(); len(blocks) != 1 {
		t.Errorf("got %d code blocks, want 1", len(blocks))
	}
	editor.NotifyExtensions(nil, buf, ext)
	buf.Mutate().Insert(2, content.TextLine("```"))
	editor.NotifyExtensions(nil, buf, ext)

	fresh := &document{LogEmbed: &ext.LogEmbed, buf: buf, lang: doc.lang, hl: doc.lang.newHighlighter()}
	fresh.parse(nil)
	fresh.build()
	doc.SyntaxSpans(0, "```d2")
	if !slices.Equal(doc.spans, fresh.spans) {
		t.Errorf("got spans\n%v\nwant\n%v", doc.spans, fresh.spans)
	}
	if blocks := doc.CodeBlocks(); len(blocks) != 2 {
		t.Errorf("got %d code blocks, want 2", len(blocks))
	}
}
