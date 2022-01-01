package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMaximumBinaryTreeCore(nums, 0, len(nums) - 1)
}

func constructMaximumBinaryTreeCore(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	// find max index
	maxIndex := left
	for i := left; i <= right; i ++ {
		if nums[i] > nums[maxIndex] { maxIndex = i }
	}
	root := &TreeNode{
		Val:nums[maxIndex],
	}
	root.Left = constructMaximumBinaryTreeCore(nums, left, maxIndex - 1)
	root.Right = constructMaximumBinaryTreeCore(nums, maxIndex + 1, right)
	return root
}
