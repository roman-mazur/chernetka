package extlsp

// Offline quality evaluation of inline completion against a real gopls.
//
//	CHE_LSP_EVAL=1 go test -run TestEvalCompletion -v ./internal/editor/extlsp/ -timeout 60m
//
// The tests in this file talk to a real gopls: they are slow and depend on
// the environment, so they are skipped unless CHE_LSP_EVAL is set.
//
// Each sampled line of a real Go file is "retyped" character by character: at
// every column c the document holds the line truncated to c (everything else
// intact) and a completion is requested at c (with the signature help in the
// sessions asking for it). A strategy maps the result to a ghost text which is
// scored against what the author actually typed next.
//
// The production integration is the FINAL strategy of the prod session; the
// others are what it evolved from, and the alternatives tried. To score some
// of them only, break them down by where the keystrokes go, and check
// another sample of lines:
//
//	CHE_LSP_EVAL=1 go test -run 'TestEvalCompletion$' -v ./internal/editor/extlsp/ -timeout 60m \
//		-args -eval.sessions=prod -eval.cache=/tmp/eval -eval.strategies='/(ident|FINAL)$' \
//		-eval.breakdown=FINAL -eval.seed=2 -eval.lines=150
//
// Findings on this repository (2 samples of 120 and 150 lines):
//   - The server latency (~15ms) is far below the typing pace: the late
//     completions barely matter here (saveT%, saveTR%).
//   - Where an argument starts, the server ranks by the type only, and its best
//     candidate is right ~10% of the time; the variables named as the
//     parameters are right ~95% of the time, though they are rarely there.
//     The only variable of the parameter type is right ~65% of the time.
//   - When the best candidates disagree, the best one is right ~70% of the
//     time if it's a variable, a constant, a keyword, or a type, and ~20%
//     otherwise.
//   - Ranking the candidates by their occurrences in the document makes it
//     worse.
//   - The server has no candidates in a fifth of the identifier positions:
//     strings, trailing comments, the names being declared. The words of the
//     document help there.

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"maps"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"
	"unicode"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/lsp"
)

// skipUnlessEval skips a test using a real language server unless asked for.
func skipUnlessEval(t *testing.T) {
	t.Helper()
	if os.Getenv("CHE_LSP_EVAL") == "" {
		t.Skip("set CHE_LSP_EVAL=1 to run tests against a real gopls")
	}
}

var (
	evalLines   = flag.Int("eval.lines", 120, "number of sampled lines")
	evalSeed    = flag.Int64("eval.seed", 1, "sampling seed")
	evalFilter  = flag.String("eval.sessions", "", "comma-separated session names to run (all when empty)")
	evalCache   = flag.String("eval.cache", "", "directory caching raw gopls responses per session")
	evalVerbose = flag.Bool("eval.trace", false, "print per-position trace for the first session")
)

type evalSession struct {
	name  string
	opts  lsp.Options
	stale bool // the document lags one keystroke behind the completion request
	sig   bool // the signature help is requested too
}

// candidate extracts ghost-text candidates (best first) for a typing position.
type candidate func(in evalInput) []string

// evalInput is what is known at a typing position when the completion arrives.
type evalInput struct {
	Items []protocol.CompletionItem
	Line  string // the line up to the cursor
	Cx    int
	Sig   *protocol.SignatureHelp
	Words map[string]int // identifiers of the document, except the line being typed
	Doc   string         // the document, except the line being typed
}

// count is how many times word occurs in the document.
func (in evalInput) count(word string) int {
	before := in.Line[:len(in.Line)-len(identTrailing(in.Line))] // without the word being typed
	return in.Words[word] + countWord(before, word)
}

// itemsCandidate extracts ghost-text candidates from completion items only.
type itemsCandidate func(items []protocol.CompletionItem, line string, cx int) []string

func lift(f itemsCandidate) candidate {
	return func(in evalInput) []string { return f(in.Items, in.Line, in.Cx) }
}

// trigger reports whether a completion request is issued after typing line[:cx].
type trigger func(line string, cx int) bool

type evalStrategy struct {
	name    string
	session string
	trigger trigger
	extract candidate
}

// point is a completion result for a single typing position.
type point struct {
	Line   string // full ground truth line
	Cx     int
	Items  []protocol.CompletionItem
	Lat    time.Duration
	Sig    *protocol.SignatureHelp `json:",omitempty"` // only in the sessions asking for it
	SigLat time.Duration           `json:",omitempty"`
}

// lat is how long the integration waits for the server at the point: the
// signature help is asked for where the arguments start only.
func (p point) lat() time.Duration {
	if afterArgStart(p.Line[:p.Cx]) {
		return p.Lat + p.SigLat
	}
	return p.Lat
}

