package editor_test

import (
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
	"testing/synctest"
	"time"
	"unicode/utf8"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/txtar"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/symbol"
	"rmazur.io/chernetka/internal/editor"
)

const (
	// simHangTimeout is how long a whole simulation may take before the editor is considered hung.
	simHangTimeout = 2 * time.Minute
	// simClickDelay separates the clicks so they are not taken for a double click.
	simClickDelay = 600 * time.Millisecond
	// simTraceLen is the number of the last gestures reported on a failure.
	simTraceLen = 24
	// simMaxUndos limits the undos to get back to the text before a hunk.
	simMaxUndos = 512
)

// simDefaultSeeds are the simulations of each pair run by a plain go test.
var simDefaultSeeds = []uint64{1, 2, 3}

// TestSimulatedEditing edits real texts the way a person does. Each txtar archive in testdata/sim
// holds a file before and after a change, mostly taken from the history of this repository,
// as before.<ext> and after.<ext>. The extension picks the extensions of the editor engaged,
// and the comment of the archive says where the text comes from. The simulation opens the before file, and turns it into the after file with the keys
// and the mouse: it goes through the changed hunks in a random order, and picks a random
// way to do each one. Between the hunks it scrolls, saves, undoes, and cuts and pastes.
//
// After each gesture the text and the cursor must match a plain model of the expected text.
// The end result must match the after file, in the buffer and once saved, undoing everything
// must give the before file, and rendering it must match what a new editor shows for that file.
//
// Everything the typist does is decided ahead of the editor, with two exceptions that keep the
// editor the source of truth. The indentation after Enter is read back, and the cursor is
// read back after the gestures not modelled, like undo and scrolling.
//
// Set CHESIM_SEED to repeat one simulation, and CHESIM_RUNS to run that many random ones:
//
//	CHESIM_RUNS=100 go test -run TestSimulatedEditing ./internal/editor
func TestSimulatedEditing(t *testing.T) {
	seeds := simSeeds(t)
	archives, err := filepath.Glob("testdata/sim/*.txtar")
	if err != nil {
		t.Fatal(err)
	}
	if len(archives) == 0 {
		t.Fatal("no simulation texts")
	}
	for _, archive := range archives {
		t.Run(strings.TrimSuffix(filepath.Base(archive), ".txtar"), func(t *testing.T) {
			pair := loadSimPair(t, archive)
			for _, seed := range seeds {
				t.Run(fmt.Sprintf("seed=%d", seed), func(t *testing.T) { pair.run(t, seed) })
			}
		})
	}
}

func simSeeds(t *testing.T) []uint64 {
	if s := os.Getenv("CHESIM_SEED"); s != "" {
		seed, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			t.Fatalf("bad CHESIM_SEED: %s", err)
		}
		return []uint64{seed}
	}
	s := os.Getenv("CHESIM_RUNS")
	if s == "" {
		return simDefaultSeeds
	}
	runs, err := strconv.Atoi(s)
	if err != nil {
		t.Fatalf("bad CHESIM_RUNS: %s", err)
	}
	seeds := make([]uint64, runs)
	for i := range seeds {
		seeds[i] = rand.Uint64()
	}
	return seeds
}

// simPair is the text of a file before and after the editing.
type simPair struct {
	ext           string
	before, after string
}

func loadSimPair(t *testing.T, archive string) simPair {
	ar, err := txtar.ParseFile(archive)
	if err != nil {
		t.Fatal(err)
	}
	if len(ar.Files) != 2 {
		t.Fatalf("%s has %d files, want before.<ext> and after.<ext>", archive, len(ar.Files))
	}
	before, after := ar.Files[0], ar.Files[1]
	ext := filepath.Ext(before.Name)
	if before.Name != "before"+ext || after.Name != "after"+ext {
		t.Fatalf("%s has files %s and %s, want before.<ext> and after.<ext>", archive, before.Name, after.Name)
	}
	return simPair{ext: ext, before: string(before.Data), after: string(after.Data)}
}

