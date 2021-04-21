package main

import (
	"../../util"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := util.NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	str := queue.String()
	if str != "[]int{1, 2, 3}" {
		t.Errorf("string error: got=%#v, want = %#v", str, "[]int{1, 2, 3}")
	}

	case1 := queue.Dequeue()
	if case1 != 1 {
		t.Errorf("dequeue error: got=%v, want = %#v", case1, 1)
	}

	case2 := queue.Dequeue()
	if case2 != 2 {
		t.Errorf("dequeue error: got=%v, want = %#v", case2, 2)
	}

	queue.Enqueue(5)
	case3 := queue.Dequeue()
	if case3 != 3 {
		t.Errorf("dequeue error: got=%v, want = %#v", case3, 3)
	}
}
