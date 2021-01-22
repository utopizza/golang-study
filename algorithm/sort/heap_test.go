package sorting

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	source1 := []int{5, 4, 3, 2, 1}
	target1 := []int{1, 2, 3, 4, 5}
	heapSort(&source1)
	if !equal(source1, target1) {
		t.Fail()
	}
	fmt.Println(source1)
}
