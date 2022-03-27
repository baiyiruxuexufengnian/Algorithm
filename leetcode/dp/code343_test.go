package dp

// 记忆化搜索版本
// func integerBreak(n int) int {
//     var max func(int, int) int
//     max = func(i1, i2 int) int {
//         if i1 > i2 {
//             return i1
//         }
//         return i2
//     }
//     mem := make([]int, n + 1)
//     for i, _ := range mem {
//         mem[i] = -1
//     }
//     var integerBreakCore func(int) int
//     integerBreakCore = func(n int) int {
//         if n == 1 {
//             return 1
//         }
//         result := -1
//         if mem[n] != -1 {
//             return mem[n]
//         }
//         for i := 1; i <= n-1; i ++ {
//             result = max(max(result, i * (n - i)), i * integerBreakCore(n - i))
//         }
//         mem[n] = result
//         return mem[n]
//     }
//     return integerBreakCore(n)
// }

// dp 版本
func integerBreak(n int) int {
	// dp数组dp[n+1],dp[i]表示每第i个元素被拆分成至少两个元素后可取得的最大乘积值
	dp := make([]int, n + 1)
	// dp[0],dp[1]不理会是因为完全没有意义这两个值，题目要求了n大于等于2
	dp[2] = 1
	var max func(int, int) int
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	for i := 3; i <= n; i ++ {
		// 枚举j的时候，是从1开始的。i是从3开始，这样dp[i - j]就是dp[2]正好可以通过我们初始化的数值求出来
		for j := 1; j < i; j ++ {
			dp[i] = max(dp[i], max(j * (i - j), j * dp[i - j]))
		}
	}
	return dp[n]
}
