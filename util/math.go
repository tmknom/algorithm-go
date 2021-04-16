package util

import "math"

func IntMax(x int, y int) int {
	return (int)(math.Max(float64(x), float64(y)))
}

func IntMin(x int, y int) int {
	return (int)(math.Min(float64(x), float64(y)))
}

func IntAbs(v int) int {
	return (int)(math.Abs(float64(v)))
}

func IntChooseMax(x *int, y int) {
	if *x < y {
		*x = y
	}
}

func IntChooseMin(x *int, y int) {
	if *x > y {
		*x = y
	}
}