func (p simPair) run(t *testing.T, seed uint64) {
	dir := t.TempDir()
	stopHangCheck := failOnHang(simHangTimeout)
	defer stopHangCheck()
	synctest.Test(t, func(t *testing.T) {
		h, path := openHarnessFile(t, dir, p.ext, p.before)
		h.Run(t)
		sim := newSimulator(t, h, path, p, seed)
		sim.editHunks()
		sim.finish()
		h.CloseInput()
	})
}

// simHunk replaces the del lines from line with ins.
type simHunk struct {
	line, del int
	ins       []string
}

// diffLines returns the hunks turning a into b, in the order of the lines.
func diffLines(a, b []string) []simHunk {
	// lcs[i][j] is the length of the longest common subsequence of a[i:] and b[j:].
	lcs := make([][]int, len(a)+1)
	for i := range lcs {
		lcs[i] = make([]int, len(b)+1)
	}
	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {
			if a[i] == b[j] {
				lcs[i][j] = lcs[i+1][j+1] + 1
			} else {
				lcs[i][j] = max(lcs[i+1][j], lcs[i][j+1])
			}
		}
	}
	var (
		hunks []simHunk
		open  bool // the last hunk ends at the current lines
		i, j  int
	)
	add := func() *simHunk {
		if !open {
			hunks = append(hunks, simHunk{line: i})
			open = true
		}
		return &hunks[len(hunks)-1]
	}
	for i < len(a) || j < len(b) {
		switch {
		case i < len(a) && j < len(b) && a[i] == b[j]:
			open = false
			i, j = i+1, j+1
		case j < len(b) && (i == len(a) || lcs[i][j+1] >= lcs[i+1][j]):
			h := add()
			h.ins = append(h.ins, b[j])
			j++
		default:
			add().del++
			i++
		}
	}
	return hunks
}

// simulator is the person editing the file, and what they expect to see in the editor.
type simulator struct {
	t    *testing.T
	h    *editor.TestHarness
	path string
	pair simPair
	seed uint64
	rnd  *rand.Rand

	// The expected state of the editor.
	lines  []string
	cursor content.Position

	autoClosers int      // closing brackets the editor added right after the cursor
	pending     []byte   // the input of the current gesture
	steps       int      // gestures so far
	trace       []string // the last gestures
}

func newSimulator(t *testing.T, h *editor.TestHarness, path string, p simPair, seed uint64) *simulator {
	return &simulator{
		t: t, h: h, path: path, pair: p, seed: seed,
		rnd:   rand.New(rand.NewPCG(seed, 0)),
		lines: strings.Split(p.before, "\n"),
	}
}

// chance is true with the probability of p.
func (s *simulator) chance(p float64) bool { return s.rnd.Float64() < p }

// simState is what the editor holds.
type simState struct {
	text   string
	cursor content.Position
	mode   editor.Mode
}

func (s *simulator) read() simState {
	s.t.Helper()
	synctest.Wait()
	if !s.h.Running() {
		s.fatalf("the editor quit")
	}
	var st simState
	s.h.Post(s.t, editor.CommandFunc(func(e *editor.Editor) {
		buf := e.Top()
		st.text = buf.Text()
		st.cursor.Col, st.cursor.Line = buf.Pos()
		st.mode = buf.Mode()
	}))
	return st
}

func (s *simulator) fatalf(format string, args ...any) {
	s.t.Helper()
	s.t.Fatalf("%s\nCHESIM_SEED=%d, step %d, the last gestures:\n  %s",
		fmt.Sprintf(format, args...), s.seed, s.steps, strings.Join(s.trace, "\n  "))
}

// flush sends the keys of the gesture to the editor without checking the result.
func (s *simulator) flush(name string) {
	s.t.Helper()
	s.steps++
	s.trace = append(s.trace, fmt.Sprintf("%d %s %q", s.steps, name, s.pending))
	if len(s.trace) > simTraceLen {
		s.trace = s.trace[1:]
	}
	if len(s.pending) > 0 {
		s.h.WriteInput(s.t, s.pending)
		s.pending = s.pending[:0]
	}
}

