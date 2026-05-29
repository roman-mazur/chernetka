package editor

import (
	"bufio"
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"rmazur.io/x/edit/internal/content"
	"rmazur.io/x/edit/internal/editor/escape"
)

func printFmt(out *bufio.Writer, data content.Interface, i, j int, tabSize int, cy int) {
	lines := data.Lines()
	_ = lines[i:j] // boundary check

	tab := strings.Repeat(" ", tabSize)
	nl := nlDigitsLen(j)
	nlFmt := "%" + strconv.Itoa(nl) + "d "

	for x := i; x < j; x++ {
		line := lines[x].String()
		line = strings.ReplaceAll(line, "\t", tab)
		escape.ClearLine(out)
		nlColor := color.Gray{Y: 100}
		if x == cy {
			nlColor.Y = 200
		}
		escape.ColorText(out, fmt.Sprintf(nlFmt, x+1), nlColor)
		out.WriteString(line)
		out.WriteString("\r\n")
	}
}

func nlDigitsLen(x int) int {
	l := 0
	for x > 0 {
		x /= 10
		l++
	}
	return l
}
