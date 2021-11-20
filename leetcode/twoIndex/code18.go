package twoIndex

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ret := make([][]int, 0)
	if len(nums) < 4 {
		return ret
	}
	sort.Ints(nums)
	index1 := 0
	for index1 < len(nums) {
		if index1 > 0 && nums[index1] == nums[index1 - 1] {
			index1 ++
			continue
		}

		index2 := index1 + 1
		for index2 < len(nums) {
			if index2 > index1 + 1 && nums[index2] == nums[index2 - 1] {
				index2 ++
				continue
			}
			left := index2 + 1
			right := len(nums) - 1
			for left < right {
				sum := nums[index1] + nums[index2] + nums[left] + nums[right]
				if sum == target {
					ret = append(ret, []int{nums[index1], nums[index2], nums[left], nums[right]})
					for left + 1 < right && nums[left] == nums[left + 1] {
						left ++
					}
					for right - 1 > 0 && nums[right] == nums[right - 1] {
						right --
					}
					left ++
					right --
				} else if sum > target {
					right --
				} else {
					left ++
				}
			}
			index2 ++
		}
		index1 ++
	}
	return ret
}

