package main

import "fmt"

// Basic queue data structure FIFO
type Queue struct {
	Values []interface{}
}

// init dynamic array with placeholder of any type and len of 0
func CreateNewQueue() *Queue {
	queue := Queue{}
	return &queue
}

// Get Item and queue it at the back
func (q *Queue) EnqueueNewValue(v interface{}) {
	q.Values = append(q.Values, v)
}

// Get Item at the front and dequeue it
func (q *Queue) DequeueOldestValue() interface{} {
	if q.IsQueueEmpty() == true {
		return "nil"
	}
	// <!> Full slice expressions
	dequeuedValue := q.PeekOldestValue()
	q.Values = q.Values[0:]
	return dequeuedValue
}

// Get Item at the front
func (q *Queue) PeekOldestValue() interface{} {
	return q.Values[0]
}

// Check if queue is empty
func (q *Queue) IsQueueEmpty() bool {
	// <!> Is this the correct var
	// <!> If possible to One-Liner
	if len(q.Values) == 0 {
		return true
	}
	return false
}

func main() {
	queue := CreateNewQueue()
	queue.EnqueueNewValue(5)
	queue.EnqueueNewValue('2')
	queue.EnqueueNewValue("text")
	fmt.Println("Peek:", queue.PeekOldestValue())
	fmt.Println("Pop:", queue.DequeueOldestValue())
	fmt.Println("Is empty?", queue.IsQueueEmpty())
}
