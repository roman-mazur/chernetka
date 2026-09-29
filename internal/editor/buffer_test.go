package editor

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/editor/styles"
	"rmazur.io/chernetka/internal/vt/escape"
)

var thousandLines = slices.Repeat(content.FullText{content.TextLine("test")}, 1000)

func TestBuffer_Render(t *testing.T) {
	cases := []struct {
		name        string
		buf         *Buffer
		suggestions []string
		prefs       RenderPrefs
		contains    []string
		absent      []string
	}{
		{
			name: "normal mode shows numbered lines",
			buf: &Buffer{
				Path: "test.txt",
				Content: &content.FullText{
					content.TextLine("hello"),
					content.TextLine("world"),
				},
				mode: ModeNormal,
				w:    40,
				h:    5,
			},
			prefs:    RenderPrefs{TabSize: 4},
			contains: []string{"1 hello", "2 world"},
		},
		{
			name: "hide line numbers",
			buf: &Buffer{
				Path: "test.txt",
				Content: &content.FullText{
					content.TextLine("hello"),
					content.TextLine("world"),
				},

				hideLineNumbers: true,
				mode:            ModeNormal,
				w:               40,
				h:               5,
			},
			prefs:    RenderPrefs{TabSize: 4},
			contains: []string{"hello", "world"},
			absent:   []string{"1 hello", "2 world"},
		},
		{
			name: "tab expands to TabSize spaces",
			buf: &Buffer{
				Path:    "test.txt",
				Content: &content.FullText{content.TextLine("a\tb")},
				mode:    ModeNormal,
				w:       40,
				h:       3,
			},
			prefs:    RenderPrefs{TabSize: 4},
			contains: []string{"1 a    b"},
		},
		{
			name: "inline suggestion rendered as ghost text after cursor",
			buf: &Buffer{
				Path:    "test.go",
				Content: &content.FullText{content.TextLine("Pri")},
				mode:    ModeInsert,
				c:       content.Position{3, 0},
				w:       40,
				h:       3,
			},
			suggestions: []string{"ntln"},
			prefs:       RenderPrefs{TabSize: 4},
			contains:    []string{"1 Println"},
		},
		{
			name: "2 digits line numbers",
			buf: &Buffer{
				Path:    "test.txt",
				Content: &thousandLines,
				w:       40,
				h:       10,
				offset:  5,
			},
			contains: []string{" 9 test", "10 test"},
		},
		{
			name: "3 digits line numbers",
			buf: &Buffer{
				Path:    "test.txt",
				Content: &thousandLines,
				w:       40,
				h:       10,
				offset:  95,
			},
			contains: []string{" 99 test", "100 test"},
		},
		{
			name: "4 digits line numbers",
			buf: &Buffer{
				Path:    "test.txt",
				Content: &thousandLines,
				w:       40,
				h:       10,
				offset:  995,
			},
			contains: []string{" 999 test", "1000 test"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.buf.ext.extend("lsp", testSuggestionExt(tc.suggestions))

			var out bytes.Buffer
			tc.buf.Render(&out, &tc.prefs)
			got := escape.Clean(out.String())

			for _, want := range tc.contains {
				if !strings.Contains(got, want) {
					t.Errorf("output missing %q; got:\n%s", want, got)
				}
			}
			for _, unwanted := range tc.absent {
				if strings.Contains(got, unwanted) {
					t.Errorf("output should not contain %q; got:\n%s", unwanted, got)
				}
			}
		})
	}

	t.Run("selections", func(t *testing.T) {
		for n, tc := range []struct {
			name       string
			content    string
			selections []content.Span
			render     []escape.Span
			selText    string

			enableCurrentHighlight bool
		}{
			{
				name:    "selection start",
				content: "line 1",
				selections: []content.Span{
					{content.Position{1, 0}, content.Position{1, 0}},
				},
			},
			{
				name:    "one line",
				content: "line 1",
				selections: []content.Span{
					{content.Position{1, 0}, content.Position{3, 0}},
				},
				selText: "in",
			},
			{
				name:    "one line reversed",
				content: "line 1",
				selections: []content.Span{
					{content.Position{3, 0}, content.Position{1, 0}},
				},
				render: []escape.Span{
					{escape.Position{Offset: 1}, escape.Position{Offset: 3}},
				},
				selText: "in",
			},
			{
				name:    "one line with hl",
				content: "line 1",
				selections: []content.Span{
					{content.Position{1, 0}, content.Position{4, 0}},
				},
				render: []escape.Span{
					{escape.Position{Line: 0, Offset: 0}, escape.Position{Line: 0, Offset: 1}},
					{escape.Position{Line: 0, Offset: 1}, escape.Position{Line: 0, Offset: 4}},
					{escape.Position{Line: 0, Offset: 4}, escape.Position{Line: 0, Offset: 6}},
					{escape.Position{Line: 0, Offset: 6}, escape.Position{Line: 0, Offset: 40}},
				},
				enableCurrentHighlight: true,
				selText:                "ine",
			},
			{
				name:    "multiple lines",
				content: "line 1\nline 2\nline 3",
				selections: []content.Span{
					{content.Position{3, 2}, content.Position{1, 0}},
				},
				render: []escape.Span{
					{escape.Position{Line: 0, Offset: 1}, escape.Position{Line: 0, Offset: 6}},
					{escape.Position{Line: 1, Offset: 0}, escape.Position{Line: 1, Offset: 6}},
					{escape.Position{Line: 2, Offset: 0}, escape.Position{Line: 2, Offset: 3}},
				},
				selText: "ine 1\nline 2\nlin",
			},
			{
				name:    "multiple lines reversed",
				content: "line 1\nline 2\nline 3",
				selections: []content.Span{
					{content.Position{1, 0}, content.Position{3, 2}},
				},
				render: []escape.Span{
					{escape.Position{Line: 0, Offset: 1}, escape.Position{Line: 0, Offset: 6}},
					{escape.Position{Line: 1, Offset: 0}, escape.Position{Line: 1, Offset: 6}},
					{escape.Position{Line: 2, Offset: 0}, escape.Position{Line: 2, Offset: 3}},
				},
				selText: "ine 1\nline 2\nlin",
			},
		} {
			t.Run(fmt.Sprintf("%d/%s", n, tc.name), func(t *testing.T) {
				var buf Buffer
				testContent, err := content.LoadFullText(strings.NewReader(tc.content))
				if err != nil {
					t.Fatal(err)
				}
				buf.Content = &testContent

				buf.sel = tc.selections
				buf.w = 40                    // not too long for the test output
				buf.h = testContent.Len() + 1 // print all test input
				buf.hideLineNumbers = true
				buf.noCurrentLineHL = !tc.enableCurrentHighlight

				var out bytes.Buffer
				buf.Render(&out, &RenderPrefs{TabSize: 2})
				sOut := out.String()
				t.Log("output:\n" + sOut)
				t.Log("raw:\n" + strings.ReplaceAll(sOut, "\x1b", "^"))

				actuals := escape.ScanPositions(sOut, "48;2;", "0")
				expected := tc.render
				if expected == nil {
					expected = make([]escape.Span, len(tc.selections))
					for i := range expected {
						expected[i] = makeSpan(tc.selections[i])
					}
				}

				if diff := cmp.Diff(expected, actuals); diff != "" {
					t.Errorf("rendering mismatch (-want +got):\n%s", diff)
				}

				if diff := cmp.Diff(tc.selText, buf.SelectedText()); diff != "" {
					t.Errorf("SelectedText() mismatch (-want +got):\n%s", diff)
				}
			})
		}
	})
}

