package tree

import "math"

func isBalanced(root *TreeNode) bool {
	if isBalancedCore(root) == -1 {
		return false
	}
	return true
}

// 获取树的高度差
func isBalancedCore(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := isBalancedCore(root.Left)
	if l == -1 {
		return -1
	}
	r := isBalancedCore(root.Right)
	if r == -1 {
		return -1
	}
	if math.Abs(float64(l - r)) > 1 {
		return -1
	} else {
		if l > r {
			return l + 1
		}
		return r + 1
	}
}

// 通过栈实现
func getBalanceDepth(root *TreeNode) int {
	depth := 0
	maxDepth := 0
	stack := make([]*TreeNode, 0)
	cur := root
	for len(stack) != 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
			depth ++
		} else {
			length := len(stack)
			top := stack[length - 1]
			// pop
			stack = stack[:length - 1]
			depth --
			cur = top.Right
		}
		if depth > maxDepth {
			maxDepth = depth
		}
	}
	return depth
}

// 通过队列实现
func getBalanceDepth2(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) != 0 {
		size := len(que)
		depth ++
		for i := 0; i < size; i ++ {
			if que[i].Left != nil {
				que = append(que, que[i].Left)
			}
			if que[i].Right != nil {
				que = append(que, que[i].Right)
			}
		}
		que = que[size:]
	}
	return depth
}

// 非递归
func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		length := len(stack)
		top := stack[length - 1]
		if math.Abs(float64(getBalanceDepth(top.Left) - getBalanceDepth(top.Right))) > 1 {
			return false
		}
		// pop
		stack = stack[:length - 1]
		if top.Right != nil { stack = append(stack, top.Right) }
		if top.Left != nil { stack = append(stack, top.Left) }
	}
	return true
}
