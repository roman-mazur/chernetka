package input

import (
	"bufio"
	"bytes"
	"io"
)

// ConsumeClipboardPaste reads the bracketed paste starting b until its end, reading the rest
// of it from in. The rest is the input read after the end of the paste, from b or from in.
func ConsumeClipboardPaste(b []byte, in io.Reader) (content string, rest []byte, detected bool, err error) {
	if !checkClipboardPaste(b, "200~") {
		return
	}
	detected = true
	var (
		input     = bufio.NewReader(io.MultiReader(bytes.NewReader(b[clipboardPasteCmdLen:]), in))
		markerBuf [clipboardPasteCmdLen - 1]byte
		data      []byte
		res       bytes.Buffer
	)
	for {
		data, err = input.ReadBytes(escByte)
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
		res.WriteByte(escByte)
		res.Write(markerBuf[:])
	}

	content = res.String()
	rest, _ = input.Peek(input.Buffered())
	return
}

const clipboardPasteCmdLen = 6

func checkClipboardPaste(b []byte, marker string) bool {
	return len(b) >= clipboardPasteCmdLen && b[0] == escByte && b[1] == '[' &&
		string(b[2:clipboardPasteCmdLen]) == marker
}
