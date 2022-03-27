package dp

func rob(nums []int) int {
	// dp数组：dp[i]表示，从第0~i个house中打劫可获取得到的最大利润为：dp[i]
	// 递推公式：dp[i] = max((dp[i - 2] + nums[i]), dp[i - 1])
	// 解释：第i-1这个house因为与第i这个house相邻，此时要打劫第i这个house(nums[i])那么就不能打劫第i-1这个house
	// 只能打劫0到i-2双闭区间的house，即：dp[i - 2]+nums[i]，若是不打劫第i这个house，那么直接取0到i-1这段的最大利润即可
	// 从公式可以推出，公式有意义的条件就是i从2开始
	var max func(int, int) int
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	// 1 <= size <= 100
	size := len(nums)
	if size <= 1 {
		return nums[0]
	}
	dp := make([]int, size)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < size; i ++ {
		dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
	}
	return dp[size - 1]
}
