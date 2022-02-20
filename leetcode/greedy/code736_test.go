package greedy

func partitionLabels(s string) []int {
	strMapMaxIndex := make(map[byte]int, 0)
	sbt := []byte(s)
	for i, ch := range sbt {
		strMapMaxIndex[ch] = i
	}
	preIndex := -1
	maxIndex := 0
	results := make([]int, 0)
	var maxInt func(int, int) int
	maxInt = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	for index, val := range sbt {
		maxIndex = maxInt(strMapMaxIndex[val], maxIndex)
		if index == maxIndex {
			results = append(results, index - preIndex)
			preIndex = index
		}
	}
	return results
}
