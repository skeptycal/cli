// +build !appengine,!js,windows

package terminal

// the isTerminal functionality from logrus is used here.
// MIT Licence Copyright (c) 2014 Simon Eskildsen
// https://github.com/sirupsen/logrus

import (
	"io"
	"os"

	"golang.org/x/sys/windows"
)

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		handle := windows.Handle(v.Fd())
		var mode uint32
		if err := windows.GetConsoleMode(handle, &mode); err != nil {
			return false
		}
		mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
		if err := windows.SetConsoleMode(handle, mode); err != nil {
			return false
		}
		return true
	}
	return false
}

func getWinsize() (*winsize, error) {
	ws := new(winsize)
	fd := os.Stdout.Fd()
	var info windows.ConsoleScreenBufferInfo
	if err := windows.GetConsoleScreenBufferInfo(windows.Handle(fd), &info); err != nil {
		return nil, err
	}

	ws.Col = uint16(info.Window.Right - info.Window.Left + 1)
	ws.Row = uint16(info.Window.Bottom - info.Window.Top + 1)

	return ws, nil
}
