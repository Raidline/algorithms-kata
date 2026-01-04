package algorithms

import "algos/kata/algorithms/model"

func walkPreOrder(curr *model.BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}

	//pre
	path = append(path, curr.Value)
	//recurse
	path = walkPreOrder(curr.Left, path)
	path = walkPreOrder(curr.Right, path)
	// post
	return path
}

func PreOrderSearch(head *model.BinaryNode) []int {
	return walkPreOrder(head, []int{})
}
