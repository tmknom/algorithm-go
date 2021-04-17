package main

import (
	"../../util"
	"fmt"
)

func Frog(h []int) int {
	size := len(h)
	dp := util.SinglePositiveIntSlice

	dp.SetFirst(0)
	fmt.Printf("dp[%02d] = %02v\n", 0, dp[0])
	for i := 1; i < size; i++ {
		if i == 1 {
			dp[i] = util.IntAbs(h[i] - h[i-1])
			fmt.Printf("dp[%02d] = %02v\n", i, dp[i])
		} else {
			// ひとつ手前から移動する場合のコスト
			single := dp[i-1] + util.IntAbs(h[i]-h[i-1])
			// ふたつ手前から移動するコスト
			double := dp[i-2] + util.IntAbs(h[i]-h[i-2])
			// コストの最小値
			dp[i] = util.IntMin(single, double)

			fmt.Printf("dp[%02d] = %02v = min(%02v(=dp[%02d]+%02v), %02v(=dp[%02d]+%02v))\n",
				i, dp[i],
				single, i-1, util.IntAbs(h[i]-h[i-1]),
				double, i-2, util.IntAbs(h[i]-h[i-2]))
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

	fmt.Println("")
	run([]int{2, 9, 4, 15, 8, 21, 16, 5, 1, 6, 10})
}
