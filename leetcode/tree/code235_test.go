package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// 核心：p,q 两个节点一定是存在于该树之上的
func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var node *TreeNode
	if root.Val > p.Val && root.Val > q.Val {
		node = lowestCommonAncestor235(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		node = lowestCommonAncestor235(root.Right, p, q)
	} else {
		node = root
	}
	return node
}
