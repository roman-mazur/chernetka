package lsp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

// The navigation and refactoring requests: references, call hierarchy, rename,
// and code actions.

// DidClose tells the server the document is not open anymore: the server reads
// the file from the disk from now on.
func (c *Client) DidClose(ctx context.Context, fileURI uri.URI) error {
	return c.conn.Notify(ctx, protocol.MethodTextDocumentDidClose, &protocol.DidCloseTextDocumentParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: fileURI},
	})
}

// DidChangeWatchedFiles tells the server about the files changed on the disk,
// which it asked to watch with Options.WatchFiles.
func (c *Client) DidChangeWatchedFiles(ctx context.Context, changes []protocol.FileEvent) error {
	params := protocol.DidChangeWatchedFilesParams{Changes: make([]*protocol.FileEvent, len(changes))}
	for i := range changes {
		params.Changes[i] = &changes[i]
	}
	return c.conn.Notify(ctx, protocol.MethodWorkspaceDidChangeWatchedFiles, &params)
}

// References returns the locations of the references to the symbol at the
// given zero-based position, including its declaration.
func (c *Client) References(ctx context.Context, fileURI uri.URI, pos protocol.Position) ([]protocol.Location, error) {
	var raw json.RawMessage
	if _, err := c.conn.Call(ctx, protocol.MethodTextDocumentReferences, &protocol.ReferenceParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: fileURI},
			Position:     pos,
		},
		Context: protocol.ReferenceContext{IncludeDeclaration: true},
	}, &raw); err != nil {
		return nil, err
	}
	return decodeLocations(raw)
}

// PrepareCallHierarchy returns the functions at the given zero-based position
// to get the calls of with IncomingCalls and OutgoingCalls. It's empty if
// there is none.
func (c *Client) PrepareCallHierarchy(ctx context.Context, fileURI uri.URI, pos protocol.Position) ([]protocol.CallHierarchyItem, error) {
	var items []protocol.CallHierarchyItem
	if _, err := c.conn.Call(ctx, protocol.MethodTextDocumentPrepareCallHierarchy, &protocol.CallHierarchyPrepareParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: fileURI},
			Position:     pos,
		},
	}, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// IncomingCalls returns the functions calling the item, with the ranges of the calls in them.
func (c *Client) IncomingCalls(ctx context.Context, item protocol.CallHierarchyItem) ([]protocol.CallHierarchyIncomingCall, error) {
	var calls []protocol.CallHierarchyIncomingCall
	if _, err := c.conn.Call(ctx, protocol.MethodCallHierarchyIncomingCalls, &protocol.CallHierarchyIncomingCallsParams{Item: item}, &calls); err != nil {
		return nil, err
	}
	return calls, nil
}

// OutgoingCalls returns the functions the item calls, with the ranges of the calls in the item.
func (c *Client) OutgoingCalls(ctx context.Context, item protocol.CallHierarchyItem) ([]protocol.CallHierarchyOutgoingCall, error) {
	var calls []protocol.CallHierarchyOutgoingCall
	if _, err := c.conn.Call(ctx, protocol.MethodCallHierarchyOutgoingCalls, &protocol.CallHierarchyOutgoingCallsParams{Item: item}, &calls); err != nil {
		return nil, err
	}
	return calls, nil
}

// RenameTarget is what Rename at a position changes: the range of the text to
// edit and the text itself, which the user edits into the new name. The text is
// usually the identifier, but gopls makes it the import path of the package to
// move on a package clause, and the signature of the function to change on its
// func keyword.
type RenameTarget struct {
	Range       protocol.Range
	Placeholder string
}

// PrepareRename returns what can be renamed at the given zero-based position:
// nil if nothing. The placeholder is the text of the range if the server
// doesn't give one, so text is the text of the document.
func (c *Client) PrepareRename(ctx context.Context, fileURI uri.URI, pos protocol.Position, text string) (*RenameTarget, error) {
	var raw json.RawMessage
	if _, err := c.conn.Call(ctx, protocol.MethodTextDocumentPrepareRename, &protocol.PrepareRenameParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: fileURI},
			Position:     pos,
		},
	}, &raw); err != nil {
		return nil, err
	}
	return decodeRenameTarget(raw, text, pos)
}

// decodeRenameTarget decodes a prepareRename result: null, a range, a range
// with a placeholder, or the default behavior (the word at the position).
func decodeRenameTarget(raw json.RawMessage, text string, pos protocol.Position) (*RenameTarget, error) {
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 || string(raw) == "null" {
		return nil, nil
	}
	var res struct {
		Range           *protocol.Range    `json:"range"`
		Placeholder     string             `json:"placeholder"`
		DefaultBehavior bool               `json:"defaultBehavior"`
		Start           *protocol.Position `json:"start"` // of a bare range
		End             *protocol.Position `json:"end"`
	}
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, fmt.Errorf("prepareRename result %s: %w", raw, err)
	}
	target := &RenameTarget{Placeholder: res.Placeholder}
	switch {
	case res.Range != nil:
		target.Range = *res.Range
	case res.DefaultBehavior:
		target.Range = wordRange(text, pos)
	case res.Start != nil && res.End != nil:
		target.Range = protocol.Range{Start: *res.Start, End: *res.End}
	default:
		return nil, nil
	}
	if target.Placeholder == "" {
		target.Placeholder = rangeText(text, target.Range)
	}
	if target.Placeholder == "" {
		return nil, nil
	}
	return target, nil
}

