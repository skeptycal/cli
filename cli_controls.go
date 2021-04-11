package cli

import (
	"fmt"
	"strings"
)

type CLIControls interface {
	CLS()
	ClearLine()
	Hr()
	Br()
	CursorControls
	UseColor(b bool)
	DevMode(b bool)
}

const (
	ansiCLS            string = "\033[2J"  // ANSI escape code - clear screen
	ansiClearLine      string = "\r\033[K" // Carriage return and ANSI clear line code
	defaultHrChar      string = "="        // repeat character for Hr()
	defaultScreenWidth int    = 80         // default screen width if not specified
)

// CLS clears the screen.
func (t *Terminal) CLS() {
	t.Print(ansiCLS)
}

// ClearLine clears moves the cursor to the start of the
// the current line and clears the line.
func (t *Terminal) ClearLine() {
	t.Print(ansiClearLine)
}

// LineBreak creates a CLI line break (dotted line, underline, etc.)
// by repeating the string c enough times to span the screen width.
//
// This is useful in delimiting lines of text in terminal output.
func (t *Terminal) LineBreak(c string) {
	if c == "" {
		c = defaultHrChar
	}
	width := t.screenWidth()

	repeatCount := width/len(c) + 1

	s := strings.Repeat(c, repeatCount)[:width]
	t.Println(s)
}

// Hr creates a hard return using the default character.
func (t *Terminal) Hr() {
	t.LineBreak(defaultHrChar)
}

// Br creates a line break.
func (t *Terminal) Br() {
	fmt.Fprintln(t.Writer, "")
}

func (t *Terminal) screenWidth() int {
	// todo - add support for variable width
	return defaultScreenWidth
}

// UseColor manually sets the use of ANSI color output to true or false.
func (t *Terminal) UseColor(b bool) {
	if b == t.useColor {
		return
	}
	if !b {
		fmt.Fprint(t.Writer, Reset)
	}
	t.useColor = b
}

const defaultDevMode = false

// DevMode manually sets the Dev mode to true (for debugging)
// or false (for production). Default is false.
func (t *Terminal) DevMode(b bool) {
	t.devMode = b
}
