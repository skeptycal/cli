package cli

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/skeptycal/cli/terminal"
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

func CheckIfTerminal(w io.Writer) bool {
	return terminal.CheckIfTerminal(w)
}

// Columns returns the number of columns in the terminal,
// similar to the COLUMNS environment variable on macOS
// and Linux systems.
func Columns() int {
	ws, err := GetWinSize()
	if err != nil {
		return 0
	}
	return int(ws.Col)
}

// Wrap splits a string into lines no longer than width.
func Wrap(s string, width int) string {

	if len(s) <= width {
		return s
	}

	sb := strings.Builder{}
	defer sb.Reset()

	for {
		i := width
		for {
			if s[i] == ' ' {
				break
			}
			i--
		}

		part := strings.TrimSpace(s[:i])
		sb.WriteString(part)
		sb.WriteByte(newline)

		s = strings.TrimSpace(s[i:])
		if len(s) <= width {
			sb.WriteString(s)
			sb.WriteByte(newline)
			break
		}
	}
	return sb.String()
}
