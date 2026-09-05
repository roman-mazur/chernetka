package inputs

import (
	"errors"
	"strings"
	"testing"
	"testing/iotest"
)

func TestConsumeClipboardPaste_NotAPaste(t *testing.T) {
	for _, tc := range []struct {
		name string
		b    []byte
	}{
		{"too short", []byte{Escape, '['}},
		{"empty", nil},
		{"no escape", []byte("[200~hello")},
		{"wrong marker", []byte("\x1b[201~hello")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			content, err := ConsumeClipboardPaste(tc.b, strings.NewReader(""))
			if err != nil {
				t.Fatalf("err = %v, want nil", err)
			}
			if content != "" {
				t.Errorf("content = %q, want empty", content)
			}
		})
	}
}

func TestConsumeClipboardPaste(t *testing.T) {
	for _, tc := range []struct {
		name string
		b    []byte // the bytes already read before the paste was recognized
		rest string // the rest of the stream, split byte-by-byte on read
		want string
	}{
		{
			name: "empty paste",
			b:    []byte("\x1b[200~\x1b[201~"),
			want: "",
		},
		{
			name: "everything in the initial buffer",
			b:    []byte("\x1b[200~hello\x1b[201~"),
			want: "hello",
		},
		{
			name: "content split between buffer and reader",
			b:    []byte("\x1b[200~hel"),
			rest: "lo\x1b[201~",
			want: "hello",
		},
		{
			name: "terminator marker split between buffer and reader",
			b:    []byte("\x1b[200~hello\x1b[2"),
			rest: "01~",
			want: "hello",
		},
		{
			name: "multiline content",
			b:    []byte("\x1b[200~line1\nline2\n\x1b[201~"),
			want: "line1\nline2\n",
		},
		{
			name: "content contains a lone escape not followed by the terminator",
			b:    []byte("\x1b[200~before\x1bmiddle\x1b[201~"),
			want: "before\x1bmiddle",
		},
		{
			name: "content contains an escape almost matching the terminator",
			// "\x1b[201-" looks like the terminator until the last byte.
			b:    []byte("\x1b[200~before\x1b[201-after\x1b[201~"),
			want: "before\x1b[201-after",
		},
		{
			name: "trailing text after the paste is left in the reader",
			b:    []byte("\x1b[200~hello\x1b[201~ignored"),
			want: "hello",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// iotest.OneByteReader forces every Read to return at most one
			// byte, exercising the case where the terminator marker arrives
			// fragmented across multiple reads.
			content, err := ConsumeClipboardPaste(tc.b, iotest.OneByteReader(strings.NewReader(tc.rest)))
			if err != nil {
				t.Fatalf("err = %v, want nil", err)
			}
			if content != tc.want {
				t.Errorf("content = %q, want %q", content, tc.want)
			}
		})
	}
}

func TestConsumeClipboardPaste_ReaderError(t *testing.T) {
	for _, tc := range []struct {
		name string
		b    []byte
		rest string
	}{
		{"missing terminator entirely", []byte("\x1b[200~hello"), ""},
		{"cut off mid-marker", []byte("\x1b[200~hello\x1b[201"), ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			wantErr := errors.New("boom")
			r := iotest.ErrReader(wantErr)
			if tc.rest != "" {
				r = iotest.TimeoutReader(strings.NewReader(tc.rest))
			}
			content, err := ConsumeClipboardPaste(tc.b, r)
			if err == nil {
				t.Fatal("err = nil, want an error")
			}
			if content != "" {
				t.Errorf("content = %q, want empty on error", content)
			}
		})
	}
}

func TestCheckClipboardPaste(t *testing.T) {
	for _, tc := range []struct {
		name   string
		b      []byte
		marker string
		want   bool
	}{
		{"matches", []byte("\x1b[200~rest"), "200~", true},
		{"too short", []byte("\x1b[20"), "200~", false},
		{"wrong marker", []byte("\x1b[201~rest"), "200~", false},
		{"no escape prefix", []byte("x[200~rest"), "200~", false},
		{"no bracket", []byte("\x1bx200~rest"), "200~", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := checkClipboardPaste(tc.b, tc.marker); got != tc.want {
				t.Errorf("checkClipboardPaste(%q, %q) = %v, want %v", tc.b, tc.marker, got, tc.want)
			}
		})
	}
}
