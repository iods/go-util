package strutil

import "strings"

/*
Different functions for changing the case of a string.
TODO finish the docs for this package (formatting)
*/

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	if len(s) == 1 {
		return strings.ToUpper(s)
	}
}

// Title is the alias to strings.ToTitle() and returns?
func Title(s string) string {
	return strings.ToTitle(s)
}

// Lower is the alias to strings.ToLower() and returns?
func Lower(s string) string {
	return strings.ToLower(s)
}

// Lowercase is the alias to strings.ToLower() and returns?
func Lowercase(s string) string {
	return strings.ToLower(s)
}

// Upper is the alias to strings.ToUpper() and returns?
func Upper(s string) string {
	return strings.ToUpper(s)
}

// Uppercase is the alias to strings.ToUpper() and returns?
func Uppercase(s string) string {
	return strings.ToUpper(s)
}
