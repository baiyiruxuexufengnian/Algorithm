package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	var parent *TreeNode
	cur := root
	for cur != nil && cur.Val != key {
		parent = cur
		if key > cur.Val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	if cur == nil {
		return root
	}
	node := cur.Right
	if node != nil {
		for node.Left != nil {
			node = node.Left
		}
		node.Left = cur.Left
	} else {
		cur.Right = cur.Left
	}
	if parent == nil {
		return cur.Right
	}
	if parent.Val < key {
		parent.Right = cur.Right
	} else {
		parent.Left = cur.Right
	}
	return root
}