func makePos(p content.Position) escape.Position { return escape.Position{Line: p.Line, Offset: p.Col} }
func makeSpan(s content.Span) escape.Span        { return escape.Span{makePos(s.Start), makePos(s.End)} }

type testSuggestionExt []string

func (tse testSuggestionExt) TextSuggestion() (s code.Suggestion) {
	if len(tse) == 0 {
		return
	}
	return code.Suggestion{Text: tse[0]}
}

func TestBuffer_AcceptSuggestion(t *testing.T) {
	buf := &Buffer{
		Content: &content.FullText{content.TextLine("Pri")},
		c:       content.Position{3, 0},
	}

	buf.AcceptSuggestion("ntln", 4)
	if got := buf.Content.Lines()[0].String(); got != "Println" {
		t.Errorf("line = %q, want %q", got, "Println")
	}
	if buf.c.Col != 7 {
		t.Errorf("cx = %d, want 7", buf.c.Col)
	}

	buf.AcceptSuggestion("()", 1)
	if got := buf.Content.Lines()[0].String(); got != "Println()" {
		t.Errorf("line = %q, want %q", got, "Println()")
	}
	if buf.c.Col != 8 {
		t.Errorf("cx = %d, want 8 (inside the brackets)", buf.c.Col)
	}
}

func TestBuffer_ReplaceText(t *testing.T) {
	pos := func(line, col int) content.Position { return content.Position{Col: col, Line: line} }
	cases := []struct {
		name       string
		lines      []string
		cursor     content.Position
		start, end content.Position
		text       string
		wantText   string
		wantCursor content.Position
	}{
		{
			name:  "insert import line above the cursor",
			lines: []string{"import (", "\t\"fmt\"", ")", "", "strings.Sp"},
			// gopls style: insert before the closing paren.
			cursor: pos(4, 10), start: pos(1, 6), end: pos(1, 6), text: "\n\t\"strings\"",
			wantText:   "import (\n\t\"fmt\"\n\t\"strings\"\n)\n\nstrings.Sp",
			wantCursor: pos(5, 10),
		},
		{
			name:   "collapse lines above the cursor",
			lines:  []string{"a", "b", "c", "x"},
			cursor: pos(3, 1), start: pos(0, 1), end: pos(2, 0), text: "",
			wantText:   "ac\nx",
			wantCursor: pos(1, 1),
		},
		{
			name:   "edit before the cursor on its line",
			lines:  []string{"foo(bar)"},
			cursor: pos(0, 7), start: pos(0, 0), end: pos(0, 3), text: "fmt.Println",
			wantText:   "fmt.Println(bar)",
			wantCursor: pos(0, 15),
		},
		{
			name:   "edit after the cursor",
			lines:  []string{"ab", "cd"},
			cursor: pos(0, 1), start: pos(1, 0), end: pos(1, 2), text: "x\ny",
			wantText:   "ab\nx\ny",
			wantCursor: pos(0, 1),
		},
		{
			name:   "cursor inside replaced range",
			lines:  []string{"hello world"},
			cursor: pos(0, 3), start: pos(0, 0), end: pos(0, 5), text: "bye",
			wantText:   "bye world",
			wantCursor: pos(0, 3),
		},
		{
			name:   "reversed span",
			lines:  []string{"hello world"},
			cursor: pos(0, 11), start: pos(0, 5), end: pos(0, 0), text: "bye",
			wantText:   "bye world",
			wantCursor: pos(0, 9),
		},
		{
			name:   "out of range is ignored",
			lines:  []string{"ab"},
			cursor: pos(0, 1), start: pos(0, 0), end: pos(1, 0), text: "zz",
			wantText:   "ab",
			wantCursor: pos(0, 1),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var text content.FullText
			for _, l := range tc.lines {
				text = append(text, content.TextLine(l))
			}
			buf := &Buffer{Content: &text, c: tc.cursor}
			buf.ReplaceText(content.Span{Start: tc.start, End: tc.end}, tc.text)
			if got := buf.Text(); got != tc.wantText {
				t.Errorf("text = %q, want %q", got, tc.wantText)
			}
			if buf.c != tc.wantCursor {
				t.Errorf("cursor = %s, want %s", buf.c, tc.wantCursor)
			}
		})
	}
}

