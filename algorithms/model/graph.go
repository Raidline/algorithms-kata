package model

type GraphEdge struct {
	To   int
	From int
}

type WeightedAdjacencyList struct {
	graph [][]GraphEdge
}

type WeightedAdjacencyMatrix struct {
	graph [][]int
}
