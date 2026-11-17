package content

import (
	"strings"
	"testing"
)

func TestTree(t *testing.T) {
	var jumped []string
	jump := func(to string) LineAction {
		return LineActionFunc(func() { jumped = append(jumped, to) })
	}
	var loaded []string
	load := func(it *TreeItem) {
		loaded = append(loaded, it.Text)
		it.SetChildren(&TreeItem{Text: "loaded of " + it.Text, Action: jump(it.Text + "/loaded")})
	}

	tree := NewTree(
		&TreeItem{Text: "title"},
		&TreeItem{Text: "a.go", Expanded: true, Action: jump("a.go"), Load: load},
		&TreeItem{Text: "b.go", Load: load},
	)
	tree.Add(&TreeItem{Text: "c.go", Expanded: true})
	tree.Roots()[1].SetChildren(
		&TreeItem{Text: "1:2 x := a()", Action: jump("a.go:1")},
		&TreeItem{Text: "3:4 a()", Action: jump("a.go:3")},
	)

	check := func(step string, want ...string) {
		t.Helper()
		var got []string
		for _, line := range tree.Lines() {
			got = append(got, line.String())
		}
		if tree.Len() != len(got) || strings.Join(got, "\n") != strings.Join(want, "\n") {
			t.Errorf("%s: tree (%d lines)\n%s\nwant\n%s", step, tree.Len(), strings.Join(got, "\n"), strings.Join(want, "\n"))
		}
	}
	check("initial",
		"  title",
		"~ a.go",
		"    1:2 x := a()",
		"    3:4 a()",
		"> b.go",
		"  c.go",
	)

	// Engaging a leaf runs its action, a title line does nothing.
	tree.Lines()[0].(LineAction).Engage()
	tree.Lines()[2].(LineAction).Engage()
	tree.Lines()[3].(LineAction).Engage()
	if want := "a.go:1,a.go:3"; strings.Join(jumped, ",") != want {
		t.Errorf("jumped %q, want %q", jumped, want)
	}

	// Engaging an item with children hides them, and shows them again; the action is not run.
	tree.Lines()[1].(LineAction).Engage()
	check("a.go collapsed", "  title", "> a.go", "> b.go", "  c.go")
	tree.Lines()[1].(LineAction).Engage()
	check("a.go expanded again", "  title", "~ a.go", "    1:2 x := a()", "    3:4 a()", "> b.go", "  c.go")
	if len(jumped) != 2 || len(loaded) != 0 {
		t.Errorf("jumped %q, loaded %q after toggling an item with children", jumped, loaded)
	}

	// The children are loaded on the first expansion only.
	tree.Lines()[4].(LineAction).Engage()
	check("b.go loaded", "  title", "~ a.go", "    1:2 x := a()", "    3:4 a()", "~ b.go", "    loaded of b.go", "  c.go")
	tree.Lines()[5].(LineAction).Engage()
	tree.Lines()[4].(LineAction).Engage()
	tree.Lines()[4].(LineAction).Engage()
	if want := "b.go"; strings.Join(loaded, ",") != want {
		t.Errorf("loaded %q, want %q", loaded, want)
	}
	if want := "a.go:1,a.go:3,b.go/loaded"; strings.Join(jumped, ",") != want {
		t.Errorf("jumped %q, want %q", jumped, want)
	}
	if line := tree.Lines()[5]; line.Len() != len(line.String()) || line.MimeType() != MimeTypeTextPlain {
		t.Errorf("line %q: len %d, mime type %q", line, line.Len(), line.MimeType())
	}
}

func TestTree_LoadLater(t *testing.T) {
	var pending *TreeItem
	tree := NewTree(&TreeItem{Text: "f", Load: func(it *TreeItem) { pending = it }})
	tree.Lines()[0].(LineAction).Engage()
	if pending == nil || tree.Len() != 1 || tree.Lines()[0].String() != "~ f" {
		t.Fatalf("after engaging: pending %v, lines %v", pending, tree.Lines())
	}
	// The children arrive from the background later, nested ones included.
	child := &TreeItem{Text: "child", Expanded: true}
	child.SetChildren(&TreeItem{Text: "grandchild"})
	pending.SetChildren(child)
	var got []string
	for _, line := range tree.Lines() {
		got = append(got, line.String())
	}
	if want := "~ f|  ~ child|      grandchild"; strings.Join(got, "|") != want {
		t.Errorf("lines %q, want %q", got, want)
	}
}

func TestTree_Style(t *testing.T) {
	tree := &Tree{Indent: "\t", MimeType: "text/x"}
	child := &TreeItem{Text: "child"}
	tree.Add(&TreeItem{Text: "dir", Expanded: true})
	tree.Roots()[0].SetChildren(child)
	if got, want := tree.Lines()[1].String(), "\t  child"; got != want || child.MimeType() != "text/x" {
		t.Errorf("line %q (%s), want %q", got, child.MimeType(), want)
	}
}
