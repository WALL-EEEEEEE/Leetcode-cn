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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right != nil {
		return int(math.Min(float64(minDepth(root.Left)+1), float64(minDepth(root.Right)+1)))
	} else if root.Left != nil {
		return minDepth(root.Left) + 1
	} else if root.Right != nil {
		return minDepth(root.Right) + 1
	} else {
		return 1
	}
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
	}
	for _, test_case := range test_cases {
		depth := minDepth(test_case)
		fmt.Printf("树：%v, 最小深度为：%v\n", test_case, depth)
	}
}
