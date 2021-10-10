package arrayList

func moveZeroes(nums []int)  {
	if len(nums) < 1 {
		return
	}
	i := 0
	for j := i + 1; j < len(nums); j ++ {
		if nums[i] == 0 && nums[j] == 0 {
			continue
		}
		if nums[i] == 0 && nums[j] != 0 {
			nums[i] = nums[j]
			nums[j] = 0
		}
		i ++
	}
}
