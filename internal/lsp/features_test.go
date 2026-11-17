package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

// TestClient_Features drives the navigation and refactoring requests against
// a fake server, which also sends the requests of its own a refactoring
// involves: the edits to apply, the document to show, and the files to watch.
func TestClient_Features(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientConn, serverConn := pipeConns(t)
	fileURI := uri.File("/tmp/x.go")
	loc := protocol.Location{URI: fileURI, Range: protocol.Range{
		Start: protocol.Position{Line: 1, Character: 2}, End: protocol.Position{Line: 1, Character: 5},
	}}
	fnItem := protocol.CallHierarchyItem{Name: "foo", Kind: protocol.SymbolKindFunction, URI: fileURI, Range: loc.Range, SelectionRange: loc.Range}

	var (
		notified = make(chan string, 10) // the notifications received, as "method params"
		executed = make(chan string, 1)
	)
	serverConn.Go(ctx, jsonrpc2.AsyncHandler(jsonrpc2.ReplyHandler(func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		switch req.Method() {
		case protocol.MethodInitialize:
			var p protocol.InitializeParams
			if err := jsonrpc2DecodeParams(req, &p); err != nil {
				return reply(ctx, nil, err)
			}
			caps := p.Capabilities
			switch {
			case !caps.TextDocument.Rename.PrepareSupport:
				return reply(ctx, nil, fmt.Errorf("no prepareRename support"))
			case !caps.Workspace.ApplyEdit, !caps.Workspace.WorkspaceEdit.DocumentChanges, len(caps.Workspace.WorkspaceEdit.ResourceOperations) != 3:
				return reply(ctx, nil, fmt.Errorf("no workspace edit support: %+v", caps.Workspace))
			case !caps.Workspace.DidChangeWatchedFiles.DynamicRegistration:
				return reply(ctx, nil, fmt.Errorf("no file watching"))
			case !caps.Window.ShowDocument.Support:
				return reply(ctx, nil, fmt.Errorf("no show document support"))
			case !caps.TextDocument.CodeAction.DataSupport, caps.TextDocument.CodeAction.ResolveSupport.Properties[0] != "edit":
				return reply(ctx, nil, fmt.Errorf("no code action resolve support"))
			}
			return reply(ctx, &protocol.InitializeResult{}, nil)
		case protocol.MethodInitialized:
			// Ask to watch the files, like gopls does once the workspace is loaded.
			go func() {
				_, _ = serverConn.Call(ctx, protocol.MethodClientRegisterCapability, &protocol.RegistrationParams{
					Registrations: []protocol.Registration{{
						ID: "1", Method: protocol.MethodWorkspaceDidChangeWatchedFiles,
						RegisterOptions: map[string]any{"watchers": []any{
							map[string]any{"globPattern": "/tmp/**/*.{go,mod}"},
							map[string]any{"globPattern": map[string]any{"baseUri": "file:///tmp/vendor", "pattern": "**/*.go"}},
						}},
					}},
				}, nil)
			}()
			return reply(ctx, nil, nil)
		case protocol.MethodTextDocumentDidClose, protocol.MethodWorkspaceDidChangeWatchedFiles:
			notified <- req.Method() + " " + string(req.Params())
			return reply(ctx, nil, nil)
		case protocol.MethodTextDocumentReferences:
			var p protocol.ReferenceParams
			if err := jsonrpc2DecodeParams(req, &p); err != nil || !p.Context.IncludeDeclaration || p.Position.Line != 1 {
				return reply(ctx, nil, fmt.Errorf("unexpected params %+v (%v)", p, err))
			}
			return reply(ctx, []protocol.Location{loc, loc}, nil)
		case protocol.MethodTextDocumentPrepareCallHierarchy:
			return reply(ctx, []protocol.CallHierarchyItem{fnItem}, nil)
		case protocol.MethodCallHierarchyIncomingCalls:
			var p protocol.CallHierarchyIncomingCallsParams
			if err := jsonrpc2DecodeParams(req, &p); err != nil || p.Item.Name != "foo" {
				return reply(ctx, nil, fmt.Errorf("unexpected params %+v (%v)", p, err))
			}
			return reply(ctx, []protocol.CallHierarchyIncomingCall{{From: fnItem, FromRanges: []protocol.Range{loc.Range}}}, nil)
		case protocol.MethodCallHierarchyOutgoingCalls:
			return reply(ctx, []protocol.CallHierarchyOutgoingCall{{To: fnItem, FromRanges: []protocol.Range{loc.Range}}}, nil)
		case protocol.MethodTextDocumentPrepareRename:
			var p protocol.PrepareRenameParams
			_ = jsonrpc2DecodeParams(req, &p)
			switch p.Position.Line {
			case 0:
				return reply(ctx, map[string]any{"range": loc.Range, "placeholder": "func(x int)"}, nil)
			case 1:
				return reply(ctx, loc.Range, nil)
			case 2:
				return reply(ctx, map[string]any{"defaultBehavior": true}, nil)
			default:
				return reply(ctx, nil, nil)
			}
		case protocol.MethodTextDocumentRename:
			var p protocol.RenameParams
			if err := jsonrpc2DecodeParams(req, &p); err != nil || p.NewName != "bar" {
				return reply(ctx, nil, fmt.Errorf("unexpected params %+v (%v)", p, err))
			}
			return reply(ctx, json.RawMessage(`{"documentChanges": [
				{"textDocument": {"uri": "file:///tmp/x.go", "version": 3}, "edits": [{"range": {"start": {"line": 1, "character": 2}, "end": {"line": 1, "character": 5}}, "newText": "bar"}]},
				{"kind": "rename", "oldUri": "file:///tmp/x.go", "newUri": "file:///tmp/y.go"},
				{"kind": "create", "uri": "file:///tmp/z.go", "options": {"overwrite": true}},
				{"kind": "delete", "uri": "file:///tmp/w.go"}
			]}`), nil)
		case protocol.MethodTextDocumentCodeAction:
			var p protocol.CodeActionParams
			if err := jsonrpc2DecodeParams(req, &p); err != nil || len(p.Context.Diagnostics) != 1 || p.Context.Diagnostics[0].Message != "unused" {
				return reply(ctx, nil, fmt.Errorf("unexpected params %+v (%v)", p, err))
			}
			return reply(ctx, []any{
				map[string]any{"title": "Remove x", "kind": "quickfix", "edit": map[string]any{"changes": map[string]any{
					"file:///tmp/x.go": []protocol.TextEdit{{Range: loc.Range}},
				}}},
				map[string]any{"title": "Extract function", "kind": "refactor.extract.function", "data": map[string]any{"command": "gopls.apply_fix"}},
				map[string]any{"title": "Browse documentation", "kind": "source.doc", "command": protocol.Command{Title: "Doc", Command: "gopls.doc"}},
			}, nil)
		case methodCodeActionResolve:
			var action CodeAction
			if err := jsonrpc2DecodeParams(req, &action); err != nil || string(action.Data) != `{"command":"gopls.apply_fix"}` {
				return reply(ctx, nil, fmt.Errorf("unexpected action %+v (%v)", action, err))
			}
			action.Edit = &WorkspaceEdit{Changes: []DocumentChange{{URI: fileURI, Edits: []protocol.TextEdit{{NewText: "extracted"}}}}}
			return reply(ctx, action, nil)
		case protocol.MethodWorkspaceExecuteCommand:
			var p protocol.ExecuteCommandParams
			_ = jsonrpc2DecodeParams(req, &p)
			executed <- p.Command
			// The command makes its edits through the client, and shows a document.
			// The handlers run asynchronously: the replies of the client are read meanwhile.
			var res protocol.ApplyWorkspaceEditResponse
			if _, err := serverConn.Call(ctx, protocol.MethodWorkspaceApplyEdit, &ApplyEditRequest{
				Label: "doc", Edit: WorkspaceEdit{Changes: []DocumentChange{{URI: fileURI, Edits: []protocol.TextEdit{{NewText: "edited"}}}}},
			}, &res); err != nil || !res.Applied {
				return reply(ctx, nil, fmt.Errorf("apply edit: %+v, %v", res, err))
			}
			var shown protocol.ShowDocumentResult
			if _, err := serverConn.Call(ctx, protocol.MethodShowDocument, &protocol.ShowDocumentParams{URI: "http://localhost/doc", External: true}, &shown); err != nil || !shown.Success {
				return reply(ctx, nil, fmt.Errorf("show document: %+v, %v", shown, err))
			}
			return reply(ctx, nil, nil)
		}
		return reply(ctx, nil, nil)
	})))

	var (
		applied = make(chan ApplyEditRequest, 1)
		shown   = make(chan protocol.ShowDocumentParams, 1)
		watched = make(chan []string, 1)
	)
	opts := Options{
		ApplyEdit:    func(req ApplyEditRequest) error { applied <- req; return nil },
		ShowDocument: func(p protocol.ShowDocumentParams) error { shown <- p; return nil },
		WatchFiles:   func(patterns []string) { watched <- patterns },
	}
	c, err := newClient(ctx, clientConn, nil, "/tmp", opts)
	if err != nil {
		t.Fatalf("newClient: %v", err)
	}

	select {
	case patterns := <-watched:
		if want := "/tmp/**/*.{go,mod} /tmp/vendor/**/*.go"; strings.Join(patterns, " ") != want {
			t.Errorf("watch patterns %q, want %q", patterns, want)
		}
	case <-ctx.Done():
		t.Fatal("the files to watch are not registered")
	}

	locs, err := c.References(ctx, fileURI, protocol.Position{Line: 1, Character: 3})
	if err != nil || len(locs) != 2 || locs[0] != loc {
		t.Errorf("References = %+v, %v", locs, err)
	}

	items, err := c.PrepareCallHierarchy(ctx, fileURI, loc.Range.Start)
	if err != nil || len(items) != 1 || items[0].Name != "foo" {
		t.Fatalf("PrepareCallHierarchy = %+v, %v", items, err)
	}
	in, err := c.IncomingCalls(ctx, items[0])
	if err != nil || len(in) != 1 || in[0].From.Name != "foo" || len(in[0].FromRanges) != 1 {
		t.Errorf("IncomingCalls = %+v, %v", in, err)
	}
	out, err := c.OutgoingCalls(ctx, items[0])
	if err != nil || len(out) != 1 || out[0].To.Name != "foo" {
		t.Errorf("OutgoingCalls = %+v, %v", out, err)
	}

	text := "package main\n\nfunc main() {}\nvar xyz = 1\n"
	for _, tc := range []struct {
		line uint32
		want *RenameTarget
	}{
		{0, &RenameTarget{Range: loc.Range, Placeholder: "func(x int)"}},
		{1, nil}, // The range is on line 1: "" (an empty line), so there is nothing to rename.
		{2, &RenameTarget{Placeholder: "main", Range: protocol.Range{Start: protocol.Position{Line: 2, Character: 5}, End: protocol.Position{Line: 2, Character: 9}}}},
		{3, nil},
	} {
		got, err := c.PrepareRename(ctx, fileURI, protocol.Position{Line: tc.line, Character: 6}, text)
		if err != nil {
			t.Errorf("PrepareRename on line %d: %v", tc.line, err)
			continue
		}
		if (got == nil) != (tc.want == nil) || (got != nil && *got != *tc.want) {
			t.Errorf("PrepareRename on line %d = %+v, want %+v", tc.line, got, tc.want)
		}
	}

	edit, err := c.Rename(ctx, fileURI, loc.Range.Start, "bar")
	if err != nil {
		t.Fatalf("Rename: %v", err)
	}
	want := []DocumentChange{
		{URI: fileURI, Edits: []protocol.TextEdit{{Range: loc.Range, NewText: "bar"}}},
		{Kind: ChangeRename, URI: fileURI, NewURI: uri.File("/tmp/y.go")},
		{Kind: ChangeCreate, URI: uri.File("/tmp/z.go"), Overwrite: true},
		{Kind: ChangeDelete, URI: uri.File("/tmp/w.go")},
	}
	if fmt.Sprint(edit.Changes) != fmt.Sprint(want) {
		t.Errorf("Rename = %+v, want %+v", edit.Changes, want)
	}

	actions, err := c.CodeActions(ctx, fileURI, loc.Range, []protocol.Diagnostic{{Message: "unused"}})
	if err != nil || len(actions) != 3 {
		t.Fatalf("CodeActions = %+v, %v", actions, err)
	}
	if a := actions[0]; a.Title != "Remove x" || a.Kind != protocol.QuickFix || a.Edit == nil || len(a.Edit.Changes) != 1 || a.Edit.Changes[0].URI != fileURI {
		t.Errorf("action 0 = %+v", a)
	}
	if a := actions[1]; a.Edit != nil || a.Command != nil || len(a.Data) == 0 {
		t.Errorf("action 1 = %+v", a)
	}
	resolved, err := c.ResolveCodeAction(ctx, actions[1])
	if err != nil || resolved.Edit == nil || resolved.Edit.Changes[0].Edits[0].NewText != "extracted" {
		t.Errorf("ResolveCodeAction = %+v, %v", resolved, err)
	}
	if a := actions[2]; a.Command == nil || a.Command.Command != "gopls.doc" {
		t.Fatalf("action 2 = %+v", a)
	}
	if err := c.ExecuteCommand(ctx, *actions[2].Command); err != nil {
		t.Fatalf("ExecuteCommand: %v", err)
	}
	if got := <-executed; got != "gopls.doc" {
		t.Errorf("executed %q", got)
	}
	if req := <-applied; req.Label != "doc" || len(req.Edit.Changes) != 1 || req.Edit.Changes[0].Edits[0].NewText != "edited" {
		t.Errorf("applied %+v", req)
	}
	if p := <-shown; p.URI != "http://localhost/doc" || !p.External {
		t.Errorf("shown %+v", p)
	}

	if err := c.DidClose(ctx, fileURI); err != nil {
		t.Fatalf("DidClose: %v", err)
	}
	if err := c.DidChangeWatchedFiles(ctx, []protocol.FileEvent{{URI: fileURI, Type: protocol.FileChangeTypeChanged}}); err != nil {
		t.Fatalf("DidChangeWatchedFiles: %v", err)
	}
	for _, want := range []string{
		`textDocument/didClose {"textDocument":{"uri":"file:///tmp/x.go"}}`,
		`workspace/didChangeWatchedFiles {"changes":[{"type":2,"uri":"file:///tmp/x.go"}]}`,
	} {
		if got := <-notified; got != want {
			t.Errorf("notified %s, want %s", got, want)
		}
	}
}

