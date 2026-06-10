// Package editor implements the internal state of a text editor.
package editor

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/term"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/escape"
	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/watch/dirwatch"
)

// Mode represents that editor mode (normal vs insert).
type Mode int

const (
	ModeNormal Mode = iota
	ModeInsert
	ModeCommand
)

const debugInput = false

func (m Mode) String() string {
	switch m {
	case ModeNormal:
		return "NORMAL"
	case ModeInsert:
		return "INSERT"
	case ModeCommand:
		return "COMMAND"
	default:
		return fmt.Sprintf("UNKNOWN_%d", int(m))
	}
}

// Content represents the data visualized by the editor - the text lines.
type Content []string

// Editor represents the editor internal state.
type Editor struct {
	top *bufEntry // stack of open buffers

	layoutRequested bool
	cmdChannel      chan Command
	quitRequested   bool

	termFd int
	rPrefs RenderPrefs

	x []Extension // extensions
}

func (e *Editor) Extend(ext Extension) {
	e.x = append(e.x, ext)
}

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
	b.xData = make(map[string]BufferExtData, len(e.x))
	for _, ext := range e.x {
		b.xData[ext.ID()] = ext.MakeBufferData(b)
	}
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
	}
	e.push(buf)

	changes := make(chan string)
	go dirwatch.Watch(path, changes)

	go func() {
		for range changes {
			folder := content.LoadFolder(path, open)
			folder.SyncState(buf.Content.(*content.FsContent))
			e.Post(CommandFunc(func(e *Editor) {
				buf.Content = folder
				e.layoutRequested = true
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

func (e *Editor) findAndActivateBuffer(p string) bool {
	for entry := range e.buffers() {
		if entry.matches(p) {
			e.selectBuffer(entry)
			return true
		}
	}
	return false
}

func (e *Editor) Post(cmd Command) {
	e.cmdChannel <- cmd
}

func (e *Editor) push(buf *Buffer) {
	e.prepareExt(buf)

	entry := &bufEntry{
		b:    buf,
		next: e.top,
	}
	if e.top != nil {
		e.top.prev = entry
	}
	e.top = entry
}

func (e *Editor) pop() (empty bool) {
	if e.top == nil {
		return true
	}

	e.top = e.top.next
	if e.top == nil {
		return true
	}
	e.top.prev = nil
	return false
}

// selectBuffer moves the selected buffer to the top of the stack.
func (e *Editor) selectBuffer(entry *bufEntry) {
	if entry == e.top {
		return
	}
	if e.top.prev != nil {
		panic("top has prev")
	}

	if entry.next != nil {
		entry.next.prev = entry.prev
	}
	if entry.prev != nil {
		entry.prev.next = entry.next
	}

	entry.prev = nil
	entry.next = e.top

	e.top.prev = entry
	e.top = entry
}

func (e *Editor) buffers() iter.Seq[*bufEntry] {
	cur := e.top
	return func(yield func(*bufEntry) bool) {
		for cur != nil {
			if !yield(cur) {
				return
			}
			cur = cur.next
		}
	}
}

type bufEntry struct {
	b          *Buffer
	next, prev *bufEntry
}

func (be *bufEntry) matches(p string) bool {
	return be.b.Path == p
}

func (e *Editor) Run(f *os.File, logf logger.Func) {
	start := time.Now()
	defer func() {
		logf("session done %s", time.Since(start))
		for _, ext := range e.x {
			if c, ok := ext.(io.Closer); ok {
				_ = c.Close()
			}
		}
	}()

	restoreAltBuffer := escape.EnableAlternativeBuffer(f)
	defer restoreAltBuffer()
	restoreLineWrapping := escape.DisableLineWrapping(f)
	defer restoreLineWrapping()

	e.rPrefs = newRenderPrefs()

	e.termFd = int(f.Fd())
	state, err := term.MakeRaw(e.termFd)
	if err != nil {
		logf("cannot initialize terminal (fd %d): %s", e.termFd, err)
		return
	}
	defer term.Restore(e.termFd, state)

	if e.cmdChannel == nil {
		e.cmdChannel = make(chan Command)
	}

	var inputsStop atomic.Bool
	defer inputsStop.Store(true)
	go e.readAndHandleInput(f, logf, &inputsStop)

	out := bufio.NewWriter(os.Stdout)

	var lastRenderTime time.Time
	const renderDelay = 10 * time.Millisecond

	e.layoutRequested = true
	for {
		// Close the current buffer if necessary.
		if e.quitRequested {
			if e.pop() {
				break // All done.
			}
			e.quitRequested = false
			e.layoutRequested = true
		}

		// Render.
		if e.layoutRequested {
			for buf := range e.layout() {
				buf.clampCursor(&e.rPrefs)
				buf.render(out, &e.rPrefs)
			}
			e.layoutRequested = false
		}
		lastRenderTime = time.Now()

		// Handle commands, including inputs.
	loop:
		for {
			select {
			case cmd := <-e.cmdChannel:
				cmd.DoOnEditor(e)
			default:
				if passed := time.Since(lastRenderTime); passed < renderDelay {
					time.Sleep(renderDelay - passed)
				}
				break loop
			}
		}
	}
}

func (e *Editor) readAndHandleInput(f *os.File, logf logger.Func, stop *atomic.Bool) {
	bPool := sync.Pool{New: func() any { return make([]byte, 64) }}
	var inBuf [64]byte
	for !stop.Load() {
		// Get input.
		n, err := f.Read(inBuf[:])
		if err != nil {
			e.cmdChannel <- commandQuit
			break
		}
		if n == 0 {
			continue
		}

		input := bPool.Get().([]byte)
		copy(input, inBuf[:n])
		if debugInput {
			logf("input: %v", input[:n])
		}
		// Handle input.
		e.cmdChannel <- CommandFunc(func(e *Editor) {
			quit := e.handleInput(input[:n])
			if quit {
				commandQuit(e)
			} else {
				e.layoutRequested = true
			}
			bPool.Put(input)
		})
	}
}

func (e *Editor) handleInput(input []byte) (quit bool) {
	buf := e.top.b

	// Ctrl+S saves the current buffer in any mode.
	if len(input) == 1 && input[0] == 0x13 {
		(&Save{buf.Path}).DoOnBuffer(buf, e.rPrefs)
		return false
	}

	switch buf.mode {
	case ModeNormal:
		return normalInput(buf, input, &e.rPrefs)
	case ModeInsert:
		handled, changed := e.extHandleInsert(buf, input)
		if !handled {
			changed = insertInput(buf, input, &e.rPrefs)
		}
		if changed {
			e.extAfterEdit(buf)
		}
		return false
	case ModeCommand:
		return commandInput(buf, input, &e.rPrefs)
	default:
		return false
	}
}

func (e *Editor) extHandleInsert(buf *Buffer, b []byte) (handled, changed bool) {
	for _, ext := range e.x {
		handled, changed = ext.HandleInsertInput(buf, &e.rPrefs, b)
		if handled {
			return
		}
	}
	return
}

func (e *Editor) extAfterEdit(buf *Buffer) {
	for _, ext := range e.x {
		ext.AfterEdit(e, buf)
	}
}

func (e *Editor) RequestLayout() {
	e.layoutRequested = true
}

// layout selects the next buffer to render configuring its dimensions.
func (e *Editor) layout() iter.Seq[*Buffer] {
	state := layoutState{e}
	return state.Pass()
}

type layoutState struct {
	editor *Editor
}

func (lps *layoutState) Pass() iter.Seq[*Buffer] {
	done := false
	w, h, _ := term.GetSize(lps.editor.termFd)
	return func(yield func(*Buffer) bool) {
		if done {
			return
		}
		// TODO: consider rendering multiple buffers.
		buf := lps.editor.top.b
		buf.w, buf.h = w, h

		done = true
		yield(buf)
	}
}

type RenderPrefs struct {
	TabSize int
}

func newRenderPrefs() RenderPrefs {
	return RenderPrefs{TabSize: 4}
}

func (rp *RenderPrefs) tabsScaleUp()   { rp.TabSize = min(rp.TabSize*2, 8) }
func (rp *RenderPrefs) tabsScaleDown() { rp.TabSize = max(rp.TabSize/2, 1) }
