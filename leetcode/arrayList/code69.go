package arrayList

func mySqrt(x int) int {
	if x * x == x {
		return x
	}
	var l, r float64
	l = 0
	r = float64(x)
	for {
		pos := (r + l) / 2
		if float64(x) - pos * pos > -0.01 &&
			float64(x) - pos * pos < 0.01 {
			return int(pos)
		} else if pos * pos > float64(x) {
			r = pos
		} else {
			l = pos
		}
	}
}
