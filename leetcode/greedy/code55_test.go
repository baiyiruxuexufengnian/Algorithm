package greedy

//如果能到达某个位置，那一定能到达它前面的所有位置
func canJump(nums []int) bool {
	maxIndex := 0
	size := len(nums)
	if size == 1 {
		return true
	}
	for i, _ := range nums {
		index := i + nums[i]
		if maxIndex >= i && index > maxIndex {
			maxIndex = index
		}
		if maxIndex >= size - 1 {
			return true
		}
	}
	return false
}
