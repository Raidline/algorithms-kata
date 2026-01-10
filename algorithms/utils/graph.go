package utils

import "algos/kata/algorithms/model"

//	  (1) --- (4) ---- (5)
//	/  |       |       /|
//
// (0)   | ------|------- |
//
//	\  |/      |        |
//	  (2) --- (3) ---- (6)
func List1() model.WeightedAdjacencyList {
	list1 := model.WeightedAdjacencyList{}
	list1.Graph = append(list1.Graph, []model.GraphEdge{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	})
	list1.Graph = append(list1.Graph, []model.GraphEdge{
		{To: 0, Weight: 3},
		{To: 2, Weight: 4},
		{To: 4, Weight: 1},
	})

	list1.Graph = append(list1.Graph, []model.GraphEdge{
		{To: 1, Weight: 4},
		{To: 3, Weight: 7},
		{To: 0, Weight: 1},
	})

	list1.Graph = append(list1.Graph, []model.GraphEdge{
		{To: 2, Weight: 7},
		{To: 4, Weight: 5},
		{To: 6, Weight: 1},
	})

	list1.Graph = append(list1.Graph, []model.GraphEdge{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	})

	list1.Graph = append(list1.Graph, []model.GraphEdge{
		{To: 6, Weight: 1},
		{To: 4, Weight: 2},
		{To: 2, Weight: 18},
	})

	list1.Graph = append(list1.Graph, []model.GraphEdge{
		{To: 3, Weight: 1},
		{To: 5, Weight: 1},
	})

	return list1
}

//	 >(1)<--->(4) ---->(5)
//	/          |       /|
//
// (0)     ------|------- |
//
//	\   v      v        v
//	 >(2) --> (3) <----(6)
func List2() model.WeightedAdjacencyList {
	list2 := model.WeightedAdjacencyList{}
	list2.Graph = append(list2.Graph, []model.GraphEdge{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	})
	list2.Graph = append(list2.Graph, []model.GraphEdge{
		{To: 4, Weight: 1},
	})

	list2.Graph = append(list2.Graph, []model.GraphEdge{
		{To: 3, Weight: 7},
	})

	list2.Graph = append(list2.Graph, []model.GraphEdge{})

	list2.Graph = append(list2.Graph, []model.GraphEdge{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	})

	list2.Graph = append(list2.Graph, []model.GraphEdge{
		{To: 2, Weight: 18},
		{To: 6, Weight: 1},
	})

	list2.Graph = append(list2.Graph, []model.GraphEdge{
		{To: 3, Weight: 1},
	})

	return list2
}

//	 >(1)<--->(4) ---->(5)
//	/          |       /|
//
// (0)     ------|------- |
//
//	\   v      v        v
//	 >(2) --> (3) <----(6)
func Matrix2() model.WeightedAdjacencyMatrix {
	matrix := model.WeightedAdjacencyMatrix{}

	//the outer index is the node we are seeing
	// each index is other node in the graph
	// the value in the index is the Weight related To that node from the outer index (curr node)
	matrix.Graph = append(matrix.Graph, []int{
		0, 3, 1, 0, 0, 0, 0, //0
	})

	matrix.Graph = append(matrix.Graph, []int{
		0, 0, 0, 0, 1, 0, 0,
	})

	matrix.Graph = append(matrix.Graph, []int{
		0, 0, 7, 0, 0, 0, 0,
	})

	matrix.Graph = append(matrix.Graph, []int{
		0, 0, 0, 0, 0, 0, 0,
	})

	matrix.Graph = append(matrix.Graph, []int{
		0, 1, 0, 5, 0, 2, 0,
	})

	matrix.Graph = append(matrix.Graph, []int{
		0, 0, 18, 0, 0, 0, 1,
	})

	matrix.Graph = append(matrix.Graph, []int{
		0, 0, 0, 1, 0, 0, 1,
	})

	return matrix
}
