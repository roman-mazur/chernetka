package editor

// Extension represents an editor extension.
type Extension interface {
	ID() string
	MakeBufferData(buf *Buffer) BufferExtData

	AfterEdit(e *Editor, buf *Buffer)
	HandleInsertInput(buf *Buffer, prefs *RenderPrefs, b []byte) (handled, changed bool)
}

// BufferExtData represents data associated with a Buffer and managed by an Extension.
type BufferExtData interface {
	TextSuggestion() string
}
