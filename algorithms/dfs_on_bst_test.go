package algorithms

import (
	"algos/kata/algorithms/utils"
	"testing"
)

func TestDfsOnBst(t *testing.T) {

	if !DepthFirstSearch(utils.GenerateTree(), 45) {
		t.Errorf("Testing Finding 45 on BST by DFS, got false want true")
		t.FailNow()
	}

	if DepthFirstSearch(utils.GenerateTree(), 7) {
		t.Errorf("Testing Finding 7 on BST by DFS, got true want false")
		t.FailNow()
	}

	if !DepthFirstSearch(utils.GenerateTree(), 69) {
		t.Errorf("Testing Finding 69 on BST by DFS, got false want true")
		t.FailNow()
	}
}
