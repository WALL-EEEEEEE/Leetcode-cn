package main

import (
	. "algorithm/structs"
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left_depth, right_depth int
	if root.Left != nil {
		left_depth = maxDepth(root.Left) + 1
	}
	if root.Right != nil {
		right_depth = maxDepth(root.Right) + 1
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return int(math.Max(float64(left_depth), float64(right_depth)))
}

func main() {
	test_cases := []*TreeNode{
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		/*
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 5,
							},
						},
					},
				},
			},
			&TreeNode{
				Val: 1,
			},
			&TreeNode{},
		*/
	}
	for _, test_case := range test_cases {
		depth := maxDepth(test_case)
		fmt.Printf("树：%v, 层序遍历为：%v\n", test_case, depth)
	}

}
