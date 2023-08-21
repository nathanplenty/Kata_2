package main

import "fmt"

// Basic queue data structure FIFO
type Queue struct {
	// store one Values of any type
	Values []interface{}
}

func NewQueue() *Queue {
	// init dynamic array with placeholder of any type and len of 0
	queue := Queue{}
	return &queue
}

func (q *Queue) Enqueue(v interface{}) {
	// Get Item and queue it at the back
	q.Values = append(q.Values, v)
}

func (q *Queue) Dequeue() int {
	// Get Item at the front and dequeue it
	i := 0
	if q.IsEmpty() == true {
		return i
	}
	q.Values = q.Values[i:]
	return i
}

func (q *Queue) Peek() int {
	// Get Item at the front
	i := 0
	return i
}

func (q *Queue) IsEmpty() bool {
	// Check if queue is empty
	b := false
	// !! Is this the correct var? !!
	if len(q.Values) == 0 {
		b = true
	}
	return b
}

func main() {
	queue := NewQueue()
	queue.Enqueue(5)
	queue.Enqueue('2')
	queue.Enqueue("text")
	fmt.Println("Peek:", queue.Values[queue.Peek()])
	queue.Peek()
	queue.Dequeue()
	fmt.Println("Is empty?", queue.IsEmpty())
}
