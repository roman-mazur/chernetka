package emulation

import (
	"math/rand/v2"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var (
	pasteRe = regexp.MustCompile(`(?s)\x1b\[200~.*?\x1b\[201~`)
	mouseRe = regexp.MustCompile(`\x1b\[<(\d+);\d+;\d+[Mm]`)
	csiRe   = regexp.MustCompile(`\x1b\[[0-9;]*[A-Za-z~]?`)
)

func TestInputs_Ban(t *testing.T) {
	ban := Ban{Runes: "qy", Ctrl: "cxv", CtrlMouse: true}
	in := newInputs(NewChoices(randomBytes(1, 2<<20)), ban, 80, 40, []string{"\x1bh", ":w\r"})
	var mouseEvents, pastes int
	for !in.rnd.Exhausted() {
		input := in.next()

		pastes += len(pasteRe.FindAllString(input, -1))
		keys := pasteRe.ReplaceAllString(input, "") // The pastes may have any characters.
		for _, m := range mouseRe.FindAllStringSubmatch(keys, -1) {
			mouseEvents++
			if b, _ := strconv.Atoi(m[1]); b&mouseCtrl != 0 {
				t.Fatalf("mouse event with Ctrl: %q", input)
			}
		}
		keys = mouseRe.ReplaceAllString(keys, "")
		keys = csiRe.ReplaceAllString(keys, "") // The sequences have letters too.
		if strings.ContainsAny(keys, ban.Runes) {
			t.Fatalf("banned character in %q", input)
		}
		for _, c := range ban.Ctrl {
			if strings.ContainsRune(keys, c-'a'+1) {
				t.Fatalf("banned Ctrl+%c in %q", c, input)
			}
		}
	}
	if mouseEvents == 0 || pastes == 0 {
		t.Errorf("%d mouse events and %d pastes are generated", mouseEvents, pastes)
	}
}

// randomBytes returns n pseudo-random bytes, the same ones for the same seed.
func randomBytes(seed uint64, n int) []byte {
	data := make([]byte, n)
	_, _ = rand.NewChaCha8([32]byte{byte(seed)}).Read(data)
	return data
}

func TestChoices(t *testing.T) {
	c := NewChoices([]byte{7, 200, 1, 2, 3})
	for i, tc := range []struct{ n, want int }{
		{n: 5, want: 2},      // 7 % 5
		{n: 1, want: 0},      // No byte is consumed.
		{n: 256, want: 200},  // One byte is enough.
		{n: 1000, want: 258}, // Two bytes: (1<<8 | 2) % 1000.
		{n: 10, want: 3},     // The last byte.
		{n: 10, want: 0},     // Exhausted.
	} {
		if got := c.IntN(tc.n); got != tc.want {
			t.Errorf("decision %d: IntN(%d) = %d, want %d", i, tc.n, got, tc.want)
		}
	}
	if !c.Exhausted() {
		t.Error("choices are not exhausted")
	}
}
