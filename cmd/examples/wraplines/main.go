package main

import (
	"strings"

	"github.com/skeptycal/ansi"
	"github.com/skeptycal/cli"
)

var (
	out   = cli.New()
	green = ansi.NewColor(2, 0, 0)
	fake  = "abcdefghijklmnopqrstuvwxyz01234567890: "
)

func main() {

	// out.CLS()
	out.SetColor(green)

	fake = strings.Repeat("fake stuff ", 42)

	testWrap(fake, cli.Columns())
	testWrap(fake, 120)
	testWrap(fake, 79)
	testWrap(fake, 1000)
}

func testWrap(s string, width int) {
	out.Br()
	out.Printf("Wrap string of length %v to width of %v\n", len(s), width)
	out.Println("------------------------------------------")
	out.Println(cli.Wrap(s, width))
}
