package arrayList

// 写的太罗嗦了，效率上是一样的
func removeDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	var i, j int
	for j < len(nums) && i + 1 < len(nums) {
		if nums[i] == nums[i + 1] {
			break
		}
		i ++
		j ++
	}

	for j < len(nums) {
		if nums[j] != nums[i] {
			nums[i + 1] = nums[j]
			i ++
		}
		j ++
	}
	return i + 1
}

func removeDuplicates2(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	i := 0
	for j := i + 1; j < len(nums); j ++ {
		if nums[i] != nums[j] {
			i ++
			nums[i] = nums[j]
		}
	}
	return i + 1
}