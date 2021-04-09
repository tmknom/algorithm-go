package main

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
)

type Stack struct {
	Values  []int
	Top     int
	MaxSize int
}

func NewStack(size int) *Stack {
	return &Stack{Values: make([]int, size), Top: 0, MaxSize: size}
}

func (s *Stack) IsEmpty() bool {
	return s.Top == 0
}

func (s *Stack) IsFull() bool {
	return s.Top == s.MaxSize
}

func (s *Stack) Push(value int) error {
	if s.IsFull() {
		return errors.New(fmt.Sprintf("stack is full: cannot push %v", value))
	}
	s.Values[s.Top] = value
	s.Top += 1
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return math.MinInt64, errors.New("stack is empty: cannot pop")
	}
	s.Top -= 1
	return s.Values[s.Top], nil
}

func (s *Stack) Print() {
	fmt.Printf("stack = [")
	for i := s.Top - 1; i >= 0; i-- {
		fmt.Printf("%#v, ", s.Values[i])
	}
	fmt.Println("]")
}

func main() {
	stack := NewStack(10)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Print()

	stack.Pop()
	stack.Pop()
	stack.Push(5)

	stack.Print()
}
