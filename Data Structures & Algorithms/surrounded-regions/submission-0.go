func solve(board [][]byte) {
	rows, cols := len(board), len(board[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		if board[r][c] != 'O' {
			return
		}
		board[r][c] = 'S'
		for _, d := range directions {
			dfs(r + d[0], c + d[1])
		}
	}
	
	for r := range rows {
		if board[r][0] == 'O' {
			dfs(r, 0)
		}
		if board[r][cols - 1] == 'O' {
			dfs(r, cols - 1)
		}
	}
	for c := range cols {
		if board[0][c] == 'O' {
			dfs(0, c)
		}
		if board[rows - 1][c] == 'O' {
			dfs(rows - 1, c)
		}
	}

	for r := range rows {
		for c := range cols {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == 'S' {
				board[r][c] = 'O'
			}
		}
	}
}
