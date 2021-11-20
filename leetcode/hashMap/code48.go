package hashMap

func rotate(matrix [][]int)  {
	if len(matrix) == 0 || len(matrix[0]) == 0 || len(matrix) != len(matrix[0]) {
		return
	}
	// 沿左上角到右下角的对角线对折
	n := len(matrix)
	for i := 0; i < n; i ++ {
		for j := i; j < n; j ++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}

	// 左右对折
	for i := 0; i < n; i ++ {
		for j := 0; j < n / 2; j ++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][n - j - 1]
			matrix[i][n - j - 1] = tmp
		}
	}
}

