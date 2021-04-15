package main

import "fmt"

const INF = 1 << 60

func MinLinearSearch(inputs []int) {
	result := INF
	for _, input := range inputs {
		if result > input {
			result = input
		}
	}

	fmt.Printf("min of %v: %v\n", inputs, result)
}

func main() {
	MinLinearSearch([]int{5, 10, 6, 100, 23})
	MinLinearSearch([]int{100, 15, 12, 6, 8, 15})
}