// done ends a gesture: the editor must show the text and the cursor expected.
func (s *simulator) done(name string) {
	s.t.Helper()
	s.flush(name)
	st := s.read()
	s.checkText(st)
	if st.cursor != s.cursor {
		s.fatalf("cursor is at %+v, want %+v", st.cursor, s.cursor)
	}
}

func (s *simulator) checkText(st simState) {
	s.t.Helper()
	if diff := cmp.Diff(strings.Join(s.lines, "\n"), st.text); diff != "" {
		s.fatalf("unexpected text (-want +got):\n%s", diff)
	}
}

// doneAnywhere ends a gesture leaving the cursor where the editor puts it.
func (s *simulator) doneAnywhere(name string) {
	s.t.Helper()
	s.flush(name)
	st := s.read()
	s.checkText(st)
	s.cursor = st.cursor
}

func (s *simulator) key(in string) { s.pending = append(s.pending, in...) }

// Typing in the insert mode. The expected text follows insertInput.

func (s *simulator) line() string { return s.lines[s.cursor.Line] }

func (s *simulator) insertAtCursor(text string, advance int) {
	line := s.line()
	s.lines[s.cursor.Line] = line[:s.cursor.Col] + text + line[s.cursor.Col:]
	s.cursor.Col += advance
}

// repeatedBracket is true when typing ch steps over it, as it closes the pair at the cursor.
func (s *simulator) repeatedBracket(ch byte, opening byte) bool {
	line, col := s.line(), s.cursor.Col
	return 0 < col && col < len(line) && line[col] == ch && line[col-1] == opening
}

// typeRune types r, returning the number of closing brackets the editor adds.
func (s *simulator) typeRune(r rune) (closers int) {
	s.key(string(r))
	switch r {
	case '{', '(', '[':
		s.insertAtCursor(string(r)+string(closingOf(r)), 1)
		return 1
	case '}', ')', ']':
		if s.repeatedBracket(byte(r), byte(openingOf(r))) {
			s.cursor.Col++
			return 0
		}
	case '"', '\'', '`':
		if s.repeatedBracket(byte(r), byte(r)) {
			s.cursor.Col++
			return 0
		}
		s.insertAtCursor(string(r)+string(r), 1)
		return 1
	}
	s.insertAtCursor(string(r), utf8.RuneLen(r))
	return 0
}

func closingOf(r rune) rune {
	return map[rune]rune{'{': '}', '(': ')', '[': ']'}[r]
}

func openingOf(r rune) rune {
	return map[rune]rune{'}': '{', ')': '(', ']': '['}[r]
}

func (s *simulator) backspace() {
	s.key("\x7f")
	line, c := s.line(), s.cursor
	if c.Col > 0 {
		start := symbol.PrevClusterStart(line, c.Col)
		s.lines[c.Line] = line[:start] + line[c.Col:]
		s.cursor.Col = start
		return
	}
	prev := s.lines[c.Line-1]
	s.lines[c.Line-1] = prev + line
	s.lines = slices.Delete(s.lines, c.Line, c.Line+1)
	s.cursor = content.Position{Line: c.Line - 1, Col: len(prev)}
}

func (s *simulator) arrowRight() {
	s.key("\x1b[C")
	s.stepRight()
}

func (s *simulator) stepRight() { s.cursor.Col = symbol.NextClusterEnd(s.line(), s.cursor.Col) }

func (s *simulator) stepLeft() { s.cursor.Col = symbol.PrevClusterStart(s.line(), s.cursor.Col) }

