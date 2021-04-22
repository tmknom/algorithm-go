package main

import (
	"../../util"
	"fmt"
)

func dfs(graph *util.Graph, id int, colors map[int]color, color color) bool {
	colors[id] = color
	edge := graph.GetEdge(id)

	for _, connected := range edge.Connected {
		if colors[connected.Id] != unknown {
			if colors[connected.Id] == color {
				return false
			}
			continue
		}
		if !dfs(graph, connected.Id, colors, color.reverse()) {
			return false
		}
	}
	return true
}

func run(graph *util.Graph) {
	colors := make(map[int]color, graph.Len())
	for _, edge := range graph.Edges {
		colors[edge.Id] = unknown
	}

	result := true
	for _, edge := range graph.Edges {
		if colors[edge.Id] != unknown {
			continue
		}
		if !dfs(graph, 0, colors, left) {
			result = false
		}
	}

	if result {
		fmt.Printf("result: yes\n")
	} else {
		fmt.Printf("result: no\n")
	}
}

type color int

const (
	_ color = iota
	unknown
	left
	right
)

func (c color) reverse() color {
	if c == left {
		return right
	} else if c == right {
		return left
	}
	return unknown
}

func main() {
	inputs := [][]int{
		{1, 0},
		{1, 2},
		{1, 4},
		{0, 1},
		{0, 3},
		{2, 1},
		{3, 0},
		{3, 4},
		{4, 1},
		{4, 3},
	}
	graph := util.InitGraph(inputs)
	fmt.Println(graph.Dump())
	run(graph)
}
