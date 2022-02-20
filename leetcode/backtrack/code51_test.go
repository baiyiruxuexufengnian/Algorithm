package backtrack

import "testing"

func solveNQueens(n int) [][]string {
	results := make([][]string, 0)
	path := make([][]byte, n)
	for i := 0; i < n; i ++ {
		path[i] = make([]byte, n)
	}
	for i := 0; i < n; i ++ {
		for j := 0; j < n; j ++ {
			path[i][j] = '.'
		}
	}
	var solveNQueensBacktrack func(int, int)
	solveNQueensBacktrack = func(row, col int) {
		if row == n{
			tmp := make([]string, 0)
			for _, v := range path {
				tmp = append(tmp, string(v))
			}
			results = append(results, tmp)
			return
		}
		for i := 0; i < col; i ++ {
			if !isValidQueens(row, i, path) {
				continue
			}
			path[row][i] = 'Q'
			solveNQueensBacktrack(row + 1, col)
			path[row][i] = '.'
		}
	}
	solveNQueensBacktrack(0, n)
	return results
}

func isValidQueens(row, col int, path [][]byte) bool {
	n := len(path)
	// check col
	for i := 0; i < row; i ++ {
		if path[i][col] == 'Q' {
			return false
		}
	}
	// check row
	for i := 0; i < n; i ++ {
		if path[row][i] == 'Q' {
			return false
		}
	}
	// check 45'
	for i, j := row, col; i >= 0 && j >= 0; i, j = i - 1, j - 1{
		if path[i][j] == 'Q' {
			return false
		}
	}
	// check 135'
	for i, j := row, col; i >= 0 && j < n; i, j = i - 1, j + 1{
		if path[i][j] == 'Q' {
			return false
		}
	}
	return true
}



func TestSolveNQueens(t *testing.T) {
	results := solveNQueens(4)
	t.Logf("count: %v, results: %v", len(results), results)
}