// Rename returns the edit renaming the symbol at the given zero-based position
// to newName. The server refuses the renames breaking the code with an error.
func (c *Client) Rename(ctx context.Context, fileURI uri.URI, pos protocol.Position, newName string) (*WorkspaceEdit, error) {
	var edit *WorkspaceEdit
	if _, err := c.conn.Call(ctx, protocol.MethodTextDocumentRename, &protocol.RenameParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: fileURI},
			Position:     pos,
		},
		NewName: newName,
	}, &edit); err != nil {
		return nil, err
	}
	return edit, nil
}

// CodeAction is a fix or a refactoring the server offers for a range of a
// document. It carries an edit, a command the server runs to make the change
// (sending the edits back with Options.ApplyEdit), or data to resolve the edit
// with ResolveCodeAction.
type CodeAction struct {
	Title       string                  `json:"title"`
	Kind        protocol.CodeActionKind `json:"kind,omitempty"`
	IsPreferred bool                    `json:"isPreferred,omitempty"`
	Edit        *WorkspaceEdit          `json:"edit,omitempty"`
	Command     *protocol.Command       `json:"command,omitempty"`
	Data        json.RawMessage         `json:"data,omitempty"`
}

// CodeActions returns the fixes and the refactorings offered for the range of
// the document. The diagnostics are the ones the server published for the
// range: the fixes are matched to them.
func (c *Client) CodeActions(ctx context.Context, fileURI uri.URI, rng protocol.Range, diagnostics []protocol.Diagnostic) ([]CodeAction, error) {
	if diagnostics == nil {
		diagnostics = []protocol.Diagnostic{}
	}
	var actions []CodeAction
	if _, err := c.conn.Call(ctx, protocol.MethodTextDocumentCodeAction, &protocol.CodeActionParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: fileURI},
		Range:        rng,
		Context:      protocol.CodeActionContext{Diagnostics: diagnostics},
	}, &actions); err != nil {
		return nil, err
	}
	return actions, nil
}

// ResolveCodeAction asks the server for the edit of the action it returned with data only.
func (c *Client) ResolveCodeAction(ctx context.Context, action CodeAction) (CodeAction, error) {
	var resolved CodeAction
	if _, err := c.conn.Call(ctx, methodCodeActionResolve, &action, &resolved); err != nil {
		return CodeAction{}, err
	}
	return resolved, nil
}

// ExecuteCommand runs the command of a code action on the server. The server
// sends the edits it makes with Options.ApplyEdit.
func (c *Client) ExecuteCommand(ctx context.Context, cmd protocol.Command) error {
	var res json.RawMessage
	_, err := c.conn.Call(ctx, protocol.MethodWorkspaceExecuteCommand, &protocol.ExecuteCommandParams{
		Command:   cmd.Command,
		Arguments: cmd.Arguments,
	}, &res)
	return err
}

// WorkspaceEdit is a change of the files of the workspace: the text edits of
// the documents, and the files created, renamed, and deleted, in the order they
// are to be applied. protocol.WorkspaceEdit doesn't decode the file operations.
type WorkspaceEdit struct {
	Changes []DocumentChange
}

// DocumentChange is an edit of a document or an operation on a file.
type DocumentChange struct {
	Kind   ChangeKind
	URI    uri.URI             // the document edited, the file created or deleted, or the old name of the renamed file
	NewURI uri.URI             // the new name of the renamed file
	Edits  []protocol.TextEdit // the edits of the document, for ChangeEdit

	Overwrite bool // create and rename over an existing file; otherwise the operation fails if it exists
}

// ChangeKind is the kind of a DocumentChange.
type ChangeKind int

const (
	ChangeEdit ChangeKind = iota
	ChangeCreate
	ChangeRename
	ChangeDelete
)

func (k ChangeKind) String() string {
	switch k {
	case ChangeEdit:
		return "edit"
	case ChangeCreate:
		return "create"
	case ChangeRename:
		return "rename"
	case ChangeDelete:
		return "delete"
	}
	return fmt.Sprintf("ChangeKind(%d)", int(k))
}

