package cli

type CursorControls interface {
	Up(n int)
	Down(n int)
	Left(n int)
	Right(n int)
	Show()
	Hide()
}

// Format strings for ANSI escape codes for cursor control
const (
	fmtAnsiUp      string = "\x1b[%dA"
	fmtAnsiDown    string = "\x1b[%dB"
	fmtAnsiRight   string = "\x1b[%dC"
	fmtAnsiLeft    string = "\x1b[%dD"
	fmtAnsiNext    string = "\x1b[%dE"
	fmtAnsiPrev    string = "\x1b[%dF"
	fmtAnsiX       string = "\x1b[%dG"
	ansiShowCursor string = "\x1b[?25h"
	ansiHideCursor string = "\x1b[?25l"
)

// Up moves the cursor n cells to up.
func (t *Terminal) Up(n int) {
	t.Printf(fmtAnsiUp, n)
}

// Down moves the cursor n cells to down.
func (t *Terminal) Down(n int) {
	t.Printf(fmtAnsiDown, n)
}

// Right moves the cursor n cells to right.
func (t *Terminal) Right(n int) {
	t.Printf(fmtAnsiRight, n)
}

// Left moves the cursor n cells to left.
func (t *Terminal) Left(n int) {
	t.Printf(fmtAnsiLeft, n)
}

// NextLine moves the cursor to beginning of the line n lines down.
func (t *Terminal) NextLine(n int) {
	t.Printf(fmtAnsiNext, n)
}

// PreviousLine moves the cursor to beginning of the line n lines up.
func (t *Terminal) PreviousLine(n int) {
	t.Printf(fmtAnsiPrev, n)
}

// CursorX moves the cursor horizontally to x.
func (t *Terminal) CursorX(x int) {
	t.Printf(fmtAnsiX, x)
}

// Show the cursor.
func (t *Terminal) Show() {
	t.Print(ansiShowCursor)
}

// Hide the cursor.
func (t *Terminal) Hide() {
	t.Print(ansiHideCursor)
}
