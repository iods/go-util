package strutil

import "strings"

type (
	// Value is a strings value.
	Value string
	// StringValue is a strings value, an alias of Value.
	StringValue = Value
)

var (
	Equal = strings.EqualFold
)

// IsEmpty will check whether a string value is empty.
func (s *Value) IsEmpty() bool {
	return string(*s) == ""
}

func (s *Value) IsNumeric() bool {

}
