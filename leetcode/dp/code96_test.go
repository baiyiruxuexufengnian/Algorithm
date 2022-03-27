package dp

// https://leetcode-cn.com/problems/unique-binary-search-trees/solution/shou-hua-tu-jie-san-chong-xie-fa-dp-di-gui-ji-yi-h/
func numTrees(n int) int {
	// 1到i为节点组成的二叉搜索树的个数为dp[i]
	dp := make([]int, n + 1) // 因为要求n即包含n，所以要求到n的下一位
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i ++ {
		for j := 0; j <= i - 1; j ++ {
			dp[i] += dp[j]*dp[i - j - 1]
		}
	}
	return dp[n]
}
