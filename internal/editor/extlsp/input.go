package extlsp

import (
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

	var arrow inputs.CursorArrow
	if inputs.IsArrow(b, &arrow) {
		switch arrow {
		case inputs.CursorArrowUp:
			data.SuggestPrev()
			handled = true
		case inputs.CursorArrowDown:
			data.SuggestNext()
			handled = true
		default:
			data.ResetSuggestions()
		}
		return
	}

	if inputs.IsEscape(b) {
		data.ResetSuggestions()
		handled = true
	}

	if inputs.IsTab(b) {
		text := data.CurrentSuggestion()
		data.ResetSuggestions()
		buf.AcceptSuggestion(text)
		handled = true
	}

	return
}
