package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Graph represents a weighted graph
type Graph struct {
	nodes map[int][]Edge
}

// Edge represents a connection between nodes
type Edge struct {
	to     int
	weight int
}

// PriorityQueue is a min-heap for the priority queue
type PriorityQueue []*Item

// Item is an element of the priority queue
type Item struct {
	node     int
	priority int
	index    int // The index of the item in the heap
}

// Len returns the number of items in the priority queue
func (pq PriorityQueue) Len() int { return len(pq) }

// Less compares two items based on their priority
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// Swap swaps two items in the priority queue
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds an item to the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes the item with the highest priority (smallest distance)
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Dijkstra finds the shortest paths from a start node to all other nodes
func (g *Graph) Dijkstra(start int) map[int]int {
	distances := make(map[int]int)
	for node := range g.nodes {
		distances[node] = math.MaxInt
	}
	distances[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)

		for _, edge := range g.nodes[current.node] {
			newDist := distances[current.node] + edge.weight
			if newDist < distances[edge.to] {
				distances[edge.to] = newDist
				heap.Push(pq, &Item{node: edge.to, priority: newDist})
			}
		}
	}

	return distances
}

func main() {
	g := Graph{
		nodes: map[int][]Edge{
			1: {{3, 5}, {2, 2}, {6, 10}},
			2: {{1, 2}, {5, 8}, {4, 2}},
			3: {{1, 5}, {6, 5}},
			4: {{2, 2}, {6, 1}},
			5: {{2, 8}, {6, 6}},
			6: {{1, 10}, {3, 5}, {4, 1}, {5, 6}},
		},
	}

	distances := g.Dijkstra(5)
	fmt.Println("Shortest distances from node 0:", distances)
}
