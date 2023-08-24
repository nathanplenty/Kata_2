package main

import "fmt"

type Queue struct {
	Values []interface{}
}

func CreateNewQueue() *Queue {
	queue := Queue{}
	return &queue
}

func (queue *Queue) EnqueueNewValue(value interface{}) {
	queue.Values = append(queue.Values, value)
}

func (queue *Queue) DequeueOldestValue() interface{} {
	if queue.IsQueueEmpty() == true {
		return nil
	}
	dequeuedValue := queue.PeekOldestValue()
	lengthOfValue := len(queue.Values)
	queue.Values = queue.Values[1:lengthOfValue:lengthOfValue]
	return dequeuedValue
}

func (queue *Queue) PeekOldestValue() interface{} {
	if queue.IsQueueEmpty() == true {
		return nil
	}
	return queue.Values[0]
}

func (queue *Queue) IsQueueEmpty() bool {
	return queue.Values == nil
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
