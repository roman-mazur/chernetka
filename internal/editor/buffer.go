package editor

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/escape"
	"rmazur.io/chernetka/internal/editor/inputs"
)

type Buffer struct {
	Path    string
	Content content.Interface

	hideLineNumbers bool
	noCurrentLineHL bool

	mode    Mode
	dirty   bool   // if buffer content is different from the source file
	cmdline string // Text typed after ':' while in ModeCommand

	cx, cy int // cursor position: cy is the line index, cx is the byte offset within that line
	offset int // first visible row (scroll)
	w, h   int // terminal window dimensions

	xData map[string]BufferExtData // data associated with the extensions

	_mutated   bool   // If Buffer was mutated since the last check. Don't use outside resetMutated and setMutated.
	_textCache string // cached result for Text()
}

// NewScratchBuffer constructs a new Buffer with empty content.
func NewScratchBuffer() *Buffer {
	return &Buffer{
		Content: content.Empty(),
	}
}

// ExtensionData returns an object managed by the Extension with the provided id.
// Can be nil if such object of Extension does not exist.
func (b *Buffer) ExtensionData(id string) BufferExtData {
	if b.xData == nil {
		return nil
	}
	return b.xData[id]
}

// Pos provides a pair of coordinates pointing to the current cursor location.
// cx is the symbol offset within the content line.
// cy is the line index in the Content.
func (b *Buffer) Pos() (cx, cy int) { return b.cx, b.cy }

// Close propagates the call to the Content and extension objects if they implement io.Closer.
func (b *Buffer) Close() error {
	allErrors := make([]error, 0, len(b.xData)+1)
	if closer, ok := b.Content.(io.Closer); ok {
		allErrors = append(allErrors, closer.Close())
	}
	for _, x := range b.xData {
		if closer, ok := x.(io.Closer); ok {
			allErrors = append(allErrors, closer.Close())
		}
	}
	return errors.Join(allErrors...)
}

func (b *Buffer) clampCursor(_ *RenderPrefs) {
	lines := b.Content.Lines()

	b.cy = max(0, min(b.cy, len(lines)-1))

	var line string
	if len(lines) > 0 {
		line = lines[b.cy].String()
	}
	maxX := len(line)
	if b.mode == ModeNormal && maxX > 0 {
		_, sz := utf8.DecodeLastRuneInString(line)
		maxX -= sz // Normal mode: cursor sits on a character, not past the last one.
	}
	b.cx = max(0, min(b.cx, maxX))
	// Vertical movement may land cx mid-rune; snap back to the rune start.
	for b.cx > 0 && b.cx < len(line) && !utf8.RuneStart(line[b.cx]) {
		b.cx--
	}

	// Adjust scroll so cursor is visible.
	if b.cy < b.offset {
		b.offset = b.cy
	}
	if b.cy >= b.offset+b.viewHeight() {
		b.offset = b.cy - b.viewHeight() + 1
	}
}

// runeToScreenCol returns the visual column at the given byte index, expanding
// tabs by tabSize. Used once per render to project cx onto the terminal.
func runeToScreenCol(line string, runeIdx, tabSize int) int {
	if runeIdx > len(line) {
		runeIdx = len(line)
	}
	col := 0
	for i, r := range line {
		if i >= runeIdx {
			break
		}
		if r == '\t' {
			col += tabSize
		} else {
			col++
		}
	}
	return col
}

// viewHeight is the number of text rows (terminal height minus the status bar).
func (b *Buffer) viewHeight() int { return b.h - 1 }

