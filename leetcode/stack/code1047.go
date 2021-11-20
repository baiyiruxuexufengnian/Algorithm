package stack

import "strings"

func removeDuplicates(s string) string {
	stacks := make([]byte, 0)
	bs := []byte(s)
	for _, v := range bs {
		length := len(stacks)
		if length > 0 && stacks[length - 1] == v {
			stacks = stacks[:length - 1]

		} else {
			stacks = append(stacks, v)
		}
	}
	return string(stacks)
}

// 这个提不能用递归一方面当字符串很长的时候，递归用的系统栈很大，而且耗时也高
func removeDuplicates2(s string) string {
	ret := findSameStr(s)
	if ret == "" {
		return s
	}
	s = strings.ReplaceAll(s, ret, "")
	return removeDuplicates2(s)
}

func findSameStr(s string) string {
	bs := []byte(s)
	for i := 0; i < len(bs); i ++ {
		if i - 1 >= 0 && bs[i] == bs[i - 1] {
			return string([]byte{bs[i - 1], bs[i]})
		}
	}
	return ""
}
