package cli

import (
	"fmt"
	"io"
)

func (t *Terminal) Write(p []byte) (n int, err error) {
	// TODO - save current ansi colors

	// t.on()

	n, err = t.Writer.Write(p)
	if err != nil {
		return n, err
	}

	// TODO instead of Reset(() - restore saved ansi colors
	t.Reset()
	return
}

func (t *Terminal) WriteString(s string) (n int, err error) {
	return t.Write([]byte(s))
}

// Print wraps args in ANSI 8-bit color codes (256 color codes)
func (t *Terminal) Print(args ...interface{}) (n int, err error) {
	sum := 0
	n, err = fmt.Fprint(t.Writer, t.colorBytes)
	if err != nil {
		return 0, err
	}
	sum += n
	n, err = fmt.Fprint(t.Writer, args...)
	if err != nil {
		return sum, err
	}
	sum += n
	n, err = fmt.Fprint(t.Writer, Reset)
	if err != nil {
		return sum, err
	}
	return sum + n, nil
}

// Printf wraps args in ANSI 8-bit color codes (256 color codes)
func (t *Terminal) Printf(s string, args ...interface{}) (n int, err error) {
	sum := 0
	n, err = fmt.Fprint(t.Writer, t.colorBytes)
	if err != nil {
		return 0, err
	}
	sum += n
	n, err = fmt.Fprintf(t.Writer, s, args...)
	if err != nil {
		return sum, err
	}
	sum += n
	n, err = fmt.Fprint(t.Writer, Reset)
	if err != nil {
		return sum, err
	}
	return sum + n, nil
}

// Println wraps args in ANSI 8-bit color codes (256 color codes)
// and adds a newline character
func (t *Terminal) Println(args ...interface{}) (n int, err error) {
	sum := 0

	n, err = t.Print(args...)
	if err != nil {
		return 0, err
	}
	sum += n
	n, err = t.Print("\n")
	if err != nil {
		return sum, err
	}
	return sum + n, nil
}

func doCheckColor(w io.Writer, p []byte) (n int, err error) {
	return w.Write(p)
}

func noOp(w io.Writer, p []byte) (n int, err error) {
	return 0, nil
}
