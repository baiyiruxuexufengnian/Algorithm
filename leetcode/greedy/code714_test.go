package greedy

// 画折线图进行推导帮助理解，要求最大利润，就是找到所有单调递增的折现片段，同时满足售价大于买价加手续费，
// 将这些端累加起来，最小买价需要记录更新
func maxProfitII(prices []int, fee int) int {
	minPrices := prices[0]
	result := 0
	for i := 0; i < len(prices); i ++ {
		if prices[i] < minPrices {
			minPrices = prices[i]
		}
		// 无法获取利润的价格直接跳过,手续费是可以为0的
		if prices[i] >= minPrices && prices[i] <= fee + minPrices {
			continue
		}
		if prices[i] > minPrices && prices[i] > fee + minPrices {
			result += prices[i] - minPrices - fee
			// result累加只是单调递增线上各个线段的增量值,，为什么这里最小价格需要上移这么多，很难说清楚，就是画折线图加数学公式推导
			minPrices = prices[i] - fee
		}
	}
	return result
}
