package sorting

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	source1 := []int{2, 4, 2, 6, 3, 7, 8, 4, 2}
	target1 := []int{2, 2, 2, 3, 4, 4, 6, 7, 8}
	SelectionSort(source1)
	if !equal(source1, target1) {
		t.Error(source1)
	}
}
