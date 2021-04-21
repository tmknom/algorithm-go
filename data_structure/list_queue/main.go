package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	values *list.List
}

func NewQueue() *Queue {
	return &Queue{values: list.New()}
}

func (q *Queue) Enqueue(value int) {
	q.values.PushBack(value)
}

func (q *Queue) Dequeue() int {
	return q.values.Remove(q.values.Front()).(int)
}

func (q *Queue) Print() {
	fmt.Printf("queue = [")
	element := q.values.Front()
	for {
		if element == nil {
			break
		}

		fmt.Printf("%#v, ", element.Value)
		element = element.Next()
	}
	fmt.Printf("]\n")
}

func main() {
	queue := NewQueue()

	queue.Enqueue(1)
	queue.Print()
	queue.Enqueue(2)
	queue.Print()
	queue.Enqueue(3)
	queue.Print()

	fmt.Printf("Dequeue = %d\n", queue.Dequeue())
	queue.Print()
	fmt.Printf("Dequeue = %d\n", queue.Dequeue())
	queue.Print()
	queue.Enqueue(5)
	queue.Print()
	fmt.Printf("Dequeue = %d\n", queue.Dequeue())
	queue.Print()
}
