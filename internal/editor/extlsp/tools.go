package extlsp

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode/utf8"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/extrefactor"
	"rmazur.io/chernetka/internal/editor/navigation"
)

// The navigation tools listing what the server finds in the tool buffer: the
// usages of a symbol, and the callers of a function.

// FindUsages lists the references to the symbol at pos found by the server in
// the tool buffer, by file, with a jump to each.
func (lbd *BufferData) FindUsages(loop editor.Sender, pos content.Position) {
	lines := lbd.buf.Content.Lines()
	if pos.Line >= len(lines) {
		return
	}
	line := lines[pos.Line].String()
	at := lspPosition(line, pos)
	symbol := identAround(line, pos.Col)
	text := lbd.buf.Text()
	srv := lbd.srv
	lbd.request(loop, func(ctx context.Context, client lspClient) {
		locs, err := client.References(ctx, lbd.docUri, at)
		if err != nil {
			srv.showMessage(loop, fmt.Sprintf("usages of %s: %s", symbol, err))
			return
		}
		if len(locs) == 0 {
			srv.showMessage(loop, fmt.Sprintf("no usages of %s", symbol))
			return
		}
		files := srv.filesOf(locs, lbd.docUri, text)
		tree := content.NewTree(&content.TreeItem{Text: fmt.Sprintf("%d usages of %s", len(locs), symbol)})
		for _, f := range files.sorted() {
			item := &content.TreeItem{Text: srv.displayPath(f), Expanded: true}
			item.SetChildren(files.jumps(loop, f)...)
			tree.Add(item)
		}
		loop.Send(&editor.ShowTool{Content: tree})
	})
}

// FindCallers lists the functions calling the one at pos found by the server in
// the tool buffer, with a jump to each call. The callers of a caller are listed
// below it on demand.
func (lbd *BufferData) FindCallers(loop editor.Sender, pos content.Position) {
	lines := lbd.buf.Content.Lines()
	if pos.Line >= len(lines) {
		return
	}
	line := lines[pos.Line].String()
	at := lspPosition(line, pos)
	symbol := identAround(line, pos.Col)
	srv := lbd.srv
	lbd.request(loop, func(ctx context.Context, client lspClient) {
		items, err := client.PrepareCallHierarchy(ctx, lbd.docUri, at)
		if err != nil {
			srv.showMessage(loop, fmt.Sprintf("callers of %s: %s", symbol, err))
			return
		}
		if len(items) == 0 {
			srv.showMessage(loop, fmt.Sprintf("no function at %s", symbol))
			return
		}
		fn := items[0]
		callers, err := srv.callersOf(ctx, client, loop, fn)
		if err != nil {
			srv.showMessage(loop, fmt.Sprintf("callers of %s: %s", fn.Name, err))
			return
		}
		tree := content.NewTree(&content.TreeItem{Text: fmt.Sprintf("%d callers of %s", len(callers), fn.Name)})
		tree.Add(callers...)
		loop.Send(&editor.ShowTool{Content: tree})
	})
}

// callersOf returns the items listing the callers of fn: each one is expanded to
// the calls it makes, and holds an item loading its own callers.
func (srv *server) callersOf(ctx context.Context, client lspClient, loop editor.Sender, fn protocol.CallHierarchyItem) ([]*content.TreeItem, error) {
	calls, err := client.IncomingCalls(ctx, fn)
	if err != nil {
		return nil, err
	}
	var res []*content.TreeItem
	for _, call := range calls {
		caller := call.From
		var locs []protocol.Location
		for _, r := range call.FromRanges {
			locs = append(locs, protocol.Location{URI: caller.URI, Range: r})
		}
		files := srv.filesOf(locs, "", "")
		item := &content.TreeItem{
			Text:     fmt.Sprintf("%s  %s:%d", caller.Name, srv.displayPath(caller.URI.Filename()), caller.SelectionRange.Start.Line+1),
			Expanded: true,
		}
		children := files.jumps(loop, caller.URI.Filename())
		children = append(children, &content.TreeItem{
			Text: "callers of " + caller.Name,
			Load: func(it *content.TreeItem) { srv.loadCallers(loop, it, caller) },
		})
		item.SetChildren(children...)
		res = append(res, item)
	}
	return res, nil
}

