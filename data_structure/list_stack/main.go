package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	values *list.List
}

func NewStack() *Stack {
	return &Stack{values: list.New()}
}

func (s *Stack) Push(value int) {
	s.values.PushBack(value)
}

func (s *Stack) Pop() int {
	return s.values.Remove(s.values.Back()).(int)
}

func (s *Stack) Print() {
	fmt.Printf("stack = [")
	element := s.values.Front()
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
	stack := NewStack()

	stack.Push(1)
	stack.Print()
	stack.Push(2)
	stack.Print()
	stack.Push(3)
	stack.Print()

	fmt.Printf("pop = %d\n", stack.Pop())
	stack.Print()
	fmt.Printf("pop = %d\n", stack.Pop())
	stack.Print()
	stack.Push(5)
	stack.Print()
	fmt.Printf("pop = %d\n", stack.Pop())
	stack.Print()
}
