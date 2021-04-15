package main

import (
	"../../util"
	"fmt"
)

func GCD(m int, n int, indent int) int {
	util.IndentPrintf(indent, "call GCD(%d, %d)\n", m, n)
	if n == 0 {
		util.IndentPrintf(indent, "return GCD(%d, %d) = %d\n", m, n, m)
		return m
	}

	result := GCD(n, m%n, indent+1)
	util.IndentPrintf(indent, "return GCD(%d, %d) = %d\n", m, n, result)
	return result
}

func main() {
	GCD(51, 15, 0)

	fmt.Println("")
	GCD(15, 51, 0)
}
