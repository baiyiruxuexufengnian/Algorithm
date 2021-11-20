package hashMap

func getSum(n int) int {
	sum := 0
	for n != 0 {
		remainder := n % 10
		squares := remainder * remainder
		sum += squares
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	if n < 0 { return false }
	mp := make(map[int]bool, 0)
	sum := n
	for {
		sum = getSum(sum)
		if _, ok := mp[sum]; ok {
			return false
		}
		if sum == 1 {
			return true
		}
		mp[sum] = true
	}
}
