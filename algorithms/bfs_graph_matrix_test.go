package algorithms

import (
	"algos/kata/algorithms/utils"
	"testing"
)

func TestBfsGraphMatrix(t *testing.T) {
	path, err := BfsGraph(utils.Matrix2(), 0, 6)

	if err != nil {
		t.Errorf("Testing BFS graph, got err as was not supposed to : %s", err.Error())
		t.FailNow()
	}

	utils.AssertArrays(t, "BFS Graph", []int{0,
		1,
		4,
		5,
		6}, path)

	_, err2 := BfsGraph(utils.Matrix2(), 6, 0)

	if err2 == nil {
		t.Error("Testing BFS graph, did not got err as was supposed to")
		t.FailNow()
	}
}
