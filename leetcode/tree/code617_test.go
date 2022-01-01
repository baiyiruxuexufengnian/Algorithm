package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 采用递归的方式好一些
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// 此时 root1 和 root2 均不可能为空

	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

// 核心就是两个树的节点成双成对的加入到队列中去
func mergeTreesIteration(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root2 == nil {
		return root1
	}
	if root1 == nil {
		return root2
	}
	que := make([]*TreeNode, 0)
	que = append(que, root1, root2)
	for len(que) != 0 {
		top1 := que[0]
		top2 := que[1]
		// 队列里面的节点都不会为空，不会往里面放的空节点
		top1.Val += top2.Val
		// pop
		que = que[2:]
		if top1.Left != nil && top2.Left != nil  {
			que = append(que, top1.Left, top2.Left)
		}
		if top1.Right != nil && top2.Right != nil  {
			que = append(que, top1.Right, top2.Right)
		}
		// 第一个树的队顶节点如果不是空，而第二个树的为空，说明什么也不用做
		if top1.Left == nil && top2.Left != nil {
			top1.Left = top2.Left
		}
		if top1.Right == nil && top2.Right != nil {
			top1.Right = top2.Right
		}
	}
	return root1
}
