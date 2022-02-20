package greedy

import "strconv"

// 这是一道很明显的贪心题目。既然要尽可能的大，那么这个数从高位开始要尽可能地保持不变。
//那么我们找到从高到低第一个满足 str[i] > str[i+1]str[i]>str[i+1] 的位置，
//然后把 str[i] - 1str[i]−1 ，再把后面的位置都变成 99 即可。对应可看下面的例子。
// n   = 1234321
// res = 1233999

// 但是由于减小了 str[i]str[i] 以后，可能不满足 str[i-1] <= str[i]str[i−1]<=str[i] 了，
// 所以我们在分析下这种情况怎么处理。我们看下这种情况的例子：
// n    = 2333332
// res  = 2299999



func monotoneIncreasingDigits(n int) int {
	numStr := []byte(strconv.Itoa(n))
	size := len(numStr)
	start := -1
	for i := size - 2; i >= 0; i -- {
		if numStr[i] > numStr[i + 1] {
			numStr[i] --
			start = i + 1
		}
	}
	if start != -1 {
		for i := start; i < size; i ++ {
			numStr[i] = '9'
		}
	}
	res, _ := strconv.Atoi(string(numStr))
	return res
}
