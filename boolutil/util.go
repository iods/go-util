package boolutil

// Flip returns the opposite of the given boolean.
func Flip(value bool) bool {
	return !value
}

// ToInt returns the int equivalent of the given boolean.
func ToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}
