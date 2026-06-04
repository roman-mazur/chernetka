package editor

// TestHarness wraps an Editor so tests — including those in extension packages —
// can drive it without the terminal event loop. It exposes the internals (the
// active buffer, the command queue, input dispatch) that the production Editor
// keeps private to Run.
type TestHarness struct {
	*Editor
}

// NewTestHarness returns a harness around a fresh editor with an initialized
// command queue, so Post never blocks even without a running event loop.
func NewTestHarness() *TestHarness {
	return &TestHarness{Editor: &Editor{cmdChannel: make(chan Command, 1)}}
}

// ActiveBuffer returns the buffer currently on top of the stack, or nil when no
// buffer is open.
func (h *TestHarness) ActiveBuffer() *Buffer {
	if h.top == nil {
		return nil
	}
	return h.top.b
}

// Commands returns the channel onto which Post delivers commands. Receiving from
// it lets a test observe commands the editor posts asynchronously.
func (h *TestHarness) Commands() <-chan Command { return h.cmdChannel }

// SetMode switches the active buffer to the given mode.
func (h *TestHarness) SetMode(m Mode) {
	if b := h.ActiveBuffer(); b != nil {
		SwitchMode(m).DoOnBuffer(b, h.rPrefs)
	}
}

// MoveCursorToLineEnd places the cursor just past the last character of the
// current line, where a completion request would typically originate.
func (h *TestHarness) MoveCursorToLineEnd() {
	if b := h.ActiveBuffer(); b != nil {
		MoveEnd.DoOnBuffer(b, h.rPrefs)
	}
}

// SendInput feeds raw input bytes through the editor's input handler, exactly as
// the run loop would. It returns whether the editor requested to quit.
func (h *TestHarness) SendInput(b []byte) (quit bool) {
	return h.handleInput(b)
}