// UnmarshalJSON decodes the LSP WorkspaceEdit: documentChanges if given,
// otherwise the changes by document.
func (we *WorkspaceEdit) UnmarshalJSON(data []byte) error {
	var raw struct {
		Changes         map[uri.URI][]protocol.TextEdit `json:"changes"`
		DocumentChanges []json.RawMessage               `json:"documentChanges"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	we.Changes = nil
	for _, rawChange := range raw.DocumentChanges {
		var dc struct {
			Kind         string                          `json:"kind"`
			URI          uri.URI                         `json:"uri"`
			OldURI       uri.URI                         `json:"oldUri"`
			NewURI       uri.URI                         `json:"newUri"`
			TextDocument protocol.TextDocumentIdentifier `json:"textDocument"`
			Edits        []protocol.TextEdit             `json:"edits"`
			Options      struct {
				Overwrite bool `json:"overwrite"`
			} `json:"options"`
		}
		if err := json.Unmarshal(rawChange, &dc); err != nil {
			return err
		}
		change := DocumentChange{Overwrite: dc.Options.Overwrite}
		switch dc.Kind {
		case "":
			change.URI, change.Edits = dc.TextDocument.URI, dc.Edits
		case "create":
			change.Kind, change.URI = ChangeCreate, dc.URI
		case "rename":
			change.Kind, change.URI, change.NewURI = ChangeRename, dc.OldURI, dc.NewURI
		case "delete":
			change.Kind, change.URI = ChangeDelete, dc.URI
		default:
			return fmt.Errorf("unknown document change kind %q", dc.Kind)
		}
		we.Changes = append(we.Changes, change)
	}
	if len(raw.DocumentChanges) == 0 {
		for docURI, edits := range raw.Changes {
			we.Changes = append(we.Changes, DocumentChange{URI: docURI, Edits: edits})
		}
	}
	return nil
}

// MarshalJSON encodes the edit as documentChanges, like it was received, so
// a code action can be sent back to be resolved.
func (we WorkspaceEdit) MarshalJSON() ([]byte, error) {
	changes := make([]map[string]any, 0, len(we.Changes))
	for _, c := range we.Changes {
		switch c.Kind {
		case ChangeEdit:
			changes = append(changes, map[string]any{
				"textDocument": protocol.TextDocumentIdentifier{URI: c.URI},
				"edits":        c.Edits,
			})
		case ChangeCreate:
			changes = append(changes, map[string]any{"kind": "create", "uri": c.URI, "options": map[string]bool{"overwrite": c.Overwrite}})
		case ChangeRename:
			changes = append(changes, map[string]any{"kind": "rename", "oldUri": c.URI, "newUri": c.NewURI, "options": map[string]bool{"overwrite": c.Overwrite}})
		case ChangeDelete:
			changes = append(changes, map[string]any{"kind": "delete", "uri": c.URI})
		}
	}
	return json.Marshal(map[string]any{"documentChanges": changes})
}

// ApplyEditRequest is what the server asks the client to change with workspace/applyEdit.
type ApplyEditRequest struct {
	Label string        `json:"label,omitempty"`
	Edit  WorkspaceEdit `json:"edit"`
}

// wordRange returns the range of the identifier at pos in text, as the default
// behavior of prepareRename asks. It's empty if there is no identifier.
func wordRange(text string, pos protocol.Position) protocol.Range {
	line := lineAt(text, pos.Line)
	runes := []rune(line)
	// Work in runes: the characters outside of the basic plane are rare in identifiers.
	i := int(pos.Character)
	if i > len(runes) {
		return protocol.Range{Start: pos, End: pos} // Past the end of the line: no identifier.
	}
	start, end := i, i
	for start > 0 && isIdent(runes[start-1]) {
		start--
	}
	for end < len(runes) && isIdent(runes[end]) {
		end++
	}
	return protocol.Range{
		Start: protocol.Position{Line: pos.Line, Character: utf16Len(string(runes[:start]))},
		End:   protocol.Position{Line: pos.Line, Character: utf16Len(string(runes[:end]))},
	}
}

func isIdent(r rune) bool {
	return r == '_' || r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r > 0x7f
}

// rangeText returns the text of the range on a single line of text.
func rangeText(text string, rng protocol.Range) string {
	if rng.Start.Line != rng.End.Line {
		return ""
	}
	line := lineAt(text, rng.Start.Line)
	start, end := utf16Offset(line, rng.Start.Character), utf16Offset(line, rng.End.Character)
	return line[start:end]
}

// lineAt returns the line of text with the zero-based number, without the line ending.
func lineAt(text string, n uint32) string {
	for ; n > 0; n-- {
		i := bytes.IndexByte([]byte(text), '\n')
		if i < 0 {
			return ""
		}
		text = text[i+1:]
	}
	if i := bytes.IndexByte([]byte(text), '\n'); i >= 0 {
		text = text[:i]
	}
	return text
}

// utf16Offset converts the UTF-16 offset in line to a byte offset.
func utf16Offset(line string, units uint32) int {
	var n uint32
	for i, r := range line {
		if n >= units {
			return i
		}
		n++
		if r >= 0x10000 {
			n++
		}
	}
	return len(line)
}

func utf16Len(s string) (n uint32) {
	for _, r := range s {
		n++
		if r >= 0x10000 {
			n++
		}
	}
	return n
}

// methodCodeActionResolve is the codeAction/resolve method, which protocol doesn't name.
const methodCodeActionResolve = "codeAction/resolve"
