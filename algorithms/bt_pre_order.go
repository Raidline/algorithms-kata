package algorithms

import "algos/kata/algorithms/model"

func walkPreOrder(curr *model.BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}

	//pre -> for the other types of traversal just move this to the the correct lines
	path = append(path, curr.Value)
	//recurse
	path = walkPreOrder(curr.Left, path)
	// path = append(path, curr.Value) -> inOrder
	path = walkPreOrder(curr.Right, path)
	//path = append(path, curr.Value) -> postOrder
	// post
	return path
}

func PreOrderSearch(head *model.BinaryNode) []int {
	return walkPreOrder(head, []int{})
}
