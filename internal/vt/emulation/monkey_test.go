//go:build darwin || linux

package emulation

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
	"time"

	"golang.org/x/term"
)

// echoEnv makes the test binary run echoApp instead of the tests.
const echoEnv = "EMULATION_ECHO_APP"

func TestMain(m *testing.M) {
	if os.Getenv(echoEnv) == "1" {
		echoApp()
		return
	}
	os.Exit(m.Run())
}

// echoApp writes its input back to the terminal in raw mode until it reads a NUL byte.
func echoApp() {
	if _, err := term.MakeRaw(0); err != nil {
		panic(err)
	}
	_, _ = os.Stdout.WriteString("ready")
	buf := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			panic(err)
		}
		if i := bytes.IndexByte(buf[:n], 0); i >= 0 {
			os.Exit(0)
		}
		_, _ = os.Stdout.Write(buf[:n])
	}
}

func TestMonkey(t *testing.T) {
	for _, tc := range []struct {
		name     string
		data     []byte
		duration time.Duration
	}{
		{name: "until the choices are exhausted", data: randomBytes(1, 4<<10)},
		{name: "for the duration", data: randomBytes(2, 1<<20), duration: time.Second},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command(os.Args[0])
			cmd.Env = append(os.Environ(), echoEnv+"=1")
			c := NewChoices(tc.data)
			m := Monkey{
				Cmd:      cmd,
				Choices:  c,
				Duration: tc.duration,
				Snippets: []string{"hello"},
				Ban:      Ban{Runes: "q"},
				Quit:     "\x00",
			}
			started := time.Now()
			m.Run(t)
			if m.count < 10 {
				t.Errorf("typed %d inputs", m.count)
			}
			if tc.duration == 0 && !c.Exhausted() {
				t.Error("choices are not exhausted")
			}
			if tc.duration > 0 && c.Exhausted() {
				t.Errorf("choices are exhausted in %s", time.Since(started))
			}
		})
	}
}
