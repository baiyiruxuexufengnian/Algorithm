package stringClass

func reverseWords(s string) string {
	if s == "" { return s }
	bs := []byte(s)
	length := len(s)
	// 删除头部冗余空格
	fastIndex, slowIndex := 0, 0
	for _, v := range bs {
		if v == ' ' {
			fastIndex ++
			continue
		} else {
			break
		}
	}

	// 删除中间部分冗余空格
	for fastIndex < length {
		if bs[fastIndex] == ' ' && fastIndex > 0 && bs[fastIndex] == bs[fastIndex - 1] {
			fastIndex ++
			continue
		}
		bs[slowIndex] = bs[fastIndex]
		slowIndex ++
		fastIndex ++
	}
	// 删除末尾冗余空格
	if slowIndex - 1 > 0 && bs[slowIndex - 1] == ' '{
		bs = bs[:slowIndex - 1]
	} else {
		bs = bs[:slowIndex]
	}
	l := 0
	r := len(bs) - 1
	reverseCore(bs, l, r)
	for i, v := range bs {
		if v == ' ' {
			reverseCore(bs, l, i - 1)
			l = i + 1
		}
	}
	reverseCore(bs, l, r)
	return string(bs)
}

func reverseCore(bs []byte, l, r int) {
	for l < r {
		swap(bs, l, r)
		l ++
		r --
	}
}
