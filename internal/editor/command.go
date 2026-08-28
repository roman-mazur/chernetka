package editor

import (
	"strings"

	"rmazur.io/chernetka/internal/editor/inputs"
)

func commandInput(buf *Buffer, b []byte, prefs *RenderPrefs) (quit bool) {
	var (
		arrow inputs.CursorArrow
		mod   inputs.Modifier
	)
	if inputs.IsArrow(b, &arrow, &mod) {
		// TODO: Handle history on up/down, move on left/right.
		return false
	}

	if len(b) != 1 {
		return false
	}
	switch b[0] {
	case 0x1b: // Esc — cancel.
		buf.cmdline = ""
		buf.mode = ModeNormal
	case '\r': // Enter — execute.
		cmd := buf.cmdline
		buf.cmdline = ""
		buf.mode = ModeNormal
		return runExCommand(buf, cmd, prefs)
	case 0x7f, 0x08: // Backspace.
		if n := len(buf.cmdline); n > 0 {
			buf.cmdline = buf.cmdline[:n-1]
		} else {
			buf.mode = ModeNormal
		}
	default:
		if b[0] >= 0x20 {
			buf.cmdline += string(b[0]) // TODO: Review to properly support utf8.
		}
	}
	return false
}

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
