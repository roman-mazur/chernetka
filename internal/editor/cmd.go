package editor

import (
	"os"

	"rmazur.io/x/edit/internal/content"
)

// BufferCommand performs some action on a Buffer.
type BufferCommand interface {
	DoOnBuffer(buf *Buffer, prefs RenderPrefs)
}

// Command performs an action on an Editor.
type Command interface {
	DoOnEditor(e *Editor)
}

// RelMove changes the cursor by Dx and Dy.
type RelMove struct {
	Dx, Dy int
}

func (r RelMove) DoOnBuffer(buf *Buffer, prefs RenderPrefs) {
	if r.Dx != 0 {
		r.moveDx(buf, prefs.TabSize)
	}
	if r.Dy != 0 {
		buf.cy += r.Dy
	}
}

func (r RelMove) moveDx(b *Buffer, tabSize int) {
	line := b.content.Lines()[b.cy].String()
	d := r.Dx
	for d != 0 {
		ti := screenToTextIdx(line, b.cx, tabSize)
		switch {
		case d > 0:
			if ti >= len(line) {
				return
			}
			if line[ti] == '\t' {
				b.cx += tabSize
			} else {
				b.cx++
			}
			d--

		case d < 0:
			if ti == 0 {
				return
			}
			// '\t' is single-byte ASCII (0x09); continuation bytes (0x80–0xBF) in
			// UTF-8 can never equal 0x09, so checking line[ti-1] is safe here.
			if line[ti-1] == '\t' {
				b.cx -= tabSize
			} else {
				b.cx--
			}
			d++
		}
	}
}

// screenToTextIdx returns the byte index of the character displayed at the
// given visual column. Each tab counts as tabSize columns. Returns len(line)
// when screenCol is at or past the visual end of the line.
func screenToTextIdx(line string, screenCol int, tabSize int) int {
	sc := 0
	for i, c := range line {
		w := 1
		if c == '\t' {
			w = tabSize
		}
		sc += w
		if sc > screenCol {
			return i
		}
	}
	return len(line)
}

type BufferCommandFunc func(b *Buffer, prefs RenderPrefs)

func (f BufferCommandFunc) DoOnBuffer(buf *Buffer, prefs RenderPrefs) { f(buf, prefs) }

var (
	// MoveHome moves the cursor the beginning of the line.
	MoveHome = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) { b.cx = 0 })
	// MoveEnd moves the cursor the end of the line.
	MoveEnd = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) { b.cx = b.content.Lines()[b.cy].Len() })
)

// OpenFile opens a new file via Editor.OpenReader.
type OpenFile struct {
	Path string
}

func (of *OpenFile) DoOnEditor(e *Editor) {
	e.layoutRequested = true

	f, err := os.Open(of.Path)
	if err != nil {
		of.handleError(e, err)
		return
	}
	defer f.Close() // TODO: Delegate to the editor.

	if err := e.OpenReader(of.Path, f); err != nil {
		of.handleError(e, err)
	}
}

func (of *OpenFile) handleError(e *Editor, err error) {
	e.push(&Buffer{
		path:    of.Path,
		content: &content.ErrorContent{Error: err},
	})
}

type CommandFunc func(e *Editor)

func (f CommandFunc) DoOnEditor(e *Editor) { f(e) }

var (
	commandQuit = CommandFunc(func(e *Editor) { e.quitRequested = true })
)
