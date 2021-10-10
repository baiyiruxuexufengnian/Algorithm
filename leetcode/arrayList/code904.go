package arrayList

// 这道题目可以理解为求只包含两种元素的最长连续子序列

func totalFruit(fruits []int) int {
	if len(fruits) == 0 {
		return 0
	}
	basket := make(map[int]int, 0)
	i := 0
	j := 0
	maxWin := -2
	for j < len(fruits) {
		basket[fruits[j]] ++
		for len(basket) > 2 {
			basket[fruits[i]] --
			if basket[fruits[i]] == 0 {
				delete(basket, fruits[i])
			}
			i ++
		}
		if j - i + 1 > maxWin {
			maxWin = j - i + 1
		}
		j ++
	}
	return maxWin
}