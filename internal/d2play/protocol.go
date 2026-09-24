// Package d2play defines how to ask a running d2play process to show a diagram.
// It does not depend on d2 itself, so the editor can use it without the diagram rendering machinery.
package d2play

import (
	"errors"
	"fmt"

	"rmazur.io/chernetka/internal/remotectl"
)

// Endpoint is the socket d2play listens on.
const Endpoint remotectl.Endpoint = "d2play"

const (
	actionFile   = "file"
	actionSource = "source"
)

// Diagram points to the d2 source to visualize.
// It's either a file at Path or the Source text itself. When both are set, the Source is shown
// and Path is used to resolve imports and to name the diagram.
type Diagram struct {
	Path   string
	Source string
}

// Command encodes the request to show the diagram.
func (d Diagram) Command() *remotectl.CommandData {
	if d.Source == "" {
		return &remotectl.CommandData{Action: actionFile, Args: []string{d.Path}}
	}
	return &remotectl.CommandData{Action: actionSource, Args: []string{d.Source, d.Path}}
}

// Name describes the diagram for the user.
func (d Diagram) Name() string {
	if d.Path != "" {
		return d.Path
	}
	return "(source)"
}

var errBadCommand = errors.New("bad d2play command")

// ParseCommand decodes the Diagram from the data encoded with Diagram.Command.
func ParseCommand(cmd remotectl.CommandData) (Diagram, error) {
	switch {
	case cmd.Action == actionFile && len(cmd.Args) == 1 && cmd.Args[0] != "":
		return Diagram{Path: cmd.Args[0]}, nil
	case cmd.Action == actionSource && len(cmd.Args) == 2:
		return Diagram{Source: cmd.Args[0], Path: cmd.Args[1]}, nil
	default:
		return Diagram{}, fmt.Errorf("%w: %s with %d args", errBadCommand, cmd.Action, len(cmd.Args))
	}
}

// Show sends the diagram to the running d2play process.
// It fails if there is no such process.
func Show(d Diagram) error {
	return remotectl.SendCommand(Endpoint, d.Command())
}
