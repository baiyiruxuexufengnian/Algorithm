package tree

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	arr := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		length := len(stack)
		top := stack[length - 1]
		arr = append(arr, top.Val)
		stack = stack[:length - 1]
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return arr
}
