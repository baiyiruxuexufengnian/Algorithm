package backtrack

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	length := len(candidates)
	results := make([][]int, 0)
	oneResult := make([]int, 0)
	used := make([]bool, length)
	sum := 0
	sort.Ints(candidates)
	var combinationSum2Greedy func(int, int)
	combinationSum2Greedy = func(begin, target int) {
		if sum > target {
			return
		}
		if sum == target {
			tmp := make([]int, 0)
			tmp = append(tmp, oneResult...)
			results = append(results, tmp)
			return
		}
		for i := begin; i < length; i ++ {
			if i - 1 >= 0 && candidates[i] == candidates[i - 1] && used[i - 1] == false {
				continue
			}
			used[i] = true
			sum += candidates[i]
			oneResult = append(oneResult, candidates[i])
			combinationSum2Greedy(i + 1, target)
			oneResult = oneResult[:len(oneResult) - 1]
			used[i] = false
			sum -= candidates[i]
		}
	}
	combinationSum2Greedy(0, target)
	return results
}

// 改进易懂点
func combinationSum2Self(candidates []int, target int) [][]int {
	length := len(candidates)
	results := make([][]int, 0)
	oneResult := make([]int, 0)
	// 用于记录递归树每一层已经使用过的元素
	used := make([]bool, length)
	sum := 0
	sort.Ints(candidates)
	var combinationSum2Greedy func(int, int)
	combinationSum2Greedy = func(begin, target int) {
		if sum > target {
			return
		}
		if sum == target {
			tmp := make([]int, 0)
			tmp = append(tmp, oneResult...)
			results = append(results, tmp)
			return
		}
		for i := begin; i < length; i ++ {
			if i - 1 >= 0 && candidates[i] == candidates[i - 1] && used[i - 1] == true {
				used[i] = true
				continue
			}
			used[i] = false
			sum += candidates[i]
			oneResult = append(oneResult, candidates[i])
			combinationSum2Greedy(i + 1, target)
			oneResult = oneResult[:len(oneResult) - 1]
			used[i] = true
			sum -= candidates[i]
		}
	}
	combinationSum2Greedy(0, target)
	return results
}