func TestEvalCompletion(t *testing.T) {
	skipUnlessEval(t)
	root, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	samples := sampleLines(t, root, *evalLines, *evalSeed)
	t.Logf("sampled %d lines", len(samples))

	deep := map[string]any{"deepCompletion": false}
	sessions := []evalSession{
		{name: "race", opts: lsp.Options{}, stale: true},
		{name: "default", opts: lsp.Options{}},
		{name: "casesens", opts: lsp.Options{InitializationOptions: map[string]any{"matcher": "CaseSensitive"}}},
		{name: "nodeep", opts: lsp.Options{InitializationOptions: deep}},
		{name: "budget30", opts: lsp.Options{InitializationOptions: map[string]any{"completionBudget": "30ms"}}},
		{name: "snippets", opts: lsp.Options{SnippetSupport: true}},
		{name: "snipnodeep", opts: lsp.Options{SnippetSupport: true, InitializationOptions: deep}},
		// What the integration runs: snippets, and the signature help at the arguments.
		{name: "prod", opts: lsp.Options{SnippetSupport: true}, sig: true},
	}
	if *evalFilter != "" {
		keep := map[string]bool{}
		for n := range strings.SplitSeq(*evalFilter, ",") {
			keep[n] = true
		}
		var filtered []evalSession
		for _, s := range sessions {
			if keep[s.name] {
				filtered = append(filtered, s)
			}
		}
		sessions = filtered
	}

	results := make(map[string][][]point) // session -> line -> positions
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)
	for _, s := range sessions {
		wg.Go(func() {
			pts := cachedSession(t, root, s, samples)
			mu.Lock()
			results[s.name] = pts
			mu.Unlock()
		})
	}
	wg.Wait()

	// What is typed elsewhere in the documents, to rank by.
	docs := make([]evalDoc, len(samples))
	for i, sm := range samples {
		docs[i].text = withLine(strings.Split(sm.text, "\n"), sm.line, "")
		docs[i].words = map[string]int{}
		for w := range identsSeq(docs[i].text) {
			docs[i].words[w]++
		}
	}

	var strategies []evalStrategy
	for _, s := range sessions {
		add := func(name string, trig trigger, extract candidate) {
			strategies = append(strategies, evalStrategy{s.name + "/" + name, s.name, trig, extract})
		}
		// The history: what the integration evolved from.
		add("current", trigAlways, lift(extractLabel))
		add("identdot", trigIdentOrDot, lift(extractLabel))
		add("identdot+edit", trigIdentOrDot, lift(extractEdit))
		add("prefix1+edit", trigPrefixOrDot(1), lift(extractEdit))
		add("prefix2+edit", trigPrefixOrDot(2), lift(extractEdit))
		add("prefix2+edit+nodot", trigPrefix(2), lift(extractEdit))
		add("p2d+edit+exact", trigPrefixOrDot(2), lift(exact(extractEdit)))
		add("p2d+edit+top3", trigPrefixOrDot(2), lift(topN(3, extractEdit)))
		add("p2d+edit+top3+exact", trigPrefixOrDot(2), lift(exact(topN(3, extractEdit))))
		add("p1d+edit+top3+exact", trigPrefixOrDot(1), lift(exact(topN(3, extractEdit))))
		add("p2+edit+top3+exact", trigPrefix(2), lift(exact(topN(3, extractEdit))))
		add("p2d+edit+lcp", trigPrefixOrDot(2), lift(lcp(extractEdit)))
		add("p1d+edit+top5+exact+lcp", trigPrefixOrDot(1), lift(lcp(exact(topN(5, extractEdit)))))
		add("dotonly+edit", trigDot, lift(extractEdit))
		add("dotonly+edit+lcp", trigDot, lift(lcp(extractEdit)))
		add("snip+p1d+top3+exact", trigPrefixOrDot(1), lift(snippets(exact(topN(3, extractEdit)))))
		add("snip+p2d+top3+exact", trigPrefixOrDot(2), lift(snippets(exact(topN(3, extractEdit)))))
		add("snip+p1d+top5+exact+lcp", trigPrefixOrDot(1), lift(snippets(lcp(exact(topN(5, extractEdit))))))
		add("snip+p2d+top5+exact+lcp", trigPrefixOrDot(2), lift(snippets(lcp(exact(topN(5, extractEdit))))))
		add("snip+p1+top5+exact+lcp", trigPrefix(1), lift(snippets(lcp(exact(topN(5, extractEdit))))))
		// What the integration shipped before the arguments, the kinds, and
		// the words of the document.
		add("ident", trigIdent, lift(extractIdent))
		// When the best candidates have nothing in common: the best one if
		// enough of the word is typed, or if it's of the given kinds.
		add("ident+top1", trigIdent, lift(top1IfDisagree(1, extractIdent)))
		add("ident+top1p2", trigIdent, lift(top1IfDisagree(2, extractIdent)))
		add("ident+top1p3", trigIdent, lift(top1IfDisagree(3, extractIdent)))
		vck := []protocol.CompletionItemKind{protocol.CompletionItemKindVariable, protocol.CompletionItemKindConstant, protocol.CompletionItemKindKeyword}
		vckc := append(slices.Clone(vck), protocol.CompletionItemKindClass)
		add("ident+top1vck", trigIdent, lift(top1IfKind(vck, extractIdent)))
		add("ident+top1vckc", trigIdent, lift(top1IfKind(vckc, extractIdent)))
		// Ranking by the occurrences in the document.
		add("ident+freq5", trigIdent, freqRank(5, lift(extractIdent)))
		add("ident+freq10", trigIdent, freqRank(10, lift(extractIdent)))
		// The words of the document when nothing extending the typed word is
		// suggested (wordsN), or when the server has no candidates (wordsEN).
		add("ident+words2", trigIdent, wordsFallback(2, lift(extractIdent)))
		add("ident+words3", trigIdent, wordsFallback(3, lift(extractIdent)))
		add("ident+wordsE1", trigIdent, wordsIfNoItems(1, lift(extractIdent)))
		add("ident+wordsE2", trigIdent, wordsIfNoItems(2, lift(extractIdent)))
		add("ident+wordsE3", trigIdent, wordsIfNoItems(3, lift(extractIdent)))
		// The candidate named as the parameter first, and the arguments named
		// as the parameters (all of them in a row, or the first one).
		add("ident+param", trigIdent, paramFirst(lift(extractIdent)))
		add("args", trigArgs, argsOnly)
		add("args1", trigArgs, argsFirst)
		add("args+type", trigArgs, argsOrTyped)
		add("ident+args", trigIdentOrArgs, paramFirst(lift(extractIdent)))
		// The combinations.
		add("ident+args+top1+wordsE2", trigIdentOrArgs, paramFirst(wordsIfNoItems(2, lift(top1IfDisagree(1, extractIdent)))))
		add("ident+args+top1vck+wordsE2", trigIdentOrArgs, paramFirst(wordsIfNoItems(2, lift(top1IfKind(vck, extractIdent)))))
		add("ident+args+top1vckc+wordsE2", trigIdentOrArgs, paramFirst(wordsIfNoItems(2, lift(top1IfKind(vckc, extractIdent)))))
		add("FINAL", trigFinal, extractFinal)
	}
	if *evalStrategies != "" {
		re := regexp.MustCompile(*evalStrategies)
		strategies = slices.DeleteFunc(strategies, func(st evalStrategy) bool { return !re.MatchString(st.name) })
	}

	fmt.Printf("\n%-32s %7s %7s %7s %7s %7s %7s %7s %7s %7s\n",
		"strategy", "shown%", "prec%", "noise", "save1%", "save3%", "saveT%", "saveTR%", "wrong", "lat")
	for _, st := range strategies {
		pts := results[st.session]
		if pts == nil {
			continue
		}
		m := score(st, pts, docs)
		fmt.Printf("%-32s %7.1f %7.1f %7.2f %7.1f %7.1f %7.1f %7.1f %7d %7s\n",
			st.name, m.shownPct, m.precision, m.noisePer10, m.save1, m.save3, m.saveT, m.saveTR, m.wrong, m.p50lat)
	}
	fmt.Printf(`
shown%%:  typing positions with a visible suggestion (after the trigger fired)
prec%%:   of shown, the ones that are a prefix of what was actually typed next
noise:   wrong suggestions shown per 10 typed characters
save1%%:  keystrokes saved by a user who accepts the first suggestion with Tab when it is right
save3%%:  same, but also cycles with ↓ to the 2nd/3rd candidate (each ↓ costs a keystroke)
saveT%%:  save1%% for a user typing a key every %s, with the completions taking as long
         as they took: the ones arriving after the next key are dropped
saveTR%%: same, but the late completions are shown if they agree with what was typed since
lat:     p50 completion latency (with the signature help when it's asked for)
`, *evalTyping)

	if *evalBreakdown != "" {
		re := regexp.MustCompile(*evalBreakdown)
		for _, st := range strategies {
			if pts := results[st.session]; pts != nil && re.MatchString(st.name) {
				printBreakdown(st, pts, docs)
			}
		}
	}
	if *evalVerbose && len(strategies) > 0 {
		traceSession(strategies, results, docs)
	}
}

var (
	evalTyping     = flag.Duration("eval.typing", 120*time.Millisecond, "interval between keystrokes in the latency-aware simulation")
	evalStrategies = flag.String("eval.strategies", "", "regexp of the strategies to score (all when empty)")
	evalBreakdown  = flag.String("eval.breakdown", "", "regexp of the strategies to break down by the typing context")
)

type metrics struct {
	shownPct, precision, noisePer10 float64
	save1, save3, saveT, saveTR     float64
	wrong                           int
	p50lat                          time.Duration
}

// lineGhosts are the ghost candidates of a strategy at the typing positions of
// a line: ghosts[i] is for pts[i], and nil if the trigger didn't fire.
type lineGhosts struct {
	pts       []point
	ghosts    [][]string
	triggered []bool
}

