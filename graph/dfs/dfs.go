package main

import "fmt"

type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]int)}
}

func (g *Graph) AddEdge(v, w int) {
	g.adjList[v] = append(g.adjList[v], w)
	g.adjList[w] = append(g.adjList[w], v)
}

func (g *Graph) DFS(start int, visited map[int]bool) {
	visited[start] = true
	fmt.Printf("Visited %d\n", start)

	for _, neighbor := range g.adjList[start] {
		if !visited[neighbor] {
			g.DFS(neighbor, visited)
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

	visited := make(map[int]bool)
	fmt.Println("DFS starting from node 1:")
	g.DFS(1, visited)
}
