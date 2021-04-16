package main

import (
	"../../util"
	"fmt"
)

func Frog(h []int) int {
	size := len(h)
	dp := util.InitInfIntSlice(size)

	dp[0] = 0
	for i := 1; i < size; i++ {
		dp[i] = dp[i-1] + util.IntAbs(h[i]-h[i-1])
		if i > 1 {
			double := dp[i-2] + util.IntAbs(h[i]-h[i-2])
			util.IntChooseMin(&dp[i], double)
		}
	}

	return dp[size-1]
}

func run(h []int) {
	result := Frog(h)
	fmt.Printf("Frog(%#v) = %v\n", h, result)
}

func main() {
	run([]int{2, 9, 4, 5, 1, 6, 10})
	run([]int{2, 9, 4, 15, 8, 21, 16, 5, 1, 6, 10})
}
