// Package editor implements the internal state of a text editor.
package editor

import (
	"fmt"
	"path/filepath"
	"sync"
	"sync/atomic"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/debugflags"
	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/vt"
	"rmazur.io/watch/dirwatch"
)

// Mode represents that editor mode (normal vs insert).
type Mode int

const (
	ModeNormal Mode = iota
	ModeInsert
)

func (m Mode) String() string {
	switch m {
	case ModeNormal:
		return "NORMAL"
	case ModeInsert:
		return "INSERT"
	default:
		return fmt.Sprintf("UNKNOWN_%d", int(m))
	}
}

// Editor represents the editor internal state.
type Editor struct {
	// Root is the project directory. Files to open are searched in it.
	// The current working directory is used if it's empty.
	Root string
	// OpenPath opens the file picked with the :e command.
	// The file is opened as a text buffer if it's not set.
	OpenPath func(path string)
	// ShowDiff shows the version control changes of the file at the absolute path.
	// Ctrl+D does nothing if it's not set.
	ShowDiff func(path string)

	logger.LogEmbed
	mouseHandler

	bufs    []*Buffer // stack of open buffers, the active one is the last
	status  StatusBar // shown below the top buffer
	toolBuf *Buffer   // buffer used by tooling, e.g. to visualize search results

	renderRequested bool
	cmdChannel      chan Command
	quitRequested   bool

	doneOnce sync.Once
	done     chan struct{} // closed when the loop is over, see loopDone

	term   vt.Terminal
	rPrefs RenderPrefs

	lastAction bufferAction // the last engaged line action, can be re-run

	x []Extension // extensions

	watcher *dirwatch.Watcher // watches the files of the buffers, started on the first use

	inputSent atomic.Int64 // bytes of the terminal input sent to the loop, for the test harness
}

type bufferAction struct {
	buf    *Buffer
	action content.LineAction
}

var debugInput = debugflags.IsEnabled("loginputs")

func (e *Editor) Extend(ext Extension) {
	if ext == nil {
		panic("nil extension")
	}
	type el interface {
		EmbeddedLogger() *logger.LogEmbed
	}
	if l, ok := ext.(el); ok {
		// Propagate the logger to the extension.
		le := l.EmbeddedLogger()
		*le = logger.Embed(logger.Prefix(e.Logf, ext.ID()+": "))
		le.LogDebug = e.LogDebug
	}
	e.x = append(e.x, ext)
}

// Send submits the command to a commands channel to be executed on the Editor's loop.
// The command is dropped if the loop is over: the extensions may be waiting for
// their goroutines sending the commands to finish while they are closed.
func (e *Editor) Send(cmd Command) {
	select {
	case e.cmdChannel <- cmd:
	case <-e.loopDone():
	}
}

// loopDone returns the channel closed when the editor loop is over.
func (e *Editor) loopDone() chan struct{} {
	e.doneOnce.Do(func() { e.done = make(chan struct{}) })
	return e.done
}

// SendBufferCmd queues a command changing the active buffer. The extensions
// are notified about the change like they are about the user edits.
func (e *Editor) SendBufferCmd(cmd BufferCommand) {
	e.Send(CommandFunc(func(e *Editor) { e.execBufferCmd(cmd) }))
}

func (e *Editor) execBufferCmd(cmd BufferCommand) {
	b := e.Top()
	if b == nil {
		return
	}
	cmd.DoOnBuffer(b, e.rPrefs)
	e.renderRequested = true
	e.handleAfterEdit(b)
}

// root returns the absolute path of the project directory.
func (e *Editor) root() string {
	root := e.Root
	if root == "" {
		root = "."
	}
	if abs, err := filepath.Abs(root); err == nil {
		return abs
	}
	return root
}

func (e *Editor) handleAfterEdit(buf *Buffer) {
	if buf.selecting {
		e.renderRequested = true
	}
	if !buf.resetMutated() {
		// No edits since the last time.
		return
	}
	for _, ext := range e.x {
		ext.AfterEdit(e, buf)
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
