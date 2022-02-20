package backtrack

func permute(nums []int) [][]int {
	results := make([][]int, 0)
	path := make([]int, 0)
	length := len(nums)
	used := make([]bool, length)
	var permuteGreedy func()
	permuteGreedy = func() {
		if len(path) == length {
			tmp := make([]int, 0)
			tmp = append(tmp, path ...)
			results = append(results, tmp)
			return
		}
		for i := 0; i < length; i ++ {
			if used[i] == true {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			permuteGreedy()
			// greedy
			used[i] = false
			path = path[:len(path) - 1]
		}
	}
	permuteGreedy()
	return results
}

func isValid(board [][]string, row, col int) (res bool){
	n := len(board)
	for i:=0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	for i := 0; i < n; i++{
		if board[row][i] == "Q" {
			return false
		}
	}

	for i ,j := row, col; i >= 0 && j >=0 ; i, j = i - 1, j- 1{
		if board[i][j] == "Q"{
			return false
		}
	}
	for i, j := row, col; i >=0 && j < n; i,j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
