// The constants.go file contains type definitions and constants mainly from the ansi package.
// Reference: github.com/skeptycal/ansi

package cli

import (
	"github.com/skeptycal/ansi"
)

type Ansi = ansi.Ansi

const (
	Reset   string = ansi.Reset
	inverse byte   = ansi.Inverse
	// SetInverse string = ansi.SetInverse
	fmtANSI string = ansi.FmtANSI
	// ResetBytes []byte = []byte(ansi.Reset)
)

var (
	ResetBytes   []byte = ansi.ResetBytes
	InverseBytes []byte = []byte(simpleEncode(inverse))
)
