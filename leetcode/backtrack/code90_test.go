package backtrack

import "sort"

func subsetsWithDup(nums []int) [][]int {
	length := len(nums)
	results := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, length)
	var subsetsWithDupCore func(int, int)
	sort.Ints(nums)
	subsetsWithDupCore = func(begin, end int) {
		tmp := make([]int, 0)
		tmp = append(tmp, path ...)
		results = append(results, tmp)
		if begin > end {
			return
		}
		for i := begin; i <= end; i ++ {
			if i - 1 >= 0 && nums[i] == nums[i - 1] && used[i - 1] == true {
				used[i] = true
				continue
			}
			used[i] = false
			path = append(path, nums[i])
			subsetsWithDupCore(i + 1, end)
			// greedy
			used[i] = true
			path = path[:len(path) - 1]
		}
	}
	subsetsWithDupCore(0, length - 1)
	return results
}
