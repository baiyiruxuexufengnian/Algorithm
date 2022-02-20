package greedy

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	count := 1
	var minInt func(int, int) int
	minInt = func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}
	for i := 1; i < len(points); i ++ {
		if points[i][0] > points[i -1][1] {
			count ++
		} else {
			points[i][1] = minInt(points[i][1], points[i - 1][1])
		}
	}
	return count
}


