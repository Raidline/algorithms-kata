package model

type GraphEdge struct {
	To     int
	Weight int
}

type WeightedAdjacencyList struct {
	Graph [][]GraphEdge
}

type WeightedAdjacencyMatrix struct {
	Graph [][]int
}
