package model

// should be generic, but i do not want to deal with go generics
// single linked list
type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}
