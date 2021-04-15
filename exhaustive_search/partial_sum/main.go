package main

import "fmt"

const INF = 1 << 60

func PartialSum(w int, inputs []int) {
	length := len(inputs)
	exists := false

	for bit := 0; bit < (1 << length); bit++ {
		sum := 0
		for i := 0; i < length; i++ {
			if (bit & (1 << i)) != 0 {
				sum += inputs[i]
			}
		}

		if sum == w {
			exists = true
			break
		}
	}

	fmt.Printf("inputs: %v\n", inputs)
	fmt.Printf("w: %v\n", w)
	if exists {
		fmt.Println("yes\n")
	} else {
		fmt.Println("no\n")
	}
}

func main() {
	PartialSum(11, []int{5, 10, 6, 100, 23})
	PartialSum(17, []int{5, 10, 6, 100, 23})
	PartialSum(29, []int{5, 10, 6, 100, 23})
}