func TestBuffer_Text(t *testing.T) {
	buf := &Buffer{
		Content: &content.FullText{content.TextLine("Hello")},
	}
	doubleCheckBufferText(t, buf, "Hello")
	buf.Mutate().Insert(1, content.TextLine("World"))
	doubleCheckBufferText(t, buf, "Hello\nWorld")
}

func TestBuffer_SelectedText_MultiLineDownAndLeft(t *testing.T) {
	// Selection starts at a high column on an early, long line and ends at
	// a low column on a later, short line (a normal down-and-left drag).
	// This exercises Span.Min()/Max() with Start.Line < End.Line and
	// Start.Col > End.Col, which must not mix columns across lines.
	buf := &Buffer{
		Content: &content.FullText{
			content.TextLine("0123456789ABCDEFGHIJ"), // line 0, len 20
			content.TextLine("middle"),               // line 1
			content.TextLine("012"),                  // line 2, len 3
		},
		sel: []content.Span{{
			Start: content.Position{Col: 15, Line: 0},
			End:   content.Position{Col: 2, Line: 2},
		}},
	}

	want := "FGHIJ\nmiddle\n01"
	if got := buf.SelectedText(); got != want {
		t.Errorf("SelectedText() = %q, want %q", got, want)
	}
}

func BenchmarkBuffer_Text(b *testing.B) {
	f, err := os.Open("buffer_test.go")
	if err != nil {
		b.Fatal(err)
	}
	b.Cleanup(func() {
		_ = f.Close()
	})
	data, err := io.ReadAll(f)
	if err != nil {
		b.Fatal(err)
	}

	cnt, err := content.LoadFullText(bytes.NewReader(data))
	if err != nil {
		b.Fatal(err)
	}
	buf := Buffer{Content: &cnt}

	for b.Loop() {
		doubleCheckBufferText(b, &buf, string(data))
	}
}

