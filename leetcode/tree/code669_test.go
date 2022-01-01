package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// found root in [low, high]
	for root != nil && (root.Val < low || root.Val > high){
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	// found left tree node in left of low
	cur := root
	for cur != nil {
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right
		}
		cur = cur.Left
	}
	// found right tree node in right of high
	cur = root
	for cur != nil {
		for cur.Right != nil && cur.Right.Val > high {
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}
	return root
}
