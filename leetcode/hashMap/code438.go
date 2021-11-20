package hashMap

func findAnagrams(s string, p string) []int {
	ret := make([]int, 0)
	if s == "" || p == "" || len(s) < len(p) {
		return ret
	}
	sCount := [26]byte{}
	pCount := [26]byte{}
	left := 0
	right := -1
	for i, c := range []byte(p) {
		sCount[s[i]-97] ++
		pCount[c-97] ++
		right ++
	}
	if sCount == pCount {
		ret = append(ret, left)
	}
	for right + 1 < len(s) {
		right ++
		sCount[s[right]-97] ++
		sCount[s[left]-97] --
		left ++
		if sCount == pCount {
			ret = append(ret, left)
		}
	}
	return ret
}