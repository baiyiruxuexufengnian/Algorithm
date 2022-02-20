package dp

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// dp 数组：dp[i][j] 表示从start到point(i, j)共多少种路径
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i ++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i ++ {
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
			break
		}
		dp[i][0] = 1
	}
	for j := 0; j < n && obstacleGrid[0][j] == 0; j ++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i ++ {
		for j := 1; j < n; j ++ {
			// 如果是障碍物则到该点的路径共0种路径，不需要做动态推导
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
		}
	}
	return dp[m - 1][n - 1]
}
