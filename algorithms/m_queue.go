package algorithms

import (
	"algos/kata/algorithms/model"
)

type MQueue struct {
	Length int
	head   *model.Node
	tail   *model.Node
}

func NewQueue() *MQueue {
	return &MQueue{
		Length: 0,
	}
}

func (q *MQueue) Enqueue(value int) {
	q.Length++

	n := createNode(value, nil)
	if q.tail == nil || q.head == nil {
		q.tail = n
		q.head = n

		return
	}

	q.tail.Next = n
	q.tail = n
}

func (q *MQueue) Deque() int {
	if q.head == nil {
		return -1 //in real applications should return nil
	}
	q.Length--

	head := q.head
	q.head = q.head.Next

	head.Next = nil

	return head.Value
}

func (q *MQueue) Peek() int {
	return q.head.Value
}

func createTail(value int) *model.Node {
	return &model.Node{
		Value: value,
		Next:  nil,
	}
}

func createNode(value int, next *model.Node) *model.Node {
	return &model.Node{
		Value: value,
		Next:  next,
	}
}
