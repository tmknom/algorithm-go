package main

import "fmt"

func LinearSearch(target int, inputs []int) {
	exists := false
	for _, input := range inputs {
		if input == target {
			exists = true
		}
	}

	fmt.Printf("%v in %v: ", target, inputs)
	if exists {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func main() {
	LinearSearch(6, []int{1, 5, 10, 6, 100, 23})
	LinearSearch(13, []int{1, 5, 10, 6, 100, 23})
}
