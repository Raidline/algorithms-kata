package algorithms

import (
	"algos/kata/algorithms/utils"
	"testing"
)

func TestCompareBinaryTrees(t *testing.T) {

	if !CompareTrees(utils.GenerateTree(), utils.GenerateTree()) {
		t.Errorf("Testing Comparing Binary trees, got false want true")
		t.FailNow()
	}

	if CompareTrees(utils.GenerateTree(), utils.GenerateTree2()) {
		t.Errorf("Testing Comparing Binary trees, got true want false")
		t.FailNow()
	}
}
