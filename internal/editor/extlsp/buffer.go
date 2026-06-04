package extlsp

import "go.lsp.dev/uri"

type BufferData struct {
	docUri uri.URI

	suggestions []string
	sugIdx      int // what suggestion is picked

	version       int32
	reqCompletion int
}

func (lbd *BufferData) SetPath(p string) {
	lbd.docUri = uri.File(p)
}

func (lbd *BufferData) HasSuggestions() bool {
	return len(lbd.suggestions) > 0
}

func (lbd *BufferData) ResetSuggestions() {
	lbd.suggestions = nil
	lbd.sugIdx = 0
}

func (lbd *BufferData) Assign(suggestions []string) {
	lbd.suggestions = suggestions
	lbd.sugIdx = 0
}

func (lbd *BufferData) SuggestNext() {
	lbd.sugIdx++
	if lbd.sugIdx >= len(lbd.suggestions) {
		lbd.sugIdx = 0
	}
}

func (lbd *BufferData) SuggestPrev() {
	lbd.sugIdx--
	if lbd.sugIdx < 0 {
		lbd.sugIdx = len(lbd.suggestions) - 1
	}
}

func (lbd *BufferData) CurrentSuggestion() string { return lbd.suggestions[lbd.sugIdx] }

func (lbd *BufferData) TextSuggestion() string {
	if lbd.HasSuggestions() {
		return lbd.CurrentSuggestion()
	}
	return ""
}
