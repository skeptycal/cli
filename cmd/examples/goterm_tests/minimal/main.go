package main

import (
	"fmt"

	"github.com/buger/goterm"
)

// TODO - error in goterm results:
/*
	[Running] go run "/Users/skeptycal/go/src/github.com/skeptycal/_small_repos/cli/cmd/examples/goterm_test/main.go"
	Terminal height :  -2147483648
	Terminal width :  -1

	The problem is with the VSCode environment ... it works from command line
	"Solution" ... but not really ... run go run ./main.go from the command line ... and it works ...
*/

func main() {
	terminalHeight := goterm.Height()
	terminalWidth := goterm.Width()

	fmt.Println("Terminal height : ", terminalHeight)
	fmt.Println("Terminal width : ", terminalWidth)

}
