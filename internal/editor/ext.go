package editor

import (
	"fmt"
	"io"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/editor/input"
)

// Extension represents an editor extension.
type Extension interface {
	ID() string
	MakeBufferData(buf *Buffer) BufferExtData

	// AfterEdit runs on the editor loop after the buffer is changed. The results
	// of the work started in the background are to be sent back to the loop.
	AfterEdit(loop Sender, buf *Buffer)
}

// Sender queues the commands to run on the editor loop, where the editor and
// its buffers can be changed. *Editor implements it.
type Sender interface {
	// Send queues the command.
	Send(cmd Command)
	// SendBufferCmd queues the command changing the active buffer, the extensions are
	// notified about the changes.
	SendBufferCmd(cmd BufferCommand)
}

// InsertKeyHandler can be optionally implemented by an Extension to handle the keys
// typed in the insert mode before the editor does.
type InsertKeyHandler interface {
	HandleInsertKey(buf *Buffer, k input.Key) (handled bool)
}

// BufferExtData represents data associated with a Buffer and managed by an Extension.
type BufferExtData any

// CodeAssist can be optionally implemented by BufferExtData.
type CodeAssist interface {
	BufferExtData

	TextSuggestion() code.Suggestion
}

// Formatter can be optionally implemented by BufferExtData to format its
// buffer before it's saved.
type Formatter interface {
	BufferExtData

	// Format edits the buffer content into its canonical form, with tabs as
	// wide as prefs tell. It runs on the editor loop, so it must not take
	// long; the buffer may be left as is.
	Format(prefs RenderPrefs)
}

// DefinitionFinder can be optionally implemented by BufferExtData to find where
// the symbol at a position of its buffer is defined.
type DefinitionFinder interface {
	BufferExtData

	// FindDefinition looks for the definition of the symbol at pos in the
	// background. When it's found, a GoTo command is sent to the loop.
	FindDefinition(loop Sender, pos content.Position)
}

// SyntaxHighlighter can be optionally implemented by BufferExtData.
type SyntaxHighlighter interface {
	BufferExtData

	// SyntaxSpans returns the highlighted regions of the given line.
	// The result is sorted by Start, spans never overlap, and every span stays
	// within the bounds of line. Regions that need no highlight are omitted, so
	// an empty result means the whole line is rendered with the default color.
	SyntaxSpans(lineNumber int, line string) []SyntaxSpan
}

// SyntaxSpan is a half-open [Start, End) byte range of a single line that
// should be rendered with the color of its TokenType.
type SyntaxSpan struct {
	LineNumber int
	Start, End int
	code.TokenType
}

func (ss SyntaxSpan) String() string {
	return fmt.Sprintf("%d:%d:%d:%s", ss.LineNumber, ss.Start, ss.End, ss.TokenType)
}

type bufExtensions struct {
	xData           []extData // data associated with the extensions, in the order they were registered
	actionProviders []content.LineActions
}

type extData struct {
	id   string
	data BufferExtData
}

func (be *bufExtensions) data(id string) BufferExtData {
	for _, x := range be.xData {
		if x.id == id {
			return x.data
		}
	}
	return nil
}

func (be *bufExtensions) extend(id string, data BufferExtData) {
	if data == nil {
		return
	}
	be.xData = append(be.xData, extData{id: id, data: data})
	if p, ok := data.(content.LineActions); ok {
		be.actionProviders = append(be.actionProviders, p)
	}
}

func (be *bufExtensions) close(allErrors *[]error) {
	for _, x := range be.xData {
		if closer, ok := x.data.(io.Closer); ok {
			*allErrors = append(*allErrors, closer.Close())
		}
	}
}

// FindExtData returns the data associated with the buffer by the first registered extension
// that implements T. It lets extensions use each other's data without knowing their IDs.
func FindExtData[T any](b *Buffer) (res T, ok bool) {
	for _, x := range b.ext.xData {
		if res, ok = x.data.(T); ok {
			return res, true
		}
	}
	return res, false
}
