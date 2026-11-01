package extlsp

import (
	"iter"
	"slices"
	"strings"

	"go.lsp.dev/protocol"
)

// maxCandidates limits how deep into the server's ranking suggestions are
// taken from. Candidates further down are mostly fuzzy matches the user
// didn't mean, and they would shrink the common prefix to nothing.
const maxCandidates = 3

// minWordTyped is how many characters of a word must be typed before the
// words of the document are suggested.
const minWordTyped = 2

// completionInput is what the suggestions at the cursor are made of.
type completionInput struct {
	items []protocol.CompletionItem
	line  string
	cx    int

	// ranked is set if the items are ranked by relevance, not just by the
	// typed prefix (see rankByPrefix).
	ranked bool

	// params are the names of the call parameters from the one at the cursor
	// on, if the cursor is in the arguments of a call.
	params []string

	// text is the document, to suggest its words from if the server has no
	// candidates.
	text string
}

// suggestionsFor makes the inline suggestions at the cursor.
//
// Where an argument of a call starts, the variables named as the parameters
// are suggested: calls like f(ctx, name) are common in Go. The server ranks
// the candidates there by their type only, which picks the right one too
// rarely to be shown on its own.
//
// While typing an identifier, the server candidates are suggested (see
// extractSuggestions). Where the server has nothing, like in strings and
// comments or for the names being declared, the most frequent word of the
// document extending what's typed is suggested.
func suggestionsFor(in completionInput) []suggestion {
	in.cx = min(in.cx, len(in.line))
	before := in.line[:in.cx]
	if afterArgStart(before) {
		return argSuggestions(in.items, in.params, in.line, in.cx)
	}
	if len(in.items) == 0 {
		return wordSuggestions(in.text, before)
	}
	return extractSuggestions(in.items, in.line, in.cx, in.ranked)
}

// extractSuggestions turns completion items into inline suggestions for the
// cursor at line[:cx]. Only candidates that extend what's already typed can be
// shown inline; the rest are dropped.
//
// When the best candidates disagree, their common prefix is suggested first
// (see withCommonPrefix). When they have nothing in common, the best ranked
// one is, with the others to cycle to: unless nothing of the word is typed
// yet, as after a selector dot, or the server doesn't rank the candidates.
// Then the first one is just the first in the alphabetical order.
func extractSuggestions(items []protocol.CompletionItem, line string, cx int, ranked bool) []suggestion {
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
	if len(res) < 2 {
		return res
	}
	if common := withCommonPrefix(res); common != nil || !ranked || typed == "" {
		return common
	}
	return res // The best ranked first, the others to cycle to.
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
// available to cycle through. When the candidates have nothing in common, it
// returns nil.
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

// afterArgStart reports whether the cursor is where an argument of a call
// may start: right after an opening bracket, or after a comma and a space.
func afterArgStart(before string) bool {
	return strings.HasSuffix(before, "(") || strings.HasSuffix(before, ", ")
}

// argSuggestions suggests the variables named as the parameters at the cursor:
// as many of them in a row as there are, separated with commas, then the first
// one alone to cycle to.
func argSuggestions(items []protocol.CompletionItem, params []string, line string, cx int) []suggestion {
	var (
		args  []string
		found []protocol.CompletionItem
	)
	for _, p := range params {
		i := slices.IndexFunc(items, func(it protocol.CompletionItem) bool {
			return it.Label == p && isValueKind(it.Kind)
		})
		if i < 0 {
			break
		}
		args = append(args, p)
		found = append(found, items[i])
	}
	if len(args) == 0 {
		return nil
	}
	if first, _, ok := completionSuffix(found[0], line, cx); !ok || first != args[0] {
		return nil // Not inserted as is, like with an address taken.
	}
	all := strings.Join(args, ", ")
	res := []suggestion{{text: all, cursor: len(all), label: all, info: found[0].Detail}}
	if len(args) > 1 {
		res[0].info = "arguments"
		res = append(res, suggestion{text: args[0], cursor: len(args[0]), label: args[0], info: found[0].Detail})
	}
	return res
}

// isValueKind reports whether a completion item of the kind can be passed as
// an argument as is.
func isValueKind(k protocol.CompletionItemKind) bool {
	switch k {
	case protocol.CompletionItemKindVariable, protocol.CompletionItemKindField, protocol.CompletionItemKindConstant:
		return true
	}
	return false
}

// paramNames returns the names of the call parameters from the active one on.
// It stops at the first unnamed parameter.
func paramNames(help *protocol.SignatureHelp) []string {
	if help == nil || int(help.ActiveSignature) >= len(help.Signatures) {
		return nil
	}
	sig := help.Signatures[help.ActiveSignature]
	active := int(help.ActiveParameter)
	if n := len(sig.Parameters); active >= n {
		if n == 0 || !strings.Contains(sig.Parameters[n-1].Label, "...") {
			return nil
		}
		active = n - 1 // The rest of the arguments are variadic.
	}
	var names []string
	for _, p := range sig.Parameters[active:] {
		// Like "ctx context.Context", or a type alone if unnamed.
		name, _, ok := strings.Cut(p.Label, " ")
		if !ok || name == "_" || !isIdentifier(name) {
			break
		}
		names = append(names, name)
	}
	return names
}

func isIdentifier(s string) bool {
	return s != "" && !startsWithDigit(s) && identTrailing(s) == s
}

// wordSuggestions suggests the most frequent of the words of text extending
// the one typed at the end of before, once minWordTyped characters of it are
// typed. Of the equally frequent words, the first in the alphabetical order
// is suggested.
func wordSuggestions(text, before string) []suggestion {
	typed := identTrailing(before)
	if len(typed) < minWordTyped || startsWithDigit(typed) {
		return nil
	}
	counts := map[string]int{}
	for w := range identsSeq(text) {
		if len(w) > len(typed) && strings.HasPrefix(w, typed) {
			counts[w]++
		}
	}
	best, n := "", 0
	for w, c := range counts {
		if c > n || (c == n && w < best) {
			best, n = w, c
		}
	}
	if best == "" {
		return nil
	}
	rest := best[len(typed):]
	return []suggestion{{text: rest, cursor: len(rest), label: best, info: "word"}}
}

// identsSeq yields the identifiers of text, and the numbers.
func identsSeq(text string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for i := 0; i < len(text); {
			if !isIdentByte(text[i]) {
				i++
				continue
			}
			j := i
			for j < len(text) && isIdentByte(text[j]) {
				j++
			}
			if !yield(text[i:j]) {
				return
			}
			i = j
		}
	}
}
