package editor

import (
	"cmp"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"unicode/utf8"

	"rmazur.io/chernetka/internal/editor/inputs"
	"rmazur.io/chernetka/internal/vt/escape"
)

// quickOpenPrefix starts the command line of the file picker: ":e query".
const quickOpenPrefix = "e "

const (
	maxQuickMatches = 50    // matches kept for display
	maxProjectFiles = 20000 // files listed in the project root
)

// quickOpen is a file picker shown in the status bar while the :e command is typed.
// It lists the open buffers first and then the project files matching the query.
type quickOpen struct {
	buf    *Buffer // the buffer the picker was started from
	offset int     // the buffer scroll offset to restore when the picker is closed

	query   string
	matches []quickMatch // the best matches first, at most maxQuickMatches
	total   int          // the number of all matches
	sel     int          // the selected match
	loading bool         // project files are being listed

	tall bool // the matches take a row of the status bar on their own
}

type quickMatch struct {
	path    string // the path to open
	display string // the path relative to the project root
	open    bool   // the path is open in a buffer
}

// startQuickOpen switches the buffer to the command mode with the :e command typed.
func (e *Editor) startQuickOpen(buf *Buffer) {
	buf.mode = ModeCommand
	buf.cmdline = quickOpenPrefix
	e.syncQuickOpen(buf)
}

// syncQuickOpen starts, updates, or closes the picker to follow the command line of buf.
func (e *Editor) syncQuickOpen(buf *Buffer) {
	query, active := strings.CutPrefix(buf.cmdline, quickOpenPrefix)
	if buf.mode != ModeCommand || !active {
		e.closeQuickOpen()
		return
	}
	q := e.status.quick
	if q == nil {
		q = &quickOpen{buf: buf, offset: buf.offset}
		e.status.quick = q
		e.loadProjectFiles()
	} else if q.query == query {
		return
	}
	q.query = query
	q.refresh(e)
}

// closeQuickOpen hides the picker returning its buffer to the normal mode if necessary.
func (e *Editor) closeQuickOpen() {
	q := e.status.quick
	if q == nil {
		return
	}
	e.status.quick = nil
	if q.buf.mode == ModeCommand && strings.HasPrefix(q.buf.cmdline, quickOpenPrefix) {
		q.buf.mode = ModeNormal
		q.buf.cmdline = ""
	}
	if q.tall {
		// The taller status bar might have scrolled the content.
		q.buf.offset = q.offset
	}
	e.renderRequested = true
}

// loadProjectFiles lists the files in the project root in the background.
// The files listed before are matched in the meantime.
func (e *Editor) loadProjectFiles() {
	q, root := e.status.quick, e.root()
	q.loading = true
	go func() {
		files := listProjectFiles(root, maxProjectFiles)
		e.Send(CommandFunc(func(e *Editor) {
			e.status.projectFiles = files
			if e.status.quick == q {
				q.loading = false
				q.refresh(e)
				e.renderRequested = true
			}
		}))
	}()
}

// quickOpenInput handles the picker keys. Other input edits the command line as usual.
func (e *Editor) quickOpenInput(b []byte) (handled bool) {
	q := e.status.quick
	var (
		arrow inputs.Cursor
		mod   inputs.Modifier
	)
	switch {
	case inputs.IsCursor(b, &arrow, &mod):
		switch arrow {
		case inputs.CursorArrowUp:
			q.move(-1)
		case inputs.CursorArrowDown:
			q.move(1)
		}
	case inputs.IsTab(b), inputs.IsQuickOpenCommand(b):
		q.move(1)
	case inputs.IsBacktab(b):
		q.move(-1)
	case len(b) == 1 && b[0] == '\r':
		e.openQuickMatch()
	default:
		return false
	}
	return true
}

// openQuickMatch opens the selected match. Without matches, the query is treated as
// a path in the project root, and a new buffer is created if the file does not exist.
func (e *Editor) openQuickMatch() {
	q := e.status.quick
	var path string
	switch {
	case len(q.matches) > 0:
		path = q.matches[q.sel].path
	case filepath.IsAbs(q.query):
		path = q.query
	case q.query != "":
		path = filepath.Join(e.root(), q.query)
	}
	e.closeQuickOpen()
	if path == "" || e.findAndActivateBuffer(path) {
		return
	}
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		buf := NewScratchBuffer()
		buf.Path = path
		e.push(buf)
		return
	}
	if e.OpenPath != nil {
		e.OpenPath(path)
	} else {
		(&OpenFile{Path: path}).DoOnEditor(e)
	}
}

func (q *quickOpen) move(d int) {
	if n := len(q.matches); n > 0 {
		q.sel = (q.sel + d + n) % n
	}
}