// loadCallers fills the item with the callers of fn in the background.
func (srv *server) loadCallers(loop editor.Sender, it *content.TreeItem, fn protocol.CallHierarchyItem) {
	srv.background(func(ctx context.Context, client lspClient) {
		callers, err := srv.callersOf(ctx, client, loop, fn)
		if err != nil {
			srv.le.Logf("callers of %s: %s", fn.Name, err)
			callers = []*content.TreeItem{{Text: "error: " + err.Error()}}
		}
		if len(callers) == 0 {
			callers = []*content.TreeItem{{Text: "none"}}
		}
		loop.Send(editor.CommandFunc(func(e *editor.Editor) {
			it.SetChildren(callers...)
			e.RequestRender()
		}))
	})
}

// showMessage shows the message in the tool buffer and logs it.
func (srv *server) showMessage(loop editor.Sender, msg string) {
	srv.le.Logf("%s", msg)
	loop.Send(&editor.ShowTool{Content: content.NewTree(&content.TreeItem{Text: msg})})
}

// displayPath returns the path relative to the workspace root if it's in it.
func (srv *server) displayPath(path string) string {
	rel, err := filepath.Rel(srv.rootDir, path)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return path
	}
	return rel
}

// locatedFiles are the locations found by the server grouped by file, with the
// text of the files to show the lines and to convert the columns.
type locatedFiles struct {
	locs  map[string][]protocol.Location // by path
	texts map[string]string              // the text of the files, "" if it can't be read
}

// filesOf groups the locations by file, reading the files. The document of docUri
// has the text given: the server has it, and the file may be older.
func (srv *server) filesOf(locs []protocol.Location, docUri uri.URI, text string) *locatedFiles {
	lf := &locatedFiles{locs: make(map[string][]protocol.Location), texts: make(map[string]string)}
	for _, loc := range locs {
		path := loc.URI.Filename()
		lf.locs[path] = append(lf.locs[path], loc)
		if _, ok := lf.texts[path]; ok {
			continue
		}
		if loc.URI == docUri {
			lf.texts[path] = text
			continue
		}
		data, err := os.ReadFile(path)
		if err != nil {
			srv.le.Logf("cannot read %s: %s", path, err)
		}
		lf.texts[path] = string(data)
	}
	return lf
}

// sorted returns the paths of the files, sorted.
func (lf *locatedFiles) sorted() []string {
	paths := make([]string, 0, len(lf.locs))
	for p := range lf.locs {
		paths = append(paths, p)
	}
	slices.Sort(paths)
	return paths
}

// jumps returns the items jumping to the locations in the file, in their order,
// showing the line of each. The jump is sent to the loop from another goroutine:
// the item is engaged on the loop, which can't wait for its own queue.
func (lf *locatedFiles) jumps(loop editor.Sender, path string) []*content.TreeItem {
	locs := lf.locs[path]
	slices.SortFunc(locs, func(a, b protocol.Location) int {
		if positionBefore(a.Range.Start, b.Range.Start) {
			return -1
		}
		if positionBefore(b.Range.Start, a.Range.Start) {
			return 1
		}
		return 0
	})
	text := lf.texts[path]
	lines := strings.Split(text, "\n")
	var res []*content.TreeItem
	for i, loc := range locs {
		if i > 0 && loc.Range.Start == locs[i-1].Range.Start {
			continue // The server may report a place twice, like a method and its interface.
		}
		var line string
		if ln := int(loc.Range.Start.Line); ln < len(lines) {
			line = strings.TrimSpace(lines[ln])
		}
		goTo := &editor.GoTo{Location: navigation.Location{Path: path, Position: textPosition(text, loc.Range.Start)}}
		res = append(res, &content.TreeItem{
			Text:   fmt.Sprintf("%d:%d  %s", loc.Range.Start.Line+1, loc.Range.Start.Character+1, line),
			Action: content.LineActionFunc(func() { go loop.Send(goTo) }),
		})
	}
	return res
}

// identAround returns the identifier around the byte offset in line, or what's
// at the offset if there is none, to name what the tools are asked about.
func identAround(line string, col int) string {
	col = min(col, len(line))
	for col > 0 && col < len(line) && !utf8.RuneStart(line[col]) {
		col-- // The start of the character.
	}
	start := len(line[:col]) - len(identTrailing(line[:col]))
	end := col
	for end < len(line) && identAt(line, end) {
		_, size := utf8.DecodeRuneInString(line[end:])
		end += size
	}
	if start == end {
		return fmt.Sprintf("column %d", col+1)
	}
	return line[start:end]
}

var (
	_ extrefactor.UsageFinder  = new(BufferData)
	_ extrefactor.CallerFinder = new(BufferData)
)