func doubleCheckBufferText[T TB](t T, buf *Buffer, want string) {
	t.Helper()
	for i := range 2 {
		if res := buf.Text(); res != want {
			t.Errorf("check %d: text = %q, want %q", i, res, want)
		}
	}
}

type TB interface {
	*testing.T | *testing.B

	Helper()
	Errorf(string, ...any)
}

// TestBuffer_Render_FillsExactlyHeight pins the invariant that keeps the buffer
// visible in a short terminal pane: Render must emit exactly h rows, leaving the rest
// of the pane to the status bar. Emitting more rows than the pane holds scrolls the
// top of the content out of view.
func TestBuffer_Render_FillsExactlyHeight(t *testing.T) {
	for _, tc := range []struct {
		name  string
		lines int
		mode  Mode
		h     int
	}{
		{name: "content shorter than pane", lines: 3, h: 21},
		{name: "content longer than pane", lines: 1000, h: 21},
		{name: "single row pane", lines: 1000, h: 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var buf Buffer
			testContent := slices.Repeat(content.FullText{content.TextLine("test")}, tc.lines)
			buf.Content = &testContent
			buf.mode = tc.mode
			buf.w, buf.h = 40, tc.h

			var out bytes.Buffer
			buf.Render(&out, &RenderPrefs{TabSize: 2})
			t.Log("\n" + out.String())

			if got, want := strings.Count(out.String(), "\r\n"), tc.h; got != want {
				t.Errorf("Render emitted %d line endings, want %d", got, want)
			}
		})
	}
}

