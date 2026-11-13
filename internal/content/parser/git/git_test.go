package git

import (
	"fmt"
	"strings"
	"testing"
)

// scan returns the spans the scanner emits for the document, each as its line,
// its text and its token, checking they are within the lines.
func scan(t *testing.T, scanner func([]string, func(Span)), src string) []string {
	t.Helper()
	lines := strings.Split(src, "\n")
	var got []string
	scanner(lines, func(s Span) {
		if s.Line < 0 || s.Line >= len(lines) || s.Start < 0 || s.End > len(lines[s.Line]) || s.Start >= s.End {
			t.Fatalf("span %+v is out of range", s)
		}
		got = append(got, fmt.Sprintf("%d %q %s", s.Line, lines[s.Line][s.Start:s.End], s.Token))
	})
	return got
}

func TestScanCommitMessage(t *testing.T) {
	for _, tc := range []struct {
		name, src string
		want      []string
	}{
		{
			// The subject comes first, and the trailers after an empty line, even
			// one with no value.
			name: "message",
			src:  "pkg: change\n\nBody text.\n\nSigned-off-by: me\nKey:\n\n# comment",
			want: []string{
				`0 "pkg: change" Heading`,
				`4 "Signed-off-by:" Field`,
				`4 "me" StringLiteral`,
				`5 "Key:" Field`,
				`7 "# comment" Comment`,
			},
		},
		{
			// The value of a trailer follows the colon and the spaces after it, if any.
			name: "trailer values",
			src:  "pkg: change\n\nNo-space:x\nTabs:\t\tv",
			want: []string{
				`0 "pkg: change" Heading`,
				`2 "No-space:" Field`,
				`2 "x" StringLiteral`,
				`3 "Tabs:" Field`,
				`3 "v" StringLiteral`,
			},
		},
		{
			// The template git opens for a new commit.
			name: "template",
			src:  "\n# Please enter the commit message\n# On branch main",
			want: []string{
				`1 "# Please enter the commit message" Comment`,
				`2 "# On branch main" Comment`,
			},
		},
		{
			// The comments are stripped by git, so they don't separate the subject
			// from the body. A line that is not a trailer ends the trailers.
			name: "comment after subject",
			src:  "pkg: change\n# comment\n\n# comment\nKey: value\nnot a trailer\nKey: value",
			want: []string{
				`0 "pkg: change" Heading`,
				`1 "# comment" Comment`,
				`3 "# comment" Comment`,
				`4 "Key:" Field`,
				`4 "value" StringLiteral`,
			},
		},
		{
			name: "no distinct subject",
			src:  "first line\nsecond line\n\nKey: value\n# comment",
			want: []string{
				`4 "# comment" Comment`,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := scan(t, ScanCommitMessage, tc.src)
			if strings.Join(got, "\n") != strings.Join(tc.want, "\n") {
				t.Errorf("got\n%s\nwant\n%s", strings.Join(got, "\n"), strings.Join(tc.want, "\n"))
			}
		})
	}
}

func TestScanRebaseTodo(t *testing.T) {
	// The commands and their shortcuts, maybe after spaces, are followed by the
	// commits they take, maybe after a flag, which are followed by plain
	// subjects, and maybe a comment. The commands taking no commit, or no
	// arguments at all, are highlighted too, but not the unknown ones.
	got := scan(t, ScanRebaseTodo, strings.Join([]string{
		"pick 1234 change # note",
		"p abc fix",
		"  reword 5678 indented",
		"fixup -C 9abc amend",
		"merge -C def0 topic # Merge topic",
		"m topic",
		"x make test",
		"update-ref refs/heads/topic",
		"break",
		"nope 12 x",
		"  # Commands:",
	}, "\n"))
	want := []string{
		`0 "pick" Keyword`,
		`0 "1234" NumberLiteral`,
		`0 "# note" Comment`,
		`1 "p" Keyword`,
		`1 "abc" NumberLiteral`,
		`2 "reword" Keyword`,
		`2 "5678" NumberLiteral`,
		`3 "fixup" Keyword`,
		`3 "-C" Constant`,
		`3 "9abc" NumberLiteral`,
		`4 "merge" Keyword`,
		`4 "-C" Constant`,
		`4 "def0" NumberLiteral`,
		`4 "# Merge topic" Comment`,
		`5 "m" Keyword`,
		`6 "x" Keyword`,
		`7 "update-ref" Keyword`,
		`8 "break" Keyword`,
		`10 "# Commands:" Comment`,
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("got\n%s\nwant\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}
