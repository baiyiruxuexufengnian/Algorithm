package dp

var mem map[int]int = make(map[int]int)
// 和198号题目一样的，只是分情况讨论了，包不包含头部和尾部
func robII(nums []int) int {
	mem[2] = 0
	size := len(nums)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return nums[0]
	}
	if size == 2 {
		return maxRob(nums[0], nums[1])
	}
	return maxRob(robRange(nums[1:]), robRange(nums[:len(nums) - 1]))
}

func robRange(nums []int) int {
	size := len(nums)
	dp := make([]int, size)
	dp[0] = nums[0]
	dp[1] = maxRob(nums[0], nums[1])
	for i := 2; i < size; i ++ {
		dp[i] = maxRob(dp[i - 1], dp[i - 2] + nums[i])
	}
	return dp[size - 1]
}

func maxRob(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