// typeText types the text of a line like a person with the bracket pairing on: it steps over
// the closing brackets added by the editor, and removes those not needed.
func (s *simulator) typeText(text string) {
	for cluster := range symbol.Clusters(text, 1) {
		if s.chance(0.03) {
			s.typeRune(rune('a' + s.rnd.IntN(26))) // A typo.
			s.backspace()
		}
		if line := s.line(); s.autoClosers > 0 && s.cursor.Col < len(line) && line[s.cursor.Col:s.cursor.Col+1] == cluster {
			s.arrowRight()
			s.autoClosers--
			continue
		}
		for _, r := range cluster { // A picker sends all the code points of a character at once.
			s.autoClosers += s.typeRune(r)
		}
	}
	for ; s.autoClosers > 0; s.autoClosers-- {
		s.arrowRight()
		s.backspace()
	}
}

// enter splits the line at its end. The editor picks the indentation of the new line,
// which is changed to the indentation of the next line of the text.
func (s *simulator) enter(next string) {
	s.key("\r")
	s.flush("enter")
	st := s.read()
	got := strings.Split(st.text, "\n")
	line := s.cursor.Line
	if st.cursor.Line != line+1 || st.cursor.Line >= len(got) {
		s.fatalf("cursor is at %+v after Enter on line %d", st.cursor, line)
	}
	indent := got[line+1]
	if want := s.indentAfter(line); !strings.HasPrefix(indent, want) {
		s.fatalf("Enter on line %d indented it with %q, want it to start with %q", line, indent, want)
	}
	s.lines[line] = strings.TrimRight(s.lines[line], simBlanks)
	s.lines = slices.Insert(s.lines, line+1, indent)
	s.cursor = content.Position{Line: line + 1, Col: len(indent)}
	if indent != leadingBlanks(indent) {
		s.fatalf("Enter added %q, want only the indentation", indent)
	}
	s.setIndent(leadingBlanks(next))
	s.typeText(strings.TrimLeft(next, simBlanks))
}

// indentAfter returns the indentation Enter must keep at the end of the line: the one of
// the line, or of the closest text above an empty line. A deeper one is not predicted here.
func (s *simulator) indentAfter(line int) string {
	for ; line > 0 && s.lines[line] == ""; line-- {
	}
	return leadingBlanks(s.lines[line])
}

const simBlanks = " \t"

func leadingBlanks(s string) string { return s[:len(s)-len(strings.TrimLeft(s, simBlanks))] }

func (s *simulator) setIndent(want string) {
	have := s.line()[:s.cursor.Col]
	common := 0
	for common < len(have) && common < len(want) && have[common] == want[common] {
		common++
	}
	for range len(have) - common {
		s.backspace()
	}
	for _, r := range want[common:] {
		s.typeRune(r)
	}
}

// esc leaves the insert mode. The cursor lands on the last typed character.
func (s *simulator) esc() {
	s.key("\x1b")
	s.stepLeft()
}

// Moving the cursor in the normal mode.

// moveTo moves the cursor to the position by one of the ways to get there.
func (s *simulator) moveTo(p content.Position) {
	if p == s.cursor {
		return
	}
	switch n := s.rnd.IntN(5); {
	case n == 0 && s.click(p, false):
	case n == 1 && s.searchTo(p):
	case n == 2 && p.Line == 0:
		s.key("g")
		s.column(p, "g")
	case n == 2:
		s.key(":" + strconv.Itoa(p.Line+1) + "\r0")
		s.column(p, "go to line")
	default:
		s.walkTo(p)
	}
}

// walkTo gets to the position with the arrows or the letters.
func (s *simulator) walkTo(p content.Position) {
	up, down := "\x1b[A", "\x1b[B"
	if s.chance(0.5) {
		up, down = "k", "j"
	}
	for d := s.cursor.Line; d < p.Line; d++ {
		s.key(down)
	}
	for d := p.Line; d < s.cursor.Line; d++ {
		s.key(up)
	}
	s.key("0")
	s.column(p, "walk")
}

