package tree

// 这是递归版
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func invertTreeUseQueue(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) != 0 {
		size := len(que)
		for i := 0; i < size; i ++ {
			que[i].Left, que[i].Right = que[i].Right, que[i].Left
			if que[i].Left != nil {
				que = append(que, que[i].Left)
			}
			if que[i].Right != nil {
				que = append(que, que[i].Right)
			}
		}
		// 前size个元素依次出队列
		que = que[size:]
	}
	return root
}

func invertTreeUseStack(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		l := len(stack)
		top := stack[l - 1]
		top.Left, top.Right = top.Right, top.Left
		// pop stack
		stack = stack[: l - 1]
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return root
}