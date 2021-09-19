package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/skeptycal/ansi"
	"github.com/skeptycal/cli"
)

const (
	title     string = "-------------------------> Color Test <-------------------------\n"
	fmtString string = "%s %3d %s"
)

var out = cli.New()

func main() {
	// out.CLS()
	n, err := ColorTest()
	if err != nil {
		log.Fatalf("error writing color test (%v bytes written) to terminal: %v", n, err)
	}
}

func fakeColorTest() (n int, err error) {
	return 42, errors.New("fake error for testing (should say 42 bytes)")
}

func ColorTest() (n int, err error) {

	nn, err := out.Println(title)
	if err != nil {
		return nn, err
	}
	n += nn

	nn, err = fgTest()
	if err != nil {
		return nn, err
	}

	n += nn

	nn, err = bgTest()
	if err != nil {
		return nn, err
	}

	n += nn

	return n, nil
}

func fgTest() (n int, err error) {
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

	return out.Println(sb.String())
}

func bgTest() (n int, err error) {
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

	return out.Println(sb.String())
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
