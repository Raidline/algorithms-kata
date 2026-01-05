package algorithms

import (
	"algos/kata/algorithms/utils"
	"testing"
)

func TestBtBts(t *testing.T) {
	tree := utils.GenerateTree()
	if !Bfs(tree, 45) {
		t.Errorf("Testing Breath First Search, did not find 45")
		t.FailNow()
	}
	if !Bfs(tree, 7) {
		t.Errorf("Testing Breath First Search, did not find 7")
		t.FailNow()
	}
	if Bfs(tree, 69) {
		t.Errorf("Testing Breath First Search, found 69 and was not supposed to")
		t.FailNow()
	}
}