func TestBuffer_ScreenToContentPosition(t *testing.T) {
	const tabSize = 2
	for i, tc := range []struct {
		x, y      int
		pos       content.Position
		content   string
		bufOffset int
	}{
		{
			content: "hello world",
		},
		{
			content:   "hello\nworld",
			bufOffset: 1,
			pos:       content.Position{Line: 1},
		},
		{
			content:   "hello\nworld",
			bufOffset: 1,
			x:         4,
			y:         5,
			pos:       content.Position{Line: 1, Col: 2},
		},
		{
			content: "hello\nworld",
			x:       0,
			y:       5,
			pos:     content.Position{Line: 1},
		},
		{
			content: "hello\nworld",
			x:       5,
			y:       0,
			pos:     content.Position{Col: 3},
		},
		{
			content: "hello\nworld",
			x:       100,
			y:       0,
			pos:     content.Position{Col: 5},
		},
		{
			content: "\t\t\thello",
			x:       2,
			pos:     content.Position{},
		},
		{
			content: "\t\t\thello",
			x:       3,
			pos:     content.Position{Col: 1},
		},
		{
			content: "\t\t\thello",
			x:       2 + tabSize,
			pos:     content.Position{Col: 1},
		},
		{
			content: "\t\t\thello",
			x:       2 + tabSize + 1,
			pos:     content.Position{Col: 2},
		},
		{
			content: "\t\t\thello",
			x:       2 + 3*tabSize + 2,
			pos:     content.Position{Col: 5},
		},
		{
			content: "привіт",
			x:       5,
			pos:     content.Position{Col: 6},
		},
	} {
		t.Run(fmt.Sprintf("%d/x=%d/y=%d/pos/%s", i, tc.x, tc.y, tc.pos), func(t *testing.T) {
			data, err := content.LoadFullText(strings.NewReader(tc.content))
			if err != nil {
				t.Fatal(err)
			}
			buf := Buffer{
				Content: &data,
				w:       40,
				h:       10,
				offset:  tc.bufOffset,
			}
			if res := buf.screenToContentPosition(tc.y, tc.x, tabSize); res != tc.pos {
				t.Errorf("screenToContentPosition(%d, %d, %d) = %s, want %s", tc.y, tc.x, 2, res, tc.pos)
			}
		})
	}
}

// testActionsExt is extension data that provides a counting action for the listed lines.
type testActionsExt map[int]*countingAction

func (tae testActionsExt) LineAction(lineNumber int) content.LineAction {
	if action, ok := tae[lineNumber]; ok {
		return action
	}
	return nil
}

type countingAction struct{ engaged int }

func (ca *countingAction) Engage() { ca.engaged++ }

func newActionsTestBuffer(text string, actionLines ...int) (*Buffer, testActionsExt) {
	data, err := content.LoadFullText(strings.NewReader(text))
	if err != nil {
		panic(err)
	}
	actions := make(testActionsExt)
	for _, ln := range actionLines {
		actions[ln] = new(countingAction)
	}
	buf := &Buffer{
		Content: &data,
		w:       20,
		h:       data.Len() + 1,
	}
	buf.ext.extend("actions", actions)
	return buf, actions
}

func TestBuffer_Render_ActionMarker(t *testing.T) {
	for _, tc := range []struct {
		name        string
		actionLines []int
		cursorLine  int
		offset      int
		hide        bool
		wantMarked  []bool // for every rendered row
	}{
		{
			name:       "no actions",
			wantMarked: []bool{false, false, false},
		},
		{
			name:        "marked lines",
			actionLines: []int{0, 2},
			cursorLine:  1,
			wantMarked:  []bool{true, false, true},
		},
		{
			name:        "marked current line",
			actionLines: []int{1},
			cursorLine:  1,
			wantMarked:  []bool{false, true, false},
		},
		{
			// The marker is positioned on the screen row, not on the content line.
			name:        "scrolled",
			actionLines: []int{2},
			offset:      1,
			wantMarked:  []bool{false, true},
		},
		{
			name:        "hidden markers",
			actionLines: []int{0, 1, 2},
			hide:        true,
			wantMarked:  []bool{false, false, false},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			buf, _ := newActionsTestBuffer("first\nsecond line that does not fit the window\nthird", tc.actionLines...)
			buf.hideLineActions = tc.hide
			buf.c.Line = tc.cursorLine
			buf.offset = tc.offset

			var out bytes.Buffer
			buf.Render(&out, &RenderPrefs{TabSize: 2})
			t.Log("\n" + out.String())

			rows := strings.Split(out.String(), "\r\n")
			for i, wantMarked := range tc.wantMarked {
				var cursorPosSeq strings.Builder
				escape.SetCursorPosition(&cursorPosSeq, i+1, buf.w)
				// The marker is placed in the last window column regardless of the line length.
				marked := strings.HasSuffix(escape.Clean(rows[i]), actionMarker) &&
					strings.Contains(rows[i], cursorPosSeq.String())
				if marked != wantMarked {
					t.Errorf("line %d %q marked = %t, want %t", i, escape.Clean(rows[i]), marked, wantMarked)
				}
			}
		})
	}
}

