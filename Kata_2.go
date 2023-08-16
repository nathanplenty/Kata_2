package main

import "fmt"

// Basic queue data structure FIFO

type Queue struct {
	// store one value of any type
	Value []interface{}
}

func NewQueue() *Queue {
	// init dynamic array with placeholder of any type and len of 0
	queue := Queue{make([]interface{}, 0)}
	fmt.Println("New queue created!")
	return &queue
}

func (q *Queue) Enqueue(v interface{}) {
	// Get Item and queue it at the back
	q.Value = append(q.Value, v)
	fmt.Println("Value [", v, "] added in queue!")
}

func (q *Queue) Dequeue() (int, bool) {
	// Get Item at the front and queue it
	return 0, false
}

func (q *Queue) Peek() (int, bool) {
	// Get Item at the front

	return 0, false
}

func (q *Queue) IsEmpty() bool {
	// Check if queue is empty
	b := false
	// !! Is this the correct var? !!
	if len(q.Value) == 0 {
		b = true
	}
	return b
}

func main() {
	queue := NewQueue()
	queue.Enqueue(5)
	queue.Enqueue(10)
	queue.Enqueue(15)
	val, _ := queue.Peek()
	fmt.Println(val) // Output: 5
	val, _ = queue.Dequeue()
	fmt.Println(val)                          // Output: 5
	fmt.Println("Is empty?", queue.IsEmpty()) // Output: false
}
