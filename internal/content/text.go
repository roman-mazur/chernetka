package content

import (
	"bufio"
	"errors"
	"io"
	"slices"
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
		line = dropRN(line)
		ft = append(ft, TextLine(line))
	}
	return ft, nil
}

func (ft *FullText) Len() int      { return len(*ft) }
func (ft *FullText) Lines() []Line { return *ft }

func (ft *FullText) Insert(pos int, line Line) {
	*ft = slices.Insert(*ft, pos, line)
}
func (ft *FullText) Update(pos int, line Line) {
	(*ft)[pos] = line
}
func (ft *FullText) Delete(pos int) {
	*ft = slices.Delete(*ft, pos, pos+1)
}

// Empty returns an implementation of an empty content.
func Empty() Interface {
	ft := FullText{TextLine("")}
	return &ft
}

func dropRN(line string) string {
	if len(line) > 0 && line[len(line)-1] == '\n' {
		line = line[:len(line)-1]
	}
	if len(line) > 0 && line[len(line)-1] == '\r' {
		line = line[:len(line)-1]
	}
	return line
}
