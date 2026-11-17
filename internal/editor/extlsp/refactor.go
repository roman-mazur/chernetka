package extlsp

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"

	"go.lsp.dev/protocol"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/extrefactor"
	"rmazur.io/chernetka/internal/editor/navigation"
	"rmazur.io/chernetka/internal/lsp"
)

// The refactoring tools: renaming, and the code actions (the fixes and the
// refactorings the server offers).
//
// The server asks first what is renamed at the position (prepareRename): the
// identifier, but gopls also makes it the import path of the package to move
// when the position is in the package clause, and the signature of the function
// to change when it's on its func keyword. What the server says is edited in
// the command line, and the server computes the edits of the renaming.
//
// A code action comes with the edits, or with data to get them from the server
// when the action is picked, or with a command the server runs itself, sending
// the edits back with workspace/applyEdit. The edits are applied to the open
// buffers, and to the other files on the disk; the files are also created,
// renamed, and deleted, as the extraction to a new file or a package move needs.

// Rename asks what is renamed at pos, prompts for the new name, and applies the renaming.
func (lbd *BufferData) Rename(loop editor.Sender, pos content.Position) {
	lines := lbd.buf.Content.Lines()
	if pos.Line >= len(lines) {
		return
	}
	at := lspPosition(lines[pos.Line].String(), pos)
	text := lbd.buf.Text()
	buf, srv := lbd.buf, lbd.srv
	lbd.request(loop, func(ctx context.Context, client lspClient) {
		target, err := client.PrepareRename(ctx, lbd.docUri, at, text)
		if err != nil {
			srv.showMessage(loop, "rename: "+err.Error())
			return
		}
		if target == nil {
			srv.showMessage(loop, "nothing to rename at the cursor")
			return
		}
		loop.Send(editor.CommandFunc(func(e *editor.Editor) {
			if e.Top() != buf || e.Prompting(buf) {
				return
			}
			e.PromptInput(buf, "rename: ", target.Placeholder, func(newName string) {
				if newName == target.Placeholder {
					return
				}
				lbd.request(e, func(ctx context.Context, client lspClient) {
					edit, err := client.Rename(ctx, lbd.docUri, at, newName)
					if err != nil {
						srv.showMessage(loop, "rename: "+err.Error())
						return
					}
					if edit == nil || len(edit.Changes) == 0 {
						srv.showMessage(loop, "rename: nothing to change")
						return
					}
					srv.applyEditOnLoop(loop, *edit, "rename "+target.Placeholder+" to "+newName)
				})
			})
		}))
	})
}

// ShowCodeActions asks the server for the fixes and the refactorings of the span,
// lets one be picked in the command line, and applies it.
func (lbd *BufferData) ShowCodeActions(loop editor.Sender, span content.Span) {
	lines := lbd.buf.Content.Lines()
	if span.Max().Line >= len(lines) {
		return
	}
	rng := protocol.Range{
		Start: lspPosition(lines[span.Min().Line].String(), span.Min()),
		End:   lspPosition(lines[span.Max().Line].String(), span.Max()),
	}
	// The fixes are matched to the problems: the ones of the lines of the span.
	diagRange := protocol.Range{
		Start: protocol.Position{Line: rng.Start.Line},
		End:   protocol.Position{Line: rng.End.Line + 1},
	}
	buf, srv := lbd.buf, lbd.srv
	lbd.request(loop, func(ctx context.Context, client lspClient) {
		actions, err := client.CodeActions(ctx, lbd.docUri, rng, srv.publishedDiagnostics(lbd.docUri, &diagRange))
		if err != nil {
			srv.showMessage(loop, "code actions: "+err.Error())
			return
		}
		if len(actions) == 0 {
			srv.showMessage(loop, "no fixes or refactorings here")
			return
		}
		titles := make([]string, len(actions))
		for i, a := range actions {
			titles[i] = a.Title
		}
		loop.Send(editor.CommandFunc(func(e *editor.Editor) {
			if e.Top() != buf || e.Prompting(buf) {
				return
			}
			e.PromptChoice(buf, "action: ", titles, func(i int) { lbd.runAction(e, actions[i]) })
		}))
	})
}

