package main

import (
	"container/list"
	"fmt"
)

type Point struct {
	x, y int
}

var maze = [][]int{
	{0, 1, 0, 0, 0, 0},
	{0, 1, 0, 1, 1, 0},
	{0, 0, 0, 1, 0, 0},
	{1, 1, 1, 1, 0, 1},
	{0, 0, 0, 0, 0, 0},
}

var directions = []Point{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func isSafe(maze [][]int, visited [][]bool, x, y int) bool {
	return x >= 0 && x < len(maze) && y >= 0 && y < len(maze[0]) && maze[x][y] == 0 && !visited[x][y]
}

func bfs(maze [][]int, start, end Point) int {
	rows := len(maze)
	cols := len(maze[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	queue := list.New()
	queue.PushBack(start)
	visited[start.x][start.y] = true

	steps := 0

	for queue.Len() > 0 {
		size := queue.Len()

		for i := 0; i < size; i++ {
			current := queue.Remove(queue.Front()).(Point)

			if current == end {
				return steps
			}
			for _, dir := range directions {
				nextX, nextY := current.x+dir.x, current.y+dir.y
				if isSafe(maze, visited, nextX, nextY) {
					queue.PushBack(Point{nextX, nextY})
					visited[nextX][nextY] = true
				}
			}
		}
		steps++
	}
	return -1
}

func main() {
	start := Point{0, 0}
	end := Point{4, 5}
	result := bfs(maze, start, end)
	if result != -1 {
		fmt.Printf("找到最短路徑，步數為: %d\n", result)
	} else {
		fmt.Println("無法到達終點")
	}
}
