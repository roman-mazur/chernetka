package content

import "strings"

// Tree is a Document of items that may have children: the directories and the
// files of a folder, or the usages of a symbol by file. An item with children
// shows or hides them when it's engaged, and the children may be loaded when
// the item is expanded for the first time. The other items run their action
// when engaged.
//
// A Tree is changed on the editor loop only.
type Tree struct {
	// Indent is repeated before an item for every level of nesting. Two spaces if empty.
	Indent string
	// MimeType is the one of the lines, MimeTypeTextPlain if empty.
	MimeType string

	roots []*TreeItem
	lines []Line // the visible items, in the order they are shown
}

// NewTree creates a Tree of the items.
func NewTree(items ...*TreeItem) *Tree {
	t := new(Tree)
	t.Add(items...)
	return t
}

// TreeItem is what Tree consists of. The exported fields are set before the item is added
// to a Tree, or to the children of an item: the Tree uses them to compose a text Line for display.
type TreeItem struct {
	// Text is shown on the line, after the indentation of the item and the marker of its children.
	Text string
	// Expanded shows the children of the item.
	Expanded bool
	// Action runs when the item is engaged, unless the item has children, or loads them.
	Action LineAction
	// Load fetches the children of the item when it's expanded for the first time: in the
	// background, if it takes a while. The children are set with SetChildren.
	Load func(it *TreeItem)

	tree     *Tree
	level    int
	children []*TreeItem
	loaded   bool   // Load was called
	display  string // the cached result of String
}

func (t *Tree) Len() int      { return len(t.lines) }
func (t *Tree) Lines() []Line { return t.lines }

// Roots returns the items of the top level.
func (t *Tree) Roots() []*TreeItem { return t.roots }

// Add appends the items to the top level of the tree.
func (t *Tree) Add(items ...*TreeItem) {
	for _, it := range items {
		it.attach(t, 0)
	}
	t.roots = append(t.roots, items...)
	t.flatten()
}

// attach makes the item and its children a part of the tree at the level.
func (it *TreeItem) attach(t *Tree, level int) {
	it.tree, it.level, it.display = t, level, ""
	for _, child := range it.children {
		child.attach(t, level+1)
	}
}

// flatten lists the visible items. The lines returned before are left as they
// were: a caller may hold them while it engages the items.
func (t *Tree) flatten() {
	t.lines = make([]Line, 0, len(t.lines))
	var visit func(items []*TreeItem)
	visit = func(items []*TreeItem) {
		for _, it := range items {
			t.lines = append(t.lines, it)
			if it.Expanded {
				visit(it.children)
			}
		}
	}
	visit(t.roots)
}

// SetChildren replaces the children of the item. The children are shown if the item is expanded,
// and they are not loaded anymore.
func (it *TreeItem) SetChildren(children ...*TreeItem) {
	it.children = children
	it.loaded = true
	it.display = ""
	if it.tree != nil {
		for _, child := range children {
			child.attach(it.tree, it.level+1)
		}
		it.tree.flatten()
	}
}

// Children returns the children of the item.
func (it *TreeItem) Children() []*TreeItem { return it.children }

// expandable reports whether the item shows children when it's engaged.
func (it *TreeItem) expandable() bool { return len(it.children) > 0 || it.Load != nil }

// Engage shows or hides the children of the item, loading them the first time, or runs
// its action if it has no children.
func (it *TreeItem) Engage() {
	if !it.expandable() {
		if it.Action != nil {
			it.Action.Engage()
		}
		return
	}
	it.Expanded = !it.Expanded
	it.display = ""
	if it.Expanded && it.Load != nil && !it.loaded {
		it.loaded = true
		it.Load(it)
	}
	if it.tree != nil {
		it.tree.flatten()
	}
}

// SingleShot marks the item not to be re-run as the last engaged action: engaging it
// again would hide what it has shown, or jump to the same place.
func (it *TreeItem) SingleShot() {}

// The markers of an item with children, hidden and shown.
const (
	markerCollapsed = "> "
	markerExpanded  = "~ "
	markerNone      = "  "
	defaultIndent   = "  "
)

func (it *TreeItem) String() string {
	if it.display == "" {
		marker := markerNone
		switch {
		case !it.expandable():
		case it.Expanded:
			marker = markerExpanded
		default:
			marker = markerCollapsed
		}
		indent := defaultIndent
		if it.tree != nil && it.tree.Indent != "" {
			indent = it.tree.Indent
		}
		it.display = strings.Repeat(indent, it.level) + marker + it.Text
	}
	return it.display
}

func (it *TreeItem) Len() int { return len(it.String()) }

func (it *TreeItem) MimeType() string {
	if it.tree != nil && it.tree.MimeType != "" {
		return it.tree.MimeType
	}
	return MimeTypeTextPlain
}

var (
	_ Document             = new(Tree)
	_ LineAction           = new(TreeItem)
	_ LineActionSingleShot = new(TreeItem)
)
