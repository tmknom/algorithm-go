package util

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

func (q *Queue) IsEmpty() bool {
	return q.values.Len() == 0
}

func (q *Queue) Get(index int) int {
	return q.ToIntSlice()[index]
}

func (q *Queue) ToIntSlice() []int {
	result := make([]int, q.values.Len())
	element := q.values.Front()
	for i := 0; i < len(result); i++ {
		result[i] = element.Value.(int)
		element = element.Next()
	}
	return result
}

func (q *Queue) String() string {
	return fmt.Sprintf("%#v", q.ToIntSlice())
}

func (q *Queue) Print() {
	fmt.Printf("queue = %s\n", q)
}
