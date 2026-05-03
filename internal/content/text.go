package content

import (
	"bufio"
	"errors"
	"io"
)

// TextLine is a text Line implementation.
type TextLine string

func (ti TextLine) String() string   { return string(ti) }
func (ti TextLine) MimeType() string { return "text/plain" }
func (ti TextLine) Len() int         { return len(ti) }

// FullText implements the Interface exposing a text file fully loaded in the memory.
type FullText []Line

// LoadFullText reads the input reader until the end using the system line delimiter.
// In the success scenario the input reader is fully consumed.
func LoadFullText(in io.Reader) (FullText, error) {
	r := bufio.NewReader(in)
	var (
		ft  FullText
		end bool
	)
	for !end {
		line, readError := r.ReadString('\n')
		end = errors.Is(readError, io.EOF)
		if readError != nil && !end {
			return nil, readError
		}
		line = dropR(line)
		ft = append(ft, TextLine(line))
	}
	return ft, nil
}

func (ft FullText) Len() int      { return len(ft) }
func (ft FullText) Lines() []Line { return ft }

// Empty returns an implementation of an empty content.
func Empty() Interface {
	return make(FullText, 0)
}

func dropR(line string) string {
	if len(line) > 0 && line[len(line)-1] == '\r' {
		return line[:len(line)-1]
	}
	return line
}