// runAction applies the code action: its edits, resolved if needed, or the
// command the server runs for it.
func (lbd *BufferData) runAction(loop editor.Sender, action lsp.CodeAction) {
	srv := lbd.srv
	lbd.request(loop, func(ctx context.Context, client lspClient) {
		if action.Edit == nil && action.Command == nil && len(action.Data) > 0 {
			resolved, err := client.ResolveCodeAction(ctx, action)
			if err != nil {
				srv.showMessage(loop, action.Title+": "+err.Error())
				return
			}
			action = resolved
		}
		if action.Edit != nil {
			srv.applyEditOnLoop(loop, *action.Edit, action.Title)
		}
		if action.Command != nil {
			// The edits come back with workspace/applyEdit, see applyEditRequest.
			if err := client.ExecuteCommand(ctx, *action.Command); err != nil {
				srv.showMessage(loop, action.Title+": "+err.Error())
			}
		}
		if action.Edit == nil && action.Command == nil {
			srv.showMessage(loop, action.Title+": the server has nothing to apply")
		}
	})
}

// applyEditRequest receives the edits the server makes on its own, like the ones
// of a command. It's called on the connection goroutine: the edits are applied
// on the editor loop later, and the server is told they are applied.
func (srv *server) applyEditRequest(req lsp.ApplyEditRequest) error {
	if srv.loop == nil {
		return errors.New("no editor")
	}
	srv.applyEditOnLoop(srv.loop, req.Edit, req.Label)
	return nil
}

// applyEditOnLoop queues the edit to be applied on the editor loop. It doesn't wait.
func (srv *server) applyEditOnLoop(loop editor.Sender, edit lsp.WorkspaceEdit, label string) {
	go loop.Send(editor.CommandFunc(func(e *editor.Editor) {
		if err := srv.applyEdit(e, edit); err != nil {
			srv.showMessage(loop, label+": "+err.Error())
		}
	}))
}

