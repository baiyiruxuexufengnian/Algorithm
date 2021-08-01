package leetcode

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	r := search(nums, target)
	t.Logf("result: %v", r)
}