// column ends a move to the line of p by going to the column. The column the line is reached
// at depends on the line, as "0" stops at the indentation first, so it's read from the editor.
func (s *simulator) column(p content.Position, name string) {
	s.t.Helper()
	s.flush(name)
	st := s.read()
	s.checkText(st)
	if st.cursor.Line != p.Line {
		s.fatalf("cursor is on line %d, want %d", st.cursor.Line, p.Line)
	}
	s.cursor = st.cursor
	s.walkColumn(p.Col)
	s.done("column")
}

// walkColumn moves the cursor along the line.
func (s *simulator) walkColumn(col int) {
	left, right := "\x1b[D", "\x1b[C"
	if s.chance(0.5) {
		left, right = "h", "l"
	}
	for s.cursor.Col < col {
		s.key(right)
		s.stepRight()
	}
	for s.cursor.Col > col {
		s.key(left)
		s.stepLeft()
	}
}

var simWord = regexp.MustCompile(`\p{L}{5,}`)

// searchTo gets to the position by searching for a word in the line that is not anywhere else,
// and going from the word. It returns false if there is no such word.
func (s *simulator) searchTo(p content.Position) bool {
	text := strings.Join(s.lines, "\n")
	for _, m := range simWord.FindAllStringIndex(s.lines[p.Line], -1) {
		word := s.lines[p.Line][m[0]:m[1]]
		if strings.Count(text, word) != 1 {
			continue
		}
		s.key("/" + word + "\r")
		s.cursor = content.Position{Line: p.Line, Col: m[0]}
		s.done("search " + word)
		s.key("\x1b") // Clear the highlighted matches.
		s.walkColumn(p.Col)
		s.done("from the word")
		return true
	}
	return false
}

// click moves the cursor with the mouse if the position is on the screen.
func (s *simulator) click(p content.Position, shift bool) bool {
	synctest.Wait()
	var (
		x, y    int
		visible bool
	)
	s.h.Post(s.t, editor.CommandFunc(func(*editor.Editor) { x, y, visible = s.h.ScreenPosition(p) }))
	if !visible {
		return false
	}
	time.Sleep(simClickDelay)
	button := 0
	if shift {
		button = 4
	}
	s.key(fmt.Sprintf("\x1b[<%d;%d;%dM\x1b[<%d;%d;%dm", button, x, y, button, x, y))
	s.cursor = p
	s.done("click")
	return true
}

// Applying the hunks.

// editHunks makes all the changes, in a random order, with interruptions between them.
func (s *simulator) editHunks() {
	hunks := diffLines(strings.Split(s.pair.before, "\n"), strings.Split(s.pair.after, "\n"))
	s.rnd.Shuffle(len(hunks), func(i, j int) { hunks[i], hunks[j] = hunks[j], hunks[i] })
	for i, hunk := range hunks {
		// The pause keeps the hunk out of the changes undone at once with the ones before.
		paused := s.chance(0.4)
		if paused {
			time.Sleep(fuzzPauseDelay)
		}
		beforeHunk := strings.Join(s.lines, "\n")

		s.applyHunk(hunk)

		delta := len(hunk.ins) - hunk.del
		for j := i + 1; j < len(hunks); j++ {
			if hunks[j].line > hunk.line {
				hunks[j].line += delta
			}
		}
		if paused && s.chance(0.5) {
			s.undoAndRedo(beforeHunk)
		}
		s.interrupt()
	}
}

func (s *simulator) applyHunk(h simHunk) {
	if h.del == 1 && len(h.ins) == 1 {
		s.editLine(h.line, h.ins[0])
		return
	}
	if h.del > 0 {
		s.deleteLines(h.line, h.del)
	}
	if len(h.ins) > 0 {
		s.insertLines(h.line, h.ins)
	}
}

