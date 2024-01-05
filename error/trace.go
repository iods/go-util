package error

import (
	"runtime/debug"
)

// GetTrace will return the current trace of the error
func (e *Error) GetTrace() string {
	if e.trace == "" {
		return string(debug.Stack())
	}
	return e.trace
}
