package greedy

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		if math.Abs(float64(nums[i])) > math.Abs(float64(nums[j])) {
			return true
		}
		return false
	})
	size := len(nums)
	for i := 0; i < size; i ++ {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k -= 1
		}
	}
	if k % 2 == 1 { nums[size - 1] *= -1 }
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}
