package main

import (
	"fmt"
)

type Queue struct {
	Values []int
	Head   int
	Tail   int
}

func NewQueue(size int) *Queue {
	return &Queue{Values: make([]int, size), Head: 0, Tail: 0}
}

func (q *Queue) IsEmpty() bool {
	return q.Head == q.Tail
}

func (q *Queue) IsFull() bool {
	return q.Head == (q.Tail+1)%q.MaxSize()
}

func (q *Queue) MaxSize() int {
	return len(q.Values)
}

func (q *Queue) Enqueue(value int) {
	if q.IsFull() {
		panic(fmt.Sprintf("queue is full: cannot enqueue %#v\n", q))
	}

	q.Values[q.Tail] = value
	if q.Tail+1 == q.MaxSize() {
		q.Tail = 0
	} else {
		q.Tail += 1
	}

	fmt.Printf("enqueue(%d): ", value)
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		panic(fmt.Sprintf("queue is empty: cannot dequeue %#v\n", q))
	}

	result := q.Values[q.Head]
	q.Values[q.Head] = 0
	if q.Head+1 == q.MaxSize() {
		q.Head = 0
	} else {
		q.Head += 1
	}

	fmt.Printf("dequeue():  ")

	return result
}

func (q *Queue) Print() {
	fmt.Printf("queue = %#v, head = %d, tail = %d\n", q.Values, q.Head, q.Tail)
}

func main() {
	stack := NewQueue(10)
	stack.Enqueue(1)
	stack.Print()
	stack.Enqueue(2)
	stack.Print()
	stack.Enqueue(3)
	stack.Print()

	stack.Dequeue()
	stack.Print()
	stack.Dequeue()
	stack.Print()
	stack.Enqueue(5)
	stack.Print()
}
