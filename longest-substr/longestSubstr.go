package longestSubstr

import (
	"math"
)

/*
	Leetcode: 3
	输入: s = "pwwkew"
	输出: 3
	解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。

	输入: s = ""
	输出: 0

	输入: s = "dvdf"
	输出: 3
*/

func lengthOfLongestSubstring(s string) int {
	ss := []byte(s)
	begin := 0
	end := -1
	maxWin := 0
	slideWin := make(map[byte]bool)
	for end < len(ss) {
		if end < begin {
			end ++
			continue
		}
		if _, ok := slideWin[ss[end]]; !ok {
			slideWin[ss[end]] = true
			if len(slideWin) > maxWin {
				maxWin = len(slideWin)
			}
			end ++
		} else {
			delete(slideWin, ss[begin])
			begin ++
		}
	}
	return maxWin
}

func lengthOfLongestSubstring2(s string) int {
	ss := []byte(s)
	begin := 0
	end := 0
	maxWin := 0
	slideWin := make(map[byte]int)
	for end < len(ss) {
		begin = int(math.Max(float64(begin), float64(slideWin[ss[end]])))
		slideWin[ss[end]] = end + 1
		maxWin = int(math.Max(float64(maxWin), float64(end - begin + 1)))
		end ++
	}
	return maxWin
}


func lengthOfLongestSubstring1(s string) int {
	ss := []byte(s)
	begin := 0
	end := 0
	res := 0
	slideWin := make([]int, 128)
	for ; end < len(ss); end ++ {
		begin = int(math.Max(float64(begin), float64(slideWin[ss[end]])))
		slideWin[ss[end]] = end + 1
		res = int(math.Max(float64(res), float64(end - begin + 1)))
	}
	return res
}

func Solution(s string) int {
	return lengthOfLongestSubstring(s)
}
