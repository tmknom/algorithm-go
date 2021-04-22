package main

import (
	"../../util"
	"fmt"
)

func bfs(graph *util.Graph, startId int) map[int]int {
	dist := map[int]int{}
	queue := util.NewQueue()

	dist[startId] = 0
	queue.Enqueue(startId)

	for {
		if queue.IsEmpty() {
			break
		}

		id := queue.Dequeue()
		edge := graph.GetEdge(id)
		for _, connected := range edge.Connected {
			_, ok := dist[connected.Id]
			if ok {
				continue
			}

			dist[connected.Id] = dist[id] + 1
			queue.Enqueue(connected.Id)
		}
	}
	return dist
}

func main() {
	inputs := [][]int{
		{0, 1},
		{0, 4},
		{0, 2},
		{1, 0},
		{1, 3},
		{1, 4},
		{1, 8},
		{2, 0},
		{2, 5},
		{3, 1},
		{3, 8},
		{3, 7},
		{4, 0},
		{4, 1},
		{4, 8},
		{5, 2},
		{5, 8},
		{5, 6},
		{6, 5},
		{6, 7},
		{7, 3},
		{7, 6},
		{8, 1},
		{8, 4},
		{8, 3},
	}
	graph := util.InitGraph(inputs)
	fmt.Println(graph.Dump())

	startId := 0
	dist := bfs(graph, startId)
	fmt.Printf("dist = %#v\n", dist)
}
