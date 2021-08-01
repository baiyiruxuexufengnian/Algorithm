package leetcode

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l < len(nums) && r < len(nums) && l <= r {
		pos := l + (r-l)/2
		if nums[pos] == target {
			return pos
		} else if nums[pos] > target {
			r = pos - 1
		} else {
			l = pos + 1
		}
	}
	return -1
}
