package content

import "testing"

type testAction string

func (testAction) Engage() {}

// actionLine is a Line that is actionable by itself.
type actionLine struct {
	TextLine
	testAction
}

// actionsMap provides actions for the lines by their index.
type actionsMap map[int]LineAction

func (am actionsMap) LineAction(lineNumber int) LineAction { return am[lineNumber] }

// actionableDoc is a Document that provides actions for its lines.
type actionableDoc struct {
	FullText
	actionsMap
}

func TestActionAt(t *testing.T) {
	self := actionLine{TextLine: "self", testAction: "line"}
	lines := FullText{
		TextLine("plain"),
		self,
		TextLine("provided"),
	}

	for _, tc := range []struct {
		name      string
		doc       Document
		providers []LineActions
		line      int
		want      LineAction
	}{
		{name: "plain line", doc: &lines, line: 0, want: nil},
		{name: "line is actionable", doc: &lines, line: 1, want: self},
		{name: "negative index", doc: &lines, line: -1, want: nil},
		{name: "index out of range", doc: &lines, line: 3, want: nil},
		{
			name:      "provider",
			doc:       &lines,
			providers: []LineActions{actionsMap{2: testAction("provider")}},
			line:      2,
			want:      testAction("provider"),
		},
		{
			name:      "line wins over provider",
			doc:       &lines,
			providers: []LineActions{actionsMap{1: testAction("provider")}},
			line:      1,
			want:      self,
		},
		{
			name:      "document wins over provider",
			doc:       &actionableDoc{FullText: lines, actionsMap: actionsMap{2: testAction("doc")}},
			providers: []LineActions{actionsMap{2: testAction("provider")}},
			line:      2,
			want:      testAction("doc"),
		},
		{
			name: "first provider with an action wins",
			doc:  &lines,
			providers: []LineActions{
				actionsMap{},
				actionsMap{0: testAction("second")},
				actionsMap{0: testAction("third")},
			},
			line: 0,
			want: testAction("second"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := ActionAt(tc.doc, tc.line, tc.providers...); got != tc.want {
				t.Errorf("ActionAt(%d) = %v, want %v", tc.line, got, tc.want)
			}
		})
	}
}
