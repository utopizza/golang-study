package sorting

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	source1 := []int{5, 4, 3, 2, 1}
	target1 := []int{1, 4, 3, 2, 5}
	swap(&source1, 0, len(source1)-1)
	if !equal(source1, target1) {
		t.Fail()
	}
	fmt.Println(source1)
}

func TestMin(t *testing.T) {
	if min(1, -1) != -1 {
		t.Fail()
	}
}

func TestShuffle(t *testing.T) {
	source1 := []int{5, 4, 3, 2, 1}
	target1 := []int{5, 4, 3, 2, 1}
	shuffle(&source1)
	if equal(source1, target1) {
		t.Fail()
	}
	fmt.Println(source1)
}
