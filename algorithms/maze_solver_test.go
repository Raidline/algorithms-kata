package algorithms

import (
	"algos/kata/algorithms/model"
	"runtime/debug"
	"strings"
	"testing"
)

func TestMazeSolver(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	mazeResult := []*model.Point{
		makePoint(10, 0),
		makePoint(10, 1),
		makePoint(10, 2),
		makePoint(10, 3),
		makePoint(10, 4),
		makePoint(9, 4),
		makePoint(8, 4),
		makePoint(7, 4),
		makePoint(6, 4),
		makePoint(5, 4),
		makePoint(4, 4),
		makePoint(3, 4),
		makePoint(2, 4),
		makePoint(1, 4),
		makePoint(1, 5),
	}

	result := MazeSolver(maze, 'x', makePoint(10, 0), makePoint(1, 5))

	if !compareMazes(drawPath(maze, mazeResult), drawPath(maze, result)) {
		t.Errorf("Mazes did not match")
		debug.PrintStack()
		t.FailNow()
	}
}

func compareMazes(expected []string, actual []string) bool {
	for i := range expected {
		if expected[i] != actual[i] {
			return false
		}
	}

	return true
}

func drawPath(data []string, path []*model.Point) []string {
	data2 := make([][]string, len(data))

	for i := range data {
		data2[i] = strings.Split(data[i], "")
	}

	for _, p := range path {
		if data2[p.Y] != nil && inBounds(data2, p) {
			data2[p.Y][p.X] = "*"
		}
	}

	finalData := make([]string, len(data2))

	for i, v := range data2 {
		finalData[i] = strings.Join(v, "")
	}

	return finalData
}

func inBounds(data [][]string, p *model.Point) bool {
	return (p.Y > 0 || p.Y < len(data)) || (p.X > 0 || p.X < len(data[p.Y]))
}

func makePoint(x int, y int) *model.Point {
	return &model.Point{
		X: x,
		Y: y,
	}
}
