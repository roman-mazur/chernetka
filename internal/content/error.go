package content

// ErrorContent visualizes an error.
type ErrorContent struct {
	Error error
}

func (ec *ErrorContent) Lines() []Line {
	return []Line{
		TextLine("ERROR"),
		TextLine("\t" + ec.Error.Error()),
	}
}

func (ec *ErrorContent) Len() int { return 2 }
