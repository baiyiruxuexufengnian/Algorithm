package greedy

import "sort"

// func merge(intervals [][]int) [][]int {
//     sort.Slice(intervals, func(i, j int) bool {
//         return intervals[i][1] < intervals[j][1]
//     })
//     size := len(intervals)
//     results := make([][]int, 0)
//     var minInt func(int, int) int
//     minInt = func(i1, i2 int) int {
//         if i1 > i2 {
//             return i2
//         }
//         return i1
//     }
//     flag := false
//     for i := size - 2; i >=0; i -- {
//         // 判断是否重叠，重叠就需要合并，合并的话需要合并多个区间，需要使用for循环
//         end := intervals[i+1][1]
//         start := intervals[i+1][0] // 上一个区间的左边界
//         for i >= 0 && intervals[i][1] >= start {
//             // 此时该区间的右边界，若是大于等于上一个区间的左边界，说明需要合并，即不断更新区间的左边界
//             start = minInt(intervals[i][0], start)
//             if i == 0 {
//                 flag = true
//             }
//             i --
//         }
//         results = append(results, []int{start, end})
//     }
//     if flag == false {
//         results = append(results, intervals[0])
//     }
//     sort.Slice(results, func(i, j int) bool {
//         return results[i][0] < results[j][0]
//     })
//     return results
// }

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	size := len(intervals)
	results := make([][]int, 0)
	var maxInt func(int, int) int
	maxInt = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	flag := false
	for i := 1; i < size; i ++ {
		start := intervals[i - 1][0]    // 上一个区间的左边界
		end := intervals[i- 1][1]    // 上一个区间的右边界
		for i < size && intervals[i][0] <= end {
			end = maxInt(end, intervals[i][1])
			if i == size - 1 {
				flag = true
			}
			i ++
		}
		results = append(results, []int{start, end})
	}
	if flag == false {
		results = append(results, intervals[size - 1])
	}
	return results
}
