package twoIndex

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) < 3 {
		return ret
	}

	sort.Ints(nums)
	for index, _ := range nums {
		if index > 0 && nums[index] == nums[index - 1] {
			continue
		}
		left := index + 1
		right := len(nums) - 1
		for left < right {
			if nums[index] > 0 {
				break
			}
			sum := nums[index] + nums[left] + nums[right]
			if sum == 0 {
				ret = append(ret, []int{nums[index], nums[left], nums[right]})
				for left < right && nums[left] == nums[left + 1] {
					left ++
				}
				for right > left && nums[right] == nums[right - 1] {
					right --
				}
				left ++
				right --
			} else if sum > 0 {
				right --
			} else {
				left ++
			}
		}
	}
	return ret
}
