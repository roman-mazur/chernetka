// Package extd2 implements an editor extension that makes d2 diagrams actionable.
//
// The first line of a .d2 file and the line opening a ```d2 code block in a Markdown file
// can be engaged to show the diagram with a Viewer. The diagram source is taken from the buffer,
// so unsaved changes are visible too.
//
// Markdown code blocks are found by the syntax highlighter extension, which has to be enabled as well.
package extd2

import (
	"path/filepath"
	"slices"
	"strings"

	"rmazur.io/chernetka/internal/cheimg"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/logger"
)

// Viewer shows diagrams, usually in the che-img process.
type Viewer interface {
	// Show displays the diagram. It's called by the editor and must not block.
	Show(d cheimg.Item)
}

// Integration implements an editor.Extension that provides line actions for the d2 diagrams.
type Integration struct {
	logger.LogEmbed
	Viewer Viewer
}

func (in *Integration) ID() string { return "d2" }

func (in *Integration) MakeBufferData(buf *editor.Buffer) editor.BufferExtData {
	if in.Viewer == nil || !holdsPlainText(buf) {
		return nil
	}
	switch strings.ToLower(filepath.Ext(buf.Path)) {
	case ".d2":
		return &diagramFile{Integration: in, buf: buf}
	case ".md", ".markdown":
		return &markdown{Integration: in, buf: buf}
	default:
		return nil
	}
}

func (in *Integration) AfterEdit(_ *editor.Editor, buf *editor.Buffer) {
	if md, ok := buf.ExtensionData(in.ID()).(*markdown); ok {
		md.sync()
	}
}

func (in *Integration) HandleInsertInput(*editor.Buffer, *editor.RenderPrefs, []byte) (handled bool) {
	return false
}

func (in *Integration) show(d cheimg.Item) {
	if abs, err := filepath.Abs(d.Path); err == nil {
		// The viewer runs in another process that may have a different working directory.
		d.Path = abs
	}
	in.Logf("show %s", d.Name())
	in.Viewer.Show(d)
}

// holdsPlainText reports whether the buffer holds ordinary text lines rather than, for example,
// a directory listing named after a directory with a matching extension.
func holdsPlainText(buf *editor.Buffer) bool {
	lines := buf.Content.Lines()
	return len(lines) == 0 || lines[0].MimeType() == content.MimeTypeTextPlain
}

// diagramFile makes the first line of a d2 file actionable.
type diagramFile struct {
	*Integration
	buf *editor.Buffer
}

func (df *diagramFile) LineAction(lineNumber int) content.LineAction {
	if lineNumber != 0 {
		return nil
	}
	return content.LineActionFunc(func() {
		df.show(cheimg.Item{Path: df.buf.Path, Source: df.buf.Text()})
	})
}

// markdown makes the lines opening d2 code blocks actionable.
// The blocks are found by another extension providing code.Blocks for the buffer: the syntax highlighter.
type markdown struct {
	*Integration
	buf *editor.Buffer

	synced bool
	blocks []code.Block // d2 blocks on the last sync
	lines  []string     // content on the last sync

	// followed is the last engaged action. It keeps pointing to its block while the content is edited,
	// so the diagram can be shown again after the lines above it are changed.
	followed *blockAction
}

// sync takes the d2 blocks of the current content and moves the followed action to its block.
// It's called after every edit, so the content differs from the last sync in one contiguous range.
func (md *markdown) sync() {
	lines := md.buf.Content.Lines()
	cur := make([]string, len(lines))
	for i, line := range lines {
		cur[i] = line.String()
	}

	var blocks []code.Block
	if provider, ok := editor.FindExtData[code.Blocks](md.buf); ok {
		for _, block := range provider.CodeBlocks() {
			if block.Lang == "d2" {
				blocks = append(blocks, block)
			}
		}
	}
	md.blocks = blocks

	if f := md.followed; f != nil && !f.lost {
		line, ok := followLine(f.line, md.lines, cur)
		f.line, f.lost = line, !ok || !md.opensBlock(line)
	}
	md.lines = cur
	md.synced = true
}

// ensureSynced performs the first sync lazily: when the buffer data is created,
// the extension providing the blocks may not have its data yet.
func (md *markdown) ensureSynced() {
	if !md.synced {
		md.sync()
	}
}

func (md *markdown) opensBlock(lineNumber int) bool {
	_, found := slices.BinarySearchFunc(md.blocks, lineNumber, compareOpenLine)
	return found
}

func (md *markdown) LineAction(lineNumber int) content.LineAction {
	md.ensureSynced()
	if !md.opensBlock(lineNumber) {
		return nil
	}
	return &blockAction{md: md, line: lineNumber}
}

// blockAction shows the diagram of the block opened on the line.
type blockAction struct {
	md   *markdown
	line int
	lost bool // the block was removed
}

func (ba *blockAction) Engage() {
	md := ba.md
	md.ensureSynced()
	if md.followed != ba {
		// A new action points to a line of the current content.
		md.followed = ba
		ba.lost = !md.opensBlock(ba.line)
	}
	if ba.lost {
		md.Logf("d2 block is not found anymore")
		return
	}
	md.showBlock(ba.line)
}

func (md *markdown) showBlock(lineNumber int) {
	i, found := slices.BinarySearchFunc(md.blocks, lineNumber, compareOpenLine)
	if !found {
		return
	}
	start, end := md.blocks[i].ContentLines()
	md.show(cheimg.Item{Path: md.buf.Path, Source: strings.Join(md.lines[start:end], "\n")})
}

func compareOpenLine(b code.Block, lineNumber int) int { return b.Start.Line - lineNumber }

// followLine finds where the line of the previous content revision is in the current one.
// The revisions are assumed to differ in one contiguous range of lines, as after a single edit.
// The lines above the edit keep their position, the lines below it are shifted.
// A line within the edited range keeps its position as long as the range is not shrunk below it.
func followLine(line int, prev, cur []string) (int, bool) {
	n := min(len(prev), len(cur))
	start := 0
	for start < n && prev[start] == cur[start] {
		start++
	}
	tail := 0
	for tail < n-start && prev[len(prev)-1-tail] == cur[len(cur)-1-tail] {
		tail++
	}

	switch {
	case line < start:
		return line, true
	case line >= len(prev)-tail:
		return line + len(cur) - len(prev), true
	default:
		return line, line < len(cur)-tail
	}
}
