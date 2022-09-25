package hot100

import "fmt"

// 560. 和为 K 的子数组
func subarraySum(nums []int, k int) int {
	prefixMap := make(map[int]int)
	prefixMap[0] = 1
	sum := 0
	count := 0
	for _, v := range nums {
		sum += v
		if vv, ok := prefixMap[sum - k]; ok {
			count += vv
		}
		prefixMap[sum] ++
	}
	return count
}

// 494. 目标和
// Method1: 递归回溯
func _findTargetSumWays(nums []int, target int) int {
	if len(nums) <= 0 { return 0 }
	var dfs func([]int, int, int) int
	dfs = func(nums []int, index, curSum int) int {
		if index == len(nums) {
			if curSum == target { return 1 }
			return 0
		}
		lRes := dfs(nums, index + 1, curSum + nums[index])
		rRes := dfs(nums, index + 1, curSum - nums[index])
		return lRes + rRes
	}
	return dfs(nums, 0, 0)
}
// Method2: 记忆化搜索
func findTargetSumWays(nums []int, target int) int {
	if len(nums) <= 0 { return 0 }
	remembers := make(map[string]int)
	var dfs func([]int, int, int) int
	dfs = func(nums []int, index, curSum int) int {
		if index == len(nums) {
			if curSum == target { return 1 }
			return 0
		}
		if val, ok := remembers[fmt.Sprintf("%v_%v", index, curSum)]; ok {
			return val
		}
		lRes := dfs(nums, index + 1, curSum + nums[index])
		rRes := dfs(nums, index + 1, curSum - nums[index])
		remembers[fmt.Sprintf("%v_%v", index, curSum)] = lRes + rRes
		return lRes + rRes
	}
	return dfs(nums, 0, 0)
}
