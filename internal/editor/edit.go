// Package editor implements the internal state of a text editor.
package editor

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"os"
	"path/filepath"
	"slices"
	"time"

	"golang.org/x/term"
	"rmazur.io/x/edit/internal/content"
	"rmazur.io/x/edit/internal/logger"
)

// Mode represents that editor mode (normal vs insert).
type Mode int

const (
	ModeNormal Mode = iota
	ModeInsert
)

const debugInput = false

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

// Content represents the data visualized by the editor - the text lines.
type Content []string

// Editor represents the editor internal state.
type Editor struct {
	b  []*Buffer // open buffers
	bi int       // current buffer

	termFd int
	rPrefs renderPrefs
}

// OpenReader adds a new buffer to the Editor by reading the full content.
func (e *Editor) OpenReader(path string, in io.Reader) error {
	data, err := content.LoadFullText(in)
	if err != nil {
		return err
	}

	e.push(&Buffer{
		path:    path,
		content: &data,
	})
	return nil
}

// OpenDir adds a new buffer to the Editor by reading the directory content.
func (e *Editor) OpenDir(path string) {
	displayPath := path
	absPath, err := filepath.Abs(path)
	if err == nil {
		displayPath = filepath.Base(absPath)
	}
	e.push(&Buffer{
		path:    displayPath,
		content: content.LoadFolder(path),
	})
}

// New creates a new scratch buffer that can be later written to a file.
func (e *Editor) New() { e.push(NewScratchBuffer()) }

func (e *Editor) push(buf *Buffer) { e.b = slices.Insert(e.b, e.bi, buf) }

func (e *Editor) Run(f *os.File, logf logger.Func) {
	start := time.Now()
	defer func() {
		logf("session done %s", time.Since(start))
	}()

	const (
		escAltBufferOn  = "\x1b[?1049h"
		escAltBufferOff = "\x1b[?1049l"
	)
	_, _ = fmt.Fprint(f, escAltBufferOn)
	defer fmt.Fprint(f, escAltBufferOff)

	e.rPrefs = newRenderPrefs()

	e.termFd = int(f.Fd())
	state, err := term.MakeRaw(e.termFd)
	if err != nil {
		logf("cannot initialize terminal (fd %d): %s", e.termFd, err)
		return
	}
	defer term.Restore(e.termFd, state)

	var inBuf [64]byte
	out := bufio.NewWriter(os.Stdout)
	quit := false

	for !quit {
		// Render.
		for buf := range e.layout() {
			buf.clampCursor()
			buf.render(out, &e.rPrefs)
		}

		// Get input.
		n, err := f.Read(inBuf[:])
		if err != nil {
			break
		}
		if n == 0 {
			continue
		}
		input := inBuf[:n]
		if debugInput {
			logf("input: %v", input)
		}

		// Handle input.
		quit = e.handleInput(e.b[e.bi].mode, input)
	}
}

func (e *Editor) handleInput(mode Mode, input []byte) (quit bool) {
	buf := e.b[e.bi]

	if isArrow(input) {
		buf.handleCursor(input, &e.rPrefs)
		return false
	}

	switch mode {
	case ModeNormal:
		switch string(input) {
		case "+", "=":
			e.rPrefs.tabsScaleUp()
			return false
		case "-":
			e.rPrefs.tabsScaleDown()
			return false
		}
		return normalInput(buf, input, &e.rPrefs)
	case ModeInsert:
		insertInput(buf, input)
		return false
	default:
		return false
	}
}

// isArrow reports whether b is an ANSI arrow escape sequence.
func isArrow(b []byte) bool {
	return len(b) == 3 && b[0] == 0x1b && b[1] == '[' &&
		(b[2] == 'A' || b[2] == 'B' || b[2] == 'C' || b[2] == 'D')
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
		buf := lps.editor.b[lps.editor.bi]
		buf.w, buf.h = w, h

		done = true
		yield(buf)
	}
}

type renderPrefs struct {
	tabSize int
}

func newRenderPrefs() renderPrefs {
	return renderPrefs{tabSize: 4}
}

func (rp *renderPrefs) tabsScaleUp()   { rp.tabSize = min(rp.tabSize*2, 8) }
func (rp *renderPrefs) tabsScaleDown() { rp.tabSize = max(rp.tabSize/2, 1) }
