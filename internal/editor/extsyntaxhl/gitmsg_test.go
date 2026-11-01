package extsyntaxhl

import (
	"slices"
	"testing"

	"rmazur.io/chernetka/internal/content/code"
)

func TestGitMessageHighlight(t *testing.T) {
	for _, tc := range []struct {
		name string
		doc  hlDoc
	}{
		{"message", hlDoc{
			{"pkg: change", []string{"0:11:Heading"}},
			{"", nil},
			{"Body text.", nil},
			{"", nil},
			{"Signed-off-by: me", []string{"0:14:Field", "15:17:StringLiteral"}},
			{"", nil},
			{"# comment", []string{"0:9:Comment"}},
		}},
		// The template git opens for a new commit.
		{"template", hlDoc{
			{"", nil},
			{"# Please enter the commit message", []string{"0:33:Comment"}},
			{"# On branch main", []string{"0:16:Comment"}},
		}},
		{"comment first", hlDoc{
			{"# comment", []string{"0:9:Comment"}},
			{"pkg: change", []string{"0:11:Heading"}},
			{"", nil},
			{"Key: value", []string{"0:4:Field", "5:10:StringLiteral"}},
		}},
		// Comments are stripped by git, so they don't separate the subject from the body.
		{"comment after subject", hlDoc{
			{"pkg: change", []string{"0:11:Heading"}},
			{"# comment", []string{"0:9:Comment"}},
			{"", nil},
			{"# comment", []string{"0:9:Comment"}},
			{"Key: value", []string{"0:4:Field", "5:10:StringLiteral"}},
		}},
		{"no distinct subject", hlDoc{
			{"first line", nil},
			{"second line", nil},
			{"", nil},
			{"Key: value", nil},
			{"# comment", []string{"0:9:Comment"}},
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			buf, ext := openDoc(t, ".git/COMMIT_EDITMSG", tc.doc.text())
			tc.doc.check(t, highlighterOf(t, buf, ext))
		})
	}
}

func TestGitRebaseReparse(t *testing.T) {
	var gr gitRebase
	for _, text := range []string{"pick 1234 change\n# comment", "pick 1234 change"} {
		if err := gr.reparse(newSource(text)); err != nil {
			t.Fatal(err)
		}
	}
	// Only the spans of the last text: no comment line left over from the first one.
	var spans []rawSpan
	gr.spans(nil, func(s rawSpan) { spans = append(spans, s) })
	want := []rawSpan{
		{EndCol: 4, TokenType: code.TtKeyword},
		{StartCol: 5, EndCol: 16, TokenType: code.TtNumberLiteral},
	}
	if !slices.Equal(spans, want) {
		t.Errorf("got spans %v, want %v", spans, want)
	}
}
