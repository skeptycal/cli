package cli

import (
	"bytes"
	"io"
)

// Parser implements a parser for []byte values. Each parser
// generates output that has been processed by its unique
// process() method to produce the desired output.
//
// The most direct use case is to set src and dst to files
// and instantiate the Parser.
//
// Output may be obtained as a string or slice of bytes.
type Parser interface {
	io.ReadWriteCloser
	Bytes() []byte
	String() string
	ReadAll() ([]byte, error)
	WriteAll() (n int, err error)
}

func NewParser(src io.Reader, dst io.WriteCloser) Parser {
	p := &parser{src, dst, nil}
	p.process()
	return p
}

// parser reads []byte data from src, processes it, and writes
// it to dst.
type parser struct {
	src io.Reader
	dst io.WriteCloser
	buf *bytes.Buffer // buffered parser ...
}

// process is the place where various parsing and editing functionality
// can be added to make specific Parsers or Wrappers
//
// If an error occurs, the message is loaded into the
// returned []byte value.
//
// This is a hacky solution and is generally not recommended
// as it makes errors much more opaque. It is a trade-off.
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
func (w *parser) process() []byte {
	return nil
}

// ReadAll reads the entire contents of src and returns it (unchanged) as []byte.
func (w *parser) ReadAll() ([]byte, error) {
	return io.ReadAll(w.src)
}

// WriteAll writes the entire processed []byte to the destination.
func (w *parser) WriteAll() (n int, err error) {
	return w.dst.Write(w.Bytes())
}

// Bytes returns the wrapped bytes.
//
// If an error occurs, the message is loaded into the
// returned []byte value.
//
// This is a hacky solution and is generally not recommended
// as it makes errors much more opaque.
func (w *parser) Bytes() []byte {
	return w.process()
}

// String returns the wrapped string
func (w *parser) String() string {
	return string(w.Bytes())
}

func (w *parser) Close() error {
	return w.dst.Close()
}

// Read reads processed text directly from the underlying buffer.
func (w *parser) Read(p []byte) (n int, err error) {
	return w.src.Read(p)
}

// Write writes directly to the destination without modifications.
func (w *parser) Write(p []byte) (n int, err error) {
	return w.dst.Write(p)
}
