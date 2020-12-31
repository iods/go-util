package util

/*
Add Returns the sum (+) of two numbers for testing tests. */
func Add(a float64, b float64) float64 {
	return a + b
}

/*
Subtract Returns the difference (-) of two numbers for testing tests. */
func Subtract(a float64, b float64) float64 {
	return a - b
}

/*
Multiply Returns the product (*) of two numbers. */
func Multiply(a float64, b float64) float64 {
	return a * b
}

/*
Divide Returns the quotient (/) of two numbers. */
func Divide(a float64, b float64) float64 {
	return a / b
}

/*
Remainder Returns the remainder (%) after dividing two numbers. */
func Remainder(a int, b int) int {
	return a % b
}

/*
Negate Returns the expression with it's sign negated.
*/
func Negate(a float64) float64 {
	return -(a)
}