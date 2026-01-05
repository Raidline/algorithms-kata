package algorithms

import "algos/kata/algorithms/model"

func Bfs(head *model.BinaryNode, needle int) bool {
	//queue := NewQueue() // should use the queue i created but it does not support generics, sadge
	//using array for that
	queue := make([]*model.BinaryNode, 0)

	queue = append(queue, head)

	for len(queue) != 0 {

		next := queue[len(queue)-1]
		queue = remove(queue, len(queue)-1)

		if next.Value == needle {
			return true
		}

		if next.Left != nil {
			queue = append(queue, next.Left)
		}

		if next.Right != nil {
			queue = append(queue, next.Right)
		}
	}

	return false
}

// not the more perfomant and prettiest but gets the job done
func remove(slice []*model.BinaryNode, s int) []*model.BinaryNode {
	return append(slice[:s], slice[s+1:]...)
}
