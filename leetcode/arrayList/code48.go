package arrayList

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	chash := make(map[byte]int)
	str := []byte(s)
	begin := 0
	end := 0
	chash[str[end]] = end
	maxWin := math.MinInt32
	for end < len(str) {
		if end - begin + 1 > maxWin {
			maxWin = end - begin + 1
		}
		if end + 1 >= len(str) {
			break
		}
		if pos, ok := chash[str[end + 1]]; ok {
			if chash[str[pos]] + 1 >= begin {
				begin = chash[str[pos]] + 1
			}
			end ++
			chash[str[pos]] = end
		} else {
			end ++
			chash[str[end]] = end
		}
	}
	if maxWin == math.MinInt32 {
		maxWin = 0
	}
	return maxWin
}