// Render visualizes the buffer UI content writing to the provided output.
// This includes presenting the visible part of the content, the status line and the command line at the bottom.
func (b *Buffer) Render(out *bufio.Writer, prefs *RenderPrefs) {
	defer out.Flush()

	resetSyncOutput := escape.SyncOutput(out)
	defer resetSyncOutput()

	showCursor := escape.HideCursor(out, b.mode != ModeNormal)
	defer showCursor()
	escape.MoveTopLeft(out)

	lines := b.Content.Lines()

	printableCount := min(b.viewHeight(), len(lines)-b.offset)
	var cr contentPrinter
	cr.prepare(b, b.offset, b.offset+printableCount, prefs)
	cr.render(out)

	for row := printableCount; row < b.viewHeight(); row++ {
		escape.ClearLine(out)
		out.WriteString("\r\n")
	}

	var cursorLine string
	if len(lines) > 0 {
		cursorLine = lines[b.cy].String()
	}
	screenCol := runeToScreenCol(cursorLine, b.cx, prefs.TabSize)

	// Status bar (reverse colors).
	restoreColors := escape.ReverseVideo(out)
	defer restoreColors()

	if b.mode == ModeCommand {
		// In command mode, the status bar becomes the command line.
		escape.ClearLine(out)
		out.WriteString(":")
		out.WriteString(b.cmdline)
		escape.SetCursorPosition(out, b.h, len(b.cmdline)+2)
		return
	}

	modeLabel := b.mode.String()
	dirtyMark := ""
	if b.dirty {
		dirtyMark = " [*]"
	}
	status := fmt.Sprintf(" %s  %s%s", modeLabel, b.Path, dirtyMark)
	pos := fmt.Sprintf("%d:%d ", b.cy+1, screenCol+1)
	padding := max(b.w-len(status)-len(pos), 0)
	out.WriteString(status)
	out.WriteString(strings.Repeat(" ", padding))
	out.WriteString(pos)
	restoreColors()

	// Reposition and show cursor.
	screenRow := b.cy - b.offset + 1
	var numDisplayWidth int
	if !b.hideLineNumbers {
		numDisplayWidth = nlDigitsLen(b.offset+printableCount) + 1
	}
	escape.SetCursorPosition(out, screenRow, screenCol+numDisplayWidth+1)
}

// AcceptSuggestion inserts the active inline suggestion at the cursor and
// advances past it. It is a no-op when there is no suggestion or the buffer is
// read-only.
func (b *Buffer) AcceptSuggestion(sug string) {
	if sug == "" || !b.canEdit() {
		return
	}
	line := b.Content.Lines()[b.cy].String()
	if b.cx > len(line) {
		b.cx = len(line)
	}
	b.Mutate().Update(b.cy, content.TextLine(line[:b.cx]+sug+line[b.cx:]))
	b.cx += len(sug)
}

func (b *Buffer) canEdit() bool {
	_, ok := b.Content.(content.Mutable)
	return ok
}

func (b *Buffer) handleCursor(arrow inputs.CursorArrow, prefs *RenderPrefs) {
	switch arrow {
	case inputs.CursorArrowUp:
		RelMove{Dy: -1}.DoOnBuffer(b, *prefs)
	case inputs.CursorArrowDown:
		RelMove{Dy: 1}.DoOnBuffer(b, *prefs)
	case inputs.CursorArrowRight:
		RelMove{Dx: 1}.DoOnBuffer(b, *prefs)
	case inputs.CursorArrowLeft:
		RelMove{Dx: -1}.DoOnBuffer(b, *prefs)
	}
}

// Text returns a string representation of the full buffer content.
// It can be used, for example, to send the buffer content to an LSP server.
func (b *Buffer) Text() string {
	if b._textCache != "" {
		return b._textCache
	}

	var buf bytes.Buffer
	_ = content.SaveToWriter(b.Content, &buf)
	b._textCache = buf.String()
	return b._textCache
}

// Mutate is used to start changing the buffer content.
// If underlying Content is not mutable, the returned implementation is a noop.
func (b *Buffer) Mutate() content.Mutable {
	m, _ := b.Content.(content.Mutable)
	return bufMutator{
		Mutable: m,
		Buffer:  b,
	}
}

func (b *Buffer) setMutated() {
	b.dirty = true
	b._mutated = true
	b._textCache = ""
}

func (b *Buffer) resetMutated() (prev bool) {
	prev = b._mutated
	b._mutated = false
	return
}

type bufMutator struct {
	content.Mutable
	*Buffer
}

func (bm bufMutator) Insert(pos int, line content.Line) {
	if bm.Mutable == nil {
		return
	}
	bm.Mutable.Insert(pos, line)
	bm.setMutated()
}
func (bm bufMutator) Update(pos int, line content.Line) {
	if bm.Mutable == nil {
		return
	}
	bm.Mutable.Update(pos, line)
	bm.setMutated()
}
func (bm bufMutator) Delete(pos int) {
	if bm.Mutable == nil {
		return
	}
	bm.Mutable.Delete(pos)
	bm.setMutated()
}
