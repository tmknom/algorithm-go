package main

import (
	"../../util"
	"testing"
)

func TestStack(t *testing.T) {
	stack := util.NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	str := stack.String()
	if str != "[]int{1, 2, 3}" {
		t.Errorf("string error: got=%#v, want = %#v", str, "[]int{1, 2, 3}")
	}

	case1 := stack.Pop()
	if case1 != 3 {
		t.Errorf("pop error: got=%v, want = %#v", case1, 3)
	}

	case2 := stack.Pop()
	if case2 != 2 {
		t.Errorf("pop error: got=%v, want = %#v", case2, 2)
	}

	stack.Push(5)

	case3 := stack.Pop()
	if case3 != 5 {
		t.Errorf("pop error: got=%v, want = %#v", case3, 5)
	}
}
