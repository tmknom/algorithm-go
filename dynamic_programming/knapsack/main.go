package main

import (
	"../../util"
	"fmt"
)

func Knapsack(maxWeight int, inputs [][]int) int {
	sumValues := util.DoubleNegativeIntSlice()
	for i := 0; i <= maxWeight; i++ {
		sumValues[0][i] = 0
	}

	for i := 1; i < len(inputs)+1; i++ {
		weight := inputs[i-1][0]
		value := inputs[i-1][1]
		for sumWeight := 0; sumWeight <= maxWeight; sumWeight++ {
			if sumWeight-weight >= 0 {
				util.IntChooseMax(&sumValues[i][sumWeight], value+sumValues[i-1][sumWeight-weight])
			}
			util.IntChooseMax(&sumValues[i][sumWeight], sumValues[i-1][sumWeight])
		}
	}

	fmt.Println(sumValues.Dump(len(inputs)+1, maxWeight+1))

	return sumValues[len(inputs)][maxWeight]
}

func main() {
	inputs := [][]int{
		{2, 3},
		{1, 2},
		{3, 6},
		{2, 1},
		{1, 3},
		{5, 85},
	}
	result := Knapsack(10, inputs) // if maxWeight == 10 => result = 96
	fmt.Println(result)
}
