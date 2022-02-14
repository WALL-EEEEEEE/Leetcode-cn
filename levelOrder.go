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
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return [][]int{}
	}
	queue := list.New()
	queue.PushBack(root)
	var level_result []int
	for queue.Len() > 0 {
		qsize := queue.Len()
		for i := 0; i < qsize; i++ {
			item := queue.Front()
			queue.Remove(item)
			node := item.Value.(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			level_result = append(level_result, node.Val)
		}
		result = append(result, level_result)
		level_result = []int{}
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
		level_order := levelOrder(test_case)
		fmt.Printf("树：%v, 层序遍历为：%v\n", test_case, level_order)
	}

}
