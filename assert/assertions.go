package assert

// Assertions is a struct for the Testing interface.
type Assertions struct {
	t    Testing
	last bool // asser result from previous test.
}

// New creates a new Assertions object for the specified Testing struct.
func New(t Testing) *Assertions {
	return &Assertions{
		t: t,
	}
}

// Tried for the last test.
func (a *Assertions) Tried() bool {
	return a.last
}

// Failed for the last test.
func (a *Assertions) Failed() bool {
	return !a.last
}
