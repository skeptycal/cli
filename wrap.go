package cli

import (
	"bytes"
	"io"
	"strings"
)

const (
	NotNewLine = " \t\r\n\v"
	Vowels     = "aeiouy"
)

func NewBasicWrapper(src io.Reader, dst io.WriteCloser) Wrapper {
	return &wrapper{&parser{src, dst, nil}, 0} // , &bytes.Buffer{}}
}

type Wrapper interface {
	Parser
	Wrap(width int) string
}

type wrapper struct {
	*parser
	width int
}

// Wrap is used to set a new width and return the wrapped text.
// The result is not saved to dst until Flush() is called.
func (w *wrapper) Wrap(width int) string {
	w.width = width
	return w.String()
}

func (w *wrapper) save() error {
	n, err := w.dst.Write(w.buf.Bytes())
	if err != nil {
		return err
	}

	if n != w.buf.Len() {
		return io.ErrShortWrite
	}

	return nil
}

func (w *wrapper) Flush() error {
	return w.save()
}

// process is the place where various parsing and editing functionality
// can be added to make specific Parsers or Wrappers
//
// If an error occurs, the message is loaded into the
// returned []byte value.
//
// This is a hacky solution and is generally not recommended
// as it makes errors much more opaque.
//
// The addition of a buffer instead of the data variable may result
// in performance increases in some cases where the same data set is
// used as the source for several different operations. e.g.
//  if w.buf.Len() == 0 {
//  	data, err := w.load()
//  	if err := nil {...}
//      w.buf = wrapBytes(data, w.width)
//  }
//  return w.buf.Bytes()
func (w *wrapper) process() []byte {
	data, err := w.ReadAll()
	if err != nil {
		return []byte(err.Error())
	}

	w.buf.ReadFrom(w.src)
	return wrapBytes(data, w.width)
}

type filewrapper struct {
	wrapper
	fileName string
	buf      *bytes.Buffer
}

// Wrap splits a string into lines no longer than width.
func Wrap(s string, width int) string {
	// TODO - test differences between string and bytes processing
	return wrapString(s, width)
}

func wrapString(s string, width int) string {

	// if the string is within width, return it immediately
	if len(s) <= width {
		return s
	}

	sb := strings.Builder{}
	defer sb.Reset()

	var a int = 0
	var b int
	var nl int
	var useVowel string = ""
	var substr string = ""

	for {
		b = a + width

		substr = s[a:b]

		// if there is no newline in the substring, search
		// for another breakpoint
		if nl = strings.LastIndexByte(substr, newline); nl == -1 {

			// if no newline, find other whitespace
			if nl = strings.LastIndexAny(substr, NotNewLine); nl == -1 {

				useVowel = "-"
				nl = b
			}
		}
		sb.WriteString(strings.TrimSpace(s[a : a+nl]))
		sb.WriteString(useVowel)
		sb.WriteByte(newline)

		// set new substring start and end
		a += nl + 1

		// if the remaining string length is less than width, write
		// the final substring and return
		if len(s)-a <= width {
			sb.WriteString(strings.TrimSpace(s[a:]))
			// sb.WriteByte(newline)
			return sb.String()
		}
	}
}

func wrapBytes(s []byte, width int) []byte {

	if len(s) <= width {
		return s
	}

	sb := bytes.NewBuffer(s)
	defer sb.Reset()

	var a int = 0
	var b int
	var nl int
	var useVowel string = ""
	var substr []byte = []byte{}

	for {
		b = a + width

		substr = s[a:b]

		if nl = bytes.LastIndexByte(substr, newline); nl == -1 {
			// if there is no newline in the substring, search
			// for another breakpoint
			// if no newline, find other whitespace
			if nl = bytes.LastIndexAny(substr, NotNewLine); nl == -1 {

				useVowel = "-"
				nl = b
			}
		}
		sb.Write(bytes.TrimSpace(s[a : a+nl]))
		sb.WriteString(useVowel)
		sb.WriteByte(newline)

		// set new substring start and end
		a += nl + 1

		// if the remaining string length is less than width, write
		// the final substring and return
		if len(s)-a <= width {
			sb.Write(bytes.TrimSpace(s[a:]))
			// sb.WriteByte(newline)
			return sb.Bytes()
		}
	}
}
