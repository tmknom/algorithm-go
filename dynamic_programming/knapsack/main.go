package main

import (
	"../../util"
	"fmt"
)

func Knapsack(threshold int, weights []int, values []int, dp [][]int) int {
	size := len(weights)
	for i := 0; i < size; i++ {
		for w := 0; w <= threshold; w++ {
			// i番目を選ばない
			dp[i+1][w] = dp[i][w]
			// i番目を選ぶ
			if threshold-weights[i] >= 0 {
				util.IntChooseMax(&dp[i+1][w], dp[i][threshold-weights[i]]+values[i])
			}
		}
	}

	return dp[size][threshold]
}

func run(threshold int, weights []int, values []int) {
	if len(weights) != len(values) {
		panic("len not equaled")
	}

	dp := make([][]int, len(weights)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, threshold+1)
	}

	result := Knapsack(threshold, weights, values, dp)
	fmt.Printf("Knapsack = %v\n", result)
	for i := 0; i < len(dp); i++ {
		fmt.Printf("dp[%d] = %#03v\n", i, dp[i])
	}
}

func main() {
	run(10, []int{2, 1, 3, 2, 1, 5}, []int{3, 2, 6, 1, 3, 85})
}
