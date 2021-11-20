package stringClass

func replaceSpace(s string) string {
	spaceCount := 0
	bs := []byte(s)
	for _, v := range bs {
		if v == ' ' {
			spaceCount ++
		}
	}
	index1 := len(s) - 1
	index2 := index1 + 2 * spaceCount
	bs = append(bs, make([]byte, 2 * spaceCount)...)
	for index1 >= 0 {
		if bs[index1] == ' ' {
			bs[index2] = '0'
			bs[index2 - 1] = '2'
			bs[index2 - 2] = '%'
			index2 -= 3
			index1 --
		} else {
			bs[index2] = bs[index1]
			index1 --
			index2 --
		}
	}
	return string(bs)
}