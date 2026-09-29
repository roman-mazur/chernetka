package input

import (
	"slices"
	"testing"
)

func TestKeys(t *testing.T) {
	up, down := Move(CursorArrowUp, 0), Move(CursorArrowDown, 0)
	for _, tc := range []struct {
		name  string
		input string
		keys  []Key
		n     int // -1 for len(input)
	}{
		{name: "empty", input: ""},
		{name: "text", input: "ab", keys: []Key{Rune('a'), Rune('b')}, n: -1},
		{name: "utf8", input: "кіт", keys: []Key{Rune('к'), Rune('і'), Rune('т')}, n: -1},
		{name: "incomplete utf8", input: "a\xd0", keys: []Key{Rune('a')}, n: 1},
		{name: "invalid utf8", input: "\xffa", keys: []Key{Rune('a')}, n: -1},
		{name: "two arrows", input: "\x1b[A\x1b[B", keys: []Key{up, down}, n: -1},
		{name: "arrow and text", input: "\x1b[Ax", keys: []Key{up, Rune('x')}, n: -1},
		{name: "shift+right", input: "\x1b[1;2C", keys: []Key{Move(CursorArrowRight, ModShift)}, n: -1},
		{name: "cmd+left", input: "\x1b[1;9D", keys: []Key{Move(CursorHome, ModMotion)}, n: -1},
		{name: "home end", input: "\x1b[H\x1b[F\x01\x05",
			keys: []Key{Move(CursorHome, 0), Move(CursorEnd, 0), Move(CursorHome, 0), Move(CursorEnd, 0)}, n: -1},
		{name: "specials", input: "\r\t\x1b[Z\x7f\x08",
			keys: []Key{Of(Enter), Of(Tab), Of(Backtab), Of(Backspace), Of(Backspace)}, n: -1},
		{name: "focus", input: "\x1b[I\x1b[O", keys: []Key{Of(FocusIn), Of(FocusOut)}, n: -1},
		{name: "ctrl", input: "\x13\x0f\x06\n", keys: []Key{Ctrl('s'), Ctrl('o'), Ctrl('f'), Ctrl('j')}, n: -1},
		{name: "ctrl slash", input: "\x1f", keys: []Key{Ctrl('/')}, n: -1},
		{name: "esc", input: "\x1b", keys: []Key{Of(Esc)}, n: -1},
		{name: "esc and command", input: "\x1bi", keys: []Key{Of(Esc), Rune('i')}, n: -1},
		{name: "unknown sequence", input: "\x1b[3~a", keys: []Key{Rune('a')}, n: -1},
		{name: "incomplete sequence", input: "a\x1b[1;", keys: []Key{Rune('a')}, n: 1},
		{name: "before mouse", input: "a\x1b[<0;1;2M", keys: []Key{Rune('a')}, n: 1},
		{name: "before paste", input: "a\x1b[200~b", keys: []Key{Rune('a')}, n: 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			keys, n := Keys([]byte(tc.input))
			if tc.n < 0 {
				tc.n = len(tc.input)
			}
			if !slices.Equal(keys, tc.keys) || n != tc.n {
				t.Errorf("Keys(%q) = %v, %d; want %v, %d", tc.input, keys, n, tc.keys, tc.n)
			}
		})
	}
}

func TestKey_String(t *testing.T) {
	for k, want := range map[Key]string{
		Rune('a'):                         "a",
		Ctrl('s'):                         "Ctrl+s",
		Of(Enter):                         "Enter",
		Move(CursorArrowLeft, ModShift):   "Shift+ArrowLeft",
		Move(CursorEnd, ModMotion|ModAlt): "Alt+Motion+End",
	} {
		if got := k.String(); got != want {
			t.Errorf("%#v.String() = %q, want %q", k, got, want)
		}
	}
}