// evalDoc is the document of a sampled line, except the line.
type evalDoc struct {
	text  string
	words map[string]int // the identifiers of text
}

func (d evalDoc) input(p point) evalInput {
	return evalInput{Items: p.Items, Line: p.Line[:p.Cx], Cx: p.Cx, Sig: p.Sig, Words: d.words, Doc: d.text}
}

func ghostsOf(st evalStrategy, pts []point, doc evalDoc) lineGhosts {
	lg := lineGhosts{pts: pts, ghosts: make([][]string, len(pts)), triggered: make([]bool, len(pts))}
	for i, p := range pts {
		if !st.trigger(p.Line[:p.Cx], p.Cx) {
			continue
		}
		lg.triggered[i] = true
		lg.ghosts[i] = st.extract(doc.input(p))
	}
	return lg
}

func score(st evalStrategy, lines [][]point, docs []evalDoc) metrics {
	var (
		positions, shown, right, wrong int
		typed, cost1, cost3            int
		costT, costTR                  int
		lats                           []time.Duration
	)
	for li, pts := range lines {
		lg := ghostsOf(st, pts, docs[li])
		for i, p := range pts {
			lats = append(lats, p.lat())
			if !lg.triggered[i] {
				continue
			}
			positions++
			if g := lg.ghosts[i]; len(g) > 0 {
				shown++
				if isUseful(g[0], p.Line[p.Cx:]) {
					right++
				} else {
					wrong++
				}
			}
		}
		if len(pts) == 0 {
			continue
		}
		line := pts[0].Line
		start := pts[0].Cx - 1 // pts[0] is right after the first typed character
		typed += len(line) - start
		cost1 += simulate(line, start, lg, 1, nil)
		cost3 += simulate(line, start, lg, 3, nil)
		costT += simulateTimed(line, start, lg, *evalTyping, false)
		costTR += simulateTimed(line, start, lg, *evalTyping, true)
	}
	slices.Sort(lats)
	var m metrics
	if positions > 0 {
		m.shownPct = 100 * float64(shown) / float64(positions)
	}
	if shown > 0 {
		m.precision = 100 * float64(right) / float64(shown)
	}
	if typed > 0 {
		saved := func(cost int) float64 { return 100 * (1 - float64(cost)/float64(typed)) }
		m.noisePer10 = 10 * float64(wrong) / float64(typed)
		m.save1, m.save3, m.saveT, m.saveTR = saved(cost1), saved(cost3), saved(costT), saved(costTR)
	}
	m.wrong = wrong
	if len(lats) > 0 {
		m.p50lat = lats[len(lats)/2].Round(time.Millisecond)
	}
	return m
}

func isUseful(ghost, rest string) bool {
	return ghost != "" && strings.HasPrefix(rest, ghost)
}

// contextStats are the counts of a typing context, see typingContext.
type contextStats struct {
	typed, saved        int // characters typed in the context, keystrokes saved by accepting there
	shown, right, wrong int
}

// simulate returns the number of keystrokes needed to type line from start
// for a user who accepts a correct suggestion among the first topK. The
// completions are there right away. If stats is set, it gets the counts of
// the typing contexts.
func simulate(line string, start int, lg lineGhosts, topK int, stats map[string]*contextStats) int {
	cost := 0
	for c := start; c < len(line); {
		i := c - start - 1 // ghosts at cursor c were computed for pts[c-start-1]
		var st *contextStats
		if stats != nil {
			ctx := typingContext(line[:c])
			if st = stats[ctx]; st == nil {
				st = &contextStats{}
				stats[ctx] = st
			}
			st.typed++
		}
		accepted := false
		if i >= 0 && i < len(lg.ghosts) {
			if st != nil && len(lg.ghosts[i]) > 0 {
				st.shown++
				if isUseful(lg.ghosts[i][0], line[c:]) {
					st.right++
				} else {
					st.wrong++
				}
			}
			for k, g := range lg.ghosts[i] {
				if k >= topK {
					break
				}
				if isUseful(g, line[c:]) {
					cost += 1 + k // ↓ k times, then Tab
					if st != nil {
						st.saved += len(g) - 1 - k
						// The accepted characters count as typed in this context.
						st.typed += len(g) - 1
					}
					c += len(g)
					accepted = true
					break
				}
			}
		}
		if !accepted {
			cost++
			c++
		}
	}
	return cost
}

// simulateTimed is simulate (with the first suggestion only) for a user typing
// a key every interval while the completions take as long as they took in the
// session. Like the integration does, one completion is requested at a time:
// the edits made meanwhile wait, and only the latest of them gets a completion
// when the server is done. A suggestion stays visible while what's typed
// agrees with it, until a newer completion replaces it. A completion arriving
// after the cursor moved on is dropped, unless rebase is set: then it's shown
// if it agrees with what was typed since it was requested.
func simulateTimed(line string, start int, lg lineGhosts, interval time.Duration, rebase bool) int {
	type request struct {
		at    int // the cursor position the completion is for
		ready time.Duration
	}
	var (
		cost      int
		now       time.Duration
		visible   []string // the suggestions anchored at visibleAt
		visibleAt int
		inflight  *request
		queued    = -1 // the latest edit waiting for the server to be done
	)
	idx := func(c int) int { return c - start - 1 }
	send := func(c int, at time.Duration) {
		p := lg.pts[idx(c)]
		inflight = &request{at: c, ready: at + p.lat()}
	}
	// typedThrough is what's left of the suggestions anchored at `at` after
	// the user typed up to c.
	typedThrough := func(ghosts []string, at, c int) []string {
		var res []string
		for _, g := range ghosts {
			if typed := line[at:c]; len(g) > len(typed) && strings.HasPrefix(g, typed) {
				res = append(res, g[len(typed):])
			}
		}
		return res
	}
	for c := start; c < len(line); {
		for inflight != nil && inflight.ready <= now {
			r := *inflight
			inflight = nil
			switch ghosts := lg.ghosts[idx(r.at)]; {
			case r.at == c:
				visible, visibleAt = ghosts, c
			case rebase && len(typedThrough(ghosts, r.at, c)) > 0:
				visible, visibleAt = ghosts, r.at
			}
			if queued >= 0 {
				send(queued, r.ready)
				queued = -1
			}
		}

		if cur := typedThrough(visible, visibleAt, c); len(cur) > 0 && isUseful(cur[0], line[c:]) {
			cost++
			c += len(cur[0])
			visible = nil
		} else {
			cost++
			c++
		}

		if i := idx(c); i >= 0 && i < len(lg.pts) && lg.triggered[i] {
			if inflight == nil {
				send(c, now)
			} else {
				queued = c
			}
		} else {
			// No completion for this edit, and the one waiting is superseded.
			visible, queued = nil, -1
		}
		now += interval
	}
	return cost
}

// typingContext classifies a typing position by what's before the cursor.
func typingContext(before string) string {
	p := identTrailing(before)
	switch {
	case p != "" && startsWithDigit(p):
		return "number"
	case p != "":
		return "ident"
	case strings.TrimSpace(before) == "":
		return "line start"
	case strings.HasSuffix(before, "."):
		return "after ."
	case strings.HasSuffix(before, "("):
		return "after ("
	case strings.HasSuffix(before, ", "):
		return "after ,"
	case strings.HasSuffix(before, " "):
		return "after space"
	}
	return "other"
}

