package cli

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/sys/unix"
	// log "github.com/sirupsen/logrus"
)

const newline byte = '\n'

type Terminal struct {

	// io.Writer is the interface that wraps the basic Write method.
	//
	// Write writes len(p) bytes from p to the underlying data stream.
	// It returns the number of bytes written from p (0 <= n <= len(p))
	// and any error encountered that caused the write to stop early.
	// Write must return a non-nil error if it returns n < len(p).
	// Write must not modify the slice data, even temporarily.
	//
	// Implementations must not retain p.
	w io.Writer `default:"defaultWriter"`

	// useColor is true if the terminal supports color and colored
	// output is desired. It is true by default on terminals that
	// support it.
	useColor bool `default:"true"`

	// on is an internal function that is called when the terminal
	// color is requested. The function is mapped here in lieu of
	// doing numerous checks of the useColor flag.
	//
	// on is set to point to doCheckColor() if useColor is true
	// and is set to point to noOp() if useColor is false.
	on func(w io.Writer, p []byte) (n int, err error)

	// devMode is a flag that indicates whether the terminal is being
	// used in a development environment where additional verbose
	// output and logging is often desired.
	devMode bool

	// colorBytes is the []byte representation of the current ANSI
	// color sequence(s) to be used when wrapping the terminal
	// output.
	colorBytes []byte

	// defaultForeground is the default ANSI 8bit foreground color
	// that is used when no other foreground is specified. This
	// requires a byte (int8) that is mapped to the ANSI code
	// with the appropriate format string:
	// FMTAnsi = "\033[38;5;%v;m"
	defaultForeground byte `default:"15"`
	defaultBackground byte // default is Zero Value (0)
	defaultEffect     byte // default is Zero Value (0)
	// useLog            bool // TODO - add logging
	// log               *log.Logger
}

// Reset sets the ANSI foreground, background, and effect to default.
func (t *Terminal) Reset() (n int, err error) {
	if t.useColor {
		return fmt.Fprint(t.w, Reset)
	}
	return 0, nil
}

func (t *Terminal) color() string {
	if t.useColor {
		return string(t.colorBytes)
	}
	return ""
}

// SetColor sets the ANSI foreground, background, and effect codes
// for upcoming output.
func (t *Terminal) SetColor(color Ansi) {
	if t.useColor {
		t.colorBytes = []byte(color.String())
	}
}

// String describes the terminal. If devMode is true, it generates a
// list of dev info.
func (t *Terminal) String() string {
	sb := &strings.Builder{}
	defer sb.Reset()
	if t.devMode {
		sb.WriteString(t.devinfo())
	}
	if !t.useColor {
		sb.WriteString("ANSI terminal - no color output.")
	} else {
		fmt.Fprintf(sb, "%vANSI color output (fg = %v, bg = %v, ef = %v) %v\n", t.color(), t.defaultForeground, t.defaultBackground, t.defaultEffect, Reset)
	}
	return sb.String()
}

func (t *Terminal) devinfo() string {
	sb := &strings.Builder{}
	defer sb.Reset()
	if t.devMode {
		t.Hr()
		fmt.Fprintf(sb, "CLI variable (DefaultForeground): %v\n", t.defaultForeground)
		fmt.Fprintf(sb, "CLI variable (DefaultBackground): %v\n", t.defaultBackground)
		fmt.Fprintf(sb, "CLI variable (DefaultEffect): %v\n", t.defaultEffect)
		fmt.Fprintf(sb, "CLI variable (Color): %q\n", t.color())
		fmt.Fprintf(sb, "CLI variable (UseColor): %t\n", t.useColor)
		fmt.Fprintf(sb, "CLI variable (devMode): %t\n", t.devMode)
		fmt.Fprintf(sb, "CLI writer pointer: %v\n\n", &t.w)
	}
	return sb.String()
}

// GetTerminalSize returns device caps for the terminal.
// The Winsize struct returned includes:
//  Row, Col, Xpixel, and Ypixel.
func GetTerminalSize() (*unix.Winsize, error) {
	return getWinsize()
}

// Columns returns the number of columns in the terminal,
// similar to the COLUMNS environment variable on macOS
// and Linux systems.
func Columns() int {
	ws, err := GetTerminalSize()
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
