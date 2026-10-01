package extlsp

import (
	"fmt"
	"strings"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/editor"
)

type BufferData struct {
	srv    *server // the server of the buffer language
	buf    *editor.Buffer
	docUri uri.URI

	suggestions []suggestion
	sugIdx      int    // what suggestion is picked
	anchor      anchor // where the suggestions apply

	version       int32
	reqCompletion int

	// The server state of the document, only used by the sync loop.
	serverOpen   bool
	serverText   string // the text the server has, to send it only the changes
	serverSynced bool   // false if the server text is unknown
}

// suggestion is a completion candidate shown inline at the cursor.
type suggestion struct {
	text   string              // what gets inserted at the cursor
	cursor int                 // where the cursor lands in the text after the insertion
	label  string              // the full candidate name, e.g. "Println"
	info   string              // details like the candidate's signature
	edits  []protocol.TextEdit // extra edits to apply on accept, e.g. an import
}

// anchor is a cursor position together with the line text around it.
type anchor struct {
	line   string
	cx, cy int
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

// Assign sets plain text suggestions.
func (lbd *BufferData) Assign(suggestions []string) {
	res := make([]suggestion, len(suggestions))
	for i, s := range suggestions {
		res[i] = suggestion{text: s, cursor: len(s)}
	}
	lbd.assign(res, lbd.anchor)
}

func (lbd *BufferData) assign(suggestions []suggestion, at anchor) {
	lbd.suggestions = suggestions
	lbd.sugIdx = 0
	lbd.anchor = at
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

func (lbd *BufferData) TextSuggestion() code.Suggestion {
	var s code.Suggestion
	if lbd.HasSuggestions() {
		cur := lbd.suggestions[lbd.sugIdx]
		s.Text = cur.text

		parts := make([]string, 0, 3)
		if cur.info != "" {
			parts = append(parts, cur.info)
		}
		if len(cur.edits) > 0 {
			parts = append(parts, "+import")
		}
		if n := len(lbd.suggestions); n > 1 {
			parts = append(parts, fmt.Sprintf("↑↓ %d/%d", lbd.sugIdx+1, n))
		}
		s.Info = strings.Join(parts, "  ")

	}
	return s
}

// typeThrough updates the suggestions after an edit that moved the cursor from
// the anchor to cx on line cy. If the user typed characters matching the
// current suggestion, they are consumed from it; candidates that contradict
// what was typed are dropped. Any other edit invalidates the suggestions.
func (lbd *BufferData) typeThrough(line string, cx, cy int) {
	if !lbd.HasSuggestions() {
		return
	}
	a := lbd.anchor
	if cy != a.cy || cx <= a.cx || cx > len(line) || a.cx > len(a.line) ||
		line[:a.cx] != a.line[:a.cx] || line[cx:] != a.line[a.cx:] {
		lbd.ResetSuggestions()
		return
	}

	typed := line[a.cx:cx]
	current := lbd.suggestions[lbd.sugIdx]
	kept := lbd.suggestions[:0]
	idx := 0
	for _, s := range lbd.suggestions {
		if len(s.text) <= len(typed) || !strings.HasPrefix(s.text, typed) {
			continue
		}
		if s.label == current.label && s.text == current.text {
			idx = len(kept)
		}
		s.text = s.text[len(typed):]
		s.cursor = max(s.cursor-len(typed), 0)
		kept = append(kept, s)
	}
	lbd.suggestions = kept
	lbd.sugIdx = idx
	lbd.anchor = anchor{line: line, cx: cx, cy: cy}
	if len(kept) == 0 {
		lbd.ResetSuggestions()
	}
}

// Format formats the buffer with the language server before it's saved.
func (lbd *BufferData) Format(prefs editor.RenderPrefs) {
	if lbd.srv != nil {
		lbd.srv.format(lbd.buf, lbd, prefs)
	}
}

// FindDefinition makes the server look for the definition of the symbol at pos
// once it has the current text. The editor moves to it when it's found.
func (lbd *BufferData) FindDefinition(loop editor.Sender, pos content.Position) {
	if lbd.srv == nil || lbd.srv.failed.Load() || lbd.docUri == "" {
		return
	}
	lines := lbd.buf.Content.Lines()
	if pos.Line >= len(lines) {
		return
	}
	lbd.version++
	lbd.srv.queue.push(syncReq{
		editor:     loop,
		buf:        lbd.buf,
		bufData:    lbd,
		text:       lbd.buf.Text(),
		version:    lbd.version,
		definition: &definitionReq{at: lspPosition(lines[pos.Line].String(), pos)},
	})
}

// Diagnostics returns the latest problems the server found in the buffer.
func (lbd *BufferData) Diagnostics() []code.Diagnostic {
	if lbd.srv == nil || lbd.docUri == "" {
		return nil
	}
	return lbd.srv.diagnostics(lbd.docUri)
}

var (
	_ editor.CodeAssist          = new(BufferData) // enforce code assist interface implementation
	_ editor.Formatter           = new(BufferData)
	_ editor.DefinitionFinder    = new(BufferData)
	_ editor.DiagnosticsProvider = new(BufferData)
)
