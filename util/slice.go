package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type SingleIntSlice []int
type DoubleIntSlice [][]int

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

func (s SingleIntSlice) Dump(x int) string {
	str := make([]string, x)
	for i := 0; i < x; i++ {
		str[i] = strconv.Itoa(s[i])
	}
	return fmt.Sprintf("[ %s ]\n", strings.Join(str, ", "))
}

func SinglePositiveIntSlice() SingleIntSlice {
	return InitSingleIntSlice(columnSize, PositiveInf)
}

func SingleNegativeIntSlice() SingleIntSlice {
	return InitSingleIntSlice(columnSize, NegativeInf)
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

func (s DoubleIntSlice) Dump(x int, y int) string {
	result := ""
	for i := 0; i < x; i++ {
		str := make([]string, y)
		for j := 0; j < y; j++ {
			str[j] = strconv.Itoa(s[i][j])
		}
		result += fmt.Sprintf("[ %s ]\n", strings.Join(str, ", "))
	}
	return result
}

func DoublePositiveIntSlice() DoubleIntSlice {
	return InitDoubleIntSlice(columnSize, rowSize, PositiveInf)
}

func DoubleNegativeIntSlice() DoubleIntSlice {
	return InitDoubleIntSlice(columnSize, rowSize, NegativeInf)
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
