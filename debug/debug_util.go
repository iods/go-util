package debug

import (
	"runtime"
	"strings"
	"time"
)

// do i need to add this to interfaces?
const threshold = time.Second * 3

// options for debugging utility at build
var (
	debug   bool
	verbose bool
	trace   bool
)

var prev time.Time

var (
// when sourcePath is set at compile time, relative paths are enabled
// taking out a prefix.
)

type Frame struct {
	File   string
	Path   string
	LineNo int
	Name   string
	Func   *runtime.Func
}

func Stacktrace() {

}

// isTrue checks whether a string is empty or not.
func isTrue(value ...string) bool {
	for _, e := range value {
		if strings.EqualFold("true", e) {
			return true
		}
		if strings.EqualFold("t", e) {
			return true
		}
		if strings.EqualFold("yes", e) {
			return true
		}
		if strings.EqualFold("y", e) {
			return true
		}
	}
	return false
}
