package sorting

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	source1 := []int{6, 4, 5, 3, 2, 1}
	target1 := []int{1, 2, 3, 4, 5, 6}
	QuickSort(source1)
	if !equal(source1, target1) {
		t.Error(source1)
	}
}
