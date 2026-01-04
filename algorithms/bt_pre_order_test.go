package algorithms

import (
	"algos/kata/algorithms/utils"
	"testing"
)

func TestBTPreOrder(t *testing.T) {
	want := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}

	utils.AssertArrays(t, "PreOrderSearch", want, PreOrderSearch(utils.GenerateTree()))
}
