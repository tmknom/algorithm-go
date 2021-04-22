package main

import (
	"../../util"
	"fmt"
)

func dfs(graph *util.Graph, startId int) map[int]int {
	dist := map[int]int{}
	stack := util.NewStack()

	dist[startId] = 0
	stack.Push(startId)

	for {
		if stack.IsEmpty() {
			break
		}

		id := stack.Pop()
		edge := graph.GetEdge(id)
		for _, connected := range edge.Connected {
			_, ok := dist[connected.Id]
			if ok {
				continue
			}

			dist[connected.Id] = dist[id] + 1
			stack.Push(connected.Id)
		}
	}
	return dist
}

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

func stSearchByDFS(graph *util.Graph, startId int, targetId int) {
	output(dfs(graph, startId), startId, targetId)
}

func stSearchByBFS(graph *util.Graph, startId int, targetId int) {
	output(bfs(graph, startId), startId, targetId)
}

func output(dist map[int]int, startId int, targetId int) {
	val, ok := dist[targetId]
	if ok {
		fmt.Printf("exist %d to %d, dist = %d\n", startId, targetId, val)
	} else {
		fmt.Printf("not exist %d to %d\n", startId, targetId)
	}
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
	graph := util.InitGraph(inputs)
	fmt.Println(graph.Dump())

	stSearchByBFS(graph, 4, 0)
	stSearchByBFS(graph, 0, 4)

	stSearchByDFS(graph, 4, 0)
	stSearchByDFS(graph, 0, 4)
}
