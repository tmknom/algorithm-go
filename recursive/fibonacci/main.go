package main

import (
	"../../util"
	"fmt"
)

func Fibonacci(x int, indent int) int {
	util.CallPrintf(indent, "Fibonacci(%d)\n", x)
	if x == 0 || x == 1 {
		util.ReturnPrintf(indent, "Fibonacci(%d) = %d\n", x, x)
		return x
	}

	result := Fibonacci(x-1, indent+1) + Fibonacci(x-2, indent+2)
	util.ReturnPrintf(indent, "Fibonacci(%d) = %d\n", x, result)
	return result
}

func main() {
	Fibonacci(5, 0)

	fmt.Println("")
	Fibonacci(6, 0)
}
