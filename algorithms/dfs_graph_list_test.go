package algorithms

import (
	"algos/kata/algorithms/utils"
	"testing"
)

func TestDfsGraphList(t *testing.T) {
	path, err := DfsGraphList(utils.List2(), 0, 6)

	if err != nil {
		t.Errorf("Testing DFS graph, got err as was not supposed to : %s", err.Error())
		t.FailNow()
	}

	utils.AssertArrays(t, "DFS Graph", []int{0,
		1,
		4,
		5,
		6}, path)

	_, err2 := DfsGraphList(utils.List2(), 6, 0)

	if err2 == nil {
		t.Error("Testing DFS graph, did not got err as was supposed to")
		t.FailNow()
	}
}
