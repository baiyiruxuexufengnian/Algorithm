package arrayList

func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}

	if num == num * num {
		return true
	}
	l := 0
	r := num
	for l <= r {
		pos := l + (r - l) / 2
		if pos * pos == num {
			return true
		} else if pos * pos > num {
			r = pos - 1
		} else {
			l = pos + 1
		}
	}
	return false
}
