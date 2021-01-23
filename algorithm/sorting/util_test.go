package sorting

import (
	"testing"
)

func TestSwap(t *testing.T) {
	source1 := []int{5, 4, 3, 2, 1}
	target1 := []int{1, 4, 3, 2, 5}
	swap(source1, 0, len(source1)-1)
	if !equal(source1, target1) {
		t.Error(source1)
	}
}

func TestMin(t *testing.T) {
	if min(1, -1) != -1 {
		t.Error()
	}
}

func TestShuffle(t *testing.T) {
	source1 := []int{5, 4, 3, 2, 1}
	target1 := []int{5, 4, 3, 2, 1}
	shuffle(source1)
	if equal(source1, target1) {
		t.Error(source1)
	}
}
