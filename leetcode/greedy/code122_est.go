package greedy

// 核心：假设p[3]-p[1]是最大利润，那么：p[3]-p[1] = (p[1]-p[2])+(p[2]-p[1])+(p[3]-p[2])
func maxProfit(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i ++ {
		diff := prices[i] - prices[i - 1]
		if diff > 0 {
			result += diff
		}
	}
	return result
}
