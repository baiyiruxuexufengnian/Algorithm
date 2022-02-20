package greedy

// 其实就是要满足左右两个规则
// 左规则：AB相邻，若是B>A，则B比A分的糖多， 若是A>B，则保持一样
// 右规则：AB相邻，若是A>B，则A比B分的糖多， 若是B>A，则保持一样
// 两个规则需要同时满足
func candy(ratings []int) int {
	size := len(ratings)
	sum := 0
	nums := make([]int, 0)
	for i := 0; i < size; i ++ {
		nums = append(nums, 1)
	}
	// Left rule
	for i := 0; i < size; i++ {
		if i - 1 > -1 && ratings[i] > ratings[i - 1] {
			nums[i] = nums[i - 1] + 1
		}
	}
	var maxInt func(int, int) int
	maxInt = func(i int, i2 int) int {
		if i > i2 {
			return i
		}
		return i2
	}
	// Right rule
	for i := size - 2; i >= 0; i -- {
		if ratings[i] > ratings[i + 1] {
			nums[i] = maxInt(nums[i], nums[i + 1] + 1)
		}
	}
	for i := 0; i < size; i ++ {
		sum += nums[i]
	}
	return sum
}
