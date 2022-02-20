package backtrack

func findSubsequences(nums []int) [][]int {
	//used := make([]bool, len(nums))
	path := make([]int, 0)
	results := make([][]int, 0)
	var findSubsequencesGreedy func(int, int)
	findSubsequencesGreedy = func(begin, end int) {
		if len(path) > 1 {
			tmp := make([]int, 0)
			tmp = append(tmp, path ...)
			results = append(results, tmp)
		}
		if begin > end {
			return
		}
		mset := make(map[int]bool, 0)
		for i := begin; i <= end; i ++ {
			pathLen := len(path)
			// if (i - 1 >= 0 && mset[nums[i]] == true && used[i - 1] == true) || (pathLen > 0 && nums[i] < path[pathLen - 1]) {
			if (mset[nums[i]] == true) || (pathLen > 0 && nums[i] < path[pathLen - 1]) {
				//used[i] = true
				continue
			}
			path = append(path, nums[i])
			//used[i] = false
			mset[nums[i]] = false
			findSubsequencesGreedy(i + 1, end)
			// greedy
			path = path[:len(path) - 1]
			//used[i] = true
			mset[nums[i]] = true
		}
	}
	findSubsequencesGreedy(0, len(nums) - 1)
	return results
}
