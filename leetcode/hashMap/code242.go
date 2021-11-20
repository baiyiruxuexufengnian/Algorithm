package hashMap

func isAnagram(s string, t string) bool {
	if s == "" || t == "" || len(s) < len(t) {
		return false
	}
	mp := make(map[byte]int)
	for _, str := range []byte(s) {
		mp[str] ++
	}

	for _, c := range []byte(t) {
		if _, ok := mp[c]; ok {
			mp[c] --
		}
		if mp[c] == 0 {
			delete(mp, c)
		}
	}
	if len(mp) == 0 {
		return true
	}
	return false
}
