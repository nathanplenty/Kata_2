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
func (queue *Queue) EnqueueNewValue(v interface{}) {
	queue.Values = append(queue.Values, v)
}

// Get Item at the front and dequeue it
func (queue *Queue) DequeueOldestValue() interface{} {
	if queue.IsQueueEmpty() == true {
		return "nil"
	}
	// <!> Full slice expressions
	dequeuedValue := queue.PeekOldestValue()
	queue.Values = queue.Values[1:]
	return dequeuedValue
}

// Get Item at the front
func (queue *Queue) PeekOldestValue() interface{} {
	return queue.Values[0]
}

// Check if queue is empty
func (queue *Queue) IsQueueEmpty() bool {
	// <!> Is this the correct var
	// <!> If possible to One-Liner
	if len(queue.Values) == 0 {
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
