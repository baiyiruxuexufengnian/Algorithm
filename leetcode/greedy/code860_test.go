package greedy

func lemonadeChange(bills []int) bool {
	fiveCount, tenCount := 0, 0
	for _, v := range bills {
		if v == 5 {
			fiveCount ++
		}
		if v == 10 {
			if fiveCount <= 0 {
				return false
			}
			tenCount ++
			fiveCount --
		}
		if v == 20 {
			// 优先先把10元的消费掉，因为5元的相对来说更加万能一些
			if tenCount > 0 && fiveCount > 0 {
				tenCount --
				fiveCount --
				continue
			}
			if fiveCount < 3 {
				// 找不开钱了
				return false
			}
			fiveCount -= 3
		}
	}
	return true
}
