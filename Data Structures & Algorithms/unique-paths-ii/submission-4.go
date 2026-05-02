func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	memo := make(map[string]int)
	var helper func(m, n int) int
	helper = func(m, n int) int {
		if obstacleGrid[m][n] == 1 {
			return 0
		}
		if m == 0 && n == 0 {
			return 1
		}
		key := fmt.Sprintf("%d,%d", m, n)
		if val, ok := memo[key]; ok {
			return val
		}
		if m > 0 {
			memo[key] = helper(m - 1, n)
		}
		if n > 0 {
			memo[key] += helper(m, n - 1)
		}
		return memo[key]
	}
	return helper(len(obstacleGrid) - 1, len(obstacleGrid[0]) - 1)
}
