package search

import "testing"

func TestRBTree(t *testing.T) {
	rbTree := NewRBTree()

	// put
	rbTree.Put("b", 2)
	rbTree.Put("a", 1)
	rbTree.Put("c", 3)
	rbTree.Put("d", 4)

	// get
	if v := rbTree.Get("c"); v.(int) != 3 {
		t.Error(v)
	}
	if v := rbTree.Get("d"); v.(int) != 4 {
		t.Error(v)
	}
	if v := rbTree.Get("f"); v != nil {
		t.Error(v)
	}

}
