package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

// 广度优先
func bfsWalk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Quene := []point{start}
	for len(Quene) > 0 {
		cur := Quene[0]
		Quene = Quene[1:]
		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			val := next.at(maze)
			if val == -1 || val == 1 {
				continue
			}
			val = next.at(steps)
			if val != 0 {
				continue
			}
			if next == start {
				continue
			}
			steps[next.i][next.j] = cur.at(steps) + 1
			Quene = append(Quene, next)
		}
	}
	return steps
}

// 深度优先
func dfsWalk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Stack := []point{start}
	for len(Stack) > 0 {
		cur := Stack[len(Stack)-1]
		Stack = Stack[0 : len(Stack)-1]
		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			val := next.at(maze)
			if val == -1 || val == 1 {
				continue
			}
			val = next.at(steps)
			if val != 0 {
				continue
			}
			if next == start {
				continue
			}
			steps[next.i][next.j] = cur.at(steps) + 1
			Stack = append(Stack, next)
		}
	}
	return steps
}
func (p point) add(r point) point {
	p.i += r.i
	p.j += r.j
	return p
}
func (p point) at(grid [][]int) int {
	if p.i < 0 || p.i >= len(grid) {
		return -1
	}
	if p.j < 0 || p.j >= len(grid[p.j]) {
		return -1
	}
	return grid[p.i][p.j]
}
func main() {
	/* 迷宫地图
	   6 5
	   0 1 0 0 0
	   0 0 0 1 0
	   0 1 0 1 0
	   0 1 1 0 0
	   0 0 0 0 1
	   0 1 0 0 0
	*/
	maze := readMaze("/Users/alvin/whx/learn/code/test/func/data.txt")
	for _, row := range maze {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
	fmt.Println()
	step := bfsWalk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range step {
		for _, col := range row {
			fmt.Printf("%3d ", col)
		}
		fmt.Println()
	}
	fmt.Println()
	step = dfsWalk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range step {
		for _, col := range row {
			fmt.Printf("%3d ", col)
		}
		fmt.Println()
	}
}
