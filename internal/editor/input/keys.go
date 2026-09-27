package input

import (
	"bytes"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Key is a key pressed on the keyboard, decoded from the terminal input.
// Keys are comparable: a keymap can check them with ==, like key == Ctrl('s').
type Key struct {
	Special Special  // Text for a typed character
	Rune    rune     // the typed character, or the letter pressed with Ctrl
	Cursor  Cursor   // the cursor key if Special is CursorMove
	Mod     Modifier // Shift, Ctrl, etc. for the cursor keys, Ctrl for the letters
}

// Special tells the keys that don't type text apart.
type Special byte

//go:generate go run golang.org/x/tools/cmd/stringer -type=Special

const (
	Text       Special = iota // a character in Key.Rune, which is a control one with ModCtrl
	Enter                     // Enter (\r)
	Tab                       // Tab
	Backtab                   // Shift+Tab
	Backspace                 // Backspace
	Esc                       // Esc
	CursorMove                // one of the arrows, Home, or End, see Key.Cursor
	FocusIn                   // the terminal window got the focus (with focus reporting enabled)
	FocusOut                  // the terminal window lost the focus
)

// Modifier bits.
const (
	ModShift  Modifier = 1
	ModAlt    Modifier = 2
	ModCtrl   Modifier = 4
	ModMotion Modifier = 8
)

// Rune returns the key typing the character r.
func Rune(r rune) Key { return Key{Rune: r} }

// Ctrl returns the key pressed with Ctrl: Ctrl('s') is Ctrl+S.
func Ctrl(letter rune) Key { return Key{Rune: letter, Mod: ModCtrl} }

// Of returns the special key with no modifiers.
func Of(s Special) Key { return Key{Special: s} }

// Move returns the cursor key with the modifiers.
func Move(c Cursor, mod Modifier) Key { return Key{Special: CursorMove, Cursor: c, Mod: mod} }

// Keys decodes the keys in the terminal input. It stops before a mouse event or
// a bracketed paste (see ReadMouse and ConsumeClipboardPaste), and before an incomplete
// sequence at the end of b, which is to be completed by the next read: n is the number
// of the decoded bytes. Unknown escape sequences are skipped.
//
// Esc followed by another key in the same input is decoded as two keys, so a quick
// Esc and a command in the normal mode work. Alt+key is not supported for this reason.
func Keys(b []byte) (keys []Key, n int) {
	for n < len(b) {
		c := b[n]
		switch {
		case c == escByte:
			if n+1 == len(b) || b[n+1] != '[' {
				keys = append(keys, Of(Esc))
				n++
				continue
			}
			if IsMouseInput(b[n:]) || bytes.HasPrefix(b[n:], []byte(pasteStart)) {
				return keys, n
			}
			k, size, ok := decodeCSI(b[n:])
			if size == 0 {
				return keys, n // Incomplete.
			}
			if ok {
				keys = append(keys, k)
			}
			n += size

		case c == '\r':
			keys = append(keys, Of(Enter))
			n++
		case c == '\t':
			keys = append(keys, Of(Tab))
			n++
		case c == 0x7f || c == 0x08:
			keys = append(keys, Of(Backspace))
			n++
		case c == 0x01: // Ctrl+A
			keys = append(keys, Move(CursorHome, 0))
			n++
		case c == 0x05: // Ctrl+E
			keys = append(keys, Move(CursorEnd, 0))
			n++
		case 0x01 <= c && c <= 0x1a:
			keys = append(keys, Ctrl(rune('a'+c-1)))
			n++
		case c < 0x20:
			n++ // Other control characters are not used.

		default:
			if !utf8.FullRune(b[n:]) {
				return keys, n // The rest of the character comes with the next read.
			}
			r, size := utf8.DecodeRune(b[n:])
			if r != utf8.RuneError || size > 1 {
				keys = append(keys, Rune(r))
			}
			n += size
		}
	}
	return keys, n
}

const (
	escByte    = 0x1b
	pasteStart = "\x1b[200~"
)

// decodeCSI decodes the control sequence "ESC [ params final" at the start of seq.
// The size is 0 if the sequence is incomplete, and ok is false if it's not a known key.
func decodeCSI(seq []byte) (k Key, size int, ok bool) {
	i := 2
	for i < len(seq) && 0x20 <= seq[i] && seq[i] <= 0x3f {
		i++
	}
	if i == len(seq) {
		return Key{}, 0, false
	}
	final, params := seq[i], string(seq[2:i])
	if final < 0x40 || final > 0x7e {
		return Key{}, i, false // Malformed: skip it.
	}
	size = i + 1

	switch final {
	case 'Z':
		return Of(Backtab), size, params == ""
	case 'I':
		return Of(FocusIn), size, params == ""
	case 'O':
		return Of(FocusOut), size, params == ""
	}

	var c Cursor
	switch final {
	case 'A':
		c = CursorArrowUp
	case 'B':
		c = CursorArrowDown
	case 'C':
		c = CursorArrowRight
	case 'D':
		c = CursorArrowLeft
	case 'H':
		c = CursorHome
	case 'F':
		c = CursorEnd
	default:
		return Key{}, size, false
	}
	var mod Modifier
	if params != "" {
		// "1;m": the modifiers are encoded as m-1.
		first, m, found := strings.Cut(params, ";")
		v, err := strconv.Atoi(m)
		if first != "1" || !found || err != nil || v < 1 {
			return Key{}, size, false
		}
		mod = Modifier(v - 1)
	}
	// Cmd+Left and Cmd+Right in some terminals.
	if mod.HasMotion() && c == CursorArrowLeft {
		c = CursorHome
	}
	if mod.HasMotion() && c == CursorArrowRight {
		c = CursorEnd
	}
	return Move(c, mod), size, true
}

func (k Key) String() string {
	var sb strings.Builder
	for _, m := range []struct {
		mod  Modifier
		name string
	}{{ModShift, "Shift"}, {ModAlt, "Alt"}, {ModCtrl, "Ctrl"}, {ModMotion, "Motion"}} {
		if k.Mod&m.mod != 0 {
			sb.WriteString(m.name + "+")
		}
	}
	switch k.Special {
	case Text:
		sb.WriteRune(k.Rune)
	case CursorMove:
		sb.WriteString(k.Cursor.String())
	default:
		sb.WriteString(k.Special.String())
	}
	return sb.String()
}
