package dp

// f(i, c) = max(f(i-1, c - weight[i]) + value[i], f(i, c))
func canPartition(nums []int) bool {
	// nums 每个元素代表01背包中每个商品的价值，也是商品的重量
	// nums元素总和的一半代表背包的空间大小
	sum := 0
	size := len(nums)
	for i := 0; i < size; i ++ {
		sum += nums[i]
	}
	if sum & 1 == 1 {
		return false
	}
	sum /= 2
	dp := make([][]int, 2)
	for i := 0; i < 2; i ++ {
		dp[i] = make([]int, sum + 1)
	}
	var max func(int, int) int
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	// 初始化第一行
	for j := 0; j <= sum; j ++ {
		// 商品价值和重量均为:nums[0]
		if j >= nums[0] {
			dp[0][j] = nums[0]
		}
	}
	for i := 1; i < size; i ++ {
		weight,value := nums[i], nums[i]
		for j := 0; j <= sum; j ++ {
			if j < weight {
				dp[i & 1][j] = dp[(i - 1) & 1][j]
			} else {
				dp[i & 1][j] = max( dp[(i - 1) & 1][j], value + dp[(i - 1) & 1][j - weight])
			}
		}
	}
	return dp[(size - 1) & 1][sum] == sum
}
