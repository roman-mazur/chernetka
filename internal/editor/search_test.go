package editor

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/styles"
	"rmazur.io/chernetka/internal/vt/escape"
)

func TestParseSearch(t *testing.T) {
	for _, tc := range []struct {
		input, pattern, template string
		replace                  bool
	}{
		{input: "", pattern: ""},
		{input: "foo", pattern: "foo"},
		{input: `a\.b`, pattern: `a\.b`},
		{input: `a\/b`, pattern: "a/b"},
		{input: "foo/", pattern: "foo", replace: true},
		{input: "foo/bar", pattern: "foo", template: "bar", replace: true},
		{input: "foo/bar/", pattern: "foo", template: "bar", replace: true},
		{input: `foo/a\/b/`, pattern: "foo", template: "a/b", replace: true},
		{input: `(\w+)=(\w+)/\2=\1`, pattern: `(\w+)=(\w+)`, template: "${2}=${1}", replace: true},
		{input: "x/[&]", pattern: "x", template: "[${0}]", replace: true},
		{input: `x/\&$`, pattern: "x", template: "&$$", replace: true},
		{input: `x/a\nb\t`, pattern: "x", template: "a\nb\t", replace: true},
	} {
		t.Run(tc.input, func(t *testing.T) {
			pattern, template, replace := parseSearch(tc.input)
			if pattern != tc.pattern || template != tc.template || replace != tc.replace {
				t.Errorf("parseSearch(%q) = %q, %q, %t; want %q, %q, %t",
					tc.input, pattern, template, replace, tc.pattern, tc.template, tc.replace)
			}
		})
	}
}

// newSearchEditor returns an editor with a buffer showing the lines, and the input function.
func newSearchEditor(t *testing.T, lines ...string) (*Editor, func(keys ...string)) {
	t.Helper()
	var e Editor
	if err := e.OpenReader("", strings.NewReader(strings.Join(lines, "\n"))); err != nil {
		t.Fatal(err)
	}
	return &e, func(keys ...string) {
		for _, k := range keys {
			if len(k) > 1 && k[0] != '\x1b' {
				for _, c := range []byte(k) {
					e.handleInput([]byte{c})
				}
				continue
			}
			e.handleInput([]byte(k))
		}
	}
}

func pos(line, col int) content.Position { return content.Position{Line: line, Col: col} }

