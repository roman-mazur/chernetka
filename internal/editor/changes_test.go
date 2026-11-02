package editor_test

import (
	"slices"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/change"
	"rmazur.io/chernetka/internal/editor"
)

// changesCheckExt checks that the changes passed to AfterEdit turn the text the extension
// saw the last time into the current one, for the buffers holding a text file.
type changesCheckExt struct {
	t       testing.TB
	changes *int // counts the changes checked, if not nil
}

// changesCheckData is the text of the buffer the extension saw the last time.
type changesCheckData struct {
	lines []string
}

func (ce changesCheckExt) ID() string { return "changes check" }

func (ce changesCheckExt) MakeBufferData(_ editor.Sender, buf *editor.Buffer) editor.BufferExtData {
	if !buf.HoldsTextFile() {
		return nil
	}
	return &changesCheckData{lines: lineStrings(buf.Content)}
}

func (ce changesCheckExt) AfterEdit(_ editor.Sender, buf *editor.Buffer, changes []change.Item) {
	data, ok := buf.ExtensionData(ce.ID()).(*changesCheckData)
	if !ok {
		return
	}
	if ce.changes != nil {
		*ce.changes += len(changes)
	}
	replayed := slices.Clone(data.lines)
	for _, ch := range changes {
		var err string
		if replayed, err = replayChange(replayed, ch); err != "" {
			ce.t.Errorf("%s: change %+v: %s", buf.Path, ch, err)
			break
		}
	}
	data.lines = lineStrings(buf.Content)
	if !slices.Equal(replayed, data.lines) {
		ce.t.Errorf("%s: the changes %+v make\n%q\nof the text, while it's\n%q",
			buf.Path, changes, strings.Join(replayed, "\n"), strings.Join(data.lines, "\n"))
	}
}

func lineStrings(doc content.Document) []string {
	var lines []string
	for _, line := range doc.Lines() {
		lines = append(lines, line.String())
	}
	return lines
}

// replayChange applies the change to the lines, checking that it describes them.
// It returns a description of the mismatch if it does not.
func replayChange(lines []string, ch change.Item) ([]string, string) {
	pos := ch.Span.Start.Line
	switch ch.Operation {
	case change.OperationInsert:
		if pos < 0 || pos > len(lines) || len(ch.New) == 0 {
			return lines, "no lines to insert there"
		}
		return slices.Insert(lines, pos, lineStrings((*content.FullText)(&ch.New))...), ""
	case change.OperationDelete:
		end := pos + len(ch.Old)
		if pos < 0 || end > len(lines) || len(ch.Old) == 0 {
			return lines, "no lines to delete there"
		}
		if !slices.Equal(lines[pos:end], lineStrings((*content.FullText)(&ch.Old))) {
			return lines, "the deleted lines differ"
		}
		return slices.Delete(lines, pos, end), ""
	case change.OperationUpdate:
		if pos < 0 || pos >= len(lines) || len(ch.Old) != 1 || len(ch.New) != 1 {
			return lines, "no line to update there"
		}
		if lines[pos] != ch.Old[0].String() {
			return lines, "the old line differs"
		}
		lines[pos] = ch.New[0].String()
		return lines, ""
	}
	return lines, "unknown operation"
}

// TestAfterEditChanges checks the changes passed to the extensions for the kinds of edits
// the fuzz test may miss: these are the edits through the editor loop.
func TestAfterEditChanges(t *testing.T) {
	for _, tc := range []struct {
		name, text string
		inputs     []string
	}{
		{name: "typing", text: "one", inputs: []string{"A two\rthree\x1b"}},
		{name: "joining lines", text: "one\ntwo", inputs: []string{"ji\x7f\x7f\x1b"}},
		{name: "deleting the selected lines", text: "one\ntwo\nthree\nfour", inputs: []string{"l\x1b[1;2B\x1b[1;2Bx", "G\x1b[1;2Ax"}},
		{name: "deleting the text", text: "one\ntwo", inputs: []string{"\x1b[1;2B\x1b[1;2Fx"}},
		{name: "paste", text: "one\n", inputs: []string{"A\x1b[200~ pasted\n\ttext\nmore\x1b[201~\x1b"}},
		{name: "undo and redo", text: "one\ntwo", inputs: []string{"otext\x1b", "\x1a", "\x1b[122;6u", "\x1a\x1a"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := editor.NewTestHarness(t)
			var changes int
			h.Extend(changesCheckExt{t: t, changes: &changes})
			if err := h.OpenReader("file.txt", strings.NewReader(tc.text)); err != nil {
				t.Fatal(err)
			}
			h.Run(t)
			for _, in := range tc.inputs {
				h.SendInput(t, []byte(in))
			}
			if changes == 0 {
				t.Error("no changes passed to AfterEdit")
			}
		})
	}
}
