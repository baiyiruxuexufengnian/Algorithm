package arrayList

// 序列本身有序，如果整个序列每个元素都出现两次，那么相邻的两个相等元素一定是：偶数下标的值等于奇数下标的值（先偶后奇）
// 一旦有一个元素只出现一次，那么将破坏这一原则，即该元素出现的一侧，将导致相邻的两个相等元素，奇数下标的值等于偶数下标的值（先奇后偶）
func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums) - 1
	for left < right {
		mid := left + (right - left) / 2
		// mid 作为偶数下标
		if mid % 2 == 0 {
			if nums[mid] != nums[mid + 1] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid - 1] != nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		}
	}
	return nums[left]
}
