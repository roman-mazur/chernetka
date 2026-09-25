// Package cheimg defines how to ask a running che-img process to show an image or a diagram.
// It does not depend on d2 itself, so the editor can use it without the diagram rendering machinery.
package cheimg

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"rmazur.io/chernetka/internal/remotectl"
)

// Endpoint is the socket che-img listens on.
const Endpoint remotectl.Endpoint = "che-img"

// Kind tells what che-img displays.
type Kind int

const (
	KindUnsupported Kind = iota
	KindImage            // a PNG image
	KindDiagram          // a d2 diagram
)

// KindOf reports what che-img can display from the file at path judging by its extension.
func KindOf(path string) Kind {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".png":
		return KindImage
	case ".d2":
		return KindDiagram
	default:
		return KindUnsupported
	}
}

// Item is what che-img shows: an image or a diagram file at Path, or the Source of a d2 diagram.
// When both are set, the Source is shown and Path is used to resolve imports and to name the diagram.
type Item struct {
	Path   string
	Source string
}

// Kind tells what the item is.
func (it Item) Kind() Kind {
	if it.Source != "" {
		return KindDiagram
	}
	return KindOf(it.Path)
}

// Name describes the item for the user.
func (it Item) Name() string {
	if it.Path != "" {
		return it.Path
	}
	return "(source)"
}

const (
	actionFile   = "file"
	actionSource = "source"
)

// Command encodes the request to show the item.
func (it Item) Command() *remotectl.CommandData {
	if it.Source == "" {
		return &remotectl.CommandData{Action: actionFile, Args: []string{it.Path}}
	}
	return &remotectl.CommandData{Action: actionSource, Args: []string{it.Source, it.Path}}
}

var errBadCommand = errors.New("bad che-img command")

// ParseCommand decodes the Item from the data encoded with Item.Command.
func ParseCommand(cmd remotectl.CommandData) (Item, error) {
	switch {
	case cmd.Action == actionFile && len(cmd.Args) == 1 && cmd.Args[0] != "":
		return Item{Path: cmd.Args[0]}, nil
	case cmd.Action == actionSource && len(cmd.Args) == 2:
		return Item{Source: cmd.Args[0], Path: cmd.Args[1]}, nil
	default:
		return Item{}, fmt.Errorf("%w: %s with %d args", errBadCommand, cmd.Action, len(cmd.Args))
	}
}

// Show sends the item to the running che-img process.
// It fails if there is no such process.
func Show(it Item) error {
	return remotectl.SendCommand(Endpoint, it.Command())
}
