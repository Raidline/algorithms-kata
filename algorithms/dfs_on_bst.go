package algorithms

import "algos/kata/algorithms/model"

func search(curr *model.BinaryNode, needle int) bool {

	if curr == nil {
		return false
	}

	if curr.Value == needle {
		return true
	}

	if curr.Value < needle {
		return search(curr.Right, needle)
	} else {
		return search(curr.Left, needle)
	}

}

func DepthFirstSearch(head *model.BinaryNode, needle int) bool {
	return search(head, needle)
}
