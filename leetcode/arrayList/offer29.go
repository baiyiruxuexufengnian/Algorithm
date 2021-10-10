package arrayList

// case: 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//       输出：[1,2,3,6,9,8,7,4,5]

func spiralOrder(matrix [][]int) []int {
	ret := make([]int, 0)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ret
	}
	col := 0
	row := 0
	count := len(matrix) * len(matrix[0])

	// col代表纵坐标：matrix[0]...matrix[3]...
	// row代表横坐标：matrix[col][0]...matrix[col][3]...
	for count > 0 {
		// 打印第一行
		for i := row; i <= len(matrix[0]) - row - 1 && count > 0; i ++ {
			ret = append(ret, matrix[col][i])
			count --
		}
		// 打印最后一列
		for i := col + 1; i <= len(matrix) - col - 1 && count > 0; i ++ {
			ret = append(ret, matrix[i][len(matrix[0]) - row - 1])
			count --
		}
		// 打印最后一行
		for i := len(matrix[0]) - row - 2; i >= row && count > 0; i -- {
			ret = append(ret, matrix[len(matrix) - col - 1][i])
			count --
		}
		// 打印第一列
		for i := len(matrix) - col - 2; i > col && count > 0; i -- {
			ret = append(ret, matrix[i][row])
			count --
		}
		if col + 1 < len(matrix) {
			col ++
		}
		if row + 1 < len(matrix[0]) {
			row ++
		}
	}
	return ret
}
