package basic

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()

	// push
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// size
	if s.Size() != 3 {
		t.Error()
	}

	// top
	if s.Top().(int) != 3 {
		t.Error()
	}

	// iterate
	for i := s.Iterator(); i.HasNext(); {
		fmt.Println(i.Next().(int))
	}

	// pop
	if s.Pop().(int) != 3 {
		t.Error()
	}
	if s.Pop().(int) != 2 {
		t.Error()
	}
	if s.Pop().(int) != 1 {
		t.Error()
	}
}
