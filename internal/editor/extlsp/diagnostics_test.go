package extlsp

import (
	"reflect"
	"testing"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
	"rmazur.io/chernetka/internal/content/code"
)

func TestIntegration_Diagnostics(t *testing.T) {
	var le Integration
	h, _, data := newGoBuffer(t, &le, &fakeLSP{}, "package main\n\nfunc main() {\n\tx\n}\n")

	diag := func(line uint32, sev protocol.DiagnosticSeverity, msg string) protocol.Diagnostic {
		return protocol.Diagnostic{
			Range:    protocol.Range{Start: protocol.Position{Line: line}, End: protocol.Position{Line: line, Character: 2}},
			Severity: sev,
			Message:  msg,
		}
	}
	// Published before the editor loop runs: it must not wait for it.
	data.srv.publishDiagnostics(protocol.PublishDiagnosticsParams{
		URI: data.docUri,
		Diagnostics: []protocol.Diagnostic{
			diag(3, protocol.DiagnosticSeverityError, "x (variable of type int) is not used"),
			diag(2, protocol.DiagnosticSeverityHint, "could be simplified"),
			diag(0, protocol.DiagnosticSeverityWarning, "package comment is missing"),
			diag(1, 0, "no severity"),
			diag(1, protocol.DiagnosticSeverityInformation, "for your information"),
		},
	})
	data.srv.publishDiagnostics(protocol.PublishDiagnosticsParams{
		URI:         uri.File("/elsewhere/other.go"),
		Diagnostics: []protocol.Diagnostic{diag(0, protocol.DiagnosticSeverityError, "not this buffer")},
	})
	h.Run(t)

	want := []code.Diagnostic{
		{Line: 0, Severity: code.SeverityWarning, Message: "package comment is missing"},
		{Line: 1, Severity: code.SeverityError, Message: "no severity"},
		{Line: 3, Severity: code.SeverityError, Message: "x (variable of type int) is not used"},
	}
	if got := onLoop(t, h, data.Diagnostics); !reflect.DeepEqual(got, want) {
		t.Errorf("diagnostics = %+v, want %+v", got, want)
	}

	// The server clears them when the problems are fixed.
	data.srv.publishDiagnostics(protocol.PublishDiagnosticsParams{URI: data.docUri})
	if got := onLoop(t, h, data.Diagnostics); len(got) != 0 {
		t.Errorf("diagnostics = %+v after clearing", got)
	}
	if got := data.srv.diagnostics(uri.File("/elsewhere/other.go")); len(got) != 1 {
		t.Errorf("diagnostics of another document = %+v", got)
	}
}
