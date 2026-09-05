package inputs

import (
	"bufio"
	"bytes"
	"io"
)

func ConsumeClipboardPaste(b []byte, in io.Reader) (content string, err error) {
	if !checkClipboardPaste(b, "200~") {
		return
	}
	var (
		input     = bufio.NewReader(io.MultiReader(bytes.NewReader(b[clipboardPasteCmdLen:]), in))
		markerBuf [clipboardPasteCmdLen - 1]byte
		data      []byte
		res       bytes.Buffer
	)
	for {
		data, err = input.ReadBytes(Escape)
		if err != nil {
			return
		}
		res.Write(data[:len(data)-1]) // omit Escape
		clear(markerBuf[:])
		if _, err = io.ReadFull(input, markerBuf[:]); err != nil {
			return
		}
		if string(markerBuf[:]) == "[201~" {
			break
		}
		// Not the end-of-paste marker: the Escape and the probed bytes
		// are part of the pasted content, keep them.
		res.WriteByte(Escape)
		res.Write(markerBuf[:])
	}

	content = res.String()
	return
}

const clipboardPasteCmdLen = 6

func checkClipboardPaste(b []byte, marker string) bool {
	return len(b) >= clipboardPasteCmdLen && IsEscape(b[:1]) && b[1] == '[' &&
		string(b[2:clipboardPasteCmdLen]) == marker
}
