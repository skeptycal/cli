package cli

import (
	"fmt"
	"io"
)

// Write writes len(p) bytes from p to the Terminal data stream.
// The bytes are wrapped in ansi escape codes using the current
// Terminal colorBytes field and the Reset constant. This ensures
// that the terminal is not left in an unknown state if other
// programs write to it concurrently.
//
// It returns the number of bytes written from p (0 <= n <= len(p)).
// The bytes sent to set and reset ANSI escape codes are not
// included in the returned value. This maintains compatibility
// with the io.Writer interface.
//
// Any error encountered that caused the write to stop early is
// also returned.
// Write returns io.ErrShortWrite if n < len(p) and no other
// explicit error was identified.
//
// As specified by io.Writer, Write does not modify the slice
// data, even temporarily and does not retain p.
func (t *Terminal) Write(p []byte) (n int, err error) {
	// TODO - save current ansi colors??

	_, err = t.w.Write(t.colorBytes)
	if err != nil {
		return 0, err
	}

	// this value of n is the one that is returned to provide
	// consistant behavor for the io.Writer interface.
	n, err = t.w.Write(p)
	if err != nil {
		return n, err
	}

	// TODO instead of Reset(() - restore saved ansi colors??

	_, err = t.w.Write([]byte(Reset))

	if n < len(p) {
		return n, io.ErrShortWrite
	}

	if err != nil {
		return n, err
	}

	return n, nil
}

// WriteString writes the contents of s to the underlying data stream.
//
// It uses the io.Writer interface and follows standard conventions:
//
// It returns the number of bytes written from s (0 <= n <= len(s))
// and any error encountered that caused the write to stop early.
// WriteString must return a non-nil error if it returns n < len(s).
// Write must not modify the string data, even temporarily.
//
// Implementations must not retain s.
func (t *Terminal) WriteString(s string) (n int, err error) {
	// if sw, ok := t.w.(io.StringWriter); ok {
	// 	return sw.WriteString(s)
	// }
	return t.Write([]byte(s))
}

// Print wraps args in ANSI 8-bit color codes (256 color codes)
func (t *Terminal) Print(args ...interface{}) (n int, err error) {
	sum := 0
	n, err = fmt.Fprint(t.w, t.colorBytes)
	if err != nil {
		return 0, err
	}
	sum += n
	n, err = fmt.Fprint(t.w, args...)
	if err != nil {
		return sum, err
	}
	sum += n
	n, err = fmt.Fprint(t.w, Reset)
	if err != nil {
		return sum, err
	}
	return sum + n, nil
}

// Printf wraps args in ANSI 8-bit color codes (256 color codes)
func (t *Terminal) Printf(s string, args ...interface{}) (n int, err error) {
	sum := 0
	n, err = fmt.Fprint(t.w, t.colorBytes)
	if err != nil {
		return 0, err
	}
	sum += n
	n, err = fmt.Fprintf(t.w, s, args...)
	if err != nil {
		return sum, err
	}
	sum += n
	n, err = fmt.Fprint(t.w, Reset)
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

func (t *Terminal) colorWrite(w io.Writer, p []byte) (n int, err error) {
	return t.Write(p)
}

func (t *Terminal) noOp(w io.Writer, p []byte) (n int, err error) {
	return 0, nil
}
