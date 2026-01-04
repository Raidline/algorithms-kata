package model

type BinaryNode struct { //should be generic, but do not want to deal with generics in g
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}

type GeneralNode struct {
	Value    int
	Children []*GeneralNode
}
