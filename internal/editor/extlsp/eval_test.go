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
// intact) and a completion is requested at c. A strategy maps the result to a
// ghost text which is scored against what the author actually typed next.

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
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
}

// candidate extracts ghost-text candidates (best first) from completion items.
type candidate func(items []protocol.CompletionItem, line string, cx int) []string

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
	Line  string // full ground truth line
	Cx    int
	Items []protocol.CompletionItem
	Lat   time.Duration
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

	var strategies []evalStrategy
	for _, s := range sessions {
		strategies = append(strategies,
			evalStrategy{s.name + "/current", s.name, trigAlways, extractLabel},
			evalStrategy{s.name + "/identdot", s.name, trigIdentOrDot, extractLabel},
			evalStrategy{s.name + "/identdot+edit", s.name, trigIdentOrDot, extractEdit},
			evalStrategy{s.name + "/prefix1+edit", s.name, trigPrefixOrDot(1), extractEdit},
			evalStrategy{s.name + "/prefix2+edit", s.name, trigPrefixOrDot(2), extractEdit},
			evalStrategy{s.name + "/prefix2+edit+nodot", s.name, trigPrefix(2), extractEdit},
			evalStrategy{s.name + "/p2d+edit+exact", s.name, trigPrefixOrDot(2), exact(extractEdit)},
			evalStrategy{s.name + "/p2d+edit+top3", s.name, trigPrefixOrDot(2), topN(3, extractEdit)},
			evalStrategy{s.name + "/p2d+edit+top3+exact", s.name, trigPrefixOrDot(2), exact(topN(3, extractEdit))},
			evalStrategy{s.name + "/p1d+edit+top3+exact", s.name, trigPrefixOrDot(1), exact(topN(3, extractEdit))},
			evalStrategy{s.name + "/p2+edit+top3+exact", s.name, trigPrefix(2), exact(topN(3, extractEdit))},
			evalStrategy{s.name + "/p2d+edit+lcp", s.name, trigPrefixOrDot(2), lcp(extractEdit)},
			evalStrategy{s.name + "/p1d+edit+top5+exact+lcp", s.name, trigPrefixOrDot(1), lcp(exact(topN(5, extractEdit)))},
			evalStrategy{s.name + "/dotonly+edit", s.name, trigDot, extractEdit},
			evalStrategy{s.name + "/FINAL", s.name, trigFinal, extractFinal},
			evalStrategy{s.name + "/FINAL+snip", s.name, trigFinal, snippets(extractFinal)},
			evalStrategy{s.name + "/snip+p1d+top3+exact", s.name, trigPrefixOrDot(1), snippets(exact(topN(3, extractEdit)))},
			evalStrategy{s.name + "/snip+p2d+top3+exact", s.name, trigPrefixOrDot(2), snippets(exact(topN(3, extractEdit)))},
			evalStrategy{s.name + "/snip+p1d+top3+exact+lcp", s.name, trigPrefixOrDot(1), snippets(lcp(exact(topN(3, extractEdit))))},
			evalStrategy{s.name + "/snip+p1d+top5+exact+lcp", s.name, trigPrefixOrDot(1), snippets(lcp(exact(topN(5, extractEdit))))},
			evalStrategy{s.name + "/snip+p2d+top5+exact+lcp", s.name, trigPrefixOrDot(2), snippets(lcp(exact(topN(5, extractEdit))))},
			evalStrategy{s.name + "/snip+p1+top5+exact+lcp", s.name, trigPrefix(1), snippets(lcp(exact(topN(5, extractEdit))))},
			evalStrategy{s.name + "/dotonly+edit+lcp", s.name, trigDot, lcp(extractEdit)},
		)
	}

	fmt.Printf("\n%-28s %7s %7s %7s %7s %7s %7s %7s\n",
		"strategy", "shown%", "prec%", "noise", "save1%", "save3%", "wrong", "lat")
	for _, st := range strategies {
		pts := results[st.session]
		if pts == nil {
			continue
		}
		m := score(st, pts)
		fmt.Printf("%-28s %7.1f %7.1f %7.2f %7.1f %7.1f %7d %7s\n",
			st.name, m.shownPct, m.precision, m.noisePer10, m.save1, m.save3, m.wrong, m.p50lat)
	}
	fmt.Println(`
shown%: typing positions with a visible suggestion (after the trigger fired)
prec%:  of shown, the ones that are a prefix of what was actually typed next
noise:  wrong suggestions shown per 10 typed characters
save1%: keystrokes saved by a user who accepts the first suggestion with Tab when it is right
save3%: same, but also cycles with ↓ to the 2nd/3rd candidate (each ↓ costs a keystroke)
lat:    p50 completion latency`)

	if *evalVerbose && len(strategies) > 0 {
		traceSession(strategies, results)
	}
}

