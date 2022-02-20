package greedy

import "math"

func maxSubArray(nums []int) int {
	tmp := 0
	result := math.MinInt64
	for i := 0; i < len(nums); i ++ {
		tmp += nums[i]
		if tmp > result {
			result = tmp
		}
		if tmp <= 0 {
			tmp = 0
		}
	}
	return result
}
