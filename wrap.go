package cli

import (
	"strings"
)

const (
	NotNewLine = " \t\r\n\v"
	Vowels     = "aeiouy"
)

// Wrap splits a string into lines no longer than width.
func Wrap(s string, width int) string {

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

		// if there is a newline in the substring, use it as a break location
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

	// return sb.String()
}
