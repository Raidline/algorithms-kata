package model

type Stack interface {
	Push(value int)
	Pop() int
	Peek() int
}
