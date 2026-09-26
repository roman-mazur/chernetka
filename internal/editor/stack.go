package editor

import (
	"io"
	"iter"
	"os"
	"path/filepath"
	"slices"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/watch/dirwatch"
)

// OpenReader adds a new buffer to the Editor by reading the full content.
func (e *Editor) OpenReader(path string, in io.Reader) error {
	if e.findAndActivateBuffer(path) {
		return nil
	}

	data, err := content.LoadFullText(in)
	if err != nil {
		return err
	}

	buf := &Buffer{
		Path:    path,
		Content: &data,
	}
	e.push(buf)
	return nil
}

func (e *Editor) prepareExt(b *Buffer) {
	b.ext.xData = make(map[string]BufferExtData, len(e.x))
	for _, ext := range e.x {
		b.ext.extend(ext.ID(), ext.MakeBufferData(b))
	}
}

// OpenFile opens a new file via Editor.OpenReader.
// The buffer content is reloaded when the file is changed outside the editor.
type OpenFile struct {
	Path string
}

func (of *OpenFile) DoOnEditor(e *Editor) {
	e.renderRequested = true
	if e.findAndActivateBuffer(of.Path) {
		return
	}

	f, err := os.Open(of.Path)
	if err != nil {
		of.handleError(e, err)
		return
	}
	defer func() { _ = f.Close() }()

	if err := e.OpenReader(of.Path, f); err != nil {
		of.handleError(e, err)
		return
	}
	buf := e.Top()
	buf.fileText = buf.Text()
	e.watchFile(buf)
}

func (of *OpenFile) handleError(e *Editor, err error) {
	e.push(&Buffer{
		Path:    of.Path,
		Content: &content.ErrorContent{Error: err},
	})
}

// OpenDir adds a new buffer to the Editor by reading the directory content.
func (e *Editor) OpenDir(path string, open content.OpenFile) {
	displayPath := path
	absPath, err := filepath.Abs(path)
	if err == nil {
		displayPath = filepath.Base(absPath)
	}

	buf := &Buffer{
		Path:    displayPath,
		Content: content.LoadFolder(path, open),

		hideLineNumbers: true,
		hideLineActions: true,
	}
	e.push(buf)

	changes := make(chan string)
	go dirwatch.Watch(path, changes)

	go func() {
		for range changes {
			folder := content.LoadFolder(path, open)
			folder.SyncState(buf.Content.(*content.FsContent))
			e.Send(CommandFunc(func(e *Editor) {
				buf.Content = folder
				e.renderRequested = true
			}))
		}
	}()
}

// New creates a new scratch buffer that can be later written to a file.
func (e *Editor) New() { e.push(NewScratchBuffer()) }

// OpenBuffer adds the provided buffer to the stack.
func (e *Editor) OpenBuffer(b *Buffer) {
	if !e.findAndActivateBuffer(b.Path) {
		e.push(b)
	}
}

// Top returns the currently active Buffer.
func (e *Editor) Top() *Buffer {
	if len(e.bufs) == 0 {
		return nil
	}
	return e.bufs[len(e.bufs)-1]
}

func (e *Editor) findAndActivateBuffer(p string) bool {
	i := slices.IndexFunc(e.bufs, func(b *Buffer) bool { return samePath(b.Path, p) })
	if i >= 0 {
		e.selectBuffer(i)
	}
	return i >= 0
}

func (e *Editor) push(buf *Buffer) {
	e.cancelCmdLine()
	e.saveTop()
	e.prepareExt(buf)
	e.bufs = append(e.bufs, buf)
}

func (e *Editor) pop() (empty bool) {
	top := e.Top()
	if top == nil {
		return true
	}
	e.cancelCmdLine()
	_ = top.Close() // TODO: log/handle the error.
	e.bufs = e.bufs[:len(e.bufs)-1]
	return len(e.bufs) == 0
}

// selectBuffer moves the buffer with index i to the top of the stack.
func (e *Editor) selectBuffer(i int) {
	if i == len(e.bufs)-1 {
		return
	}
	e.cancelCmdLine()
	e.saveTop()
	buf := e.bufs[i]
	e.bufs = append(slices.Delete(e.bufs, i, i+1), buf)
}

// buffers iterates over the open buffers from the top of the stack.
func (e *Editor) buffers() iter.Seq[*Buffer] {
	return func(yield func(*Buffer) bool) {
		for _, b := range slices.Backward(e.bufs) {
			if !yield(b) {
				return
			}
		}
	}
}

// samePath checks whether both paths point to the same file, resolving relative paths
// against the working directory.
func samePath(a, b string) bool {
	if a == b {
		return true
	}
	if a == "" || b == "" {
		return false
	}
	absA, errA := filepath.Abs(a)
	absB, errB := filepath.Abs(b)
	return errA == nil && errB == nil && absA == absB
}

// saveTop saves the unsaved changes of the top buffer if it's backed by a file.
func (e *Editor) saveTop() {
	buf := e.Top()
	if buf == nil || !buf.dirty || buf.Path == "" {
		return
	}
	if _, isDir := buf.Content.(*content.FsContent); isDir {
		return // Its path is only a display name.
	}
	e.execBufferCmd(&Save{buf.Path}) // Formatting may change it.
}
