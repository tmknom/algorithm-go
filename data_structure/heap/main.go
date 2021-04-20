package main

import (
	"fmt"
)

type Heap struct {
	values []int
}

func NewDefaultHeap() *Heap {
	return NewHeap(1000)
}

func NewHeap(size int) *Heap {
	return &Heap{values: make([]int, 0, size)}
}

func (h *Heap) Top() int {
	if !h.isEmpty() {
		return h.values[0]
	}
	return -1
}

func (h *Heap) Push(val int) {
	fmt.Printf("Push(%d): ", val)

	// 最後尾に挿入
	h.values = append(h.values, val)
	// 挿入されたインデックス
	index := h.lastIndex()

	for {
		if index <= 0 {
			break
		}

		// 親インデックス
		parentIndex := (index - 1) / 2

		// 親要素のほうが大きければ終了
		if h.values[parentIndex] >= val {
			break
		}
		// 親要素のほうが小さければ、子要素のインデックスに親をセットする（親子を入れ替える）
		h.values[index] = h.values[parentIndex]
		index = parentIndex
	}
	// 挿入対象の要素の位置が確定したので格納する
	h.values[index] = val
}

func (h *Heap) Pop() int {
	fmt.Printf("Pop():   ")

	if h.isEmpty() {
		return -1
	}
	// 先頭の値を取得
	result := h.values[0]

	// 末尾の値を暫定的に親要素にする
	lastValue := h.values[h.lastIndex()]

	index := 0
	for {
		leftIndex := index*2 + 1
		rightIndex := index*2 + 2

		// 子要素がなければループを抜ける
		if leftIndex > h.length()-1 {
			break
		}

		// 子要素同士を比較して、大きい方を選択
		childIndex := leftIndex
		if rightIndex < h.length()-1 && h.values[rightIndex] > h.values[leftIndex] {
			childIndex = rightIndex
		}

		// 親要素より子要素のほうが小さければループを抜ける
		if h.values[childIndex] <= lastValue {
			break
		}

		// 子要素と親要素を入れ替える
		h.values[index] = h.values[childIndex]
		index = childIndex
	}

	// 最後尾の値を適切な位置にセット
	h.values[index] = lastValue

	// スライスから末尾の値を削除
	h.values = h.values[0 : h.length()-1]
	return result
}

func (h *Heap) Print() {
	fmt.Printf("heap = %#v\n", h.values)
}

func (h *Heap) lastIndex() int {
	return h.length() - 1
}

func (h *Heap) isEmpty() bool {
	return h.length() == 0
}

func (h *Heap) length() int {
	return len(h.values)
}

func main() {
	heap := NewDefaultHeap()

	heap.Push(19)
	heap.Print()
	heap.Push(12)
	heap.Print()
	heap.Push(15)
	heap.Print()
	heap.Push(10)
	heap.Print()
	heap.Push(7)
	heap.Print()
	heap.Push(6)
	heap.Print()
	heap.Push(1)
	heap.Print()
	heap.Push(3)
	heap.Print()
	heap.Push(7)
	heap.Print()
	heap.Push(5)
	heap.Print()
	heap.Push(3)
	heap.Print()
	heap.Push(2)
	heap.Print()

	fmt.Println("")
	heap.Pop()
	heap.Print()
	heap.Pop()
	heap.Print()
	heap.Pop()
	heap.Print()
}
