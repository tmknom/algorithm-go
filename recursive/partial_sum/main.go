package main

import (
	"../../util"
	"fmt"
)

const INF = 1 << 60

// 総和wがN個の配列inputsの部分和で計算できるか判定
// 1. inputs[N-1]を含まない：inputs[0]...inputs[N-2]の部分和「w」が存在する
// 2. inputs[N-1]を含む：inputs[0]...inputs[N-2]の部分和「w-inputs[N-1]」が存在する
func PartialSum(i int, w int, inputs *[]int, indent int) bool {
	util.CallPrintf(indent, "PartialSum(%v, %v)\n", i, w)

	// ベースケース
	// ゼロ個の配列の総和がゼロなら部分和が存在する
	if i == 0 {
		if w == 0 {
			util.ReturnPrintf(indent, "PartialSum(%v, %v) = true\n", i, w)
			return true
		} else {
			util.ReturnPrintf(indent, "PartialSum(%v, %v) = false\n", i, w)
			return false
		}
	}

	// 1. inputs[N-1]を含まない：inputs[0]...inputs[N-2]の部分和「w」が存在する
	if PartialSum(i-1, w, inputs, indent+1) {
		util.ReturnPrintf(indent, "PartialSum(%v, %v) = true\n", i-1, w)
		return true
	}

	// 2. inputs[N-1]を含む：inputs[0]...inputs[N-2]の部分和「w-inputs[N-1]」が存在する
	if PartialSum(i-1, w-(*inputs)[i-1], inputs, indent+1) {
		util.ReturnPrintf(indent, "PartialSum(%v, %v) = true\n", i-1, w-(*inputs)[i-1])
		return true
	}

	// 1・2番の条件に合致しない
	util.ReturnPrintf(indent, "PartialSum(%v, %v) = false\n", i, w)
	return false
}

func run(w int, inputs *[]int) {
	result := PartialSum(len(*inputs), w, inputs, 0)
	fmt.Printf("PartialSum(%v, %#v) = %v\n", w, *inputs, result)
}

func main() {
	run(4, &[]int{1, 2, 3})

	fmt.Println("")
	run(11, &[]int{5, 10, 6, 100, 23})

	fmt.Println("")
	run(17, &[]int{5, 10, 6, 100, 23})
}
