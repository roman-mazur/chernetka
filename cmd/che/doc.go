// Command che starts an interactive text editor in your terminal.
//
// Usage:
//
//	che [-root dir] [file]
//	some-command | che
//
// When no file is given and stdin is a pipe, edit opens the piped content in a
// read-only buffer. When no file and no pipe are present, a blank scratch
// buffer is opened instead.
//
// Opening a directory is also supported. Editor will display a terminal UI allowing
// to navigate the directory tree. Pressing Enter on a file will attempt opening this
// file in the right pane of your terminal. The main editor opened this way searches files
// (see Quick open) in the same directory.
//
// The -root flag sets the project directory to search files in. It defaults to the
// opened directory or the working one.
//
// # Images and diagrams
//
// PNG images and d2 diagrams are shown with che-img (see its docs in cmd/che-img).
// If che-img is not running, che launches it in a new terminal pane below.
// A .d2 file is also opened for editing. When che is started with an image and has
// nothing else open, it turns into che-img in the same pane.
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
//	0        jump to the first non-blank character of the line, or to its start
//	         if the cursor is already there or within the indentation (same for Home)
//	$        jump to end of line
//	i        enter insert mode before the cursor
//	a        enter insert mode after the cursor
//	A        enter insert mode at end of line
//	o        open a new line below and enter insert mode
//	x        delete the character under the cursor
//	+ / =    increase tab width
//	-        decrease tab width
//	:        enter command mode
//	/        search (see below)
//	n / N    move to the next / previous search match
//	Esc      clear the selection and the search highlights
//	Ctrl+O   quick open a file or switch to a buffer (in any mode)
//	Ctrl+F   search (in any mode)
//	Ctrl+D   save the current file and show its git diff in a new pane on the right (in any mode)
//	Enter    engage the line action if the line is marked with ▶, move down otherwise
//	Ctrl+R   re-run the last engaged line action
//	q        quit
//
// # Line actions
//
// Some lines can be engaged with Enter, they are marked with a green ▶ on the right.
//
// d2 diagrams: the first line of a .d2 file and the line opening a ```d2 code block in
// a Markdown file show the diagram with che-img. After editing the diagram, press Ctrl+R
// to show it again.
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
//	Ctrl+R     re-run the last engaged line action
//
// # Code suggestions
//
// Suggestions come from a language server: gopls for Go files, and the one
// built into the cue command for CUE files.
//
//	Tab        accept the suggestion (adding a missing import if needed)
//	↑ ↓        cycle through the alternatives, if there are several
//	Esc        dismiss the suggestion (press again to return to normal mode)
//
// Go and CUE files are formatted with their language server when saved.
//
// Ctrl+click on a symbol goes to its definition, opening its file if needed.
// Ctrl+O and Enter switch back to the previous buffer.
//
// Unsaved changes of a file are saved automatically when switching to another
// buffer or when the terminal loses focus (if the terminal supports focus reporting).
// When an open file is changed outside the editor, its buffer is reloaded,
// discarding the unsaved changes.
//
// # Search and replace
//
// Type a regular expression (Go syntax) after / to search. The cursor moves to the
// first match from its position as you type, and all the matches are highlighted.
//
//	Tab / ↓ / Ctrl+F  move to the next match
//	Shift+Tab / ↑     move to the previous match
//	Enter             keep the cursor at the match and the matches highlighted
//	Esc               cancel, returning the cursor and the previous search highlights
//
// Like in sed, /pattern/replacement replaces the matches when Enter is pressed:
// all of them in the buffer, like with the g flag. In the replacement, & is the
// whole match, \1..\9 are the groups, \n and \t are a new line and a tab, and
// \/ is a slash (in the pattern too).
//
// # Command mode
//
// Entered by pressing : in normal mode. Type a command and press Enter to run it.
// Backspace removes the last character; on an empty command line, it closes it.
// The command line, the search and the file picker opened from insert mode return
// to insert mode when they are closed.
//
//	:q       quit
//	:w       save the current buffer
//	:w path  save the current buffer to path
//	:wq      save and quit
//	:e query quick open a file (see below)
//	Esc      cancel
//
// # Quick open
//
// Ctrl+O types :e in the command line. The matches for the query that follows are
// shown in the status bar: the open buffers first, starting from the previously
// used one, then the files in the project directory. The query characters must appear
// in the path in the same order, not necessarily next to each other. Files ignored
// by git are not listed.
//
//	Tab / ↓ / Ctrl+O   select the next match
//	Shift+Tab / ↑      select the previous match
//	Enter              open the selected match; without matches, open the query
//	                   as a path in the project directory (a new file if it does not exist)
//	Esc                cancel
//
// Ctrl+O and Enter switch to the previous buffer.
package main
