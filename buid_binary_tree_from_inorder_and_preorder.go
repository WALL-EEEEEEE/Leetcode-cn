package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeTuple struct {
	preorder interface{}
	inorder  interface{}
}

var inorder_node_map map[int]int

func makeTree(preorder []int, inorder []int, preorder_left int, preorder_right int, inorder_left int, inorder_right int) *TreeNode {
	//通过前序遍历获取根节点
	preorder_len := preorder_right - preorder_left
	inorder_len := inorder_right - inorder_left
	if preorder_left >= preorder_right {
		return nil
	}
	root_val := preorder[preorder_left]
	node := &TreeNode{
		Val: root_val,
	}
	root_index := inorder_node_map[root_val]
	size_left_subtree := root_index - inorder_left
	node.Left = makeTree(preorder, inorder, preorder_left+1, preorder_left+size_left_subtree, inorder_left, root_index)
	node.Right = makeTree(preorder, inorder, preorder_left+size_left_subtree, preorder_len, root_index, inorder_len)
	return node
}
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(inorder_node_map) < 1 {
		for i, v := range inorder {
			inorder_node_map[v] = i
		}
	}
	return makeTree(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}
func printTree(tree *TreeNode) string {
	str := ""
	if tree == nil {
		return ""
	}
	str += fmt.Sprintf("%v, ", tree.Val)
	str += printTree(tree.Left)
	str += printTree(tree.Right)
	return str
}

func main() {
	trees := []NodeTuple{
		NodeTuple{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
		},
		NodeTuple{
			[]int{-1},
			[]int{-1},
		},
	}
	for _, tree := range trees {
		inorder_node_map = make(map[int]int)
		preorder, inorder := tree.preorder.([]int), tree.inorder.([]int)
		fmt.Printf("二叉树的前序遍历为：%v, 后序遍历为: %v, ", preorder, inorder)
		binTree := buildTree(preorder, inorder)
		fmt.Printf("构建的二叉树为：%v\n", printTree(binTree))
	}

}
