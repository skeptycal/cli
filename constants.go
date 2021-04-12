// The constants.go file contains type definitions and constants mainly from the ansi package.
// Reference: github.com/skeptycal/ansi

package cli

import (
	"fmt"

	"github.com/skeptycal/ansi"
)

const (
	Reset   string = ansi.Reset
	Inverse byte   = ansi.Inverse
	// SetInverse string = ansi.SetInverse
	fmtANSI string = ansi.FmtANSI
)

type Ansi = ansi.Ansi

func simpleEncode(b byte) string {
	return fmt.Sprintf(fmtANSI, b)
}
