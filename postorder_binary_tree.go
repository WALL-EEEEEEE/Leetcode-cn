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

func postorderTraversal(root *TreeNode) []int {
	var stack *list.List = list.New()
	var ans []int
	if root == nil {
		return ans
	}
	stack.PushBack(root)
	for stack.Len() > 0 {
		root = stack.Back().Value.(*TreeNode)
		if root.Left != nil {
			stack.PushBack(root.Left)
			root.Left = nil
			continue
		}
		if root.Right != nil {
			stack.PushBack(root.Right)
			root.Right = nil
			continue
		}
		if root.Left == nil && root.Right == nil {
			ans = append(ans, root.Val)
			stack.Remove(stack.Back())
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
		fmt.Printf("后序遍历为：%v\n", postorderTraversal(tree))
	}
}
