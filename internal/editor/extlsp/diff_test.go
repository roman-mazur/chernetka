package extlsp

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"go.lsp.dev/protocol"
)

func TestDiff(t *testing.T) {
	pos := func(line, char uint32) protocol.Position { return protocol.Position{Line: line, Character: char} }
	cases := []struct {
		name          string
		before, after string
		start, end    protocol.Position
		text          string
	}{
		{"insert a char", "fmt.Pr\n", "fmt.Pri\n", pos(0, 6), pos(0, 6), "i"},
		{"delete a char", "a\nfmt.Pri\n", "a\nfmt.Pr\n", pos(1, 6), pos(1, 7), ""},
		{"new line", "a\nb", "a\n\nb", pos(1, 0), pos(1, 0), "\n"},
		{"join lines", "a\nb\nc", "a\nbc", pos(1, 1), pos(2, 0), ""},
		{"no change", "abc", "abc", pos(0, 3), pos(0, 3), ""},
		{"from empty", "", "package main\n", pos(0, 0), pos(0, 0), "package main\n"},
		{"to empty", "x\ny", "", pos(0, 0), pos(1, 1), ""},
		{"utf16 columns", "日本 x", "日本 xy", pos(0, 4), pos(0, 4), "y"},
		{"astral rune", "\U0001F600a", "\U0001F600ba", pos(0, 2), pos(0, 2), "b"},
		// é (c3 a9) -> è (c3 a8): the common first byte must not split the rune.
		{"same leading byte", "é", "è", pos(0, 0), pos(0, 1), "è"},
		{"repeated text", "aaaa", "aaaaa", pos(0, 4), pos(0, 4), "a"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := diff(tc.before, tc.after)
			if got.Range == nil || got.Range.Start != tc.start || got.Range.End != tc.end || got.Text != tc.text {
				t.Errorf("diff(%q, %q) = %+v %q, want %v-%v %q", tc.before, tc.after, got.Range, got.Text, tc.start, tc.end, tc.text)
			}
			if applied := applyChange(tc.before, got); applied != tc.after {
				t.Errorf("applying the diff gives %q, want %q", applied, tc.after)
			}
		})
	}
}

// TestDiff_Random applies diffs of random edits to check they always
// reproduce the new text.
func TestDiff_Random(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	alphabet := []rune("ab\n\té日\U0001F600")
	randText := func(n int) string {
		var sb strings.Builder
		for range n {
			sb.WriteRune(alphabet[r.Intn(len(alphabet))])
		}
		return sb.String()
	}
	text := randText(20)
	for i := range 2000 {
		runes := []rune(text)
		from := r.Intn(len(runes) + 1)
		to := from + r.Intn(len(runes)-from+1)
		next := string(runes[:from]) + randText(r.Intn(4)) + string(runes[to:])

		if got := applyChange(text, diff(text, next)); got != next {
			t.Fatalf("step %d: diff of %q -> %q applies as %q", i, text, next, got)
		}
		text = next
	}
}

// TestCommonPrefixSuffix checks the chunked comparisons against a byte loop,
// with differences around the chunk boundaries of long texts.
func TestCommonPrefixSuffix(t *testing.T) {
	bytePrefix := func(a, b string) int {
		i := 0
		for i < len(a) && i < len(b) && a[i] == b[i] {
			i++
		}
		return i
	}
	byteSuffix := func(a, b string) int {
		i := 0
		for i < len(a) && i < len(b) && a[len(a)-1-i] == b[len(b)-1-i] {
			i++
		}
		return i
	}
	for _, base := range []string{
		strings.Repeat("x", 300), // prefix and suffix may overlap
		strings.Repeat("func f() {\n\t日本\n}\n", 20),
	} {
		for at := 0; at <= len(base); at++ {
			for _, other := range []string{
				base[:at] + "y" + base[at:],             // insertion
				base[:at] + base[min(at+1, len(base)):], // deletion
				base[:at],                               // truncation
			} {
				if got, want := commonPrefix(base, other), bytePrefix(base, other); got != want {
					t.Fatalf("commonPrefix at %d = %d, want %d", at, got, want)
				}
				if got, want := commonSuffix(base, other), byteSuffix(base, other); got != want {
					t.Fatalf("commonSuffix at %d = %d, want %d", at, got, want)
				}
				if got := applyChange(base, diff(base, other)); got != other {
					t.Fatalf("diff at %d applies as %q, want %q", at, got, other)
				}
			}
		}
	}
}

// BenchmarkDiff measures a keystroke in documents of different sizes: the
// cost is in finding the common prefix and suffix.
func BenchmarkDiff(b *testing.B) {
	for _, lines := range []int{800, 5000, 20000} {
		var sb strings.Builder
		for i := range lines {
			fmt.Fprintf(&sb, "\tresult = append(result, compute(value%d, other)) // line %d\n", i, i)
		}
		before := sb.String()
		for _, where := range []struct {
			name string
			at   int
		}{{"start", 10}, {"middle", len(before) / 2}, {"end", len(before) - 10}} {
			after := before[:where.at] + "x" + before[where.at:]
			b.Run(fmt.Sprintf("%dlines/%s", lines, where.name), func(b *testing.B) {
				if lines == 5000 && where.name == "middle" {
					// A noisy one, see internal/cmd/benchmarks.
					b.Log("tolerance=20%")
				}
				for b.Loop() {
					_ = diff(before, after)
				}
			})
		}
	}
}
