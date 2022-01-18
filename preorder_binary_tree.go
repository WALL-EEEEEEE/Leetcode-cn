package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack *list.List = list.New()
	if root == nil {
		return ans
	}
	stack.PushBack(root)
	for stack.Len() > 0 {
		root = stack.Back().Value.(*TreeNode)
		stack.Remove(stack.Back())
		ans = append(ans, root.Val)
		if root.Right != nil {
			stack.PushBack(root.Right)
		}
		if root.Left != nil {
			stack.PushBack(root.Left)
		}
	}
	return ans
}

func main() {
	trees := []*TreeNode{
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		},
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		},
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	for _, tree := range trees {
		fmt.Printf("前序遍历为：%v\n", preorderTraversal(tree))
	}
}
