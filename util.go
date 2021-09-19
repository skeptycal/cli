package cli

import (
	"io"
	"strings"

	"github.com/buger/goterm"
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

// BasicEncode encodes a basic (3-4 bit) ANSI color code.
// The code may be passed in as a byte, int, uint, float32,
// float64, or string and will be converted to the best
// guess of a byte (uint8) value before encoding.
//
// It is best to simply use a byte value ...
//
// The format is "\x1b[%dm"
func BasicEncode(b interface{}) string {
	c, err := byteEncode(b)

	if err != nil {
		return ""
	}

	return simpleEncode(c)
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

func CheckIfTerminal(w io.Writer) bool {

	// TODO not working ... see issues
	return terminal.CheckIfTerminal(w)
}

// Columns returns the number of columns in the terminal,
// similar to the COLUMNS environment variable on macOS
// and Linux systems.
func Columns() int {
	return goterm.Width()
}

// Rows returns the number of rows in the terminal,
func Rows() int {
	return goterm.Height()
}

// func XPixels() int {
// 	ws, err := GetWinSize()
// 	if err != nil {
// 		return 0
// 	}
// 	return int(ws.Xpixel)
// }

// func YPixels() int {
// 	ws, err := GetWinSize()
// 	if err != nil {
// 		return 0
// 	}
// 	return int(ws.Ypixel)
// }
