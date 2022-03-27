package dp

import (
	"testing"
)

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i ++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j ++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i ++ {
		for j := 1; j < n; j ++ {
			dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
		}
	}
	return dp[m - 1][n - 1]
}

func findMinPath(in [][]int) int {
	row := len(in)
	col := len(in[0])
	results := make([]int, 0)
	visited := make([][]int, row)
	for i := 0; i < row; i ++ {
		visited[i] = make([]int, col)
	}
	sum := 0
	var findMinPathBacktrack func(int, int)
	findMinPathBacktrack = func(i int, j int) {
		if i < 0 || i >= row || j < 0 || j >= col || visited[i][j] == 1 {
			return
		}
		sum += in[i][j]
		visited[i][j] = 1
		if i == row - 1 && j == col - 1 {
			results = append(results, sum)
			return
		}
		findMinPathBacktrack(i, j + 1)
		findMinPathBacktrack(i + 1, j)
		findMinPathBacktrack(i, j - 1)
		findMinPathBacktrack(i - 1, j)
		visited[i][j] = 0
		sum -= in[i][j]
		return
	}
	findMinPathBacktrack(0, 0)
	sum = 0
	for _, v := range results {
		sum = min(sum, v)
	}
	return sum
}

func TestFind(t *testing.T) {
	in := [][]int{
		   {1,9,3,4,5},
		   {1,9,1,1,1},
		   {1,1,1,9,1},
	}
	t.Logf("value: %v", findMinPath(in))
}


func min(i, j int) int {
	if i > j {
		return i
	}
	return j
}

