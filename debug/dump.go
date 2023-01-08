package debug

import (
	"encoding/json"
	"io"
)

// Dump is a formatted JSON output of a stacktrace.
func Dump(v ...interface{}) {

}

// Dumpf is a non-formatted dump of a JSON stacktrace?
//
// what is the io.Writer in this?
func Dumpf(w io.Writer, v ...interface{}) {
	e := json.NewEncoder(w)
	e.SetIndent("", " ")

	for _, x := range v {
		// add in some error handling here
		err := e.Encode(x)
		if err != nil {
			return
		}
	}
}