// refresh matches the query against the open buffers and the project files.
func (q *quickOpen) refresh(e *Editor) {
	root := e.root()
	q.matches, q.total, q.sel = q.matches[:0], 0, 0

	// Open buffers go first, from the most recently used. The current one is the last:
	// switching to the previous buffer takes just Enter.
	open := make(map[string]bool)
	var current []quickMatch
	for entry := range e.buffers() {
		b := entry.b
		if b.Path == "" {
			continue
		}
		m := quickMatch{path: b.Path, display: displayPath(root, b.Path), open: true}
		open[m.display] = true
		if _, ok := fuzzyScore(q.query, m.display); !ok {
			continue
		}
		if b == q.buf {
			current = append(current, m)
		} else {
			q.matches = append(q.matches, m)
		}
	}
	q.matches = append(q.matches, current...)
	q.total = len(q.matches)

	type scoredFile struct {
		path  string
		score int
	}
	var files []scoredFile
	for _, f := range e.status.projectFiles {
		if open[f] {
			continue
		}
		if score, ok := fuzzyScore(q.query, f); ok {
			files = append(files, scoredFile{f, score})
		}
	}
	slices.SortFunc(files, func(a, b scoredFile) int {
		return cmp.Or(
			cmp.Compare(b.score, a.score),
			cmp.Compare(len(a.path), len(b.path)),
			strings.Compare(a.path, b.path),
		)
	})
	q.total += len(files)
	for _, f := range files[:min(len(files), max(maxQuickMatches-len(q.matches), 0))] {
		q.matches = append(q.matches, quickMatch{path: filepath.Join(root, f.path), display: f.path})
	}
}

// displayPath returns p relative to the root if it's in the root, or p itself otherwise.
func displayPath(root, p string) string {
	abs, err := filepath.Abs(p)
	if err != nil {
		return p
	}
	rel, err := filepath.Rel(root, abs)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return p
	}
	return rel
}

// fuzzyScore checks whether the query characters appear in the candidate in the same order,
// ignoring the case. Higher scores are given to consecutive characters, characters at the
// start of path segments or words, and matches in the file name.
func fuzzyScore(query, candidate string) (score int, ok bool) {
	if query == "" {
		return 0, true
	}
	q, c := strings.ToLower(query), strings.ToLower(candidate)
	score, ok = subsequenceScore(q, c, 0)
	if !ok {
		return 0, false
	}
	base := strings.LastIndexByte(c, filepath.Separator) + 1
	if base > 0 {
		if baseScore, ok := subsequenceScore(q, c, base); ok {
			const fileNameBonus = 10
			score = max(score, baseScore+fileNameBonus)
		}
	}
	return score, true
}

// subsequenceScore matches the query characters in c[from:]. Every occurrence of the first
// character is tried as the start, and the rest of the characters are matched greedily.
func subsequenceScore(q, c string, from int) (score int, ok bool) {
	for start := from; ; start++ {
		k := strings.IndexByte(c[start:], q[0])
		if k < 0 {
			return score, ok
		}
		start += k
		if s, matched := greedyScore(q, c, start); matched {
			score, ok = max(score, s), true
		} else {
			return score, ok // Later starts can't match either.
		}
	}
}

func greedyScore(q, c string, from int) (score int, ok bool) {
	const consecutiveBonus, boundaryBonus = 5, 3
	prev := -2
	for i := 0; i < len(q); i++ {
		k := strings.IndexByte(c[from:], q[i])
		if k < 0 {
			return 0, false
		}
		p := from + k
		score++
		if p == prev+1 {
			score += consecutiveBonus
		}
		if p == 0 || strings.IndexByte("/._- ", c[p-1]) >= 0 {
			score += boundaryBonus
		}
		prev, from = p, p+1
	}
	return score, true
}

// listProjectFiles returns up to limit file paths relative to the root.
// Git is asked for the files in a repository to skip the ignored ones.
func listProjectFiles(root string, limit int) []string {
	out, err := exec.Command("git", "-C", root, "ls-files", "-z", "--cached", "--others", "--exclude-standard").Output()
	if err == nil {
		var files []string
		for f := range strings.SplitSeq(string(out), "\x00") {
			f = filepath.FromSlash(f)
			if f == "" || (len(files) > 0 && files[len(files)-1] == f) {
				continue // Unmerged files are listed several times in a row.
			}
			files = append(files, f)
			if len(files) == limit {
				break
			}
		}
		return files
	}

	var files []string
	_ = filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			if p != root && isSkippedDir(d.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		if !d.Type().IsRegular() {
			return nil
		}
		if rel, err := filepath.Rel(root, p); err == nil {
			files = append(files, rel)
		}
		if len(files) == limit {
			return filepath.SkipAll
		}
		return nil
	})
	return files
}

