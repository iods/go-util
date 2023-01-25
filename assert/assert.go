// Package assert adds and extends default tool functions available for testing in Go.
package assert

// Testing wraps and extends the *testing.TB pointer.
type Testing interface {
	Helper()
	Name() string
	Error(args ...any)
}
