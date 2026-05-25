func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])

	var dfs func(int, int, int) bool
	dfs = func(r, c, index int) bool {
		if index == len(word) {
			return true
		}
		if r < 0 || r >= row || c < 0 || c >= col || board[r][c] != word[index] {
			return false
		}
		temp := board[r][c]
		board[r][c] = '#'
		found := dfs(r - 1, c, index + 1) || dfs(r, c - 1, index + 1) || 
					dfs(r + 1, c, index + 1) || dfs(r, c + 1, index + 1)
		board[r][c] = temp
		return found
	}

	for r := range row {
		for c := range col {
			if board[r][c] == word[0] && dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}
