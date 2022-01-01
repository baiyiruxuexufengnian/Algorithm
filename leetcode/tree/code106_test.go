package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	return buildTreeCore(0, len(inorder) - 1, inorder,
		0, len(postorder) - 1, postorder)
}

func buildTreeCore(inorderBegin, inorderEnd int, inorder []int,
	postorderBegin, postorderEnd int, postorder []int) *TreeNode {
	if postorderBegin > postorderEnd { return nil }
	root := &TreeNode{
		Val: postorder[postorderEnd],
	}
	if postorderEnd - postorderBegin == 0 { return root }
	inIndex := findIndex(inorder, postorder[postorderEnd])
	postorderEnd --
	// 就是记录一个间隔，自己琢磨去吧
	offset := inIndex - inorderBegin - 1
	root.Left = buildTreeCore(inorderBegin, inIndex - 1, inorder,
		postorderBegin, postorderBegin + offset, postorder)
	root.Right = buildTreeCore(inIndex + 1, inorderEnd, inorder,
		postorderBegin + offset + 1, postorderEnd, postorder)
	return root
}

func findIndex(arr []int, val int) int {
	for i := 0; i < len(arr); i ++ {
		if val == arr[i] {
			return i
		}
	}
	return -1
}
