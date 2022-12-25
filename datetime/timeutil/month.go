package timeutil

func ConvertDayToString(i uint16) string {
	day := []string{
		"first", "second", "third", "fourth", "fifth",
		"sixth", "seventh", "eighth", "ninth", "tenth",
		"eleventh", "twelfth", "thirteenth", "fourteenth", "fifteenth",
		"sixteenth", "seventeenth", "eighteenth", "nineteenth", "twentieth",
	}
	dayZero := []string{
		"tenth", "twentieth", "thirtieth", "fortieth", "fiftieth",
	}
	dayTen := []string{
		"ten", "twenty", "thirty", "forty", "fifty",
	}

	if i < 21 {
		return day[i]
	}
	if i > 59 {
		panic("E_OUT_OF_RANGE")
	}
	q, r := i/10, i%10
	if r == 0 {
		return dayZero[q-1]
	}

	return dayTen[q-1] + " " + day[r]
}
