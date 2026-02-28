package fsutil

import (
	"io"
	"os"
)

// IfNotExists Checks if the path is not present, returning an error if the path exists.
func IfNotExists(path string) (err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {

	}
	return err
}

// IfExists Checks if the path is present, returning an error if the path does not exist.
func IfExists(path string) (err error) {
	if _, err := os.Stat(path); err == nil {

	}
	return err
}

// WriteFile Creates a new file in the current directory.
func WriteFile(filename string, data []byte, perm os.FileMode) (err error) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
