package editor

import (
	"os"
	"unicode/utf8"

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

func (r RelMove) DoOnBuffer(buf *Buffer, _ RenderPrefs) {
	if r.Dx != 0 {
		r.moveDx(buf)
	}
	if r.Dy != 0 {
		buf.cy += r.Dy
	}
}

func (r RelMove) moveDx(b *Buffer) {
	line := b.Content.Lines()[b.cy].String()
	d := r.Dx
	for d > 0 && b.cx < len(line) {
		_, sz := utf8.DecodeRuneInString(line[b.cx:])
		b.cx += sz
		d--
	}
	for d < 0 && b.cx > 0 {
		_, sz := utf8.DecodeLastRuneInString(line[:b.cx])
		b.cx -= sz
		d++
	}
}

type BufferCommandFunc func(b *Buffer, prefs RenderPrefs)

func (f BufferCommandFunc) DoOnBuffer(buf *Buffer, prefs RenderPrefs) { f(buf, prefs) }

var (
	// MoveHome moves the cursor the beginning of the line.
	MoveHome = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) { b.cx = 0 })
	// MoveEnd moves the cursor the end of the line.
	MoveEnd = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) { b.cx = b.Content.Lines()[b.cy].Len() })
)

// Save stores the boffer content in the destination path.
type Save struct {
	DstPath string
}

func (s *Save) DoOnBuffer(buf *Buffer, _ RenderPrefs) {
	if s.DstPath == "" {
		return
	}
	err := content.Save(buf.Content, s.DstPath)
	if err == nil {
		buf.dirty = false
	}
	// TODO: Visualize the error.
}

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
		Path:    of.Path,
		Content: &content.ErrorContent{Error: err},
	})
}

type CommandFunc func(e *Editor)

func (f CommandFunc) DoOnEditor(e *Editor) { f(e) }

var (
	commandQuit = CommandFunc(func(e *Editor) { e.quitRequested = true })
)

type SwitchMode Mode

func (s SwitchMode) DoOnBuffer(buf *Buffer, _ RenderPrefs) { buf.mode = Mode(s) }
