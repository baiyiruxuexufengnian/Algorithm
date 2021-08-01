package leetcode

import "log"

func searchRange(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		pos := l + (r-l)/2
		if target == nums[pos] {
			l = pos
			r = pos
			ret := make([]int, 0)
			log.Printf("l=%v, r=%v, pos=%v", l, r, pos)
			for target == nums[l] {
				if l <= 0 || nums[l] != target {
					ret = append(ret, l)
				}
				l--
			}
			for target == nums[r] {
				if r >= len(nums)-1 || nums[r] != target {
					ret = append(ret, r)
				}
				r++
			}
			return ret
		} else if nums[pos] > target {
			r = pos - 1
		} else {
			l = pos + 1
		}
	}
	return []int{-1, -1}
}
