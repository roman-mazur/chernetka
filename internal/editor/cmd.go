package editor

// Command performs some action on a Buffer.
type Command interface {
	Do(buf *Buffer, prefs RenderPrefs)
}

// RelMove changes the cursor by Dx and Dy.
type RelMove struct {
	Dx, Dy int
}

func (r RelMove) Do(buf *Buffer, prefs RenderPrefs) {
	if r.Dx != 0 {
		r.moveDx(buf, prefs.TabSize)
	}
	if r.Dy != 0 {
		buf.cy += r.Dy
	}
}

func (r RelMove) moveDx(b *Buffer, tabSize int) {
	line := b.content.Lines()[b.cy].String()
	d := r.Dx
	for d != 0 {
		ti := screenToTextIdx(line, b.cx, tabSize)
		switch {
		case d > 0:
			if ti >= len(line) {
				return
			}
			if line[ti] == '\t' {
				b.cx += tabSize
			} else {
				b.cx++
			}
			d--

		case d < 0:
			if ti == 0 {
				return
			}
			// '\t' is single-byte ASCII (0x09); continuation bytes (0x80–0xBF) in
			// UTF-8 can never equal 0x09, so checking line[ti-1] is safe here.
			if line[ti-1] == '\t' {
				b.cx -= tabSize
			} else {
				b.cx--
			}
			d++
		}
	}
}

type CommandFunc func(b *Buffer, prefs RenderPrefs)

func (f CommandFunc) Do(buf *Buffer, prefs RenderPrefs) { f(buf, prefs) }

var (
	// MoveHome moves the cursor the beginning of the line.
	MoveHome = CommandFunc(func(b *Buffer, _ RenderPrefs) { b.cx = 0 })
	// MoveEnd moves the cursor the end of the line.
	MoveEnd = CommandFunc(func(b *Buffer, _ RenderPrefs) { b.cx = b.content.Lines()[b.cy].Len() })
)
