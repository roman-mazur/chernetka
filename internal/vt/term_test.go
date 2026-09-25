package vt

import "testing"

func TestWindowSize_CellSize(t *testing.T) {
	for _, tc := range []struct {
		name   string
		size   WindowSize
		w, h   float64
		wantOK bool
	}{
		{name: "pixels reported", size: WindowSize{Cols: 100, Rows: 50, XPixel: 1000, YPixel: 1100}, w: 10, h: 22, wantOK: true},
		{name: "fractional", size: WindowSize{Cols: 3, Rows: 2, XPixel: 10, YPixel: 5}, w: 10.0 / 3, h: 2.5, wantOK: true},
		{name: "no pixels", size: WindowSize{Cols: 100, Rows: 50}},
		{name: "no width in pixels", size: WindowSize{Cols: 100, Rows: 50, YPixel: 1000}},
		{name: "empty window", size: WindowSize{XPixel: 1000, YPixel: 1000}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w, h, ok := tc.size.CellSize()
			if w != tc.w || h != tc.h || ok != tc.wantOK {
				t.Errorf("CellSize() = %v, %v, %t; want %v, %v, %t", w, h, ok, tc.w, tc.h, tc.wantOK)
			}
		})
	}
}

func TestTestTerminal_Size(t *testing.T) {
	size, err := TestTerminal(80, 24, nil).Size()
	if err != nil || size != (WindowSize{Cols: 80, Rows: 24}) {
		t.Errorf("Size() = %+v, %v", size, err)
	}
}
