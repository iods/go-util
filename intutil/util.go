package intutil

// Absolute takes a number and returns its absolute value, returning its positive magnitude regardless of sign.
func Absolute(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
