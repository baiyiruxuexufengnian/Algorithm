package tree

func inorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	if root == nil {
		return arr
	}
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			arr = append(arr, top.Val)
			cur = top.Right
		}
	}
	return arr
}

func inorderTraversalCore(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	arr := make([]int, 0)
	arr = append(arr, inorderTraversalCore(node.Left)...)
	arr = append(arr, node.Val)
	arr = append(arr, inorderTraversalCore(node.Right)...)
	return arr
}
