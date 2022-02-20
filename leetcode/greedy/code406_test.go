package greedy

import (
	"sort"
)

// 其实就是重新排序，排序有两个比较规则：
// 1、身高
// 2、有k个身高更高或者相同的人排在第i个人的前面(因为这个规则说有k个不低于他的人在前面，
// 那么一定是需要先按照身高排序，让高的先在前面去，然后因为是有k个，那么直接从高的开始遍历，
// 按k作为下标，进行插入即可，这样即能保证规则二的满足)
func reconstructQueue(people [][]int) [][]int {
	//先将身高从大到小排序，确定最大个子的相对位置
	sort.Slice(people,func(i,j int)bool{
		if people[i][0]==people[j][0]{
			return people[i][1]<people[j][1]//这个才是当身高相同时，将K按照从小到大排序
		}
		return people[i][0]>people[j][0]//这个只是确保身高按照由大到小的顺序来排，并不确定K是按照从小到大排序的
	})
	size := len(people)
	que := make([][]int, 0)

	var sliceInsert func(int, []int)
	sliceInsert = func(i int, value []int) {
		if len(que) == 0 {
			que = append(que, value)
		} else {
			tmp := append([][]int{}, que[i:]...)
			que = append(que[:i], value)
			que = append(que, tmp ...)
		}
	}
	// 按照第二条规则进行插入排序
	for i := 0; i < size; i ++ {
		sliceInsert(people[i][1], people[i])
	}
	return que
}
