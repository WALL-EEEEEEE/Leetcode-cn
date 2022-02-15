package main

import (
	. "algorithm/structs"
	"container/list"
	"fmt"
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
	stack := list.New()
	stack.PushBack(root)
	depth := 0
	for stack.Len() > 0 {
		size := stack.Len()
		for i := 0; i < size; i++ {
			node := stack.Front().Value.(*TreeNode)
			stack.Remove(stack.Front())
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			if node.Left == nil && node.Right == nil {
				return depth + 1
			}
		}
		depth++
	}
	return depth
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
