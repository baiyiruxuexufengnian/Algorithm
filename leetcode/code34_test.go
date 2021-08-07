package leetcode

import "testing"

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 5
	r := searchRange2(nums, target)
	t.Logf("result: %v", r)
}
