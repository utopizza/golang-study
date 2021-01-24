package search

import "testing"

func TestBST(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Put("b", 2)
	bst.Put("a", 1)
	bst.Put("c", 3)
	bst.Put("d", 4)
	if v := bst.Get("c"); v.(int) != 3 {
		t.Error(v)
	}
	if v := bst.Get("d"); v.(int) != 4 {
		t.Error(v)
	}
}
