package stringClass

func reverseString(s []byte)  {
	right := len(s) - 1
	left := 0
	for left < right {
		swap(s, left, right)
		left ++
		right --
	}

}

func swap(s []byte, i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}
