package util

import (
	"fmt"
	"strings"
)

type digit int

const (
	_            = iota
	digit8 digit = 8 << (iota - 1)
	digit16
	digit32
	digit64
)

const delimiterSize = 4

func ToBinary8(src uint) string {
	return toBinary(src, digit8)
}

func ToBinary16(src uint) string {
	return toBinary(src, digit16)
}

func ToBinary32(src uint) string {
	return toBinary(src, digit32)
}

func ToBinary64(src uint) string {
	return toBinary(src, digit64)
}

func toBinary(src uint, digit digit) string {
	str := fmt.Sprintf("%064b", src)

	size := len(str) / delimiterSize
	results := make([]string, size)
	for i := 0; i < size; i++ {
		results[i] = str[i*delimiterSize : i*delimiterSize+delimiterSize]
	}

	start := size - (int)(digit/delimiterSize)
	return strings.Join(results[start:size], " ")
}
