package main

import (
	"fmt"
)

type Queue struct {
	values []int
	size   int
}

func NewDefaultQueue() *Queue {
	return NewQueue(1000)
}

func NewQueue(size int) *Queue {
	return &Queue{values: make([]int, 0, size), size: 0}
}

func (q *Queue) Enqueue(value int) {
	q.values = append(q.values, value)
	q.size += 1
}

func (q *Queue) Dequeue() int {
	if q.size == 0 {
		panic(fmt.Sprintf("cannot dequeue: %#v\n", q))
	}

	result := q.values[0]
	q.values = q.values[1:]
	q.size -= 1

	return result
}

func (q *Queue) Print() {
	fmt.Printf("queue = %#v\n", q)
}

func main() {
	stack := NewDefaultQueue()

	stack.Enqueue(1)
	stack.Enqueue(2)
	stack.Enqueue(3)
	stack.Print()

	stack.Dequeue()
	stack.Dequeue()
	stack.Print()

	stack.Enqueue(5)
	stack.Print()

	stack.Dequeue()
	stack.Print()
}
