// Package extrefactor binds the keys of the review mode to the code navigation and
// refactoring tools the other extensions provide for their buffers, like the language
// servers: the usages and the callers of a symbol listed in the tool buffer, renaming,
// and the code actions (the fixes and the refactorings offered).
//
// The tools are found in the extension data of the buffer implementing the interfaces
// of this package: an extension implements them without the editor knowing about them.
package extrefactor

import (
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/change"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/input"
)

// UsageFinder can be optionally implemented by editor.BufferExtData to find where the
// symbol at a position of its buffer is used.
type UsageFinder interface {
	editor.BufferExtData

	// FindUsages looks for the usages of the symbol at pos in the background.
	// They are shown in the tool buffer with an editor.ShowTool command sent to the loop.
	FindUsages(loop editor.Sender, pos content.Position)
}

// CallerFinder can be optionally implemented by editor.BufferExtData to find the
// functions calling the one at a position of its buffer.
type CallerFinder interface {
	editor.BufferExtData

	// FindCallers looks for the callers of the function at pos in the background.
	// They are shown in the tool buffer with an editor.ShowTool command sent to the loop.
	FindCallers(loop editor.Sender, pos content.Position)
}

// Renamer can be optionally implemented by editor.BufferExtData to rename the symbol
// at a position of its buffer.
type Renamer interface {
	editor.BufferExtData

	// Rename finds out what can be renamed at pos in the background, asks for
	// the new name with Editor.PromptInput, and applies the renaming to the
	// buffers and the files.
	Rename(loop editor.Sender, pos content.Position)
}

// CodeActionProvider can be optionally implemented by editor.BufferExtData to offer
// the fixes and the refactorings of a part of its buffer.
type CodeActionProvider interface {
	editor.BufferExtData

	// ShowCodeActions finds the actions for the span in the background, which may
	// be empty at the cursor, lets one be picked with Editor.PromptChoice, and
	// applies it to the buffers and the files.
	ShowCodeActions(loop editor.Sender, span content.Span)
}

// Integration is the editor extension binding the keys, which are the ones of the
// JetBrains IDEs. It has no buffer data.
//
//	Alt+F7       list the usages of the symbol under the cursor in the tool buffer
//	Ctrl+Alt+H   list the callers of the function under the cursor in the tool buffer
//	Shift+F6     rename the symbol under the cursor: the new name is typed in the command line
//	Alt+Enter    pick a fix or a refactoring of the selection, or of the cursor position
//
// Ctrl+Alt+H and Alt+Enter need a terminal with the kitty keyboard protocol: the
// other terminals send Esc for Alt.
type Integration struct{}

// The keys.
var (
	KeyUsages      = input.Fn(7, input.ModAlt)
	KeyCallers     = input.Key{Rune: 'h', Mod: input.ModCtrl | input.ModAlt}
	KeyRename      = input.Fn(6, input.ModShift)
	KeyCodeActions = input.Key{Special: input.Enter, Mod: input.ModAlt}
)

func (*Integration) ID() string { return "refactor" }

func (*Integration) MakeBufferData(editor.Sender, *editor.Buffer) editor.BufferExtData { return nil }

func (*Integration) AfterEdit(editor.Sender, *editor.Buffer, []change.Item) {}

// HandleKey implements editor.KeyHandler.
func (*Integration) HandleKey(loop editor.Sender, buf *editor.Buffer, k input.Key) (handled bool) {
	cx, cy := buf.Pos()
	pos := content.Position{Line: cy, Col: cx}
	// The key is handled only if the buffer has the tool.
	switch k {
	case KeyUsages:
		f, ok := editor.FindExtData[UsageFinder](buf)
		if ok {
			f.FindUsages(loop, pos)
		}
		return ok
	case KeyCallers:
		f, ok := editor.FindExtData[CallerFinder](buf)
		if ok {
			f.FindCallers(loop, pos)
		}
		return ok
	case KeyRename:
		r, ok := editor.FindExtData[Renamer](buf)
		if ok {
			r.Rename(loop, pos)
		}
		return ok
	case KeyCodeActions:
		p, ok := editor.FindExtData[CodeActionProvider](buf)
		if !ok {
			return false
		}
		span := content.Span{Start: pos, End: pos}
		if sel := buf.Selection(); len(sel) > 0 {
			span = content.Span{Start: sel[0].Start, End: sel[len(sel)-1].End}
		}
		p.ShowCodeActions(loop, span)
		return true
	}
	return false
}

var _ editor.KeyHandler = new(Integration)
