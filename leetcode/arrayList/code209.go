package arrayList

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	if len(nums) < 0 {
		return 0
	}
	minWin := math.MaxInt32
	begin := 0
	end := -1
	sum := 0
	for end < len(nums) {
		if sum >= target {
			// 找到一个窗口
			if (end - begin + 1) < minWin {
				minWin = end - begin + 1
			}
			sum -= nums[begin]
			begin ++
		} else {
			end ++
			if end < len(nums) {
				sum += nums[end]
			}
		}
	}
	if minWin ==math.MaxInt32 {
		minWin = 0
	}
	return minWin
}
