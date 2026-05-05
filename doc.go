// Command edit starts an interactive text editor in your terminal.
//
// Usage:
//
//	edit [file]
//	some-command | edit
//
// When no file is given and stdin is a pipe, edit opens the piped content in a
// read-only buffer. When no file and no pipe are present, a blank scratch
// buffer is opened instead.
//
// # Normal mode
//
// The editor starts in normal mode. The cursor rests on a character (never
// past the last one).
//
//	h / ←    move left
//	l / →    move right
//	j / ↓    move down
//	k / ↑    move up
//	0        jump to start of line
//	$        jump to end of line
//	i        enter insert mode before the cursor
//	a        enter insert mode after the cursor
//	A        enter insert mode at end of line
//	o        open a new line below and enter insert mode
//	x        delete the character under the cursor
//	+ / =    increase tab width (step 2, max 8)
//	-        decrease tab width (step 2, min 0)
//	q        quit
//
// # Insert mode
//
//	Esc        return to normal mode
//	(any)      insert printable ASCII character at the cursor
//	Backspace  delete the character before the cursor;
//	           if at column 0, join the line with the one above
//	Enter      split the line at the cursor
//	← → ↑ ↓   move the cursor without leaving insert mode
package main
