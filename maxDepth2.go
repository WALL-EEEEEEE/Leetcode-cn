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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	var depth int
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
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
		depth := maxDepth(test_case)
		fmt.Printf("树：%v, 最大深度为：%v\n", test_case, depth)
	}

}
