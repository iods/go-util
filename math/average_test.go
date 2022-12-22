package math

import (
	"testing"
)

func TestAverage(t *testing.T) {
	if Average(2, 2) != 2 {
		t.Error("Average of the two is incorrect, brother.")
	}
}
