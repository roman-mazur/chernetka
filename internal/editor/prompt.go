package editor

import (
	"fmt"
	"io"

	"rmazur.io/chernetka/internal/content/symbol"
	"rmazur.io/chernetka/internal/editor/input"
)

// The prompts the extensions open in the status bar: a text to type, like the
// new name of a symbol, and a choice of items, like the refactorings offered.
// They are opened on the editor loop, by the commands the extensions send.

// PromptInput shows the command line for buf with the label before the text,
// which the user edits. submit gets the submitted text; it's not called if the
// prompt is canceled with Esc. A prompt that is open is canceled.
func (e *Editor) PromptInput(buf *Buffer, label, text string, submit func(text string)) {
	e.openCmdLine(buf, text, func(c *cmdLine) prompt {
		return &inputPrompt{c: c, label: label, submitFn: submit}
	})
}

type inputPrompt struct {
	c        *cmdLine
	label    string
	submitFn func(text string)
}

func (p *inputPrompt) prefix() string                { return p.label }
func (p *inputPrompt) input(*Editor, input.Key) bool { return false }
func (p *inputPrompt) changed(*Editor)               {}
func (p *inputPrompt) cancel(*Editor)                {}
func (p *inputPrompt) height(int) int                { return 1 }
func (p *inputPrompt) render(out io.Writer, w int)   { renderCmdline(out, w, p.prefix()+p.c.text, "") }
func (p *inputPrompt) submit(*Editor) (quit bool) {
	if p.c.text != "" {
		p.submitFn(p.c.text)
	}
	return false
}

// PromptChoice shows the items to pick one of in the status bar of buf, with the
// label before the text typed to narrow them down. The arrows and Tab select an
// item, Enter picks it: pick gets its index in items. It's not called if the
// prompt is canceled with Esc. A prompt that is open is canceled.
func (e *Editor) PromptChoice(buf *Buffer, label string, items []string, pick func(i int)) {
	e.openCmdLine(buf, "", func(c *cmdLine) prompt {
		return &choicePrompt{c: c, label: label, all: items, pickFn: pick, offset: buf.offset}
	})
}

type choicePrompt struct {
	c      *cmdLine
	label  string
	all    []string // the items to pick from
	pickFn func(i int)
	offset int // the buffer scroll offset to restore when the prompt is closed

	indexes []int // the indexes in all of the items shown
	pickList
}

func (p *choicePrompt) prefix() string { return p.label }

// input selects the items. Other input narrows them down.
func (p *choicePrompt) input(_ *Editor, k input.Key) (handled bool) {
	switch {
	case k == input.Move(input.CursorArrowUp, 0), k == input.Of(input.Backtab):
		p.move(-1)
	case k == input.Move(input.CursorArrowDown, 0), k == input.Of(input.Tab):
		p.move(1)
	case k.Special == input.CursorMove:
		// Other cursor keys are ignored.
	default:
		return false
	}
	return true
}

// changed lists the items matching the typed text.
func (p *choicePrompt) changed(*Editor) {
	var items []string
	p.indexes = p.indexes[:0]
	for i, item := range p.all {
		if _, ok := fuzzyScore(p.c.text, item); ok {
			items = append(items, item)
			p.indexes = append(p.indexes, i)
		}
	}
	p.reset(items, len(items))
}

// submit picks the selected item.
func (p *choicePrompt) submit(*Editor) (quit bool) {
	p.cancel(nil)
	if len(p.indexes) > 0 {
		p.pickFn(p.indexes[p.sel])
	}
	return false
}

// cancel restores the buffer scroll offset: the taller status bar might have scrolled the content.
func (p *choicePrompt) cancel(*Editor) {
	if p.tall {
		p.c.buf.offset = p.offset
	}
}

func (p *choicePrompt) height(w int) int {
	return p.pickList.height(w, symbol.DisplayWidth(p.prefix()+p.c.text, symbol.NoTabExpansion))
}

func (p *choicePrompt) render(out io.Writer, w int) {
	info := "no matches"
	if len(p.items) > 0 {
		info = fmt.Sprintf("%d/%d", p.sel+1, p.total)
	}
	p.pickList.render(out, w, p.prefix()+p.c.text, info)
}

// Prompting reports whether a prompt opened with PromptInput or PromptChoice is
// shown for the buffer, for the extensions not to open another one over it.
func (e *Editor) Prompting(buf *Buffer) bool {
	c := e.status.cmd
	if c == nil || c.buf != buf {
		return false
	}
	return prompting[*inputPrompt](e) || prompting[*choicePrompt](e)
}
