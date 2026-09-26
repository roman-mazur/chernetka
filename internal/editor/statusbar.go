package editor

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"rmazur.io/chernetka/internal/vt/escape"
)

// StatusBar shows the state of the buffer below its content. In the command mode,
// it becomes the command line, which includes the file picker for the :e command.
type StatusBar struct {
	buf *Buffer // the buffer to show the state of

	quick        *quickOpen // the :e file picker, if it's shown
	projectFiles []string   // files in the project root, listed when the picker was shown last time
}

// Height returns the number of rows the status bar takes. It's rendered right below
// the buffer content, and its width is the buffer width.
func (s *StatusBar) Height() int {
	if q := s.picker(); q != nil {
		return q.height(s.buf.w)
	}
	return 1
}

// picker returns the file picker if it's shown for the buffer.
func (s *StatusBar) picker() *quickOpen {
	if s.quick != nil && s.quick.buf == s.buf && s.buf.mode == ModeCommand {
		return s.quick
	}
	return nil
}

// RenderCursorPosition places the cursor at the end of the command line.
func (s *StatusBar) RenderCursorPosition(out io.Writer) {
	escape.SetCursorPosition(out, s.buf.h+s.Height(), len(s.buf.cmdline)+2)
}

// Render prints the status bar into the provided output.
func (s *StatusBar) Render(out io.Writer) {
	// Set terminal title.
	if s.buf.Path != "" {
		title := "che: " + filepath.Base(s.buf.Path)
		escape.SetTermTitle(out, title)
	}

	if q := s.picker(); q != nil {
		q.render(out, s.buf.w)
		return
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
