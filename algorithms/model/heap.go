package model

type Heap interface {
	Insert(value int)
	Delete() int
	Length() int
}
