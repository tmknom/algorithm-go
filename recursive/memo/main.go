package main

import (
	"../../util"
	"fmt"
)

var memo []int

const InitialValue = -1

func Fibonacci(x int, indent int) int {
	if memo[x] != InitialValue {
		return memo[x]
	}

	util.CallPrintf(indent, "Fibonacci(%d)\n", x)
	result := Fibonacci(x-1, indent+1) + Fibonacci(x-2, indent+2)
	util.ReturnPrintf(indent, "Fibonacci(%d) = %d\n", x, result)
	return result
}

func initMemo(x int) {
	memo = make([]int, x+1)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i < len(memo); i++ {
		memo[i] = InitialValue
	}
}

func run(x int) {
	initMemo(x)
	Fibonacci(x, 0)
}

func main() {
	run(5)

	fmt.Println("")
	run(6)
}
