package input

import (
	"bytes"
	"io"
	"slices"
)

// ConsumeClipboardPaste reads the bracketed paste starting b until its end, reading the rest
// of it from in. The rest is the input read after the end of the paste, from b or from in.
func ConsumeClipboardPaste(b []byte, in io.Reader) (content string, rest []byte, detected bool, err error) {
	if !checkClipboardPaste(b, "200~") {
		return
	}
	detected = true
	// The paste is read into one buffer, which is searched for the end marker as it grows.
	// The marker may be split between the reads: the search resumes before the bytes
	// that can start it.
	data := append(make([]byte, 0, pasteReadSize), b[clipboardPasteCmdLen:]...)
	searched := 0
	for {
		if i := bytes.Index(data[searched:], []byte(pasteEnd)); i >= 0 {
			i += searched
			return string(data[:i]), data[i+len(pasteEnd):], true, nil
		}
		if err != nil {
			return "", nil, true, err
		}
		searched = max(searched, len(data)-len(pasteEnd)+1)
		if len(data) == cap(data) {
			data = slices.Grow(data, len(data))
		}
		var n int
		n, err = in.Read(data[len(data):cap(data)])
		data = data[:len(data)+n]
	}
}

const (
	pasteEnd = "\x1b[201~"
	// pasteReadSize is the initial size of the paste buffer, which is doubled for bigger pastes.
	pasteReadSize = 4096
)

const clipboardPasteCmdLen = 6

func checkClipboardPaste(b []byte, marker string) bool {
	return len(b) >= clipboardPasteCmdLen && b[0] == escByte && b[1] == '[' &&
		string(b[2:clipboardPasteCmdLen]) == marker
}
