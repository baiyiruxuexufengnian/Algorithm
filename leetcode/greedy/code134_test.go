package greedy

// 两个关键点：
// 1、gas总和即总油量要大于cost总和即耗油量才有解
// 2、rest表示从每个i点出发，所需要的剩余油量，若是值小于0则表示，从该点出发无法闭环可达，需要从i+1尝试出发
// 若是从该点出发不断累加剩余油量，当后面有更大的负数值导致累加的值变成负数，则需要将累加值复位为：0，然后从i+1位置开始
func canCompleteCircuit(gas []int, cost []int) int {
	size := len(gas)
	// 默认是从初始未知出发的
	start := 0
	totalSum, partSum := 0, 0
	for i := 0; i < size; i ++ {
		rest := gas[i] - cost[i]
		totalSum += rest
		partSum += rest
		if partSum < 0 {
			start = i + 1
			partSum = 0
		}
	}
	if totalSum < 0 {
		start = -1
	}
	return start
}
