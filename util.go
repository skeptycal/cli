package cli

import (
	"fmt"
	"strconv"
)

type AnsiColor struct {
	byte
}

func (a AnsiColor) String() string {
	return BasicEncode(a)
}

const nibbleMask = 1<<4 - 1

// BasicEncode encodes a basic (3-4 bit) ANSI color code.
// The code may be passed in as a byte, int, uint, float32,
// float64, or string and will be converted to the best
// guess of a byte (uint8) value before encoding.
//
// It is best to use a byte value ...
//
// The format is "\x1b[%dm"
func BasicEncode(b interface{}) string {
	switch v := b.(type) {
	case byte:
		return fmt.Sprintf(fmtANSI, v)
	case int:
		return fmt.Sprintf(fmtANSI, v&255)
	case uint:
		return fmt.Sprintf(fmtANSI, v&255)
	case float32, float64:
		return fmt.Sprintf(fmtANSI, byte(v.(float64)))
	case string:
		i, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return ""
		}
		return fmt.Sprintf(fmtANSI, i)
	default:
		return ""
	}

}
