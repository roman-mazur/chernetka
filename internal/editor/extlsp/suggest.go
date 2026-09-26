package extlsp

import (
	"strings"

	"go.lsp.dev/protocol"
)

// maxCandidates limits how deep into the server's ranking suggestions are
// taken from. Candidates further down are mostly fuzzy matches the user
// didn't mean, and they would shrink the common prefix to nothing.
const maxCandidates = 3

// extractSuggestions turns completion items into inline suggestions for the
// cursor at line[:cx]. Only candidates that extend what's already typed can be
// shown inline; the rest are dropped.
func extractSuggestions(items []protocol.CompletionItem, line string, cx int) []suggestion {
	cx = min(cx, len(line))
	typed := identTrailing(line[:cx])
	if len(items) > 0 && typed != "" && isComplete(items[0], typed) {
		// The best candidate is what's already typed: the word is complete,
		// and proposing a longer one is more likely a distraction.
		return nil
	}

	var res []suggestion
	for _, item := range items[:min(maxCandidates, len(items))] {
		text, cursor, ok := completionSuffix(item, line, cx)
		if !ok || containsSuggestion(res, text) {
			continue
		}
		res = append(res, suggestion{
			text:   text,
			cursor: cursor,
			label:  item.Label,
			info:   item.Detail,
			edits:  item.AdditionalTextEdits,
		})
	}
	return withCommonPrefix(res)
}

// rankByPrefix orders completion items of a server that doesn't rank them
// (cue lsp returns every candidate in alphabetical order) for the cursor at
// line[:cx]: the items that are exactly what's already typed come first, as a
// ranking server would put them, then the ones extending it. The rest can't be
// shown inline and are dropped, so they don't push the useful items out of
// the few extractSuggestions looks at.
func rankByPrefix(items []protocol.CompletionItem, line string, cx int) []protocol.CompletionItem {
	cx = min(cx, len(line))
	typed := identTrailing(line[:cx])
	var complete, extending []protocol.CompletionItem
	for _, item := range items {
		if typed != "" && isComplete(item, typed) {
			complete = append(complete, item)
		} else if _, _, ok := completionSuffix(item, line, cx); ok {
			extending = append(extending, item)
		}
	}
	return append(complete, extending...)
}

// withCommonPrefix makes the common prefix of the candidates the first
// suggestion when they disagree, so what is shown by default is right no
// matter which candidate the user is after. The candidates themselves remain
// available to cycle through. When the candidates have nothing in common,
// nothing is suggested: whatever is picked is more likely wrong than right.
func withCommonPrefix(res []suggestion) []suggestion {
	if len(res) < 2 {
		return res
	}
	common := res[0].text
	labels := make([]string, 0, len(res))
	for _, s := range res {
		i := 0
		for i < len(common) && i < len(s.text) && common[i] == s.text[i] {
			i++
		}
		common = common[:i]
		labels = append(labels, s.label)
	}
	if common == "" {
		return nil
	}
	for i, s := range res {
		if s.text == common {
			// One of the candidates is the common part: it goes first.
			copy(res[1:i+1], res[:i])
			res[0] = s
			return res
		}
	}
	head := suggestion{text: common, cursor: len(common), info: strings.Join(labels, " · ")}
	return append([]suggestion{head}, res...)
}

// completionSuffix returns what should be inserted at the cursor to complete
// item, and where in the inserted text the cursor lands. Items that would
// change the text before or after the cursor in a way other than extending it
// are rejected.
func completionSuffix(item protocol.CompletionItem, line string, cx int) (text string, cursor int, ok bool) {
	start := cx - len(identTrailing(line[:cx]))
	if te := item.TextEdit; te != nil {
		if te.Range.Start.Line != te.Range.End.Line ||
			byteOffset(line, te.Range.End.Character) != cx {
			return "", 0, false
		}
		start = byteOffset(line, te.Range.Start.Character)
	}
	if start > cx {
		return "", 0, false
	}
	newText, newCursor := candidateText(item)
	typed := line[start:cx]
	if len(newText) <= len(typed) || !strings.HasPrefix(newText, typed) {
		return "", 0, false
	}
	return newText[len(typed):], max(newCursor-len(typed), 0), true
}

// candidateText is the plain text an item inserts, and the cursor position in
// it after the insertion.
func candidateText(item protocol.CompletionItem) (string, int) {
	text := item.Label
	if te := item.TextEdit; te != nil {
		text = te.NewText
	} else if item.InsertText != "" {
		text = item.InsertText
	}
	if item.InsertTextFormat == protocol.InsertTextFormatSnippet {
		return expandSnippet(text)
	}
	return text, len(text)
}

// isComplete reports whether typed is already the whole item.
func isComplete(item protocol.CompletionItem, typed string) bool {
	text, _ := candidateText(item)
	return item.Label == typed || text == typed
}

func containsSuggestion(list []suggestion, text string) bool {
	for _, s := range list {
		if s.text == text {
			return true
		}
	}
	return false
}
