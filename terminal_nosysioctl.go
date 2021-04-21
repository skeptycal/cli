// +build plan9 solaris

package cli

// getWinsize contains code from the goterm package
// Reference: https://github.com/buger/goterm (MIT License)
// Reference: https://www.darkcoding.net/software/pretty-command-line-console-output-on-unix-in-python-and-go-lang/

func getWinsize() (*winsize, error) {
	ws := new(winsize)

	ws.Col = 80
	ws.Row = 24

	return ws, nil
}
