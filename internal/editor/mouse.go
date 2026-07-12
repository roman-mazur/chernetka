package editor

import (
	"fmt"
	"time"

	"rmazur.io/chernetka/internal/editor/inputs"
)

// mouseEvent wraps the raw mouse input extending it with the event type derived from the saved state.
type mouseEvent struct {
	inputs.Mouse
	event mouseEventType
}

// mouseEventType indicates an event derived from the manipulations with a mouse like click, double-click, etc.
type mouseEventType byte

const (
	mouseEventTypeRaw         mouseEventType = iota // transmitting the raw press/release input
	mouseEventTypeClick                             // quick press and release of the same button
	mouseEventTypeDoubleClick                       // press/release that happen twice quickly
	mouseEventTypeTripleClick                       // three press/release cycles
	mouseEventTypeDragStart                         // long press and start moving
	mouseEventTypeDragEnd                           // release after a long press and move
)

func (et mouseEventType) String() string {
	switch et {
	case mouseEventTypeRaw:
		return "Raw"
	case mouseEventTypeClick:
		return "Click"
	case mouseEventTypeDoubleClick:
		return "DoubleClick"
	case mouseEventTypeTripleClick:
		return "TripleClick"
	case mouseEventTypeDragStart:
		return "DragStart"
	case mouseEventTypeDragEnd:
		return "DragEnd"
	default:
		return fmt.Sprintf("mouseEventType(%d)", byte(et))
	}
}

type mouseHandler struct {
	now         func() time.Time
	lastButton  inputs.MouseButton
	lastPressAt time.Time
	clickCount  int
	isDragging  bool
}

func (mh *mouseHandler) currentTime() time.Time {
	if mh.now != nil {
		return mh.now()
	}
	return time.Now()
}

func (mh *mouseHandler) transformInput(in inputs.Mouse) mouseEvent {
	if in.Mod.HasMotion() {
		if in.Pressed && !mh.isDragging {
			mh.isDragging = true
			return mouseEvent{Mouse: in, event: mouseEventTypeDragStart}
		}
		return mouseEvent{Mouse: in, event: mouseEventTypeRaw}
	}

	if in.Pressed {
		now := mh.currentTime()
		withinDelay := !mh.lastPressAt.IsZero() && now.Sub(mh.lastPressAt) <= mouseDelayDoubleClick
		if in.Button == mh.lastButton && withinDelay {
			mh.clickCount++
		} else {
			mh.clickCount = 1
			mh.isDragging = false
		}
		mh.lastButton = in.Button
		mh.lastPressAt = now
		return mouseEvent{Mouse: in, event: mouseEventTypeRaw}
	}

	// Button release.
	if mh.isDragging {
		mh.isDragging = false
		mh.clickCount = 0
		return mouseEvent{Mouse: in, event: mouseEventTypeDragEnd}
	}

	switch mh.clickCount {
	case 0:
		return mouseEvent{Mouse: in, event: mouseEventTypeRaw}
	case 1:
		return mouseEvent{Mouse: in, event: mouseEventTypeClick}
	case 2:
		return mouseEvent{Mouse: in, event: mouseEventTypeDoubleClick}
	default:
		mh.clickCount = 0
		return mouseEvent{Mouse: in, event: mouseEventTypeTripleClick}
	}
}

const mouseDelayDoubleClick = 500 * time.Millisecond
