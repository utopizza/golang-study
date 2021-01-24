package sorting

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	source1 := []int{5, 4, 3, 2, 1}
	target1 := []int{1, 2, 3, 4, 5}
	InsertionSort(source1)
	if !equal(source1, target1) {
		t.Error(source1)
	}
}
