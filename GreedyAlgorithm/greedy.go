package GreedyAlgorithm

import "sort"

// 贪心算法系列

func findContentChildren(g []int, s []int) int {
	child := 0
	cookie := 0
	sort.Ints(g)
	sort.Ints(s)
	for child < len(g) && cookie < len(s) {
		if g[child] <= s[cookie] { child ++ }
		cookie ++
	}
	return child
}
