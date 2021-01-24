package search

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	// put
	trie.Put("hello", 100)
	trie.Put("hexo", 50)
	trie.Put("l a u g h", "hahaha")

	// get
	if v := trie.Get("hello"); v.(int) != 100 {
		t.Error(v)
	}
	if v := trie.Get("l a u g h"); v.(string) != "hahaha" {
		t.Error(v)
	}
	if v := trie.Get("laugh"); v != nil {
		t.Error(v)
	}

	// keysWithPrefix
	keys1 := trie.KeysWithPrefix("h")
	if len(keys1) != 2 {
		t.Error(keys1)
	}

	// keys
	keys2 := trie.Keys()
	if len(keys2) != 3 {
		t.Error(keys2)
	}

	// delete
	trie.Delete("hello")
	keys3 := trie.Keys()
	if len(keys3) != 2 {
		t.Error(keys3)
	}
}
