package model

type Queue interface {
	Enqueue(value int)
	Deque() int
	Peek() int
}
