package arrayList

import (
	"sort"
)

func sortedSquares(nums []int) []int {
	if len(nums) < 0 {
		return nums
	}
	for i, _ := range nums {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

func sortedSquares2(nums []int) []int {
	if len(nums) < 0 {
		return nums
	}
	ret := make([]int, 0)
	var i, j int
	for j < len(nums) {
		if nums[j] >= 0 {
			break
		}
		j ++
	}
	i = j - 1
	for i >= 0 && j <= len(nums) - 1 {
		if -nums[i] < nums[j] {
			ret = append(ret, -nums[i])
			i --
		} else {
			ret = append(ret, nums[j])
			j ++
		}
	}
	if i >= 0 {
		for i >= 0 {
			ret = append(ret, nums[i])
			i --
		}
	}
	if j <= len(nums) - 1 {
		ret = append(ret, nums[j:]...)
	}
	for index, v := range ret {
		ret[index] = v * v
	}
	return ret
}
