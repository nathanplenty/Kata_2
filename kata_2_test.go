package main

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	stack := CreateNewQueue()
	b := stack.IsQueueEmpty()
	if b != true {
		t.Errorf("b = %v; want true", b)
	}
}

func TestEnqueuePeek(t *testing.T) {
	stack := CreateNewQueue()
	stack.EnqueueNewValue(5)
	v := stack.PeekOldestValue()
	if v != 5 {
		t.Errorf("v = %v; want 5", v)
	}
}

func TestDequeue(t *testing.T) {
	stack := CreateNewQueue()
	stack.EnqueueNewValue(5)
	stack.EnqueueNewValue(10)
	v := stack.DequeueOldestValue()
	if v != 5 {
		t.Errorf("v = %v; want 5", v)
	}
}
