package main

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	b := false
	stack := NewQueue()
	b = stack.IsEmpty()
	if b != true {
		t.Errorf("b = %v; want true", b)
	}
}

func TestEnqueuePeek(t *testing.T) {
	stack := NewQueue()
	stack.Enqueue(5)
	i := stack.Peek()
	v := stack.Values[i]
	if v != 5 {
		t.Errorf("v = %v; want 5", v)
	}
}

func TestDequeue(t *testing.T) {
	stack := NewQueue()
	stack.Enqueue(5)
	stack.Enqueue(10)
	i := stack.Dequeue()
	v := stack.Values[i]
	if v != 5 {
		t.Errorf("v = %v; want 5", v)
	}
}
