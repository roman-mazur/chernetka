package content

import (
	"io/fs"
	"os"
	"path"
	"slices"
	"strings"
)

type OpenFile interface {
	OpenFile(p string)
}

// LoadFolder lists the directory as a Tree: the directories are expanded when they
// are engaged, and the files are opened with open.
func LoadFolder(dir string, open OpenFile) *FsContent {
	fc := &FsContent{
		Tree:     Tree{Indent: "\t", MimeType: MimeTypeFileName},
		rootPath: dir,
		root:     os.DirFS(dir),
		open:     open,
		entries:  make(map[*TreeItem]fsEntry),
	}
	fc.Add(fc.list(".")...)
	return fc
}

// MimeTypeFileName is the mime type of the lines of a folder listing.
const MimeTypeFileName = "text/filename"

// FsContent implements the content Document for a particular directory path.
type FsContent struct {
	Tree

	rootPath string
	root     fs.FS
	open     OpenFile

	entries map[*TreeItem]fsEntry // what the items of the tree show
}

// fsEntry is what a line of the listing shows: an entry of a directory, or the
// error of listing the directory if entry is nil.
type fsEntry struct {
	dir   string // slash-separated, relative to the root
	entry fs.DirEntry
}

// RootPath returns the directory the content was loaded from.
func (fc *FsContent) RootPath() string { return fc.rootPath }

// DirAt returns the directory selected with the line, as a slash-separated path relative
// to the root ("." for the root itself). It's the directory shown on the line,
// or the parent of the entry shown there otherwise.
func (fc *FsContent) DirAt(line int) string {
	if line < 0 || line >= fc.Len() {
		return "."
	}
	e := fc.entries[fc.Lines()[line].(*TreeItem)]
	if e.entry != nil && e.entry.IsDir() {
		return path.Join(e.dir, e.entry.Name())
	}
	return e.dir
}

// ExpandedDirs returns the directories whose entries are listed, as slash-separated paths
// relative to the root. The root (".") is always expanded.
func (fc *FsContent) ExpandedDirs() []string {
	dirs := []string{"."}
	var visit func(items []*TreeItem)
	visit := func(items []*TreeItem) {
		for _, it := range items {
			if it.Expanded {
				dirs = append(dirs, fc.DirAt(slices.Index(fc.Lines(), Line(it))))
				visit(it.Children())
			}
		}
	}
	visit(fc.Roots())
	return dirs
}

// DirEntries returns the sorted names of the entries listed for the directory, a slash-separated
// path relative to the root, as the ones returned by ExpandedDirs.
func (fc *FsContent) DirEntries(dir string) []string {
	var names []string
	for it, e := range fc.entries {
		if e.dir == dir && e.entry != nil && slices.Contains(fc.Lines(), Line(it)) {
			names = append(names, e.entry.Name())
		}
	}
	slices.Sort(names)
	return names
}

// SyncState expands the directories of fc that are expanded in origin, which lists
// the same root.
func (fc *FsContent) SyncState(origin *FsContent) {
	if fc.rootPath != origin.rootPath {
		return
	}
	var sync func(items, originItems []*TreeItem)
	sync = func(items, originItems []*TreeItem) {
		for _, o := range originItems {
			if !o.Expanded {
				continue
			}
			i := slices.IndexFunc(items, func(it *TreeItem) bool { return it.Text == o.Text && fc.entries[it].entry != nil })
			if i < 0 {
				continue
			}
			if !items[i].Expanded {
				items[i].Engage()
			}
			sync(items[i].Children(), o.Children())
		}
	}
	sync(fc.Roots(), origin.Roots())
}

// list returns the items of the entries of the directory, the directories first,
// or the item showing the error of listing it.
func (fc *FsContent) list(dirPath string) []*TreeItem {
	entries, err := fs.ReadDir(fc.root, dirPath)
	if err != nil {
		it := &TreeItem{Text: err.Error()}
		fc.entries[it] = fsEntry{dir: dirPath}
		return []*TreeItem{it}
	}
	slices.SortFunc(entries, func(a, b fs.DirEntry) int {
		if a.IsDir() != b.IsDir() {
			if a.IsDir() {
				return -1
			}
			return 1
		}
		return strings.Compare(a.Name(), b.Name())
	})

	items := make([]*TreeItem, len(entries))
	for i, entry := range entries {
		it := &TreeItem{Text: entry.Name()}
		entryPath := path.Join(dirPath, entry.Name())
		if entry.IsDir() {
			it.Load = func(it *TreeItem) { it.SetChildren(fc.list(entryPath)...) }
		} else {
			it.Action = LineActionFunc(func() {
				if fc.open != nil {
					fc.open.OpenFile(entryPath)
				}
			})
		}
		fc.entries[it] = fsEntry{dir: dirPath, entry: entry}
		items[i] = it
	}
	return items
}