func TestEditor_OpenDir_HidesActionMarkers(t *testing.T) {
	var e Editor
	e.OpenDir(t.TempDir(), nil)
	if !e.Top().hideLineActions {
		t.Error("directory buffer is configured to mark actionable lines")
	}
}

// actionsExt is an Extension providing actions for the listed lines of every buffer.
type actionsExt struct {
	noopExt
	actions testActionsExt
}

func (ae *actionsExt) MakeBufferData(*Buffer) BufferExtData { return ae.actions }

type noopExt struct{}

func (noopExt) ID() string                { return "actions" }
func (noopExt) AfterEdit(Sender, *Buffer) {}

func TestEditor_RerunAction(t *testing.T) {
	const ctrlR = 0x12
	first, second := new(countingAction), new(countingAction)
	h := NewTestHarness()
	h.Extend(&actionsExt{actions: testActionsExt{0: first, 2: second}})
	if err := h.OpenReader("a.txt", strings.NewReader("action\ntext\naction")); err != nil {
		t.Fatal(err)
	}
	h.Run(t)

	check := func(wantFirst, wantSecond int) {
		t.Helper()
		// Commands are drained by SendInput: the counters are not modified concurrently.
		if first.engaged != wantFirst || second.engaged != wantSecond {
			t.Errorf("engaged %d and %d times, want %d and %d",
				first.engaged, second.engaged, wantFirst, wantSecond)
		}
	}

	h.SendInput(t, []byte{ctrlR})
	check(0, 0) // Nothing to re-run yet.

	h.SendInput(t, []byte{'\r'})
	check(1, 0)

	h.SendInputSequence(t, "jj")
	h.SendInput(t, []byte{ctrlR})
	check(2, 0) // The cursor position does not matter.

	h.SendInput(t, []byte{'\r'})
	h.SendInput(t, []byte{ctrlR})
	check(2, 2)

	h.SendInput(t, []byte{'i'})
	h.SendInput(t, []byte{ctrlR})
	check(2, 3) // Works in insert mode too.

	h.Post(t, CommandFunc(func(e *Editor) {
		if text := e.Top().Text(); text != "action\ntext\naction" {
			t.Errorf("Ctrl+R changed the text: %q", text)
		}
	}))
}

// singleShotAction is a countingAction that is not re-run.
type singleShotAction struct{ countingAction }

func (*singleShotAction) SingleShot() {}

// singleShotExt provides a countingAction on the first line and a singleShotAction on the second.
type singleShotExt struct {
	noopExt
	rerun  *countingAction
	single *singleShotAction
}

func (se *singleShotExt) MakeBufferData(*Buffer) BufferExtData { return se }

func (se *singleShotExt) LineAction(lineNumber int) content.LineAction {
	switch lineNumber {
	case 0:
		return se.rerun
	case 1:
		return se.single
	}
	return nil
}

func TestEditor_RerunAction_SkipsSingleShot(t *testing.T) {
	const ctrlR = 0x12
	rerun, single := new(countingAction), new(singleShotAction)
	h := NewTestHarness()
	h.Extend(&singleShotExt{rerun: rerun, single: single})
	if err := h.OpenReader("a.txt", strings.NewReader("action\nsingle")); err != nil {
		t.Fatal(err)
	}
	h.Run(t)

	h.SendInput(t, []byte{'\r'}) // Engage the first action, the cursor stays.
	h.SendInputSequence(t, "j")
	h.SendInput(t, []byte{'\r'})
	h.SendInput(t, []byte{ctrlR})
	if rerun.engaged != 2 || single.engaged != 1 {
		t.Errorf("engaged %d and %d times, want 2 and 1", rerun.engaged, single.engaged)
	}
}

