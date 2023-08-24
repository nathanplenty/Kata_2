package main

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	queue := CreateNewQueue()
	boolean := queue.IsQueueEmpty()
	if boolean != true {
		t.Errorf("boolean = %v; want true", boolean)
	}
}

func TestEnqueuePeek(t *testing.T) {
	queue := CreateNewQueue()
	queue.EnqueueNewValue(5)
	value := queue.PeekOldestValue()
	if value != 5 {
		t.Errorf("value = %v; want 5", value)
	}
}

func TestDequeue(t *testing.T) {
	queue := CreateNewQueue()
	queue.EnqueueNewValue(5)
	queue.EnqueueNewValue(10)
	value := queue.DequeueOldestValue()
	if value != 5 {
		t.Errorf("value = %v; want 5", value)
	}
}

func TestForNil(t *testing.T) {
	queue := CreateNewQueue()
	value := queue.PeekOldestValue()
	if value != nil {
		t.Errorf("value = %v; want <nil>", value)
	}
}
