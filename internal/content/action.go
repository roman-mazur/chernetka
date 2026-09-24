package content

// LineAction extends the Line adding a possibility to interact with it (after pressing Enter).
type LineAction interface {
	Engage()
}

// LineActions provides actions for the lines that are not actionable by themselves.
// It lets an action depend on the document as a whole rather than on the line alone:
// for example, only the first line of a diagram file triggers visualizing the diagram.
type LineActions interface {
	// LineAction returns the action associated with the line at the given index or nil.
	LineAction(lineNumber int) LineAction
}

// ActionAt returns the action associated with the line of the document at the given index or nil.
// The line itself is checked first, then the document, and then the additional providers in order.
func ActionAt(doc Document, lineNumber int, providers ...LineActions) LineAction {
	lines := doc.Lines()
	if lineNumber < 0 || lineNumber >= len(lines) {
		return nil
	}
	if action, ok := lines[lineNumber].(LineAction); ok {
		return action
	}
	if p, ok := doc.(LineActions); ok {
		if action := p.LineAction(lineNumber); action != nil {
			return action
		}
	}
	for _, p := range providers {
		if action := p.LineAction(lineNumber); action != nil {
			return action
		}
	}
	return nil
}
