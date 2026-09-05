package extsyntaxhl

import (
	"strings"

	"rmazur.io/chernetka/internal/editor"
)

type gitMessage struct {
	subject *rawSpan
	footers [][2]rawSpan // footers as key/value pairs
}

var _ = highlighter(&gitMessage{}) // ensure interface implementation

func (gm *gitMessage) reparse(s *source) {
	gm.subject, gm.footers = nil, nil
	sLen := len(s.lines)
	if sLen == 0 {
		return
	}
	if s.line(1) != "" {
		return
	}
	gm.subject = &rawSpan{
		EndCol:    len(s.line(0)),
		TokenType: editor.TtHeading,
	}
	if sLen <= 2 {
		return
	}
	footersStart := findGitMsgFootersStart(s)
	if footersStart == 0 {
		return
	}
	for i := footersStart; i < sLen; i++ {
		parts := strings.Split(s.lines[i], ":")
		if len(parts) < 2 && strings.ContainsAny(parts[0], " \t") {
			continue
		}
		gm.footers = append(gm.footers, [2]rawSpan{
			{
				StartLine: i,
				EndLine:   i,
				EndCol:    len(parts[0]) + 1,
				TokenType: editor.TtComment,
			},
			{
				StartLine: i,
				StartCol:  len(parts[0]) + 2,
				EndLine:   i,
				EndCol:    len(s.lines[i]),
				TokenType: editor.TtStringLiteral,
			},
		})
	}
}

func findGitMsgFootersStart(s *source) int {
	_ = s.lines[2:] // bounds check

	footersStart := 0
	afterEmptyLine := false
	for i := 2; i < len(s.lines); i++ {
		switch {
		case s.lines[i] == "":
			afterEmptyLine = true
		case afterEmptyLine:
			afterEmptyLine = false
			parts := strings.Split(s.lines[i], ":")
			if len(parts) > 1 && !strings.ContainsAny(parts[0], " \t") {
				footersStart = i
				break
			}
		}
	}
	return footersStart
}

func (gm *gitMessage) spans(_ *source, emit func(rawSpan)) {
	if gm.subject != nil {
		emit(*gm.subject)
	}
	for _, footer := range gm.footers {
		emit(footer[0])
		emit(footer[1])
	}
}

func (gm *gitMessage) Close() error { return nil }
