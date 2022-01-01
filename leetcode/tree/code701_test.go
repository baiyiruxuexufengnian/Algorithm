package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:val,
		}
	}
	cur := root
	parent := root
	for cur != nil {
		parent = cur
		if cur.Val >= val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	// 一定是按照正确的路径走到了对应的空节点之上了
	node := &TreeNode{
		Val: val,
	}
	if parent.Val >= val {
		parent.Left = node
		return root
	}
	parent.Right = node
	return root
}
