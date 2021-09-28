package util

import (
	"os"
	"path/filepath"

	"github.com/skeptycal/ansi"
	"github.com/skeptycal/cli"
)

const (
	optionList = "FILE"
)

var (
	NewColor          = ansi.NewColor
	Out      cli.CLI  = cli.New()
	Green    cli.Ansi = NewColor(2, 0, 0)
	White             = NewColor(15, 0, 0)
	Me       string   = filepath.Base(os.Args[0])
)

func init() {
	if len(os.Args) != 2 {
		usage()
	}
}

func usage() {
	Out.Printf("usage: %v%s %v%s\n", White, Me, Green, optionList)
	// os.Exit(0)
}
