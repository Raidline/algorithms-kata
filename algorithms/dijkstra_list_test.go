package algorithms

import (
	"algos/kata/algorithms/utils"
	"testing"
)

func TestDijkstraList(t *testing.T) {
	utils.AssertArrays(t, "Dijkstra", []int{0, 1, 4, 5, 6}, DijkstraList(0, 6, utils.List1()))
}
