package algorithms

import "algos/kata/algorithms/model"

func CompareTrees(a *model.BinaryNode, b *model.BinaryNode) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	return CompareTrees(a.Left, b.Left) && CompareTrees(a.Right, b.Right)
}
