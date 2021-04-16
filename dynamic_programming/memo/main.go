package main

import (
	"../../util"
	"fmt"
)

func Frog(i int, h []int, dp []int, indent int) int {
	if dp[i] != util.Inf {
		return dp[i]
	}

	util.CallPrintf(indent, "Frog(%d)\n", i)

	result := Frog(i-1, h, dp, indent+1) + util.IntAbs(h[i]-h[i-1])
	if i > 1 {
		double := Frog(i-2, h, dp, indent+1) + util.IntAbs(h[i]-h[i-2])
		util.IntChooseMin(&result, double)
	}

	util.ReturnPrintf(indent, "Frog(%d) = %d\n", i, result)
	dp[i] = result
	return result
}

func run(h []int) {
	dp := util.InitInfIntSlice(len(h))
	dp[0] = 0
	result := Frog(len(h)-1, h, dp, 0)
	fmt.Printf("Frog(%#v) = %v, dp = %#v\n", h, result, dp)
}

func main() {
	run([]int{2, 9, 4, 5, 1, 6, 10})

	fmt.Println("")
	run([]int{2, 9, 4, 15, 8, 21, 16, 5, 1, 6, 10})
}
