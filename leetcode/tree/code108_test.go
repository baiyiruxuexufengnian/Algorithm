package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var sortedArrayToBSTCore func(int, int) *TreeNode
	sortedArrayToBSTCore = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right - left) / 2
		root := &TreeNode{
			Val:nums[mid],
		}
		root.Left = sortedArrayToBSTCore(left, mid - 1)
		root.Right = sortedArrayToBSTCore(mid + 1, right)
		return root
	}
	return sortedArrayToBSTCore(0, len(nums) - 1)
}
