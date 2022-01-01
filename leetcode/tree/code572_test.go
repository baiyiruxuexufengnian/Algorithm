package tree

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root != nil && subRoot == nil {
		return false
	}
	nodes := findNode(root, subRoot)
	for _, v := range nodes {
		res := isSubTreeCore(v, subRoot)
		if res == true {
			return true
		}
	}
	return false
}

func isSubTreeCore(node, subRoot *TreeNode) bool {
	if node == nil && subRoot != nil {
		return false
	} else if node != nil && subRoot == nil {
		return false
	} else if node == nil && subRoot == nil {
		return true
	} else if node != nil  && subRoot != nil && node.Val != subRoot.Val {
		return false
	}
	return isSubTreeCore(node.Left, subRoot.Left) && isSubTreeCore(node.Right, subRoot.Right)
}

func findNode(root, target *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	res := make([]*TreeNode, 0)
	if root.Val == target.Val {
		res = append(res, root)
	}
	res = append(res, findNode(root.Left, target)...)
	res = append(res, findNode(root.Right, target)...)
	return res
}
