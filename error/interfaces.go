package error

// as interface will cast an error to an interface
type as interface {
	As(interface{}) bool
}

// is interface will check if the error matches what is provided
type is interface {
	Is(err error) bool
}

// unwrap interface is used to unwrap the error
type unwrap interface {
	Unwrap() error
}

// wrap interface is used to wrap the error
type wrap interface {
	Wrap(err error) error
}
