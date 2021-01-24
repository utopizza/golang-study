package search

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	if found := BinarySearch(arr, 2); found != 1 {
		t.Error(found)
	}
	if found := BinarySearch(arr, 0); found != -1 {
		t.Error(found)
	}
}
