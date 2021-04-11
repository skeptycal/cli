package cli

import (
	"fmt"

	"github.com/skeptycal/ansi"
)

const (
	Reset      string = ansi.Reset
	SetInverse string = ansi.SetInverse
	fmtANSI    string = ansi.FmtANSI
)

type Ansi = ansi.Ansi

func simpleEncode(b byte) string {
	return fmt.Sprintf(fmtANSI, b)
}
