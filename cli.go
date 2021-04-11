package cli

import (
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/ansi"
)

const (
// defaultCLIforeground byte = 15
// defaultCLIbackground byte = 0
// defaultCLIeffect          = 0
)

var (
	defaultWriter      io.Writer = newAnsiStdout()
	defaultErrorWriter io.Writer = newAnsiStderr()
)

// newAnsiStdout returns stdout which converts escape sequences
// to Windows API calls on Windows environment.
func newAnsiStdout() io.Writer {
	return os.Stdout
}

// newAnsiStderr returns stdout which converts escape sequences
// to Windows API calls on Windows environment.
func newAnsiStderr() io.Writer {
	return os.Stderr
}

// CLI implements an ANSI compatible terminal interface.
type CLI interface {
	io.Writer
	io.StringWriter
	String() string
	Print(args ...interface{}) (n int, err error)
	Printf(format string, args ...interface{}) (n int, err error)
	Println(args ...interface{}) (n int, err error)
	SetColor(color ansi.Ansi)
	Reset()
	CLIControls
}

func NewStdout(w io.Writer) CLI {
	checkColor := true // TODO check if color is supported - bring this code over ...
	devMode := defaultDevMode
	if w == nil {
		w = defaultWriter
	}
	if _, ok := w.(io.Writer); !ok {
		w = defaultWriter
	}

	t := &Terminal{
		Writer:   w,
		useColor: checkColor,
		devMode:  devMode,
	}

	if checkColor {
		t.on = doCheckColor
	} else {
		t.on = noOp
	}

	return t
}

func NewStderr(w io.Writer) CLI {
	checkColor := true // todo bring this code over ...
	devMode := defaultDevMode
	if w == nil {
		w = defaultErrorWriter
	}
	if _, ok := w.(io.Writer); !ok {
		w = defaultErrorWriter
	}

	return &Terminal{
		Writer:   w,
		useColor: checkColor,
		devMode:  devMode,
	}
}

type Terminal struct {
	Writer                           io.Writer `default:"defaultWriter"`
	useColor                         bool      `default:"true"`
	on                               func(w io.Writer, p []byte) (n int, err error)
	devMode                          bool
	colorBytes                       []byte
	useLog                           bool
	log                              *log.Logger
	DefaultForeground                byte `default:"15"`
	DefaultBackground, DefaultEffect byte // default is Zero Value (0)
}

// Inverse sets the inverse ANSI effect if the terminal supports it.
func (t *Terminal) Inverse() {
	if t.useColor {
		t.Print(simpleEncode(Inverse))
	}
}

// Reset sets the ANSI foreground, background, and effect to default.
func (t *Terminal) Reset() {
	if t.useColor {
		fmt.Fprint(t.Writer, Reset)
	}
}

func (t *Terminal) SetDefault() {

}

func (t *Terminal) Color() string {
	if t.useColor {
		return string(t.colorBytes)
	}
	return ""
}

func (t *Terminal) Info() {
	if t.useLog {
		t.log.Info()
	}
}

// func (t *Terminal) SetLogging(b bool) {
// 	t.useLog = b
// 	if ok := xxx.(log.Logger); ok {

// 	}
// }

// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
// If w implements StringWriter, its WriteString method is invoked directly.
// Otherwise, w.Write is called exactly once.
func WriteString(w io.Writer, s string) (n int, err error) {
	if sw, ok := w.(io.StringWriter); ok {
		return sw.WriteString(s)
	}
	return w.Write([]byte(s))
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
		fmt.Fprintf(sb, "%vANSI color output (fg = %v, bg = %v, ef = %v) %v\n", t.Color(), t.DefaultForeground, t.DefaultBackground, t.DefaultEffect, Reset)
	}
	return sb.String()
}

func (t *Terminal) devinfo() string {
	sb := &strings.Builder{}
	defer sb.Reset()
	if t.devMode {
		t.Hr()
		fmt.Fprintf(sb, "CLI variable (DefaultForeground): %v\n", t.DefaultForeground)
		fmt.Fprintf(sb, "CLI variable (DefaultBackground): %v\n", t.DefaultBackground)
		fmt.Fprintf(sb, "CLI variable (DefaultEffect): %v\n", t.DefaultEffect)
		fmt.Fprintf(sb, "CLI variable (Color): %q\n", t.Color())
		fmt.Fprintf(sb, "CLI variable (UseColor): %t\n", t.useColor)
		fmt.Fprintf(sb, "CLI variable (devMode): %t\n", t.devMode)
		fmt.Fprintf(sb, "CLI writer pointer: %v\n\n", &t.Writer)
	}
	return sb.String()
}
