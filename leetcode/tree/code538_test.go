package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归版
func convertBST1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	pre := 0
	var convertBSTCore func(*TreeNode)
	convertBSTCore = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		// 开始反中序遍历
		convertBSTCore(cur.Right)
		cur.Val += pre
		pre = cur.Val
		convertBSTCore(cur.Left)
	}
	convertBSTCore(root)
	return root
}

func convertBST(root *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0)
	pre := 0
	cur := root
	for len(stack) != 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		} else {
			// top
			length := len(stack)
			top := stack[length - 1]
			top.Val += pre
			pre = top.Val
            // pop
            stack = stack[:length - 1]
            cur = top.Left
		}
	}
	return root
}
