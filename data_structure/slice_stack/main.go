package main

import (
	"fmt"
)

type Stack struct {
	values []int
	top    int
	size   int
}

func NewStack(size int) *Stack {
	return &Stack{values: make([]int, 0, size), top: 0, size: size}
}

func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack) IsFull() bool {
	return s.top == s.size
}

func (s *Stack) Push(value int) {
	if s.IsFull() {
		panic(fmt.Sprintf("cannot push stack: %v => %#v", value, s))
	}
	s.values = append(s.values, value)
	s.top += 1

	// fmt.Printf("push(%d)\n", value)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic(fmt.Sprintf("cannot pop stack: %#v", s))
	}
	s.top -= 1
	result := s.values[s.top]
	s.values = s.values[0:s.top]

	// fmt.Printf("pop() = %d\n", result)
	return result
}

func (s *Stack) Print() {
	fmt.Printf("stack = %#v\n\n", s)
}

func main() {
	stack := NewStack(3)

	stack.Push(1)
	stack.Print()
	stack.Push(2)
	stack.Print()
	stack.Push(3)
	stack.Print()

	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Push(5)
	stack.Print()
	stack.Pop()
	stack.Print()
}
