package code

// Diagnostic is a problem found on a line of a buffer.
type Diagnostic struct {
	Line     int
	Severity Severity
	Message  string
}

// Severity is how bad a Diagnostic is.
type Severity int

const (
	SeverityError Severity = iota
	SeverityWarning
)
