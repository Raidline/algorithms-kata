package algorithms

import (
	"algos/kata/algorithms/utils"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{9, 420, 7, 4, 69, 3, 42}
	expected := []int{3, 4, 7, 9, 42, 69, 420}

	QuickSort(arr)

	utils.AssertArrays(t, "QuickSort", expected, arr)
}
