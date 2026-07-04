package inputs

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestIsMouseInput(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []byte
		want  bool
	}{
		{"valid prefix", []byte{0x1b, '[', '<'}, true},
		{"valid prefix with more bytes", []byte{0x1b, '[', '<', '0'}, true},
		{"too short", []byte{0x1b, '['}, false},
		{"empty", []byte{}, false},
		{"wrong second byte", []byte{0x1b, 'O', '<'}, false},
		{"wrong third byte", []byte{0x1b, '[', 'A'}, false},
		{"no escape", []byte{'[', '<', '0'}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsMouseInput(tc.input); got != tc.want {
				t.Errorf("IsMouseInput(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestReadMouse(t *testing.T) {
	for _, tc := range []struct {
		name    string
		input   string
		want    Mouse
		wantErr string
		restTxt string
	}{
		{
			name:  "left button press",
			input: "\x1b[<0;10;20M",
			want:  Mouse{Button: MouseButtonLeft, X: 10, Y: 20, Pressed: true},
		},
		{
			name:  "left button release",
			input: "\x1b[<0;10;20m",
			want:  Mouse{Button: MouseButtonLeft, X: 10, Y: 20, Pressed: false},
		},
		{
			name:  "middle button press",
			input: "\x1b[<1;5;15M",
			want:  Mouse{Button: MouseButtonMiddle, X: 5, Y: 15, Pressed: true},
		},
		{
			name:  "right button release",
			input: "\x1b[<2;80;24m",
			want:  Mouse{Button: MouseButtonRight, X: 80, Y: 24, Pressed: false},
		},
		{
			name:  "large coordinates",
			input: "\x1b[<0;1920;1080M",
			want:  Mouse{Button: MouseButtonLeft, X: 1920, Y: 1080, Pressed: true},
		},
		{
			name:  "single-digit coordinates",
			input: "\x1b[<0;1;1M",
			want:  Mouse{Button: MouseButtonLeft, X: 1, Y: 1, Pressed: true},
		},
		{
			name:    "not a mouse input",
			input:   "\x1b[A",
			wantErr: "not a mouse input",
			restTxt: "\x1b[A",
		},
		{
			name:    "invalid terminator",
			input:   "\x1b[<0;1;2X",
			wantErr: "expected 'm' or 'M'",
		},
		{
			name:    "remaining text",
			input:   "\x1b[<0;1;2Mhello",
			want:    Mouse{Button: MouseButtonLeft, X: 1, Y: 2, Pressed: true},
			restTxt: "hello",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tc.input))

			data, err := ReadMouse(reader)
			if (err != nil) != (tc.wantErr != "") {
				t.Fatalf("ReadMouse() error = %q, wantErr %q", err, tc.wantErr)
			}
			if err != nil && !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("ReadMouse() error = %q, wantErr %q", err, tc.wantErr)
			}
			if err == nil && data != tc.want {
				t.Errorf("ReadMouse() = %v, want %v", data, tc.want)
			}

			remainingData, _ := io.ReadAll(reader)
			if string(remainingData) != tc.restTxt {
				t.Errorf("remaining text = %q, want %q", string(remainingData), tc.restTxt)
			}
		})
	}
}
