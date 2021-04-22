package util

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

func (s *Stack) IsEmpty() bool {
	return s.values.Len() == 0
}

func (s *Stack) Get(index int) int {
	return s.ToIntSlice()[index]
}

func (s *Stack) ToIntSlice() []int {
	result := make([]int, s.values.Len())
	element := s.values.Front()
	for i := 0; i < len(result); i++ {
		result[i] = element.Value.(int)
		element = element.Next()
	}
	return result
}

func (s *Stack) String() string {
	return fmt.Sprintf("%#v", s.ToIntSlice())
}

func (s *Stack) Print() {
	fmt.Printf("stack = %s\n", s)
}
