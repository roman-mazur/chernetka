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
	// The commands and their shortcuts are followed by their arguments, and maybe
	// a comment. An unknown command is not highlighted.
	got := scan(t, ScanRebaseTodo, "pick 1234 change # note\np abc fix\nx make test\nnope 12 x\n# Commands:")
	want := []string{
		`0 "pick" Keyword`,
		`0 "1234 change" NumberLiteral`,
		`0 " # note" Comment`,
		`1 "p" Keyword`,
		`1 "abc fix" NumberLiteral`,
		`2 "x" Keyword`,
		`2 "make test" NumberLiteral`,
		`4 "# Commands:" Comment`,
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("got\n%s\nwant\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}
