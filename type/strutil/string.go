package strutil

import (
	"strings"
)

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

// Reverse takes a string and reverses it.
func Reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

// Match takes and compares two arrays of strings and compares them.
func Match(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	diff := make(map[string]bool)
	for _, dim := range a {
		diff[dim] = true
	}
	for _, dim := range b {
		if !diff[dim] {
			return false
		}
	}
	return true
}
