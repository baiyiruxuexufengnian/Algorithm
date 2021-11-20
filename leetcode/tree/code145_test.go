package tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	arr := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		top := stack[len(stack) - 1]
		arr = append(arr, top.Val)
		stack = stack[:len(stack) - 1]
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}

	left := 0
	right := len(arr) - 1
	for left < right {
		tmp := arr[left]
		arr[left] = arr[right]
		arr[right] = tmp
		left ++
		right --
	}
	return arr
}

func postorderTraversalCore(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	arr := make([]int, 0)
	arr = append(arr, postorderTraversalCore(node.Left)...)
	arr = append(arr, postorderTraversalCore(node.Right)...)
	arr = append(arr, node.Val)
	return arr
}

