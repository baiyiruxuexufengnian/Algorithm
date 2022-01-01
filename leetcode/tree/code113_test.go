package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	arr := make([]int, 0)
	results := make([][]int, 0)
	pathSumCore(root, targetSum, 0, arr, &results)
	return results
}

func pathSumCore(root *TreeNode, targetSum, sum int, arr []int, results *[][]int) {
	sum += root.Val
	arr = append(arr, root.Val)
	if root.Left == nil && root.Right == nil {
		if sum == targetSum {
			// arr 是一个切片，实质上就是指针嘛，在递归自动回溯的时候，
			// 原本已经保存在results里面的那个切片会被修改覆盖，所有这里需要整个
			// 副本，深拷贝将结果加入到结果集才不会有问题
			tmp := make([]int, 0)
			tmp = append(tmp, arr ...)
			*results = append(*results, tmp)
		}
		return
	}
	if root.Left != nil {
		pathSumCore(root.Left, targetSum, sum, arr, results)
	}
	if root.Right != nil {
		pathSumCore(root.Right, targetSum, sum, arr, results)
	}
}