package search

import (
	"testing"
)

func TestBST(t *testing.T) {
	bst := NewBinarySearchTree()

	// put
	bst.Put("b", 2)
	bst.Put("a", 1)
	bst.Put("c", 3)
	bst.Put("d", 4)

	// get
	if v := bst.Get("c"); v.(int) != 3 {
		t.Error(v)
	}
	if v := bst.Get("d"); v.(int) != 4 {
		t.Error(v)
	}
	if v := bst.Get("f"); v != nil {
		t.Error(v)
	}

	// keys
	expectedKeys := []string{"a", "b", "c", "d"}
	actualKeys := bst.Keys()
	if equal(actualKeys, expectedKeys) {
		t.Error(actualKeys)
	}

	// min
	if k, _ := bst.Min(); k != "a" {
		t.Error(k)
	}

	// max
	if k, _ := bst.Max(); k != "d" {
		t.Error(k)
	}

	// detele
	bst.Delete("c")
	expectedKeys = []string{"a", "b", "d"}
	actualKeys = bst.Keys()
	if equal(actualKeys, expectedKeys) {
		t.Error(actualKeys)
	}
}

func equal(arr []string, target []string) bool {
	if len(arr) != len(target) {
		return false
	}
	for i := range arr {
		if arr[i] != target[i] {
			return false
		}
	}
	return true
}
