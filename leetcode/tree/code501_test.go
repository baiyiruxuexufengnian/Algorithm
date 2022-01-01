package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	results := make([]int, 0)
	var pre *TreeNode
	var maxCount,count int
	var findModeCore func(*TreeNode)
	findModeCore = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		// 前
		findModeCore(cur.Left)
		// 中
		if pre == nil {
			count = 1
		} else if pre.Val == cur.Val {
			count ++
		} else {
			// 与递归上一层的节点值不同，即：开始新一轮统计计数
			count = 1
		}
		pre = cur
		if maxCount == count {
			// 当前元素的频率恰好等于当前的最大频率即可入队
			results = append(results, cur.Val)
		}
		if count > maxCount {
			// 有一个新的频率最大值元素出现 clear source slice, set new maxCount
			maxCount = count
			results = results[:0]
			results = append(results, cur.Val)
		}
		// 后
		findModeCore(cur.Right)
	}
	findModeCore(root)
	return results
}
