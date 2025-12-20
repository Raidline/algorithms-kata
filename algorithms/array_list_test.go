package algorithms

import (
	"algos/kata/algorithms/utils"
	"testing"
)

func TestArrayList(t *testing.T) {
	list := New(3)
	utils.TestList(list, t)
}
