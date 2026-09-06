package extsyntaxhl

import (
	"strings"

	"rmazur.io/chernetka/internal/editor"
)

type gitMessage struct {
	subject *rawSpan
	parts   []rawSpan // footers and comments
}

var _ = highlighter(&gitMessage{}) // ensure interface implementation

func (gm *gitMessage) reparse(s *source) {
	gm.subject, gm.parts = nil, nil
	sLen := len(s.lines)
	if sLen == 0 {
		return
	}
	gm.subject = &rawSpan{
		EndCol:    len(s.line(0)),
		TokenType: editor.TtHeading,
	}
	if len(s.lines) != 1 && s.line(1) != "" {
		gm.subject = nil // no distinct subject line
		return
	}
	if sLen <= 2 {
		return
	}
	gm.parseBody(s)
}

func (gm *gitMessage) parseBody(s *source) {
	_ = s.lines[2:] // bounds check

	afterEmptyLine, insideTags := true, false
	for i := 2; i < len(s.lines); i++ {
		switch {
		case s.lines[i] == "":
			afterEmptyLine = true

		case s.lines[i][0] == '#':
			afterEmptyLine = false
			gm.parts = append(gm.parts, rawSpan{
				StartLine: i,
				EndLine:   i,
				EndCol:    len(s.lines[i]),
				TokenType: editor.TtComment,
			})

		case afterEmptyLine || insideTags:
			afterEmptyLine = false
			parts := strings.SplitN(s.lines[i], ":", 2)
			insideTags = len(parts) == 2 && !strings.ContainsAny(parts[0], " \t")
			if insideTags {
				gm.parts = append(gm.parts,
					rawSpan{
						StartLine: i,
						EndLine:   i,
						EndCol:    len(parts[0]) + 1,
						TokenType: editor.TtField,
					},
					rawSpan{
						StartLine: i,
						StartCol:  len(parts[0]) + 2,
						EndLine:   i,
						EndCol:    len(s.lines[i]),
						TokenType: editor.TtStringLiteral,
					},
				)
			}
		}
	}
}

func (gm *gitMessage) spans(_ *source, emit func(rawSpan)) {
	if gm.subject != nil {
		emit(*gm.subject)
	}
	for _, p := range gm.parts {
		emit(p)
	}
}

func (gm *gitMessage) Close() error { return nil }
