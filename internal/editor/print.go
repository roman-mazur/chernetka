package editor

import (
	"bufio"
	"strings"

	"rmazur.io/x/edit/internal/content"
)

func printContent(out *bufio.Writer, data content.Interface, i, j int, prefs *renderPrefs) {
	lines := data.Lines()
	_ = lines[i:j] // boundary check

	tab := strings.Repeat(" ", prefs.tabSize)
	for x := i; x < j; x++ {
		line := lines[x].String()
		line = strings.ReplaceAll(line, "\t", tab)
		out.WriteString("\x1b[2K") // Clear the line.
		out.WriteString(line)
		out.WriteString("\r\n")
	}
}
