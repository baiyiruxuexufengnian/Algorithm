package leetcode

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法

例：
输入: nums = [1,3,5,6], target = 5
输出: 2

输入: nums = [1,3,5,6], target = 2
输出: 1

输入: nums = [1,3,5,6], target = 7
输出: 4

输入: nums = [1,3,5,6], target = 0
输出: 0

输入: nums = [1], target = 0
输出: 0
*/

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	ret := 0
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l >= r {
		if nums[l] >= target {
			ret = l
		} else {
			ret = l + 1
		}
	}
	return ret
}
