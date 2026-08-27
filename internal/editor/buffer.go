package editor

import (
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

	mode       Mode
	dirty      bool   // if buffer content is different from the source file
	noKeyboard bool   // a flag that indicates that the last input was not from the keyboard
	cmdline    string // Text typed after ':' while in ModeCommand

	c      position // cursor position
	offset int      // first visible row (scroll)
	w, h   int      // terminal window dimensions

	sel       []span // selected text
	selecting bool

	xData map[string]BufferExtData // data associated with the extensions

	_mutated   bool   // If Buffer was mutated since the last check. Don't use outside resetMutated and setMutated.
	_textCache string // cached result for Text()
}

type span struct {
	start, end position
}

func (s *span) min() position {
	if s.start.y <= s.end.y {
		return position{min(s.start.x, s.end.x), s.start.y}
	}
	return s.end
}

func (s *span) max() position {
	if s.end.y >= s.start.y {
		return position{max(s.start.x, s.end.x), s.end.y}
	}
	return s.start
}

func (s *span) containsLine(line int) bool {
	return s.min().y <= line && line <= s.max().y
}

func (s *span) lineProjection(idx int, lineLen int) span {
	res := span{s.min(), s.max()}
	if res.start.y < idx {
		res.start = position{0, idx}
	}
	if res.end.y > idx {
		res.end = position{lineLen, idx}
	}
	return res
}

