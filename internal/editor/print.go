package editor

import (
	"bufio"
	"strings"

	"rmazur.io/x/edit/internal/content"
)

func printClearLine(out *bufio.Writer) {
	out.WriteString("\x1b[2K")
}

func printFmt(out *bufio.Writer, data content.Interface, i, j int, tabSize int) {
	lines := data.Lines()
	_ = lines[i:j] // boundary check

	tab := strings.Repeat(" ", tabSize)
	for x := i; x < j; x++ {
		line := lines[x].String()
		line = strings.ReplaceAll(line, "\t", tab)
		printClearLine(out)
		out.WriteString(line)
		out.WriteString("\r\n")
	}
}
