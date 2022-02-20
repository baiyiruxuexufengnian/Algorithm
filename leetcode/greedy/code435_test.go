package greedy

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	// 统计为重叠的区间个数，默认只有一个区间时，个数是1
	count := 1
	size := len(intervals)
	// 上一个无重叠部分的区间的有边界，不能直接使用上一个点的区间作为比较参考，是因为上一个点的区间可能已经是需要被删除的一个区间，
	// 用一个需要被删除的作为参考，得到的结果是错误的，例：[[1,100],[11,22],[1,11],[2,12]]
	preRightBorder := intervals[0][1]
	for i := 1; i < size; i ++ {
		if intervals[i][0] >= preRightBorder {
			preRightBorder = intervals[i][1]
			count ++
		}
	}
	return size - count
}