func isSkippedDir(name string) bool {
	return strings.HasPrefix(name, ".") || name == "node_modules" || name == "vendor"
}

// height returns the number of the status bar rows needed to show the picker.
// Once the matches don't fit next to the command line, they stay on a separate row
// until the picker is closed, so that the content does not jump while typing.
func (q *quickOpen) height(w int) int {
	if !q.tall {
		avail := w - q.cmdWidth() - inlineGap
		from, to := q.window(avail)
		q.tall = to-from < min(minInlineMatches, len(q.matches))
	}
	if q.tall {
		return 2
	}
	return 1
}

const (
	inlineGap        = 4 // columns between the command line and the matches
	minInlineMatches = 3 // matches to show next to the command line, a separate row is used otherwise
)

func (q *quickOpen) cmdWidth() int { return 1 + utf8.RuneCountInString(q.buf.cmdline) }

// render prints the command line and the matches.
func (q *quickOpen) render(out io.Writer, w int) {
	restore := escape.ReverseVideo(out)
	defer func() { restore() }()

	cmd := ":" + q.buf.cmdline
	escape.ClearLine(out)
	if q.tall {
		from, to := q.window(w)
		q.renderMatches(out, w, from, to, &restore)
		writePadding(out, w-q.matchesWidth(w, from, to))
		printLineEnding(out)
		escape.ClearLine(out)
	}
	_, _ = io.WriteString(out, cmd)

	avail := w - q.cmdWidth() - inlineGap
	if from, to := q.window(avail); !q.tall && to > from {
		writePadding(out, w-q.cmdWidth()-q.matchesWidth(avail, from, to))
		q.renderMatches(out, avail, from, to, &restore)
		return
	}
	info := q.info()
	writePadding(out, w-q.cmdWidth()-utf8.RuneCountInString(info)-1)
	_, _ = io.WriteString(out, info+" ")
}

// info describes the matches for the command line.
func (q *quickOpen) info() string {
	switch {
	case len(q.matches) > 0:
		return fmt.Sprintf("%d/%d", q.sel+1, q.total)
	case q.loading:
		return "loading…"
	case q.query != "":
		return "new file"
	default:
		return "no files"
	}
}

// renderMatches prints the matches in the range selected by window for w columns.
// The selected match is printed with the regular colors, restore is updated to re-enable
// the reverse video.
func (q *quickOpen) renderMatches(out io.Writer, w, from, to int, restore *func()) {
	for i := from; i < to; i++ {
		item := q.item(i, w)
		if i == q.sel {
			(*restore)()
			_, _ = io.WriteString(out, item)
			*restore = escape.ReverseVideo(out)
		} else {
			_, _ = io.WriteString(out, item)
		}
	}
	_, _ = io.WriteString(out, q.moreLabel(to))
}

// matchesWidth returns the number of columns taken by renderMatches.
func (q *quickOpen) matchesWidth(w, from, to int) int {
	sum := utf8.RuneCountInString(q.moreLabel(to))
	for i := from; i < to; i++ {
		sum += utf8.RuneCountInString(q.item(i, w))
	}
	return sum
}

// item returns the text of the match to show in w columns.
func (q *quickOpen) item(i, w int) string {
	return " " + shortenPath(q.matches[i].display, w-2) + " "
}

func (q *quickOpen) moreLabel(shown int) string {
	if shown >= q.total {
		return ""
	}
	return fmt.Sprintf(" +%d ", q.total-shown)
}

// window selects the range of matches that fit into w columns and include the selected one.
func (q *quickOpen) window(w int) (from, to int) {
	if len(q.matches) == 0 || w <= 2 {
		return 0, 0
	}
	fits := func(from, to int) bool { return q.matchesWidth(w, from, to) <= w }
	from, to = 0, q.sel+1
	for from < q.sel && !fits(from, to) {
		from++
	}
	if !fits(from, to) {
		return from, from // Even the selected match does not fit.
	}
	for to < len(q.matches) && fits(from, to+1) {
		to++
	}
	return from, to
}

// shortenPath cuts the beginning of the path to fit it into w columns.
func shortenPath(p string, w int) string {
	n := utf8.RuneCountInString(p)
	if n <= w || w < 2 {
		return p
	}
	runes := []rune(p)
	return "…" + string(runes[n-w+1:])
}

func writePadding(out io.Writer, n int) {
	if n > 0 {
		_, _ = io.WriteString(out, strings.Repeat(" ", n))
	}
}
