package strutil

import "strings"

// UcFirst capitalizes the first rune of a string, and makes the remaining runes lowercase.
func UcFirst(s string) string {
	chars := []rune(s)
	first := ""
	if len(chars) > 0 {
		first = strings.ToUpper(string(chars[0]))
	}
	last := ""
	if len(chars) > 0 {
		last = strings.ToLower(string(chars[1:]))
	}
	return first + last
}
