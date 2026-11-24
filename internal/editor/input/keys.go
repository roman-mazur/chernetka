package input

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Key is a key pressed on the keyboard, decoded from the terminal input.
// Keys are comparable: a keymap can check them with ==, like key == Ctrl('s').
type Key struct {
	Special Special  // Text for a typed character
	Rune    rune     // the typed character, the letter pressed with Ctrl, or the number of the function key
	Cursor  Cursor   // the cursor key if Special is CursorMove
	Mod     Modifier // Shift, Ctrl, etc.; only Ctrl for the letters without the kitty keyboard protocol
}

// Special tells the keys that don't type text apart.
type Special byte

//go:generate go run golang.org/x/tools/cmd/stringer -type=Special

const (
	Text       Special = iota // a character in Key.Rune, the letter pressed with ModCtrl
	Enter                     // Enter (\r)
	Tab                       // Tab
	Backtab                   // Shift+Tab
	Backspace                 // Backspace
	Esc                       // Esc
	CursorMove                // one of the arrows, Home, or End, see Key.Cursor
	FocusIn                   // the terminal window got the focus (with focus reporting enabled)
	FocusOut                  // the terminal window lost the focus
	Function                  // one of the function keys F1 to F12, the number is in Key.Rune
	Delete                    // Delete (forward delete), sent as "CSI 3 ~"
	Insert                    // Insert, sent as "CSI 2 ~"
)

// Modifier bits.
const (
	ModShift  Modifier = 1
	ModAlt    Modifier = 2
	ModCtrl   Modifier = 4
	ModMotion Modifier = 8
	ModWheel  Modifier = 16 // mouse scrolling, see ScrollDirection
)

// Rune returns the key typing the character r.
func Rune(r rune) Key { return Key{Rune: r} }

// Ctrl returns the key pressed with Ctrl: Ctrl('s') is Ctrl+S.
func Ctrl(letter rune) Key { return Key{Rune: letter, Mod: ModCtrl} }

// CtrlShift returns the key pressed with Ctrl and Shift: CtrlShift('z') is Ctrl+Shift+Z.
// The terminal tells it from the key pressed with Ctrl only with the kitty keyboard protocol
// (see escape.EnableKittyKeyboard), and sends Ctrl+Z for it otherwise.
func CtrlShift(letter rune) Key { return Key{Rune: letter, Mod: ModCtrl | ModShift} }

// Of returns the special key with no modifiers.
func Of(s Special) Key { return Key{Special: s} }

// Fn returns the function key with the number pressed with the modifiers: Fn(6, ModShift) is Shift+F6.
func Fn(number int, mod Modifier) Key { return Key{Special: Function, Rune: rune(number), Mod: mod} }

// Move returns the cursor key with the modifiers.
func Move(c Cursor, mod Modifier) Key { return Key{Special: CursorMove, Cursor: c, Mod: mod} }

