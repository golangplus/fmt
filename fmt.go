package fmtp

import (
	"fmt"
	"io"
	"os"
)

// Printfln is similar to fmt.Printf but a newline is appended.
func Printfln(format string, a ...interface{}) (n int, err error) {
	return Fprintfln(os.Stdout, format, a...)
}

// Fprintfln is similar to fmt.Fprintf but a newline is appended.
func Fprintfln(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format+"\n", a...)
}