// printBreakdown shows where the keystrokes go and where they are saved
// (from simulate with the first suggestion).
func printBreakdown(st evalStrategy, lines [][]point, docs []evalDoc) {
	stats := map[string]*contextStats{}
	total := 0
	for li, pts := range lines {
		if len(pts) == 0 {
			continue
		}
		lg := ghostsOf(st, pts, docs[li])
		line := pts[0].Line
		start := pts[0].Cx - 1
		total += len(line) - start
		simulate(line, start, lg, 1, stats)
	}
	names := slices.Sorted(maps.Keys(stats))
	fmt.Printf("\n%s by the typing context:\n%-12s %7s %7s %7s %7s\n", st.name, "context", "typed%", "saved%", "shown", "prec%")
	for _, n := range names {
		s := stats[n]
		prec := 0.0
		if s.shown > 0 {
			prec = 100 * float64(s.right) / float64(s.shown)
		}
		fmt.Printf("%-12s %7.1f %7.1f %7d %7.1f\n", n,
			100*float64(s.typed)/float64(total), 100*float64(s.saved)/float64(total), s.shown, prec)
	}
	fmt.Println("typed%: share of the typed characters, saved%: share of all the keystrokes saved there")
}

func traceSession(strategies []evalStrategy, results map[string][][]point, docs []evalDoc) {
	st := strategies[0]
	for li, pts := range results[st.session][:min(10, len(results[st.session]))] {
		for _, p := range pts {
			g := st.extract(docs[li].input(p))
			var labels []string
			for _, it := range p.Items[:min(3, len(p.Items))] {
				te := ""
				if it.TextEdit != nil {
					te = it.TextEdit.NewText
				}
				labels = append(labels, it.Label+"|"+te)
			}
			fmt.Printf("%q ▮ ghost=%q items=%v\n", p.Line[:p.Cx], g, labels)
		}
	}
}

// --- strategies ---------------------------------------------------------

func trigAlways(line string, _ int) bool { return strings.TrimSpace(line) != "" }

func trigIdentOrDot(line string, _ int) bool {
	if line == "" {
		return false
	}
	r := rune(line[len(line)-1])
	return r == '.' || isIdent(r)
}

func trigPrefixOrDot(n int) trigger {
	return func(line string, cx int) bool {
		if strings.HasSuffix(line, ".") {
			return true
		}
		return trigPrefix(n)(line, cx)
	}
}

func trigPrefix(n int) trigger {
	return func(line string, _ int) bool {
		p := identTrailing(line)
		return len(p) >= n && !unicode.IsDigit(rune(p[0]))
	}
}

func trigDot(line string, _ int) bool { return strings.HasSuffix(line, ".") }

// topN only looks at the n best ranked items.
func topN(n int, next itemsCandidate) itemsCandidate {
	return func(items []protocol.CompletionItem, line string, cx int) []string {
		return next(items[:min(n, len(items))], line, cx)
	}
}

// exact shows nothing when the typed word is already the best ranked item.
func exact(next itemsCandidate) itemsCandidate {
	return func(items []protocol.CompletionItem, line string, cx int) []string {
		if len(items) > 0 {
			p := identTrailing(line)
			if p != "" && (items[0].Label == p || items[0].FilterText == p) {
				return nil
			}
		}
		return next(items, line, cx)
	}
}

// lcp first offers the longest common prefix of all the candidates.
func lcp(next itemsCandidate) itemsCandidate {
	return func(items []protocol.CompletionItem, line string, cx int) []string {
		res := next(items, line, cx)
		if len(res) < 2 {
			return res
		}
		common := res[0]
		for _, r := range res[1:] {
			i := 0
			for i < len(common) && i < len(r) && common[i] == r[i] {
				i++
			}
			common = common[:i]
		}
		if common == "" {
			return nil
		}
		if common == res[0] {
			return res
		}
		return append([]string{common}, res...)
	}
}

// snippets turns snippet candidates into plain text up to the first tab stop:
// "Println(${1:})" becomes "Println(" which is what the user types before
// the arguments.
func snippets(next itemsCandidate) itemsCandidate {
	return func(items []protocol.CompletionItem, line string, cx int) []string {
		cp := make([]protocol.CompletionItem, len(items))
		for i, it := range items {
			if it.InsertTextFormat == protocol.InsertTextFormatSnippet && it.TextEdit != nil {
				te := *it.TextEdit
				te.NewText = snippetPrefix(te.NewText)
				it.TextEdit = &te
			}
			cp[i] = it
		}
		return next(cp, line, cx)
	}
}

func snippetPrefix(s string) string {
	if before, _, ok := strings.Cut(s, "$"); ok {
		return before
	}
	return s
}

func isIdent(r rune) bool { return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r) }

// extractLabel is the original, label based logic.
func extractLabel(items []protocol.CompletionItem, line string, cx int) []string {
	prefix := identTrailing(line[:min(cx, len(line))])
	var res []string
	for _, item := range items {
		if item.Label == "" || (prefix != "" && !strings.HasPrefix(item.Label, prefix)) || len(prefix) == len(item.Label) {
			continue
		}
		res = append(res, item.Label[len(prefix):])
	}
	return res
}

// extractEdit uses the item's text edit: the replaced range starts somewhere
// before the cursor, and the new text must extend what's already there.
func extractEdit(items []protocol.CompletionItem, line string, cx int) []string {
	var res []string
	for _, it := range items {
		if it.TextEdit == nil {
			continue
		}
		start := int(it.TextEdit.Range.Start.Character) // ASCII in practice
		if start > cx {
			continue
		}
		typed := line[start:cx]
		if !strings.HasPrefix(it.TextEdit.NewText, typed) || len(it.TextEdit.NewText) == len(typed) {
			continue
		}
		res = append(res, it.TextEdit.NewText[len(typed):])
	}
	return res
}

// ghostTexts are the suggestions scored up to where they place the cursor:
// the closing bracket after it is auto-paired by the editor anyway.
func ghostTexts(sugs []suggestion) []string {
	var res []string
	for _, s := range sugs {
		if s.cursor > 0 {
			res = append(res, s.text[:s.cursor])
		}
	}
	return res
}

// extractIdent is the logic shipped for identifiers and selectors: the top
// candidates extending the typed word, their common prefix first.
func extractIdent(items []protocol.CompletionItem, line string, cx int) []string {
	return ghostTexts(extractSuggestions(items, line, cx, false))
}

// trigIdent is when the integration asked for completions before the
// arguments were suggested: while typing an identifier and after a dot.
func trigIdent(line string, _ int) bool {
	if strings.HasSuffix(line, ".") {
		p := identTrailing(line[:len(line)-1])
		return p != "" && !startsWithDigit(p)
	}
	p := identTrailing(line)
	return p != "" && !startsWithDigit(p)
}

// top1IfDisagree shows the best candidate when the candidates have nothing in
// common, if at least minTyped characters of the word are typed.
func top1IfDisagree(minTyped int, next itemsCandidate) itemsCandidate {
	return func(items []protocol.CompletionItem, line string, cx int) []string {
		res := next(items, line, cx)
		if len(res) > 0 || len(identTrailing(line)) < minTyped {
			return res
		}
		for _, it := range items[:min(maxCandidates, len(items))] {
			if g := next([]protocol.CompletionItem{it}, line, cx); len(g) > 0 {
				return g
			}
		}
		return nil
	}
}

