package basic

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	// enqueue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// size
	if q.Size() != 3 {
		t.Error()
	}

	// iterate
	for i := q.Iterator(); i.HasNext(); {
		fmt.Println(i.Next().(int))
	}

	// dequeue
	if q.Dequeue() != 1 {
		t.Error()
	}
	if q.Dequeue() != 2 {
		t.Error()
	}
	if q.Size() != 1 {
		t.Error()
	}
	for i := q.Iterator(); i.HasNext(); {
		fmt.Println(i.Next().(int))
	}

	// empty
	if q.Dequeue() != 3 {
		t.Error()
	}
	if q.Size() != 0 {
		t.Error()
	}
	for i := q.Iterator(); i.HasNext(); {
		fmt.Println(i.Next().(int))
	}
}
