package backtrack

func partition(s string) [][]string {
	results := make([][]string, 0)
	oneResult := make([]string, 0)
	length := len(s)
	sb := []byte(s)
	var partitionCore func(int)
	partitionCore = func(begin int) {
		if begin > length - 1 {
			tmp := make([]string, 0)
			tmp = append(tmp, oneResult ...)
			results = append(results, tmp)
			return
		}
		// 开始切分字串
		for i := begin; i < length; i ++ {
			if isPalindrome(sb, begin, i) {
				oneResult = append(oneResult, string(sb[begin:i+1]))
			} else {
				continue
			}
			partitionCore(i + 1)
			// greedy
			oneResult = oneResult[:len(oneResult) - 1]
		}
	}
	partitionCore(0)
	return results
}

func isPalindrome(in []byte, begin, end int) bool {
	for begin < end {
		if in[begin] != in[end] {
			return false
		}
		begin ++
		end --
	}
	return true
}