// editLine changes the part of the line that differs from want.
func (s *simulator) editLine(line int, want string) {
	have := s.lines[line]
	start := commonPrefix(have, want)
	end := commonSuffix(have[start:], want[start:])
	gone, typed := have[start:len(have)-end], want[start:len(want)-end]
	s.moveTo(content.Position{Line: line, Col: start})

	selected := false
	if gone != "" {
		if typed != "" && s.chance(0.5) {
			// The typed text replaces the selection.
			for range symbol.Clusters(gone, 1) {
				s.key("\x1b[1;2C")
			}
			s.key("i")
			selected = true
			s.lines[line] = have[:start] + have[start+len(gone):] // The next key removes it.
		} else {
			for range symbol.Clusters(gone, 1) {
				s.key("x")
			}
			s.lines[line] = have[:start] + have[start+len(gone):]
			s.done("delete characters")
		}
	}
	if typed == "" {
		return
	}
	if !selected {
		s.key("i")
	}
	s.typeText(typed)
	s.esc()
	s.done("type in line")
}

// commonPrefix returns the size of the start of a and b made of the same characters.
func commonPrefix(a, b string) int {
	n := 0
	for cluster := range symbol.Clusters(a, 1) {
		end := n + len(cluster)
		if !strings.HasPrefix(b[n:], cluster) || symbol.NextClusterEnd(b, n) != end {
			break
		}
		n = end
	}
	return n
}

// commonSuffix returns the size of the end of a and b made of the same characters.
func commonSuffix(a, b string) int {
	n := 0
	for n < len(a) && n < len(b) {
		aEnd, bEnd := len(a)-n, len(b)-n
		aStart, bStart := symbol.PrevClusterStart(a, aEnd), symbol.PrevClusterStart(b, bEnd)
		if a[aStart:aEnd] != b[bStart:bEnd] {
			break
		}
		n += aEnd - aStart
	}
	return n
}

// deleteLines removes the lines by selecting them or by erasing them from the end.
func (s *simulator) deleteLines(line, n int) {
	start, end := content.Position{Line: line}, content.Position{Line: line + n}
	if n <= 3 && s.chance(0.3) {
		s.moveTo(end)
		s.key("i")
		for s.cursor != start {
			s.backspace()
		}
		s.esc()
		s.done("erase lines")
		return
	}

	s.moveTo(start)
	selectedByClick := s.chance(0.4) && s.click(end, true)
	if !selectedByClick {
		for range n {
			s.key("\x1b[1;2B")
		}
	}
	if s.chance(0.5) {
		s.key("x")
	} else {
		s.key("\x18") // Ctrl+X.
	}
	s.lines = slices.Delete(s.lines, line, line+n)
	s.cursor = start
	s.done("delete selected lines")
}

// insertLines adds the lines before the given one by typing them or by pasting them.
func (s *simulator) insertLines(line int, ins []string) {
	if line == 0 || s.chance(0.2) {
		s.pasteLines(line, ins)
		return
	}
	prev := content.Position{Line: line - 1, Col: len(s.lines[line-1])}
	first := ins[0]
	if s.chance(0.5) {
		s.moveTo(content.Position{Line: line - 1})
		s.key("o")
		s.lines = slices.Insert(s.lines, line, "")
		s.cursor = content.Position{Line: line}
		s.setIndent(leadingBlanks(first))
		s.typeText(strings.TrimLeft(first, simBlanks))
	} else {
		s.moveTo(content.Position{Line: line - 1})
		s.key("A")
		s.cursor = prev
		s.enter(first)
	}
	for _, next := range ins[1:] {
		s.enter(next)
	}
	s.esc()
	s.done("type lines")
}

func (s *simulator) pasteLines(line int, ins []string) {
	s.moveTo(content.Position{Line: line})
	text := strings.Join(ins, "\n") + "\n"
	s.key("\x1b[200~" + text + "\x1b[201~")
	s.lines = slices.Insert(s.lines, line, ins...)
	s.cursor = content.Position{Line: line + len(ins)}
	s.done("paste lines")
}

