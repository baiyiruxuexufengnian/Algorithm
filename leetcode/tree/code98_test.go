package tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil{
		return true
	}
	return isBST(root,math.MinInt64,math.MaxInt64)
}

// 前序中序都可以
func isBST(root *TreeNode,min, max int) bool{
	if root == nil{
		return true
	}
	// 中序遍历
	l := isBST(root.Left,min,root.Val)  // 最左下节点要比最小min大，要比最大root value小
	if min >= root.Val || max <= root.Val{
		return false                    // 中
	}
	r := isBST(root.Right,root.Val,max) // 最右下节点要比最小min大，要比最大root value小
	return l && r
}