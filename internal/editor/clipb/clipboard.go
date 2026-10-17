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

// Clipboard stores the copied text.
type Clipboard interface {
	Write(s string)
	Read() string
}

// InMemory keeps the text in the memory of the process.
type InMemory struct {
	buf string
}

func (c *InMemory) Write(s string) { c.buf = s }
func (c *InMemory) Read() string   { return c.buf }

// System uses the system clipboard, with a fallback to the memory when it's not available.
type System struct {
	fallback InMemory
}

func (c *System) Write(s string) {
	if systemSupported {
		_, err := clipboard.Write(context.Background(), clipboard.FmtText, []byte(s))
		if err == nil {
			return
		}
	}
	c.fallback.Write(s)
}

func (c *System) Read() string {
	if systemSupported {
		res, err := clipboard.Read(context.Background(), clipboard.FmtText)
		if err == nil {
			return string(res)
		}
	}
	return c.fallback.Read()
}
