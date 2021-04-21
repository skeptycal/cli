package main

import (
	"fmt"
	"log"
	"os/exec"
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

func GetScreenWidth() (int, error) {
	ws, err := cli.GetTerminalSize()
	if err != nil {
		return 0, err
	}
	return int(ws.Col), nil
}

func ColorTest() error {

	width, err := GetScreenWidth()
	if err != nil {
		return err
	}

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

	lines := SplitLinesSize(sb.String(), width)
	sb.Reset()

	sb.WriteString("Color Test\n\n")
	sb.WriteString(lines)
	sb.WriteByte(newline)
	retval := sb.String()

	out.Println(retval)

	return nil
}

func ShellTest() {
	app := "echo"

	arg0 := "-e"
	arg1 := "Hello world"
	arg2 := "\n\tfrom"
	arg3 := "${COLUMNS}"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}

// Cmd 封装exec ，有shell= true 这样的选项
// Reference: https://stackoverflow.com/a/27764262
func Cmd(cmd string, shell bool) []byte {

	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out

	}
}
