package extlsp

import (
	"strings"
	"unicode/utf16"

	"go.lsp.dev/protocol"
	"rmazur.io/chernetka/internal/content"
)

// The editor and LSP locate text the same way, except for columns: the editor
// counts the bytes of a line (content.Position.Col), while LSP counts UTF-16
// code units (protocol.Position.Character). Lines are 0-based in both.

// spanOf converts an LSP range in the document made of lines to a span. It
// reports false if the range is outside of the document.
func spanOf(lines []content.Line, r protocol.Range) (content.Span, bool) {
	start, ok1 := positionOf(lines, r.Start)
	end, ok2 := positionOf(lines, r.End)
	return content.Span{Start: start, End: end}, ok1 && ok2
}

func positionOf(lines []content.Line, p protocol.Position) (content.Position, bool) {
	if int(p.Line) >= len(lines) {
		return content.Position{}, false
	}
	return content.Position{Line: int(p.Line), Col: byteOffset(lines[p.Line].String(), p.Character)}, true
}

// textPosition converts an LSP position in text to an editor one. A line past
// the end of text is kept as is, for the editor to clamp.
func textPosition(text string, p protocol.Position) content.Position {
	lines := strings.SplitN(text, "\n", int(p.Line)+2)
	if int(p.Line) >= len(lines) {
		return content.Position{Line: int(p.Line)}
	}
	return content.Position{Line: int(p.Line), Col: byteOffset(lines[p.Line], p.Character)}
}

// lspPosition converts an editor position to an LSP one. line is the text of
// the line the position is on.
func lspPosition(line string, p content.Position) protocol.Position {
	return protocol.Position{Line: uint32(p.Line), Character: utf16Len(line[:min(p.Col, len(line))])}
}

func utf16Len(s string) (n uint32) {
	for _, r := range s {
		n += uint32(utf16.RuneLen(r))
	}
	return n
}

// byteOffset converts a UTF-16 based LSP character offset in line to a byte offset.
func byteOffset(line string, utf16Offset uint32) int {
	var n uint32
	for i, r := range line {
		if n >= utf16Offset {
			return i
		}
		n += uint32(utf16.RuneLen(r))
	}
	return len(line)
}
