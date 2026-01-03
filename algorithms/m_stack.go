package algorithms

import "algos/kata/algorithms/model"

type MStack struct {
	Length int
	head   *model.Node
}

func NewStack() *MStack {
	return &MStack{
		Length: 0,
		head:   nil,
	}
}

func (s *MStack) Push(value int) {
	s.Length++
	node := createStackNode(value)
	if s.head == nil {
		s.head = node

		return
	}

	node.Prev = s.head
	s.head = node
}

func (s *MStack) Pop() int {
	if s.head == nil {
		return -1
	}
	s.Length--

	prevNode := s.head.Prev

	retValue := s.head.Value
	s.head = nil
	s.head = prevNode

	return retValue
}

func (s *MStack) Peek() int {
	if s.head == nil {
		return -1
	}

	return s.head.Value
}

func createStackNode(value int) *model.Node {
	return &model.Node{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
}