// top1IfKind shows the best candidate of the given kinds when the candidates
// have nothing in common.
func top1IfKind(kinds []protocol.CompletionItemKind, next itemsCandidate) itemsCandidate {
	return func(items []protocol.CompletionItem, line string, cx int) []string {
		res := next(items, line, cx)
		if len(res) > 0 || identTrailing(line) == "" {
			return res
		}
		for _, it := range items[:min(maxCandidates, len(items))] {
			if g := next([]protocol.CompletionItem{it}, line, cx); len(g) > 0 {
				if slices.Contains(kinds, it.Kind) {
					return g
				}
				return nil
			}
		}
		return nil
	}
}

// freqRank orders the n best ranked items by how often they occur in the
// document; the server ranking breaks the ties.
func freqRank(n int, next candidate) candidate {
	return func(in evalInput) []string {
		top := slices.Clone(in.Items[:min(n, len(in.Items))])
		slices.SortStableFunc(top, func(a, b protocol.CompletionItem) int {
			return in.count(b.Label) - in.count(a.Label)
		})
		in.Items = append(top, in.Items[len(top):]...)
		return next(in)
	}
}

// wordsFallback suggests the most frequent word of the document extending
// the typed word when the server has nothing, once minTyped characters of
// it are typed.
func wordsFallback(minTyped int, next candidate) candidate {
	return func(in evalInput) []string {
		if res := next(in); len(res) > 0 {
			return res
		}
		typed := identTrailing(in.Line)
		if len(typed) < minTyped || startsWithDigit(typed) {
			return nil
		}
		best, n := "", 0
		for w, c := range in.Words {
			if len(w) > len(typed) && strings.HasPrefix(w, typed) && (c > n || (c == n && w < best)) {
				best, n = w, c
			}
		}
		if best == "" {
			return nil
		}
		return []string{best[len(typed):]}
	}
}

// wordsIfNoItems is wordsFallback only when the server has no candidates at
// all, like in strings or for the names being declared.
func wordsIfNoItems(minTyped int, next candidate) candidate {
	fallback := wordsFallback(minTyped, func(evalInput) []string { return nil })
	return func(in evalInput) []string {
		if len(in.Items) == 0 {
			return fallback(in)
		}
		return next(in)
	}
}

func trigIdentOrArgs(line string, cx int) bool { return trigIdent(line, cx) || afterArgStart(line) }

// paramFirst puts the candidate named as the parameter at the cursor first.
// Right after an opening bracket or a comma, it's the only one suggested.
func paramFirst(next candidate) candidate {
	return func(in evalInput) []string {
		if identTrailing(in.Line) == "" && !strings.HasSuffix(in.Line, ".") {
			return argsOnly(in)
		}
		if names := paramNames(in.Sig); len(names) > 0 {
			in.Items = moveFirst(in.Items, names[0])
		}
		return next(in)
	}
}

// argsOnly suggests the arguments named as the parameters at the cursor.
func argsOnly(in evalInput) []string {
	return ghostTexts(argSuggestions(in.Items, paramNames(in.Sig), in.Line, in.Cx))
}

// argsFirst only suggests the argument at the cursor, not the ones after it.
func argsFirst(in evalInput) []string {
	return ghostTexts(argSuggestions(in.Items, paramNames(in.Sig)[:min(1, len(paramNames(in.Sig)))], in.Line, in.Cx))
}

// argsOrTyped is argsOnly, or if no variable is named as the parameter, the
// only variable of the parameter type.
func argsOrTyped(in evalInput) []string {
	if res := argsOnly(in); len(res) > 0 {
		return res
	}
	if in.Sig == nil || int(in.Sig.ActiveSignature) >= len(in.Sig.Signatures) {
		return nil
	}
	sig := in.Sig.Signatures[in.Sig.ActiveSignature]
	if int(in.Sig.ActiveParameter) >= len(sig.Parameters) {
		return nil
	}
	_, typ, ok := strings.Cut(sig.Parameters[in.Sig.ActiveParameter].Label, " ")
	if !ok {
		return nil
	}
	var match []string
	for _, it := range in.Items {
		if it.Kind == protocol.CompletionItemKindVariable && it.Detail == typ {
			match = append(match, it.Label)
		}
	}
	if len(match) != 1 {
		return nil
	}
	return match
}

func trigArgs(line string, _ int) bool { return afterArgStart(line) }

// moveFirst moves the item labeled as given to the front.
func moveFirst(items []protocol.CompletionItem, label string) []protocol.CompletionItem {
	i := slices.IndexFunc(items, func(it protocol.CompletionItem) bool { return it.Label == label })
	if i <= 0 {
		return items
	}
	res := slices.Clone(items)
	copy(res[1:i+1], items[:i])
	res[0] = items[i]
	return res
}

// countWord counts the occurrences of word in text as a whole identifier.
func countWord(text, word string) int {
	n := 0
	for w := range identsSeq(text) {
		if w == word {
			n++
		}
	}
	return n
}

// extractFinal is what the integration ships.
func extractFinal(in evalInput) []string {
	return ghostTexts(suggestionsFor(completionInput{
		items:  in.Items,
		line:   in.Line,
		cx:     in.Cx,
		params: paramNames(in.Sig),
		ranked: true,
		text:   in.Doc,
	}))
}

func trigFinal(line string, _ int) bool { return shouldComplete(line) }

// --- plumbing -----------------------------------------------------------

type sample struct {
	path  string
	text  string
	line  int
	start int // column where typing starts (after indentation)
}

func sampleLines(t *testing.T, root string, n int, seed int64) []sample {
	var files []string
	_ = filepath.Walk(filepath.Join(root, "internal"), func(p string, info os.FileInfo, err error) error {
		if err == nil && strings.HasSuffix(p, ".go") && !strings.HasSuffix(p, "eval_test.go") {
			files = append(files, p)
		}
		return nil
	})
	_ = filepath.Walk(filepath.Join(root, "cmd"), func(p string, info os.FileInfo, err error) error {
		if err == nil && strings.HasSuffix(p, ".go") {
			files = append(files, p)
		}
		return nil
	})
	sort.Strings(files)

	type cand struct {
		file string
		ln   int
	}
	var all []cand
	texts := map[string]string{}
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		texts[f] = string(b)
		for i, l := range strings.Split(string(b), "\n") {
			tl := strings.TrimSpace(l)
			// Code lines only, ASCII to keep column math simple.
			if len(tl) < 8 || strings.HasPrefix(tl, "//") || strings.HasPrefix(tl, "import") ||
				strings.HasPrefix(tl, "package") || strings.HasPrefix(tl, "\"") || !isASCII(l) {
				continue
			}
			all = append(all, cand{f, i})
		}
	}
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(all), func(i, j int) { all[i], all[j] = all[j], all[i] })
	var res []sample
	for _, c := range all[:min(n, len(all))] {
		l := strings.Split(texts[c.file], "\n")[c.ln]
		res = append(res, sample{c.file, texts[c.file], c.ln, len(l) - len(strings.TrimLeft(l, " \t"))})
	}
	return res
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 0x80 {
			return false
		}
	}
	return true
}

