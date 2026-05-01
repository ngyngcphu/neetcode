func uniquePaths(m int, n int) int {
	memo := make(map[string]int)
	var helper func(row, col int) int
	helper = func(row, col int) int {
		if row == 0 || col == 0 {
			return 1
		}
		key := fmt.Sprintf("%d,%d", row, col)
		if val, ok := memo[key]; ok {
			return val
		}
		memo[key] = helper(row - 1, col) + helper(row, col - 1)
		return memo[key]
	}
	return helper(m - 1, n - 1)
}
