package arrayList

var res [][]int
func combine(n int, k int) [][]int {
	if k <= 0 || k > n || n <= 0{
		return nil
	}
	res = make([][]int, 0)
	tmp := make([]int, 0)
	combineCore(1, n, k, tmp)
	return res
}

func combineCore(start, n int, k int, tmp []int) {
	if len(tmp) == k {
		res = append(res, tmp)
		return
	}

	for i := start; i <= n; i ++ {
		tmp = append(tmp, i)
		combineCore(i + 1, n, k, tmp)
		tmp = tmp[:len(tmp) - 1]
	}
}
