package cli

import (
	"fmt"
	"io"
	"strings"
	// log "github.com/sirupsen/logrus"
)

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

	useColor          bool `default:"true"`
	on                func(w io.Writer, p []byte) (n int, err error)
	devMode           bool
	colorBytes        []byte
	defaultForeground byte `default:"15"`
	defaultBackground byte // default is Zero Value (0)
	defaultEffect     byte // default is Zero Value (0)
	// useLog            bool // TODO - add logging
	// log               *log.Logger
}

// Inverse sets the inverse ANSI effect if the terminal supports it.
func (t *Terminal) Inverse() {
	if t.useColor {
		t.Print(simpleEncode(Inverse))
	}
}

// Reset sets the ANSI foreground, background, and effect to default.
func (t *Terminal) Reset() (n int, err error) {
	if t.useColor {
		return fmt.Fprint(t.w, Reset)
	}
	return 0, nil
}

// func (t *Terminal) SetDefault() {

// }

func (t *Terminal) Color() string {
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
		fmt.Fprintf(sb, "%vANSI color output (fg = %v, bg = %v, ef = %v) %v\n", t.Color(), t.defaultForeground, t.defaultBackground, t.defaultEffect, Reset)
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
		fmt.Fprintf(sb, "CLI variable (Color): %q\n", t.Color())
		fmt.Fprintf(sb, "CLI variable (UseColor): %t\n", t.useColor)
		fmt.Fprintf(sb, "CLI variable (devMode): %t\n", t.devMode)
		fmt.Fprintf(sb, "CLI writer pointer: %v\n\n", &t.w)
	}
	return sb.String()
}
