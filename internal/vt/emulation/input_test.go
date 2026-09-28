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
	in := newInputs(rand.New(rand.NewPCG(1, 2)), ban, 80, 40, []string{"\x1bh", ":w\r"})
	var mouseEvents, pastes int
	for range 100000 {
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
