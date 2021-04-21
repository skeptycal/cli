// +build linux aix zos
// +build !js

package terminal

// the isTerminal functionality from logrus is used here.
// MIT Licence Copyright (c) 2014 Simon Eskildsen
// https://github.com/sirupsen/logrus

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TCGETS

func isTerminal(fd int) bool {
	_, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	return err == nil
}