// cachedSession runs a session appending every line result to a JSONL cache
// file, so an interrupted evaluation resumes where it stopped.
func cachedSession(t *testing.T, root string, s evalSession, samples []sample) [][]point {
	var (
		res  [][]point
		out  *os.File
		path string
	)
	if *evalCache != "" {
		path = filepath.Join(*evalCache, fmt.Sprintf("%s-%d.jsonl", s.name, *evalSeed))
		if b, err := os.ReadFile(path); err == nil {
			for l := range strings.SplitSeq(string(b), "\n") {
				var pts []point
				if l != "" && json.Unmarshal([]byte(l), &pts) == nil {
					res = append(res, pts)
				}
			}
		}
		var err error
		if out, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644); err != nil {
			t.Fatal(err)
		}
		defer out.Close()
	}
	if len(res) >= len(samples) {
		return res[:len(samples)]
	}

	sessionSlots <- struct{}{}
	defer func() { <-sessionSlots }()

	const restartEvery = 20 // gopls slows down with many versions of the files
	for from := len(res); from < len(samples); from += restartEvery {
		batch := samples[from:min(from+restartEvery, len(samples))]
		for _, pts := range runSession(t, root, s, batch, from) {
			res = append(res, pts)
			if out != nil {
				b, _ := json.Marshal(pts)
				_, _ = out.Write(append(b, '\n'))
			}
		}
	}
	return res
}

var sessionSlots = make(chan struct{}, 3)

func runSession(t *testing.T, root string, s evalSession, samples []sample, offset int) [][]point {
	ctx := context.Background()
	c, err := lsp.StartWith(ctx, root, s.opts)
	if err != nil {
		t.Fatalf("%s: %v", s.name, err)
	}
	defer c.Shutdown(ctx)

	var (
		res    [][]point
		opened = map[string]int32{}
	)
	for si, sm := range samples {
		started := time.Now()
		u := uri.File(sm.path)
		if _, ok := opened[sm.path]; !ok {
			if err := c.DidOpen(ctx, u, "go", sm.text, 1); err != nil {
				t.Fatal(err)
			}
			opened[sm.path] = 1
		}
		lines := strings.Split(sm.text, "\n")
		full := lines[sm.line]
		var pts []point
		for cx := sm.start + 1; cx <= len(full); cx++ {
			docCx := cx
			if s.stale {
				docCx = cx - 1
			}
			doc := withLine(lines, sm.line, full[:docCx])
			opened[sm.path]++
			_ = c.DidChange(ctx, u, opened[sm.path], lsp.TextChange{Text: doc})
			t0 := time.Now()
			cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			items, _ := c.Completion(cctx, u, uint32(sm.line), uint32(cx))
			pt := point{Line: full, Cx: cx, Items: items, Lat: time.Since(t0)}
			if s.sig {
				t0 = time.Now()
				pt.Sig, _ = c.SignatureHelp(cctx, u, uint32(sm.line), uint32(cx))
				pt.SigLat = time.Since(t0)
			}
			cancel()
			pts = append(pts, pt)
		}
		// Restore the original content for the following samples.
		opened[sm.path]++
		_ = c.DidChange(ctx, u, opened[sm.path], lsp.TextChange{Text: sm.text})
		res = append(res, pts)
		fmt.Fprintf(os.Stderr, "%s: line %d %s:%d (%d chars) in %s\n", s.name, offset+si,
			filepath.Base(sm.path), sm.line+1, len(full), time.Since(started).Round(time.Millisecond))
	}
	return res
}

func withLine(lines []string, n int, l string) string {
	cp := append([]string(nil), lines...)
	cp[n] = l
	return strings.Join(cp, "\n")
}

