// Command che starts an interactive text editor in your terminal.
//
// Usage:
//
//	che [file]
//	some-command | che
//
// When no file is given and stdin is a pipe, edit opens the piped content in a
// read-only buffer. When no file and no pipe are present, a blank scratch
// buffer is opened instead.
//
// Opening a directory is also supported. Editor will display a terminal UI allowing
// to navigate the directory tree. Pressing Enter on a file will attempt opening this
// file in the right pane of your terminal.
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
//	+ / =    increase tab width
//	-        decrease tab width
//	:        enter command mode
//	q        quit
//
// # Insert mode
//
//	Esc        return to normal mode
//	(any)      insert printable ASCII character at the cursor
//	Tab        insert a tab character
//	Backspace  delete the character before the cursor;
//	           if at column 0, join the line with the one above
//	Enter      split the line at the cursor
//	← → ↑ ↓   move the cursor without leaving insert mode
//	Ctrl+S     save the current buffer
//
// # Command mode
//
// Entered by pressing : in normal mode. Type a command and press Enter to run it.
// Backspace removes the last character; an empty command line returns to normal mode.
//
//	:q       quit
//	:w       save the current buffer
//	:w path  save the current buffer to path
//	:wq      save and quit
//	Esc      cancel and return to normal mode
package main
