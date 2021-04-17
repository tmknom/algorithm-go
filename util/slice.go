package util

type SingleIntSlice []int
type DoubleIntSlice [][]int

const dpSize = 100010

var SingleIntDP = InitInfIntSlice(dpSize)
var DoubleIntDP = InitInfDoubleIntSlice(dpSize, 110)

func (s SingleIntSlice) SetFirst(val int) int {
	s[0] = val
	return s[0]
}

func (s SingleIntSlice) IsComputed(i int) bool {
	return s[i] != Inf
}

func InitInfIntSlice(size int) SingleIntSlice {
	return initIntSlice(size, Inf)
}

func initIntSlice(size int, val int) SingleIntSlice {
	slice := make([]int, size)
	for i, _ := range slice {
		slice[i] = val
	}
	return slice
}

func InitInfDoubleIntSlice(x int, y int) DoubleIntSlice {
	return initDoubleIntSlice(x, y, Inf)
}

func initDoubleIntSlice(x int, y int, val int) DoubleIntSlice {
	slice := make([][]int, x)
	for i, _ := range slice {
		slice[i] = make([]int, y)
		for j := 0; j < y; j++ {
			slice[i][j] = val
		}
	}
	return slice
}
