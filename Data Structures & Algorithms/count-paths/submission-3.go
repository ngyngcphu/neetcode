func uniquePaths(m int, n int) int {
	result := 1
	k := min(m, n) - 1
	total := m + n - 2
	for i := 0; i < k; i++ {
		result = result * (total - i) / (i + 1)
	}
	return result
}
