package algorithms

import "algos/kata/algorithms/model"

var dir = []model.Point{
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
	{X: 0, Y: 1},
}

func walk(maze []string, wall byte, curr *model.Point, end *model.Point, seen [][]bool, path []*model.Point) bool {

	if (curr.X < 0 || curr.X >= len(maze[0])) || (curr.Y < 0 || curr.Y >= len(maze)) {
		return false
	}

	if maze[curr.Y][curr.X] == wall {
		return false
	}

	if curr.X == end.X && curr.Y == end.Y {
		path = append(path, end)
		return true
	}

	if seen[curr.Y][curr.X] {
		return false
	}

	// 3 steps recurse
	// pre
	seen[curr.Y][curr.X] = true
	path = append(path, curr)
	// recurse
	for _, v := range dir {
		if walk(maze, wall, createPoint(curr, v), end, seen, path) {
			return true
		}
	}
	// post
	path = path[:len(path)-1]

	return false
}

func MazeSolver(maze []string, wall byte, start *model.Point, end *model.Point) []*model.Point {
	path := make([]*model.Point, 0)
	seen := make([][]bool, len(maze))

	for i := range maze {
		inner := make([]bool, len(maze[0]))

		for j := range inner {
			inner[j] = false
		}

		seen[i] = inner
	}

	walk(maze, wall, start, end, seen, path)

	return path
}

func createPoint(curr *model.Point, dir model.Point) *model.Point {
	return &model.Point{
		X: curr.X + dir.X,
		Y: curr.Y + dir.Y,
	}
}
