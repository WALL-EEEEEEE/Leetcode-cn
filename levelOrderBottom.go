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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		level_nodes := []int{}
		for i := 0; i < size; i++ {
			node := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			level_nodes = append(level_nodes, node.Val)
		}
		result = append(result, level_nodes)
	}
	for i := 0; i < len(result)/2; i++ {
		tmp := result[i]
		result[i] = result[len(result)-1-i]
		result[len(result)-1-i] = tmp
	}
	return result
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
		level_order := levelOrderBottom(test_case)
		fmt.Printf("树：%v, 层序遍历为：%v\n", test_case, level_order)
	}

}
