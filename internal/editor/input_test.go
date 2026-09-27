package editor

import (
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/editor/input"
)

// handleInput handles the keys decoded from the terminal input like the editor loop does.
func (e *Editor) handleInput(b []byte) (quit bool) {
	keys, _ := input.Keys(b)
	for _, k := range keys {
		if e.handleKey(k) {
			return true
		}
	}
	return false
}

// sendChunks passes the chunks to sendInput like they are read from the terminal,
// and runs the commands sent to the editor loop.
func sendChunks(t *testing.T, e *Editor, chunks ...string) {
	t.Helper()
	var pending []byte
	for _, chunk := range chunks {
		var err error
		pending, err = e.sendInput(append(pending, chunk...), strings.NewReader(""))
		if err != nil {
			t.Fatal(err)
		}
		for len(e.cmdChannel) > 0 {
			(<-e.cmdChannel).DoOnEditor(e)
		}
	}
	if len(pending) > 0 {
		t.Errorf("pending input %q", pending)
	}
}

func TestEditor_SendInput(t *testing.T) {
	newEditor := func(t *testing.T) *Editor {
		e := &Editor{cmdChannel: make(chan Command, 16)}
		if err := e.OpenReader("", strings.NewReader("one\ntwo\nthree\nfour")); err != nil {
			t.Fatal(err)
		}
		return e
	}

	t.Run("keys read at once", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "\x1b[B\x1b[B", "jl")
		if got := e.Top().c; got != pos(3, 1) {
			t.Errorf("cursor at %s", got)
		}
	})

	t.Run("split sequence", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "\x1b[1;", "2B")
		if b := e.Top(); b.c != pos(1, 0) || len(b.sel) == 0 {
			t.Errorf("cursor at %s, selection %v", b.c, b.sel)
		}
	})

	t.Run("split mouse event", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "j\x1b[<0;3", ";1Mi\x1b[<0;3;1m")
		if e.quitRequested {
			t.Fatal("quit on a split mouse event")
		}
		// The click at the first row moves the cursor back from the second line.
		if b := e.Top(); b.mode != ModeInsert || b.c.Line != 0 {
			t.Errorf("mode %s, cursor at %s after the click", b.mode, b.c)
		}
	})

	t.Run("paste", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "i\x1b[200~x\x1b[201~")
		if got := e.Top().Content.Lines()[0].String(); got != "xone" {
			t.Errorf("line %q", got)
		}
	})

	t.Run("keys after quit", func(t *testing.T) {
		e := newEditor(t)
		sendChunks(t, e, "qj")
		if !e.quitRequested || e.Top().c != pos(0, 0) {
			t.Errorf("quit %t, cursor at %s", e.quitRequested, e.Top().c)
		}
	})
}
