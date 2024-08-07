package logger

import (
	"io"
	"os"
)

var outWriter io.Writer = os.Stdout
var errWriter io.Writer = os.Stderr

// SetOutWriter sets the output destination for the standard logger.
func SetOutWriter(w io.Writer) {
	outWriter = w
}

// SetErrWriter sets the output destination for the standard logger.
func SetErrWriter(w io.Writer) {
	errWriter = w
}
