// +build js nacl plan9

package cli

// the isTerminal functionality from logrus is used here.
// MIT Licence Copyright (c) 2014 Simon Eskildsen
// https://github.com/sirupsen/logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
