package tree

// 外侧与外侧相比较，内测与内测相比较
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	} else if p != nil && q == nil {
		return false
	} else if p == nil && q == nil {
		return true
	} else if p != nil && q != nil && p.Val != q.Val {
		return false
	}

	// 此时排除了所有左右子树根节点是空的情况，且此时是左右根节点的值相等的时候，不相等上面直接返回false好了
	// 左子树的内侧和右子树的内测相比较
	inCompareIn := isSameTree(p.Left, q.Left)
	// 左子树的外测与右子树的外侧相比较
	outCompareOut := isSameTree(p.Right, q.Right)
	return inCompareIn && outCompareOut
}

// 外侧与外侧相比较，内测与内测相比较
func isSameTreeUseQueue(p *TreeNode, q *TreeNode) bool {
	que := make([]*TreeNode, 0)
	que = append(que, p, q)
	for len(que) != 0 {
		leftTree := que[0]
		rightTree := que[1]
		// pop
		que = que[2:]
		if leftTree == nil && rightTree == nil {
			continue
		}
		if leftTree == nil || rightTree == nil || leftTree.Val != rightTree.Val {
			return false
		}
		que = append(que, leftTree.Left, rightTree.Left, leftTree.Right, rightTree.Right)
	}
	return true
}
