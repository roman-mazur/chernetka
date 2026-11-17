package editor

import (
	"fmt"
	"io"
	"strings"

	"rmazur.io/chernetka/internal/content/symbol"
	"rmazur.io/chernetka/internal/vt/escape"
)

// pickList shows the items to pick from in the status bar next to the command line,
// or on a row of their own above it once they don't fit. It's embedded by the prompts
// listing their matches, like the file picker.
type pickList struct {
	items []string // what is shown for the matches, the best first
	total int      // the number of all matches, more than the items if some are left out
	sel   int      // the selected item

	tall bool // the items take a row of the status bar on their own
}

// reset replaces the items with the given ones, selecting the first.
func (p *pickList) reset(items []string, total int) {
	p.items, p.total, p.sel = items, total, 0
}

// move selects the item d items away from the selected one, wrapping around.
func (p *pickList) move(d int) {
	if n := len(p.items); n > 0 {
		p.sel = (p.sel + d + n) % n
	}
}

// height returns the number of the status bar rows needed to show the items next to
// a command line cmdWidth columns wide. Once the items don't fit next to the command
// line, they stay on a separate row until the prompt is closed, so that the content
// does not jump while typing.
func (p *pickList) height(w, cmdWidth int) int {
	if !p.tall {
		avail := w - cmdWidth - inlineGap
		from, to := p.window(avail)
		p.tall = to-from < min(minInlineMatches, len(p.items))
	}
	if p.tall {
		return 2
	}
	return 1
}

const (
	inlineGap        = 4 // columns between the command line and the matches
	minInlineMatches = 3 // matches to show next to the command line, a separate row is used otherwise
)

// render prints the command line and the items. The info is shown instead of the
// items when they don't fit next to the command line.
func (p *pickList) render(out io.Writer, w int, cmd, info string) {
	restore := escape.ReverseVideo(out)
	// Not deferred directly: renderMatches may replace restore.
	defer func() { restore.Undo() }()

	cmdWidth := symbol.DisplayWidth(cmd, symbol.NoTabExpansion)
	escape.ClearLine(out)
	if p.tall {
		from, to := p.window(w)
		p.renderMatches(out, w, from, to, &restore)
		writePadding(out, w-p.matchesWidth(w, from, to))
		printLineEnding(out)
		escape.ClearLine(out)
	}
	_, _ = io.WriteString(out, cmd)

	avail := w - cmdWidth - inlineGap
	if from, to := p.window(avail); !p.tall && to > from {
		writePadding(out, w-cmdWidth-p.matchesWidth(avail, from, to))
		p.renderMatches(out, avail, from, to, &restore)
		return
	}
	writePadding(out, w-cmdWidth-symbol.DisplayWidth(info, symbol.NoTabExpansion)-1)
	_, _ = io.WriteString(out, info+" ")
}

// renderMatches prints the items in the range selected by window for w columns.
// The selected item is printed with the regular colors, restore is updated to re-enable
// the reverse video.
func (p *pickList) renderMatches(out io.Writer, w, from, to int, restore *escape.Restore) {
	for i := from; i < to; i++ {
		item := p.item(i, w)
		if i == p.sel {
			restore.Undo()
			_, _ = io.WriteString(out, item)
			*restore = escape.ReverseVideo(out)
		} else {
			_, _ = io.WriteString(out, item)
		}
	}
	_, _ = io.WriteString(out, p.moreLabel(to))
}

// matchesWidth returns the number of columns taken by renderMatches.
func (p *pickList) matchesWidth(w, from, to int) int {
	sum := symbol.DisplayWidth(p.moreLabel(to), symbol.NoTabExpansion)
	for i := from; i < to; i++ {
		sum += symbol.DisplayWidth(p.item(i, w), symbol.NoTabExpansion)
	}
	return sum
}

// item returns the text of the item to show in w columns.
func (p *pickList) item(i, w int) string {
	return " " + shortenPath(p.items[i], w-2) + " "
}

func (p *pickList) moreLabel(shown int) string {
	if shown >= p.total {
		return ""
	}
	return fmt.Sprintf(" +%d ", p.total-shown)
}

// window selects the range of items that fit into w columns and include the selected one.
func (p *pickList) window(w int) (from, to int) {
	if len(p.items) == 0 || w <= 2 {
		return 0, 0
	}
	fits := func(from, to int) bool { return p.matchesWidth(w, from, to) <= w }
	from, to = 0, p.sel+1
	for from < p.sel && !fits(from, to) {
		from++
	}
	if !fits(from, to) {
		return from, from // Even the selected item does not fit.
	}
	for to < len(p.items) && fits(from, to+1) {
		to++
	}
	return from, to
}

// shortenPath cuts the beginning of the path to fit it into w columns.
func shortenPath(p string, w int) string {
	rest := symbol.DisplayWidth(p, symbol.NoTabExpansion)
	if rest <= w || w < 2 {
		return p
	}
	start := 0
	for cluster, clusterWidth := range symbol.Clusters(p, symbol.NoTabExpansion) {
		if rest <= w-1 {
			break
		}
		start += len(cluster)
		rest -= clusterWidth
	}
	return "…" + p[start:]
}

func writePadding(out io.Writer, n int) {
	if n > 0 {
		_, _ = io.WriteString(out, strings.Repeat(" ", n))
	}
}