func TestEditor_EngageSaves(t *testing.T) {
	path := filepath.Join(t.TempDir(), "a.txt")
	if err := os.WriteFile(path, []byte("a\nb"), 0o600); err != nil {
		t.Fatal(err)
	}
	action := new(countingAction)
	e := new(Editor)
	e.Extend(&actionsExt{actions: testActionsExt{1: action}})
	(&OpenFile{Path: path}).DoOnEditor(e)

	saved := func() string {
		data, _ := os.ReadFile(path)
		return string(data)
	}
	for _, k := range []string{"x", "j", "\r"} {
		e.handleInput([]byte(k))
	}
	if action.engaged != 1 || saved() != "\nb" {
		t.Errorf("action engaged %d times with the file %q, want once with the changes saved", action.engaged, saved())
	}
	e.handleInput([]byte("x"))
	e.handleInput([]byte{0x12}) // Ctrl+R
	if action.engaged != 2 || saved() != "\n" {
		t.Errorf("re-run action engaged %d times with the file %q, want twice with the changes saved", action.engaged, saved())
	}
}

// ghostAssist is a CodeAssist with a syntax highlighter and suggestion info.
type ghostAssist struct {
	suggestion, info string
	spans            []SyntaxSpan
}

func (g ghostAssist) TextSuggestion() code.Suggestion {
	return code.Suggestion{Text: g.suggestion, Info: g.info}
}
func (g ghostAssist) SyntaxSpans(int, string) []SyntaxSpan { return g.spans }

func TestBuffer_RenderGhostKeepsSyntaxHighlight(t *testing.T) {
	keyword := SyntaxSpan{Start: 0, End: 3, TokenType: code.TtKeyword}
	ext := ghostAssist{suggestion: "intln", info: "func(a ...any)", spans: []SyntaxSpan{keyword}}
	buf := &Buffer{
		Path:    "test.go",
		Content: &content.FullText{content.TextLine("fmt.Pr()")},
		mode:    ModeInsert,
		c:       content.Position{Col: 6},
		w:       40,
		h:       3,

		noCurrentLineHL: true,
	}
	buf.ext.extend("lsp", ext)

	var out bytes.Buffer
	buf.Render(&out, &RenderPrefs{TabSize: 4})

	var styled bytes.Buffer
	escape.StyleText(&styled, "fmt", styles.ResolveTokenStyle(code.TtKeyword))
	if !strings.Contains(out.String(), styled.String()) {
		t.Errorf("syntax highlight is lost with a suggestion shown:\n%q", out.String())
	}
	got := escape.Clean(out.String())
	if !strings.Contains(got, "1 fmt.Println()") {
		t.Errorf("ghost text missing:\n%s", got)
	}
	if !strings.Contains(got, "func(a ...any)") {
		t.Errorf("suggestion info missing:\n%s", got)
	}

	buf.w = 20 // Too narrow for the info.
	out.Reset()
	buf.Render(&out, &RenderPrefs{TabSize: 4})
	if got := escape.Clean(out.String()); strings.Contains(got, "func(") {
		t.Errorf("suggestion info shown on a narrow screen:\n%s", got)
	}
}

func TestBuffer_Render_HorizontalScroll(t *testing.T) {
	for _, tc := range []struct {
		name    string
		text    string
		xoff    int
		w       int
		mode    Mode
		c       content.Position
		sel     []content.Span
		sug     []string
		hlLine  bool
		want    string
		tabSize int
	}{
		{name: "not scrolled", text: "0123456789abcdef", w: 10, want: "0123456789"},
		{name: "scrolled", text: "0123456789abcdef", xoff: 4, w: 10, want: "456789abcd"},
		{name: "scrolled to the end", text: "0123456789abcdef", xoff: 10, w: 10, want: "abcdef"},
		{name: "line is scrolled away", text: "0123", xoff: 10, w: 10, want: ""},
		{name: "tab crosses the left edge", text: "a\tbcdef", xoff: 2, w: 4, tabSize: 4, want: "   b"},
		{name: "tab crosses the right edge", text: "ab\tc", xoff: 1, w: 4, tabSize: 4, want: "b   "},
		{name: "utf-8", text: "привіт світ", xoff: 3, w: 5, want: "віт с"},
		{
			name: "selection and current line", text: "0123456789abcdef", xoff: 4, w: 10, hlLine: true,
			sel:  []content.Span{{Start: content.Position{Col: 2}, End: content.Position{Col: 6}}},
			want: "456789abcd",
		},
		{name: "current line tail", text: "abc", xoff: 4, w: 10, hlLine: true, want: strings.Repeat(" ", 10)},
		{
			name: "ghost suggestion", text: "Pri", xoff: 2, w: 4, mode: ModeInsert,
			c: content.Position{Col: 3}, sug: []string{"ntln"}, want: "intl",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			buf := &Buffer{
				Content:         &content.FullText{content.TextLine(tc.text)},
				w:               tc.w,
				h:               1,
				xoff:            tc.xoff,
				mode:            tc.mode,
				c:               tc.c,
				sel:             tc.sel,
				hideLineNumbers: true,
				noCurrentLineHL: !tc.hlLine,
			}
			buf.ext.extend("lsp", testSuggestionExt(tc.sug))

			var out bytes.Buffer
			buf.Render(&out, &RenderPrefs{TabSize: max(tc.tabSize, 2)})
			got := strings.Split(escape.Clean(out.String()), "\r\n")[0]
			if got != tc.want {
				t.Errorf("rendered %q, want %q", got, tc.want)
			}
		})
	}
}

