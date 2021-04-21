package main

import (
	"github.com/skeptycal/ansi"
	"github.com/skeptycal/cli"
)

func main() {
	out := cli.New()

	out.CLS()

	out.Print("CLI example")
	out.Br()
	var blue = ansi.NewColor(4, 0, 1).String()

	var green = ansi.NewColor(2, 0, 1).String()

	out.Printf("%s%s%s%s\n", blue, "Blue stuff...", ansi.Reset, "   reset ...")
	out.Printf("%s%s%s%s\n", green, "Green stuff...", ansi.Reset, "   reset ...")

	out.Println("")

	out.Print(cli.InverseBytes, "Inverse")
	out.Print("CLI inverse example")
	out.Println()

	counter := 0
	for f := 0; f < 255; f++ {
		color := ansi.NewColor(byte(f), 0, 0)
		out.Printf("%s %4d ", color.String(), f)
		counter += 1
		if counter%20 == 0 {
			out.Println()
			counter = 0
		}
	}

	out.Println()
	out.Println()

	counter = 0
	for f := 0; f < 255; f++ {
		color := ansi.NewColor(0, byte(f), 0)
		out.Printf("%s %4d ", color.String(), f)
		counter += 1
		if counter%20 == 0 {
			out.Println()
			counter = 0
		}
	}

	out.Println()
	out.Println()

	// for b := 0; b < 255; b++ {
	// 	for e := 0; e < 9; e++ {

	// 	}
	// }
}
