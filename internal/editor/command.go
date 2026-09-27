package editor

import (
	"io"
	"strings"

	"rmazur.io/chernetka/internal/editor/input"
)

// exPrompt runs the typed command on Enter, see runExCommand.
// Typing ":e " turns it into the file picker.
type exPrompt struct {
	c *cmdLine
}

func newExPrompt(c *cmdLine) prompt { return &exPrompt{c: c} }

func (p *exPrompt) prefix() string { return ":" }

func (p *exPrompt) input(*Editor, input.Key) bool { return false }

func (p *exPrompt) changed(e *Editor) {
	if strings.HasPrefix(p.c.text, quickOpenPrefix) {
		p.c.prompt = newQuickOpen(p.c)
		p.c.prompt.changed(e)
	}
}

func (p *exPrompt) submit(e *Editor) (quit bool) { return runExCommand(p.c.buf, p.c.text, &e.rPrefs) }

func (p *exPrompt) cancel(*Editor) {}

func (p *exPrompt) height(int) int { return 1 }

func (p *exPrompt) render(out io.Writer, w int) { renderCmdline(out, w, p.prefix()+p.c.text, "") }

func runExCommand(buf *Buffer, cmd string, prefs *RenderPrefs) (quit bool) {
	for len(cmd) > 0 {
		key := cmd[0:1]
		cmd = cmd[1:]

		switch key {
		case "q":
			return true

		case "w":
			dstPath := buf.Path
			if len(cmd) > 2 {
				cmd, dstPath, _ = strings.Cut(cmd, " ")
			}
			(&Save{DstPath: dstPath}).DoOnBuffer(buf, *prefs)

		case "p":
			switch cmd {
			case "bcopy":
				ClipboardCopy.DoOnBuffer(buf, *prefs)
			case "bcut":
				ClipboardCut.DoOnBuffer(buf, *prefs)
			case "bpaste":
				ClipboardPaste.DoOnBuffer(buf, *prefs)
			}
		}
	}
	return false
}
