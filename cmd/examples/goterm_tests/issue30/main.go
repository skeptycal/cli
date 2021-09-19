package main

import (
	"github.com/buger/goterm"
)

func main() {
	println("before")
	goterm.Println("foo")
	goterm.Flush()
	println("after", goterm.Width(), goterm.Height())
}
