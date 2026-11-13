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
	h := editor.NewTestHarness(t)
	h.Extend(rec)
	if err := h.OpenReader("text.txt", strings.NewReader("one\ntwo\nthree")); err != nil {
		t.Fatal(err)
	}
	buf := h.Top()

	rnd := rand.New(rand.NewPCG(1, 2))
	for i := range 2000 {
		old := buf.Text()
		var lines []string // the document may have no lines, unlike the text split by "\n"
		for _, line := range buf.Content.Lines() {
			lines = append(lines, line.String())
		}
		for range 1 + rnd.IntN(3) {
			randomEdit(rnd, buf.Mutate())
		}
		rec.changes = nil
		h.NotifyExtensions(buf)
		if rec.changes == nil {
			continue // Not mutated, like deleting an empty span.
		}
		cur := buf.Text()

		edits, ok := textEdits(lineLens(old), newSource(cur), rec.changes)
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
			h.NotifyExtensions(buf)
		}
	}
}

// lineLens returns the lengths of the lines of the text.
func lineLens(text string) (lens []int) {
	for _, line := range newSource(text).lines {
		lens = append(lens, len(line))
	}
	return lens
}

func TestTextEdits_Mismatch(t *testing.T) {
	old, cur := "one\ntwo", newSource("one\ntwo!")
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
		if edits, ok := textEdits(lineLens(old), cur, tc.changes); ok {
			t.Errorf("%s: got edits %+v for the changes that don't match the texts", tc.name, edits)
		}
	}
	if edits, ok := textEdits(lineLens("one\ntwo"), newSource("one\ntwo"), nil); !ok || len(edits) > 0 {
		t.Errorf("got edits %+v, %t for the same text, want none", edits, ok)
	}
}

// TestIncrementalParse edits the documents, checking that the incremental parse highlights them
// as the parse from scratch does.
//
// The Go input is a fixed copy of a source file, so changing the source doesn't change the test.
// Tree-sitter may recover from the syntax errors differently on an incremental parse: it reuses
// the tokens outside the edits, which keep the meaning they got in the error recovery of the old
// tree. For example, "*func\nfunc" edited into "*0\nfunc\nfunc" ends with an identifier in an
// ERROR node, while the parse from scratch has the func keyword there. Other inputs may hit it.
func TestIncrementalParse(t *testing.T) {
	for _, tc := range []struct{ path, src string }{
		{"edited.go", "testdata/incremental.go.txt"},
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
			h := editor.NewTestHarness(t)
			h.Extend(ext)
			h.Extend(rec)
			if err := h.OpenReader(tc.path, strings.NewReader(text)); err != nil {
				t.Fatal(err)
			}
			buf := h.Top()
			t.Cleanup(func() { _ = buf.Close() })
			doc := buf.ExtensionData(ext.ID()).(*document)
			if _, ok := doc.hl.(*tsHighlighter); !ok {
				t.Fatalf("%s is not highlighted with tree-sitter", tc.path)
			}
			doc.ensureParsed() // Like the first render does.

			rnd := rand.New(rand.NewPCG(3, 4))
			for i := range 200 {
				oldLens, oldText, parses := slices.Clone(doc.lineLens), buf.Text(), doc.parses
				// The keys handled between two renders.
				rec.changes = nil
				for range 1 + rnd.IntN(3) {
					for range 1 + rnd.IntN(3) {
						randomEdit(rnd, buf.Mutate())
					}
					h.NotifyExtensions(buf)
				}
				if len(rec.changes) > 0 && (doc.changesLost || !doc.stale) {
					t.Fatalf("edit %d: the document does not have all the changes", i)
				}
				doc.ensureParsed() // Like the render does.
				if buf.Text() == oldText {
					continue // Like a line deleted and inserted back.
				}
				if doc.parses != parses+1 {
					t.Fatalf("edit %d: the document is parsed %d times", i, doc.parses-parses)
				}
				if _, ok := textEdits(oldLens, newSource(buf.Text()), rec.changes); !ok {
					t.Fatalf("edit %d: the document is parsed from scratch", i)
				}

				fresh := &document{LogEmbed: &ext.LogEmbed, buf: buf, lang: doc.lang, hl: doc.lang.newHighlighter()}
				fresh.parse(nil)
				_ = fresh.Close()
				if !slices.Equal(doc.spans, fresh.spans) {
					t.Fatalf("edit %d: the incremental parse highlights the text differently:\n%s", i, buf.Text())
				}
			}
		})
	}
}

