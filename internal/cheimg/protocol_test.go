package cheimg

import (
	"errors"
	"testing"

	"rmazur.io/chernetka/internal/remotectl"
)

func TestItem_Command(t *testing.T) {
	for _, it := range []Item{
		{Path: "/tmp/a.d2"},
		{Path: "/tmp/a.png"},
		{Source: "a -> b\nb -> c"},
		{Source: "a -> b", Path: "/tmp/README.md"},
	} {
		got, err := ParseCommand(*it.Command())
		if err != nil {
			t.Errorf("ParseCommand(%+v.Command()): %s", it, err)
		}
		if got != it {
			t.Errorf("ParseCommand(%+v.Command()) = %+v", it, got)
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

func TestItem_Kind(t *testing.T) {
	for _, tc := range []struct {
		item Item
		want Kind
	}{
		{Item{Path: "a.png"}, KindImage},
		{Item{Path: "/x/Photo.PNG"}, KindImage},
		{Item{Path: "arch.d2"}, KindDiagram},
		{Item{Path: "README.md"}, KindUnsupported},
		{Item{Path: "png"}, KindUnsupported},
		{Item{Path: "README.md", Source: "a -> b"}, KindDiagram},
		{Item{Source: "a -> b"}, KindDiagram},
	} {
		if got := tc.item.Kind(); got != tc.want {
			t.Errorf("%+v.Kind() = %d, want %d", tc.item, got, tc.want)
		}
	}
}
