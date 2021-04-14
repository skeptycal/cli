package main

import (
	"github.com/skeptycal/cli"
)

func main() {
	out := cli.New()

	out.CLS()

	out.Print("CLI example")
	out.Br()

	out.Print(cli.InverseBytes)
	out.Print("CLI inverse example")

}