func TestSearch(t *testing.T) {
	const (
		ctrlF = "\x06"
		esc   = "\x1b"
		enter = "\r"
		up    = "\x1b[A"
		down  = "\x1b[B"
		tab   = "\t"
	)
	sample := []string{
		"one two",
		"three two",
		"four",
		"two",
	}

	t.Run("incremental", func(t *testing.T) {
		e, input := newSearchEditor(t, sample...)
		b := e.Top()
		b.c = pos(1, 0)
		input("/")
		if b.mode != ModeCommand || b.cmdline != "/" {
			t.Fatalf("mode %s, cmdline %q", b.mode, b.cmdline)
		}
		input("t")
		if b.c != pos(1, 0) {
			t.Errorf("cursor at %s, want the match at the cursor", b.c)
		}
		input("w")
		if b.c != pos(1, 6) {
			t.Errorf("cursor at %s, want the first match after the cursor", b.c)
		}
		input("x")
		if b.c != pos(1, 0) || b.search.failure != "no matches" {
			t.Errorf("cursor at %s, failure %q", b.c, b.search.failure)
		}
		input("\x7f", enter)
		if b.mode != ModeNormal || b.c != pos(1, 6) || b.search.re == nil {
			t.Errorf("mode %s, cursor at %s, pattern %v", b.mode, b.c, b.search.re)
		}
	})

	t.Run("navigation", func(t *testing.T) {
		e, input := newSearchEditor(t, sample...)
		b := e.Top()
		input("/two")
		if b.c != pos(0, 4) {
			t.Fatalf("cursor at %s", b.c)
		}
		input(tab)
		if b.c != pos(1, 6) {
			t.Errorf("tab: cursor at %s", b.c)
		}
		input(down, ctrlF)
		if b.c != pos(0, 4) {
			t.Errorf("down, ctrl+f: cursor at %s, want wrapping to the first match", b.c)
		}
		input(up)
		if b.c != pos(3, 0) {
			t.Errorf("up: cursor at %s, want wrapping to the last match", b.c)
		}
		input(enter, "n")
		if b.c != pos(0, 4) {
			t.Errorf("n: cursor at %s", b.c)
		}
		input("N", "N")
		if b.c != pos(1, 6) {
			t.Errorf("N: cursor at %s", b.c)
		}
		input(esc)
		if b.search.re != nil {
			t.Error("search is not cleared with Esc")
		}
		input("n")
		if b.c != pos(1, 6) {
			t.Errorf("n without search: cursor at %s", b.c)
		}
	})

	t.Run("single match on the line", func(t *testing.T) {
		e, input := newSearchEditor(t, "a", "b")
		b := e.Top()
		input("/a", enter)
		input("n")
		if b.c != pos(0, 0) {
			t.Errorf("n: cursor at %s", b.c)
		}
		input("N")
		if b.c != pos(0, 0) {
			t.Errorf("N: cursor at %s", b.c)
		}
	})

	t.Run("cancel", func(t *testing.T) {
		e, input := newSearchEditor(t, sample...)
		b := e.Top()
		b.c = pos(2, 1)
		input("/two", esc)
		if b.mode != ModeNormal || b.c != pos(2, 1) || b.search.re != nil || b.cmdline != "" {
			t.Errorf("mode %s, cursor at %s, pattern %v, cmdline %q", b.mode, b.c, b.search.re, b.cmdline)
		}

		input("/t", "\x7f", "\x7f")
		if b.mode != ModeNormal || b.search.typing {
			t.Errorf("deleting the prefix: mode %s, typing %t", b.mode, b.search.typing)
		}
	})

	t.Run("does not run commands", func(t *testing.T) {
		e, input := newSearchEditor(t, "quit")
		input("/quit", enter)
		if e.quitRequested || e.Top() == nil {
			t.Error("the pattern is run as a command")
		}
	})

	t.Run("invalid pattern", func(t *testing.T) {
		e, input := newSearchEditor(t, sample...)
		b := e.Top()
		input("/tw(")
		if b.search.failure != "invalid pattern" || b.search.re.String() != "tw" {
			t.Errorf("failure %q, pattern %v", b.search.failure, b.search.re)
		}
		input(enter)
		if b.search.re != nil || b.mode != ModeNormal {
			t.Errorf("pattern %v, mode %s", b.search.re, b.mode)
		}
	})

	t.Run("ctrl+f in insert mode", func(t *testing.T) {
		e, input := newSearchEditor(t, sample...)
		b := e.Top()
		input("i", ctrlF)
		if b.mode != ModeCommand || !b.search.typing {
			t.Fatalf("mode %s, typing %t", b.mode, b.search.typing)
		}
		input("four", enter)
		if b.mode != ModeInsert || b.c != pos(2, 0) {
			t.Errorf("mode %s, cursor at %s", b.mode, b.c)
		}
		input(ctrlF, esc)
		if b.mode != ModeInsert {
			t.Errorf("cancel: mode %s", b.mode)
		}
	})

	t.Run("replace", func(t *testing.T) {
		e, input := newSearchEditor(t, sample...)
		b := e.Top()
		b.c = pos(1, 0)
		input("/t(w)o/[&\\1]")
		if b.Text() != strings.Join(sample, "\n") {
			t.Fatalf("replaced before Enter: %q", b.Text())
		}
		input(enter)
		want := "one [twow]\nthree [twow]\nfour\n[twow]"
		if b.Text() != want {
			t.Errorf("text %q, want %q", b.Text(), want)
		}
		if b.c != pos(1, 6) || b.mode != ModeNormal || b.search.re != nil || !b.dirty {
			t.Errorf("cursor at %s, mode %s, pattern %v, dirty %t", b.c, b.mode, b.search.re, b.dirty)
		}
	})

	t.Run("replace with new lines", func(t *testing.T) {
		e, input := newSearchEditor(t, "a,b", "c")
		b := e.Top()
		input(`/,/\n`, enter)
		if want := "a\nb\nc"; b.Text() != want {
			t.Errorf("text %q, want %q", b.Text(), want)
		}
	})

	t.Run("replace read-only", func(t *testing.T) {
		var e Editor
		e.OpenBuffer(&Buffer{Content: &content.ErrorContent{Error: errors.New("test")}})
		for _, k := range []string{"/", "e", "/", "x", enter} {
			e.handleInput([]byte(k)) // Must not panic.
		}
	})
}