// TestRealGopls drives the editor with the production integration and a real
// gopls: type a call, check the ghost text, accept it.
func TestRealGopls(t *testing.T) {
	skipUnlessEval(t)
	dir := t.TempDir()
	_ = os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/x\n\ngo 1.22\n"), 0o644)
	src := "package main\n\nimport \"fmt\"\n\nfunc main() {\n\t\n}\n"
	path := filepath.Join(dir, "main.go")
	_ = os.WriteFile(path, []byte(src), 0o644)

	var le Integration
	h := editor.NewTestHarness(t)
	h.LogEmbed = logger.Embed(logger.Prefix(t.Logf, "editor: "))
	h.Extend(&le)
	if err := h.OpenReader(path, strings.NewReader(src)); err != nil {
		t.Fatal(err)
	}
	defer le.Close()
	buf := h.Top()
	data := buf.ExtensionData(le.ID()).(*BufferData)
	h.Run(t)
	h.Post(t, editor.CommandFunc(func(*editor.Editor) {
		for range 5 {
			editor.RelMove{Dy: 1}.DoOnBuffer(buf, editor.RenderPrefs{})
		}
		h.MoveCursorToLineEnd()
		h.SetMode(editor.ModeInsert)
	}))

	check := func(typed string) {
		h.SendInputSequence(t, typed)
		deadline := time.Now().Add(10 * time.Second)
		var ghost, info string
		for time.Now().Before(deadline) {
			h.Post(t, editor.CommandFunc(func(*editor.Editor) { sug := data.TextSuggestion(); ghost, info = sug.Text, sug.Info }))
			if ghost != "" {
				break
			}
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Printf("typed %-10q ghost=%q info=%q\n", typed, ghost, info)
	}
	check("fmt.Pri")
	check("ntl")
	h.SendInput(t, []byte("\t"))
	var line string
	var cx int
	h.Post(t, editor.CommandFunc(func(*editor.Editor) {
		cx, _ = buf.Pos()
		line = buf.Content.Lines()[5].String()
	}))
	fmt.Printf("after Tab: %q cursor at %d (%q|%q)\n", line, cx, line[:cx], line[cx:])
	if line != "\tfmt.Println()" || cx != len("\tfmt.Println(") {
		t.Errorf("accepted line = %q, cursor %d", line, cx)
	}
}

// TestRealGoplsSuggestsArguments types a call with the production integration
// and a real gopls, and checks the variables named as the parameters are
// suggested as the arguments.
func TestRealGoplsSuggestsArguments(t *testing.T) {
	skipUnlessEval(t)
	dir := t.TempDir()
	_ = os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/x\n\ngo 1.22\n"), 0o644)
	src := "package main\n\nimport \"context\"\n\nfunc greet(ctx context.Context, name string, n int) {}\n\n" +
		"func main() {\n\tctx := context.Background()\n\tname := \"x\"\n\t\n\t_, _ = ctx, name\n}\n"
	path := filepath.Join(dir, "main.go")
	_ = os.WriteFile(path, []byte(src), 0o644)

	var le Integration
	h := editor.NewTestHarness(t)
	h.LogEmbed = logger.Embed(logger.Prefix(t.Logf, "editor: "))
	h.Extend(&le)
	if err := h.OpenReader(path, strings.NewReader(src)); err != nil {
		t.Fatal(err)
	}
	defer le.Close()
	buf := h.Top()
	data := buf.ExtensionData(le.ID()).(*BufferData)
	h.Run(t)
	h.Post(t, editor.CommandFunc(func(*editor.Editor) {
		for range 9 {
			editor.RelMove{Dy: 1}.DoOnBuffer(buf, editor.RenderPrefs{})
		}
		h.MoveCursorToLineEnd()
		h.SetMode(editor.ModeInsert)
	}))
	h.SendInputSequence(t, "greet(")
	var ghost, info string
	for deadline := time.Now().Add(10 * time.Second); time.Now().Before(deadline); time.Sleep(50 * time.Millisecond) {
		h.Post(t, editor.CommandFunc(func(*editor.Editor) { sug := data.TextSuggestion(); ghost, info = sug.Text, sug.Info }))
		if ghost != "" {
			break
		}
	}
	fmt.Printf("typed \"greet(\" ghost=%q info=%q\n", ghost, info)
	if ghost != "ctx, name" {
		t.Errorf("suggested %q, want the arguments", ghost)
	}
}

var goplsLog = flag.String("eval.gopls-log", "", "run a dedicated gopls logging here")

func TestRealGoplsRepoFile(t *testing.T) {
	skipUnlessEval(t)
	path := "../../../cmd/che/main.go"
	src, _ := os.ReadFile(path)
	var le Integration
	if *goplsLog != "" {
		le.Starter = func(ctx context.Context, _, root string) (lspClient, error) {
			return lsp.StartWith(ctx, root, lsp.Options{
				Args:                  append(strings.Fields(os.Getenv("EVAL_GOPLS_FLAGS")), "-rpc.trace", "-logfile", *goplsLog, "serve"),
				SnippetSupport:        true,
				InitializationOptions: map[string]any{"deepCompletion": false},
			})
		}
	}
	h := editor.NewTestHarness(t)
	h.LogEmbed = logger.Embed(t.Logf)
	h.Extend(&le)
	if err := h.OpenReader(path, strings.NewReader(string(src))); err != nil {
		t.Fatal(err)
	}
	defer le.Close()
	buf := h.Top()
	data := buf.ExtensionData(le.ID()).(*BufferData)
	h.Run(t)
	h.Post(t, editor.CommandFunc(func(*editor.Editor) {
		for range 99 {
			editor.RelMove{Dy: 1}.DoOnBuffer(buf, editor.RenderPrefs{})
		}
		h.MoveCursorToLineEnd()
		h.SetMode(editor.ModeInsert)
	}))
	h.SendInput(t, []byte("\r"))
	for _, typed := range []string{"fm", "t.", "Spr"} {
		h.SendInputSequence(t, typed)
		time.Sleep(3 * time.Second)
		var ghost, info string
		h.Post(t, editor.CommandFunc(func(*editor.Editor) { sug := data.TextSuggestion(); ghost, info = sug.Text, sug.Info }))
		fmt.Printf("typed %q ghost=%q info=%q\n", typed, ghost, info)
	}
}

// TestSyncModes compares full and incremental document sync with a real gopls:
// the latency of a didChange followed by a completion, while typing.
func TestSyncModes(t *testing.T) {
	skipUnlessEval(t)
	dir := t.TempDir()
	_ = os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/x\n\ngo 1.22\n"), 0o644)
	gen := func(funcs int) string {
		var sb strings.Builder
		sb.WriteString("package main\n\nimport \"fmt\"\n\n")
		for i := range funcs {
			fmt.Fprintf(&sb, "func f%d(a, b int) int {\n\tif a > b {\n\t\treturn a - b\n\t}\n\tfmt.Println(a, b)\n\treturn a + b\n}\n\n", i)
		}
		sb.WriteString("func main() {\n\t\n}\n")
		return sb.String()
	}
	repoFile, _ := os.ReadFile("../buffer_test.go")
	files := map[string]string{
		"repo_buffer_test.go (21KB)": string(repoFile),
		"gen_5k_lines.go":            gen(600),
		"gen_20k_lines.go":           gen(2500),
	}

	ctx := context.Background()
	c, err := lsp.StartWith(ctx, dir, lsp.Options{SnippetSupport: true})
	if err != nil {
		t.Fatal(err)
	}
	defer c.Shutdown(ctx)

	names := make([]string, 0, len(files))
	for n := range files {
		names = append(names, n)
	}
	sort.Strings(names)
	for i, name := range names {
		text := files[name]
		path := filepath.Join(dir, fmt.Sprintf("f%d.go", i))
		if strings.HasPrefix(name, "repo_") {
			// Without its package it doesn't type check, but it's parsed as usual.
			path = filepath.Join(dir, "repo", "buffer_test.go")
			_ = os.MkdirAll(filepath.Dir(path), 0o755)
		}
		_ = os.WriteFile(path, []byte(text), 0o644)
		u := uri.File(path)
		version := int32(1)
		_ = c.DidOpen(ctx, u, "go", text, version)

		// Type in main (or at the end of the file) and measure each keystroke.
		lines := strings.Split(text, "\n")
		ln := len(lines) - 3
		const typed = "fmt.Println(a)"
		var full, incr []time.Duration
		var fullBytes, incrBytes int
		for round := range 6 {
			incremental := round%2 == 1
			base := text
			for k := 1; k <= len(typed); k++ {
				cur := withLine(lines, ln, lines[ln]+typed[:k])
				change := lsp.TextChange{Text: cur}
				if incremental {
					change = diff(base, cur)
				}
				b, _ := json.Marshal(change)
				version++
				t0 := time.Now()
				_ = c.DidChange(ctx, u, version, change)
				_, _ = c.Completion(ctx, u, uint32(ln), uint32(len(lines[ln])+k))
				d := time.Since(t0)
				if round >= 2 { // skip warm up
					if incremental {
						incr, incrBytes = append(incr, d), incrBytes+len(b)
					} else {
						full, fullBytes = append(full, d), fullBytes+len(b)
					}
				}
				base = cur
			}
			// Back to the original text.
			version++
			_ = c.DidChange(ctx, u, version, lsp.TextChange{Text: text})
		}
		fmt.Printf("%-28s %6d lines  full: p50 %-7s %8d B/keystroke   incremental: p50 %-7s %4d B/keystroke\n",
			name, len(lines), median(full), fullBytes/len(full), median(incr), incrBytes/len(incr))
	}
}

func median(ds []time.Duration) time.Duration {
	s := append([]time.Duration(nil), ds...)
	slices.Sort(s)
	return s[len(s)/2].Round(100 * time.Microsecond)
}

// TestProbeOrganizeImports shows what gopls does to imports with a package
// used without an import, and an unused import.
func TestProbeOrganizeImports(t *testing.T) {
	skipUnlessEval(t)
	dir := t.TempDir()
	_ = os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/x\n\ngo 1.22\n"), 0o644)
	src := "package main\n\nimport (\n\t\"fmt\"\n\t\"io\"\n)\n\nfunc main() {\n\tfmt.Println(strconv.Itoa(1))\n}\n"
	path := filepath.Join(dir, "main.go")
	_ = os.WriteFile(path, []byte(src), 0o644)

	ctx := context.Background()
	c, err := lsp.Start(ctx, dir, lsp.Options{Logf: t.Logf})
	if err != nil {
		t.Fatal(err)
	}
	defer c.Shutdown(ctx)
	u := uri.File(path)
	_ = c.DidOpen(ctx, u, "go", src, 1)

	started := time.Now()
	edits, err := c.OrganizeImports(ctx, u)
	fmt.Printf("organize imports in %s: err=%v\n", time.Since(started).Round(time.Millisecond), err)
	after := src
	for i := len(edits) - 1; i >= 0; i-- {
		fmt.Printf("   %+v %q\n", edits[i].Range, edits[i].NewText)
		after = applyChange(after, lsp.TextChange{Range: &edits[i].Range, Text: edits[i].NewText})
	}
	fmt.Printf("result:\n%s", after)
}

// TestRealGoplsAddsImports types calls of packages that are not imported and
// shows the imports added by the production integration with a real gopls.
func TestRealGoplsAddsImports(t *testing.T) {
	skipUnlessEval(t)
	dir := t.TempDir()
	_ = os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/x\n\ngo 1.22\n"), 0o644)
	src := "package main\n\nimport (\n\t\"fmt\"\n\t\"io\"\n)\n\nfunc main() {\n\t\n}\n"
	path := filepath.Join(dir, "main.go")
	_ = os.WriteFile(path, []byte(src), 0o644)

	var le Integration
	h := editor.NewTestHarness(t)
	h.LogEmbed = logger.Embed(logger.Prefix(t.Logf, "editor: "))
	h.Extend(&le)
	if err := h.OpenReader(path, strings.NewReader(src)); err != nil {
		t.Fatal(err)
	}
	defer le.Close()
	buf := h.Top()
	h.Run(t)
	h.Post(t, editor.CommandFunc(func(*editor.Editor) {
		for range 8 {
			editor.RelMove{Dy: 1}.DoOnBuffer(buf, editor.RenderPrefs{})
		}
		h.MoveCursorToLineEnd()
		h.SetMode(editor.ModeInsert)
	}))
	time.Sleep(3 * time.Second) // let gopls load the module

	h.SendInputSequence(t, "fmt.Println(strconv.Itoa(1")
	time.Sleep(2 * time.Second)
	h.Post(t, editor.CommandFunc(func(*editor.Editor) { h.MoveCursorToLineEnd() }))
	h.SendInputSequence(t, "\r\tn := rand.Intn(10")
	time.Sleep(2 * time.Second)
	var text string
	h.Post(t, editor.CommandFunc(func(*editor.Editor) { text = buf.Text() }))
	fmt.Printf("result:\n%s\n", text)
	for _, want := range []string{"\t\"io\"\n", "\t\"strconv\"\n", "\t\"math/rand\"\n"} {
		if !strings.Contains(text, want) {
			t.Errorf("missing import line %q", want)
		}
	}
}

// TestRealGoplsFormatsOnSave saves a badly formatted file with the production
// integration and a real gopls.
func TestRealGoplsFormatsOnSave(t *testing.T) {
	skipUnlessEval(t)
	dir := t.TempDir()
	_ = os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/x\n\ngo 1.22\n"), 0o644)
	src := "package main\nimport \"fmt\"\nfunc main() {\n  x:=1\n    fmt.Println( x )\n}\n"
	want := "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tx := 1\n\tfmt.Println(x)\n}\n"
	path := filepath.Join(dir, "main.go")
	_ = os.WriteFile(path, []byte(src), 0o644)

	var le Integration
	h := editor.NewTestHarness(t)
	h.LogEmbed = logger.Embed(logger.Prefix(t.Logf, "editor: "))
	h.Extend(&le)
	if err := h.OpenReader(path, strings.NewReader(src)); err != nil {
		t.Fatal(err)
	}
	defer le.Close()
	srv := h.Top().ExtensionData(le.ID()).(*BufferData).srv
	h.Run(t)
	for !srv.ready.Load() {
		time.Sleep(10 * time.Millisecond)
	}

	started := time.Now()
	h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
		(&editor.Save{DstPath: path}).DoOnBuffer(e.Top(), editor.RenderPrefs{})
	}))
	fmt.Printf("saved in %s\n", time.Since(started).Round(time.Millisecond))
	if got, _ := os.ReadFile(path); string(got) != want {
		t.Errorf("saved\n%s\nwant\n%s", got, want)
	}

	// Once the workspace is loaded.
	started = time.Now()
	h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
		(&editor.Save{DstPath: path}).DoOnBuffer(e.Top(), editor.RenderPrefs{})
	}))
	fmt.Printf("saved again in %s\n", time.Since(started).Round(time.Millisecond))
}

