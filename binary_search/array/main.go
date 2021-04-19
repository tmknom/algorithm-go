package main

import "fmt"

func BinarySearch(inputs []int, key int) {
	left := 0
	right := len(inputs) - 1

	result := -1
	for {
		if left > right {
			break
		}

		middle := (left + right) / 2
		fmt.Printf("inputs = %#v, inputs[middle] = %d\n", inputs[left:right+1], inputs[middle])
		if inputs[middle] == key {
			result = middle
			break
		} else if inputs[middle] > key {
			right = middle - 1
		} else if inputs[middle] < key {
			left = middle + 1
		}
	}

	fmt.Printf("result = %d, key =%d\n\n", result, key)
}

func main() {
	BinarySearch([]int{3, 5, 8, 10, 14, 17, 21, 39}, 14)
	BinarySearch([]int{3, 5, 8, 10, 14, 17, 21, 39}, 39)
	BinarySearch([]int{3, 5, 8, 10, 14, 17, 21, 39}, 23)
}
