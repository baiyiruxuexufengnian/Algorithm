package leetcode

import (
	"fmt"
	"sort"
)

/**
----------------------------------------------------------------------------------
40. 组合总和 II
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

注意：解集不能包含重复的组合。

输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
================================================================================
输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
]
----------------------------------------------------------------------------------
**/

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	sort.Ints(candidates)
	ret := make([][]int, len(candidates))
	for i, _ := range candidates {
		ret[i] = make([]int, 0)
	}

	for i, _ := range candidates {
		fmt.Println(candidates[i:])
		ret = append(ret, combinationSum2Core(candidates[i:], target)...)
	}
	return ret
}

func combinationSum2Core(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	if candidates[0] > target {
		return nil
	}
	ret := make([][]int, len(candidates))
	for i, _ := range candidates {
		ret[i] = make([]int, 0)
	}
	for _, v := range candidates {
		if v == target {
			ret = append(ret, []int{v})
		}
	}
	fmt.Println("ret: ", ret)
	ret = combinationSum2Core(candidates[1:], target - candidates[0])
	ret = append(ret, ret...)
	return ret
}

func Solution(candidates []int, target int) [][]int {
	return combinationSum2(candidates, target)
}
