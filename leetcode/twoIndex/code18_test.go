package twoIndex

import (
	"testing"
)

func TestFourSum(t *testing.T) {
	nums := []int{1,0,-1,0,-2,2}
	t.Log(fourSum(nums, 0))

	nums = []int{2,2,2,2,2}
	t.Log(fourSum(nums, 8))

	nums = []int{1,-2,-5,-4,-3,3,3,5}
	t.Log(fourSum(nums, -11))
}
