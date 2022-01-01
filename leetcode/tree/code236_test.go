package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 重点：题目描述隐含了p、q节点一定是在该树中存在的，
// 即公共p、q要么一个在左子树一个在右子树，要么两个全在左子树要么全在右子树
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q { return root }
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil { return root }
	if left == nil && right != nil { return right }
	if left != nil && right == nil { return left }
	return nil
}
