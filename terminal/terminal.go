// Package terminal provides information about the state of the terminal.
package terminal

import (
	"io"

	"golang.org/x/sys/unix"
)

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

type winsize = Winsize

func GetWinsize() (*unix.Winsize, error) {
	return getWinsize()
}

func CheckIfTerminal(w io.Writer) bool {
	return checkIfTerminal(w)
}

func IsTerminal(fd int) bool {
	return isTerminal(fd)
}
