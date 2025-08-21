package errors

import (
	"io"
)

// ErrorReporter provides an interface to report errors accross the interpreter, making it possible to
// swap out different reporting strategies.
type ErrorReporter interface {
	// Report reports the error in the GLox main struct, while writes the type of error, file and line where the error occurred to a io.Writer.
	Report(errorType string, file string, line string, w io.Writer)
}

// Error is the basic error type, which implements the ErrorReporter interface and writes to the standard input.
type Error struct { }

func (e *Error) Report(errorType string, file string, line string, w io.Writer) {

}