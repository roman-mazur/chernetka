package clipb

import (
	"context"

	"golang.design/x/clipboard"
)

var systemSupported = false

func init() {
	err := clipboard.Init()
	if err == nil {
		systemSupported = true
	}
}

type Clipboard struct {
	buf string
}

func (c *Clipboard) Write(s string) {
	fallback := true
	if systemSupported {
		fallback = false
		_, err := clipboard.Write(context.Background(), clipboard.FmtText, []byte(s))
		if err != nil {
			fallback = true
		}
	}
	if fallback {
		c.buf = s
	}
}

func (c *Clipboard) Read() string {
	if systemSupported {
		res, err := clipboard.Read(context.Background(), clipboard.FmtText)
		if err != nil {
			return c.buf
		}
		return string(res)
	}
	return c.buf
}
