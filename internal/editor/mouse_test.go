package editor

import (
	"testing"
	"time"

	"rmazur.io/chernetka/internal/editor/inputs"
)

func TestMouseHandlerTransformInput(t *testing.T) {
	type step struct {
		advance time.Duration
		in      inputs.Mouse
		want    mouseEventType
	}

	const (
		withinDelay = 100 * time.Millisecond
		pastDelay   = 600 * time.Millisecond
	)

	press := func(btn inputs.MouseButton) inputs.Mouse {
		return inputs.Mouse{Button: btn, Pressed: true}
	}
	release := func(btn inputs.MouseButton) inputs.Mouse {
		return inputs.Mouse{Button: btn}
	}
	drag := func(btn inputs.MouseButton) inputs.Mouse {
		return inputs.Mouse{Button: btn, Pressed: true, Mod: inputs.MouseModifier(8)}
	}
	hover := func() inputs.Mouse {
		return inputs.Mouse{Button: inputs.MouseButtonNone, Mod: inputs.MouseModifier(8)}
	}

	left := inputs.MouseButtonLeft
	right := inputs.MouseButtonRight

	for _, tc := range []struct {
		name  string
		steps []step
	}{
		{
			name: "single click",
			steps: []step{
				{in: press(left), want: mouseEventTypeRaw},
				{in: release(left), want: mouseEventTypeClick},
			},
		},
		{
			name: "double click",
			steps: []step{
				{in: press(left), want: mouseEventTypeRaw},
				{in: release(left), want: mouseEventTypeClick},
				{advance: withinDelay, in: press(left), want: mouseEventTypeRaw},
				{in: release(left), want: mouseEventTypeDoubleClick},
			},
		},
		{
			name: "triple click",
			steps: []step{
				{in: press(left), want: mouseEventTypeRaw},
				{in: release(left), want: mouseEventTypeClick},
				{advance: withinDelay, in: press(left), want: mouseEventTypeRaw},
				{in: release(left), want: mouseEventTypeDoubleClick},
				{advance: withinDelay, in: press(left), want: mouseEventTypeRaw},
				{in: release(left), want: mouseEventTypeTripleClick},
			},
		},
		{
			name: "click after timeout is single",
			steps: []step{
				{in: press(left), want: mouseEventTypeRaw},
				{in: release(left), want: mouseEventTypeClick},
				{advance: pastDelay, in: press(left), want: mouseEventTypeRaw},
				{in: release(left), want: mouseEventTypeClick},
			},
		},
		{
			name: "different button resets click count",
			steps: []step{
				{in: press(left), want: mouseEventTypeRaw},
				{in: release(left), want: mouseEventTypeClick},
				{advance: withinDelay, in: press(right), want: mouseEventTypeRaw},
				{in: release(right), want: mouseEventTypeClick},
			},
		},
		{
			name: "drag and release",
			steps: []step{
				{in: press(left), want: mouseEventTypeRaw},
				{in: drag(left), want: mouseEventTypeDragStart},
				{in: drag(left), want: mouseEventTypeRaw},
				{in: release(left), want: mouseEventTypeDragEnd},
			},
		},
		{
			name: "drag resets click count",
			steps: []step{
				{in: press(left), want: mouseEventTypeRaw},
				{in: drag(left), want: mouseEventTypeDragStart},
				{in: release(left), want: mouseEventTypeDragEnd},
				{advance: withinDelay, in: press(left), want: mouseEventTypeRaw},
				{in: release(left), want: mouseEventTypeClick},
			},
		},
		{
			name: "hover",
			steps: []step{
				{in: hover(), want: mouseEventTypeRaw},
			},
		},
		{
			name: "release without press",
			steps: []step{
				{in: release(left), want: mouseEventTypeRaw},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			now := time.Date(2022, 2, 24, 3, 55, 0, 0, time.UTC)
			mh := &mouseHandler{now: func() time.Time { return now }}

			for i, s := range tc.steps {
				now = now.Add(s.advance)
				got := mh.transformInput(s.in)
				if got.event != s.want {
					t.Errorf("step %d: transformInput() = %v, want %v", i, got.event, s.want)
				}
			}
		})
	}
}
