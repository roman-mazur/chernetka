package extrefactor

import (
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/change"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/input"
)

// toolsExt provides the buffer data recording the tool calls.
type toolsExt struct{ calls []string }

type toolsData struct{ ext *toolsExt }

func (x *toolsExt) ID() string { return "tools" }
func (x *toolsExt) MakeBufferData(editor.Sender, *editor.Buffer) editor.BufferExtData {
	return &toolsData{ext: x}
}
func (x *toolsExt) AfterEdit(editor.Sender, *editor.Buffer, []change.Item) {}

func (d *toolsData) record(what string, p content.Position) {
	d.ext.calls = append(d.ext.calls, what+"@"+p.String())
}
func (d *toolsData) FindUsages(_ editor.Sender, p content.Position)  { d.record("usages", p) }
func (d *toolsData) FindCallers(_ editor.Sender, p content.Position) { d.record("callers", p) }
func (d *toolsData) Rename(_ editor.Sender, p content.Position)      { d.record("rename", p) }
func (d *toolsData) ShowCodeActions(_ editor.Sender, s content.Span) {
	d.ext.calls = append(d.ext.calls, "actions@"+s.Start.String()+"-"+s.End.String())
}

func TestIntegration_Keys(t *testing.T) {
	tools := new(toolsExt)
	var e editor.Editor
	e.Extend(tools)
	e.Extend(new(Integration))
	if err := e.OpenReader("a.go", strings.NewReader("package a\n\nfunc f() {}\n")); err != nil {
		t.Fatal(err)
	}
	buf := e.Top()
	ext := new(Integration)

	press := func(k input.Key) bool { return ext.HandleKey(&e, buf, k) }
	editor.RelMove{Dy: 2, Dx: 5}.DoOnBuffer(buf, editor.RenderPrefs{})
	for _, k := range []input.Key{KeyUsages, KeyCallers, KeyRename, KeyCodeActions} {
		if !press(k) {
			t.Errorf("%s not handled", k)
		}
	}
	for _, k := range []input.Key{input.Rune('u'), input.Of(input.Enter), input.Fn(7, 0), input.Fn(6, input.ModCtrl)} {
		if press(k) {
			t.Errorf("%s handled", k)
		}
	}

	// The selection is passed to the code actions.
	editor.StartTextSelection.DoOnBuffer(buf, editor.RenderPrefs{})
	editor.RelMove{Dy: -2}.DoOnBuffer(buf, editor.RenderPrefs{})
	editor.StopTextSelection.DoOnBuffer(buf, editor.RenderPrefs{})
	press(KeyCodeActions)

	want := []string{"usages@l=2/c=5", "callers@l=2/c=5", "rename@l=2/c=5", "actions@l=2/c=5-l=2/c=5", "actions@l=0/c=5-l=2/c=5"}
	if strings.Join(tools.calls, " ") != strings.Join(want, " ") {
		t.Errorf("calls %q, want %q", tools.calls, want)
	}

	// A buffer without the tools.
	var plain editor.Editor
	plain.Extend(ext)
	if err := plain.OpenReader("b.txt", strings.NewReader("text")); err != nil {
		t.Fatal(err)
	}
	if ext.HandleKey(&plain, plain.Top(), KeyUsages) {
		t.Error("usages handled without a usage finder")
	}
}
