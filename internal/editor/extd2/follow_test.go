package extd2

import (
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl"
)

func TestFollowLine(t *testing.T) {
	prev := []string{"a", "b", "c", "d", "e"}
	for _, tc := range []struct {
		name   string
		cur    []string
		line   int
		want   int
		wantOK bool
	}{
		{name: "no changes", cur: prev, line: 2, want: 2, wantOK: true},
		{name: "insert above", cur: []string{"a", "x", "y", "b", "c", "d", "e"}, line: 2, want: 4, wantOK: true},
		{name: "insert below", cur: []string{"a", "b", "c", "x", "d", "e"}, line: 2, want: 2, wantOK: true},
		{name: "delete above", cur: []string{"b", "c", "d", "e"}, line: 2, want: 1, wantOK: true},
		{name: "update above", cur: []string{"a", "x", "c", "d", "e"}, line: 2, want: 2, wantOK: true},
		{name: "update the line", cur: []string{"a", "b", "x", "d", "e"}, line: 2, want: 2, wantOK: true},
		{name: "replace lines around", cur: []string{"a", "x", "y", "z", "e"}, line: 2, want: 2, wantOK: true},
		{name: "delete the line", cur: []string{"a", "b", "d", "e"}, line: 2, want: 2, wantOK: false},
		{name: "delete range with the line", cur: []string{"a", "e"}, line: 2, want: 2, wantOK: false},
		{name: "append to the end", cur: []string{"a", "b", "c", "d", "e", "f"}, line: 4, want: 4, wantOK: true},
		{name: "insert at the start", cur: []string{"x", "a", "b", "c", "d", "e"}, line: 0, want: 1, wantOK: true},
		{name: "clear", cur: nil, line: 2, want: 2, wantOK: false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := followLine(tc.line, prev, tc.cur)
			if got != tc.want || ok != tc.wantOK {
				t.Errorf("followLine(%d) = %d, %t; want %d, %t", tc.line, got, ok, tc.want, tc.wantOK)
			}
		})
	}
}

func TestMarkdown_FollowBlock(t *testing.T) {
	ext, viewer := newIntegration(t)
	buf := openDoc(t, ext, "a.md", strings.Join([]string{
		"# Notes", // 0
		"```d2",   // 1
		"a -> b",  // 2
		"```",     // 3
		"```d2",   // 4
		"x -> y",  // 5
		"```",     // 6
	}, "\n"))

	action := buf.ExtensionData(ext.ID()).(content.LineActions).LineAction(4)
	rerun := func(wantSource string) {
		t.Helper()
		action.Engage()
		if got := viewer.receive(t); got.Source != wantSource {
			t.Errorf("shown source %q, want %q", got.Source, wantSource)
		}
	}
	rerun("x -> y")

	// Insert a new block above, like a paste does.
	m := buf.Mutate()
	m.Insert(0, content.TextLine("```d2"))
	m.Insert(1, content.TextLine("new"))
	m.Insert(2, content.TextLine("```"))
	ext.AfterEdit(nil, buf)
	rerun("x -> y")

	// Edit the diagram and the text above it. The editor calls AfterEdit after every input.
	m.Update(7, content.TextLine("```d2 {x}"))
	ext.AfterEdit(nil, buf)
	m.Update(8, content.TextLine("x -> z"))
	ext.AfterEdit(nil, buf)
	m.Delete(3)
	ext.AfterEdit(nil, buf)
	rerun("x -> z")

	// Remove the block.
	for range 3 {
		m.Delete(6)
	}
	ext.AfterEdit(nil, buf)
	action.Engage()
	select {
	case d := <-viewer:
		t.Errorf("shown %+v for a removed block", d)
	default:
	}

	// The same lines are added back, but the removed block is not resurrected.
	m.Insert(6, content.TextLine("```d2"))
	m.Insert(7, content.TextLine("```"))
	ext.AfterEdit(nil, buf)
	action.Engage()
	select {
	case d := <-viewer:
		t.Errorf("shown %+v for a removed block", d)
	default:
	}
}

// TestRerunWithCtrlR checks that re-running the action in the editor follows the block.
func TestRerunWithCtrlR(t *testing.T) {
	const ctrlR = 0x12
	ext, viewer := newIntegration(t)
	h := editor.NewTestHarness()
	h.Extend(new(extsyntaxhl.Integration))
	h.Extend(ext)
	if err := h.OpenReader("a.md", strings.NewReader("```d2\na -> b\n```")); err != nil {
		t.Fatal(err)
	}
	h.Run(t)

	h.SendInput(t, []byte{'\r'})
	viewer.receive(t)

	// Add a line above the block: Enter at the start of the first line in insert mode.
	// The cursor stays on the fence line.
	h.SendInputSequence(t, "i\r\x1b")
	// Edit the diagram.
	h.SendInputSequence(t, "j$a -> c\x1b")
	h.SendInput(t, []byte{ctrlR})
	if got := viewer.receive(t); got.Source != "a -> b -> c" {
		t.Errorf("shown source %q", got.Source)
	}
}
