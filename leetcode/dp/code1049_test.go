package dp

// 01背包具备几大要素：背包容量、物品重量、物品价值，诉求为获取最大价值
// 本题物品重量对应：石头重量，物品价值对应：石头重量
// 因为两个物品相撞要么是两者直接粉碎y-x==0，要么是重量大的那个留下残体：y-x>0作为新的一个石头继续战斗相撞
// 所以假设尽可能的将这堆石头分为两堆PK，最后剩下的石头只会有一个，因为在PK过程中即使质量大赢了业会是一个残体
// 作为一个新的石头继续PK，所以直接将两堆看作一个整体来进行PK
// 因为要尽可能的将其分为均等的两堆，所以需要一个一直标准容器进行分堆操作，这个容器就看做是背包，因为石头
// 每个个体是不可切割的，所以这个容器只能是整数容量，不可能是小数，因为要尽可能的分成两堆，所以这个容器的容量为：sum(stones) / 2
func lastStoneWeightII(stones []int) int {
	sum := 0
	size := len(stones)
	for i := 0; i < size; i ++ {
		sum += stones[i]
	}
	// dp[i][j] ==> weight:stones[i], capacity:j
	dp := make([][]int, 2)
	for i := 0; i < 2; i ++ {
		dp[i] = make([]int, (sum / 2) + 1)
	}
	var max func(int, int) int
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	for j := 0; j <= (sum / 2); j ++ {
		if j >= stones[0] {
			dp[0][j] = stones[0]
		}
	}
	for i := 1; i < size; i ++ {
		for j := 0; j <= (sum / 2); j ++ {
			if j < stones[i] {
				dp[i & 1][j] = dp[(i - 1) & 1][j]
			} else {
				dp[i & 1][j] = max(dp[(i - 1) & 1][j], dp[(i - 1) & 1][j - stones[i]] + stones[i])
			}
		}
	}
	return sum - 2 * dp[(size - 1) & 1][sum / 2]
}

