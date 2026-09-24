package escape

import (
	"bytes"
	"encoding/base64"
	"io"
	"strconv"
)

// GraphicsQuery asks the terminal whether it supports the graphics protocol.
// It's followed by a primary device attributes request that all terminals answer to.
// So, a terminal without graphics support answers only the second request.
const GraphicsQuery = "\x1b_Gi=31,s=1,v=1,a=q,t=d,f=24;AAAA\x1b\\" + "\x1b[c"

// ParseGraphicsQueryResponse checks the terminal input received after sending the GraphicsQuery.
// done reports that the terminal answered both requests of the query.
func ParseGraphicsQueryResponse(in []byte) (supported, done bool) {
	supported = bytes.Contains(in, []byte("\x1b_Gi=31;OK\x1b\\"))
	start := bytes.Index(in, []byte("\x1b[?"))
	done = start >= 0 && bytes.IndexByte(in[start:], 'c') >= 0
	return
}

// ShowPNG transmits the image and displays it at the cursor position.
// Non-zero cols or rows scale the image to occupy this number of cells.
// When only one of them is set, the image keeps its aspect ratio.
func ShowPNG(out io.Writer, png io.Reader, cols, rows int) {
	w := graphChunkWriter{
		Writer: out,
		Cols:   cols,
		Rows:   rows,
	}
	defer w.Close()

	_, _ = io.Copy(&w, png)
}

// DeleteImages removes all the images displayed with ShowPNG.
func DeleteImages(out io.Writer) {
	_, _ = io.WriteString(out, "\x1b_Ga=d,q=2\x1b\\")
}

// graphicsChunkSize is the maximum size of the payload in one escape sequence.
const graphicsChunkSize = 4096

type graphChunkWriter struct {
	io.Writer
	Rows, Cols int

	controlsDone bool
	buf          [graphicsChunkSize / 4 * 3]byte // smaller than the chunk size as we base64 it
	n            int                             // buffer size
}

func (w *graphChunkWriter) Write(p []byte) (n int, err error) {
	offset := 0
	for offset < len(p) {
		// A full buffer is flushed only when more data comes, so that the last chunk is written on Close.
		if w.n == len(w.buf) {
			if err := w.writeChunk(true); err != nil {
				return offset, err
			}
		}
		copySize := copy(w.buf[w.n:], p[offset:])
		offset += copySize
		w.n += copySize
	}
	return offset, nil
}

// writePrefix starts a chunk escape sequence. The first one also carries the image controls,
// the following ones have only the m key.
func (w *graphChunkWriter) writePrefix() error {
	if _, err := w.Writer.Write([]byte("\x1b_G")); err != nil {
		return err
	}
	if !w.controlsDone {
		w.controlsDone = true
		if err := w.writeControls(); err != nil {
			return err
		}
		_, err := w.Writer.Write([]byte(","))
		return err
	}
	return nil
}

func (w *graphChunkWriter) writeControls() error {
	// Transmit and display a PNG (f=100) suppressing terminal responses (q=2).
	_, err := w.Writer.Write([]byte("a=T,f=100,q=2"))
	if err != nil {
		return err
	}
	var (
		buf     [16]byte
		sizeCmd = buf[:0]
	)
	if w.Cols > 0 {
		sizeCmd = append(sizeCmd, []byte(",c=")...)
		sizeCmd = strconv.AppendInt(sizeCmd, int64(w.Cols), 10)
	}
	if w.Rows > 0 {
		sizeCmd = append(sizeCmd, []byte(",r=")...)
		sizeCmd = strconv.AppendInt(sizeCmd, int64(w.Rows), 10)
	}
	_, err = w.Writer.Write(sizeCmd)
	return err
}

func (w *graphChunkWriter) writeChunk(more bool) error {
	if err := w.writePrefix(); err != nil {
		return err
	}
	if more {
		_, err := w.Writer.Write([]byte("m=1;"))
		if err != nil {
			return err
		}
	} else {
		_, err := w.Writer.Write([]byte("m=0;"))
		if err != nil {
			return err
		}
	}
	if err := w.encodeBuffer(); err != nil {
		return err
	}
	w.n = 0 // free the buffer once it's written downstream
	_, err := w.Writer.Write([]byte("\x1b\\"))
	return err
}

func (w *graphChunkWriter) encodeBuffer() (err error) {
	if w.n == 0 {
		return nil
	}

	enc := base64.NewEncoder(base64.StdEncoding, w.Writer)
	defer func() {
		closeError := enc.Close()
		if err == nil {
			err = closeError
		}
	}()

	k := 0
	for k < w.n {
		n, err := enc.Write(w.buf[k:w.n])
		if err != nil {
			return err
		}
		k += n
	}
	return
}

func (w *graphChunkWriter) Close() error {
	return w.writeChunk(false)
}
