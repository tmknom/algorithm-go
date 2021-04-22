package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Edge struct {
	Id        int
	Connected []*Edge
}

func NewEdge(id int) *Edge {
	return &Edge{Id: id, Connected: []*Edge{}}
}

func (e *Edge) Connect(other *Edge) {
	e.Connected = append(e.Connected, other)
}

func (e *Edge) Dump() string {
	str := make([]string, len(e.Connected))
	for i, edge := range e.Connected {
		str[i] = strconv.Itoa(edge.Id)
	}
	return fmt.Sprintf("&Edge{Id: %d, Connected: []{%s}}", e.Id, strings.Join(str, ", "))
}

type Graph struct {
	Edges map[int]*Edge
}

func NewGraph() *Graph {
	return &Graph{Edges: map[int]*Edge{}}
}

func (g *Graph) GetOrAddEdge(id int) *Edge {
	edge := g.GetEdge(id)
	if edge != nil {
		return edge
	}
	return g.AddEdge(id)
}

func (g *Graph) GetEdge(id int) *Edge {
	return g.Edges[id]
}

func (g *Graph) AddEdge(id int) *Edge {
	g.Edges[id] = NewEdge(id)
	return g.Edges[id]
}

func (g *Graph) Len() int {
	return len(g.Edges)
}

func (g *Graph) Dump() string {
	result := "&Graph{\n"
	for id, edge := range g.Edges {
		result += fmt.Sprintf("  {%d: %s},\n", id, edge.Dump())
	}
	result += "}\n"
	return result
}

func dfs(graph *Graph, edge *Edge, seen map[int]bool) {
	seen[edge.Id] = true
	fmt.Printf("dfs(%v)\n", edge.Id)

	for _, connected := range edge.Connected {
		if seen[connected.Id] {
			continue
		}
		dfs(graph, connected, seen)
	}
}

func initGraph(inputs [][]int) *Graph {
	graph := NewGraph()

	for _, input := range inputs {
		edge := graph.GetOrAddEdge(input[0])
		edge.Connect(graph.GetOrAddEdge(input[1]))
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
	graph := initGraph(inputs)
	fmt.Printf("inputs = %#v\n", inputs)
	fmt.Println(graph.Dump())

	seen := make(map[int]bool, graph.Len())
	for _, edge := range graph.Edges {
		seen[edge.Id] = false
	}

	dfs(graph, graph.GetEdge(4), seen)
}
