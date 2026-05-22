package content

import (
	"io/fs"
	"os"
	"path"
	"slices"
	"strings"
)

// LineAction extends the Line adding a possibility to interact with it (after pressing Enter).
type LineAction interface {
	Engage()
}

type OpenFile interface {
	OpenFile(p string)
}

func LoadFolder(dir string, open OpenFile) Interface {
	fc := &fsContent{
		rootPath: dir,
		root:     os.DirFS(dir),
		open:     open,
	}
	fc.insert(-1, ".")
	return fc
}

type fsContent struct {
	rootPath string
	root     fs.FS
	open     OpenFile

	entries []fs.DirEntry
	lines   []Line
}

func (fc *fsContent) insert(pos int, dirPath string) {
	level := 0
	if pos >= 0 && pos < len(fc.lines) {
		level = fc.lines[pos].(*fsEntryLine).level + 1
	}

	entries, err := fs.ReadDir(fc.root, dirPath)
	if err != nil {
		fc.insertLines(pos+1, []Line{&fsEntryLine{dir: dirPath, err: err, pos: pos + 1, level: level}})
		return
	}

	lines := make([]Line, len(entries))
	for i, entry := range entries {
		lines[i] = &fsEntryLine{
			dir:   dirPath,
			level: level,
			entry: entry,
			pos:   pos + i + 1,
			fc:    fc,
		}
	}
	fc.insertLines(pos+1, lines)
}

func (fc *fsContent) insertLines(pos int, lines []Line) {
	for i := pos; i < len(fc.lines); i++ {
		fc.lines[i].(*fsEntryLine).pos += len(lines)
	}
	fc.lines = slices.Insert(fc.lines, pos, lines...)
}

func (fc *fsContent) collapse(pos int) {
	if pos < 0 || pos >= len(fc.lines)-1 {
		return
	}
	level := fc.lines[pos].(*fsEntryLine).level
	end := pos + 1
	for ; end < len(fc.lines) && fc.lines[end].(*fsEntryLine).level > level; end++ {
	}
	fc.lines = slices.Delete(fc.lines, pos+1, end)

	removedLen := end - pos - 1
	movedPart := fc.lines[pos+1:]
	for i := range movedPart {
		movedPart[i].(*fsEntryLine).pos -= removedLen
	}
}

func (fc *fsContent) Len() int {
	return len(fc.lines)
}

func (fc *fsContent) Lines() []Line {
	return fc.lines
}

type fsEntryLine struct {
	dir   string
	pos   int
	level int
	err   error // error to display
	entry fs.DirEntry

	fc *fsContent

	// mutable state, depends on the Engage action
	display []byte
	expIdx  int
}

const (
	signExpanded = ">~"
)

func (fl *fsEntryLine) String() string {
	if fl.display == nil {
		prefix := strings.Repeat("\t", fl.level) + "  "
		if fl.err != nil {
			fl.display = []byte(prefix + fl.err.Error())
		} else {
			fl.display = []byte(prefix + fl.entry.Name())
		}
	}
	if fl.entry.IsDir() {
		fl.display[fl.level] = signExpanded[fl.expIdx]
	}
	return string(fl.display)
}

func (fl *fsEntryLine) Len() int { return len(fl.display) }

func (fl *fsEntryLine) MimeType() string { return "text/filename" }

func (fl *fsEntryLine) Engage() {
	if fl.err != nil || fl.fc == nil {
		return
	}

	if !fl.entry.IsDir() {
		if open := fl.fc.open; open != nil {
			open.OpenFile(path.Join(fl.dir, fl.entry.Name()))
		}
		return
	}

	fl.expIdx = 1 - fl.expIdx
	switch fl.expIdx {
	case 0:
		fl.fc.collapse(fl.pos)
	case 1:
		fl.fc.insert(fl.pos, path.Join(fl.dir, fl.entry.Name()))
	}
}
