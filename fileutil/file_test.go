package fileutil

import "testing"

func TestIfNotExists(t *testing.T) {
	err := IfNotExists(".config")
	if err != nil {
		t.Error("it looks like it did not succeed.")
	}
}

func TestIfExists(t *testing.T) {
	err := IfExists(".config/eddie")
	if err != nil {
		t.Error("it looks like it did not succeed.")
	}
}
