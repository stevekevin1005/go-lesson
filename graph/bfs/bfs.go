package main

import (
	"fmt"
)

type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]int)}
}

func (g *Graph) AddEdge(v, w int) {
	g.adjList[v] = append(g.adjList[v], w)
	g.adjList[w] = append(g.adjList[w], v) // 無向圖
}

func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	visited[start] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("Visited %d\n", node)

		for _, neighbor := range g.adjList[node] {
			// if !visited[neighbor] {
			visited[neighbor] = true
			queue = append(queue, neighbor)
			// }
		}
	}
}

func main() {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)

	fmt.Println("BFS starting from node 1:")
	g.BFS(1)
}
