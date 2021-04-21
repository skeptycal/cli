package main

import (
	"strings"

	"github.com/skeptycal/ansi"
	"github.com/skeptycal/cli"
)

func main() {

	var (
		out   = cli.New()
		green = ansi.NewColor(2, 0, 0)
		fake  = strings.Repeat("fake stuff ", 32)
	)

	out.CLS()
	out.SetColor(green)

	width := cli.Columns()
	out.Println("Terminal current width : ", width)
	out.Println(cli.Wrap(fake, width))

	out.Br()
	width = 120
	out.Println("Screen width: ", width)
	out.Println(cli.Wrap(fake, width))

	width = 79
	out.Br()
	out.Println("Screen width: ", width)
	out.Println(cli.Wrap(fake, width))
}
