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
func (ti TextLine) MimeType() string { return MimeTypeTextPlain }
func (ti TextLine) Len() int         { return len(ti) }

func (ti TextLine) copy() Line { return ti }

const MimeTypeTextPlain = "text/plain"

// FullText implements the Document exposing a text file fully loaded in the memory.
type FullText []Line

// LoadFullText reads the input reader until the end using the system line delimiter.
// In the success scenario the input reader is fully consumed.
func LoadFullText(in io.Reader) (doc FullText, fingerprint string, err error) {
	hasher := newFingerprintHasher()
	r := bufio.NewReader(io.TeeReader(in, hasher))
	var end bool
	for !end {
		line, readError := r.ReadString('\n')
		end = errors.Is(readError, io.EOF)
		if readError != nil && !end {
			return nil, "", readError
		}
		line = dropRN(line)
		doc = append(doc, TextLine(line))
	}
	fingerprint = encodeFingerprint(hasher.Sum(nil))
	return
}

func (ft *FullText) Len() int      { return len(*ft) }
func (ft *FullText) Lines() []Line { return *ft }

func (ft *FullText) copy() Document {
	res := slices.Clone(*ft)
	return &res
}

func (ft *FullText) Insert(pos int, lines ...Line) {
	*ft = slices.Insert(*ft, pos, lines...)
}
func (ft *FullText) Update(pos int, line Line) {
	(*ft)[pos] = line
}
func (ft *FullText) Delete(startPos, endPos int) {
	*ft = slices.Delete(*ft, startPos, endPos)
}

// Empty returns an implementation of an empty content.
func Empty() Document {
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
