package search

import "testing"

func TestHashMap(t *testing.T) {
	m := NewHashMap()

	// put
	m.Put("haha", 123)
	m.Put("haha", 321)
	m.Put("abc", "xxx")

	// get
	if v := m.Get("haha"); v.(int) != 321 {
		t.Error(v)
	}
	if v := m.Get("abc"); v.(string) != "xxx" {
		t.Error(v)
	}

	// delete
	if m.Delete("ahah") {
		t.Error()
	}
	if !m.Delete("haha") {
		t.Error()
	}
	if v := m.Get("haha"); v != nil {
		t.Error(v)
	}
}