type metrics struct {
	shownPct, precision, noisePer10, save1, save3 float64
	wrong                                         int
	p50lat                                        time.Duration
}

func score(st evalStrategy, lines [][]point) metrics {
	var (
		positions, shown, right, wrong int
		typed, cost1, cost3            int
		lats                           []time.Duration
	)
	for _, pts := range lines {
		// Per position ghost candidates.
		ghosts := make([][]string, len(pts))
		for i, p := range pts {
			lats = append(lats, p.Lat)
			if !st.trigger(p.Line[:p.Cx], p.Cx) {
				continue
			}
			ghosts[i] = st.extract(p.Items, p.Line[:p.Cx], p.Cx)
			positions++
			rest := p.Line[p.Cx:]
			if len(ghosts[i]) > 0 {
				shown++
				if isUseful(ghosts[i][0], rest) {
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
		cost1 += simulate(line, start, pts, ghosts, 1)
		cost3 += simulate(line, start, pts, ghosts, 3)
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
		m.noisePer10 = 10 * float64(wrong) / float64(typed)
		m.save1 = 100 * (1 - float64(cost1)/float64(typed))
		m.save3 = 100 * (1 - float64(cost3)/float64(typed))
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

// simulate returns the number of keystrokes needed to type line from start
// for a user who accepts a correct suggestion among the first topK.
func simulate(line string, start int, pts []point, ghosts [][]string, topK int) int {
	cost := 0
	for c := start; c < len(line); {
		i := c - start - 1 // ghosts at cursor c were computed for pts[c-start-1]
		accepted := false
		if i >= 0 && i < len(ghosts) {
			for k, g := range ghosts[i] {
				if k >= topK {
					break
				}
				if isUseful(g, line[c:]) {
					cost += 1 + k // ↓ k times, then Tab
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

func traceSession(strategies []evalStrategy, results map[string][][]point) {
	st := strategies[0]
	for _, pts := range results[st.session][:min(10, len(results[st.session]))] {
		for _, p := range pts {
			g := st.extract(p.Items, p.Line[:p.Cx], p.Cx)
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
func topN(n int, next candidate) candidate {
	return func(items []protocol.CompletionItem, line string, cx int) []string {
		return next(items[:min(n, len(items))], line, cx)
	}
}

// exact shows nothing when the typed word is already the best ranked item.
func exact(next candidate) candidate {
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
func lcp(next candidate) candidate {
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
func snippets(next candidate) candidate {
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

// extractFinal is what the integration ships. A suggestion is scored up to
// where it places the cursor: the closing bracket after it is auto-paired
// by the editor anyway.
func extractFinal(items []protocol.CompletionItem, line string, cx int) []string {
	var res []string
	for _, s := range extractSuggestions(items, line, cx) {
		if s.cursor > 0 {
			res = append(res, s.text[:s.cursor])
		}
	}
	return res
}

func trigFinal(line string, _ int) bool { return shouldComplete(line) }

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
			cancel()
			pts = append(pts, point{Line: full, Cx: cx, Items: items, Lat: time.Since(t0)})
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
	h := editor.NewTestHarness()
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
	h := editor.NewTestHarness()
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
	h := editor.NewTestHarness()
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
	h := editor.NewTestHarness()
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
	h := editor.NewTestHarness()
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
		select {
		case cmd := <-h.Commands():
			return cmd.(*editor.GoTo)
		case <-time.After(30 * time.Second):
			t.Fatal("no definition found")
			return nil
		}
	}

	if got := find(6, "func main() { f"); got.Path != path || got.Pos != (content.Position{Line: 4, Col: len("func ")}) {
		t.Errorf("foo is defined at %s:%v", got.Path, got.Pos)
	}
	got := find(4, "func foo() string { return strings.To")
	fmt.Printf("strings.ToUpper is defined at %s:%v\n", got.Path, got.Pos)
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
	h := editor.NewTestHarness()
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
