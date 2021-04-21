package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/skeptycal/ansi"
	"github.com/skeptycal/cli"
)

var out = cli.New()

const newline byte = '\n'

const (
	fmtString string = "%s %3d %s"
	fmtLength int    = 6
)

type colorString int

func (c colorString) fg() string {
	fg := ansi.NewColor(byte(c), 0, 0).String()
	return fmt.Sprintf(fmtString, fg, c, ansi.Reset)
}

func (c colorString) bg() string {
	bg := ansi.NewColor(0, byte(c), 0).String()
	return fmt.Sprintf(fmtString, bg, c, ansi.Reset)
}

func (c colorString) String() string {
	return c.fg()
}

func main() {
	out.CLS()
	err := ColorTest()
	if err != nil {
		log.Fatal(err)
	}
}

func ColorTest() error {

	width := cli.Columns()

	unitWidth := width / fmtLength

	out.Printf("width: %v\n", width)
	out.Printf("fmtLength: %v\n", fmtLength)
	out.Printf("unitWidth: %v\n", unitWidth)

	sb := strings.Builder{}
	defer sb.Reset()

	for f := 0; f < 255; f++ {
		sb.WriteString(colorString(f).fg())
	}
	sb.WriteString("\n")
	for f := 0; f < 255; f++ {
		sb.WriteString(colorString(f).bg())
	}

	lines := sb.String()
	sb.Reset()

	sb.WriteString("Color Test\n\n")
	sb.WriteString(lines)
	sb.WriteByte(newline)
	retval := sb.String()

	out.Println(retval)

	return nil
}
