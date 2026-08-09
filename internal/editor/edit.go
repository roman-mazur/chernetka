// Package editor implements the internal state of a text editor.
package editor

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"iter"
	"os"
	"path/filepath"
	"sync"
	"time"

	"golang.org/x/term"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/escape"
	"rmazur.io/chernetka/internal/editor/inputs"
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

const debugInput = true

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

// InOut exposes Reader and Writer for the Editor.
// Reader is used to receive user input. Writer is used to show the terminal UI.
type InOut struct {
	io.Reader
	io.Writer

	WindowChangeSignal <-chan struct{}
}

// Editor represents the editor internal state.
type Editor struct {
	mouseHandler

	top *bufEntry // stack of open buffers

	renderRequested bool
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
	if e.top == nil {
		return nil
	}
	return e.top.b
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

	_ = e.top.b.Close() // TODO: log/handle the error.

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

func (e *Editor) Run(t *InOut, logf logger.Func) {
	start := time.Now()
	defer func() {
		logf("session done %s", time.Since(start))
		for _, ext := range e.x {
			if c, ok := ext.(io.Closer); ok {
				_ = c.Close()
			}
		}
	}()

	termCleanup := e.initTerminal(t, logf)
	defer termCleanup()

	e.rPrefs = newRenderPrefs()

	if e.cmdChannel == nil {
		e.cmdChannel = make(chan Command)
	}

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	go e.readAndHandleInput(ctx, bufio.NewReader(t), logf)
	go e.handleWindowChange(ctx, t.WindowChangeSignal)

	out := bufio.NewWriter(t)

	var lastRenderTime time.Time
	const renderDelay = 10 * time.Millisecond

	e.renderRequested = true
	for {
		// Close the current buffer if necessary.
		if e.quitRequested {
			if e.pop() {
				break // All done.
			}
			e.quitRequested = false
			e.renderRequested = true
		}

		// Render.
		if e.renderRequested {
			for buf := range e.layout() {
				buf.clampCursor()
				buf.Render(out, &e.rPrefs)
				buf.noKeyboard = false
			}
			e.renderRequested = false
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

func (e *Editor) initTerminal(t *InOut, logf logger.Func) (cleanup func()) {
	f, ok := t.Writer.(*os.File)
	if !ok {
		logf("not a terminal")
		return func() {}
	}

	var cleanupOps []func()
	cleanup = func() {
		for i := len(cleanupOps) - 1; i >= 0; i-- {
			cleanupOps[i]()
		}
	}

	cleanupOps = append(cleanupOps, escape.EnableAlternativeBuffer(f))
	cleanupOps = append(cleanupOps, escape.DisableLineWrapping(f))

	e.termFd = int(f.Fd())
	state, err := term.MakeRaw(e.termFd)
	if err != nil {
		logf("cannot initialize terminal (fd %d): %s", e.termFd, err)
		return
	}
	cleanupOps = append(cleanupOps, func() {
		_ = term.Restore(e.termFd, state)
	})

	cleanupOps = append(cleanupOps, escape.EnableMouse(f))
	return
}

// handleWindowChange reads signals typically wired to SIGWINCH and propagates a layout request.
// It ensures these signals are propagated at most every 10ms.
func (e *Editor) handleWindowChange(ctx context.Context, s <-chan struct{}) {
	const maxFrequency = 10 * time.Millisecond
	var (
		lastTime  time.Time
		timerChan <-chan time.Time
		timer     time.Timer
	)
	for {
		select {
		case _, ok := <-s:
			if !ok {
				return
			}
			if time.Since(lastTime) > maxFrequency {
				lastTime = time.Now()
				e.Post(commandRequestLayout)
			} else if timerChan == nil {
				timer.Reset(maxFrequency)
				timerChan = timer.C
			}

		case <-timerChan:
			e.Post(commandRequestLayout)
			timerChan = nil

		case <-ctx.Done():
			return
		}
	}
}

func (e *Editor) readAndHandleInput(ctx context.Context, in *bufio.Reader, logf logger.Func) {
	const bufSize = 64
	bPool := sync.Pool{New: func() any { return make([]byte, bufSize) }}
	var inBuf [bufSize]byte

	for ctx.Err() == nil {
		// Get input.
		n, err := in.Read(inBuf[:])
		if err != nil {
			e.cmdChannel <- commandQuit
			break
		}
		if n == 0 {
			continue
		}
		buf := bPool.Get().([]byte)
		copy(buf, inBuf[:n])
		input := buf[:n]
		if debugInput {
			logf("input: %v", input)
		}

		// Check for mouse input.
		_, k, err := e.readAndHandleMouse(input, logf)
		if err != nil {
			e.cmdChannel <- commandQuit
			break
		}
		input = input[k:]
		if len(input) == 0 {
			bPool.Put(buf)
			continue
		}

		// Handle input.
		e.cmdChannel <- CommandFunc(func(e *Editor) {
			quit := e.handleInput(input)
			if quit {
				commandQuit(e)
			} else {
				e.renderRequested = true
			}
			bPool.Put(buf)
		})
	}
}

func (e *Editor) readAndHandleMouse(input []byte, logf logger.Func) (handled bool, n int, err error) {
	var data inputs.Mouse
	data, n, err = inputs.ReadMouse(input)

	if err != nil {
		if errors.Is(err, inputs.ErrorNotMouse) {
			err = nil
		}
		return
	}

	handled = true
	e.cmdChannel <- CommandFunc(func(e *Editor) { e.handleMouse(data, logf) })
	return
}

func (e *Editor) handleMouse(data inputs.Mouse, logf logger.Func) {
	event := e.transformInput(data)
	if debugInput {
		logf("mouse: %s", event)
	}

	buf := e.Top()
	if buf == nil {
		return
	}
	buf.noKeyboard = true

	switch event.eventType {
	case mouseEventTypeScroll:
		dir := event.Mod.SrollDirection(event.Mouse)
		Scroll(dir).DoOnBuffer(buf, e.rPrefs)
		logf("offset: %d, dir: %d", buf.offset, dir)
		e.renderRequested = true
		return

	case mouseEventTypeRaw:
		if event.Button != inputs.MouseButtonLeft || !event.Pressed {
			return
		}
		buf.c.x = data.X - 1
		buf.c.y = buf.offset + data.Y - 1
		e.renderRequested = true
		return
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
		quit = normalInput(buf, input, &e.rPrefs)
		e.handleAfterEdit(buf)
		return

	case ModeInsert:
		handled := e.extHandleInsert(buf, input)
		if !handled {
			insertInput(buf, input, &e.rPrefs)
		}
		e.handleAfterEdit(buf)
		return false

	case ModeCommand:
		return commandInput(buf, input, &e.rPrefs)
	default:
		return false
	}
}

func (e *Editor) extHandleInsert(buf *Buffer, b []byte) (handled bool) {
	for _, ext := range e.x {
		handled = ext.HandleInsertInput(buf, &e.rPrefs, b)
		if handled {
			return
		}
	}
	return
}

func (e *Editor) handleAfterEdit(buf *Buffer) {
	if !buf.resetMutated() {
		// No edits since the last time.
		return
	}
	for _, ext := range e.x {
		ext.AfterEdit(e, buf)
	}
}

func (e *Editor) RequestLayout() {
	e.renderRequested = true
}

// layout selects the next buffer to render configuring its dimensions.
func (e *Editor) layout() iter.Seq[*Buffer] {
	state := layoutState{e}
	return state.Pass()
}

type layoutState struct {
	editor *Editor
}

func (lps *layoutState) resolveWindowSize() (w int, h int) {
	if lps.editor.termFd == 0 {
		return 80, 40 // real terminal is not resolved - test/mock environment
	}
	w, h, _ = term.GetSize(lps.editor.termFd)
	return
}

func (lps *layoutState) Pass() iter.Seq[*Buffer] {
	done := false
	w, h := lps.resolveWindowSize()
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
