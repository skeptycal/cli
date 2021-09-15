package cli

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/skeptycal/cli/terminal"
)

var GetWinSize = terminal.GetWinsize

func NewColor(in byte) *AnsiColor {
	return &AnsiColor{in, ""}
}

// AnsiColor is a buffered, encoded ANSI color string typically used
// for CLI output. The encoded ANSI color code string is JIT
// buffered at the time of the first output request to eliminate
// repeated fmt.Sprintf (or similar) calls.
type AnsiColor struct {
	Color byte
	out   string
}

func (a AnsiColor) String() string {
	if a.out == "" {
		a.out = BasicEncode(a.Color)
	}
	return a.out
}

func encode(b byte) string {
	// return ansiPrefix + string([]byte{b}) + ansiSuffix
	return fmt.Sprintf(fmtANSI, b)
}

func byteEncode(b Any) (byte, error) {
	switch v := b.(type) {
	case byte:
		return v, nil
	case int:
		return byte(v & 255), nil
	case uint:
		return byte(v & 255), nil
	case float32, float64:
		return byte(v.(float64)), nil
	case string:
		i, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return 0, err
		}
		return byte(i), nil
	default:
		return 0, errors.New("encoding error: unable to convert input to byte")
	}
}

// BasicEncode encodes a basic (3-4 bit) ANSI color code.
// The code may be passed in as a byte, int, uint, float32,
// float64, or string and will be converted to the best
// guess of a byte (uint8) value before encoding.
//
// It is best to use a byte value ...
//
// The format is "\x1b[%dm"
func BasicEncode(b interface{}) string {
	c, err := byteEncode(b)

	if err != nil {
		return ""
	}

	return encode(c)
}

func CheckIfTerminal(w io.Writer) bool { return terminal.CheckIfTerminal(w) }

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

func Rows() int {
	ws, err := GetWinSize()
	if err != nil {
		return 0
	}
	return int(ws.Row)
}

func XPixels() int {
	ws, err := GetWinSize()
	if err != nil {
		return 0
	}
	return int(ws.Xpixel)
}

func YPixels() int {
	ws, err := GetWinSize()
	if err != nil {
		return 0
	}
	return int(ws.Ypixel)
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
