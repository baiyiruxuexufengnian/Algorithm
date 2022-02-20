package backtrack

func combinationSum3(k int, n int) [][]int {
	results := make([][]int, 0)
	oneResult := make([]int, 0)
	sum := 0
	var combinationSum3Greedy func(int, int, int)
	combinationSum3Greedy = func(n, k, begin int) {
		if sum == n && len(oneResult) == k {
			tmp := make([]int, 0)
			tmp = append(tmp, oneResult...)
			results = append(results, tmp)
		}
		for i := begin; i <= 9; i ++ {
			oneResult = append(oneResult, i)
			sum += i
			combinationSum3Greedy(n, k, i + 1)
			// greedy
			oneResult = oneResult[:len(oneResult) - 1]
			sum -= i
		}
	}
	combinationSum3Greedy(n, k, 1)
	return results
}
