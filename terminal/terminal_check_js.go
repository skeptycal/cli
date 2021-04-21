// +build js

package terminal

// the isTerminal functionality from logrus is used here.
// MIT Licence Copyright (c) 2014 Simon Eskildsen
// https://github.com/sirupsen/logrus

func isTerminal(fd int) bool {
	return false
}