// applyEdit applies the edit on the editor loop: the open buffers are edited,
// and the other files on the disk, which are also created, renamed, and
// deleted. The server is told about the files changed on the disk.
func (srv *server) applyEdit(e *editor.Editor, edit lsp.WorkspaceEdit) error {
	var (
		errs    []error
		changed []protocol.FileEvent
		pending = make(map[*editor.Buffer][]protocol.TextEdit) // the edits of the buffers, applied together
		order   []*editor.Buffer
	)
	// flush applies the pending edits of the buffer as one change, undone at once.
	flush := func(buf *editor.Buffer) {
		edits := pending[buf]
		if len(edits) == 0 {
			return
		}
		delete(pending, buf)
		e.EditBuffer(buf, func(buf *editor.Buffer) {
			buf.MutateAsGroup(func(content.Mutable) { applyEdits(buf, edits) })
		})
	}
	for _, c := range edit.Changes {
		path := c.URI.Filename()
		buf := e.FindBuffer(path)
		var err error
		switch c.Kind {
		case lsp.ChangeEdit:
			if buf != nil {
				if _, ok := pending[buf]; !ok {
					order = append(order, buf)
				}
				pending[buf] = append(pending[buf], c.Edits...)
			} else if err = editFile(path, c.Edits); err == nil {
				changed = append(changed, protocol.FileEvent{URI: c.URI, Type: protocol.FileChangeTypeChanged})
			}

		case lsp.ChangeCreate:
			if _, statErr := os.Stat(path); statErr == nil && !c.Overwrite {
				break // Keep what's there.
			}
			if err = os.MkdirAll(filepath.Dir(path), 0o755); err == nil {
				err = os.WriteFile(path, nil, 0o644)
			}
			if err == nil {
				changed = append(changed, protocol.FileEvent{URI: c.URI, Type: protocol.FileChangeTypeCreated})
			}

		case lsp.ChangeRename:
			newPath := c.NewURI.Filename()
			if _, statErr := os.Stat(newPath); statErr == nil && !c.Overwrite {
				err = fmt.Errorf("%s exists", newPath)
				break
			}
			if buf != nil {
				flush(buf)
				e.SaveBuffer(buf)
			}
			if err = os.MkdirAll(filepath.Dir(newPath), 0o755); err == nil {
				err = os.Rename(path, newPath)
			}
			if err != nil {
				break
			}
			changed = append(changed,
				protocol.FileEvent{URI: c.URI, Type: protocol.FileChangeTypeDeleted},
				protocol.FileEvent{URI: c.NewURI, Type: protocol.FileChangeTypeCreated})
			if buf != nil {
				// The buffer follows the file: it's reopened at the new path.
				cx, cy := buf.Pos()
				(&editor.GoTo{Location: navigation.Location{Path: newPath, Position: content.Position{Line: cy, Col: cx}}}).DoOnEditor(e)
				e.CloseBuffer(buf)
			}

		case lsp.ChangeDelete:
			if err = os.Remove(path); err == nil {
				changed = append(changed, protocol.FileEvent{URI: c.URI, Type: protocol.FileChangeTypeDeleted})
				if buf != nil {
					delete(pending, buf)
					e.CloseBuffer(buf)
				}
			}
		}
		if err != nil {
			errs = append(errs, fmt.Errorf("%s %s: %w", c.Kind, srv.displayPath(path), err))
		}
	}
	for _, buf := range order {
		flush(buf)
	}
	if len(changed) > 0 {
		// Not waiting for the file watcher: the server is to see the files before the next request.
		srv.background(func(ctx context.Context, client lspClient) {
			if err := client.DidChangeWatchedFiles(ctx, changed); err != nil {
				srv.le.Debugf("didChangeWatchedFiles failed: %s", err)
			}
		})
	}
	return errors.Join(errs...)
}

// editFile applies the edits to the file on the disk, keeping its permissions.
func editFile(path string, edits []protocol.TextEdit) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	text := string(data)
	edits = append([]protocol.TextEdit(nil), edits...)
	sort.SliceStable(edits, func(i, j int) bool { return positionBefore(edits[j].Range.Start, edits[i].Range.Start) })
	for _, te := range edits {
		text = applyChange(text, lsp.TextChange{Range: &te.Range, Text: te.NewText})
	}
	return os.WriteFile(path, []byte(text), info.Mode().Perm())
}

// background runs the request with the client of the server, which must have
// started, in the background. It's for the requests not about a buffer text.
func (srv *server) background(then func(ctx context.Context, client lspClient)) {
	if !srv.ready.Load() {
		return
	}
	go srv.runThen(srv.client, then)
}

// showDocument opens the document the server asks to show: a web page with what
// the server renders, like the documentation of a package, in the browser, or a
// file in the editor. It's called on the connection goroutine.
func (srv *server) showDocument(params protocol.ShowDocumentParams) error {
	if params.External {
		go func() {
			if err := openInBrowser(string(params.URI)); err != nil {
				srv.le.Logf("cannot open %s: %s", params.URI, err)
			}
		}()
		return nil
	}
	docUri := params.URI
	if docUri.Filename() == "" || srv.loop == nil {
		return fmt.Errorf("not a file: %s", params.URI)
	}
	go func() {
		var pos content.Position
		if params.Selection != nil {
			if data, err := os.ReadFile(docUri.Filename()); err == nil {
				pos = textPosition(string(data), params.Selection.Start)
			}
		}
		srv.loop.Send(&editor.GoTo{Location: navigation.Location{Path: docUri.Filename(), Position: pos}})
	}()
	return nil
}

// openInBrowser opens the URL with the default browser of the system.
func openInBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	return cmd.Run()
}

var (
	_ extrefactor.Renamer            = new(BufferData)
	_ extrefactor.CodeActionProvider = new(BufferData)
)
