func numEnclaves(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	visited := make(map[string]bool)
	count := 0
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(r, c int) 
	dfs = func(r, c int) {
		key := fmt.Sprintf("%d,%d", r, c)
		if visited[key] {
			return
		}
		if r < 0 || r >= row || c < 0 || c >= col {
			return
		}
		if grid[r][c] != 1 {
			return
		}
		visited[key] = true
		for _, d := range dirs {
			dfs(r + d[0], c + d[1])
		}
	}

	for r := range row {
		if grid[r][0] == 1 {
			dfs(r, 0)
		}
		if grid[r][col - 1] == 1 {
			dfs(r, col - 1)
		}
	}

	for c := range col {
		if grid[0][c] == 1 {
			dfs(0, c)
		}
		if grid[row - 1][c] == 1 {
			dfs(row - 1, c)
		}
	}

	for r := range row {
		for c := range col {
			key := fmt.Sprintf("%d,%d", r, c)
			if grid[r][c] == 1 && !visited[key] {
				count++
			}
		}
	}
	return count
}
