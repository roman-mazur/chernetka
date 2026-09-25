package main

import (
	"io"
	"time"

	"rmazur.io/chernetka/internal/vt"
	"rmazur.io/chernetka/internal/vt/escape"
)

// setupTerminal switches the terminal to raw mode and an alternative screen.
func setupTerminal() (terminal vt.Terminal, err error) {
	term, err := vt.SystemTerminal()
	if err != nil {
		return nil, err
	}
	term.Configure(
		escape.EnableAlternativeBuffer,
		func(out io.Writer) (restore func()) {
			return escape.HideCursor(out, false)
		},
	)
	return term, nil
}

// readInput delivers chunks of input read from in until an error happens.
func readInput(in io.Reader) <-chan []byte {
	res := make(chan []byte)
	go func() {
		defer close(res)
		for {
			buf := make([]byte, 256)
			n, err := in.Read(buf)
			if n > 0 {
				res <- buf[:n]
			}
			if err != nil {
				return
			}
		}
	}()
	return res
}

// detectGraphics queries the terminal whether it can display images and waits for the response on input.
func detectGraphics(out io.Writer, input <-chan []byte, timeout time.Duration) bool {
	if _, err := io.WriteString(out, escape.GraphicsQuery); err != nil {
		return false
	}
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	var response []byte
	for {
		select {
		case chunk, ok := <-input:
			if !ok {
				return false
			}
			response = append(response, chunk...)
			if supported, done := escape.ParseGraphicsQueryResponse(response); done {
				return supported
			}
		case <-timer.C:
			return false
		}
	}
}

// isQuitInput checks whether the user asked to quit with q or Ctrl+C.
func isQuitInput(in []byte) bool {
	return len(in) == 1 && (in[0] == 'q' || in[0] == 0x03)
}
