package debug

import "fmt"

type Color int32

const (
	None    Color = 0
	Green   Color = 1
	White   Color = 2
	Yellow  Color = 3
	Red     Color = 4
	Blue    Color = 5
	Magenta Color = 6
	Cyan    Color = 7
	Reset   Color = 8
)

const (
	BgNone    Color = 0
	GreenBg   Color = 11
	WhiteBg   Color = 12
	YellowBg  Color = 13
	RedBg     Color = 14
	BlueBg    Color = 15
	MagentaBg Color = 16
	CyanBg    Color = 17
)

var colors = map[Color]string{
	Green:   string([]byte{27, 91, 51, 50, 109}),
	White:   string([]byte{27, 91, 51, 55, 109}),
	Yellow:  string([]byte{27, 91, 51, 51, 109}),
	Red:     string([]byte{27, 91, 51, 49, 109}),
	Blue:    string([]byte{27, 91, 51, 49, 109}),
	Magenta: string([]byte{27, 91, 51, 49, 109}),
	Cyan:    string([]byte{27, 91, 51, 49, 109}),
	Reset:   string([]byte{27, 91, 48, 109}),
}

var bgcolors = map[Color]string{
	GreenBg:   string([]byte{27, 91, 57, 55, 59, 52, 50, 109}),
	WhiteBg:   string([]byte{27, 91, 57, 48, 59, 52, 55, 109}),
	YellowBg:  string([]byte{27, 91, 57, 48, 59, 52, 51, 109}),
	RedBg:     string([]byte{27, 91, 57, 55, 59, 52, 49, 109}),
	BlueBg:    string([]byte{27, 91, 57, 55, 59, 52, 52, 109}),
	MagentaBg: string([]byte{27, 91, 57, 55, 59, 52, 53, 109}),
	CyanBg:    string([]byte{27, 91, 57, 55, 59, 52, 54, 109}),
}

// OutputGreen returns the provided string in a Green text color with no background color.
func OutputGreen(msg string) string {
	return format(Green, msg)
}

// OutputWhite returns the provided string in a White text color with no background color.
func OutputWhite(msg string) string {
	return format(White, msg)
}

// OutputYellow returns the provided string in a Yellow text color with no background color.
func OutputYellow(msg string) string {
	return format(Yellow, msg)
}

// OutputRed returns the provided string in a Red text color with no background color.
func OutputRed(msg string) string {
	return format(Red, msg)
}

// OutputBlue returns the provided string in a Blue text color with no background color.
func OutputBlue(msg string) string {
	return format(Blue, msg)
}

// OutputMagenta returns the provided string in a Magenta text color with no background color.
func OutputMagenta(msg string) string {
	return format(Magenta, msg)
}

// OutputCyan returns the provided string in a Cyan text color with no background color.
func OutputCyan(msg string) string {
	return format(Cyan, msg)
}

// color returns the matching color value or an empty string.
func color(c Color) string {
	if cs, ok := colors[c]; ok {
		return cs
	}
	return ""
}

// bgcolor returns the matching bgcolors value or an empty string.
func bgcolor(c Color) string {
	if cs, ok := bgcolors[c]; ok {
		return cs
	}
	return ""
}

// format returns the output formatted with color (should this go in error formatting too?)
func format(c Color, msg string) string {
	return fmt.Sprintf("%s%s%s", color(c), msg, color(Reset))
}
