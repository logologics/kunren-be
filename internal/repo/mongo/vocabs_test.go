package mongo_test

import (
	"testing"

	. "github.com/logologics/kunren-be/internal/repo/mongo"
)

func TestMergeAndSort(t *testing.T) {
	a := []string{"x", "a", "k"}
	b := []string{"y", "a", "l", "m"}
	expected := []string{"a", "k", "l", "m", "x", "y"}

	res := MergeAndSort(a, b)
	if !Equal(res, expected) {
		t.Errorf("Expected %v, but got %v", expected, res)
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []string) bool {
	if len(a) != len(b) {
			return false
	}
	for i, v := range a {
			if v != b[i] {
					return false
			}
	}
	return true
}