func TestBuffer_ClampCursor_FollowsHorizontally(t *testing.T) {
	buf := &Buffer{
		Content: &content.FullText{
			content.TextLine(strings.Repeat("x", 100)),
			content.TextLine("short"),
		},
		w: 20, h: 5,
		hideLineNumbers: true,
	}
	const margin = 20 / 3
	for _, step := range []struct {
		name       string
		c          content.Position
		noKeyboard bool
		want       int
	}{
		{name: "line end", c: content.Position{Col: 100}, want: 100 - 20 + margin + 1},
		{name: "line start", c: content.Position{Col: 0}, want: 0},
		{name: "visible without scroll", c: content.Position{Col: 20 - margin - 1}, want: 0},
		{name: "past the right margin", c: content.Position{Col: 50}, want: 50 - 20 + margin + 1},
		{name: "within the view", c: content.Position{Col: 45}, want: 50 - 20 + margin + 1},
		{name: "past the left margin", c: content.Position{Col: 40}, want: 40 - margin},
		{name: "mouse input keeps the scroll", c: content.Position{Col: 0}, noKeyboard: true, want: 40 - margin},
		{name: "short line", c: content.Position{Line: 1, Col: 40}, want: 0},
	} {
		buf.c, buf.noKeyboard = step.c, step.noKeyboard
		buf.clampCursor(4)
		if buf.xoff != step.want {
			t.Errorf("%s: xoff = %d, want %d", step.name, buf.xoff, step.want)
		}
	}
}

func TestBuffer_HorizontalScroll_Coordinates(t *testing.T) {
	const tabSize = 4
	for _, tc := range []struct {
		text string
		xoff int
		x    int
		want int
	}{
		{text: "0123456789", xoff: 5, x: 2, want: 5},
		{text: "0123456789", xoff: 5, x: 4, want: 7},
		{text: "\t\tab", xoff: 6, x: 2, want: 2}, // inside the tab: the next rune
		{text: "\t\tab", xoff: 6, x: 5, want: 3},
	} {
		buf := &Buffer{Content: &content.FullText{content.TextLine(tc.text)}, w: 20, h: 5, xoff: tc.xoff}
		if got := buf.screenToContentPosition(0, tc.x, tabSize); got.Col != tc.want {
			t.Errorf("%q scrolled by %d: click at %d = %s, want col %d", tc.text, tc.xoff, tc.x, got, tc.want)
		}

		// The cursor goes back to the clicked column.
		buf.c.Col = tc.want
		var out bytes.Buffer
		buf.RenderCursorPosition(&out, &RenderPrefs{TabSize: tabSize})
		wantCol := runeToScreenCol(tc.text, tc.want, tabSize) - tc.xoff + buf.lineNumberPrefixWidth() + 1
		if want := fmt.Sprintf("\x1b[1;%dH", wantCol); out.String() != want {
			t.Errorf("%q scrolled by %d: cursor position %q, want %q", tc.text, tc.xoff, out.String(), want)
		}
	}
}
