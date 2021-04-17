package main

import (
	"../../util"
	"fmt"
)

func Vacation(inputs [][]int) int {
	dp := util.DoubleNegativeIntSlice()

	// 初日のコストを初期値としてもたせる
	for i := 0; i < len(inputs[0]); i++ {
		dp[0][i] = inputs[0][i]
	}

	// 初日を計算済みなので二日目以降を計算
	days := len(inputs)
	for i := 1; i < days; i++ {
		util.IntChooseMax(&dp[i][0], inputs[i][0]+dp[i-1][1])
		util.IntChooseMax(&dp[i][0], inputs[i][0]+dp[i-1][2])

		util.IntChooseMax(&dp[i][1], inputs[i][1]+dp[i-1][0])
		util.IntChooseMax(&dp[i][1], inputs[i][1]+dp[i-1][2])

		util.IntChooseMax(&dp[i][2], inputs[i][2]+dp[i-1][0])
		util.IntChooseMax(&dp[i][2], inputs[i][2]+dp[i-1][1])
	}

	// 最終日の最大値を取得
	result := 0
	actions := len(inputs[0])
	for i := 0; i < actions; i++ {
		util.IntChooseMax(&result, dp[days-1][i])
	}

	return result
}

func run(inputs [][]int) {
	result := Vacation(inputs)
	fmt.Printf("Vacation(%#v) = %v\n", inputs, result)
}

func main() {
	run([][]int{
		{10, 40, 70},
		{20, 50, 80},
		{30, 60, 90},
	}) // => 210

	fmt.Println("")
	run([][]int{
		{100, 10, 1},
	}) // => 100

	fmt.Println("")
	run([][]int{
		{6, 7, 8},
		{8, 8, 3},
		{2, 5, 2},
		{7, 8, 6},
		{4, 6, 8},
		{2, 3, 4},
		{7, 5, 1},
	}) // => 46
}
