package util

import "math"

type SingleIntSlice []int
type DoubleIntSlice [][]int

var (
	SinglePositiveIntSlice = InitSingleIntSlice(columnSize, PositiveInf)
	SingleNegativeIntSlice = InitSingleIntSlice(columnSize, NegativeInf)
	DoublePositiveIntSlice = InitDoubleIntSlice(columnSize, rowSize, PositiveInf)
	DoubleNegativeIntSlice = InitDoubleIntSlice(columnSize, rowSize, NegativeInf)
)

const (
	PositiveInf = math.MaxInt64 >> 3 // +1,152,921,504,606,846,975
	NegativeInf = math.MinInt64 >> 3 // -1,152,921,504,606,846,976
)

const (
	columnSize = 10010
	rowSize    = 110
)

func (s SingleIntSlice) SetFirst(val int) int {
	s[0] = val
	return s[0]
}

func (s SingleIntSlice) Length() int {
	return len(s)
}

func (s SingleIntSlice) IsComputed(i int) bool {
	return (s[i] != PositiveInf) && (s[i] != NegativeInf)
}

func InitSingleIntSlice(size int, val int) SingleIntSlice {
	slice := make([]int, size)
	for i, _ := range slice {
		slice[i] = val
	}
	return slice
}

func (s DoubleIntSlice) ColumnLength() int {
	return len(s)
}

func (s DoubleIntSlice) RowLength() int {
	if s.ColumnLength() == 0 {
		return 0
	}
	return len(s[0])
}

func InitDoubleIntSlice(x int, y int, val int) DoubleIntSlice {
	slice := make([][]int, x)
	for i, _ := range slice {
		slice[i] = make([]int, y)
		for j := 0; j < y; j++ {
			slice[i][j] = val
		}
	}
	return slice
}
