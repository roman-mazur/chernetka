package inputs

type Cursor byte

//go:generate go run golang.org/x/tools/cmd/stringer -type=Cursor -trimprefix=Cursor

const (
	CursorArrowUp Cursor = iota
	CursorArrowDown
	CursorArrowLeft
	CursorArrowRight
	CursorHome
	CursorEnd
)

// Modifier is a bit mask giving more context about the Mouse or keyboard event.
// bit	meaning
// 0	Shift
// 1	Alt
// 2	Ctrl
// 3	Motion / command
// 4	Scroll wheel
type Modifier byte

func (m Modifier) HasShift() bool  { return m&1 != 0 }
func (m Modifier) HasAlt() bool    { return m&2 != 0 }
func (m Modifier) HasCtrl() bool   { return m&4 != 0 }
func (m Modifier) HasMotion() bool { return m&8 != 0 }
func (m Modifier) HasWheel() bool  { return m&16 != 0 }

func (m Modifier) SrollDirection(mouse Mouse) ScrollDirection { return ScrollDirection(mouse.Button) }

func (m Modifier) String() string {
	var data [5]byte
	for i := range data {
		data[i] = '-'
	}
	if m.HasShift() {
		data[0] = 's'
	}
	if m.HasAlt() {
		data[1] = 'a'
	}
	if m.HasCtrl() {
		data[2] = 'c'
	}
	if m.HasMotion() {
		data[3] = 'm'
	}
	if m.HasWheel() {
		data[4] = 'w'
	}
	return string(data[:])
}
