package randstring

import (
	"testing"
)

func TestCreate(t *testing.T) {
	lengths := []int{1, 16, 42, 1337}

	for _, length := range lengths {
		s := Create(length)
		if len(s) != length {
			t.Errorf("String length should be %d, but is %d", length, len(s))
		}
	}
}
