package main

import (
	"bytes"
	"os"

	"github.com/skeptycal/cli"
	. "github.com/skeptycal/cli/cmd/utilities/wrapper/util"
)

const fakeFileData = "abcdefghijklmnopqrstuvwxyz01234567890: "

func main() {
	Out.CLS()

	buf := bytes.NewBuffer(bytes.Repeat([]byte(fakeFileData), 42))

	w := cli.NewBasicWrapper(buf, os.Stdout)
	Out.Print(w.String())
}
