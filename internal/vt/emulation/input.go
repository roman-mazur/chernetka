package emulation

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
)

// Ban lists the input the monkey must not send.
type Ban struct {
	Runes     string // the characters never typed, the pastes may still have them
	Ctrl      string // the letters never pressed with Ctrl
	CtrlMouse bool   // no mouse events with Ctrl, like Ctrl+click
}

// contains reports whether the snippet has a banned character.
func (b Ban) contains(snippet string) bool {
	return b.Runes != "" && strings.ContainsAny(snippet, b.Runes)
}

// inputs generates random terminal input: typed text, special and Ctrl keys, cursor keys
// with modifiers, pastes, mouse events in the SGR encoding, malformed sequences,
// and the snippets of the tested app.
type inputs struct {
	rnd         *rand.Rand
	cols, rows  int
	snippets    []string
	noCtrlMouse bool

	runes []rune // the characters to type
	keys  []string
}

// The special keys to press, some several times to press them more often.
var specialKeys = []string{
	"\x1b", "\x1b", "\x1b", "\x7f", "\x7f", "\x7f", "\t", "\x1b[Z", "\r", "\r",
	"\x1b[H", "\x1b[F", "\x1b[2~", "\x1b[3~", "\x1b[5~", "\x1b[6~", // Home, End, Insert, Delete, Page Up and Down
	"\x1b[I", "\x1b[O", // Focus in and out.
	"\x1b[", "\x1b[1;5", "\x1b[99", // Malformed and incomplete sequences.
}

var (
	cursorKeys      = []byte("ABCD")
	cursorModifiers = []string{"", "", "1;2", "1;3", "1;5", "1;6", "1;7", "1;9", "1;10"}
)

func newInputs(rnd *rand.Rand, ban Ban, cols, rows int, snippets []string) *inputs {
	in := &inputs{rnd: rnd, cols: cols, rows: rows, snippets: snippets, noCtrlMouse: ban.CtrlMouse}
	for r := rune(' '); r <= '~'; r++ {
		in.runes = append(in.runes, r)
	}
	in.runes = append(in.runes, []rune("éїж世界😀")...)
	in.runes = filterRunes(in.runes, ban.Runes)

	in.keys = slices.Clone(specialKeys)
	for c := byte('a'); c <= 'z'; c++ {
		if !strings.ContainsRune(ban.Ctrl, rune(c)) {
			in.keys = append(in.keys, string(c-'a'+1))
		}
	}
	return in
}

func filterRunes(runes []rune, banned string) []rune {
	var res []rune
	for _, r := range runes {
		if !strings.ContainsRune(banned, r) {
			res = append(res, r)
		}
	}
	return res
}

// next returns the next random input.
func (in *inputs) next() string {
	r := in.rnd
	switch n := r.IntN(100); {
	case n < 35:
		return in.text(1 + r.IntN(8))
	case n < 50 && len(in.snippets) > 0:
		return in.snippets[r.IntN(len(in.snippets))]
	case n < 65:
		return in.keys[r.IntN(len(in.keys))]
	case n < 80:
		mod := cursorModifiers[r.IntN(len(cursorModifiers))]
		return "\x1b[" + mod + string(cursorKeys[r.IntN(len(cursorKeys))])
	case n < 88:
		return "\x1b[200~" + in.pasteText(r.IntN(200)) + "\x1b[201~"
	default:
		return in.mouse()
	}
}

// text returns n random characters to type.
func (in *inputs) text(n int) string {
	var sb strings.Builder
	for range n {
		sb.WriteRune(in.runes[in.rnd.IntN(len(in.runes))])
	}
	return sb.String()
}

var pastedUnicode = []string{"é", "世", "😀"}

// pasteText returns random text of n characters, any printable ones, tabs, and new lines.
func (in *inputs) pasteText(n int) string {
	var sb strings.Builder
	for range n {
		switch in.rnd.IntN(20) {
		case 0:
			sb.WriteByte('\n')
		case 1:
			sb.WriteByte('\t')
		case 2:
			sb.WriteString(pastedUnicode[in.rnd.IntN(len(pastedUnicode))])
		default:
			sb.WriteByte(byte(' ' + in.rnd.IntN('~'-' '+1)))
		}
	}
	return sb.String()
}

// Mouse buttons and modifiers of the SGR encoding.
const (
	mouseLeft   = 0
	mouseRight  = 2
	mouseNone   = 3
	mouseShift  = 4
	mouseAlt    = 8
	mouseCtrl   = 16
	mouseMotion = 32
	mouseWheel  = 64
)

// mouse returns random mouse events: presses, releases, drags, hovering, scrolling,
// and quick series of clicks, some slightly out of the window.
func (in *inputs) mouse() string {
	r := in.rnd
	x, y := 1+r.IntN(in.cols+2), 1+r.IntN(in.rows+2)
	mods := []int{0, 0, 0, mouseShift, mouseAlt, mouseCtrl}
	if in.noCtrlMouse {
		mods = mods[:len(mods)-1]
	}
	mod := mods[r.IntN(len(mods))]
	event := func(b int, pressed bool) string {
		end := 'M'
		if !pressed {
			end = 'm'
		}
		return fmt.Sprintf("\x1b[<%d;%d;%d%c", b|mod, x, y, end)
	}
	switch r.IntN(6) {
	case 0: // Clicks.
		var sb strings.Builder
		for range 1 + r.IntN(3) {
			sb.WriteString(event(mouseLeft, true))
			sb.WriteString(event(mouseLeft, false))
		}
		return sb.String()
	case 1: // Drag.
		var sb strings.Builder
		sb.WriteString(event(mouseLeft, true))
		for range 1 + r.IntN(5) {
			x, y = max(1, x+r.IntN(11)-5), max(1, y+r.IntN(5)-2)
			sb.WriteString(event(mouseLeft|mouseMotion, true))
		}
		sb.WriteString(event(mouseLeft, false))
		return sb.String()
	case 2:
		return event(mouseWheel+r.IntN(4), true)
	case 3:
		return event(mouseNone|mouseMotion, true)
	case 4:
		return event([]int{mouseLeft, mouseRight}[r.IntN(2)], r.IntN(2) == 0)
	default:
		return event(mouseLeft|mouseMotion, true) // Motion without a press.
	}
}
