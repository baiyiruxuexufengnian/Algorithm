package stringClass

func reverseLeftWords(s string, n int) string {
	length := len(s)
	if n >= length {
		return s
	}
	bs := []byte(s)
	reverseLeftWordsCore(bs, 0, n - 1)
	reverseLeftWordsCore(bs, n, length - 1)
	reverseLeftWordsCore(bs, 0, length - 1)
	return string(bs)
}

func reverseLeftWordsCore(in []byte, l, r int) {
	for l < r {
		in[l], in[r] = in[r], in[l]
		l ++
		r --
	}
}

func reverseLeftWords2(s string, n int) string {
	length := len(s)
	if n >= length {
		return s
	}
	bs := []byte(s)
	tmp := bs[:n]
	bs = bs[n:]
	bs = append(bs, tmp...)
	return string(bs)
}