// TestRealGoplsGoesToDefinition finds definitions in the same file and in the
// standard library with the production integration and a real gopls.
func TestRealGoplsGoesToDefinition(t *testing.T) {
	skipUnlessEval(t)
	dir := t.TempDir()
	_ = os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/x\n\ngo 1.22\n"), 0o644)
	src := "package main\n\nimport \"strings\"\n\nfunc foo() string { return strings.ToUpper(\"x\") }\n\nfunc main() { foo() }\n"
	path := filepath.Join(dir, "main.go")
	_ = os.WriteFile(path, []byte(src), 0o644)

	var le Integration
	h := editor.NewTestHarness(t)
	h.LogEmbed = logger.Embed(logger.Prefix(t.Logf, "editor: "))
	h.Extend(&le)
	if err := h.OpenReader(path, strings.NewReader(src)); err != nil {
		t.Fatal(err)
	}
	defer le.Close()
	data := h.Top().ExtensionData(le.ID()).(*BufferData)

	find := func(line int, before string) *editor.GoTo {
		t.Helper()
		data.FindDefinition(h.Editor, content.Position{Line: line, Col: len(before)})
		timeout := time.After(30 * time.Second)
		for {
			select {
			case cmd := <-h.Commands():
				if goTo, ok := cmd.(*editor.GoTo); ok {
					return goTo
				}
				// Like a render of the diagnostics published meanwhile.
			case <-timeout:
				t.Fatal("no definition found")
				return nil
			}
		}
	}

	if got := find(6, "func main() { f"); got.Path != path || got.Position != (content.Position{Line: 4, Col: len("func ")}) {
		t.Errorf("foo is defined at %s:%v", got.Path, got.Position)
	}
	got := find(4, "func foo() string { return strings.To")
	fmt.Printf("strings.ToUpper is defined at %s:%v\n", got.Path, got.Position)
	if filepath.Base(got.Path) != "strings.go" {
		t.Errorf("strings.ToUpper is defined in %s", got.Path)
	}
}

// TestRealGoplsDiagnostics opens a file with an error with the production
// integration and a real gopls, and waits for the error to be reported.
func TestRealGoplsDiagnostics(t *testing.T) {
	skipUnlessEval(t)
	dir := t.TempDir()
	_ = os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/x\n\ngo 1.22\n"), 0o644)
	src := "package main\n\nfunc main() {\n\tx := 1\n}\n"
	path := filepath.Join(dir, "main.go")
	_ = os.WriteFile(path, []byte(src), 0o644)

	var le Integration
	h := editor.NewTestHarness(t)
	h.LogEmbed = logger.Embed(logger.Prefix(t.Logf, "editor: "))
	h.Extend(&le)
	if err := h.OpenReader(path, strings.NewReader(src)); err != nil {
		t.Fatal(err)
	}
	defer le.Close()
	data := h.Top().ExtensionData(le.ID()).(*BufferData)
	h.Run(t)

	started := time.Now()
	for {
		diags := onLoop(t, h, data.Diagnostics)
		if len(diags) > 0 {
			fmt.Printf("diagnostics in %s: %+v\n", time.Since(started).Round(time.Millisecond), diags)
			if d := diags[0]; d.Line != 3 || d.Severity != code.SeverityError || !strings.Contains(d.Message, "declared and not used") {
				t.Errorf("diagnostics = %+v", diags)
			}
			return
		}
		if time.Since(started) > 30*time.Second {
			t.Fatal("no diagnostics")
		}
		time.Sleep(20 * time.Millisecond)
	}
}
