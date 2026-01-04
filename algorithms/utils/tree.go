package utils

import "algos/kata/algorithms/model"

func GenerateTree() *model.BinaryNode {
	return &model.BinaryNode{
		Value: 20,
		Right: &model.BinaryNode{
			Value: 50,
			Right: &model.BinaryNode{
				Value: 100,
				Right: nil,
				Left:  nil,
			},
			Left: &model.BinaryNode{
				Value: 30,
				Right: &model.BinaryNode{
					Value: 45,
					Right: nil,
					Left:  nil,
				},
				Left: &model.BinaryNode{
					Value: 29,
					Right: nil,
					Left:  nil,
				},
			},
		},
		Left: &model.BinaryNode{
			Value: 10,
			Right: &model.BinaryNode{
				Value: 15,
				Right: nil,
				Left:  nil,
			},
			Left: &model.BinaryNode{
				Value: 5,
				Right: &model.BinaryNode{
					Value: 7,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
		},
	}
}

func GenerateTree2() *model.BinaryNode {
	return &model.BinaryNode{
		Value: 20,
		Right: &model.BinaryNode{
			Value: 50,
			Right: nil,
			Left: &model.BinaryNode{
				Value: 30,
				Right: &model.BinaryNode{
					Value: 45,
					Right: &model.BinaryNode{
						Value: 49,
						Left:  nil,
						Right: nil,
					},
					Left: nil,
				},
				Left: &model.BinaryNode{
					Value: 29,
					Right: nil,
					Left: &model.BinaryNode{
						Value: 21,
						Right: nil,
						Left:  nil,
					},
				},
			},
		},
		Left: &model.BinaryNode{
			Value: 10,
			Right: &model.BinaryNode{
				Value: 15,
				Right: nil,
				Left:  nil,
			},
			Left: &model.BinaryNode{
				Value: 5,
				Right: &model.BinaryNode{
					Value: 7,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
		},
	}
}
