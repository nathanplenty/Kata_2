package main

import (
	"fmt"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	b := false
	stack := NewQueue()
	b = stack.IsEmpty()
	if b != true {
		t.Errorf("b = %v; want true", b)
	}
	fmt.Println(stack)
}

func TestEnqueuePeek(t *testing.T) {
	stack := NewQueue()
	stack.Enqueue(5)
	i, b := stack.Peek()
	v := stack.Value[i]
	if v != 5 {
		t.Errorf("v = %v; want 5", v)
	}
	if b != true {
		t.Errorf("b = %v; want true", b)
	}
	fmt.Println(stack)
}

func TestDequeue(t *testing.T) {
	stack := NewQueue()
	stack.Enqueue(5)
	stack.Enqueue(10)
	i, b := stack.Dequeue()
	v := stack.Value[i]
	if v != 5 {
		t.Errorf("v = %v; want 5", v)
	}
	if b != true {
		t.Errorf("b = %v; want true", b)
	}
	fmt.Println(stack)
}
