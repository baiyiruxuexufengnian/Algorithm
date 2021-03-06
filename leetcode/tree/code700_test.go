package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
	// return searchBSTIteration(root, val)
}

func searchBSTIteration(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	for root != nil {
		if root.Val == val {
			return root
		} else if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}
