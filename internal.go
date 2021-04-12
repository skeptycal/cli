package cli

import "fmt"

func simpleEncode(b byte) string {
	return fmt.Sprintf(fmtANSI, b)
}
