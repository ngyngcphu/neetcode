func numIslands(grid [][]byte) int {
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
		if grid[r][c] != '1' {
			return
		}
		visited[key] = true
		for _, d := range dirs {
			dfs(r + d[0], c + d[1])
		}
	}

	for r := range row {
		for c := range col {
			key := fmt.Sprintf("%d,%d", r, c)
			if grid[r][c] == '1' && !visited[key] {
				dfs(r, c)
				count++
			}
		}
	}
	return count
}
