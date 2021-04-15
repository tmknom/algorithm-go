package main

import "fmt"

const INF = 1 << 60

func MinPair(threshold int, aSlice []int, bSlice []int) {
	if len(aSlice) != len(bSlice) {
		panic("different size")
	}

	resultSum := INF
	resultA := INF
	resultB := INF
	for _, a := range aSlice {
		for _, b := range bSlice {
			sum := a + b
			if sum < threshold {
				continue
			}

			if resultSum > sum {
				resultSum = sum
				resultA = a
				resultB = b
			}
		}
	}

	fmt.Printf("aSlice: %v\n", aSlice)
	fmt.Printf("bSlice: %v\n", bSlice)
	fmt.Printf("threshold: %v\n", threshold)
	fmt.Printf("resultSum of min: %v (%v, %v)\n\n", resultSum, resultA, resultB)
}

func main() {
	MinPair(10, []int{5, 10, 6, 100, 23}, []int{15, 12, 6, 8, 15})
	MinPair(20, []int{5, 10, 6, 100, 23}, []int{15, 12, 6, 8, 15})
	MinPair(30, []int{5, 10, 6, 100, 23}, []int{15, 12, 6, 8, 15})
}
