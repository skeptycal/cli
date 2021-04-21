// +build !windows,!plan9,!solaris

package cli

// getWinsize contains code from the goterm package
// Reference: https://github.com/buger/goterm (MIT License)
// Reference: https://www.darkcoding.net/software/pretty-command-line-console-output-on-unix-in-python-and-go-lang/

import (
	"os"

	"golang.org/x/sys/unix"
)

func getWinsize() (*unix.Winsize, error) {

	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, os.NewSyscallError("GetWinsize", err)
	}

	return ws, nil
}
