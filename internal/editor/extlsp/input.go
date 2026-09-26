package extlsp

import (
	"sort"

	"go.lsp.dev/protocol"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/inputs"
)

func (le *Integration) HandleInsertInput(buf *editor.Buffer, _ *editor.RenderPrefs, b []byte) (handled bool) {
	data, ok := buf.ExtensionData(le.ID()).(*BufferData)
	if !ok {
		return
	}

	if !data.HasSuggestions() {
		return
	}

	var (
		arrow inputs.Cursor
		mod   inputs.Modifier
	)
	if inputs.IsCursor(b, &arrow, &mod) {
		multiple := len(data.suggestions) > 1
		switch {
		case arrow == inputs.CursorArrowUp && multiple:
			data.SuggestPrev()
			handled = true
		case arrow == inputs.CursorArrowDown && multiple:
			data.SuggestNext()
			handled = true
		default:
			data.ResetSuggestions()
		}
		return
	}

	if inputs.IsEscape(b) {
		// Only dismiss: the next Esc leaves the insert mode.
		data.ResetSuggestions()
		handled = true
	}

	if inputs.IsTab(b) {
		sug := data.suggestions[data.sugIdx]
		data.ResetSuggestions()
		applyEdits(buf, sug.edits)
		buf.AcceptSuggestion(sug.text, sug.cursor)
		handled = true
	}

	return
}

// applyEdits applies additional LSP text edits that come with a completion,
// like adding a missing import. Edits never overlap, so applying them from the
// last to the first keeps the positions of the remaining ones valid.
func applyEdits(buf *editor.Buffer, edits []protocol.TextEdit) {
	edits = append([]protocol.TextEdit(nil), edits...)
	sort.Slice(edits, func(i, j int) bool {
		a, b := edits[i].Range.Start, edits[j].Range.Start
		return a.Line > b.Line || (a.Line == b.Line && a.Character > b.Character)
	})
	for _, te := range edits {
		if span, ok := spanOf(buf.Content.Lines(), te.Range); ok {
			buf.ReplaceText(span, te.NewText)
		}
	}
}
