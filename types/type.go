package types

import (
	"fmt"
	"io"
)

// BoolCheckFunc will check a boolean functions type.
type BoolCheckFunc func(v bool) error

// ByteStringWriter interface.
type ByteStringWriter interface {
	io.Writer
	io.ByteWriter
	io.StringWriter
	fmt.Stringer
}

// Errors type is available for a multi-error list.
type Errors []error

// Float64able interface
type Float64able interface {
	Float64() (float64, error)
}

// FormatData provides an interface for data formatting.
type FormatData interface {
	Format() string
	FormatTo(w io.Writer)
}

// FormatDefault struct for data formatting.
type FormatDefault struct {
	// ow ByteStringWriter
	// Out formatted to the writer
	Out io.Writer
	// Src data(array, map, struct) for format
	Src any
	// MaxDepth limits the depth of an array
	MaxDepth int
	// Prefix string for an element
	Prefix string
	// Indent a string for each element
	Indent string
	// Suffix add a string for the last element "]", "}"
	Suffix string
}

// Int interface type.
type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Int64able interface type.
type Int64able interface {
	Int64() (int64, error)
}

// IntCheckFunc will check an integer functions type.
type IntCheckFunc func(v int64) error

type (
	// MarshalFunc definition.
	MarshalFunc func(v any) ([]byte, error)

	// UnmarshalFunc definition.
	UnmarshalFunc func(v any) ([]byte, error)
)

// Matcher interface
type Matcher[T any] interface {
	Match(s T) bool
}

// MatchFunc definition implementing the Matcher interface
type MatchFunc[T any] func(v T) bool

// SafeStringFunc will safe convert any value to a string.
type SafeStringFunc func(v any) string

// StringCheckFunc will check a string functions type.
type StringCheckFunc func(v string) error

// StringHandler interface
type StringHandler interface {
	HandleString(s string) string
}

// StringHandlerFunc definition
type StringHandlerFunc func(s string) string

// StringMatcher interface
type StringMatcher interface {
	Match(s string) bool
}

// StringWriteStringer interface.
type StringWriteStringer interface {
	io.StringWriter
	fmt.Stringer
}

// ToStringFunc will convert any value to a string and return an error on failure.
type ToStringFunc func(v any) (string, error)

// ToTypeFunc converts a value to a defined type.
type ToTypeFunc[T any] func(v any) (T, error)
