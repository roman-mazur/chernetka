package editor

import (
	"io"
	"strings"
	"unicode"
	"unicode/utf8"

	"rmazur.io/chernetka/internal/editor/input"
	"rmazur.io/chernetka/internal/vt/escape"
)

// cmdLine is typed in the status bar in place of the buffer state: an ex command,
// a file to open, or a search pattern. What the typed text means is up to its prompt.
type cmdLine struct {
	buf    *Buffer // the buffer the command line is opened for
	text   string  // typed after the prompt prefix
	prompt prompt

	resumeInsert bool // the buffer returns to the insert mode when the command line is closed
}

// prompt gives the meaning to the text typed in the command line.
type prompt interface {
	// prefix is shown before the typed text.
	prefix() string
	// input handles the keys specific to the prompt. Other keys edit the text,
	// Enter submits it, and Esc cancels the prompt.
	input(e *Editor, k input.Key) (handled bool)
	// changed follows the edits of the text. It may replace the prompt of the command line.
	changed(e *Editor)
	// submit runs after the command line is closed with Enter.
	submit(e *Editor) (quit bool)
	// cancel reverts what the prompt changed while the text was typed.
	cancel(e *Editor)
	// height returns the number of the status bar rows, w columns wide.
	height(w int) int
	// render prints the status bar rows, w columns wide.
	render(out io.Writer, w int)
}

// openCmdLine shows the command line for buf with the prompt and the initial text.
// The command line that is already open is canceled.
func (e *Editor) openCmdLine(buf *Buffer, text string, newPrompt func(c *cmdLine) prompt) {
	e.cancelCmdLine()
	c := &cmdLine{buf: buf, text: text, resumeInsert: buf.mode == ModeInsert}
	// The extensions don't treat typing a command as editing the buffer.
	buf.mode = ModeNormal
	c.prompt = newPrompt(c)
	e.status.cmd = c
	c.prompt.changed(e)
	e.renderRequested = true
}

// cancelCmdLine closes the command line reverting the changes of its prompt.
func (e *Editor) cancelCmdLine() {
	c := e.status.cmd
	if c == nil {
		return
	}
	e.status.cmd = nil
	c.prompt.cancel(e)
	e.resume(c)
}

// submitCmdLine closes the command line running its prompt.
func (e *Editor) submitCmdLine() (quit bool) {
	c := e.status.cmd
	e.status.cmd = nil
	quit = c.prompt.submit(e)
	e.resume(c)
	return quit
}

// resume returns the buffer of the closed command line to the insert mode if it was in it.
func (e *Editor) resume(c *cmdLine) {
	if c.resumeInsert && c.buf.canEdit() {
		c.buf.mode = ModeInsert
	}
	e.renderRequested = true
}

// cmdLineInput edits the text of the command line or passes the key to its prompt.
func (e *Editor) cmdLineInput(k input.Key) (quit bool) {
	c := e.status.cmd
	if c.prompt.input(e, k) {
		return false
	}
	switch k.Special {
	case input.Esc:
		e.cancelCmdLine()
	case input.Enter:
		return e.submitCmdLine()
	case input.Backspace:
		if c.text == "" {
			e.cancelCmdLine()
			return false
		}
		_, sz := utf8.DecodeLastRuneInString(c.text)
		c.text = c.text[:len(c.text)-sz]
		c.prompt.changed(e)
	case input.Text:
		if k.Mod == 0 && unicode.IsPrint(k.Rune) {
			c.text += string(k.Rune)
			c.prompt.changed(e)
		}
	}
	return false
}

// prompting reports whether the command line is open with a prompt of type P.
func prompting[P prompt](e *Editor) bool {
	if e.status.cmd == nil {
		return false
	}
	_, ok := e.status.cmd.prompt.(P)
	return ok
}

// renderCmdline prints the command line with the info aligned to the right.
func renderCmdline(out io.Writer, w int, cmd, info string) {
	restoreColors := escape.ReverseVideo(out)
	defer restoreColors.Undo()

	escape.ClearLine(out)
	_, _ = io.WriteString(out, cmd)
	if info != "" {
		padding := max(w-utf8.RuneCountInString(cmd)-utf8.RuneCountInString(info)-1, 1)
		_, _ = io.WriteString(out, strings.Repeat(" ", padding))
		_, _ = io.WriteString(out, info+" ")
	}
}
