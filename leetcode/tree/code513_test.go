package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	result := -1
	if root == nil {
		return result
	}
	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) != 0 {
		size := len(que)
		result = que[0].Val
		for i := 0; i < size; i ++ {
			if que[i].Left != nil {
				que = append(que, que[i].Left)
			}
			if que[i].Right != nil {
				que = append(que, que[i].Right)
			}
		}
		// pop
		que = que[size:]
	}
	return result
}
