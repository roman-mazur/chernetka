package editor

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"rmazur.io/chernetka/internal/vt/escape"
)

type StatusBar struct {
	buf *Buffer
}

// Render prints the status bar into the provided output.
func (s *StatusBar) Render(out io.Writer) {
	// Set terminal title.
	if s.buf.Path != "" {
		title := "che: " + filepath.Base(s.buf.Path)
		escape.SetTermTitle(out, title)
	}

	restoreColors := escape.ReverseVideo(out)
	defer restoreColors()

	if s.buf.mode == ModeCommand {
		// In command mode, the status bar becomes the command line.
		escape.ClearLine(out)
		_, _ = io.WriteString(out, ":")
		_, _ = io.WriteString(out, s.buf.cmdline)
		return
	}

	modeLabel := s.buf.mode.String()
	dirtyMark := ""
	if s.buf.dirty {
		dirtyMark = " [*]"
	}
	status := fmt.Sprintf(" %s  %s%s", modeLabel, s.buf.Path, dirtyMark)
	pos := fmt.Sprintf("%d:%d ", s.buf.c.Line+1, s.buf.c.Col+1)
	padding := max(s.buf.w-len(status)-len(pos), 0)
	_, _ = io.WriteString(out, status)
	_, _ = io.WriteString(out, strings.Repeat(" ", padding))
	_, _ = io.WriteString(out, pos)
}
