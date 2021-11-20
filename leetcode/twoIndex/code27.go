package twoIndex

func removeElement(nums []int, val int) int {
	slowIndex, fastIndex := 0, 0
	for fastIndex < len(nums) {
		if nums[fastIndex] != val {
			nums[slowIndex] = nums[fastIndex]
			slowIndex ++
		}
		fastIndex ++
	}
	return slowIndex
}
