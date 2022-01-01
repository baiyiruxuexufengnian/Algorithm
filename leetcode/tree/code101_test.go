package tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 递归做法
	// return isSymmetricCore(root.Left, root.Right)
	// 采用队列做法的核心就是：左右子树的根节点成双成对入队，即使是空节点也是要入队的
	que := make([]*TreeNode, 0)
	que = append(que, root.Left, root.Right)
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
		// 逻辑走到这里左右子树不会是空指针，空指针的情况已经排除了
		// 入队顺序：左子树左孩子，右子树右孩子，左子树右孩子，右子树左孩子
		que = append(que, leftTree.Left, rightTree.Right, leftTree.Right, rightTree.Left)
	}
	return true
}

// 递归遍历的不是一棵二叉树了，而是在同时遍历两棵二叉树了
func isSymmetricCore(leftTreeRoot, rightTreeRoot *TreeNode) bool {
	if leftTreeRoot != nil && rightTreeRoot == nil {
		return false
	} else if leftTreeRoot == nil && rightTreeRoot != nil {
		return false
	} else if leftTreeRoot == nil && rightTreeRoot == nil {
		return true
	} else if leftTreeRoot != nil && rightTreeRoot != nil && leftTreeRoot.Val != rightTreeRoot.Val {
		return false
	}
	// 此时 leftTreeRoot 和 rightTreeRoot 不可能为空指针，前面把空指针的情况已经排除在外了
	// 此时是左右子树根节点的值相同的情况了，相同则可以继续深度递归
	outSide := isSymmetricCore(leftTreeRoot.Left, rightTreeRoot.Right)
	inSide := isSymmetricCore(leftTreeRoot.Right, rightTreeRoot.Left)
	return outSide && inSide
}
