package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 递归和迭代效率感觉差不多
	return hasPathSumIteration(root, targetSum)
	// return hasPathSumCore(root, targetSum, 0)
}

func hasPathSumCore(root *TreeNode, targetSum int, sum int) bool {
	sum += root.Val
	if root.Left == nil && root.Right == nil {
		if targetSum == sum {
			return true
		}
		return false
	}
	var lRet, rRet bool
	if root.Left != nil {
		lRet = hasPathSumCore(root.Left, targetSum, sum)
	}
	if root.Right != nil {
		rRet = hasPathSumCore(root.Right, targetSum, sum)
	}
	return lRet || rRet
}

func hasPathSumIteration(root *TreeNode, targetSum int) bool {
	stack := make([]*TreeNode, 0)
	m := map[*TreeNode]int { root:root.Val }
	stack = append(stack, root)
	for len(stack) != 0 {
		length := len(stack)
		top := stack[length - 1]
		sum := m[top]
		if top.Left == nil && top.Right == nil && sum == targetSum {
			return true
		}
		// pop
		stack = stack[:length - 1]
		delete(m, top)
		// push
		if top.Right != nil {
			stack = append(stack, top.Right)
			m[top.Right] = sum + top.Right.Val
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
			m[top.Left] = sum + top.Left.Val
		}
	}
	return false
}
