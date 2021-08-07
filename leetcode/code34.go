package leetcode

import "log"

func searchRange(nums []int, target int) []int {
	l := findLeft(nums, target)
	r := findRight(nums, target)
	log.Printf("l = %v, r = %v", l, r)
	if l == -2 || r == -2 {
		return []int{-1, -1}
	}

	if r - l >= 2 {
		return []int{l + 1, r - 1}
	}

	return []int{-1, -1}
}

func findLeft(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	ret  := -2
	for l <= r {
		pos := l + (r -l) / 2
		if nums[pos] >= target {
			r = pos - 1
			ret = r
		} else {
			l = pos + 1
		}
	}
	return ret
}

func findRight(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	ret  := -2
	for l <= r {
		pos := l + (r -l) / 2
		if nums[pos] <= target {
			l = pos + 1
			ret = l
		} else {
			r = pos - 1
		}
	}
	return ret
}

func searchRange2(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		pos := l + (r -l) / 2
		if nums[pos] == target {
			// 处理相等的情况
			l = pos
			r = pos
			for l >= 0 && nums[l] == target {
				l --
			}
			for r <= len(nums) - 1  && nums[r] == target {
				r ++
			}
			return []int{l + 1, r - 1}
		} else if nums[pos] > target {
			r = pos - 1
		} else {
			l = pos + 1
		}
	}
	return []int{-1, -1}
}