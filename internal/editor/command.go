package editor

import "strings"

func commandInput(buf *Buffer, b []byte, prefs *RenderPrefs) (quit bool) {
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
	if cmd == "q" {
		return true
	}

	if strings.HasPrefix(cmd, "w") {
		dstPath := buf.path
		if len(cmd) > 2 {
			_, dstPath, _ = strings.Cut(cmd, " ")
		}
		(&Save{DstPath: dstPath}).DoOnBuffer(buf, *prefs)
	}
	return false
}
