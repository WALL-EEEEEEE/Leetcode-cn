package structs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NewType int8

const (
	PREORDER NewType = iota
	POSTORDER
	INORDER
	LEVELORDER
	DEEPOrder
)

func makeTreeLevelOrder(arr []interface{}) *TreeNode {
	if len(arr) < 1 {
		return nil
	}
	root := &TreeNode{}
	for inode, vnode := range arr[0:] {
		if inode == 0 {
			root.Val = vnode.(int)
		} else if inode%2 == 0 && vnode != nil {
			root.Left = &TreeNode{
				Val: vnode.(int),
			}
		} else if inode%2 == 1 {
			if vnode != nil {
				root.Right = &TreeNode{
					Val: vnode.(int),
				}
			}
			root = root.Left
		}
	}
	return root
}

func NewTree(arr []interface{}, new_type NewType) *TreeNode {
	if new_type == LEVELORDER {
		return makeTreeLevelOrder(arr)
	}
	return nil
}