// Keys decodes the keys in the terminal input. It stops before a mouse event or
// a bracketed paste (see ReadMouse and ConsumeClipboardPaste), and before an incomplete
// sequence at the end of b, which is to be completed by the next read: n is the number
// of the decoded bytes. Unknown escape sequences are skipped.
//
// Esc followed by another key in the same input is decoded as two keys, so a quick
// Esc and a command in the review mode work. Alt+key is not supported for this reason.
// With the kitty keyboard protocol, Esc and the keys pressed with the modifiers come as
// "CSI code ; modifiers u" instead, and are not ambiguous.
func Keys(b []byte) (keys []Key, n int) {
	for n < len(b) {
		c := b[n]
		switch {
		case c == escByte:
			if n+1 < len(b) && b[n+1] == 'O' {
				// SS3: F1 to F4 as "ESC O P" to "ESC O S" in the terminals without the kitty keyboard protocol.
				if n+2 == len(b) {
					return keys, n // Incomplete.
				}
				if fn := ss3Function(b[n+2]); fn != 0 {
					keys = append(keys, Fn(fn, 0))
					n += 3
					continue
				}
			}
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
		case c == 0x1f: // Ctrl+/ (also Ctrl+_)
			keys = append(keys, Ctrl('/'))
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
	case 'u':
		k, ok := decodeKittyKey(params)
		return k, size, ok
	case 'Z':
		return Of(Backtab), size, params == ""
	case 'I':
		return Of(FocusIn), size, params == ""
	case 'O':
		return Of(FocusOut), size, params == ""
	case '~':
		k, ok := decodeTildeKey(params)
		return k, size, ok
	case 'P', 'Q', 'R', 'S':
		// F1 to F4 as "CSI P" or "CSI 1 ; modifiers P". "CSI row ; col R" is a cursor
		// position report, which is not expected: only "1" is taken as the first parameter.
		first, m, found := strings.Cut(params, ";")
		mod, ok := csiModifier(m, found)
		if !ok || (first != "" && first != "1") {
			return Key{}, size, false
		}
		return Fn(int(final-'P')+1, mod), size, true
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
		var ok bool
		if mod, ok = csiModifier(m, found); first != "1" || !ok {
			return Key{}, size, false
		}
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

// csiModifier decodes the modifiers parameter of a control sequence, encoded as the
// modifier bits plus one. It's zero if the parameter is not given.
func csiModifier(m string, given bool) (Modifier, bool) {
	if !given {
		return 0, m == ""
	}
	v, err := strconv.Atoi(m)
	if err != nil || v < 1 {
		return 0, false
	}
	return Modifier(v - 1), true
}

// tildeKeys are the keys sent as "CSI number ~", also with ";modifiers" before the tilde.
var tildeKeys = map[string]Key{
	"2": Of(Insert), "3": Of(Delete),
	"11": Fn(1, 0), "12": Fn(2, 0), "13": Fn(3, 0), "14": Fn(4, 0), "15": Fn(5, 0),
	"17": Fn(6, 0), "18": Fn(7, 0), "19": Fn(8, 0), "20": Fn(9, 0), "21": Fn(10, 0),
	"23": Fn(11, 0), "24": Fn(12, 0),
}

// decodeTildeKey decodes the parameters of a "CSI number ; modifiers ~" sequence.
func decodeTildeKey(params string) (Key, bool) {
	number, m, found := strings.Cut(params, ";")
	k, known := tildeKeys[number]
	mod, ok := csiModifier(m, found)
	if !known || !ok {
		return Key{}, false
	}
	k.Mod = mod
	return k, true
}

// ss3Function returns the function key sent as "ESC O final", or 0.
func ss3Function(final byte) int {
	if 'P' <= final && final <= 'S' {
		return int(final-'P') + 1
	}
	return 0
}

// decodeKittyKey decodes the params "code[:alternates][;modifiers[:event]]" of a key sent
// with the kitty keyboard protocol: https://sw.kovidgoyal.net/kitty/keyboard-protocol/.
// The code is the key without Shift: Ctrl+Shift+Z is "122;6".
func decodeKittyKey(params string) (Key, bool) {
	codes, mods, _ := strings.Cut(params, ";")
	code, _, _ := strings.Cut(codes, ":") // The shifted key and the key on the base layout follow.
	c, err := strconv.Atoi(code)
	if err != nil || c <= 0 || c > utf8.MaxRune {
		return Key{}, false
	}
	var mod Modifier
	if mods != "" {
		m, _, _ := strings.Cut(mods, ":") // The event type follows.
		v, err := strconv.Atoi(m)
		if err != nil || v < 1 {
			return Key{}, false
		}
		// Encoded as 1 + the bits, which are the same as the Modifier bits up to Super (ModMotion).
		// Hyper, Meta, Caps Lock, and Num Lock are not used.
		mod = Modifier(v-1) & (ModShift | ModAlt | ModCtrl | ModMotion)
	}

	switch r := rune(c); {
	case r == escByte:
		return Key{Special: Esc, Mod: mod}, true
	case r == '\r':
		return Key{Special: Enter, Mod: mod}, true
	case r == '\t' && mod == ModShift:
		return Of(Backtab), true
	case r == '\t':
		return Key{Special: Tab, Mod: mod}, true
	case r == 0x7f:
		return Key{Special: Backspace, Mod: mod}, true
	case mod == ModCtrl && controlByte(r) != 0:
		// The same key as the control character sent without the protocol: Ctrl+A is Home.
		keys, _ := Keys([]byte{controlByte(r)})
		return keys[0], true
	case unicode.IsPrint(r): // Not one of the functional keys in the Private Use Area.
		return Key{Rune: r, Mod: mod}, true
	}
	return Key{}, false
}

// controlByte returns the control character sent for the key pressed with Ctrl without
// the kitty keyboard protocol, or 0 if there is none.
func controlByte(r rune) byte {
	switch {
	case 'a' <= r && r <= 'z':
		return byte(r-'a') + 1
	case r == '[':
		return escByte
	case r == '/':
		return 0x1f
	}
	return 0
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
	case Function:
		sb.WriteString("F" + strconv.Itoa(int(k.Rune)))
	default:
		sb.WriteString(k.Special.String())
	}
	return sb.String()
}
