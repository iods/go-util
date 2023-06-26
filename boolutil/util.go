package boolutil

// Flip returns the opposite of the given boolean.
func Flip(b bool) bool {
	return !b
}

// ToInt returns the int equivalent of the given boolean.
func ToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
