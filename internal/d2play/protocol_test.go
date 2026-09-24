package d2play

import (
	"errors"
	"testing"

	"rmazur.io/chernetka/internal/remotectl"
)

func TestDiagram_Command(t *testing.T) {
	for _, d := range []Diagram{
		{Path: "/tmp/a.d2"},
		{Source: "a -> b\nb -> c"},
		{Source: "a -> b", Path: "/tmp/README.md"},
	} {
		got, err := ParseCommand(*d.Command())
		if err != nil {
			t.Errorf("ParseCommand(%+v.Command()): %s", d, err)
		}
		if got != d {
			t.Errorf("ParseCommand(%+v.Command()) = %+v", d, got)
		}
	}
}

func TestParseCommand_Errors(t *testing.T) {
	for _, cmd := range []remotectl.CommandData{
		{Action: "open", Args: []string{"a.d2"}},
		{Action: actionFile},
		{Action: actionFile, Args: []string{""}},
		{Action: actionSource, Args: []string{"a -> b"}},
	} {
		if _, err := ParseCommand(cmd); !errors.Is(err, errBadCommand) {
			t.Errorf("ParseCommand(%+v) error = %v", cmd, err)
		}
	}
}
