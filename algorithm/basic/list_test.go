package basic

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := NewList()

	// add
	l.Add(1)
	l.Add(2)
	l.Add(3)

	// size
	if l.Size() != 3 {
		t.Error()
	}

	// iterate
	for i := l.Iterator(); i.HasNext(); {
		data := i.Next()
		fmt.Println(data.(int))
	}
}
