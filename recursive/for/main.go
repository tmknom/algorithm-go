package main

import (
	"fmt"
	"strconv"
)

func Fibonacci(x int) int {
	buf := make([]int, x+1)
	buf[0] = 0
	buf[1] = 1

	bufstr := "0, 1"
	for i := 2; i < len(buf); i++ {
		buf[i] = buf[i-1] + buf[i-2]
		bufstr += ", " + strconv.Itoa(buf[i])
	}

	result := buf[len(buf)-1]
	fmt.Printf("Fibonacci(%d) = %d, [%v]\n", x, result, bufstr)
	return result
}

func main() {
	Fibonacci(5)
	Fibonacci(10)
}
