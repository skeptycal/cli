package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/skeptycal/ansi"
	"github.com/skeptycal/cli"
)

var out = cli.New()

const fmtString string = "%s %3d %s"

func main() {
	out.CLS()
	err := ColorTest()
	if err != nil {
		log.Fatal(err)
	}
}

func ColorTest() error {
	sb := strings.Builder{}
	defer sb.Reset()

	sb.WriteString("-------------------------> Color Test <-------------------------\n")

	out.Println(sb.String())

	fgTest()
	bgTest()
	return nil
}

func fgTest() {
	sb := strings.Builder{}
	defer sb.Reset()

	fg := func(c int) string {
		color := ansi.NewColor(byte(c), 0, 0).String()
		return fmt.Sprintf("%s %3d %s", color, c, ansi.Reset)
	}

	prefix := "            "

	sb.WriteString(colorSet(0, 15, 8, 4, prefix, fg))
	sb.WriteString("\n")
	prefix = "  "

	sb.WriteString(colorSet(16, 231, 12, 6, prefix, fg))
	sb.WriteString("\n")
	sb.WriteString(colorSet(232, 255, 12, 6, prefix, fg))
	sb.WriteString("\n")

	out.Println(sb.String())
}

func bgTest() {
	sb := strings.Builder{}
	defer sb.Reset()

	bg36 := func(c int) string {
		color := ansi.NewColor(byte(c%36+18), byte(c), 1).String()
		return fmt.Sprintf("%s %3d %s", color, c, ansi.Reset)
	}

	bg24 := func(c int) string {
		color := ansi.NewColor(byte((243-c)%24-12), byte(c), 1).String()
		return fmt.Sprintf("%s %3d %s", color, c, ansi.Reset)
	}

	bg8 := func(c int) string {
		color := ansi.NewColor(byte(15-c), byte(c), 1).String()
		return fmt.Sprintf("%s %3d %s", color, c, ansi.Reset)
	}

	sb.WriteString("\n")

	prefix := "            "

	sb.WriteString(colorSet(0, 15, 8, 4, prefix, bg8))
	sb.WriteString("\n")
	prefix = "  "

	sb.WriteString(colorSet(16, 231, 12, 6, prefix, bg36))
	sb.WriteString("\n")
	sb.WriteString(colorSet(232, 255, 12, 6, prefix, bg24))
	sb.WriteString("\n")
	out.Println(sb.String())
}

func colorSet(start, end, major, minor int, prefix string, colorFunc func(int) string) string {

	bias := end - start
	sb := strings.Builder{}
	defer sb.Reset()
	sb.WriteString(prefix)

	for i := 1; i <= bias+1; i++ {
		sb.WriteString(colorFunc(i + start - 1))
		if minor > 0 && i%minor == 0 {
			sb.WriteByte(' ')
		}
		if major > 0 && i%major == 0 {
			sb.WriteString("\n")
			sb.WriteString(prefix)
		}

	}

	return sb.String()
}

type color int

func (c color) fg() string {
	fg := ansi.NewColor(byte(c), 0, 0).String()
	return fmt.Sprintf(fmtString, fg, c, ansi.Reset)
}

func (c color) bg() string {
	bg := ansi.NewColor(0, byte(c), 0).String()
	return fmt.Sprintf(fmtString, bg, c, ansi.Reset)
}

func (c color) String() string {
	return c.fg()
}
