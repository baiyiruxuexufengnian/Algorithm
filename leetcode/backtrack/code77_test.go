package backtrack

import (
	"testing"
)

func combine(n int, k int) [][]int {
	results := make([][]int, 0)
	oneResult := make([]int, 0)
	var combineGreedy func(int, int, int)
	combineGreedy = func(n int, k int, startIndex int) {
		if len(oneResult) == k {
			tmp := make([]int, 0)
			tmp = append(tmp, oneResult...)
			results = append(results, tmp)
			return
		}
		for i := startIndex; i <= n; i ++ {
			oneResult = append(oneResult, i)
			combineGreedy(n, k, i + 1)
			// 递归出来了, 对单次结果集做回溯
			oneResult = oneResult[:len(oneResult) - 1]
		}
	}
	combineGreedy(n, k, 1)
	return results
}

func TestCombine(t *testing.T) {
	results := combine(4, 2)
	t.Log(results)
}
