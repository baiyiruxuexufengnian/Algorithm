package hashMap

func twoSum(nums []int, target int) []int {
	mapIndex := make(map[int]int, 0)
	for i, v := range nums {
		mapIndex[v] = i
	}

	for i, v := range nums {
		diff := target - v
		if _, ok := mapIndex[diff]; ok && mapIndex[diff] != i {
			return []int{i, mapIndex[diff]}
		}
	}
	return []int{}
}
