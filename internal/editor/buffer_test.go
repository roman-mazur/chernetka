package editor

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"rmazur.io/chernetka/internal/content"
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

func (tse testSuggestionExt) TextSuggestion() string {
	if len(tse) == 0 {
		return ""
	}
	return tse[0]
}

func TestBuffer_AcceptSuggestion(t *testing.T) {
	buf := &Buffer{
		Content: &content.FullText{content.TextLine("Pri")},
		c:       content.Position{3, 0},
	}

	buf.AcceptSuggestion("ntln")
	if got := buf.Content.Lines()[0].String(); got != "Println" {
		t.Errorf("line = %q, want %q", got, "Println")
	}
	if buf.c.Col != 7 {
		t.Errorf("cx = %d, want 7", buf.c.Col)
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

// TestBuffer_Render_FillsExactlyWindowHeight pins the invariant that keeps the buffer
// visible in a short terminal pane: Render must emit exactly h rows, i.e. h-1 line
// endings plus the status bar with no trailing newline. Emitting more rows than the
// pane holds scrolls the top of the content out of view.
func TestBuffer_Render_FillsExactlyWindowHeight(t *testing.T) {
	for _, tc := range []struct {
		name  string
		lines int
		mode  Mode
		h     int
	}{
		{name: "content shorter than pane", lines: 3, h: 21},
		{name: "content longer than pane", lines: 1000, h: 21},
		{name: "single row pane", lines: 1000, h: 2},
		{name: "command mode", lines: 1000, mode: ModeCommand, h: 21},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var buf Buffer
			testContent := slices.Repeat(content.FullText{content.TextLine("test")}, tc.lines)
			buf.Content = &testContent
			buf.mode = tc.mode
			buf.cmdline = "w"
			buf.w, buf.h = 40, tc.h

			var out bytes.Buffer
			buf.Render(&out, &RenderPrefs{TabSize: 2})
			t.Log("\n" + out.String())

			if got, want := strings.Count(out.String(), "\r\n"), tc.h-1; got != want {
				t.Errorf("Render emitted %d line endings, want %d (window height %d)", got, want, tc.h)
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

func TestNormalInput_EngageAction(t *testing.T) {
	buf, actions := newActionsTestBuffer("one\ntwo\nthree", 1)
	prefs := newRenderPrefs()

	normalInput(buf, []byte{'\r'}, &prefs)
	if buf.c.Line != 1 {
		t.Fatalf("Enter on a plain line moved the cursor to %d, want 1", buf.c.Line)
	}

	normalInput(buf, []byte{'\r'}, &prefs)
	if got := actions[1].engaged; got != 1 {
		t.Errorf("action engaged %d times, want 1", got)
	}
	if buf.c.Line != 1 {
		t.Errorf("Enter on an actionable line moved the cursor to %d", buf.c.Line)
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

func (noopExt) ID() string                                           { return "actions" }
func (noopExt) AfterEdit(*Editor, *Buffer)                           {}
func (noopExt) HandleInsertInput(*Buffer, *RenderPrefs, []byte) bool { return false }

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
