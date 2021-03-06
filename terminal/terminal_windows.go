// +build windows

package terminal

// getWinsize contains code from the goterm package
// Reference: https://github.com/buger/goterm (MIT License)
// Reference: https://www.darkcoding.net/software/pretty-command-line-console-output-on-unix-in-python-and-go-lang/

import (
	"golang.org/x/sys/windows"
	"os"
)

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
