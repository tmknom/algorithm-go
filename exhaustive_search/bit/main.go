package main

import (
	"fmt"
	"math"
	"strings"
)

type Digit int

const (
	_          = iota
	Int8 Digit = 8 << (iota - 1)
	Int16
	Int32
	Int64
)

func ToBinary(src int, digit Digit) string {
	const DelimiterSize = 4
	str := fmt.Sprintf("%064b", src)

	size := len(str) / DelimiterSize
	results := make([]string, size)
	for i := 0; i < size; i++ {
		results[i] = str[i*DelimiterSize : i*DelimiterSize+DelimiterSize]
	}

	start := size - (int)(digit/DelimiterSize)
	return strings.Join(results[start:size], " ")
}

func main() {
	fmt.Println(Int8)
	fmt.Println(Int16)
	fmt.Println(Int32)
	fmt.Println(Int64)

	fmt.Printf("8bits: %v\n", ToBinary(255, Int8))
	fmt.Printf("16bits: %v\n", ToBinary(256, Int16))
	fmt.Printf("32bits: %v\n", ToBinary(1<<31, Int32))
	fmt.Printf("64bits: %v\n", ToBinary(math.MaxInt64, Int64))
}
