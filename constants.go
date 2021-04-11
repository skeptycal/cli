package cli

import (
	"fmt"

	"github.com/skeptycal/ansi"
)

const (
	Reset   = ansi.Reset
	Inverse = ansi.Inverse
	fmtANSI = ansi.FmtANSI
)

type Ansi = ansi.Ansi

func simpleEncode(b byte) string {
	return fmt.Sprintf(fmtANSI, b)
}
