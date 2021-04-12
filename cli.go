// Package cli contains utility functions for dealing with cli commands within Go.
//
// The main component is the CLI interface that implements cli features through the Terminal struct.
// Reference: github.com/skeptycal/cli
//
// The ansi color terminal support is provided by the ansi package.
// Reference: github.com/skeptycal/ansi
package cli

import (
	"io"
	"os"

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
	Reset() (n int, err error)
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
		w:        w,
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
		w:        w,
		useColor: checkColor,
		devMode:  devMode,
	}
}
