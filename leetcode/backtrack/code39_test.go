package backtrack

func combinationSum(candidates []int, target int) [][]int {
	results := make([][]int, 0)
	oneResult := make([]int, 0)
	sum := 0
	end := len(candidates)
	var combinationSumGreedy func(int, int)
	combinationSumGreedy = func(begin, target int) {
		if sum > target {
			return
		}
		if sum == target {
			tmp := make([]int, 0)
			tmp = append(tmp, oneResult ...)
			results = append(results, tmp)
			return
		}
		for i := begin; i < end; i ++ {
			oneResult = append(oneResult, candidates[i])
			sum += candidates[i]
			combinationSumGreedy(i, target)
			sum -= candidates[i]
			oneResult = oneResult[:len(oneResult) - 1]
		}
	}
	combinationSumGreedy(0, target)
	return results
}