// BenchmarkKeystroke measures typing a character into a Go file: the edit, the parse, and building the spans.
func BenchmarkKeystroke(b *testing.B) {
	data, err := os.ReadFile("treesitter.go")
	if err != nil {
		b.Fatal(err)
	}
	ext := new(Integration)
	h := editor.NewTestHarness(b)
	h.Extend(ext)
	if err := h.OpenReader("edited.go", strings.NewReader(string(data))); err != nil {
		b.Fatal(err)
	}
	buf := h.Top()
	b.Cleanup(func() { _ = buf.Close() })
	doc := buf.ExtensionData(ext.ID()).(*document)
	doc.ensureParsed() // Like the first render does.
	ln := buf.Content.Len() / 2

	for i := 0; b.Loop(); i++ {
		line := buf.Content.Lines()[ln].String()
		if i%2 == 0 {
			line += "x"
		} else {
			line = line[:len(line)-1]
		}
		buf.Mutate().Update(ln, content.TextLine(line))
		h.NotifyExtensions(buf)
		doc.ensureParsed() // Like the render does.
	}
}

// TestCodeBlocksInAfterEdit checks the highlights after another extension reads the code blocks
// in its AfterEdit, which comes after the one of the syntax highlighter.
func TestCodeBlocksInAfterEdit(t *testing.T) {
	reader := new(blocksReader)
	h, buf, ext := openHarness(t, "notes.md", "# Title\n\n```go\nx := 1\n```\n", reader)
	doc := buf.ExtensionData(ext.ID()).(*document)

	buf.Mutate().Insert(0, content.TextLine("```d2"), content.TextLine("a -> b"))
	h.NotifyExtensions(buf)
	buf.Mutate().Insert(2, content.TextLine("```"))
	h.NotifyExtensions(buf)
	if got := len(reader.blocks); got != 2 || len(reader.blocks[0]) != 1 || len(reader.blocks[1]) != 2 {
		t.Errorf("the blocks read in AfterEdit: %v, want 1 block, then 2", reader.blocks)
	}

	fresh := &document{LogEmbed: &ext.LogEmbed, buf: buf, lang: doc.lang, hl: doc.lang.newHighlighter()}
	fresh.parse(nil)
	doc.SyntaxSpans(0, "```d2")
	if !slices.Equal(doc.spans, fresh.spans) {
		t.Errorf("got spans\n%v\nwant\n%v", doc.spans, fresh.spans)
	}
}

// TestIncrementalParse_EditedBuffer checks that the changes are parsed incrementally in a buffer
// edited before its document was made.
func TestIncrementalParse_EditedBuffer(t *testing.T) {
	rec := new(changesRecorder)
	h := editor.NewTestHarness(t)
	h.Extend(rec)
	if err := h.OpenReader("edited.go", strings.NewReader("package main\n\nfunc main() {}\n")); err != nil {
		t.Fatal(err)
	}
	buf := h.Top()
	t.Cleanup(func() { _ = buf.Close() })
	buf.Mutate().Insert(1, content.TextLine("// Edited."))
	h.NotifyExtensions(buf)

	ext := new(Integration)
	doc := ext.MakeBufferData(nil, buf).(*document)
	doc.ensureParsed()
	parses := doc.parses

	buf.Mutate().Update(0, content.TextLine("package edited"))
	rec.changes = nil
	h.NotifyExtensions(buf)
	doc.edited(rec.changes)
	if doc.changesLost || len(doc.pending) != len(rec.changes) {
		t.Fatalf("the document has %d changes of %d", len(doc.pending), len(rec.changes))
	}
	doc.ensureParsed()
	if doc.parses != parses+1 || len(doc.pending) != 0 {
		t.Errorf("the edit is parsed %d times, with %d changes left", doc.parses-parses, len(doc.pending))
	}
}

// TestUnknownChanges checks that the document is parsed from scratch when the changes are not
// known, like when the content is reloaded, even with the changes known before.
func TestUnknownChanges(t *testing.T) {
	h, buf, ext := openHarness(t, "edited.go", "package main\n\nfunc main() {}\n")
	doc := buf.ExtensionData(ext.ID()).(*document)
	doc.ensureParsed()

	buf.Mutate().Insert(1, content.TextLine("// One."))
	h.NotifyExtensions(buf)
	buf.Mutate().Update(0, content.TextLine("package reloaded"))
	buf.Mutate().Insert(3, content.TextLine(`var s = "two"`))
	ext.AfterEdit(nil, buf, nil) // Like the editor does after a reload.
	if !doc.changesLost || len(doc.pending) != 0 {
		t.Fatalf("the document keeps %d changes, changesLost: %t", len(doc.pending), doc.changesLost)
	}

	doc.ensureParsed()
	fresh := &document{LogEmbed: &ext.LogEmbed, buf: buf, lang: doc.lang, hl: doc.lang.newHighlighter()}
	fresh.parse(nil)
	_ = fresh.Close()
	if !slices.Equal(doc.spans, fresh.spans) {
		t.Errorf("got spans\n%v\nwant\n%v", doc.spans, fresh.spans)
	}
	if doc.changesLost {
		t.Error("the changes are still lost after the parse")
	}
}
