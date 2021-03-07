package fileutil

import (
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
