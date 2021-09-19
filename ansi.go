package cli

import (
	"errors"
	"fmt"
	"strconv"
)

// ANSI implements the interface for ANSI encoded values typically used
// for control and output in CLI applications.
type ANSI interface {
	String() string
}

func NewAnsiColor(in byte) ANSI {
	return &AnsiColor{in, ""}
}

// AnsiColor is a buffered, encoded ANSI color string typically used
// for CLI output. The encoded ANSI color code string is JIT
// buffered at the time of the first output request to eliminate
// repeated fmt.Sprintf (or similar) calls.
type AnsiColor struct {
	Color byte
	out   string
}

// String returns the ANSI formatted string representation of the AnsiColor byte.
func (a AnsiColor) String() string {
	if a.out == "" {
		a.out = BasicEncode(a.Color)
	}
	return a.out
}

// SimpleANSI produces an unbuffered ANSI encoded string typically used for terminal
// control and output in CLI applications.
type SimpleANSI struct {
	c byte
}

func (s *SimpleANSI) String() string {
	// TODO: is this faster? by a noticeable amount?
	// return ansiPrefix + string([]byte{b}) + ansiSuffix
	return fmt.Sprintf(fmtANSI, s.c)
}

func NewSimpleANSI(c byte) ANSI {
	return &SimpleANSI{c}
}

// BasicEncode encodes a basic (3-4 bit) ANSI color code.
// The code may be passed in as a byte, bool, int, uint, float32,
// float64, []byte or string and will be converted to the best
// guess of a byte (uint8) value before encoding.
//
// It is best to simply use a byte value to avoid confusing errors.
//
// The format is "\x1b[%dm"
func BasicEncode(b interface{}) string {

	switch v := b.(type) {
	case byte:
		NewSimpleANSI(v).String()
	default:
		c, err := ByteEncode(b)
		if err != nil {
			return ""
		}
		return NewSimpleANSI(c).String()
	}

	return ""
}

// ByteEncode returns the best guess encoding of the given
// interface as a byte value.
//
// -- byte values return the underlying byte value
//
// -- bool values return 0 or 1
//
// -- ints and floats are truncated to a single byte (this could be a likely source of confusing errors)
//
// -- []byte and string values are parsed for numerical values; if none is detected, 0 and an error are returned.
//
// -- other values return a value of 0 and an error
func ByteEncode(b Any) (byte, error) {
	switch v := b.(type) {
	case byte:
		return v, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		return byte(v & 255), nil
	case uint:
		return byte(v & 255), nil
	case float32, float64:
		return byte(v.(float64)), nil
	case []byte:
		i, err := strconv.ParseUint(string(v), 10, 8)
		if err != nil {
			return 0, err
		}
		return byte(i), nil
	case string:
		i, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return 0, err
		}
		return byte(i), nil
	default:
		return 0, errors.New("encoding error: unable to convert input to byte")
	}
}