// undoAndRedo undoes until the text before the hunk, and redoes as many changes.
func (s *simulator) undoAndRedo(beforeHunk string) {
	want := strings.Join(s.lines, "\n")
	undone := 0
	for ; ; undone++ {
		st := s.read()
		if st.text == beforeHunk {
			break
		}
		if undone == simMaxUndos {
			s.fatalf("the text before the hunk is not reached by undoing %d changes", undone)
		}
		s.key("\x1a") // Ctrl+Z.
		s.flush("undo")
	}
	for range undone {
		s.key("\x1b[122;6u") // Ctrl+Shift+Z.
		s.flush("redo")
	}
	st := s.read()
	if diff := cmp.Diff(want, st.text); diff != "" {
		s.fatalf("text after redoing differs (-want +got):\n%s", diff)
	}
	s.cursor = st.cursor
}

// interrupt does something that must keep the text between the changes.
func (s *simulator) interrupt() {
	switch s.rnd.IntN(8) {
	case 0:
		s.save()
	case 1:
		s.key([]string{"g", "G", " ", "b"}[s.rnd.IntN(4)])
		s.doneAnywhere("scroll")
	case 2:
		s.key("\x1b[<65;1;1M\x1b[<64;1;1M") // The mouse wheel.
		s.doneAnywhere("wheel")
	case 3:
		s.cutAndPaste()
	}
}

func (s *simulator) save() {
	s.key(":w\r")
	s.done("save")
	s.checkSaved(strings.Join(s.lines, "\n"))
}

func (s *simulator) checkSaved(want string) {
	s.t.Helper()
	data, err := os.ReadFile(s.path)
	if err != nil {
		s.t.Fatal(err)
	}
	if diff := cmp.Diff(want, string(data)); diff != "" {
		s.fatalf("saved file differs (-want +got):\n%s", diff)
	}
}

// cutAndPaste cuts some whole lines, and pastes them back.
func (s *simulator) cutAndPaste() {
	last := len(s.lines) - 1 // Selecting stops at the start of the last line.
	if last < 1 {
		return
	}
	start := s.rnd.IntN(last)
	n := 1 + s.rnd.IntN(min(5, last-start))
	s.moveTo(content.Position{Line: start})
	for range n {
		s.key("\x1b[1;2B")
	}
	s.key("\x18\x16") // Ctrl+X, Ctrl+V.
	s.cursor = content.Position{Line: start + n}
	s.done("cut and paste")
}

// finish checks the end result.
func (s *simulator) finish() {
	s.t.Helper()
	st := s.read()
	if st.mode != editor.ModeNormal {
		s.fatalf("the editor stays in mode %v", st.mode)
	}
	if diff := cmp.Diff(s.pair.after, st.text); diff != "" {
		s.fatalf("the edited text differs from the after file (-want +got):\n%s", diff)
	}
	s.key(":w\r")
	s.done("save the result")
	s.checkSaved(s.pair.after)

	s.h.Post(s.t, editor.CommandFunc(func(e *editor.Editor) { checkUndoRedo(s.t, e, s.path, s.pair.before) }))
	s.compareToFreshRender()
}

// compareToFreshRender checks that the editor kept what it derives from the text, like
// the syntax highlighting, up to date: it must show the same as an editor opening the result.
func (s *simulator) compareToFreshRender() {
	s.t.Helper()
	s.key("g0") // A new editor shows the top of the file.
	s.doneAnywhere("go to the top")
	var edited string
	s.h.Post(s.t, editor.CommandFunc(func(*editor.Editor) { edited = s.h.RenderBuffer(100) }))

	fresh, _ := openHarnessFile(s.t, s.t.TempDir(), s.pair.ext, s.pair.after)
	fresh.Run(s.t)
	synctest.Wait()
	var want string
	fresh.Post(s.t, editor.CommandFunc(func(*editor.Editor) { want = fresh.RenderBuffer(100) }))
	fresh.CloseInput()

	if diff := cmp.Diff(want, edited); diff != "" {
		s.fatalf("the edited file renders unlike the same file opened anew (-want +got):\n%s", diff)
	}
}
