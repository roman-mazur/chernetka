package editor

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"rmazur.io/chernetka/internal/vt/escape"
)

// StatusBar shows the state of the buffer below its content. The command line
// is shown in its place while a command, a file to open, or a search pattern is typed.
type StatusBar struct {
	buf *Buffer  // the buffer to show the state of
	cmd *cmdLine // the command line, if it's open

	projectFiles []string // files in the project root, listed when the picker was shown last time
}

// Height returns the number of rows the status bar takes. It's rendered right below
// the buffer content, and its width is the buffer width.
func (s *StatusBar) Height() int {
	if c := s.cmdLine(); c != nil {
		return c.prompt.height(s.buf.w)
	}
	return 1
}

// cmdLine returns the command line if it's open for the buffer.
func (s *StatusBar) cmdLine() *cmdLine {
	if s.cmd != nil && s.cmd.buf == s.buf {
		return s.cmd
	}
	return nil
}

// RenderCursorPosition places the cursor at the end of the command line.
func (s *StatusBar) RenderCursorPosition(out io.Writer) {
	var col int
	if c := s.cmdLine(); c != nil {
		col = utf8.RuneCountInString(c.prompt.prefix() + c.text)
	}
	escape.SetCursorPosition(out, s.buf.h+s.Height(), col+1)
}

// Render prints the status bar into the provided output.
func (s *StatusBar) Render(out io.Writer) {
	// Set terminal title.
	if s.buf.Path != "" {
		title := "che: " + filepath.Base(s.buf.Path)
		escape.SetTermTitle(out, title)
	}

	if c := s.cmdLine(); c != nil {
		c.prompt.render(out, s.buf.w)
		return
	}

	restoreColors := escape.ReverseVideo(out)
	defer restoreColors()

	modeLabel := s.buf.mode.String()
	dirtyMark := ""
	if s.buf.dirty {
		dirtyMark = " [*]"
	}
	status := fmt.Sprintf(" %s  %s%s", modeLabel, s.buf.Path, dirtyMark)
	if re := s.buf.search; re != nil {
		status += "  /" + re.String()
	}
	pos := fmt.Sprintf("%d:%d ", s.buf.c.Line+1, s.buf.c.Col+1)
	padding := max(s.buf.w-len(status)-len(pos), 0)
	_, _ = io.WriteString(out, status)
	_, _ = io.WriteString(out, strings.Repeat(" ", padding))
	_, _ = io.WriteString(out, pos)
}
