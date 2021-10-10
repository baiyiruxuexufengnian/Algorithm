package arrayList

func generateMatrix(n int) [][]int {
	if n <= 0 {
		return nil
	}
	ret := make([][]int, n)
	for i := 0; i < n; i ++ {
		ret[i] = make([]int, n)
	}
	var v, j int
	v = 1
	// arr[纵坐标][横坐标],哪个坐标不变哪个坐标就应该有j
	for v <= n * n {
		for i := j; i < n - j; i ++ {
			ret[j][i] = v
			v ++
		}
		for i := j + 1; i < n - j; i ++ {
			ret[i][n - j - 1] = v
			v ++
		}
		for i := n - j - 2; i >= j; i -- {
			ret[n - j - 1][i] = v
			v ++
		}
		for i := n - j - 2; i > j; i -- {
			ret[i][j] = v
			v ++
		}
		j ++
	}
	return ret
}
