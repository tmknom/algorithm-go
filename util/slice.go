package util

func InitInfIntSlice(size int) []int {
	return initIntSlice(size, Inf)
}

func initIntSlice(size int, val int) []int {
	slice := make([]int, size)
	for i, _ := range slice {
		slice[i] = val
	}
	return slice
}
