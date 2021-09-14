// The constants.go file contains type definitions and constants mainly from the ansi package.
// Reference: github.com/skeptycal/ansi

package cli

import (
	"github.com/skeptycal/ansi"
)

type Ansi = ansi.Ansi

const (
	// Reset all custom styles
	Reset string = ansi.Reset // "\033[0m"

	// ResetColor - Reset to default color
	ResetColor = "\033[32m"

	// ResetLineConst - Return cursor to start of line and clean it
	ResetLineConst = "\r\033[K"

	// inverse    byte   = ansi.Inverse
	SetInverse string = ansi.SetInverse
	fmtANSI    string = "\x1b[%dm" // ansi.FmtANSI
)

// List of possible colors
const (
	BLACK = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

var (
	ResetBytes   []byte = ansi.ResetBytes
	InverseBytes []byte = []byte(simpleEncode(ansi.Inverse))
)
