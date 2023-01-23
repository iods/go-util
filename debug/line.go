package debug

// Fileline is a type that is what?
type Fileline int

const (
	FilelineOff  Fileline = 0
	ShowFileline Fileline = 1
)

var lineSwitch Fileline = FilelineOff

// SetFileline toggles the option to include the line number of the error.
func SetFileline(l Fileline) {
	lineSwitch = l
}