func TestSearch_Highlight(t *testing.T) {
	b := &Buffer{Content: &content.FullText{content.TextLine("ab ab ab")}}
	b.startSearch()
	b.cmdline = "/ab"
	b.syncSearch()
	b.searchMove(1)
	b.sel = []content.Span{{Start: pos(0, 7), End: pos(0, 8)}}

	var cr contentPrinter
	cr.prepare(b, 0, 1, &RenderPrefs{TabSize: 4})
	got := cr.buildBgSpans(0, 8, false)

	c := styles.DefaultColors
	want := []colorSpan{
		{Span: content.Span{Start: pos(0, 0), End: pos(0, 2)}, color: c.SearchMatchBg},
		{Span: content.Span{Start: pos(0, 2), End: pos(0, 3)}},
		{Span: content.Span{Start: pos(0, 3), End: pos(0, 5)}, color: c.SearchCursorBg},
		{Span: content.Span{Start: pos(0, 5), End: pos(0, 7)}},
		{Span: content.Span{Start: pos(0, 7), End: pos(0, 8)}, color: c.TextSelectedBg},
	}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("span %d: got %v, want %v", i, got[i], want[i])
		}
	}
}

func TestStatusBar_Search(t *testing.T) {
	b := &Buffer{Content: &content.FullText{content.TextLine("abc")}, w: 40, h: 1}
	s := StatusBar{buf: b}
	b.startSearch()
	b.cmdline = "/x"
	b.syncSearch()

	var out bytes.Buffer
	s.Render(&out)
	got := escape.Clean(out.String())
	if !strings.Contains(got, "/x") || strings.Contains(got, ":/x") || !strings.Contains(got, "no matches") {
		t.Errorf("status %q", got)
	}

	b.cmdline = "/b/y"
	b.syncSearch()
	out.Reset()
	s.Render(&out)
	if got := escape.Clean(out.String()); !strings.Contains(got, "replace all") {
		t.Errorf("status %q", got)
	}

	b.cmdline = "/b"
	b.syncSearch()
	b.finishSearch()
	out.Reset()
	s.Render(&out)
	if got := escape.Clean(out.String()); !strings.Contains(got, "NORMAL") || !strings.Contains(got, "/b") {
		t.Errorf("status %q", got)
	}
}

func TestSearch_Render(t *testing.T) {
	b := &Buffer{Content: &content.FullText{
		content.TextLine("\tкіт ab"),
		content.TextLine("ab"),
	}, w: 40, h: 2}
	b.startSearch()
	b.cmdline = "/кіт|ab"
	b.syncSearch()
	b.finishSearch()

	var out bytes.Buffer
	b.Render(&out, &RenderPrefs{TabSize: 4})
	got := escape.Clean(out.String())
	if !strings.Contains(got, "    кіт ab") || !strings.Contains(got, "2 ab") {
		t.Errorf("rendered %q", got)
	}
	if n := strings.Count(out.String(), "48;2;158;122;26m"); n != 1 {
		t.Errorf("the match at the cursor is highlighted %d times, want once", n)
	}
	if n := strings.Count(out.String(), "48;2;92;75;30m"); n != 2 {
		t.Errorf("other matches are highlighted %d times, want twice", n)
	}
}
