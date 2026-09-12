package inputs

import (
	"fmt"
	"testing"
)

func TestIsCursor(t *testing.T) {
	for i, tc := range []struct {
		input    []byte
		match    bool
		typ      Cursor
		hasShift bool
	}{
		{
			input: nil,
			match: false,
		},
		{
			input: []byte{42},
			match: false,
		},
		{
			input:    []byte{27, '[', '1', ';', '1', '0', 'C'},
			match:    true,
			typ:      CursorEnd,
			hasShift: true,
		},
		{
			input:    []byte{27, '[', '1', ';', '2', 'C'},
			match:    true,
			typ:      CursorArrowRight,
			hasShift: true,
		},
		{
			input: []byte{27, '[', 'D'},
			match: true,
			typ:   CursorArrowLeft,
		},
		{
			input: []byte{27, '[', 'A'},
			match: true,
			typ:   CursorArrowUp,
		},
	} {
		t.Run(fmt.Sprintf("%d/typ=%s/shift=%t", i, tc.typ, tc.hasShift), func(t *testing.T) {
			var (
				typ Cursor
				mod Modifier
			)
			if res := IsCursor(tc.input, &typ, &mod); res != tc.match {
				t.Errorf("IsCursor() = %t, want %t", res, tc.match)
			}
			if tc.match {
				if typ != tc.typ {
					t.Errorf("typ=%s, want %s", typ, tc.typ)
				}
				if hasShift := mod.HasShift(); hasShift != tc.hasShift {
					t.Errorf("mod.HaShift()=%t, want %t", hasShift, tc.hasShift)
				}
			}
		})
	}
}
