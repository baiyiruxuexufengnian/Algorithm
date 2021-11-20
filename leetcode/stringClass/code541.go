package stringClass

// 就是找到2k字符串，然后把这2k中的前k个反转了，如果找不到2k但是能找到大于等于k小于等于2k的话，
// 也把前k个反转了，实在是连k个都找不到，这不足k个字符全部反转了
func reverseStr(s string, k int) string {
	ss := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		if i + k <= length {
			reverse(ss[i:i+k])
		} else {
			reverse(ss[i:length])
		}
	}
	return string(ss)
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}