type position struct {
	x, y int // x is a symbol offset in the line, y is a line index in the content.
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
func (b *Buffer) Pos() (cx, cy int) { return b.c.x, b.c.y }

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

func (b *Buffer) clampCursor() {
	lines := b.Content.Lines()

	b.c.y = max(0, min(b.c.y, len(lines)-1))

	var line string
	if len(lines) > 0 {
		line = lines[b.c.y].String()
	}
	maxX := len(line)
	if b.mode == ModeNormal && maxX > 0 {
		_, sz := utf8.DecodeLastRuneInString(line)
		maxX -= sz // Normal mode: cursor sits on a character, not past the last one.
	}
	b.c.x = max(0, min(b.c.x, maxX))
	// Vertical movement may land cx mid-rune; snap back to the rune start.
	for b.c.x > 0 && b.c.x < len(line) && !utf8.RuneStart(line[b.c.x]) {
		b.c.x--
	}

	// Adjust scroll so cursor is visible.
	if !b.noKeyboard {
		if b.c.y < b.offset {
			b.offset = b.c.y
		}
		if b.c.y >= b.offset+b.viewHeight() {
			b.offset = b.c.y - b.viewHeight() + 1
		}
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
func (b *Buffer) Render(out io.Writer, prefs *RenderPrefs) {
	// Main content first.
	printableCount := b.printableLinesCount()
	var cr contentPrinter
	cr.prepare(b, b.offset, b.offset+printableCount, prefs)
	cr.render(out)

	// Empty space.
	for row := printableCount; row < b.viewHeight(); row++ {
		escape.ClearLine(out)
		printLineEnding(out)
	}

	// Status bar (reverse colors).
	restoreColors := escape.ReverseVideo(out)
	defer restoreColors()

	if b.mode == ModeCommand {
		// In command mode, the status bar becomes the command line.
		escape.ClearLine(out)
		_, _ = io.WriteString(out, ":")
		_, _ = io.WriteString(out, b.cmdline)
		return
	}

	modeLabel := b.mode.String()
	dirtyMark := ""
	if b.dirty {
		dirtyMark = " [*]"
	}
	status := fmt.Sprintf(" %s  %s%s", modeLabel, b.Path, dirtyMark)
	pos := fmt.Sprintf("%d:%d ", b.c.y+1, b.c.x+1)
	padding := max(b.w-len(status)-len(pos), 0)
	_, _ = io.WriteString(out, status)
	_, _ = io.WriteString(out, strings.Repeat(" ", padding))
	_, _ = io.WriteString(out, pos)
	restoreColors()
}

// RenderCursorPosition asks the Buffer to instruct the terminal to position the cursor
// to allow input for this buffer. Usually called on the top buffer of the Editor.
func (b *Buffer) RenderCursorPosition(out io.Writer, prefs *RenderPrefs) {
	if b.mode == ModeCommand {
		escape.SetCursorPosition(out, b.h, len(b.cmdline)+2)
		return
	}

	lines := b.Content.Lines()
	var cursorLine string
	if len(lines) > 0 {
		cursorLine = lines[b.c.y].String()
	}

	screenCol := runeToScreenCol(cursorLine, b.c.x, prefs.TabSize)
	screenRow := b.c.y - b.offset + 1
	numDisplayWidth := b.lineNumberPrefixWidth()
	escape.SetCursorPosition(out, screenRow, screenCol+numDisplayWidth+1)
}

// printableLinesCount calculates how many lines of content can be visualized given current offset
func (b *Buffer) printableLinesCount() int {
	return min(b.viewHeight(), len(b.Content.Lines())-b.offset)
}

// lineNumberPrefixWidth returns the length of line numbers presented on the left.
func (b *Buffer) lineNumberPrefixWidth() int {
	if b.hideLineNumbers {
		return 0
	}
	return nlDigitsLen(b.offset+b.printableLinesCount()) + 1
}

func (b *Buffer) selectionsOnLine(line int) (res []span) {
	for i := range b.sel {
		if b.sel[i].containsLine(line) {
			lineLen := b.Content.Lines()[line].Len()
			res = append(res, b.sel[i].lineProjection(line, lineLen))
		}
	}
	return
}

// AcceptSuggestion inserts the active inline suggestion at the cursor and
// advances past it. It is a no-op when there is no suggestion or the buffer is
// read-only.
func (b *Buffer) AcceptSuggestion(sug string) {
	if sug == "" || !b.canEdit() {
		return
	}
	line := b.Content.Lines()[b.c.y].String()
	if b.c.x > len(line) {
		b.c.x = len(line)
	}
	b.Mutate().Update(b.c.y, content.TextLine(line[:b.c.x]+sug+line[b.c.x:]))
	b.c.x += len(sug)
}

func (b *Buffer) canEdit() bool {
	_, ok := b.Content.(content.Mutable)
	return ok
}

func (b *Buffer) handleCursor(arrow inputs.CursorArrow, mod inputs.Modifier, prefs *RenderPrefs) {
	if !b.selecting && mod.HasShift() {
		StartTextSelection.DoOnBuffer(b, *prefs)
	}
	if b.selecting && !mod.HasShift() {
		StopTextSelection.DoOnBuffer(b, *prefs)
	}
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
	if !b.selecting {
		b.sel = nil
	}
}

// CheckContentCoordinates returns true if on-screen coordinates are part of the displayed content.
func (b *Buffer) CheckContentCoordinates(row, col int) bool {
	// Check buffer bounds.
	if row < 0 || row >= b.viewHeight() || col < 0 || col >= b.w {
		return false
	}

	// Check visible content length.
	if row+b.offset >= b.Content.Len() {
		return false
	}
	if col < b.lineNumberPrefixWidth() {
		return false
	}

	return true
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

func (b *Buffer) SelectedText() string {
	if len(b.sel) == 0 {
		return ""
	}
	var out bytes.Buffer
	for _, sel := range b.sel {
		b.selectedText(&out, sel)
	}
	return out.String()
}

func (b *Buffer) selectedText(out *bytes.Buffer, s span) string {
	lines := b.Content.Lines()
	start, stop := s.min(), s.max()
	for i := start.y; i <= stop.y; i++ {
		line := lines[i].String()
		j, k := start.x, stop.x
		if i > start.y {
			j = 0
		}
		if i < stop.y {
			k = len(line)
		}
		if i > start.y {
			out.Write([]byte("\n"))
		}
		out.Write([]byte(line[j:k]))
	}
	return out.String()
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
