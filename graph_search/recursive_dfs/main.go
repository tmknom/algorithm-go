package main

import (
	"../../util"
	"fmt"
)

func dfs(graph util.DoubleIntSlice, val int, seen []bool) {
	seen[val] = true
	fmt.Printf("dfs(%v)\n", val)

	for _, value := range graph {
		if value[0] != val {
			continue
		}

		nextEdge := value[1]
		if seen[nextEdge] {
			continue
		}

		dfs(graph, nextEdge, seen)
	}
}

func initGraph(inputs [][]int) util.DoubleIntSlice {
	graph := util.InitDoubleIntSlice(len(inputs), 2, util.NegativeInf)
	for i := 0; i < len(inputs); i++ {
		graph[i] = inputs[i]
	}
	return graph
}

func main() {
	inputs := [][]int{
		{4, 1},
		{4, 2},
		{4, 6},
		{1, 3},
		{1, 6},
		{2, 5},
		{6, 7},
		{3, 0},
		{3, 7},
		{7, 0},
		{0, 5},
	}
	graphSize := len(inputs)
	graph := initGraph(inputs)
	fmt.Printf("graph = %#v\n\n", graph)

	seen := make([]bool, graphSize)
	for i := 0; i < len(seen); i++ {
		seen[i] = false
	}

	for v := 0; v < graphSize; v++ {
		if seen[v] {
			continue
		}
		dfs(graph, v, seen)
	}
}
