package symbol

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	testMermaid = "\U0001F9DC\u200D\u2640\uFE0F"
	testFlag    = "\U0001F1FA\U0001F1E6"
	testAccent  = "\u0430\u0301" // A Cyrillic letter with a combining stress mark.
)

func TestClusters(t *testing.T) {
	type item struct {
		cluster string
		width   int
	}
	for _, tc := range []struct {
		name string
		text string
		want []item
	}{
		{name: "empty", text: ""},
		{name: "ascii", text: "ab", want: []item{{"a", 1}, {"b", 1}}},
		{name: "tab", text: "\ta", want: []item{{"\t", 4}, {"a", 1}}},
		{name: "tab before non-ascii", text: "\tі", want: []item{{"\t", 4}, {"і", 1}}},
		{name: "cyrillic", text: "кіт", want: []item{{"к", 1}, {"і", 1}, {"т", 1}}},
		{name: "wide", text: "🌊a", want: []item{{"🌊", 2}, {"a", 1}}},
		{name: "cjk", text: "猫", want: []item{{"猫", 2}}},
		{name: "joined emoji", text: "a" + testMermaid + "b", want: []item{{"a", 1}, {testMermaid, 2}, {"b", 1}}},
		{name: "flags", text: testFlag + testFlag, want: []item{{testFlag, 2}, {testFlag, 2}}},
		{name: "combining", text: testAccent + "б", want: []item{{testAccent, 1}, {"б", 1}}},
		{name: "invalid utf-8", text: "\xffa", want: []item{{"\xff", 1}, {"a", 1}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got []item
			for cluster, width := range Clusters(tc.text, 4) {
				got = append(got, item{cluster, width})
			}
			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(item{})); diff != "" {
				t.Errorf("Clusters (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDisplayWidth(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    string
		want int
	}{
		{name: "empty", s: "", want: 0},
		{name: "ascii", s: "abc", want: 3},
		{name: "ascii with a tab", s: "a\tb", want: 6},
		{name: "wide with a tab", s: "\t🌊", want: 6},
		{name: "joined emoji", s: testMermaid, want: 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := DisplayWidth(tc.s, 4); got != tc.want {
				t.Errorf("DisplayWidth(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}

func TestClusterBoundaries(t *testing.T) {
	line := "a" + testMermaid + testFlag + "b"
	flagAt := 1 + len(testMermaid)
	bAt := flagAt + len(testFlag)

	for _, tc := range []struct {
		name       string
		col        int
		next, prev int
		start      int
	}{
		{name: "line start", col: 0, next: 1, prev: 0, start: 0},
		{name: "cluster start", col: 1, next: flagAt, prev: 0, start: 1},
		{name: "inside the emoji", col: 5, next: flagAt, prev: 1, start: 1},
		{name: "between the flags letters", col: flagAt + 4, next: bAt, prev: flagAt, start: flagAt},
		{name: "last cluster", col: bAt, next: len(line), prev: flagAt, start: bAt},
		{name: "line end", col: len(line), next: len(line), prev: bAt, start: len(line)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := NextClusterEnd(line, tc.col); got != tc.next {
				t.Errorf("NextClusterEnd = %d, want %d", got, tc.next)
			}
			if got := PrevClusterStart(line, tc.col); got != tc.prev {
				t.Errorf("PrevClusterStart = %d, want %d", got, tc.prev)
			}
			if got := ClusterStart(line, tc.col); got != tc.start {
				t.Errorf("ClusterStart = %d, want %d", got, tc.start)
			}
		})
	}
}

func TestClusterEnd(t *testing.T) {
	line := "a" + testMermaid + "b"
	for _, tc := range []struct {
		name string
		col  int
		want int
	}{
		{name: "line start", col: 0, want: 0},
		{name: "cluster start", col: 1, want: 1},
		{name: "inside a cluster", col: 5, want: 1 + len(testMermaid)},
		{name: "after a cluster", col: 1 + len(testMermaid), want: 1 + len(testMermaid)},
		{name: "line end", col: len(line), want: len(line)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := ClusterEnd(line, tc.col); got != tc.want {
				t.Errorf("ClusterEnd = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestClipColumns(t *testing.T) {
	for _, tc := range []struct {
		name     string
		s        string
		x        int
		from, to int
		want     string
	}{
		{name: "narrow", s: "abcd", from: 1, to: 3, want: "bc"},
		{name: "cyrillic", s: "кіт", from: 1, to: 2, want: "і"},
		{name: "wide cut on the left", s: "🌊a", from: 1, to: 9, want: " a"},
		{name: "wide cut on the right", s: "a🌊", from: 0, to: 2, want: "a "},
		{name: "wide hidden", s: "🌊a", from: 2, to: 9, want: "a"},
		{name: "starting off the column 0", s: "🌊a", x: 3, from: 4, to: 9, want: " a"},
		{name: "joined emoji whole", s: testMermaid + "b", from: 0, to: 3, want: testMermaid + "b"},
		{name: "combining stays with its letter", s: testAccent + "б", from: 0, to: 1, want: testAccent},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := ClipColumns(tc.s, tc.x, tc.from, tc.to); got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
