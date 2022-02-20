package greedy

func wiggleMaxLength(nums []int) int {
	preDiff := 0
	result := 1
	// 单调趋势是水平的话，直接平的那段抹平了看待，便于理解
	for i := 0; i < len(nums) - 1; i ++ {
		curDiff := nums[i + 1] - nums[i]
		if (curDiff > 0 && preDiff <= 0) || (curDiff < 0 && preDiff >= 0) {
			preDiff = curDiff
			result ++
		}
		// preDiff = curDiff 这一步的赋值操作一定需要放在上面条件内部，原因：看这个用例：[1,1,3,2,2,1,3]
	}
	return result
}