func TestWorkspaceEdit_Changes(t *testing.T) {
	// Without documentChanges, the edits by document are used.
	var edit WorkspaceEdit
	if err := json.Unmarshal([]byte(`{"changes": {"file:///a.go": [{"range": {"start": {"line": 0, "character": 0}, "end": {"line": 0, "character": 1}}, "newText": "x"}]}}`), &edit); err != nil {
		t.Fatal(err)
	}
	if len(edit.Changes) != 1 || edit.Changes[0].URI != uri.File("/a.go") || edit.Changes[0].Edits[0].NewText != "x" {
		t.Errorf("changes %+v", edit.Changes)
	}
	// Encoding and decoding the edit keeps the changes.
	edit.Changes = append(edit.Changes, DocumentChange{Kind: ChangeRename, URI: uri.File("/a.go"), NewURI: uri.File("/b.go"), Overwrite: true})
	data, err := json.Marshal(edit)
	if err != nil {
		t.Fatal(err)
	}
	var decoded WorkspaceEdit
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if fmt.Sprint(decoded.Changes) != fmt.Sprint(edit.Changes) {
		t.Errorf("decoded %+v, want %+v", decoded.Changes, edit.Changes)
	}
}

func TestWordRange(t *testing.T) {
	text := "a\nvar 日本語x = y\n"
	for _, tc := range []struct {
		pos  protocol.Position
		want string
	}{
		{protocol.Position{Line: 1, Character: 4}, "日本語x"},
		{protocol.Position{Line: 1, Character: 7}, "日本語x"},
		{protocol.Position{Line: 1, Character: 8}, "日本語x"},
		{protocol.Position{Line: 1, Character: 9}, ""},
		{protocol.Position{Line: 1, Character: 12}, "y"},
		{protocol.Position{Line: 1, Character: 99}, ""},
		{protocol.Position{Line: 5, Character: 0}, ""},
	} {
		if got := rangeText(text, wordRange(text, tc.pos)); got != tc.want {
			t.Errorf("word at %v = %q, want %q", tc.pos, got, tc.want)
		}
	}
}
