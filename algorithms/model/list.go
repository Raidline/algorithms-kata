package model

type List interface {
	InsertAt(idx int, value int)
	Append(value int)
	Get(idx int) int
	RemoveAt(idx int) int
	Remove(value int) int
	Length() int
	Prepend(value int)
}
