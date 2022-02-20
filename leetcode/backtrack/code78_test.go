package backtrack

func subsets(nums []int) [][]int {
	results := make([][]int, 0)
	path := make([]int, 0)
	var subsetsCore func(int, int)
	subsetsCore = func(begin, end int) {
		tmp := make([]int, 0)
		tmp =  append(tmp, path ...)
		results = append(results, tmp)
		if begin > end {
			return
		}
		for i := begin; i <= end; i ++ {
			path = append(path, nums[i])
			subsetsCore(i + 1, end)
			// greedy
			path = path[:len(path) - 1]
		}
	}
	subsetsCore(0, len(nums) - 1)
	return results
}
